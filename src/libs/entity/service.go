package entity

// ServiceRegister ...
type ServiceRegister struct {
	// var host, port, healthcheck string
	Host          string `json:"host"`
	Port          string `json:"port"`
	HealthCheck   string `json:"health_check"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	CheckInterval string `json:"interval"`
	CheckTimeout  string `json:"timeout"`
}
