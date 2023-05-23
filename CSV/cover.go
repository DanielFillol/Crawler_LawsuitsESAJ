package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
	"strconv"
)

const fileNameC = "Covers"

func WriteCovers(lawsuits []Crawler.EntireLawsuit) error {
	var rows [][]string

	rows = append(rows, generateCoverHeaders())

	for _, lawsuit := range lawsuits {
		rows = append(rows, tableCoverRows(lawsuit))
	}

	cf, err := createFile(folderName + "/" + fileNameC + ".csv")
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

func generateCoverHeaders() []string {
	return []string{
		"Número do Processo",
		"Aviso",
		"Titulo",
		"Tag",
		"Classe",
		"Assunto",
		"Foro",
		"Vara",
		"Magistrado",
		"Data de Distribuição",
		"Controle",
		"Área",
		"Valor da Causa",
		"Possui Autos",
		"Documento encontrado nos Autos",
		"2 Grau - Titulo",
		"2 Grau - Tag",
		"2 Grau - Classe",
		"2 Grau - Assunto",
		"2 Grau - Foro",
		"2 Grau - Vara",
		"2 Grau - Magistrado",
		"2 Grau - Data de Distribuição",
		"2 Grau - Controle",
		"2 Grau - Área",
		"2 Grau - Valor da Causa",
		"2 Grau - Possui Autos",
		"2 Grau - Documento encontrado nos Autos",
	}
}

func tableCoverRows(results Crawler.EntireLawsuit) []string {
	return []string{
		results.LawsuitNumber,
		results.Warning,
		results.FirstDegree.Cover.Title,
		results.FirstDegree.Cover.Tag,
		results.FirstDegree.Cover.Class,
		results.FirstDegree.Cover.Subject,
		results.FirstDegree.Cover.Location,
		results.FirstDegree.Cover.Unit,
		results.FirstDegree.Cover.Judge,
		results.FirstDegree.Cover.InitialDate,
		results.FirstDegree.Cover.Control,
		results.FirstDegree.Cover.Field,
		results.FirstDegree.Cover.Value,
		strconv.FormatBool(results.FirstDegree.Documents.HasDocuments),
		strconv.FormatBool(results.FirstDegree.Documents.DocumentFound),
		results.SecondDegree.Cover.Title,
		results.SecondDegree.Cover.Tag,
		results.SecondDegree.Cover.Class,
		results.SecondDegree.Cover.Subject,
		results.SecondDegree.Cover.Location,
		results.SecondDegree.Cover.Unit,
		results.SecondDegree.Cover.Judge,
		results.SecondDegree.Cover.InitialDate,
		results.SecondDegree.Cover.Control,
		results.SecondDegree.Cover.Field,
		results.SecondDegree.Cover.Value,
		strconv.FormatBool(results.SecondDegree.Documents.HasDocuments),
		strconv.FormatBool(results.SecondDegree.Documents.DocumentFound),
	}
}
