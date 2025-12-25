package services

import (
	"context"
	"errors"
	"log"
	"reflect"
	"sync"
)

const CacheServiceID = 1
const LogServiceID = 2

// FncParam is a lookup for param name e.g Key, to the type allowed.
type FncParam map[string]reflect.Type

type Service interface {
	Call(ctx context.Context, params any) (any, error)

	// For now, I'll start with Funcs that returns the service type and map of function names with associated params.
	Funcs() (int, map[string]FncParam)
}

type Store interface {
	Add(serviceID int, s Service) error
	Fetch(serviceID int) (Service, error)
}

// Simple in-mem store for now
type serviceStore struct {
	items map[int]Service
	mu    sync.Mutex
}

func (s *serviceStore) Fetch(serviceID int) (Service, error) {
	log.Println("fetching service")

	s.mu.Lock()
	defer s.mu.Unlock()

	service, found := s.items[serviceID]
	if !found {
		return nil, errors.New("service not found")
	}

	return service, nil
}

func (s *serviceStore) Add(serviceID int, service Service) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[serviceID] = service
	return nil
}

func NewStore() Store {
	return &serviceStore{
		items: map[int]Service{},
		mu:    sync.Mutex{},
	}
}
