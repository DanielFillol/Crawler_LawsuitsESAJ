package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"log"
	"strings"
	"time"
)

const (
	btnPetition      = "//*[@id=\"pbPeticionar\"]"
	titleActivePole  = "Polo ativo"
	titlePassivePole = "Polo passivo"
	titleOthersPole  = "Outras participações "
	xpathTitleInit   = "/html/body/span/main/div/div/div/form/div/div/div[2]/div/div/section"
	xpathNameInit    = "/div/div[2]/div/div/ng-include/div/div/div/div/div"
	xpathName        = "/span[1]"
	xpathDocument    = "/span[2]"
)

type Poles struct {
	Active  []Part
	Passive []Part
	Other   []Part
}

type Part struct {
	Type     string
	Name     string
	Document string
}

func getLawsuitPoles(driver selenium.WebDriver) (Poles, error) {
	btnViewDocuments, err := driver.FindElement(selenium.ByXPATH, btnPetition)
	if err != nil {
		return Poles{}, errors.New("btnPetition not found")
	}
	err = btnViewDocuments.Click()

	time.Sleep(4 * time.Second)

	pageSource, err := driver.PageSource()
	if err != nil {
		return Poles{}, errors.New("could not get page source")
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return Poles{}, errors.New("could not convert string to node html")
	}

	var activePole []Part
	var passivePole []Part
	var otherPole []Part
	existCard := htmlquery.Find(htmlPgSrc, xpathTitleInit)

	if len(existCard) > 0 {
		cards := htmlquery.Find(htmlPgSrc, xpathTitleInit)

		for i := 0; i < len(cards); i++ {
			title := htmlquery.FindOne(cards[i], "/div/div[1]/h3")
			if htmlquery.InnerText(title) == titleActivePole {
				names := htmlquery.Find(cards[i], xpathNameInit)
				for _, name := range names {
					var n string
					var d string
					existName := htmlquery.Find(name, xpathName)
					if len(existName) > 0 && !strings.Contains(htmlquery.InnerText(htmlquery.FindOne(name, xpathName)), "Incluir no polo contrário") {
						n = htmlquery.InnerText(htmlquery.FindOne(name, xpathName))
						existDocument := htmlquery.Find(name, xpathDocument)
						if len(existDocument) > 0 {
							d = htmlquery.InnerText(htmlquery.FindOne(name, xpathDocument))
						}
						activePole = append(activePole, Part{
							Type:     "Ativo",
							Name:     n,
							Document: d,
						})
					}

				}
			}
			if htmlquery.InnerText(title) == titlePassivePole {
				names := htmlquery.Find(cards[i], xpathNameInit)
				for _, name := range names {
					var n string
					var d string
					existName := htmlquery.Find(name, xpathName)
					if len(existName) > 0 && !strings.Contains(htmlquery.InnerText(htmlquery.FindOne(name, xpathName)), "Incluir no polo contrário") {
						n = htmlquery.InnerText(htmlquery.FindOne(name, xpathName))
						existDocument := htmlquery.Find(name, xpathDocument)
						if len(existDocument) > 0 {
							d = htmlquery.InnerText(htmlquery.FindOne(name, xpathDocument))
						}
						passivePole = append(passivePole, Part{
							Type:     "Passivo",
							Name:     n,
							Document: d,
						})
					}

				}
			}
			if strings.Contains(htmlquery.InnerText(title), titleOthersPole) {
				names := htmlquery.Find(cards[i], xpathNameInit)
				for _, name := range names {
					var n string
					var d string
					existName := htmlquery.Find(name, xpathName)

					if len(existName) > 0 && !strings.Contains(htmlquery.InnerText(htmlquery.FindOne(name, xpathName)), "Incluir no polo contrário") {
						n = htmlquery.InnerText(htmlquery.FindOne(name, xpathName))
						existDocument := htmlquery.Find(name, xpathDocument)
						if len(existDocument) > 0 {
							d = htmlquery.InnerText(htmlquery.FindOne(name, xpathDocument))
						}
						otherPole = append(otherPole, Part{
							Type:     "Outros",
							Name:     n,
							Document: d,
						})
					}

				}
			}
		}
	} else {
		log.Println("no parts were found")
		return Poles{}, nil
	}

	return Poles{activePole, passivePole, otherPole}, nil
}
