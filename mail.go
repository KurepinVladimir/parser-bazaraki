package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// Программа для парсинга сайта bazaraki.com
func main() {
	// Создаём контекст Chrome
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Добавим таймаут на всякий случай
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var html string

	// URL страницы
	url := "https://www.bazaraki.com/car-motorbikes-boats-and-parts/cars-trucks-and-vans/doors---20/extras---130/extras---20/extras---50/gearbox---1/mileage_max---90000/year_min---66/?price_min=6000&price_max=12500"

	// Запускаем браузер и получаем HTML после полной загрузки страницы
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second), // ждём загрузку JS
		chromedp.OuterHTML("body", &html),
	)
	if err != nil {
		log.Fatalf("Ошибка при загрузке страницы: %v", err)
	}

	// Загружаем в goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		log.Fatalf("Ошибка парсинга HTML: %v", err)
	}

	fmt.Println("Объявления:")
	// Для каждого блока объявления
	doc.Find(".advert__content").Each(func(i int, s *goquery.Selection) {
		// Название
		titleSel := s.Find("a.advert__content-title")
		title := strings.TrimSpace(titleSel.Text())
		link, _ := titleSel.Attr("href")

		// Цена
		price := strings.TrimSpace(s.Find("a.advert__content-price").Text())

		//<span><b>€</b>12.500  <span class="advert__content-price--discount"><b>€</b>13.200</span></span>
		priceNode := s.Find("a.advert__content-price > span").First()
		priceNode.Find(".advert__content-price--discount").Remove()
		price2 := strings.TrimSpace(priceNode.Text())

		// Выводим, если всё есть
		if title != "" && link != "" && price != "" {
			fmt.Printf("%d. %s\n", i+1, title)
			fmt.Println("   Ссылка: https://www.bazaraki.com" + link)
			fmt.Println("   Цена:  ", price)
			fmt.Println("   Цена (итоговая):  ", price2)
		}
	})

}
