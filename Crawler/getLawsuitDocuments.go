package Crawler

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/model"
	"os"
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
	DocumentFound bool
}

func GetLawsuitDocuments(driver selenium.WebDriver, degree string, lawsuitNumber string, searchDocument string) (Document, error) {
	if degree == "s" {
		btnViewDocuments, err := driver.FindElement(selenium.ByXPATH, xPathViewDocumentsS)
		if err != nil {
			return Document{HasDocuments: false, DocumentFound: false}, errors.New("xPathDocuments not found")
		}
		err = btnViewDocuments.Click()
		if err != nil {
			return Document{}, errors.New("error clicking btnViewDocuments")
		}
	} else {
		btnViewDocuments, err := driver.FindElement(selenium.ByXPATH, xPathViewDocumentsP)
		if err != nil {
			return Document{HasDocuments: false, DocumentFound: false}, errors.New("xPathDocuments not found")
		}

		err = btnViewDocuments.Click()
		if err != nil {
			return Document{}, errors.New("error clicking btnViewDocuments")
		}
	}

	handles, err := driver.WindowHandles()
	if err != nil {
		return Document{}, errors.New("error getting handles")
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
	time.Sleep(500 * time.Millisecond)

	err = btnContinue.Click()
	if err != nil {
		return Document{}, errors.New("error clicking btnContinue")
	}

	finishElement, err := driver.FindElement(selenium.ByXPATH, xPathSaveDocument)
	if err != nil {
		return Document{}, errors.New("error finding xPathFinish")
	}

	start := time.Now()
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
	}

	err = finishElement.Click()
	if err != nil {
		return Document{}, errors.New("error clicking finishElement")
	}

	//wait for the download to finish
	time.Sleep(total)

	h, err := driver.CurrentWindowHandle()
	if err != nil {
		return Document{}, errors.New("error getting window handle")
	}

	err = driver.CloseWindow(h)
	if err != nil {
		return Document{}, errors.New("error switching windows")
	}

	err = driver.SwitchWindow(handles[0])
	if err != nil {
		return Document{}, errors.New("error switching windows")
	}

	found, err := searchDocumentOnFile(pdfPath+lawsuitNumber+".pdf", searchDocument)
	if err != nil {
		return Document{HasDocuments: true, DocumentFound: false}, err
	}

	err = os.RemoveAll(pdfPath + lawsuitNumber + ".pdf")
	if err != nil {
		return Document{}, err
	}

	return Document{
		HasDocuments:  true,
		DocumentFound: found,
	}, nil
}

func searchDocumentOnFile(pdfPath string, searchString string) (bool, error) {
	// Enable debug-level logging to see any parsing errors.
	common.SetLogger(common.ConsoleLogger{})

	// Load the PDF file.
	pdfReader, _, err := model.NewPdfReaderFromFile(pdfPath, nil)
	if err != nil {
		return false, fmt.Errorf("failed to open PDF file: %v", err)
	}

	// Get the number of pages in the PDF.
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return false, fmt.Errorf("failed to get number of pages: %v", err)
	}

	// Iterate through each page and search for the string.
	found := false
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return false, fmt.Errorf("failed to get page %d: %v", i, err)
		}

		content, err := page.GetContentStreams()
		if err != nil {
			return false, fmt.Errorf("failed to get content streams for page %d: %v", i, err)
		}

		// Convert the content streams to a string.
		contentStr := strings.Join(content, " ")

		// Search for the string in the content.
		if strings.Contains(contentStr, searchString) {
			found = true
			break
		}
	}

	return found, nil
}
