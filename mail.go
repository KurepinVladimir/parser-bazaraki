package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiToken := "apify_api_I69tt6CZ9WJnfKDq5s4x2SUeIHD3ok4GhKKb" // вставь сюда свой токен
	url := "https://api.apify.com/v2/acts/sashkavasa~bazaraki-scrapper/runs?token=" + apiToken

	jsonData := []byte(`{
		"startUrls": [
			{
				"url": "https://www.bazaraki.com/car-motorbikes-boats-and-parts/cars-trucks-and-vans/doors---20/extras---130/extras---20/extras---50/gearbox---1/mileage_max---90000/year_min---66/?price_min=6000&price_max=12500"
			}
		]
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
}
