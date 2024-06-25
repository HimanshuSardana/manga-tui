
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

type Chapter struct {
    Name 	string
    URL		string
    UploadDate 	string
}

func main() {
    var mangas []Manga
    headers := map[string]string{
        "Host":                      "chapmanganato.to",
        "User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
        "Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-Language":           "en-US,en;q=0.5",
        "Accept-Encoding":           "gzip, deflate, br, zstd",
        "Connection":                "keep-alive",
        "Referer":                   "https://chapmanganato.to/manga-aa951409",
        "Cookie":                    "ci_session=Uj%2FMmsw3snmXDdm%2FAxXvAh30dOWXFaZawBkdrdGCA0eWDAxwI77%2FgH%2BU6TRe4RzUikBDsDMIYQTSDN%2FZ8O388NJzeHTdOsCLpCa6MoysPwk0g9fI1ntRO8qn0%2B3zZHWg6%2Be1SrOYgs0KxZU6wS9lo%2F3dej81aq1Vw%2Baz7EBeSsrYnVVqdcATFl7PhnVh65J3QEvJa8bMKCkeXsdyuGJNCOkGpkkCXlemCTNguS%2F71i2qygsuZa5G4XJqTaSBDWN4%2FzhdrcgGhggfc5wtQIsYT1qYEsXcWXcq2J32x%2BnHMiTp%2FSc9rxbm4jvPnaf8tZOG%2FVBcvbOk3ZvrjxVQiG6bY0V6uYmqM4Fm6eWaj5%2Fhik9dwWz5BJISf%2B4lJJJadzg8CcVIJht5ABZAKrytGkFpKtFhHiXFnKqW4HfV%2Fqt9HoY%3Da3468db3c15b1d69fbb668ff5a3f14ce1c5bffc6; panel-fb-comment=fb-comment-title-show",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest":            "document",
        "Sec-Fetch-Mode":            "navigate",
        "Sec-Fetch-Site":            "same-origin",
        "Sec-Fetch-User":            "?1",
        "Priority":                  "u=1",
        "TE":                        "trailers",
    }


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

    var selectedMangaURL string
    for _, manga := range mangas {
	if manga.Name == selectedManga {
		selectedMangaURL = manga.URL
	}
    }

    fmt.Println(selectedMangaURL)

    var chapters []Chapter
    c.OnHTML("li.a-h", func(e *colly.HTMLElement) {
        chapterName := e.ChildText("a.chapter-name")
        chapterURL := e.ChildAttr("a.chapter-name", "href")
        date := e.ChildText("span.chapter-time")

        chapter := Chapter{
            Name:       chapterName,
            URL:        chapterURL,
            UploadDate: date,
        }

        chapters = append(chapters, chapter)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.Visit(selectedMangaURL)

    var chapter_names []huh.Option[string]
    for _, chapter := range chapters {
        chapter_names = append(chapter_names, huh.NewOption(chapter.Name, chapter.Name))
    }

    var selectedChapter string
    form3 := huh.NewForm(
	    huh.NewGroup(
		    huh.NewSelect[string]().
		    Title("Chapter:").
		    Options(chapter_names...).
		    Value(&selectedChapter),
	    ),
    )

    err3 := form3.Run()
    if err3 != nil {
	    log.Fatal(err3)
    }

    var selectedChapterURL string

    for _, chapter := range chapters {
	if chapter.Name == selectedChapter {
		selectedChapterURL = chapter.URL
    	    }
    }

    var imageURLs []string
    c.OnHTML("*", func(e *colly.HTMLElement) {
	    //ImageURL := e.ChildAttr("img","src")
	    //imageURLs = append(imageURLs, ImageURL)
	    fmt.Println(e.DOM.Html())
	    fmt.Println("hello")
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Response Status Code:", r.StatusCode)

        // Check if the request was successful (status code 200)
        if r.StatusCode == 200 {
            // Process the response here
            fmt.Println("Successful response")
	    fmt.Println("HTML:", r.Headers)
        } else {
            fmt.Println("Request failed with status code:", r.StatusCode)
        }
    })

    c.OnRequest(func(r *colly.Request) {
	for key, value := range headers {
        	r.Headers.Set(key, value)
        }
        fmt.Println("Visiting", r.URL)
    })


    selectedChapterURL = "https://chapmanganelo.com/manga-aa88620/chapter-1118"
    c.Visit(selectedChapterURL)
    fmt.Println(imageURLs)
}
