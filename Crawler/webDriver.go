package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
)

func SeleniumWebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome", "Args": "--headless"})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, errors.New("could not create webdriver")
	}

	//driver.ResizeWindow("", 0, 0)

	return driver, nil
}
