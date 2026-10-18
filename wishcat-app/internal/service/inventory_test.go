package service

import (
	"testing"
	"wishcat-app/internal/domain"
)

func TestInMemoryInventory_AddAndGetAll(t *testing.T) {
	inv := NewInMemoryInventory()
	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)
	helmet := domain.NewHelmet("H-001", "Шлем M", "M")

	inv.Add(scooter)
	inv.Add(helmet)

	items := inv.GetAll()
	if len(items) != 2 {
		t.Errorf("Ожидалось 2 элемента, получено %d", len(items))
	}
}

func TestInMemoryInventory_EmptyInventory(t *testing.T) {
	inv := NewInMemoryInventory()

	items := inv.GetAll()
	if len(items) != 0 {
		t.Errorf("Ожидалось 0 элементов, получено %d", len(items))
	}
}

func TestInMemoryInventory_AddMultiple(t *testing.T) {
	inv := NewInMemoryInventory()

	for i := 0; i < 10; i++ {
		scooter := domain.NewElectricScooter("ES-00"+string(rune('0'+i)), "Xiaomi", 0.5, 8, 75)
		inv.Add(scooter)
	}

	items := inv.GetAll()
	if len(items) != 10 {
		t.Errorf("Ожидалось 10 элементов, получено %d", len(items))
	}
}
