package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
)

const fileNamepP = "Poles"

func WritePoles(lawsuits []Crawler.EntireLawsuit) error {
	var rows [][]string

	rows = tablePolesRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNamepP + ".csv")
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

func tablePolesRows(lawsuits []Crawler.EntireLawsuit) [][]string {
	var poles [][]string

	poles = append(poles, []string{"Processo", "Grau", "Type", "Name", "Document"})

	//fmt.Println(len(lawsuits[0].FirstDegree.Pole.Active))
	//fmt.Println(len(lawsuits[0].FirstDegree.Pole.Passive))
	//fmt.Println(len(lawsuits[0].FirstDegree.Pole.Other))
	//
	//fmt.Println(len(lawsuits[0].SecondDegree.Pole.Active))
	//fmt.Println(len(lawsuits[0].SecondDegree.Pole.Passive))
	//fmt.Println(len(lawsuits[0].SecondDegree.Pole.Other))

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].FirstDegree.Pole.Active); j++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "primeiro grau", lawsuits[i].FirstDegree.Pole.Active[j].Type, lawsuits[i].FirstDegree.Pole.Active[j].Name, lawsuits[i].FirstDegree.Pole.Active[j].Document})
		}
		for k := 0; k < len(lawsuits[i].FirstDegree.Pole.Passive); k++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "primeiro grau", lawsuits[i].FirstDegree.Pole.Passive[k].Type, lawsuits[i].FirstDegree.Pole.Passive[k].Name, lawsuits[i].FirstDegree.Pole.Passive[k].Document})
		}
		for l := 0; l < len(lawsuits[i].FirstDegree.Pole.Other); l++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "primeiro grau", lawsuits[i].FirstDegree.Pole.Other[l].Type, lawsuits[i].FirstDegree.Pole.Other[l].Name, lawsuits[i].FirstDegree.Pole.Other[l].Document})
		}
		for m := 0; m < len(lawsuits[i].SecondDegree.Pole.Active); m++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "segundo grau", lawsuits[i].SecondDegree.Pole.Active[m].Type, lawsuits[i].SecondDegree.Pole.Active[m].Name, lawsuits[i].SecondDegree.Pole.Active[m].Document})
		}
		for n := 0; n < len(lawsuits[i].SecondDegree.Pole.Passive); n++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "segundo grau", lawsuits[i].SecondDegree.Pole.Passive[n].Type, lawsuits[i].SecondDegree.Pole.Passive[n].Name, lawsuits[i].SecondDegree.Pole.Passive[n].Document})
		}
		for o := 0; o < len(lawsuits[i].SecondDegree.Pole.Other); o++ {
			poles = append(poles, []string{lawsuits[i].LawsuitNumber, "segundo grau", lawsuits[i].SecondDegree.Pole.Other[o].Type, lawsuits[i].SecondDegree.Pole.Other[o].Name, lawsuits[i].SecondDegree.Pole.Other[o].Document})
		}
	}
	return poles
}
