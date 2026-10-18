package domain

import (
	"testing"
)

var _ Item = (*ElectricScooter)(nil)
var _ Vehicle = (*ElectricScooter)(nil)
var _ EnergyConsumer = (*ElectricScooter)(nil)

var _ Item = (*Bicycle)(nil)
var _ Vehicle = (*Bicycle)(nil)

var _ Item = (*ElectricBicycle)(nil)
var _ Vehicle = (*ElectricBicycle)(nil)
var _ EnergyConsumer = (*ElectricBicycle)(nil)

var _ Item = (*Helmet)(nil)

var _ Item = (*ChargingStation)(nil)
var _ EnergyConsumer = (*ChargingStation)(nil)

func TestNewElectricScooter(t *testing.T) {
	scooter := NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)

	if scooter.GetInventoryNumber() != "ES-001" {
		t.Errorf("Ожидался номер ES-001, получено %s", scooter.GetInventoryNumber())
	}
	if scooter.GetName() != "Xiaomi" {
		t.Errorf("Ожидалось имя Xiaomi, получено %s", scooter.GetName())
	}
	if scooter.GetDailyEnergyConsumption() != 0.5 {
		t.Errorf("Ожидалась энергия 0.5, получено %f", scooter.GetDailyEnergyConsumption())
	}
	if scooter.GetSimplicity() != 8 {
		t.Errorf("Ожидалась простота 8, получено %d", scooter.GetSimplicity())
	}
	if scooter.GetDurability() != 75 {
		t.Errorf("Ожидалась прочность 75, получено %d", scooter.GetDurability())
	}
}

func TestNewBicycle(t *testing.T) {
	bike := NewBicycle("B-001", "Trek", 7, 80)

	if bike.GetInventoryNumber() != "B-001" {
		t.Errorf("Ожидался номер B-001, получено %s", bike.GetInventoryNumber())
	}
	if bike.GetSimplicity() != 7 {
		t.Errorf("Ожидалась простота 7, получено %d", bike.GetSimplicity())
	}
	if _, ok := interface{}(bike).(EnergyConsumer); ok {
		t.Error("Bicycle не должен реализовывать EnergyConsumer")
	}
}

func TestNewHelmet(t *testing.T) {
	helmet := NewHelmet("H-001", "Шлем M", "M")

	if helmet.GetInventoryNumber() != "H-001" {
		t.Errorf("Ожидался номер H-001, получено %s", helmet.GetInventoryNumber())
	}
	if _, ok := interface{}(helmet).(Vehicle); ok {
		t.Error("Helmet не должен реализовывать Vehicle")
	}
	if _, ok := interface{}(helmet).(EnergyConsumer); ok {
		t.Error("Helmet не должен реализовывать EnergyConsumer")
	}
}

func TestNewChargingStation(t *testing.T) {
	charger := NewChargingStation("CS-001", "Зарядка Pro", 2.5)

	if charger.GetInventoryNumber() != "CS-001" {
		t.Errorf("Ожидался номер CS-001, получено %s", charger.GetInventoryNumber())
	}
	if charger.GetDailyEnergyConsumption() != 2.5 {
		t.Errorf("Ожидалась энергия 2.5, получено %f", charger.GetDailyEnergyConsumption())
	}
	if _, ok := interface{}(charger).(Vehicle); ok {
		t.Error("ChargingStation не должен реализовывать Vehicle")
	}
}
