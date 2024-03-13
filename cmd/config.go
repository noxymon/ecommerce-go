package cmd

type Config struct {
	Server struct {
		Port    int  `yaml:"port"`
		PreFork bool `yaml:"pre-fork"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DbClass  string `yaml:"db-class"`
		Username string `yaml:"username"`
		Password int    `yaml:"password"`
	} `yaml:"database"`
}
