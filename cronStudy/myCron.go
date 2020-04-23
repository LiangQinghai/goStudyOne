package cronStudy

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func CronOne(cronStr string) {

	c := cron.New()

	addFunc, err := c.AddFunc("*/10 * * * ?", func() {
		fmt.Println("---------")
	})

	fmt.Println(addFunc, err)

	if err != nil {
		panic(err)
	}

	go c.Start()

	defer c.Stop()

	select {
	case <-time.After(time.Second * 60):
		return
	}

}
