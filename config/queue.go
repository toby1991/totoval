package config

import (
	. "github.com/totoval/framework/config"
	"totoval/app/models"
)

func init() {
	queue := make(map[string]interface{})

	queue["default"] = Env("QUEUE_CONNECTION", "memory")
	queue["connections"] = map[string]interface{}{
		"nsq": map[string]interface{}{
			"host": Env("QUEUE_NSQ_HOST", "127.0.0.1"),
			"port": Env("QUEUE_NSQ_PORT", "4150"),
		},
	}

	queue["failed_db_processor_model"] = &models.FailedQueue{}

	Add("queue", queue)
}
