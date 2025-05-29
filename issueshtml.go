package main

import (
	"html/template"
	"log"
	"os"

	"server1/github"
)

var issueListTemplate = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} Issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	if len(os.Args) <= 1 {
		log.Println("Будь ласка, вкажіть пошукові терміни.")
		log.Println("Приклад: go run issueshtml.go repo:golang/go commenter:gopherbot json encoder")
		os.Exit(1)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("issues.html")
	if err != nil {
		log.Fatalf("Помилка створення файлу issues.html: %v", err)
	}
	defer outputFile.Close()

	if err := issueListTemplate.Execute(outputFile, result); err != nil {
		log.Fatal(err)
	}
	log.Println("HTML звіт успішно згенеровано у файл issues.html")
}
