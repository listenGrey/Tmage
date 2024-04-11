package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode       string `mapstructure:"mode"`
	Port       int    `mapstructure:"port"`
	Name       string `mapstructure:"name"`
	Version    string `mapstructure:"version"`
	StartTime  string `mapstructure:"start_time"`
	MachineID  int    `mapstructure:"machine_id"`
	*LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() error {
	// 读取配置文件
	viper.SetConfigFile("./conf/config.yaml")
	// 读取环境变量
	viper.WatchConfig()
	// 监听配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		viper.Unmarshal(&Conf)
	})
	// 查找并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
