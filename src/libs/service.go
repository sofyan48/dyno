package libs

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/sofyan48/dyno/src/libs/entity"
)

// Client ...
var Client *api.Client

// RegisterConsul register service
func (svc *Service) RegisterConsul(client *api.Client, regis entity.ServiceRegister) {
	svc.registerConsul(client, regis)
}

// GetAgentServiceConsul get agent
func (svc *Service) GetAgentServiceConsul() *api.AgentServiceRegistration {
	return new(api.AgentServiceRegistration)
}

func (svc *Service) registerConsul(client *api.Client, regis entity.ServiceRegister) error {
	registration := svc.GetAgentServiceConsul()
	registration.ID = regis.ID
	registration.Name = regis.Name
	registration.Address = regis.Host
	prt, err := strconv.Atoi(regis.Port[1:len(regis.Port)])
	if err != nil {
		return err
	}
	registration.Port = prt
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/%s", regis.Host, prt, regis.HealthCheck)
	registration.Check.Interval = regis.CheckInterval
	registration.Check.Timeout = regis.CheckTimeout
	client.Agent().ServiceRegister(registration)
	return nil
}
