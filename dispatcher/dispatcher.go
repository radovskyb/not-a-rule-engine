package dispatcher

import (
	"log"

	"github.com/radovskyb/not-a-rule-engine/services"
)

type Client interface {
	Dispatch(services.Service, any) (any, error)
}

type client struct{}

func New() Client {
	return &client{}
}

func (d *client) Dispatch(s services.Service, payload any) (any, error) {
	// Add switch to check type of service here (Probably add shared lookup keys or service.Type method)
	log.Println("dispatching service")
	return s.Call(payload)
}
