package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	// 소켓 열어주기
	if err != nil {
		log.Println(err)
	}
	defer l.Close()
	// 프로그램 끝나기 전 소켓 닫아주기

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()
		go ConnHandler(conn)
		// 고루틴
		// 연결도 메인프로세스 종료 전에 종료
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err := conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
