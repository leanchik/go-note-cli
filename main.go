package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title string
	Body  string
}

func main() {
	var notes []Note
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню")
		fmt.Println("1. Добавить заметку")
		fmt.Println("2. Показать заголовки")
		fmt.Println("3. Показать текст заметки")
		fmt.Println("4. Удалить заметку")
		fmt.Println("5. Выход")

		fmt.Print("Выберите пункт: ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Введите заголовок заметки: ")
			titleInput, _ := reader.ReadString('\n')
			title := strings.TrimSpace(titleInput)

			fmt.Print("Введите текст заметки: ")
			bodyInput, _ := reader.ReadString('\n')
			body := strings.TrimSpace(bodyInput)

			newNote := Note{title, body}

			notes = append(notes, newNote)
			fmt.Println("Заметка добавлена!")

		case "2":
			if len(notes) == 0 {
				fmt.Println("Список заметок пуст.")
			} else {
				fmt.Println("Заголовки заметок: ")
				for i, note := range notes {
					fmt.Printf("%d. %s\n", i+1, note.Title)
				}
			}

		case "3":
			if len(notes) == 0 {
				fmt.Println("Некорректный номер заметки.")
				continue
			}
			fmt.Println("Введите номер заметки: ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)

			var index int
			_, err := fmt.Sscan(numStr, &index)
			if err != nil || index < 1 || index > len(notes) {
				fmt.Println("Некорретный номер.")
				continue
			}
			note := notes[index-1]
			fmt.Println("Заголовок: ", note.Title)
			fmt.Println("Текст: ", note.Body)

		case "4":
			if len(notes) == 0 {
				fmt.Println("Список заметок пуст.")
			}
			fmt.Println("Введите номер заметки для удаления: ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)

			var index int
			_, err := fmt.Sscan(numStr, &index)
			if err != nil || index < 1 || index > len(notes) {
				fmt.Println("Некрректный номер заметки.")
				continue
			}
			notes = append(notes[:index-1], notes[index:]...)
			fmt.Println("Заметка удалена.")

		case "5":
			fmt.Println("Пока!")
			return
		}
	}
}
