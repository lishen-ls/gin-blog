package config

type App struct {
	Env  string `json:"env" yaml:"env"`
	Port string `json:"port" yaml:"port"`
}
