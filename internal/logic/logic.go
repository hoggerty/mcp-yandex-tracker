package logic

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

type Services struct {
	Issue *issue.Service
	Queue *queue.Service
	User  *user.Service
	Board *board.Service
}

func New(client *trackerapi.Client) *Services {
	return &Services{
		Issue: issue.New(client),
		Queue: queue.New(client),
		User:  user.New(client),
		Board: board.New(client),
	}
}
