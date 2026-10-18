package domain

type BaseItem struct {
	inventoryNumber string
	name            string
}

func NewBaseItem(invNum, name string) BaseItem {
	return BaseItem{
		inventoryNumber: invNum,
		name:            name,
	}
}

func (b *BaseItem) GetInventoryNumber() string { return b.inventoryNumber }
func (b *BaseItem) GetName() string            { return b.name }

//--------------------------------------

type BaseEnergyConsumer struct {
	dailyEnergy float64
}

func NewBaseEnergyConsumer(energy float64) BaseEnergyConsumer {
	return BaseEnergyConsumer{
		dailyEnergy: energy,
	}
}

func (b *BaseEnergyConsumer) GetDailyEnergyConsumption() float64 { return b.dailyEnergy }

//--------------------------------------

type BaseVehicle struct {
	simplicity int
	durability int
}

func NewBaseVehicle(simplicity, durability int) BaseVehicle {
	return BaseVehicle{
		simplicity: simplicity,
		durability: durability,
	}
}

func (b *BaseVehicle) GetDurability() int { return b.durability }
func (b *BaseVehicle) GetSimplicity() int { return b.simplicity }
