package bootstrap

import (
	"fmt"
	"gin-blog/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	config := "config.yaml"

	// 创建新实例
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件变化
	v.WatchConfig()

	// 配置文件变化时重载到全局
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 载入配置到全局
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
