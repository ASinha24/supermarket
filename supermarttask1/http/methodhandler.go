package http

import (
	"github.com/pborman/uuid"

	supermart "github.com/alka/supermarttask1"
	"github.com/alka/supermarttask1/api"
	"github.com/alka/supermarttask1/store"
	"github.com/alka/supermarttask1/store/model"
)

type MartHandler struct {
	manager   supermart.MartManager
	martStore store.MartStore
}

func (m *MartHandler) CreateNewMart(martName string) error {
	martID := uuid.New()
	if err := m.martStore.CreateMart(&model.SuperMart{
		Name: martName,
		ID:   martID,
	}); err != nil {
		return err
	}
	return nil
}

func (m *MartHandler) GetItems(martName string) ([]api.CreateItemRespose, error) {
	items, err := m.manager.GetItems(martName)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (m *MartHandler) CreateItem(martName string, itemCreateReq *api.ItemRequest) (*api.CreateItemRespose, error) {
	resp, err := m.manager.CreateItem(martName, itemCreateReq)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (m *MartHandler) UpdateItem(martName string, itemID string, itemCreateReq *api.ItemRequest) (*api.CreateItemRespose, error) {
	resp, err := m.manager.UpdateItem(martName, itemID, itemCreateReq)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (m *MartHandler) DeleteItem(martName string, itemID string) error {
	if err := m.manager.DeleteItem(martName, itemID); err != nil {
		return err
	}
	return nil
}

func NewMartHandler(manager supermart.MartManager, martStore store.MartStore) *MartHandler {
	return &MartHandler{
		manager:   manager,
		martStore: martStore,
	}
}
