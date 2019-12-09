package entity

// ServiceRegister ...
type ServiceRegister struct {
	// var host, port, healthcheck string
	Host          string `json:"host" yaml:"host"`
	Port          string `json:"port" yaml:"port"`
	HealthCheck   string `json:"health_check" yaml:"health_check"`
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name"  yaml:"name"`
	CheckInterval string `json:"interval" yaml:"interval"`
	CheckTimeout  string `json:"timeout" yaml:"timeout"`
}
