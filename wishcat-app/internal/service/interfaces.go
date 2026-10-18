package service

import "wishcat-app/internal/domain"

type Inspector interface {
	Inspect(item domain.Item) bool
}

type Inventory interface {
	Add(item domain.Item)
	GetAll() []domain.Item
}
