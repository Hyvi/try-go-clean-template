package amqprpc

import (
	"github.com/Hyvi/go-present/try-go-clean-template/internal/usecase"
	"github.com/Hyvi/go-present/try-go-clean-template/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
