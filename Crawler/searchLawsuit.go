package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strings"
)

const (
	xpathRadio  = "//*[@id=\"interna_NUMPROC\"]/div/fieldset/label[2]"
	xpathSearch = "//*[@id=\"nuProcessoAntigoFormatado\"]"
	xpathBttp   = "//*[@id=\"botaoConsultarProcessos\"]"
	xpathBtts   = "//*[@id=\"pbConsultar\"]"
	xpathReturn = "//*[@id=\"mensagemRetorno\"]/li"
)

func SearchLawsuit(driver selenium.WebDriver, searchLink string, lawsuit string, degree string, login string, password string) (*html.Node, error) {
	err := driver.Get(searchLink)
	if err != nil {
		return nil, errors.New("url unavailable")
	}

	rc, err := existReCaptcha(driver)
	if err != nil {
		return nil, err
	}

	if rc != false {
		err = Login(driver, login, password)
		if err != nil {
			return nil, err
		}
		err = driver.Get(searchLink)
		if err != nil {
			return nil, errors.New("url unavailable")
		}
	}

	radio, err := driver.FindElement(selenium.ByXPATH, xpathRadio)
	if err != nil {
		return nil, errors.New("could not find xpathRadio")
	}
	search, err := driver.FindElement(selenium.ByXPATH, xpathSearch)
	if err != nil {
		return nil, errors.New("could not find xpathSearch")
	}

	var btt selenium.WebElement
	if degree == "p" {
		btt, err = driver.FindElement(selenium.ByXPATH, xpathBttp)
		if err != nil {
			return nil, errors.New("could not find xpathBtt")
		}
	} else {
		btt, err = driver.FindElement(selenium.ByXPATH, xpathBtts)
		if err != nil {
			return nil, errors.New("could not find xpathBtt")
		}
	}

	err = radio.Click()
	if err != nil {
		return nil, errors.New("could not click on radio button")
	}

	err = search.SendKeys(lawsuit)
	if err != nil {
		return nil, errors.New("could not input lawsuit as search parameter")
	}

	err = btt.Click()
	if err != nil {
		return nil, errors.New("could not click on button")
	}

	pageSource, err := driver.PageSource()
	if err != nil {
		return nil, errors.New("could not get page source")
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return nil, errors.New("could not convert string to node html")
	}

	exist := existLawsuit(htmlPgSrc)
	if exist != true {
		return nil, errors.New("could not find lawsuit")
	}

	return htmlPgSrc, nil
}

func existLawsuit(htmlPgSrc *html.Node) bool {
	noReturn := htmlquery.Find(htmlPgSrc, xpathReturn)

	if len(noReturn) > 0 {
		return false
	}

	return true
}

func existReCaptcha(driver selenium.WebDriver) (bool, error) {
	pageSource, err := driver.PageSource()
	if err != nil {
		return false, err
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return false, err
	}

	rcpch := htmlquery.FindOne(htmlPgSrc, xpathRecaptcha)

	if htmlquery.InnerText(rcpch) == "Identificar-se " {
		return true, nil
	}
	return false, nil
}
