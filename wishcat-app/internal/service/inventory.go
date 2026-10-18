package service

import "wishcat-app/internal/domain"

type InMemoryInventory struct {
	items []domain.Item
}

func NewInMemoryInventory() *InMemoryInventory {
	return &InMemoryInventory{
		items: make([]domain.Item, 0),
	}
}

func (i *InMemoryInventory) Add(item domain.Item) {
	i.items = append(i.items, item)
}

func (i *InMemoryInventory) GetAll() []domain.Item {
	return i.items
}
