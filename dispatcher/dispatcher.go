package dispatcher

import (
	"context"
	"log"

	"github.com/radovskyb/not-a-rule-engine/services"
)

type Client interface {
	Dispatch(context.Context, services.Service, any) (any, error)
}

type client struct{}

func New() Client {
	return &client{}
}

func (d *client) Dispatch(ctx context.Context, s services.Service, payload any) (any, error) {
	// Add switch to check type of service here (Probably add shared lookup keys or service.Type method)
	log.Println("dispatching service")
	return s.Call(ctx, payload)
}
