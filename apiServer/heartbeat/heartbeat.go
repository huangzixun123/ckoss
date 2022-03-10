package heartbeat

import (
	"ckoss/pkg/rabbitmq"
	"os"
	"strconv"
	"sync"
	"time"
)

// APIServer 用于接受数据服务的心跳，更新dataServers，删除失联心跳
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume() // 返回收到的心跳：数据服务地址列表
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			// 如果上次收到数据服务的心跳跟现在系统的时间大于10秒，删除该数据服务
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}

		mutex.Unlock()
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()

	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}
