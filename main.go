
package main

import (
    "fmt"
    "strings"
    "log"
    "github.com/gocolly/colly"
    "github.com/charmbracelet/huh"
)

type Manga struct {
    Name   string
    Author string
    URL    string
}

func main() {
    var mangas []Manga
    c := colly.NewCollector()
    baseURL := "https://mangakakalot.com/search/story/"
    var mangaName string
    form := huh.NewForm(
	    huh.NewGroup(
		    huh.NewInput().
		    Title("Manga Name:").
		    Value(&mangaName),
	    ),
    )

    err := form.Run()
    if err != nil {
	    log.Fatal(err)
    }

    var selectedManga string

    

    mangaSlug := strings.Join(strings.Split(mangaName, " "), "_")

    mangaURL := baseURL + mangaSlug
    fmt.Println(mangaURL)

    c.OnHTML("div.story_item", func(e *colly.HTMLElement) {
        name := e.ChildText(".story_name")
        author := e.ChildText("span")
        url := e.ChildAttr("a", "href")

        authorCorrected := author
        if strings.HasSuffix(author, "Updated") {
            authorCorrected = author[:len(author)-len("Updated")]
        }

        manga := Manga{
            Name:   name,
            Author: authorCorrected,
            URL:    url,
        }

        mangas = append(mangas, manga)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.Visit(mangaURL)

    var manga_names []huh.Option[string]
    for _, manga := range mangas {
        manga_names = append(manga_names, huh.NewOption(manga.Name, manga.Name))
    }

    form2 := huh.NewForm(
	    huh.NewGroup(
		    huh.NewSelect[string]().
		    Title("Manga:").
		    Options(manga_names...).
		    Value(&selectedManga),
	    ),
    )

    err2 := form2.Run()
    if err2 != nil {
	    log.Fatal(err2)
    }
}
