package config

type RestAPIConfig struct {
	Endpoint                 string
	Port                     string
	UserServiceEndpoint      string
	ContainerServiceEndpoint string
}

type ContainerServiceConfig struct {
	UserServiceEndpoint string
}
