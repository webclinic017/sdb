package main

import (
	"github.com/yemingfeng/sdb/internal/util"
	pb "github.com/yemingfeng/sdb/pkg/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var bitsetLogger = util.GetLogger("bitset")

func main() {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		bitsetLogger.Printf("faild to connect: %+v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// 连接服务器
	c := pb.NewSDBClient(conn)
	// 发起 mset 请求
	bsMSetResponse, err := c.BSMSet(context.Background(),
		&pb.BSMSetRequest{Key: []byte("hello"), Bits: []uint32{1, 2, 3}, Value: true})
	bitsetLogger.Printf("bsMSetResponse: %+v, err: %+v", bsMSetResponse, err)
	// 发起 mget 请求
	bsMGetResponse, err := c.BSMGet(context.Background(),
		&pb.BSMGetRequest{Key: []byte("hello"), Bits: []uint32{4, 1, 2, 3, 5}})
	bitsetLogger.Printf("bsMGetResponse: %+v, err: %+v", bsMGetResponse, err)
	// 发起 setrange 请求
	bsSetResponse, err := c.BSSetRange(context.Background(),
		&pb.BSSetRangeRequest{Key: []byte("hello"), Start: 10, End: 20, Value: true})
	bitsetLogger.Printf("bsSetResponse: %+v, err: %+v", bsSetResponse, err)
	// 发起 getrange 请求
	bsGetResponse, err := c.BSGetRange(context.Background(),
		&pb.BSGetRangeRequest{Key: []byte("hello"), Start: 9, End: 21})
	bitsetLogger.Printf("bsGetResponse: %+v, err: %+v", bsGetResponse, err)
	// 发起 count range 请求
	bsCountRangeResponse, err := c.BSCountRange(context.Background(),
		&pb.BSCountRangeRequest{Key: []byte("hello"), Start: 15, End: 100})
	bitsetLogger.Printf("bsCountRangeResponse: %+v, err: %+v", bsCountRangeResponse, err)
	// 发起 count 请求
	bsCountResponse, err := c.BSCount(context.Background(),
		&pb.BSCountRequest{Key: []byte("hello")})
	bitsetLogger.Printf("bsCountResponse: %+v, err: %+v", bsCountResponse, err)
	// 发起 del 请求
	bsDelResponse, err := c.BSDel(context.Background(),
		&pb.BSDelRequest{Key: []byte("hello")})
	bitsetLogger.Printf("bsDelResponse: %+v, err: %+v", bsDelResponse, err)
	//发起 count 请求
	bsCountResponse, err = c.BSCount(context.Background(),
		&pb.BSCountRequest{Key: []byte("hello")})
	bitsetLogger.Printf("bsCountResponse: %+v, err: %+v", bsCountResponse, err)
}
