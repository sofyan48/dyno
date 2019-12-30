package entity

// ServiceRegister ...
type ServiceRegister struct {
	Host        string   `json:"host" yaml:"host"`
	Port        string   `json:"port" yaml:"port"`
	HealthCheck string   `json:"health_check" yaml:"health_check"`
	ID          string   `json:"id" yaml:"id"`
	Name        string   `json:"name"  yaml:"name"`
	Interval    string   `json:"interval" yaml:"interval"`
	Timeout     string   `json:"timeout" yaml:"timeout"`
	Tags        []string `json:"tags" yaml:"tags"`
}

// ServiceRegisterYML ...
type ServiceRegisterYML struct {
	Service struct {
		ID   string   `json:"id" yaml:"id"`
		Name string   `json:"name"  yaml:"name"`
		Host string   `json:"host" yaml:"host"`
		Port string   `json:"port" yaml:"port"`
		Tags []string `json:"tags" yaml:"tags"`
	} `json:"service" yaml:"service"`

	HealthCheck struct {
		Endpoint string `json:"path" yaml:"path"`
		Interval string `json:"interval" yaml:"interval"`
		Timeout  string `json:"timeout" yaml:"timeout"`
	} `json:"healthcheck" yaml:"healthcheck"`
}
