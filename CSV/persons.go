package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
)

const fileNameP = "Persons"

func WritePersons(lawsuits []Crawler.EntireLawsuit) error {
	var rows [][]string

	rows = tablePersonsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameP + ".csv")
	if err != nil {
		return err
	}

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func tablePersonsRows(lawsuits []Crawler.EntireLawsuit) [][]string {
	var prs [][]string
	prs = append(prs, []string{"Processo", "Grau", "Polo", "Nome"})

	for _, lawsuit := range lawsuits {
		for _, person1 := range lawsuit.FirstDegree.Persons {
			prs = append(prs, []string{lawsuit.LawsuitNumber, "primeiro", person1.Pole, person1.Name})
		}
		for _, person2 := range lawsuit.SecondDegree.Persons {
			prs = append(prs, []string{lawsuit.LawsuitNumber, "segundo", person2.Pole, person2.Name})
		}

	}

	return prs
}
