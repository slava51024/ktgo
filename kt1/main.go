package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// ==================== ЗАДАЧА №1 ====================

// square вычисляет квадрат числа и отправляет результат в канал.
func square(n int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()

	ch <- n * n
}

func task1() {
	fmt.Println("Задача №1")

	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	results := make(chan int, len(numbers))

	for _, n := range numbers {
		wg.Add(1)
		go square(n, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println()
}

// ==================== ЗАДАЧА №2 ====================

// groupTemperatures объединяет температуры в группы с шагом 10 градусов.
func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temperature := range temperatures {
		// Для отрицательных температур правильно определяем
		// начало десятилетнего интервала.
		group := int(temperature/10) * 10

		if temperature < 0 && int(temperature)%10 != 0 {
			group -= 10
		}

		groups[group] = append(groups[group], temperature)
	}

	return groups
}

func task2() {
	fmt.Println("Задача №2")

	temperatures := []float64{
		-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5,
	}

	groups := groupTemperatures(temperatures)

	for group, values := range groups {
		fmt.Printf("%d: %v\n", group, values)
	}

	fmt.Println()
}

// ==================== ЗАДАЧА №3 ====================

// reverseString переворачивает строку с поддержкой Unicode.
func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func task3() {
	fmt.Println("Задача №3")

	text := "главрыба — абырвалг"

	fmt.Println(reverseString(text))
	fmt.Println()
}

// ==================== ЗАДАЧА №4 ====================

// reverseWords переворачивает порядок слов в строке.
func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func task4() {
	fmt.Println("Задача №4")

	text := "snow dog sun"

	fmt.Println(reverseWords(text))
	fmt.Println()
}

// ==================== ЗАДАЧА №5 ====================

// allUnique проверяет, что все символы строки уникальны.
// Проверка является регистронезависимой.
func allUnique(s string) bool {
	seen := make(map[rune]bool)

	for _, r := range s {
		r = unicode.ToLower(r)

		if seen[r] {
			return false
		}

		seen[r] = true
	}

	return true
}

func task5() {
	fmt.Println("Задача №5")

	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, text := range tests {
		fmt.Printf("%s — %t\n", text, allUnique(text))
	}

	fmt.Println()
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
}
