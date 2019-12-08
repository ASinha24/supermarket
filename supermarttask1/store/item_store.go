package store

import (
	"errors"

	"github.com/alka/supermarttask1/store/model"
)

type ItemStore interface {
	CreateItem(item *model.Item) (*model.Item, error)
	UpdateItem(item *model.Item) (*model.Item, error)
	DeleteItem(id string) error
	FindItemByID(id string) (*model.Item, error)
	GetAllMartItems(martID string) ([]*model.Item, error)
}

type ItemStoreInMem struct {
	items map[string]*model.Item
}

func (i *ItemStoreInMem) CreateItem(item *model.Item) (*model.Item, error) {
	i.items[item.ID] = item
	return item, nil
}

func (i *ItemStoreInMem) UpdateItem(item *model.Item) (*model.Item, error) {
	oldItem, ok := i.items[item.ID]
	if !ok {
		return nil, errors.New("item not found")
	}
	oldItem.Name = item.Name
	oldItem.Price = item.Price
	i.items[oldItem.ID] = oldItem
	return oldItem, nil
}

func (i *ItemStoreInMem) DeleteItem(id string) error {
	_, ok := i.items[id]
	if !ok {
		return errors.New("item not found")
	}
	delete(i.items, id)
	return nil
}

func (i *ItemStoreInMem) FindItemByID(id string) (*model.Item, error) {
	item, ok := i.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (i *ItemStoreInMem) GetAllMartItems(martID string) ([]*model.Item, error) {
	var items []*model.Item
	for _, item := range i.items {
		if item.SuperMartID == martID {
			items = append(items, item)
		}
	}
	return items, nil
}

func NewItemStore() ItemStore {
	return &ItemStoreInMem{items: make(map[string]*model.Item)}
}
