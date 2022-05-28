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

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func tableLawyerRows(lawsuits []Crawler.EntireLawsuit) [][]string {
	var prs [][]string

	prs = append(prs, []string{"Processo", "Grau", "Polo", "Advogado"})

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].FirstDegree.Persons); j++ {
			for k := 0; k < len(lawsuits[i].FirstDegree.Persons[j].Names); k++ {
				for l := 0; l < len(lawsuits[i].FirstDegree.Persons[j].Lawyers); l++ {
					prs = append(prs, []string{lawsuits[i].LawsuitNumber, "primeiro", lawsuits[i].FirstDegree.Persons[j].Pole, lawsuits[i].FirstDegree.Persons[j].Lawyers[l]})
				}
			}
		}
		for j := 0; j < len(lawsuits[i].SecondDegree.Persons); j++ {
			for k := 0; k < len(lawsuits[i].SecondDegree.Persons[j].Names); k++ {
				for l := 0; l < len(lawsuits[i].SecondDegree.Persons[j].Lawyers); l++ {
					prs = append(prs, []string{lawsuits[i].LawsuitNumber, "segundo", lawsuits[i].SecondDegree.Persons[j].Pole, lawsuits[i].SecondDegree.Persons[j].Lawyers[l]})
				}
			}
		}
	}
	return prs
}
