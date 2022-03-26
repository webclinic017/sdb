package main

import (
	"github.com/yemingfeng/sdb/internal/util"
	pb "github.com/yemingfeng/sdb/pkg/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
)

var stringLogger = util.GetLogger("string")

func main() {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		stringLogger.Printf("faild to connect: %+v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// 连接服务器
	c := pb.NewSDBClient(conn)
	// 发起 mset 请求
	msetResponse, err := c.MSet(context.Background(),
		&pb.MSetRequest{Keys: [][]byte{[]byte("1"), []byte("2"), []byte("3")},
			Values: [][]byte{[]byte("4"), []byte("5"), []byte("6")}})
	stringLogger.Printf("msetResponse: %+v, err: %+v", msetResponse, err)
	// 发起 setnx 请求
	setNXResponse, err := c.SetNX(context.Background(),
		&pb.SetNXRequest{Key: []byte("1"), Value: []byte("11")})
	stringLogger.Printf("setNXResponse: %+v, err: %+v", setNXResponse, err)
	// 发起 setnx 请求
	setNXResponse, err = c.SetNX(context.Background(),
		&pb.SetNXRequest{Key: []byte("10"), Value: []byte("11")})
	stringLogger.Printf("setNXResponse: %+v, err: %+v", setNXResponse, err)
	// 发起 mget 请求
	mGetResponse, err := c.MGet(context.Background(),
		&pb.MGetRequest{Keys: [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("10"), []byte("100")}})
	stringLogger.Printf("mGetResponse: %+v, err: %+v", mGetResponse, err)
	// 发起 incr 请求
	incrResponse, err := c.Incr(context.Background(),
		&pb.IncrRequest{Key: []byte("abc"), Delta: 10})
	stringLogger.Printf("incrResponse: %+v, err: %+v", incrResponse, err)
	incrResponse, err = c.Incr(context.Background(),
		&pb.IncrRequest{Key: []byte("abc"), Delta: 1})
	stringLogger.Printf("incrResponse: %+v, err: %+v", incrResponse, err)
	// 发起 get 请求
	getResponse, err := c.Get(context.Background(),
		&pb.GetRequest{Key: []byte("abc")})
	stringLogger.Printf("getResponse: %+v, err: %+v", getResponse, err)
	// 发起 delete 请求
	delResponse, err := c.Del(context.Background(),
		&pb.DelRequest{Key: []byte("h")})
	stringLogger.Printf("delResponse: %+v, err: %+v", delResponse, err)
	// 发起 set 请求
	setResponse, err := c.Set(context.Background(),
		&pb.SetRequest{Key: []byte("h"), Value: []byte("h")})
	stringLogger.Printf("setResponse: %+v, err: %+v", setResponse, err)
	// 发起 get 请求
	getResponse, err = c.Get(context.Background(),
		&pb.GetRequest{Key: []byte("h")})
	stringLogger.Printf("getResponse: %+v, err: %+v", getResponse, err)
	for i := 0; i < 100; i++ {
		_, _ = c.Set(context.Background(),
			&pb.SetRequest{Key: []byte("h111" + strconv.Itoa(i)), Value: []byte("h" + strconv.Itoa(i))})
	}
}
