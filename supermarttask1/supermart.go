package supermart

import (
	"fmt"

	"github.com/alka/supermarttask1/store"
	"github.com/alka/supermarttask1/store/model"
	"github.com/pborman/uuid"

	"github.com/alka/supermarttask1/api"
)

type MartManager interface {
	GetItems(martName string) ([]api.CreateItemRespose, error)
	CreateItem(martName string, item *api.ItemRequest) (*api.CreateItemRespose, error)
	DeleteItem(martName string, itemID string) error
	UpdateItem(martName string, itemID string, item *api.ItemRequest) (*api.CreateItemRespose, error)
}

type SuperMartService struct {
	itemStore store.ItemStore
	martStore store.MartStore
}

func (s *SuperMartService) GetItems(martName string) ([]api.CreateItemRespose, error) {
	mart, err := s.martStore.GetMartByName(martName)
	if err != nil {
		return nil, &api.MartError{Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	items, err := s.itemStore.GetAllMartItems(mart.ID)
	if err != nil {
		return nil, err
	}

	martItems := []api.CreateItemRespose{}
	for _, item := range items {
		martItems = append(martItems, api.CreateItemRespose{
			ItemRequest: &api.ItemRequest{
				Name:  item.Name,
				Price: item.Price,
			},
			ID: item.ID,
		})
	}
	return martItems, nil
}

func (s *SuperMartService) CreateItem(martName string, item *api.ItemRequest) (*api.CreateItemRespose, error) {
	mart, err := s.martStore.GetMartByName(martName)
	if err != nil {
		return nil, &api.MartError{Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	newItem, err := s.itemStore.CreateItem(&model.Item{
		ID:          uuid.NewUUID().String(),
		Name:        item.Name,
		Price:       item.Price,
		SuperMartID: mart.ID,
	})
	if err != nil {
		return nil, &api.MartError{Message: "can not create new item", Description: err.Error()}
	}
	return &api.CreateItemRespose{ID: newItem.ID, ItemRequest: item}, nil
}

func (s *SuperMartService) DeleteItem(martName string, itemID string) error {
	if err := s.checkValidMart(martName, itemID); err != nil {
		return err
	}
	return s.itemStore.DeleteItem(itemID)
}

func (s *SuperMartService) UpdateItem(martName string, itemID string, item *api.ItemRequest) (*api.CreateItemRespose, error) {
	if err := s.checkValidMart(martName, itemID); err != nil {
		return nil, err
	}
	updateItem, err := s.itemStore.UpdateItem(&model.Item{
		ID:    itemID,
		Name:  item.Name,
		Price: item.Price,
	})
	if err != nil {
		return nil, &api.MartError{Message: fmt.Sprintf("can't update item with id %s", itemID), Description: err.Error()}
	}
	return &api.CreateItemRespose{ID: updateItem.ID, ItemRequest: item}, nil
}

func (s *SuperMartService) checkValidMart(martName string, itemID string) error {
	mart, err := s.martStore.GetMartByName(martName)
	if err != nil {
		return &api.MartError{Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	item, err := s.itemStore.FindItemByID(itemID)
	if err != nil {
		return &api.MartError{Message: fmt.Sprintf("item having id %s not found", itemID), Description: err.Error()}
	}

	if item.SuperMartID != mart.ID {
		return &api.MartError{Message: fmt.Sprintf("store %s is not authorised for item having id %s", martName, itemID), Description: ""}
	}
	return nil
}

func NewSuperMartService(itemStore store.ItemStore, martStore store.MartStore) *SuperMartService {
	return &SuperMartService{
		itemStore: itemStore,
		martStore: martStore,
	}
}
