package Crawler

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

const (
	xpathTitleP    = "//*[@id=\"containerDadosPrincipaisProcesso\"]/div[1]/div/div/span[1]"
	xpathTitleS    = "//*[@id=\"numeroProcesso\"]"
	xpathTagP      = "//*[@id=\"containerDadosPrincipaisProcesso\"]/div[1]/div/div/span[2]"
	xpathTagS      = "//*[@id=\"situacaoProcesso\"]"
	xpathClassP    = "//*[@id=\"classeProcesso\"]"
	xpathClassS    = "//*[@id=\"classeProcesso\"]/span"
	xpathSubjectP  = "//*[@id=\"assuntoProcesso\"]"
	xpathSubjectS  = "//*[@id=\"assuntoProcesso\"]/span"
	xpathLocationP = "//*[@id=\"foroProcesso\"]"
	xpathLocationS = "//*[@id=\"secaoProcesso\"]/span"
	xpathUnitP     = "//*[@id=\"varaProcesso\"]"
	xpathUnitS     = "//*[@id=\"orgaoJulgadorProcesso\"]/span"
	xpathJudgeP    = "//*[@id=\"juizProcesso\"]"
	xpathJudgeS    = "//*[@id=\"relatorProcesso\"]/span"
	xpathInitDate  = "//*[@id=\"dataHoraDistribuicaoProcesso\"]"
	xpathControl   = "//*[@id=\"numeroControleProcesso\"]"
	xpathFieldP    = "//*[@id=\"areaProcesso\"]/span"
	xpathFieldS    = "//*[@id=\"areaProcesso\"]/span"
	xpathValueP    = "//*[@id=\"valorAcaoProcesso\"]"
	xpathValueS    = "//*[@id=\"valorAcaoProcesso\"]/span"
)

type LawsuitCover struct {
	Title       string
	Tag         string
	Class       string
	Subject     string
	Location    string
	Unit        string
	Judge       string
	InitialDate string
	Control     string
	Field       string
	Value       string
}

func GetLawsuitCover(htmlPgSrc *html.Node, degree string) (LawsuitCover, error) {
	var title string
	var tag string
	var class string
	var subject string
	var location string
	var unit string
	var judge string
	var initDate string
	var control string
	var field string
	var value string

	if degree == "p" {
		existTitle := htmlquery.Find(htmlPgSrc, xpathTitleP)
		if len(existTitle) > 0 {
			title = strings.Replace(htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathTitleP)), "                                                            ", "", -1)
		}

		existTag := htmlquery.Find(htmlPgSrc, xpathTagP)
		if len(existTag) > 0 {
			tag = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathTagP))
		}

		existClass := htmlquery.Find(htmlPgSrc, xpathClassP)
		if len(existClass) > 0 {
			class = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathClassP))
		}

		existSubject := htmlquery.Find(htmlPgSrc, xpathSubjectP)
		if len(existSubject) > 0 {
			subject = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathSubjectP))
		}

		existLocation := htmlquery.Find(htmlPgSrc, xpathLocationP)
		if len(existLocation) > 0 {
			location = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathLocationP))
		}

		existUnit := htmlquery.Find(htmlPgSrc, xpathUnitP)
		if len(existUnit) > 0 {
			unit = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathUnitP))
		}

		existJudge := htmlquery.Find(htmlPgSrc, xpathJudgeP)
		if len(existJudge) > 0 {
			judge = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathJudgeP))
		}

		existInitDate := htmlquery.Find(htmlPgSrc, xpathInitDate)
		if len(existInitDate) > 0 {
			initDate = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathInitDate))
		}

		existControl := htmlquery.Find(htmlPgSrc, xpathControl)
		if len(existControl) > 0 {
			control = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathControl))
		}

		existField := htmlquery.Find(htmlPgSrc, xpathFieldP)
		if len(existField) > 0 {
			field = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathFieldP))
		}

		existValueOrigin := htmlquery.Find(htmlPgSrc, xpathValueP)
		if len(existValueOrigin) > 0 {
			valueOrigin := htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathValueP))
			vOrinSplit := strings.Split(valueOrigin, "R$         ")
			if len(vOrinSplit) > 1 {
				value = vOrinSplit[1]
			} else {
				value = vOrinSplit[0]
			}

		}

	} else {
		existTitle := htmlquery.Find(htmlPgSrc, xpathTitleS)
		if len(existTitle) > 0 {
			title = strings.Replace(htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathTitleS)), "                        ", "", -1)
		}

		existTag := htmlquery.Find(htmlPgSrc, xpathTagS)
		if len(existTag) > 0 {
			tag = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathTagS))
		}

		existClass := htmlquery.Find(htmlPgSrc, xpathClassS)
		if len(existClass) > 0 {
			class = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathClassS))
		}

		existSubject := htmlquery.Find(htmlPgSrc, xpathSubjectS)
		if len(existSubject) > 0 {
			subject = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathSubjectS))
		}

		existLocation := htmlquery.Find(htmlPgSrc, xpathLocationS)
		if len(existLocation) > 0 {
			location = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathLocationS))
		}

		existUnit := htmlquery.Find(htmlPgSrc, xpathUnitS)
		if len(existUnit) > 0 {
			unit = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathUnitS))
		}

		existJudge := htmlquery.Find(htmlPgSrc, xpathJudgeS)
		if len(existJudge) > 0 {
			judge = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathJudgeS))
		}

		existField := htmlquery.Find(htmlPgSrc, xpathFieldS)
		if len(existField) > 0 {
			field = htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathFieldS))
		}

		existValueOrigin := htmlquery.Find(htmlPgSrc, xpathValueS)
		if len(existValueOrigin) > 0 {
			valueOrigin := htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathValueS))
			vOrinSplit := strings.Split(valueOrigin, "R$         ")
			if len(vOrinSplit) > 1 {
				value = vOrinSplit[1]
			} else {
				value = vOrinSplit[0]
			}
		}
	}

	return LawsuitCover{
		Title:       strings.Replace(strings.Replace(title, Dirt, "", -1), "                                                            ", "", -1),
		Tag:         tag,
		Class:       class,
		Subject:     subject,
		Location:    location,
		Unit:        unit,
		Judge:       judge,
		InitialDate: initDate,
		Control:     control,
		Field:       field,
		Value:       value,
	}, nil

}
