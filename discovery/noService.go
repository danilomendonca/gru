package discovery

import (
	"errors"
)

var ErrNoService error = errors.New("No discovery service")

type noService struct{}

func (p *noService) Name() string {
	return "noservice"
}

func (p *noService) Initialize(uri string) error {
	return ErrNoService
}

func (p *noService) Register(myUUID string, myAddress string, ttl int) error {
	return ErrNoService
}

func (p *noService) Get(key string) (map[string]string, error) {
	return nil, ErrNoService
}

func (p *noService) Set(key string, value string, ttl int) error {
	return ErrNoService
}