package ztask

import "github.com/hibiken/asynq"

type Client struct {
	Cli *asynq.Client
}

// NewClient Allow custom asynq.RedisConnOpt
func NewClient(r asynq.RedisConnOpt) *Client {
	return &Client{Cli: asynq.NewClient(r)}
}

// NewEasyClient Simple parameters to participate in Client
func NewEasyClient(param EasyParam) *Client {
	return NewClient(asynq.RedisClientOpt{
		Addr:     param.Addr,
		Password: param.Password,
		DB:       param.DB,
	})
}

func (c Client) Close() error {
	return c.Cli.Close()
}

func (c Client) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return c.Cli.Enqueue(task, opts...)
}
