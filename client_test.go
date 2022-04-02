package ztask

import (
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewEasyClient(EasyParam{Addr: "127.0.0.1:6739"})
	defer client.Close()

	gameID, source, name := "game", "source", "name"
	task, err := NewDynamicConfigTask(gameID, source, name)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
