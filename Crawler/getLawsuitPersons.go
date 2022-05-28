package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const (
	xpathPeople     = "//*[@id=\"tablePartesPrincipais\"]/tbody/tr"
	xpathPole       = "td[1]/span"
	xpathPersonName = "td[2]/span"
	xpathLawyerName = "td[2]/text()"
	Dirt            = "\n"
)

type Person struct {
	Pole    string
	Names   []string
	Lawyers []string
}

func GetLawsuitPersons(htmlPgSrc *html.Node) ([]Person, error) {
	totalPersons := htmlquery.Find(htmlPgSrc, xpathPeople)

	if len(totalPersons) > 0 {
		var personas []Person
		for _, person := range totalPersons {
			personas = append(personas, findPerson(person))
		}
		return personas, nil
	}

	return nil, errors.New("could not find persons")

}

func findPerson(person *html.Node) Person {
	pole := htmlquery.InnerText(htmlquery.FindOne(person, xpathPole))

	names := htmlquery.Find(person, xpathPersonName)
	personNames := findNames(names, person)

	lawyers := htmlquery.Find(person, xpathLawyerName)
	lawyerNames := findLawyers(lawyers, person)
	clearLawyer := clearLawyerName(personNames, lawyerNames)

	return Person{
		Pole:    strings.TrimSpace(pole),
		Names:   personNames,
		Lawyers: clearLawyer,
	}

}

func findNames(names []*html.Node, person *html.Node) []string {
	var personNames []string
	if len(names) > 1 {
		for i := 0; i < len(names); i++ {
			name := strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()["+strconv.Itoa(i+1)+"]")), Dirt, "", -1))
			if name != "" {
				personNames = append(personNames, name)
			}
		}
	} else {
		name := strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()")), Dirt, "", -1))
		if name != "" {
			personNames = append(personNames, name)
		}
	}
	return personNames
}

func findLawyers(lawyers []*html.Node, person *html.Node) []string {
	var lawyerNames []string
	if len(lawyers) > 1 {
		for i := 0; i < len(lawyers); i++ {
			name := strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()["+strconv.Itoa(i+1)+"]")), Dirt, "", -1))
			lawyerNames = append(lawyerNames, name)
		}
	} else {
		name := strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(person, "td[2]/text()")), Dirt, "", -1))
		lawyerNames = append(lawyerNames, name)
	}
	return lawyerNames
}

func clearLawyerName(names []string, lawyerNames []string) []string {
	var clearNames []string
	for _, lawyer := range lawyerNames {
		if lawyer != "" {
			for _, name := range names {
				if name != "" {
					if lawyer != name {
						clearNames = append(clearNames, lawyer)
					}
				}
			}
		}
	}
	return clearNames
}
