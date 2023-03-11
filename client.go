package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	// 서버와 연결시도
	if err != nil {
		log.Println(err)
	}

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()
	// 고루틴 생성하여 서버가 값을 던질 때까지 기다렸다가
	// 던지면 값 출력

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
	// 사용자 입력할때 까지 기다렸다가
	// 입력하면 서버로 전송
}
