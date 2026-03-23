package main

import (
	"dashboard-data/grpc/proto"
	"fmt"
)

func main() {
	// 测试是否可以创建IPOAnalysis类型
	ipo := &proto.IPOAnalysis{}
	fmt.Println(ipo)
	fmt.Println("测试成功：proto.IPOAnaalysis类型存在")
}
