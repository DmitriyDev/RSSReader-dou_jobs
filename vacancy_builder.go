package main

import "strings"

type Vacancy struct {
	Title       string
	Link        string
	Text        string
	Salary      string
	Cities      []string
	Company     string
	PublishDate string
	Guid        string
}

type VacancyBuilder struct {
	Items []Item
}

func (vb VacancyBuilder) extractVacancies() []Vacancy {
	vacancies := []Vacancy{}

	var vacancy Vacancy
	var title string
	var company string
	var salary string
	var cities []string

	for _, item := range vb.Items {
		title, company, salary, cities = vb.parseTitle(item.Title)

		vacancy = Vacancy{
			Title:       title,
			Link:        item.Link,
			Text:        item.Description,
			Salary:      salary,
			Cities:      cities,
			Company:     company,
			PublishDate: item.PubDate,
			Guid:        item.Guid,
		}

		vacancies = append(vacancies, vacancy)

	}
	return vacancies
}

func (vb VacancyBuilder) parseTitle(baseTitle string) (string, string, string, []string) {
	data := strings.Split(baseTitle, ",")

	titleStr := data[0]
	data = data[1:]
	if !strings.Contains(titleStr, " в ") {
		titleStr += ", " + data[0]
		data = data[1:]
	}

	t := strings.Split(titleStr, " в ")
	title, company := strings.TrimSpace(t[0]), strings.TrimSpace(t[1])

	salary := ""
	if strings.Contains(data[0], "$") {
		salary = strings.TrimSpace(data[0])
		data = data[1:]
	}

	cities := []string{}
	for _, city := range data {
		cities = append(cities, strings.TrimSpace(city))
	}

	return title, company, salary, cities
}
