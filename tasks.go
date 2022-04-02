package ztask

import "github.com/hibiken/asynq"

var tasks = make(map[string]asynq.Handler)

// RegisterAsynqTask Bind the pattern to the handler.
// When the server runs, it will be processed according to the pattern
func RegisterAsynqTask(pattern string, handler asynq.Handler) {
	tasks[pattern] = handler
}

func TaskHandle(mux *asynq.ServeMux) {
	for pattern, handler := range tasks {
		mux.Handle(pattern, handler)
	}
}
