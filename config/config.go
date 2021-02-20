package config

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"os"
	_ "time"
)

type Configs struct {
	Development Configuration
	Testing     Configuration
	Production  Configuration
}

type Configuration struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
	Name     string
}

var Conf Configs
var EnvConfig Configuration

func GetConfigByEnv(env string) Configuration {
	fmt.Println("all config is: ", Conf)
	switch {
	case env == "development":
		return Conf.Development
	case env == "testing":
		return Conf.Testing
	default:
		return Conf.Production
	}
}

func init() {
	v := viper.New()
	// 配置文件路径
	v.AddConfigPath(".")
	// 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	v.SetConfigName("config")
	// 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	v.SetConfigType("json")
	// 搜索并读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将配置文件绑定到Conf上
	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal error ViperConfig json: %s \n", err))
	}
	// 通过环境变量动态控制配置
	if env := os.Getenv("ENV"); env != "" {
		fmt.Printf("本地环境变量 ENV=%s\n", env)
		EnvConfig = GetConfigByEnv(env)
	} else {
		fmt.Println("本地环境变量ENV为空，默认值为 production，请确认服务器环境！")
		EnvConfig = GetConfigByEnv("production")
	}
	// 本地化初始设置
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	//bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active.en.json")
	//bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	//ViperConfig.LocaleBundle = bundle
}
