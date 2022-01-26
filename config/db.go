package config

type DB struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Database    string `yaml:"database"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Charset     string `yaml:"charset"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	EnableLog   bool   `yaml:"enableLog"`
	LogLevel    string `yaml:"logLevel"`
	LogFilename string `yaml:"logFilename"`
}
