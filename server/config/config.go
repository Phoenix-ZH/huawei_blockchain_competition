package config

import "github.com/jinzhu/configor"

type JWT struct {
	Secret string `required:"true"`
}
type Database struct {
	Dialect  string `default:"mongodb"`
	Debug    bool   `default:"false"`
	Username string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Port     int
	SSLMode bool
}
type Configuration struct {
	Database Database `required:"true"`
	JWT      JWT      `required:"true"`
}

var Main = (func() Configuration {
	var conf Configuration
	if err := configor.Load(&conf, "PATH_TO_CONFIG_FILE"); err != nil {
		panic(err.Error())
	}
	return conf
})()

