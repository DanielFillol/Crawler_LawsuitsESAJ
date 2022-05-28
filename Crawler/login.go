package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
)

const (
	SiteLogin  = "https://esaj.tjsp.jus.br/sajcas/login"
	XpathUser  = "//*[@id=\"usernameForm\"]"
	XpathPass  = "//*[@id=\"passwordForm\"]"
	XpathBtt   = "//*[@id=\"pbEntrar\"]"
	XpathError = "//*[@id=\"mensagemRetorno\"]/li"
)

func Login(driver selenium.WebDriver, login string, password string) error {
	err := driver.Get(SiteLogin)
	if err != nil {
		return errors.New("url unavailable")
	}

	userName, err := driver.FindElement(selenium.ByXPATH, XpathUser)
	if err != nil {
		return errors.New("xpathUser not found")
	}

	psw, err := driver.FindElement(selenium.ByXPATH, XpathPass)
	if err != nil {
		return errors.New("xpathPass not found")
	}

	btt, err := driver.FindElement(selenium.ByXPATH, XpathBtt)
	if err != nil {
		return errors.New("xpathBtt not found")
	}

	err = userName.SendKeys(login)
	if err != nil {
		return errors.New("could not send login parameter")
	}

	err = psw.SendKeys(password)
	if err != nil {
		return errors.New("could not send password parameter")
	}

	err = btt.Click()
	if err != nil {
		return errors.New("could not click on login button")
	}

	infoLogin, err := driver.FindElements(selenium.ByXPATH, XpathError)
	if err != nil {
		return errors.New("could not find xpathError")
	}

	if len(infoLogin) > 0 {
		for _, info := range infoLogin {
			innerText, err := info.Text()
			if err != nil {
				return errors.New("could not find inner text msg")
			}
			if innerText == "Usuário ou senha inválidos." {
				return errors.New("wrong user or password parameter")
			}
		}
	}

	return nil

}
