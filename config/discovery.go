package config

import (
	"github.com/topfreegames/pitaya/v2/config"
	"strings"
	"time"
)

type Discovery struct {
	ClusterSdEtcdDialTimeout             int      `json:"cluster-sd-etcd-dial-timeout" mapstructure:"cluster-sd-etcd-dial-timeout"`
	ClusterSdEtcdEndpoints               string   `json:"cluster-sd-etcd-endpoints" mapstructure:"cluster-sd-etcd-endpoints"`
	ClusterSdEtcdUser                    string   `json:"cluster-sd-etcd-user" mapstructure:"cluster-sd-etcd-user"`
	ClusterSdEtcdPass                    string   `json:"cluster-sd-etcd-pass" mapstructure:"cluster-sd-etcd-pass"`
	ClusterSdEtcdHeartbeatTtl            int      `json:"cluster-sd-etcd-heartbeat-ttl" mapstructure:"cluster-sd-etcd-heartbeat-ttl"`
	ClusterSdEtcdGrantLeaseTimeout       int      `json:"cluster-sd-etcd-grant-lease-timeout" mapstructure:"cluster-sd-etcd-grant-lease-timeout"`
	ClusterSdEtcdGrantLeaseMaxRetries    int      `json:"cluster-sd-etcd-grant-lease-max-retries" mapstructure:"cluster-sd-etcd-grant-lease-max-retries"`
	ClusterSdEtcdGrantLeaseRetryInterval int      `json:"cluster-sd-etcd-grant-lease-retry-interval" mapstructure:"cluster-sd-etcd-grant-lease-retry-interval"`
	ClusterSdEtcdRevokeTimeout           int      `json:"cluster-sd-etcd-revoke-timeout" mapstructure:"cluster-sd-etcd-revoke-timeout"`
	ClusterSdEtcdHeartbeatLog            bool     `json:"cluster-sd-etcd-heartbeat-log" mapstructure:"cluster-sd-etcd-heartbeat-log"`
	ClusterSdEtcdPrefix                  string   `json:"cluster-sd-etcd-prefix" mapstructure:"cluster-sd-etcd-prefix"`
	ClusterSdEtcdSyncServersInterval     int      `json:"cluster-sd-etcd-sync-servers-interval" mapstructure:"cluster-sd-etcd-sync-servers-interval"`
	ClusterSdEtcdShutdownDelay           int      `json:"cluster-sd-etcd-shutdown-delay" mapstructure:"cluster-sd-etcd-shutdown-delay"`
	ClusterSdEtcdServerTypeBlacklist     []string `json:"cluster-sd-etcd-service-http-admin-type-blacklist" mapstructure:"cluster-sd-etcd-service-http-admin-type-blacklist"`
	ClusterSdEtcdSyncServersParallelism  int      `json:"cluster-sd-etcd-sync-servers-parallelism" mapstructure:"cluster-sd-etcd-sync-servers-parallelism"`
}

func (d *Discovery) EtcdDiscoveryConfig() *config.EtcdServiceDiscoveryConfig {
	return &config.EtcdServiceDiscoveryConfig{
		Endpoints:   strings.Split(d.ClusterSdEtcdEndpoints, ","),
		User:        d.ClusterSdEtcdUser,
		Pass:        d.ClusterSdEtcdPass,
		DialTimeout: time.Duration(d.ClusterSdEtcdDialTimeout) * time.Second,
		Prefix:      d.ClusterSdEtcdPrefix,
		Heartbeat: struct {
			TTL time.Duration
			Log bool
		}{
			TTL: time.Duration(d.ClusterSdEtcdHeartbeatTtl) * time.Second,
			Log: d.ClusterSdEtcdHeartbeatLog,
		},
		SyncServers: struct {
			Interval    time.Duration
			Parallelism int
		}{
			Interval:    time.Duration(d.ClusterSdEtcdSyncServersInterval) * time.Second,
			Parallelism: d.ClusterSdEtcdSyncServersParallelism,
		},
		Revoke: struct {
			Timeout time.Duration
		}{
			Timeout: time.Duration(d.ClusterSdEtcdRevokeTimeout) * time.Second,
		},
		GrantLease: struct {
			Timeout       time.Duration
			MaxRetries    int
			RetryInterval time.Duration
		}{
			Timeout:       time.Duration(d.ClusterSdEtcdGrantLeaseTimeout) * time.Second,
			MaxRetries:    d.ClusterSdEtcdGrantLeaseMaxRetries,
			RetryInterval: time.Duration(d.ClusterSdEtcdGrantLeaseRetryInterval) * time.Second,
		},
		Shutdown: struct {
			Delay time.Duration
		}{
			Delay: time.Duration(d.ClusterSdEtcdShutdownDelay) * time.Second,
		},
		ServerTypesBlacklist: d.ClusterSdEtcdServerTypeBlacklist,
	}
}
