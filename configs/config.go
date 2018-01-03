package configs

import (
	"os"
	"path"

	"github.com/aaa59891/go_empty_gin/utils"
	"github.com/spf13/viper"
)

type config struct {
	Server   server
	Database database
}

type server struct {
	Port          string
	StaticPath    string
	StaticUriRoot string
}

type database struct {
	DriveName    string
	Account      string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

var c *config = &config{}

func init() {
	mosiGoEnv := os.Getenv("GO_ENV")
	var dir string
	if mosiGoEnv == "test" {
		dir = "../"
	} else {
		dir = utils.GetProjectRoot()
	}

	viper.SetConfigName("default")
	viper.AddConfigPath(path.Join(dir, "configs"))

	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	if len(mosiGoEnv) != 0 {
		viper.SetConfigName(mosiGoEnv)
	}

	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(c); err != nil {
		panic(err)
	}
}

func GetConfig() config {
	return *c
}
