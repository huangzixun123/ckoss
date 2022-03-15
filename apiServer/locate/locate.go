package locate

import (
	"ckoss/pkg/rabbitmq"
	"ckoss/pkg/rs"
	"ckoss/pkg/types"
	"encoding/json"
	"os"
	"time"
)

// name : object name
// 返回存储对应对象的数据服务节点的地址
// locateInfo key : 分片ID  value : 监听地址
func Locate(name string) (locateInfo map[int]string) {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}
