package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	//格林时间1970-1-1到现在的秒数
	start := now.Unix()
	//格林时间纳秒数
	startNano := now.UnixNano()

	fmt.Println(now)

	dateString := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateString)

	//休眠
	time.Sleep(10 * time.Second)
	//数字格式固定，年份必须为2006，月份须为01，日期须为02， 小时为03代表12小时制，15代表24小时制，分钟须为04，秒须为05，排序和间隔符号自由组合
	fmt.Println(now.Format("2006-01-02 03:04:05"))

	end := time.Now().Unix()
	endNano := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Println(endNano - startNano)

}
