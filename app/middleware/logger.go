package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

func LoggerConfig() middleware.LoggerConfig {
	logger := middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","uri":"${uri}","method":"${method}","user_agent":"${user_agent}"`+
			`"status":"${status}","latency_human":${latency_human},"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}
	return logger
}
