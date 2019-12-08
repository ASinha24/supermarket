package store

import (
	"context"
	"errors"

	"github.com/alka/supermart/store/model"
)

type MartStore interface {
	CreateMart(ctx context.Context, mart *model.SuperMart) error
	GetMartByName(ctx context.Context, name string) (*model.SuperMart, error)
	GetMartByID(ctx context.Context, id string) (*model.SuperMart, error)
}

type MartInMemStore struct {
	marts map[string]*model.SuperMart
}

func (m *MartInMemStore) CreateMart(ctx context.Context, mart *model.SuperMart) error {
	m.marts[mart.ID] = mart
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

func NewMartStore() MartStore {
	return &MartInMemStore{marts: make(map[string]*model.SuperMart)}
}
