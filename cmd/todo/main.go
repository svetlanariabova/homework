// Command todo — точка входа менеджера задач, сквозного проекта домашних
// заданий курса Go Basic.
package main

import (
	"bufio"
	"fmt"
	"os"
)

const tasksCount = 3

var (
	ids      = [tasksCount]int{1, 2, 3}
	titles   = [tasksCount]string{"Задача1", "Задача2", "Задача3"}
	statuses = [tasksCount]string{"Новая", "В работе", "Завершена"}
	action   = []string{"Показать все", "Фильтровать по статусу", "Выйти"}
)

func showMenu() {
	for index, value := range action {
		fmt.Printf("%-4d %-20s \n", index+1, value)
	}
	fmt.Printf("--------------------------- \n")
}

func showCatalog() {
	fmt.Printf("%-4s %-20s %s\n", "ID", "НАЗВАНИЕ", "СТАТУС")
	for i := range action {
		fmt.Printf("%-4d %-20s %s\n", ids[i], titles[i], statuses[i])
	}
}

func countTasksResult() {
	created := 0
	inProgress := 0
	completed := 0
	for i := range tasksCount {
		switch statuses[i] {
		case "Новая":
			created++
		case "В работе":
			inProgress++
		case "Завершена":
			completed++
		}
	}
	fmt.Printf("Новая: %d В работе: %d Завершена: %d\n", created, completed, inProgress)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	showMenu()
menu:
	for {
		fmt.Print("Выберите пункт меню: ")
		if !scanner.Scan() {
			break menu // ввод закончился
		}
		command := scanner.Text()

		switch command {
		case "1":
			showCatalog()
		case "2":
			countTasksResult()
		case "3":
			fmt.Println("Выход.")
			break menu
		default:
			fmt.Printf("Неизвестная команда %q\n", command)
			fmt.Printf("Допустимые команды: \n")
			showMenu()
		}
	}

}
