package config

type Log struct {
	Level     string `yaml:"level"`
	RootDir   string `yaml:"rootDir"`
	Filename  string `yaml:"filename"`
	ShowLine  bool   `yaml:"showLine"`
	Compress  bool   `yaml:"compress"`
	Format    string `yaml:"format"`
	MaxSize   int    `yaml:"maxSize"`
	MaxAge    int    `yaml:"maxAge"`
	MaxBackup int    `yaml:"maxBackup"`
}
