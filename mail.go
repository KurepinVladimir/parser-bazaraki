package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// main is the entry point of the application. It fetches the Hacker News homepage,
// checks the HTTP response status, parses the HTML using goquery, and extracts
// story titles and their links. The function prints the extracted titles and links
// to the standard output. It also includes debug print statements for tracing execution.
func main() {
	fmt.Println("test1")

	// Загружаем страницу
	//res, err := http.Get("https://news.ycombinator.com")

	adr := "https://www.bazaraki.com/car-motorbikes-boats-and-parts/cars-trucks-and-vans/doors---20/extras---130/extras---140/extras---160/extras---20/extras---50/gearbox---1/mileage_max---90000/year_min---66/?price_min=6000&price_max=12500"

	res, err := http.Get(adr)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Проверка статуса
	if res.StatusCode != 200 {
		log.Fatalf("Ошибка запроса: статус %d", res.StatusCode)
		fmt.Println("test2")
	}

	// Загружаем HTML в goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("test3")
		log.Fatal(err)

	}

	// Ищем заголовки
	fmt.Println("test4")
	doc.Find(".storylink").Each(func(i int, s *goquery.Selection) {
		fmt.Println("test5")
		title := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("%d: %s (%s)\n", i+1, title, link)
	})

	// Ищем заголовки
	doc.Find(".titleline > a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("%d: %s\n%s\n\n", i+1, title, link)
	})

}
