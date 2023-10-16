package main

//TODO: covers has another unmapped field "Outros números" this field can result in the search engine finding a different lawsuit number that was searched. Example: Search: 0000009-81.2015.8.26.0536 it returns: 0022797-06.2018.8.26.0562
//TODO: given this unmapped occurrence the getLawsuitDocuments is compromised as it searches for a file based on lawsuit searched, example: search for: 0000009-81.2015.8.26.0536.pdf but it does not exist, only: 0022797-06.2018.8.26.0562.pdf

import (
	"fmt"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/CSV"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
	"strconv"
	"time"
)

const (
	Login    = ""
	Password = ""
)

func main() {
	start1 := time.Now()

	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println(err)
	}

	defer driver.Close()

	err = Crawler.Login(driver, Login, Password)
	if err != nil {
		fmt.Println(err)
	}

	var suits []Crawler.EntireLawsuit
	for i, lawsuitNumber := range lawsuitNumbers {
		start2 := time.Now()

		lawsuit, _ := Crawler.Craw(driver, lawsuitNumber.LawsuitNumber, lawsuitNumber.DocumentNumber, Login, Password)
		suits = append(suits, lawsuit)

		t1 := time.Since(start1).String()
		t2 := time.Since(start2).String()
		m := int(time.Since(start1).Seconds()) / (i + 1)
		r := int(time.Since(start1).Seconds()) % (i + 1)
		md := strconv.Itoa(m) + "." + strconv.Itoa(r)
		fmt.Printf("processado %v | tempo: %v%v | total: %v%v | média: %vs \n", i+1, t2[0:4], t2[len(t2)-1:], t1[0:4], t1[len(t1)-1:], md)
	}

	err = CSV.WriteCSV(suits)
	if err != nil {
		fmt.Println(err)
	}

	driver.Close()

}

type SearchLawsuits struct {
	LawsuitNumber  string
	DocumentNumber []string
}

var lawsuitNumbers = []SearchLawsuits{
	{LawsuitNumber: "123456", DocumentNumber: []string{"1", "2"}},
}
