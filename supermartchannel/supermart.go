package supermart

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/alka/supermartchannel/api"
	"github.com/alka/supermartchannel/store"
	"github.com/alka/supermartchannel/store/model"
)

type MartManager interface {
	GetItems(ctx context.Context, martName string) ([]api.CreateItemRespose, error)
	CreateItem(ctx context.Context, martName string, item *api.ItemRequest) (*api.CreateItemRespose, error)
	DeleteItem(ctx context.Context, martName string, itemID string) error
	UpdateItem(ctx context.Context, martName string, itemID string, item *api.ItemRequest) (*api.CreateItemRespose, error)
}

type SuperMartService struct {
	itemStore store.ItemStore
	martStore store.MartStore
}

func (s *SuperMartService) GetItems(ctx context.Context, martName string) ([]api.CreateItemRespose, error) {
	mart, err := s.martStore.GetMartByName(ctx, martName)
	if err != nil {
		return nil, &api.MartError{Code: api.StoreNotFound, Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	items, err := s.itemStore.GetAllMartItems(ctx, mart.ID)
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

func (s *SuperMartService) CreateItem(ctx context.Context, martName string, item *api.ItemRequest) (*api.CreateItemRespose, error) {
	mart, err := s.martStore.GetMartByName(ctx, martName)
	if err != nil {
		return nil, &api.MartError{Code: api.StoreNotFound, Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	newItem, err := s.itemStore.CreateItem(ctx, &model.Item{
		ID:          uuid.NewUUID().String(),
		Name:        item.Name,
		Price:       item.Price,
		SuperMartID: mart.ID,
	})
	if err != nil {
		return nil, &api.MartError{Code: api.ItemCreationFailed, Message: "can not create new item", Description: err.Error()}
	}
	return &api.CreateItemRespose{ID: newItem.ID, ItemRequest: item}, nil
}

func (s *SuperMartService) DeleteItem(ctx context.Context, martName string, itemID string) error {
	if err := s.checkValidMart(ctx, martName, itemID); err != nil {
		return err
	}
	return s.itemStore.DeleteItem(ctx, itemID)
}

func (s *SuperMartService) UpdateItem(ctx context.Context, martName string, itemID string, item *api.ItemRequest) (*api.CreateItemRespose, error) {
	if err := s.checkValidMart(ctx, martName, itemID); err != nil {
		return nil, err
	}
	updateItem, err := s.itemStore.UpdateItem(ctx, &model.Item{
		ID:    itemID,
		Name:  item.Name,
		Price: item.Price,
	})
	if err != nil {
		return nil, &api.MartError{Code: api.ItemUpdateFailed, Message: fmt.Sprintf("can't update item with id %s", itemID), Description: err.Error()}
	}
	return &api.CreateItemRespose{ID: updateItem.ID, ItemRequest: item}, nil
}

func (s *SuperMartService) checkValidMart(ctx context.Context, martName string, itemID string) error {
	mart, err := s.martStore.GetMartByName(ctx, martName)
	if err != nil {
		return &api.MartError{Code: api.StoreNotFound, Message: fmt.Sprintf("mart having name %s not found", martName), Description: err.Error()}
	}

	item, err := s.itemStore.FindItemByID(ctx, itemID)
	if err != nil {
		return &api.MartError{Code: api.ItemNotFound, Message: fmt.Sprintf("item having id %s not found", itemID), Description: err.Error()}
	}

	if item.SuperMartID != mart.ID {
		return &api.MartError{Code: api.UnauthorisedStore, Message: fmt.Sprintf("store %s is not authorised for item having id %s", martName, itemID), Description: ""}
	}
	return nil
}

func NewSuperMartService(itemStore store.ItemStore, martStore store.MartStore) *SuperMartService {
	return &SuperMartService{
		itemStore: itemStore,
		martStore: martStore,
	}
}
