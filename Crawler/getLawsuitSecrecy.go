package Crawler

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const xpathSecrecy = "//*[@id=\"popupSenha\"]"

func GetSecrecy(htmlPgSrc *html.Node) bool {
	secrecy := htmlquery.SelectAttr(htmlquery.FindOne(htmlPgSrc, xpathSecrecy), "style")

	if secrecy == "display: block;" {
		return true
	}
	return false
}
