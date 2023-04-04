package server

import (
	"fmt"
	"go_rpc/config"
	"go_rpc/util"
	"net"
)

const HEADER_LENGTH = 10

//magic value for header on first 2byte
const HEADER_ONE = 0xAA
const HEADER_TWO = 0xBB

type RpcServer struct {
	ServerConfigInfo config.ServerConfig
}

func (server *RpcServer) start() {
	port := server.ServerConfigInfo.ListenPort
	fmt.Println(fmt.Sprintf("start to listen on port:%d", port))

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(fmt.Sprintf("error to listen on port:%d", port))
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(fmt.Sprintf("error to accept on port: %d\n"))
	}

	header := make([]byte, HEADER_LENGTH)
	_, err = conn.Read(header)

	if err != nil {
		fmt.Println("fail to read the header packet")
	}

	if header[0] != HEADER_ONE || header[1] != HEADER_TWO {
		fmt.Println("error magic value , close the client conn!!!")
		conn.Close()
		return
	}

	requestId, err := util.Byte2Int32(header[2:6])
	len, err := util.Byte2Int32(header[6:])
	fmt.Println(fmt.Sprintf("receive packet requestId:%d, len:%d", requestId, len))
	body := make([]byte, len)
	conn.Read(body)
}

func (server *RpcServer) stop() {

}
