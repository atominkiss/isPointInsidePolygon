package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// входные данные: точка на карте и полигон.
// Необходимо выяснить: 1.Находится точка внутри полигона или нет.
// сначала идет долгота потом широта [46.066426,51.572922]

func main() {
	var polygon, point string
	var coordPolygonArrStr, coordPointArrStr []string
	var longitudePolygon, latitudePolygon, coordPoint []float64 // долгота, широта
	polygon = "{\n      \"type\": \"Feature\",\n      \"properties\": {},\n      \"geometry\": {\n        \"type\": \"Polygon\",\n        \"coordinates\": [\n          [\n            [\n              43.448342084884644,\n              56.23708370966285\n            ],\n            [\n              43.44818115234375,\n              56.23471050828792\n            ],\n            [\n              43.4518826007843,\n              56.23475224928091\n            ],\n            [\n              43.45214009284973,\n              56.23667228578949\n            ],\n            [\n              43.448342084884644,\n              56.23708370966285\n            ]\n          ]\n        ]\n      }"
	point = "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [\n    {\n      \"type\": \"Feature\",\n      \"properties\": {},\n      \"geometry\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          43.45165729522705,\n          56.23481187919196\n        ]\n      }\n    }\n  ]\n}"
	coordPolygonArrStr = extractCoordFromJsonToStrArr(polygon)
	coordPointArrStr = extractCoordFromJsonToStrArr(point)
	longitudePolygon, latitudePolygon = fromStrToFloatArr(coordPolygonArrStr)
	coordPoint = fromStrToFloatArrPoint(coordPointArrStr)

	pointInside := isPointInside(latitudePolygon, longitudePolygon, coordPoint)

	fmt.Println(pointInside)

}

func isPointInside(x []float64, y []float64, coordPoint []float64) bool {
	// x - широта (latitude), y - долгота(longitude)

	var i, j int
	var c bool
	npol := len(x)
	j = npol - 1
	for i = 0; i < npol; i++ {
		if (((y[i] <= coordPoint[0]) && (coordPoint[0] < y[j])) || ((y[j] <= coordPoint[0]) && (coordPoint[0] < y[i]))) &&
			(coordPoint[1] < (x[j]-x[i])*(coordPoint[0]-y[i])/(y[j]-y[i])+x[i]) {
			c = !c
		}
		if i+2 >= npol {
			j = i + 2 - npol
		} else {
			j = i + 2
		}
	}

	return c

}

func fromStrToFloatArrPoint(str []string) []float64 {
	pointArr := make([]float64, len(str))

	for i := 0; i < len(str); i++ {
		pointArr[i], _ = strconv.ParseFloat(str[i], 64)
	}

	return pointArr
}

func fromStrToFloatArr(arrayString []string) ([]float64, []float64) {
	longit := make([]float64, len(arrayString)/2)
	latid := make([]float64, len(arrayString)/2)
	var lo, la int

	for i := 0; i < len(arrayString); i++ {
		coordinate, _ := strconv.ParseFloat(arrayString[i], 64)
		if i%2 == 0 {
			longit[lo] = coordinate
			lo++
		} else if i%2 != 0 {
			latid[la] = coordinate
			la++
		}
	}

	return longit, latid
}

func extractCoordFromJsonToStrArr(polygon string) []string {
	for i := 0; i < len(polygon); i++ {
		if unicode.IsDigit(rune(polygon[i])) || rune(polygon[i]) == '.' || rune(polygon[i]) == ',' {
			continue
		} else {
			polygon = strings.ReplaceAll(polygon, string(polygon[i]), "")
			i--
		}
	}
	polygon = strings.ReplaceAll(polygon, ",", " ")
	polygon = strings.TrimSpace(polygon)
	return strings.Split(polygon, " ")
}
