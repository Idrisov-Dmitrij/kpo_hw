package service

import (
	"errors"
	"wishcat-app/internal/domain"
)

type ParkManager struct {
	inventory Inventory
	inspector Inspector
}

func NewParkManager(inventory Inventory, inspector Inspector) *ParkManager {
	return &ParkManager{
		inventory: inventory,
		inspector: inspector,
	}
}

func (pm *ParkManager) AddToPark(item domain.Item) error {
	if !pm.inspector.Inspect(item) {
		return errors.New("Техосмотр не пройден.")
	}

	for _, existingItem := range pm.inventory.GetAll() {
		if existingItem.GetInventoryNumber() == item.GetInventoryNumber() {
			return errors.New("объект с таким инвентарным номером уже существует в парке")
		}
	}

	pm.inventory.Add(item)
	return nil
}

func (pm *ParkManager) GetTotalEnergy() float64 {
	var total float64
	items := pm.inventory.GetAll()

	for _, item := range items {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.GetDailyEnergyConsumption()
		}
	}

	return total
}

func (pm *ParkManager) GetBeginnerFriendly() []domain.Item {
	var result []domain.Item
	items := pm.inventory.GetAll()

	for _, item := range items {
		if vehicle, ok := item.(domain.Vehicle); ok {
			if vehicle.GetSimplicity() >= 6 {
				result = append(result, item)
			}
		}
	}

	return result
}

func (pm *ParkManager) GetAllInventory() []domain.Item {
	return pm.inventory.GetAll()
}
