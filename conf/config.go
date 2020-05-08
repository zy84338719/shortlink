package conf

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Conf struct {
	Redis Redis
	Mysql Mysql
}

type Redis struct {
	Addr string
	Pwd  string
	DB   int
}

type Mysql struct {
	Addr     string
	Pwd      string
	Username string
}

var C *Conf

func init() {
	env := getMode()
	realpath, _ := filepath.Abs("./")
	viper.SetConfigName("app")
	viper.AddConfigPath(realpath + "/conf/" + env)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatel error config file: %s \n", err))
	}
	C = &Conf{}
	err = viper.Unmarshal(C)
	if err != nil {
		panic(err)
	}
}

func getMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}
