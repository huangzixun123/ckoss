package locate

import (
	"ckoss/pkg/rabbitmq"
	"os"
	"strconv"
	"time"
)

// 用以接口服务收到对象GET请求时定位该对象被保存在哪个数据服务节点
func Locate(name string) string {
	// 创建消息队列
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	// 向 dataServers群发这个对象名字
	q.Publish("dataServers", name)
	c := q.Consume()
	// 一秒钟后关闭，防止无限制的等待
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	//如果一秒钟内有回应，返回数据服务节点的监听地址
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
