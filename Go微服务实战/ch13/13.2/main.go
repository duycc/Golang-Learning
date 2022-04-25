package main

import (
	"log"

	"Golang-Learning/Go微服务实战/ch13/13.2/protocol/protocol"

	"github.com/golang/protobuf/proto"
)

func main() {
	u := &protocol.UserInfo{
		Message: *proto.String("testInfo"),
		Length:  *proto.Int32(10),
	}

	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newInfo := &protocol.UserInfo{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Println(newInfo.GetMessage())
}
