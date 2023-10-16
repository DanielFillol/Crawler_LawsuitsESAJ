package Crawler

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/model"
	"strings"
	"time"
)

const (
	xPathViewDocumentsS = "//*[@id=\"pbVisualizarAutos\"]"
	xPathViewDocumentsP = "//*[@id=\"linkPasta\"]"
	xPathAllDocuments   = "//*[@id=\"selecionarButton\"]"
	xPathDownloadPdf    = "//*[@id=\"salvarButton\"]"
	xPathContinue       = "/html/body/div[6]/div/form/div[2]/p[1]/input"
	xPathSaveDocument   = "//*[@id=\"btnDownloadDocumento\"]"
)

var pdfPath = "/Users/danielfillol/Downloads/"

type Document struct {
	HasDocuments  bool
	Found         bool
	DocumentFound string
	Page          int
}

type Pdf struct {
	found         bool
	documentFound string
	page          int
}

func GetLawsuitDocuments(driver selenium.WebDriver, degree string, lawsuitNumber string, searchDocument []string) (Document, error) {
	if degree == "s" {
		btnViewDocuments, err := driver.FindElement(selenium.ByXPATH, xPathViewDocumentsS)
		if err != nil {
			return Document{}, errors.New("xPathDocuments not found")
		}
		err = btnViewDocuments.Click()
		if err != nil {
			return Document{}, errors.New("error clicking btnViewDocuments")
		}
	} else {
		btnViewDocuments, err := driver.FindElement(selenium.ByXPATH, xPathViewDocumentsP)
		if err != nil {
			return Document{}, errors.New("xPathDocuments not found")
		}

		err = btnViewDocuments.Click()
		if err != nil {
			return Document{}, errors.New("error clicking btnViewDocuments")
		}
	}

	time.Sleep(2 * time.Second)

	handles, err := driver.WindowHandles()
	if err != nil {
		return Document{}, errors.New("error getting handles")
	}

	err = driver.CloseWindow(handles[0])
	if err != nil {
		return Document{}, errors.New("error closing first")
	}

	err = driver.SwitchWindow(handles[len(handles)-1])
	if err != nil {
		return Document{}, errors.New("error switching windows")
	}

	btnAllDocuments, err := driver.FindElement(selenium.ByXPATH, xPathAllDocuments)
	if err != nil {
		return Document{}, errors.New("xPathAllDocuments not found")
	}

	err = btnAllDocuments.Click()
	if err != nil {
		return Document{}, errors.New("error clicking btnAllDocuments")
	}

	btnDownloadPdf, err := driver.FindElement(selenium.ByXPATH, xPathDownloadPdf)
	if err != nil {
		return Document{}, errors.New("xPathDownload not found")
	}

	err = btnDownloadPdf.Click()
	if err != nil {
		return Document{}, errors.New("error clicking btnAllDocuments")
	}

	btnContinue, err := driver.FindElement(selenium.ByXPATH, xPathContinue)
	if err != nil {
		return Document{}, errors.New("xPathContinue not found")
	}
	//wait for JS to load
	time.Sleep(2 * time.Second)

	err = btnContinue.Click()
	if err != nil {
		return Document{}, errors.New("error clicking btnContinue")
	}

	finishElement, err := driver.FindElement(selenium.ByXPATH, xPathSaveDocument)
	if err != nil {
		return Document{}, errors.New("error finding xPathFinish")
	}

	start := time.Now()
	moreThan60s := false
	var total time.Duration
	for {
		element, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"popupGerarDocumentoFinalizadoComSucesso\"]/table/tbody/tr/td[2]")
		if err != nil {
			return Document{}, errors.New("error finding popupGerarDocumentoFinalizadoComSucesso")
		}

		t, err := element.Text()
		if err != nil {
			return Document{}, errors.New("error getting element  text")
		}

		if t == "O documento foi gerado. Você já pode salvá-lo." {
			total = time.Since(start)
			break
		}

		if time.Since(start) >= 60*time.Second {
			moreThan60s = true
			break
		}
	}

	if moreThan60s {
		return Document{HasDocuments: true, Found: false, DocumentFound: "O download foi superior a 1min", Page: 0}, errors.New("more than 20 seconds to download")
	}

	err = finishElement.Click()
	if err != nil {
		return Document{}, errors.New("error clicking finishElement")
	}

	//wait for the download to finish
	time.Sleep(total + 5)

	pdf, err := searchDocumentOnFile(pdfPath+lawsuitNumber+".pdf", searchDocument)
	if err != nil {
		return Document{HasDocuments: true, Found: false, DocumentFound: "", Page: 0}, err
	}

	//err = os.RemoveAll(pdfPath + lawsuitNumber + ".pdf")
	//if err != nil {
	//	return Document{}, err
	//}

	return Document{
		HasDocuments:  true,
		Found:         pdf.found,
		DocumentFound: pdf.documentFound,
		Page:          pdf.page,
	}, nil

}

func searchDocumentOnFile(pdfPath string, searchString []string) (Pdf, error) {
	// Enable debug-level logging to see any parsing errors.
	common.SetLogger(common.ConsoleLogger{})

	// Load the PDF file.
	pdfReader, _, err := model.NewPdfReaderFromFile(pdfPath, nil)
	if err != nil {
		return Pdf{found: false, documentFound: "", page: 0}, fmt.Errorf("failed to open PDF file: %v", err)
	}

	// Get the number of pages in the PDF.
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return Pdf{found: false, documentFound: "", page: 0}, fmt.Errorf("failed to get number of pages: %v", err)
	}

	// Iterate through each page and search for the string.
	found := false
	pg := 0
	dcf := ""
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return Pdf{found: false, page: pg}, fmt.Errorf("failed to get page %d: %v", i, err)
		}

		content, err := page.GetContentStreams()
		if err != nil {
			return Pdf{found: false, page: pg}, fmt.Errorf("failed to get content streams for page %d: %v", i, err)
		}

		// Convert the content streams to a string.
		contentStr := strings.Join(content, " ")

		for _, s := range searchString {
			// Search for the string in the content.
			if strings.Contains(contentStr, s) {
				found = true
				pg = i
				dcf += "{" + s + "}"
			}

			// Search for the string in the content.
			if strings.Contains(contentStr, strings.Replace(strings.Replace(s, ".", "", -1), "-", "", -1)) {
				found = true
				pg = i
				dcf += "{" + strings.Replace(strings.Replace(s, ".", "", -1), "-", "", -1) + "}"
			}

			// Search for the string in the content.
			if strings.Contains(contentStr, s) {
				found = true
				pg = i
				dcf += "{" + s + "}"
			}

			if found {
				break
			}
		}
	}

	return Pdf{
		found:         found,
		documentFound: dcf,
		page:          pg,
	}, nil
}
