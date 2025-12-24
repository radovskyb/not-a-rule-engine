package services

import (
	"errors"
	"log"
	"sync"
)

type Service interface {
	// TODO: Common service methods?
	//
	// For now, service will just pass params, data, and return any response and/or error.
	Call(params any) (any, error)
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
