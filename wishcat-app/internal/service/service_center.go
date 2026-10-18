package service

import "wishcat-app/internal/domain"

type DefaultServiceCenter struct{}

func NewDefaultServiceCenter() *DefaultServiceCenter {
	return &DefaultServiceCenter{}
}

func (sc *DefaultServiceCenter) Inspect(item domain.Item) bool {
	if item.GetInventoryNumber() == "" {
		return false
	}

	if consumer, ok := item.(domain.EnergyConsumer); ok {
		if consumer.GetDailyEnergyConsumption() < 0 {
			return false
		}
	}

	if vehicle, ok := item.(domain.Vehicle); ok {
		if vehicle.GetSimplicity() < 1 || vehicle.GetSimplicity() > 10 {
			return false
		}

		if vehicle.GetDurability() < 30 || vehicle.GetDurability() > 100 {
			return false
		}
	}

	return true
}
