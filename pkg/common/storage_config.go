package common

import (
	"fmt"
	"github.com/traPtitech/neoshowcase/pkg/storage"
	"strings"
)

type StorageConfig struct {
	Type  string `mapstructure:"type" yaml:"type"`
	Local struct {
		// Dir 保存ディレクトリ
		Dir string `mapstructure:"dir" yaml:"dir"`
	} `mapstructure:"local" yaml:"local"`
	S3 struct {
		// Bucket バケット名
		Bucket string `mapstructure:"bucket" yaml:"bucket"`
	} `mapstructure:"s3" yaml:"s3"`
	Swift struct {
		// UserName ユーザー名
		UserName string `mapstructure:"username" yaml:"username"`
		// APIKey APIキー(パスワード)
		APIKey string `mapstructure:"apiKey" yaml:"apiKey"`
		// TenantName テナント名
		TenantName string `mapstructure:"tenantName" yaml:"tenantName"`
		// TenantID テナントID
		TenantID string `mapstructure:"tenantId" yaml:"tenantId"`
		// Container コンテナ名
		Container string `mapstructure:"container" yaml:"container"`
		// AuthURL 認証エンドポイント
		AuthURL string `mapstructure:"authUrl" yaml:"authUrl"`
	} `mapstructure:"swift" yaml:"swift"`
}

func (c *StorageConfig) InitStorage() (storage.Storage, error) {
	switch strings.ToLower(c.Type) {
	case "local":
		return storage.NewLocalStorage(c.Local.Dir)
	case "s3":
		return storage.NewS3Storage(c.S3.Bucket)
	case "swift":
		return storage.NewSwiftStorage(c.Swift.Container, c.Swift.UserName, c.Swift.APIKey, c.Swift.TenantName, c.Swift.TenantID, c.Swift.AuthURL)
	default:
		return nil, fmt.Errorf("unknown storage: %s", c.Type)
	}
}
