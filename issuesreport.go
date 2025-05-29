// Файл: mygoproject/issuesreport.go (або як ти назвав свій проект)
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"log"
	"os"
	"text/template"
	"time"

	// Важливо: "mygoproject" - це назва нашого модуля.
	// Якщо ти назвав модуль інакше, зміни тут "mygoproject" на свою назву.
	"server1/github" // Використовуємо той самий пакет github, що й раніше
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Ініціалізуємо та парсимо шаблон один раз при запуску програми
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}). // Реєструємо нашу функцію daysAgo
	Parse(templ))                                // Парсимо шаблон templ

func main() {
	if len(os.Args) <= 1 {
		log.Println("Будь ласка, вкажіть пошукові терміни.")
		log.Println("Приклад: go run issuesreport.go repo:golang/go is:open json decoder")
		os.Exit(1)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Виконуємо шаблон, передаючи дані (result) та куди виводити (os.Stdout)
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// Функція noMust (не використовується в main, але показує альтернативний спосіб обробки помилок парсингу)
// func noMust() {
// 	report, err := template.New("report").
// 		Funcs(template.FuncMap{"daysAgo": daysAgo}).
// 		Parse(templ)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	result, err := github.SearchIssues(os.Args[1:])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := report.Execute(os.Stdout, result); err != nil {
// 		log.Fatal(err)
// 	}
// }
