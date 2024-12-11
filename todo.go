package main

import (
	"fmt"
	"net/http"
)

var t = []string{"Study for CIE", "Write Datasheets", "Finish EL", "Finish project", "sleep", "play"}

func main() {
	// fmt.Println("Study for CIE")
	// fmt.Println("Write Datasheets")
	// fmt.Println("Finish EL")
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

	printTask(t)
}
func helloUser(writer http.ResponseWriter, request *http.Request) {
	g := "Welcome to the To-Do list application"
	fmt.Fprintln(writer, g)
}
func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, i := range t {
		fmt.Fprintln(writer, i)
	}

}
func printTask(tasks []string) {
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func addTask(tasks []string, t string) {
	tasks = append(tasks, t)
	printTask(tasks)
}
