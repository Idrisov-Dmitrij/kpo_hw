package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"wishcat-app/internal/di"
	"wishcat-app/internal/domain"
	"wishcat-app/internal/service"
)

func main() {
	container := di.NewContainer()
	manager := container.BuildParkManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		fmt.Println()

		switch choice {
		case "1":
			addElectricScooter(manager, scanner)
		case "2":
			addBicycle(manager, scanner)
		case "3":
			addElectricBicycle(manager, scanner)
		case "4":
			addHelmet(manager, scanner)
		case "5":
			addChargingStation(manager, scanner)
		case "6":
			showTotalEnergy(manager)
		case "7":
			showBeginnerFriendly(manager)
		case "8":
			showAllInventory(manager)
		case "9":
			runDemoScenario(manager)
		case "0":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("--- Меню ВышКат ---")
	fmt.Println("1. Добавить электросамокат")
	fmt.Println("2. Добавить велосипед")
	fmt.Println("3. Добавить электровелосипед")
	fmt.Println("4. Добавить шлем")
	fmt.Println("5. Добавить зарядную станцию")
	fmt.Println("6. Показать суммарное энергопотребление")
	fmt.Println("7. Показать технику для новичков")
	fmt.Println("8. Показать весь инвентарь")
	fmt.Println("9. Запустить демо-сценарий")
	fmt.Println("0. Выход")
	fmt.Print("Выберите действие: ")
}

func addElectricScooter(manager *service.ParkManager, scanner *bufio.Scanner) {
	invNum := readString(scanner, "Инвентарный номер")
	model := readString(scanner, "Модель")
	energy := readFloat(scanner, "Суточный расход энергии (кВт*ч)")
	simplicity := readInt(scanner, "Простота для новичка (1-10)")
	durability := readInt(scanner, "Прочность (0-100)")

	addToPark(manager, domain.NewElectricScooter(invNum, model, energy, simplicity, durability))
}

func addBicycle(manager *service.ParkManager, scanner *bufio.Scanner) {
	invNum := readString(scanner, "Инвентарный номер")
	model := readString(scanner, "Модель")
	simplicity := readInt(scanner, "Простота для новичка (1-10)")
	durability := readInt(scanner, "Прочность (0-100)")

	addToPark(manager, domain.NewBicycle(invNum, model, simplicity, durability))
}

func addElectricBicycle(manager *service.ParkManager, scanner *bufio.Scanner) {
	invNum := readString(scanner, "Инвентарный номер")
	model := readString(scanner, "Модель")
	energy := readFloat(scanner, "Суточный расход энергии (кВт*ч)")
	simplicity := readInt(scanner, "Простота для новичка (1-10)")
	durability := readInt(scanner, "Прочность (0-100)")

	addToPark(manager, domain.NewElectricBicycle(invNum, model, energy, simplicity, durability))
}

func addHelmet(manager *service.ParkManager, scanner *bufio.Scanner) {
	invNum := readString(scanner, "Инвентарный номер")
	model := readString(scanner, "Наименование")
	size := readString(scanner, "Размер (S/M/L)")

	addToPark(manager, domain.NewHelmet(invNum, model, size))
}

func addChargingStation(manager *service.ParkManager, scanner *bufio.Scanner) {
	invNum := readString(scanner, "Инвентарный номер")
	model := readString(scanner, "Модель")
	energy := readFloat(scanner, "Суточный расход энергии (кВт*ч)")

	addToPark(manager, domain.NewChargingStation(invNum, model, energy))
}

func addToPark(manager *service.ParkManager, item domain.Item) {
	err := manager.AddToPark(item)
	if err != nil {
		fmt.Printf("Ошибка: %v. Устройство не принято.\n", err)
	} else {
		fmt.Printf("Успешно: %s (№%s) принят в парк.\n", item.GetName(), item.GetInventoryNumber())
	}
}

func showTotalEnergy(manager *service.ParkManager) {
	total := manager.GetTotalEnergy()
	items := manager.GetAllInventory()

	energyCount := 0
	for _, item := range items {
		if _, ok := item.(domain.EnergyConsumer); ok {
			energyCount++
		}
	}

	fmt.Printf("Всего единиц на балансе: %d\n", len(items))
	fmt.Printf("Энергопотребителей: %d\n", energyCount)
	fmt.Printf("Суммарное суточное потребление: %.2f кВт*ч\n", total)
}

func showBeginnerFriendly(manager *service.ParkManager) {
	beginners := manager.GetBeginnerFriendly()

	if len(beginners) == 0 {
		fmt.Println("В парке нет техники, подходящей для новичков.")
		return
	}

	fmt.Printf("Найдено единиц: %d\n", len(beginners))
	fmt.Printf("%-15s %-25s %-10s %-10s\n", "№ Инвентарный", "Наименование", "Простота", "Прочность")
	fmt.Println(strings.Repeat("-", 65))

	for _, item := range beginners {
		vehicle := item.(domain.Vehicle)
		fmt.Printf("%-15s %-25s %-10d %-10d\n",
			item.GetInventoryNumber(),
			item.GetName(),
			vehicle.GetSimplicity(),
			vehicle.GetDurability(),
		)
	}
}

func showAllInventory(manager *service.ParkManager) {
	items := manager.GetAllInventory()

	if len(items) == 0 {
		fmt.Println("Парк пуст.")
		return
	}

	fmt.Printf("Всего единиц: %d\n", len(items))
	fmt.Printf("%-15s %-25s %-18s\n", "№ Инвентарный", "Наименование", "Тип")
	fmt.Println(strings.Repeat("-", 62))

	for _, item := range items {
		fmt.Printf("%-15s %-25s %-18s\n",
			item.GetInventoryNumber(),
			item.GetName(),
			determineType(item),
		)
	}
}

func determineType(item domain.Item) string {
	_, isVehicle := item.(domain.Vehicle)
	_, isEnergy := item.(domain.EnergyConsumer)

	switch {
	case isVehicle && isEnergy:
		return "Электротранспорт"
	case isVehicle && !isEnergy:
		return "Транспорт"
	case !isVehicle && isEnergy:
		return "Энергоустройство"
	default:
		return "Вещь"
	}
}

func runDemoScenario(manager *service.ParkManager) {
	fmt.Println("--- Запуск демо-сценария ---")

	fmt.Println("\n1. Приёмка исправной техники:")
	addToPark(manager, domain.NewElectricScooter("ES-001", "Xiaomi Mi Pro", 0.5, 8, 75))
	addToPark(manager, domain.NewBicycle("B-001", "Trek FX1", 7, 85))
	addToPark(manager, domain.NewHelmet("H-001", "Шлем городской", "M"))
	addToPark(manager, domain.NewChargingStation("CS-001", "Зарядка Pro", 2.5))

	fmt.Println("\n2. Попытки добавить бракованную технику (ожидаются ошибки):")
	addToPark(manager, domain.NewElectricScooter("", "Без номера", 0.5, 8, 75))
	addToPark(manager, domain.NewElectricScooter("ES-002", "Сломанный", -1.0, 8, 75))
	addToPark(manager, domain.NewElectricScooter("ES-003", "Странный", 0.5, 15, 75))
	addToPark(manager, domain.NewElectricScooter("ES-001", "Дубликат номера", 0.5, 8, 75))

	fmt.Println("\n3. Отчёт по энергопотреблению:")
	showTotalEnergy(manager)

	fmt.Println("\n4. Техника для новичков:")
	showBeginnerFriendly(manager)

	fmt.Println("\n5. Полный инвентарь:")
	showAllInventory(manager)

	fmt.Println("\nДемо-сценарий завершён.")
}

func readString(scanner *bufio.Scanner, prompt string) string {
	fmt.Printf("%s: ", prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readFloat(scanner *bufio.Scanner, prompt string) float64 {
	for {
		fmt.Printf("%s: ", prompt)
		scanner.Scan()
		val, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err == nil {
			return val
		}
		fmt.Println("Ошибка: введите число.")
	}
}

func readInt(scanner *bufio.Scanner, prompt string) int {
	for {
		fmt.Printf("%s: ", prompt)
		scanner.Scan()
		val, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err == nil {
			return val
		}
		fmt.Println("Ошибка: введите целое число.")
	}
}
