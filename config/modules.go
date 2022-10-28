package config

import (
	"github.com/topfreegames/pitaya/v2/config"
	"strings"
	"time"
)

type Modules struct {
	SessionUnique                        bool   `json:"session-unique" yaml:"session-unique"`
	ModulesBindingStorageEtcdEndpoints   string `json:"modules-binding-storage-etcd-endpoints" yaml:"modules-binding-storage-etcd-endpoints"`
	ModulesBindingStorageEtcdPrefix      string `json:"modules-binding-storage-etcd-prefix" yaml:"modules-binding-storage-etcd-prefix"`
	ModulesBindingStorageEtcdDialTimeout int    `json:"modules-binding-storage-etcd-dial-timeout" yaml:"modules-binding-storage-etcd-dial-timeout"`
	ModulesBindingStorageEtcdLeaseTtl    int    `json:"modules-binding-storage-etcd-lease-ttl" yaml:"modules-binding-storage-etcd-lease-ttl"`
}

func (m *Modules) ETCDBindingConfig() *config.ETCDBindingConfig {
	return &config.ETCDBindingConfig{
		DialTimeout: time.Duration(m.ModulesBindingStorageEtcdDialTimeout) * time.Second,
		Endpoints:   strings.Split(m.ModulesBindingStorageEtcdEndpoints, ","),
		Prefix:      m.ModulesBindingStorageEtcdPrefix,
		LeaseTTL:    time.Duration(m.ModulesBindingStorageEtcdLeaseTtl) * time.Hour,
	}
}
