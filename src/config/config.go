package config

// Etcd ...
type Etcd struct{}

// Consul ...
type Consul struct{}

type Config struct {
	CosulConfig Consul
	EtcdConfig  Etcd
}
