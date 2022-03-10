package heartbeat

import (
	"ckoss/pkg/rabbitmq"
	"os"
	"time"
)

// export RABBITMQ_SERVER=amqp://test:test@localhost:5672
func StartHeartbeat() {
	// 心跳信息
	// 生成一个队列，通过这个队列，将 dataServer 的监听地址发送给 exchange : apiServers
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		// 将 dataServer 的监听地址发送给 exchange : apiServers
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
