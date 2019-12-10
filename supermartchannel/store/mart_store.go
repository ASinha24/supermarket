package store

import (
	"context"
	"errors"
	"sync"

	"github.com/alka/supermartchannel/store/model"
)

type MartStore interface {
	CreateMart(ctx context.Context, mart *model.SuperMart) error
	GetMartByName(ctx context.Context, name string) (*model.SuperMart, error)
	GetMartByID(ctx context.Context, id string) (*model.SuperMart, error)
	Close()
}

//channel
type MartInMemStore struct {
	marts  map[string]*model.SuperMart
	martCh chan<- *model.SuperMart
	once   sync.Once
}

func (m *MartInMemStore) CreateMart(ctx context.Context, mart *model.SuperMart) error {
	m.martCh <- mart
	return nil
}

func (m *MartInMemStore) GetMartByName(ctx context.Context, name string) (*model.SuperMart, error) {
	for _, m := range m.marts {
		if m.Name == name {
			return m, nil
		}
	}
	return nil, errors.New("mart not found")
}

func (m *MartInMemStore) GetMartByID(ctx context.Context, id string) (*model.SuperMart, error) {
	mart, ok := m.marts[id]
	if !ok {
		return nil, errors.New("mart not found")
	}
	return mart, nil
}

func (m *MartInMemStore) Close() {
	m.once.Do(func() {
		close(m.martCh)
	})
}

//goroutine to write the data into map in sync
func NewMartStore() MartStore {
	martCh := make(chan *model.SuperMart)
	m := &MartInMemStore{marts: make(map[string]*model.SuperMart), martCh: martCh}
	go func(ch <-chan *model.SuperMart) {
		for mart := range ch {
			m.marts[mart.ID] = mart
		}
	}(martCh)
	return m
}
