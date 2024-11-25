package config

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App           App           `mapstructure:"app"`
	Log           Log           `mapstructure:"log"`
	Sqlite        Sqlite        `mapstructure:"sqlite"`
	Secrets       Secrets       `mapstructure:"secrets"`
	JsonDataFiles JsonDataFiles `mapstructure:"jsondatafiles"`
}

type Log struct {
	Level string `mapstructure:"level"`
}
type App struct {
	Name    string  `mapstructure:"name"`
	Version float64 `mapstructure:"version"`
	Port    int     `mapstructure:"port"`
	Env     string  `mapstructure:"env"`
}
type Sqlite struct {
	Name               string        `mapstructure:"dbname"`
	Path               string        `mapstructure:"dbpath"`
	MaxIdleConns       int           `mapstructure:"maxIdleConns"`
	MaxOpenConns       int           `mapstructure:"maxOpenConns"`
	MaxLifeTimeMinutes time.Duration `mapstructure:"maxLifeTimeMinutes"`
}
type JsonDataFiles struct {
	PathFile string `mapstructure:"pathFile"`
	NameFile string `mapstructure:"nameFile"`
}

type Secrets struct {
	JwtKey string `mapstructure:"jwt-key"`
}

var config Config
var configOnce sync.Once

func InitConfig() Config {
	configOnce.Do(func() {

		viper.SetConfigName("config")   // ชื่อ config file
		viper.AddConfigPath("./config") // ระบุ path ของ config file
		viper.AutomaticEnv()            // อ่าน value จาก ENV variable
		// แปลง _ underscore ใน env เป็น . dot notation ใน viper
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		err := viper.Unmarshal(&config)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

	})
	return config
}
