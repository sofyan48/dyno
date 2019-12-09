package config

import (
	"github.com/hashicorp/consul/api"
)

// New config to consul
func (cn *Consul) New(config *api.Config) (*api.Client, error) {
	consul := &api.Client{}
	consul, err := api.NewClient(config)
	if err != nil {
		return consul, err
	}
	return consul, nil

}

// Config ...
func (cn *Consul) Config() *api.Config {
	cfg := &api.Config{}
	return cfg
}
