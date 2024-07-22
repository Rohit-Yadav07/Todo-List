package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todos []Todo
var idCounter int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTodo CLI Application")
		fmt.Println("1. View all todos")
		fmt.Println("2. View a todo by ID")
		fmt.Println("3. Add a new todo")
		fmt.Println("4. Update a todo by ID")
		fmt.Println("5. Delete a todo by ID")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()
		switch choice {
		case "1":
			viewAllTodos()
		case "2":
			viewTodoByID(scanner)
		case "3":
			addNewTodo(scanner)
		case "4":
			updateTodoByID(scanner)
		case "5":
			deleteTodoByID(scanner)
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func viewAllTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}
	for _, todo := range todos {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", todo.ID, todo.Title, todo.Content)
	}
}

func viewTodoByID(scanner *bufio.Scanner) {
	fmt.Print("Enter todo ID: ")
	if !scanner.Scan() {
		return
	}
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			fmt.Printf("ID: %d, Title: %s, Content: %s\n", todo.ID, todo.Title, todo.Content)
			return
		}
	}
	fmt.Println("Todo not found.")
}

func addNewTodo(scanner *bufio.Scanner) {
	fmt.Print("Enter todo title: ")
	if !scanner.Scan() {
		return
	}
	title := scanner.Text()

	fmt.Print("Enter todo content: ")
	if !scanner.Scan() {
		return
	}
	content := scanner.Text()

	idCounter++
	todo := Todo{ID: idCounter, Title: title, Content: content}
	todos = append(todos, todo)
	fmt.Println("Todo added successfully.")
}

func updateTodoByID(scanner *bufio.Scanner) {
	fmt.Print("Enter todo ID: ")
	if !scanner.Scan() {
		return
	}
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	fmt.Print("Enter new todo title: ")
	if !scanner.Scan() {
		return
	}
	title := scanner.Text()

	fmt.Print("Enter new todo content: ")
	if !scanner.Scan() {
		return
	}
	content := scanner.Text()

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = title
			todos[i].Content = content
			fmt.Println("Todo updated successfully.")
			return
		}
	}
	fmt.Println("Todo not found.")
}

func deleteTodoByID(scanner *bufio.Scanner) {
	fmt.Print("Enter todo ID: ")
	if !scanner.Scan() {
		return
	}
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted successfully.")
			return
		}
	}
	fmt.Println("Todo not found.")
}
