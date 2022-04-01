package tasks

import "github.com/hibiken/asynq"

var tasks = make(map[string]asynq.Handler)

func RegisterAsynqTask(pattern string, handler asynq.Handler) {
	tasks[pattern] = handler
}

func TaskHandle(mux *asynq.ServeMux) {
	for pattern, handler := range tasks {
		mux.Handle(pattern, handler)
	}
}

func InitTasks() {
	RegisterAsynqTask(TypeDynamicConfig, NewDynamicConfigProcessor())
}
