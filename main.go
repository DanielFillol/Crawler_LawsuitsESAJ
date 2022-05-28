package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/CSV"
	"github.com/Darklabel91/Crawler_LawsuitsESAJ/Crawler"
	"strconv"
	"time"
)

const (
	Login    = "718.191.691-20"
	Password = "fwrs1825PC"
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

		lawsuit, _ := Crawler.Craw(driver, lawsuitNumber, Login, Password)
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

var lawsuitNumbers = []string{
	"1001290-10.2018.8.26.0614",
	"1001295-32.2018.8.26.0614",
	"1022500-82.2020.8.26.0506",
	"1505661-48.2020.8.26.0562",
	"1016453-07.2020.8.26.0114",
	"1502935-46.2019.8.26.0624",
	"0002087-43.2021.8.26.0114",
	"1024842-50.2020.8.26.0576",
	"1040316-88.2014.8.26.0053",
	"0001975-53.2020.8.26.0100",
	"1002332-82.2016.8.26.0576",
	"1002342-20.2020.8.26.0663",
	"1000412-53.2020.8.26.0311",
	"1512443-66.2019.8.26.0090",
	"1514702-34.2019.8.26.0090",
	"1016484-88.2017.8.26.0451",
	"1051679-18.2020.8.26.0100",
	"1013162-57.2020.8.26.0224",
	"1012755-85.2018.8.26.0009",
	"1009001-78.2020.8.26.0361",
}
