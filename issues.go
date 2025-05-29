package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// Важливо: "mygoproject" - це назва нашого модуля (зміни, якщо твоя інша)
	"server1/github"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Будь ласка, вкажіть пошукові терміни.")
		fmt.Println("Приклад: go run issues.go repo:golang/go is:open json decoder")
		os.Exit(1)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	// Визначення часових проміжків
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	// Створення списків для категорій
	var issuesLessThanMonth []*github.Issue
	var issuesLessThanYear []*github.Issue
	var issuesOlderThanYear []*github.Issue

	// Розподіл issues за категоріями
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			issuesLessThanMonth = append(issuesLessThanMonth, item)
		} else if item.CreatedAt.After(oneYearAgo) {
			issuesLessThanYear = append(issuesLessThanYear, item)
		} else {
			issuesOlderThanYear = append(issuesOlderThanYear, item)
		}
	}

	// Виведення категорій, якщо в них є елементи
	if len(issuesLessThanMonth) > 0 {
		fmt.Println("\n--- Менше місяця тому ---")
		for _, item := range issuesLessThanMonth {
			fmt.Printf("#%-5d %9.9s %.55s (створено: %s)\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
		}
	}

	if len(issuesLessThanYear) > 0 {
		fmt.Println("\n--- Менше року тому (але більше місяця) ---")
		for _, item := range issuesLessThanYear {
			fmt.Printf("#%-5d %9.9s %.55s (створено: %s)\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
		}
	}

	if len(issuesOlderThanYear) > 0 {
		fmt.Println("\n--- Більше року тому ---")
		for _, item := range issuesOlderThanYear {
			fmt.Printf("#%-5d %9.9s %.55s (створено: %s)\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
		}
	}
}
