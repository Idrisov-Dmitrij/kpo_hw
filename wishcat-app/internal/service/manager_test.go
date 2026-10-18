package service

import (
	"testing"
	"wishcat-app/internal/domain"
)

type MockInspector struct {
	ShouldPass bool
}

func (m *MockInspector) Inspect(item domain.Item) bool {
	return m.ShouldPass
}

type MockInventory struct {
	items        []domain.Item
	addCalled    bool
	getAllCalled bool
}

func (m *MockInventory) Add(item domain.Item) {
	m.items = append(m.items, item)
	m.addCalled = true
}

func (m *MockInventory) GetAll() []domain.Item {
	m.getAllCalled = true
	return m.items
}

func TestParkManager_AddToPark_WhenInspectionPasses(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	mockInventory := &MockInventory{}
	manager := NewParkManager(mockInventory, mockInspector)

	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)
	err := manager.AddToPark(scooter)

	if err != nil {
		t.Errorf("Не ожидалась ошибка, получено: %v", err)
	}
	if len(mockInventory.items) != 1 {
		t.Error("Самокат должен быть добавлен в хранилище")
	}
	if !mockInventory.addCalled {
		t.Error("Метод Add не был вызван у хранилища")
	}
}

func TestParkManager_AddToPark_WhenInspectionFails(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: false}
	mockInventory := &MockInventory{}
	manager := NewParkManager(mockInventory, mockInspector)

	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)
	err := manager.AddToPark(scooter)

	if err == nil {
		t.Error("Ожидалась ошибка техосмотра")
	}
	if len(mockInventory.items) != 0 {
		t.Error("Самокат НЕ должен быть добавлен в хранилище")
	}
}

func TestParkManager_GetTotalEnergy(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	inventory := NewInMemoryInventory()
	manager := NewParkManager(inventory, mockInspector)

	manager.AddToPark(domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75))
	manager.AddToPark(domain.NewBicycle("B-001", "Trek", 7, 80))
	manager.AddToPark(domain.NewHelmet("H-001", "Шлем M", "M"))
	manager.AddToPark(domain.NewChargingStation("CS-001", "Зарядка", 2.5))
	manager.AddToPark(domain.NewElectricBicycle("EB-001", "Trek", 0.3, 7, 80))

	total := manager.GetTotalEnergy()
	expected := 0.5 + 2.5 + 0.3

	if total != expected {
		t.Errorf("Ожидалась суммарная энергия %f, получено %f", expected, total)
	}
}

func TestParkManager_GetBeginnerFriendly(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	inventory := NewInMemoryInventory()
	manager := NewParkManager(inventory, mockInspector)

	manager.AddToPark(domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75))
	manager.AddToPark(domain.NewBicycle("B-001", "Trek", 5, 80))
	manager.AddToPark(domain.NewElectricBicycle("EB-001", "Trek", 0.3, 7, 80))
	manager.AddToPark(domain.NewHelmet("H-001", "Шлем M", "M"))

	beginners := manager.GetBeginnerFriendly()

	if len(beginners) != 2 {
		t.Errorf("Ожидалось 2 объекта для новичков, получено %d", len(beginners))
	}
}

func TestParkManager_GetAllInventory(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	inventory := NewInMemoryInventory()
	manager := NewParkManager(inventory, mockInspector)

	manager.AddToPark(domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75))
	manager.AddToPark(domain.NewBicycle("B-001", "Trek", 7, 80))
	manager.AddToPark(domain.NewHelmet("H-001", "Шлем M", "M"))

	all := manager.GetAllInventory()

	if len(all) != 3 {
		t.Errorf("Ожидалось 3 объекта в инвентаре, получено %d", len(all))
	}
}

func TestParkManager_GetTotalEnergy_EmptyInventory(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	inventory := NewInMemoryInventory()
	manager := NewParkManager(inventory, mockInspector)

	total := manager.GetTotalEnergy()

	if total != 0 {
		t.Errorf("Ожидалась энергия 0 для пустого инвентаря, получено %f", total)
	}
}

func TestParkManager_GetBeginnerFriendly_EmptyInventory(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	inventory := NewInMemoryInventory()
	manager := NewParkManager(inventory, mockInspector)

	beginners := manager.GetBeginnerFriendly()

	if len(beginners) != 0 {
		t.Errorf("Ожидалось 0 объектов для новичков в пустом инвентаре, получено %d", len(beginners))
	}
}

func TestParkManager_AddToPark_CallsInventoryAdd(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	mockInventory := &MockInventory{}
	manager := NewParkManager(mockInventory, mockInspector)

	scooter := domain.NewElectricScooter("ES-001", "Xiaomi", 0.5, 8, 75)
	manager.AddToPark(scooter)

	if !mockInventory.addCalled {
		t.Error("Метод Add не был вызван у хранилища")
	}
}

func TestParkManager_GetTotalEnergy_CallsInventoryGetAll(t *testing.T) {
	mockInspector := &MockInspector{ShouldPass: true}
	mockInventory := &MockInventory{}
	manager := NewParkManager(mockInventory, mockInspector)

	manager.GetTotalEnergy()

	if !mockInventory.getAllCalled {
		t.Error("Метод GetAll не был вызван у хранилища")
	}
}
