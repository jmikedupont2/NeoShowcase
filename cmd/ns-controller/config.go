package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/spf13/viper"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/domain/web"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/k8simpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb/pbconnect"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/repository"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/storage"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/webhook"
)

const (
	ModeDocker = iota
	ModeK8s
)

type Config struct {
	Port           int                    `mapstructure:"port" yaml:"port"`
	Debug          bool                   `mapstructure:"debug" yaml:"debug"`
	PrivateKeyFile string                 `mapstructure:"privateKeyFile" yaml:"privateKeyFile"`
	AdminerURL     domain.AdminerURL      `mapstructure:"adminerURL" yaml:"adminerURL"`
	Mode           string                 `mapstructure:"mode" yaml:"mode"`
	Docker         dockerimpl.Config      `mapstructure:"docker" yaml:"docker"`
	K8s            k8simpl.Config         `mapstructure:"k8s" yaml:"k8s"`
	SSH            domain.SSHConfig       `mapstructure:"ssh" yaml:"ssh"`
	Webhook        webhook.ReceiverConfig `mapstructure:"webhook" yaml:"webhook"`
	DB             repository.Config      `mapstructure:"db" yaml:"db"`
	Storage        domain.StorageConfig   `mapstructure:"storage" yaml:"storage"`
	Image          builder.ImageConfig    `mapstructure:"image" yaml:"image"`
}

func init() {
	viper.SetDefault("port", 10000)
	viper.SetDefault("debug", false)
	viper.SetDefault("privateKeyFile", "")
	viper.SetDefault("adminerURL", "http://adminer.local.trapti.tech/")
	viper.SetDefault("mode", "docker")

	viper.SetDefault("docker.confDir", "/opt/traefik/conf")
	viper.SetDefault("docker.domains", nil)
	viper.SetDefault("docker.ports", nil)
	viper.SetDefault("docker.ss.url", "")
	viper.SetDefault("docker.network", "neoshowcase_apps")
	viper.SetDefault("docker.labels", nil)
	viper.SetDefault("docker.tls.certResolver", "nsresolver")
	viper.SetDefault("docker.tls.wildcard.domains", nil)
	viper.SetDefault("docker.resources.cpus", 1.6)
	viper.SetDefault("docker.resources.memory", 1e9 /* 1GB */)
	viper.SetDefault("docker.resources.memorySwap", -1 /* unlimited swap */)
	viper.SetDefault("docker.resources.memoryReservation", 256*1e6 /* 256MB */)

	viper.SetDefault("k8s.domains", nil)
	viper.SetDefault("k8s.ports", nil)
	viper.SetDefault("k8s.ss.namespace", "default")
	viper.SetDefault("k8s.ss.kind", "Service")
	viper.SetDefault("k8s.ss.name", "")
	viper.SetDefault("k8s.ss.port", 80)
	viper.SetDefault("k8s.ss.scheme", "http")
	viper.SetDefault("k8s.namespace", "neoshowcase-apps")
	viper.SetDefault("k8s.labels", nil)
	viper.SetDefault("k8s.tls.type", "traefik")
	viper.SetDefault("k8s.tls.traefik.certResolver", "nsresolver")
	viper.SetDefault("k8s.tls.traefik.wildcard.domains", nil)
	viper.SetDefault("k8s.tls.certManager.issuer.name", "cert-issuer")
	viper.SetDefault("k8s.tls.certManager.issuer.kind", "ClusterIssuer")
	viper.SetDefault("k8s.tls.certManager.wildcard.domains", nil)
	viper.SetDefault("k8s.imagePullSecret", "")
	viper.SetDefault("k8s.scheduling.nodeSelector", nil)
	viper.SetDefault("k8s.scheduling.tolerations", nil)
	viper.SetDefault("k8s.resources.requests.cpu", "")
	viper.SetDefault("k8s.resources.requests.memory", "")
	viper.SetDefault("k8s.resources.limits.cpu", "")
	viper.SetDefault("k8s.resources.limits.memory", "")

	viper.SetDefault("ssh.host", "localhost")
	viper.SetDefault("ssh.port", 2201)

	viper.SetDefault("webhook.basePath", "/api/webhook")
	viper.SetDefault("webhook.port", 8080)

	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.database", "neoshowcase")
	viper.SetDefault("db.connection.maxOpen", 0)
	viper.SetDefault("db.connection.maxIdle", 2)
	viper.SetDefault("db.connection.lifetime", 0)

	viper.SetDefault("storage.type", "local")
	viper.SetDefault("storage.local.dir", "/neoshowcase")
	viper.SetDefault("storage.s3.bucket", "neoshowcase")
	viper.SetDefault("storage.s3.endpoint", "")
	viper.SetDefault("storage.s3.region", "")
	viper.SetDefault("storage.s3.accessKey", "")
	viper.SetDefault("storage.s3.accessSecret", "")
	viper.SetDefault("storage.swift.username", "")
	viper.SetDefault("storage.swift.apiKey", "")
	viper.SetDefault("storage.swift.tenantName", "")
	viper.SetDefault("storage.swift.tenantId", "")
	viper.SetDefault("storage.swift.container", "neoshowcase")
	viper.SetDefault("storage.swift.authUrl", "")

	viper.SetDefault("image.registry.scheme", "https")
	viper.SetDefault("image.registry.addr", "localhost")
	viper.SetDefault("image.registry.username", "")
	viper.SetDefault("image.registry.password", "")
	viper.SetDefault("image.namePrefix", "ns-apps/")
	viper.SetDefault("image.tmpNamePrefix", "ns-apps-tmp/")
}

func (c *Config) GetMode() int {
	switch strings.ToLower(c.Mode) {
	case "k8s", "kubernetes":
		return ModeK8s
	case "docker":
		return ModeDocker
	default:
		return ModeDocker
	}
}

func providePublicKey(c Config) (*ssh.PublicKeys, error) {
	bytes, err := os.ReadFile(c.PrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open private key file")
	}
	return domain.NewPublicKey(bytes)
}

func provideStorage(c domain.StorageConfig) (domain.Storage, error) {
	switch strings.ToLower(c.Type) {
	case "local":
		return storage.NewLocalStorage(c.Local.Dir)
	case "s3":
		return storage.NewS3Storage(c.S3.Bucket, c.S3.AccessKey, c.S3.AccessSecret, c.S3.Region, c.S3.Endpoint)
	case "swift":
		return storage.NewSwiftStorage(c.Swift.Container, c.Swift.UserName, c.Swift.APIKey, c.Swift.TenantName, c.Swift.TenantID, c.Swift.AuthURL)
	default:
		return nil, fmt.Errorf("unknown storage: %s", c.Type)
	}
}

type controllerServer struct {
	*web.H2CServer
}

func provideControllerServer(
	c Config,
	controller pbconnect.ControllerServiceHandler,
	builderHandler domain.ControllerBuilderService,
	ssgenHandler domain.ControllerSSGenService,
) *controllerServer {
	wc := web.H2CConfig{
		Port: c.Port,
		SetupRoute: func(mux *http.ServeMux) {
			mux.Handle(pbconnect.NewControllerServiceHandler(controller))
			mux.Handle(pbconnect.NewControllerBuilderServiceHandler(builderHandler))
			mux.Handle(pbconnect.NewControllerSSGenServiceHandler(ssgenHandler))
		},
	}
	return &controllerServer{web.NewH2CServer(wc)}
}
