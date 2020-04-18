package viperStu

import (
	"fmt"
	"github.com/spf13/viper"
)

// 读取配置文件
func read() {

	viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	get := viper.Get("app.name")

	fmt.Println(get)

	one := viper.Sub("one")

	fmt.Println(one.GetString("sub01"), one.GetString("sub02"), one.AllKeys())

	slice := viper.GetStringSlice("two")

	println(slice[0], slice[1])

}
