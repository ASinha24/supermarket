package store

import (
	"context"
	"errors"
	"sync"

	"github.com/alka/supermartchannel/store/model"
)

type ItemStore interface {
	CreateItem(ctx context.Context, item *model.Item) (*model.Item, error)
	UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error)
	DeleteItem(ctx context.Context, id string) error
	FindItemByID(ctx context.Context, id string) (*model.Item, error)
	GetAllMartItems(ctx context.Context, martID string) ([]*model.Item, error)
	Close()
}

//since we are using map which the only shared variable hence created channel to make it thread safe
type ItemStoreInMem struct {
	items  map[string]*model.Item
	itemCh chan<- *model.Item
	once   sync.Once
}

func (i *ItemStoreInMem) CreateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	i.itemCh <- item
	return item, nil
}

func (i *ItemStoreInMem) UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	oldItem, ok := i.items[item.ID]
	if !ok {
		return nil, errors.New("item not found")
	}
	oldItem.Name = item.Name
	oldItem.Price = item.Price
	i.itemCh <- oldItem
	return oldItem, nil
}

func (i *ItemStoreInMem) DeleteItem(ctx context.Context, id string) error {
	_, ok := i.items[id]
	if !ok {
		return errors.New("item not found")
	}
	delete(i.items, id)
	return nil
}

func (i *ItemStoreInMem) FindItemByID(ctx context.Context, id string) (*model.Item, error) {
	item, ok := i.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (i *ItemStoreInMem) GetAllMartItems(ctx context.Context, martID string) ([]*model.Item, error) {
	var items []*model.Item
	for _, item := range i.items {
		if item.SuperMartID == martID {
			items = append(items, item)
		}
	}
	return items, nil
}

func (i *ItemStoreInMem) Close() {
	i.once.Do(func() { close(i.itemCh) })
}

func NewItemStore() ItemStore {
	itemCh := make(chan *model.Item)
	s := &ItemStoreInMem{items: make(map[string]*model.Item), itemCh: itemCh}
	//goroutine to write the data into map in sync
	go func(ch <-chan *model.Item) {
		for i := range ch {
			s.items[i.ID] = i
		}
	}(itemCh)
	return s
}
