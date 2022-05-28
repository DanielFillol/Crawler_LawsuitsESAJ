# Crawler_LawsuitsESAJ
Projeto desenhado para capturar dados processuais de um dado processo no primeiro e segundo grau do sistema ESAJ. São capturados todos os dados de ```capa```, ```partes```, ```advogados``` e ```movimentos```. Ao final do processamento são devolvidos um arquivo csv para cada um dos tipos de dados retornados.

NOTA: o crawler funciona para os seguintes TJ:```tjac,tjal,tjam,tjce, tjms e tjsp```.
 
## Dependências

Para o esse crawler somos dependentes do projeto [selenium](https://github.com/tebeka/selenium#readme), sendo necessário a pré-instalação do [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) e do [selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/). Nesse projeto também utilizamos o [htmlquery](https://github.com/antchfx/htmlquery).

## Run
Inicie o selenium server no terminal e depois basta 

```
java -jar selenium-server-standalone.jar

go run main.go

```

## Futuro
Permitir a pesquisa não apenas pelo número do CNJ, mas sim por qualquer dos dados das partes.
