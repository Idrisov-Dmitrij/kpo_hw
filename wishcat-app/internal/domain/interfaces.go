package domain

type Item interface {
	GetInventoryNumber() string
	GetName() string
}

type EnergyConsumer interface {
	GetDailyEnergyConsumption() float64
}

type Vehicle interface {
	GetSimplicity() int
	GetDurability() int
}
