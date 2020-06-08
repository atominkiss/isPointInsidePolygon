package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// входные данные: точка на карте и полигон. Необходимо выяснить: 1.Находится точка внутри полигона или нет.
// сначала идет долгота потом широта [46.066426,51.572922]

func main(){
	var polygon, point string
	// var pointInside bool
	var coordPolygonArrStr, coordPointArrStr []string
	var longitudePolygon, latitudePolygon, coordPoint []float64 // долгота, широта
	polygon = "[{\"type\":\"Feature\",\"properties\":{},\"geometry\":{\"type\":\"Polygon\",\"coordinates\":[[[46.066426,51.572922],[46.069762,51.573089],[46.069923,51.57208],[46.068872,51.571318],[46.066362,51.569258],[46.064828,51.569218],[46.064581,51.571003],[46.064581,51.571057],[46.066641,51.571184],[46.066426,51.572922]]]}}]"
	point ="{\n  \"type\": \"FeatureCollection\",\n  \"features\": [\n    {\n      \"type\": \"Feature\",\n      \"properties\": {},\n      \"geometry\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          43.44976902008056,\n          56.23616843459138\n        ]\n      }\n    }\n  ]\n}"
	coordPolygonArrStr = extractCoordFromJsonToStrArr(polygon)
	coordPointArrStr = extractCoordFromJsonToStrArr(point)
	longitudePolygon, latitudePolygon = fromStrToFloatArr(coordPolygonArrStr)
	coordPoint = fromStrToFloatArrPoint(coordPointArrStr)
	
	pointInside := isPointInside(latitudePolygon, longitudePolygon ,  coordPoint)
	
	// fmt.Println(longitudePolygon)
	// fmt.Println(latitudePolygon)
	// fmt.Println(coordPoint)
	fmt.Println(pointInside)
	
}

func isPointInside(x []float64, y []float64, coordPoint []float64) bool {
	// многоугольник = 3 вершинам минимум. точку и отрезок игнорируем
	// x - широта (latitude), y - долгота(longitude)
	// quantity := len(x)
	// if  quantity <= 2 {
	// 	return false
	// }
	//
	// intersectionsNum := 0
	// prev := quantity - 1
	// prev_under := y[prev] < coordPoint[2]
	//
	// for i := 0; i < quantity; i++ {
	// 	cur_under := y[i] < coordPoint[2]
	//
	// 	a := polygon[prev] - point
	// 	b := polygon[i]  - point
	//
	// 	t := (a.x*(b.y - a.y) - a.y*(b.x - a.x))
	//
	// 	if cur_under && !prev_under {
	// 		if t > 0 {
	// 			intersectionsNum += 1
	// 		}
	// 		if !cur_under && prev_under {
	// 			if t < 0 {
	// 				intersectionsNum += 1
	// 			}
	// 		}
	// 	}
	//
	// 	prev = i
	// 	prev_under = cur_under
	// }
	//
	// return (intersectionsNum &1) != 0
	
	var i, j int
	var c bool
	npol := len(x)
	for i := 0; i < npol; i++ {
		// j = npol - 1  j = i++
		if (((yp [i] <= y) && (y < yp [j])) || ((yp [j] <= y) && (y < yp [i]))) &&
					(x < (xp [j] - xp [i]) * (y - yp [i]) / (yp [j] - yp [i]) + xp [i]) {
			c = !c
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
	latid:= make([]float64, len(arrayString)/2)
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
	sort.Float64s(longit)
	sort.Float64s(latid)
	
	return longit, latid
}

func extractCoordFromJsonToStrArr(polygon string) []string {
	for i :=0; i< len(polygon); i++{
		if unicode.IsDigit(rune(polygon[i])) || rune(polygon[i]) == '.' || rune(polygon[i]) == ','  {
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
