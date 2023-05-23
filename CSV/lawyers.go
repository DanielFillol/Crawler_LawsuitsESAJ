package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
)

const fileNameL = "Lawyers"

func WriteLawyers(lawsuits []Crawler.EntireLawsuit) error {
	var rows [][]string

	rows = tableLawyerRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameL + ".csv")
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

func tableLawyerRows(lawsuits []Crawler.EntireLawsuit) [][]string {
	var prs [][]string

	prs = append(prs, []string{"Processo", "Grau", "Polo", "Nome Parte", "Advogado"})

	for _, lawsuit := range lawsuits {
		for _, person1 := range lawsuit.FirstDegree.Persons {
			for _, lawyer1 := range person1.Lawyers {
				if lawyer1 != "" {
					prs = append(prs, []string{lawsuit.LawsuitNumber, "primeiro", person1.Pole, person1.Name, lawyer1})
				}
			}
		}
		for _, person2 := range lawsuit.SecondDegree.Persons {
			for _, lawyer2 := range person2.Lawyers {
				if lawyer2 != "" {
					prs = append(prs, []string{lawsuit.LawsuitNumber, "segundo", person2.Pole, person2.Name, lawyer2})
				}
			}
		}

	}

	return prs
}
