package httpStu

import (
	"fmt"
	"net"
)

func HttpServer(){
	//使用tcp协议，配置监听端口
	listener, e := net.Listen("tcp", ":8889")
	defer listener.Close()
	if e != nil{
		fmt.Println(e)
		return
	}

	for {
		//接收客户端请求，阻塞
		conn, e := listener.Accept()
		if e != nil{
			fmt.Println(e)
			return
		}else {
			fmt.Println("Client IP: ", conn.RemoteAddr())
		}

		//开启一个携程处理客户端的请求
		go func(connection net.Conn) {
			defer connection.Close()
			for {
				buf := make([]byte,1024)
				n, err := connection.Read(buf)
				if err != nil{
					fmt.Println(err)
					return
				}
				fmt.Println(string(buf[:n]))
				i, err := connection.Write([]byte("HelloWorld"))
				if err != nil{
					fmt.Println(err)
					return
				}
				fmt.Println(i)
			}

		}(conn)

	}

}

func HttpClient()  {

	conn, e := net.Dial("tcp", "127.0.0.1:8889")
	defer conn.Close()
	if e != nil{
		fmt.Println(e)
		return
	}
	//读取从终端输入信息
	//reader := bufio.NewReader(os.Stdin)
	//
	//s, e := reader.ReadString('\n')
	//if e != nil{
	//	fmt.Println(e)
	//	return
	//}
	s := "This is a message from client..."
	n, e := conn.Write([]byte(s))
	if e != nil{
		fmt.Println(e)
		return
	}

	fmt.Println(n)

}