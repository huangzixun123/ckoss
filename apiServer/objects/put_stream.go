package objects

import (
	"ckoss/apiServer/heartbeat"
	"ckoss/pkg/objectstream"
	"fmt"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
