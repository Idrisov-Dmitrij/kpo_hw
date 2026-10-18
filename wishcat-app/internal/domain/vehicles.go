package domain

type ElectricScooter struct {
	BaseItem
	BaseEnergyConsumer
	BaseVehicle
}

func NewElectricScooter(invNum, model string, energy float64, simplicity, durability int) *ElectricScooter {
	return &ElectricScooter{
		BaseItem:           NewBaseItem(invNum, model),
		BaseEnergyConsumer: NewBaseEnergyConsumer(energy),
		BaseVehicle:        NewBaseVehicle(simplicity, durability),
	}
}

//--------------------------------------

type Bicycle struct {
	BaseItem
	BaseVehicle
}

func NewBicycle(invNum, name string, simplicity, durability int) *Bicycle {
	return &Bicycle{
		BaseItem:    NewBaseItem(invNum, name),
		BaseVehicle: NewBaseVehicle(simplicity, durability),
	}
}

//--------------------------------------

type ElectricBicycle struct {
	BaseItem
	BaseEnergyConsumer
	BaseVehicle
}

func NewElectricBicycle(invNum, model string, energy float64, simplicity, durability int) *ElectricBicycle {
	return &ElectricBicycle{
		BaseItem:           NewBaseItem(invNum, model),
		BaseEnergyConsumer: NewBaseEnergyConsumer(energy),
		BaseVehicle:        NewBaseVehicle(simplicity, durability),
	}
}
