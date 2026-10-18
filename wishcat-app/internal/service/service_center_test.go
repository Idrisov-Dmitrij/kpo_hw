package service

import (
	"testing"
	"wishcat-app/internal/domain"
)

func TestDefaultServiceCenter_ValidElectricScooter(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)

	if !sc.Inspect(scooter) {
		t.Error("Валидный самокат должен пройти техосмотр")
	}
}

func TestDefaultServiceCenter_EmptyInventoryNumber(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("", "Xiaomi", 0.5, 8, 75)

	if sc.Inspect(scooter) {
		t.Error("Самокат без инвентарного номера должен быть отклонён")
	}
}

func TestDefaultServiceCenter_NegativeEnergy(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", -0.5, 8, 75)

	if sc.Inspect(scooter) {
		t.Error("Самокат с отрицательной энергией должен быть отклонён")
	}
}

func TestDefaultServiceCenter_InvalidSimplicity_TooLow(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 0, 75)

	if sc.Inspect(scooter) {
		t.Error("Самокат с простотой 0 должен быть отклонён")
	}
}

func TestDefaultServiceCenter_InvalidSimplicity_TooHigh(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 11, 75)

	if sc.Inspect(scooter) {
		t.Error("Самокат с простотой 11 должен быть отклонён")
	}
}

func TestDefaultServiceCenter_InvalidDurability_TooLow(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 29)

	if sc.Inspect(scooter) {
		t.Error("Самокат с прочностью 29 должен быть отклонён")
	}
}

func TestDefaultServiceCenter_InvalidDurability_TooHigh(t *testing.T) {
	sc := NewDefaultServiceCenter()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 101)

	if sc.Inspect(scooter) {
		t.Error("Самокат с прочностью 101 должен быть отклонён")
	}
}

func TestDefaultServiceCenter_ValidHelmet(t *testing.T) {
	sc := NewDefaultServiceCenter()
	helmet := domain.NewHelmet("H-001", "Шлем M", "M")

	if !sc.Inspect(helmet) {
		t.Error("Валидный шлем должен пройти техосмотр")
	}
}

func TestDefaultServiceCenter_ValidChargingStation(t *testing.T) {
	sc := NewDefaultServiceCenter()
	charger := domain.NewChargingStation("CS-001", "Зарядка Pro", 2.5)

	if !sc.Inspect(charger) {
		t.Error("Валидная зарядная станция должна пройти техосмотр")
	}
}

func TestDefaultServiceCenter_InvalidChargingStationEnergy(t *testing.T) {
	sc := NewDefaultServiceCenter()
	charger := domain.NewChargingStation("CS-001", "Зарядка Pro", -1.0)

	if sc.Inspect(charger) {
		t.Error("Зарядная станция с отрицательной энергией должна быть отклонена")
	}
}
