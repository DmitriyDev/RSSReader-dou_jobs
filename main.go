package main

import (
	"fmt"
	crss "github.com/DmitriyDev/go-RssCommunicator"
)

const Domain = "https://jobs.dou.ua/vacancies/feeds/"

var Categories = map[int]string{1: ".NET", 2: "1С", 3: "Analyst", 4: "Android", 5: "Blockchain", 6: "C++",
	7: "Data Science", 8: "DBA", 9: "DevOps", 10: "Embedded", 11: "ERP/CRM", 12: "Front End",
	13: "Golang", 14: "HR", 15: "iOS/macOS", 16: "Java", 17: "Node.js", 18: "Other",
	19: "PHP", 20: "Product Manager", 21: "Project Manager", 22: "Python", 23: "QA",
	24: "React Native", 25: "Ruby", 26: "Sales", 27: "Scala", 28: "Security", 29: "SEO",
	30: "Support", 31: "Technical Writer", 32: "Unity", 33: "Дизайн", 34: "Маркетинг",
	35: "Системный администратор",
}

func main() {

	rssFacade := RssFacade{
		Domain,
		crss.Communicator{},
	}

	stream := RssStream{"Киев", Categories[13]}

	items := rssFacade.ReadStream(stream)
	vacancies := VacancyBuilder{items}.extractVacancies()
	for _, vacancy := range vacancies {
		fmt.Printf("%v\t %s\t %s\t %s\n", vacancy.Cities, vacancy.Company, vacancy.Title, vacancy.Salary)
	}

}
