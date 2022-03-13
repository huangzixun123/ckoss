package locate

import (
	"ckoss/pkg/rabbitmq"
	"os"
	"strconv"
	"time"
)

// name : object name
// 返回存储对应对象的数据服务节点的地址
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
