package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
)

const fileNameM = "Movements"

func WriteMovements(lawsuits []Crawler.EntireLawsuit) error {
	var rows [][]string

	rows = tableMovementsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameM + ".csv")
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

func tableMovementsRows(lawsuits []Crawler.EntireLawsuit) [][]string {
	var mvts [][]string

	mvts = append(mvts, []string{"Processo", "Grau", "Data", "Título", "Texto"})

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].FirstDegree.Movements); j++ {
			mvts = append(mvts, []string{lawsuits[i].LawsuitNumber, "primeiro grau", lawsuits[i].FirstDegree.Movements[j].Date, lawsuits[i].FirstDegree.Movements[j].Title, lawsuits[i].FirstDegree.Movements[j].Text})
		}
		for k := 0; k < len(lawsuits[i].SecondDegree.Movements); k++ {
			mvts = append(mvts, []string{lawsuits[i].LawsuitNumber, "segundo grau", lawsuits[i].SecondDegree.Movements[k].Date, lawsuits[i].SecondDegree.Movements[k].Title, lawsuits[i].SecondDegree.Movements[k].Text})
		}
	}
	return mvts
}
