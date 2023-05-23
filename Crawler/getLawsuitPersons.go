package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const (
	xpathPeople1    = "//*[@id=\"tableTodasPartes\"]/tbody/tr"
	xpathPeople2    = "//*[@id=\"tablePartesPrincipais\"]/tbody/tr"
	xpathPole       = "td[1]/span"
	xpathLawyerName = "td[2]/text()"
	Dirt            = "\n"
)

type Person struct {
	Pole    string
	Name    string
	Lawyers []string
}

func GetLawsuitPersons(htmlPgSrc *html.Node) ([]Person, error) {
	totalPersons1 := htmlquery.Find(htmlPgSrc, xpathPeople1)
	totalPersons2 := htmlquery.Find(htmlPgSrc, xpathPeople2)

	if len(totalPersons1) > 0 {
		var personas []Person
		for i, person := range totalPersons1 {
			personas = append(personas, findPerson(person, i))
		}
		return personas, nil
	}

	if len(totalPersons2) > 0 {
		var personas []Person
		for i, person := range totalPersons2 {
			personas = append(personas, findPerson(person, i))
		}
		return personas, nil
	}

	return nil, errors.New("could not find persons")

}

func findPerson(person *html.Node, i int) Person {
	pole := htmlquery.InnerText(htmlquery.FindOne(person, xpathPole))
	name := findName(person, i)

	lawyers := findLawyers(person)

	p := Person{
		Pole:    strings.TrimSpace(pole),
		Name:    name,
		Lawyers: lawyers,
	}

	return p

}

func findName(person *html.Node, i int) string {
	elemNames := htmlquery.Find(person, "//*[@id=\"tableTodasPartes\"]/tbody/tr["+strconv.Itoa(i)+"]/td[2]")
	if len(elemNames) > 1 {
		return strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()["+strconv.Itoa(1)+"]")), Dirt, "", -1))
	} else {
		return strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()")), Dirt, "", -1))
	}
}

func findLawyers(person *html.Node) []string {
	var lawyerNames []string
	elemLawyers := htmlquery.Find(person, xpathLawyerName)
	if len(elemLawyers) > 2 {
		for i := 1; i < len(elemLawyers); i++ {
			name := strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()["+strconv.Itoa(i+1)+"]")), Dirt, "", -1))
			lawyerNames = append(lawyerNames, name)
		}
	} else {
		lawyerNames = append(lawyerNames, "no lawyer found")
	}

	return lawyerNames
}
