# Crawler_LawsuitsESAJ
Project for crawlling lawsuit data avaliable in first and seccond degree of brazilian justice system.
Data Craw:
- capa ()
- partes (parties)
- advogados (lawyer's)
- movimentos (Steps of the lawsuit)

CSV file is generate with collected data.
 
## Dependencies
- [Selenium](https://github.com/tebeka/selenium#readme)
- [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/)
- [Selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
- [htmlquery](https://github.com/antchfx/htmlquery)

## Run
```brew install chrome driver ``` (not needed if you alredy have chrome driver)

```java -jar selenium-server-standalone.jar```

```go run main.go```

- To config a new search go in **crawler.go** file, function **DayCrawler** and alter the URL parameter on **driver.Get("")**

## Notes
- This crawler works for the following courts: ```tjac,tjal,tjam,tjce, tjms and tjsp```.
- Sometimes chromedriver need a clearnce in security options on MacOS.
- Don't forget to previus install Java.

## Future
Build a search function for every parameter allowed, exemple: name of people mentioned in the lawsuit.
