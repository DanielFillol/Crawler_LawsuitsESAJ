package main

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

}

type SearchLawsuits struct {
	LawsuitNumber  string
	DocumentNumber string
}

var lawsuitNumbers = []SearchLawsuits{
	{LawsuitNumber: "1502024-40.2021.8.26.0567", DocumentNumber: "100.837.979-47"},
}
