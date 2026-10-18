package domain

type Helmet struct {
	BaseItem
	size string
}

func NewHelmet(invnum, name, size string) *Helmet {
	return &Helmet{
		BaseItem: NewBaseItem(invnum, name),
		size:     size,
	}
}

//--------------------------------------

type ChargingStation struct {
	BaseItem
	BaseEnergyConsumer
}

func NewChargingStation(invnum, name string, energy float64) *ChargingStation {
	return &ChargingStation{
		BaseItem:           NewBaseItem(invnum, name),
		BaseEnergyConsumer: NewBaseEnergyConsumer(energy),
	}
}
