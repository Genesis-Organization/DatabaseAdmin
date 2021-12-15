package config

type ConfigInterface struct {
	Port string
}

var Config = ConfigInterface{
	Port: ":3337",
}
