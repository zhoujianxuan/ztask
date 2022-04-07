package ztask

import (
	"context"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterAsynqTask(TypeDynamicConfig, &DynamicConfigProcessor{})
	ser := NewEasyServer(EasyParam{Addr: "127.0.0.1:6379"})

	ctx := context.TODO()
	ser.Run(ctx)
	select {
	case <-ctx.Done():
	}
}
