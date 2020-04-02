package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type dataregion struct {
	Name         string
	Casostotales int64
	Casosnuevos  int64
	Fallecidos   int64
}

func (app *application) Routes() {
	http.HandleFunc("/getDataMinsal", app.getDataMinsal)

}

func (app *application) getDataMinsal(w http.ResponseWriter, req *http.Request) {
	//Data:
	//Regiones:[
	//	{name:Atacama,casostotales:10,casosnuevos:0,fallecidos:0}
	//]

	fmt.Printf("Entro solicitud")

	var data map[string]map[string][]dataregion
	data = make(map[string]map[string][]dataregion)
	data["Data"] = make(map[string][]dataregion)

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
					var region string
					var casostotales int64
					var casosnuevos int64
					var fallecidos int64
					elementt.Children().Each(func(indexf int, elementf *goquery.Selection) {
						if indexf == 0 {
							// fmt.Printf("%+v\n", elementf)
							fmt.Printf("Region: %s\n", elementf.Text())
							region = elementf.Text()
						}

						if indexf == 1 {
							// fmt.Printf("Casos Nuevos:%s\n", elementf.Text())
							dato := strings.Replace(elementf.Text(), ".", "", -1)
							i2, err := strconv.ParseInt(dato, 10, 64)
							if err == nil {
								casosnuevos = i2
							}
						}

						if indexf == 2 {
							// fmt.Printf("Casos Totales:%s\n", elementf.Text())
							dato := strings.Replace(elementf.Text(), ".", "", -1)
							i2, err := strconv.ParseInt(dato, 10, 64)
							if err == nil {
								casostotales = i2
							}
						}

						if indexf == 4 {
							// fmt.Printf("Fallecidos: %s\n", elementf.Text())
							dato := strings.Replace(elementf.Text(), ".", "", -1)
							i2, err := strconv.ParseInt(dato, 10, 64)
							if err == nil {
								fallecidos = i2
							}
						}
					})
					data["Data"]["Regiones"] = append(data["Data"]["Regiones"], dataregion{Name: region, Casosnuevos: casosnuevos, Casostotales: casostotales, Fallecidos: fallecidos})
					//fmt.Printf("%+v\n", data)
				}
			})

		})
	})

	//fmt.Printf("%+v\n\n", data)
	als := data
	// fmt.Printf("%+v\n", als)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(als)
}
