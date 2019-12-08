package store

import (
	"errors"

	"github.com/alka/supermarttask1/store/model"
)

type MartStore interface {
	CreateMart(mart *model.SuperMart) error
	GetMartByName(name string) (*model.SuperMart, error)
	GetMartByID(id string) (*model.SuperMart, error)
}

type MartInMemStore struct {
	marts map[string]*model.SuperMart
}

func (m *MartInMemStore) CreateMart(mart *model.SuperMart) error {
	m.marts[mart.ID] = mart
	return nil
}

func (m *MartInMemStore) GetMartByName(name string) (*model.SuperMart, error) {
	for _, m := range m.marts {
		if m.Name == name {
			return m, nil
		}
	}
	return nil, errors.New("mart not found")
}

func (m *MartInMemStore) GetMartByID(id string) (*model.SuperMart, error) {
	mart, ok := m.marts[id]
	if !ok {
		return nil, errors.New("mart not found")
	}
	return mart, nil
}

func NewMartStore() MartStore {
	return &MartInMemStore{marts: make(map[string]*model.SuperMart)}
}
