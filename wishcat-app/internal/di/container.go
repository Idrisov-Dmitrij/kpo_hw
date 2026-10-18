package di

import "wishcat-app/internal/service"

type Container struct{}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) BuildParkManager() *service.ParkManager {
	inventory := service.NewInMemoryInventory()
	inspector := service.NewDefaultServiceCenter()

	return service.NewParkManager(inventory, inspector)
}
