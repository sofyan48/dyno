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
func (svc *Service) RegisterConsul(client *api.Client, regis entity.ServiceRegister) error {
	return svc.registerConsul(client, regis)
}

// GetAgentServiceConsul get agent
func (svc *Service) GetAgentServiceConsul() *api.AgentServiceRegistration {
	return new(api.AgentServiceRegistration)
}

// GetCheckAgentService get agent
func (svc *Service) GetCheckAgentService() *api.AgentCheckRegistration {
	return new(api.AgentCheckRegistration)
}

// CheckServiceConsul check service at consul
func (svc *Service) CheckServiceConsul(client *api.Client, regis entity.ServiceRegister) error {
	registran := svc.GetCheckAgentService()
	registran.ID = regis.ID
	registran.Name = regis.Name
	return client.Agent().CheckRegister(registran)
}

// UnRegisterConsul uregister consul
func (svc *Service) UnRegisterConsul(client *api.Client, ID string) error {
	return client.Agent().ServiceDeregister(ID)
}

// ListServiceConsul list all service
func (svc *Service) ListServiceConsul(client *api.Client) (map[string]*api.AgentService, error) {
	return client.Agent().Services()
}

// ServiceLookupConsul get Lookup
func (svc *Service) ServiceLookupConsul(client *api.Client, ID string) (string, error) {
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[ID]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

// GetHealthByIDConsul ...
func (svc *Service) GetHealthByIDConsul(client *api.Client, ID string) (string, *api.AgentServiceChecksInfo, error) {
	return client.Agent().AgentHealthServiceByID(ID)
}

func (svc *Service) registerConsul(client *api.Client, regis entity.ServiceRegister) error {
	registration := svc.GetAgentServiceConsul()
	registration.ID = regis.ID
	registration.Name = regis.Name
	registration.Address = regis.Host
	prt, err := strconv.Atoi(regis.Port[0:len(regis.Port)])
	if err != nil {
		return err
	}
	registration.Port = prt
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/%s", regis.Host, prt, regis.HealthCheck)
	registration.Check.Interval = regis.Interval
	registration.Check.Timeout = regis.Timeout
	registration.Tags = regis.Tags
	return client.Agent().ServiceRegister(registration)
}
