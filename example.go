package ztask

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeDynamicConfig = "dynamic:config"
)

type DynamicConfigPayload struct {
	GameID     string
	Source     string
	ConfigName string
}

func NewDynamicConfigTask(gameID, source, configName string) (*asynq.Task, error) {
	payload, err := json.Marshal(DynamicConfigPayload{GameID: gameID, Source: source, ConfigName: configName})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDynamicConfig, payload), nil
}

type DynamicConfigProcessor struct{}

func (processor *DynamicConfigProcessor) ProcessTask(_ context.Context, t *asynq.Task) error {
	var p DynamicConfigPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	fmt.Println(p.GameID, p.Source, p.ConfigName)
	return nil
}

func NewDynamicConfigProcessor() *DynamicConfigProcessor {
	return &DynamicConfigProcessor{}
}
