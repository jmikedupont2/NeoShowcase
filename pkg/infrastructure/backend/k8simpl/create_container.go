package k8simpl

import (
	"context"
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/util"
)

func (b *k8sBackend) CreateContainer(ctx context.Context, app *domain.Application, args domain.ContainerCreateArgs) error {
	if args.ImageTag == "" {
		args.ImageTag = "latest"
	}

	podSelector := map[string]string{
		appContainerLabel:              "true",
		appContainerApplicationIDLabel: app.ID,
	}
	labels := util.MergeLabels(args.Labels, podSelector)

	var envs []apiv1.EnvVar

	for name, value := range args.Envs {
		envs = append(envs, apiv1.EnvVar{Name: name, Value: value})
	}

	cont := apiv1.Container{
		Name:  "app",
		Image: args.ImageName + ":" + args.ImageTag,
		Env:   envs,
	}
	for _, website := range app.Websites {
		err := b.registerService(ctx, app, website, podSelector)
		if err != nil {
			return err
		}
		err = b.registerIngress(ctx, app, website)
		if err != nil {
			return err
		}
	}

	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName(app.ID),
			Namespace: appNamespace,
			Labels:    labels,
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{cont},
		},
	}

	_, err := b.client.CoreV1().Pods(appNamespace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		if args.Recreate && errors.IsAlreadyExists(err) {
			// TODO いい感じにしたい
			if err = b.client.CoreV1().Pods(appNamespace).Delete(ctx, pod.Name, metav1.DeleteOptions{}); err != nil {
				return fmt.Errorf("failed to delete pod: %w", err)
			}
			for {
				_, err := b.client.CoreV1().Pods(appNamespace).Get(ctx, pod.Name, metav1.GetOptions{})
				if err != nil {
					if errors.IsNotFound(err) {
						break
					}
					return fmt.Errorf("failed to delete pod: %w", err)
				}
				time.Sleep(500 * time.Millisecond)
			}
			if _, err := b.client.CoreV1().Pods(appNamespace).Create(ctx, pod, metav1.CreateOptions{}); err != nil {
				return fmt.Errorf("failed to create pod: %w", err)
			}
		} else {
			return fmt.Errorf("failed to create pod: %w", err)
		}
	}

	return nil

}
