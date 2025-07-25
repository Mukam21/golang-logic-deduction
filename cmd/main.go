package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GumColor string

const (
	RedOnly   GumColor = "only red"
	GreenOnly GumColor = "only green"
	Mixed     GumColor = "mixed"
)

type Machine struct {
	Label  string   // Наклейка на автомате
	Actual GumColor // Фактическое содержимое
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Все возможные варианты содержимого автоматов
	possibleContents := []GumColor{RedOnly, GreenOnly, Mixed}

	// Генерируем все уникальные перестановки
	for _, permutation := range generatePermutations(possibleContents) {
		machines := []Machine{
			{"Red", permutation[0]},
			{"Red and Green", permutation[1]},
			{"Green", permutation[2]},
		}

		if !validateConfiguration(machines) {
			continue
		}

		printSolution(machines, r)
		return
	}

	fmt.Println("Не удалось найти подходящую конфигурацию")
}

// Проверяет, что все наклейки не соответствуют содержимому
func validateConfiguration(machines []Machine) bool {
	for _, m := range machines {
		switch m.Label {
		case "Red":
			if m.Actual == RedOnly {
				return false
			}
		case "Green":
			if m.Actual == GreenOnly {
				return false
			}
		case "Red and Green":
			if m.Actual == Mixed {
				return false
			}
		}
	}
	return true
}

// Покупает жвачку из автомата
func buyGum(machine Machine, r *rand.Rand) GumColor {
	if machine.Actual == Mixed {
		if r.Intn(2) == 0 {
			return RedOnly
		}
		return GreenOnly
	}
	return machine.Actual
}

// Генерирует все возможные перестановки
func generatePermutations(items []GumColor) [][]GumColor {
	var result [][]GumColor
	permute(items, 0, &result)
	return result
}

func permute(arr []GumColor, i int, result *[][]GumColor) {
	if i > len(arr) {
		return
	}
	if i == len(arr)-1 {
		tmp := make([]GumColor, len(arr))
		copy(tmp, arr)
		*result = append(*result, tmp)
		return
	}
	for j := i; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		permute(arr, i+1, result)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Выводит решение
func printSolution(machines []Machine, r *rand.Rand) {
	fmt.Println("Найдена корректная конфигурация автоматов:")
	for _, m := range machines {
		fmt.Printf("- Автомат с наклейкой %-15s фактически содержит: %v\n", "\""+m.Label+"\"", m.Actual)
	}

	// Выбираем автомат "Red and Green"
	targetMachine := machines[1]
	fmt.Printf("\nПокупаем жвачку из автомата %q...\n", targetMachine.Label)

	pulledGum := buyGum(targetMachine, r)
	fmt.Printf("Результат: получили %s жвачку\n", pulledGum)

	fmt.Println("\nОпределяем содержимое всех автоматов:")

	// Определяем содержимое выбранного автомат
	var targetActual GumColor
	if pulledGum == RedOnly {
		targetActual = RedOnly
	} else {
		targetActual = GreenOnly
	}
	fmt.Printf("1. Автомат %q содержит %s (наклейка неверна)\n",
		targetMachine.Label, targetActual)

	// Определяем содержимое остальных автоматов
	remainingTypes := []GumColor{RedOnly, GreenOnly, Mixed}
	remainingTypes = removeItem(remainingTypes, targetActual)

	for _, m := range machines {
		if m.Label == targetMachine.Label {
			continue
		}

		// Для каждого оставшегося автомата
		possible := make([]GumColor, len(remainingTypes))
		copy(possible, remainingTypes)

		// Удаляем недопустимые варианты
		if m.Label == "Red" {
			possible = removeItem(possible, RedOnly)
		} else if m.Label == "Green" {
			possible = removeItem(possible, GreenOnly)
		}

		// Должен остаться ровно один вариант
		if len(possible) == 1 {
			fmt.Printf("%d. Автомат %q содержит %s\n",
				len(machines)-len(remainingTypes)+1, m.Label, possible[0])
			remainingTypes = removeItem(remainingTypes, possible[0])
		}
	}
}

func removeItem(slice []GumColor, item GumColor) []GumColor {
	result := []GumColor{}
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}
