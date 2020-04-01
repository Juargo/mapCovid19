package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type dataregion struct {
	name         string
	casostotales int64
	casosnuevos  int64
	fallecidos   int64
}

func (app *application) Routes() {
	http.HandleFunc("/getDataMinsal", app.getDataMinsal)

}

func (app *application) getDataMinsal(w http.ResponseWriter, req *http.Request) {
	//Data:
	//Regiones:[
	//	{name:Atacama,casostotales:10,casosnuevos:0,fallecidos:0}
	//]

	var data map[string]map[string][]dataregion
	data = make(map[string]map[string][]dataregion)
	data["Data"] = make(map[string][]dataregion)

	var region string
	var casostotales int64
	var casosnuevos int64
	var fallecidos int64

	// Make HTTP request
	response, err := http.Get("https://www.minsal.cl/nuevo-coronavirus-2019-ncov/casos-confirmados-en-chile-covid-19/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".contenido").ChildrenFiltered("table").First().Each(func(index int, element *goquery.Selection) {
		element.Children().Each(func(indexd int, elementd *goquery.Selection) {
			elementd.Children().Each(func(indext int, elementt *goquery.Selection) {
				if indext > 2 {
					elementt.Children().Each(func(indexf int, elementf *goquery.Selection) {
						if indexf == 0 {
							// fmt.Printf("%+v\n", elementf)
							fmt.Printf("Region: %s\n", elementf.Text())
							region = elementf.Text()
						}
						if indexf == 1 {
							fmt.Printf("Casos Nuevos:%s\n", elementf.Text())
							i2, err := strconv.ParseInt(elementf.Text(), 10, 64)
							if err == nil {
								casosnuevos = i2
							}
						}
						if indexf == 2 {
							fmt.Printf("Casos Totales:%s\n", elementf.Text())
							i2, err := strconv.ParseInt(elementf.Text(), 10, 64)
							if err == nil {
								casostotales = i2
							}
						}
						if indexf == 4 {
							fmt.Printf("Fallecidos: %s\n", elementf.Text())
							i2, err := strconv.ParseInt(elementf.Text(), 10, 64)
							if err == nil {
								fallecidos = i2
							}
							data["Data"]["Regiones"] = append(data["Data"]["Regiones"], dataregion{name: region, casosnuevos: casosnuevos, casostotales: casostotales, fallecidos: fallecidos})
							fmt.Printf("%+v\n", data)
						}
					})
				}
			})

		})
	})

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	als := js
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(als)
}
