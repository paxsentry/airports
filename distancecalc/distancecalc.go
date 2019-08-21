// distancecalc calculates the distance between two coordinates and returns an ordered list
package distancecalc

import (
    "httpgrabber"
    "inputhandler"
    "math"
    "sort"
    "strconv"
)

type OrderedData struct {
    Airports []Airport
}

type Airport struct {
    Name     string
    Distance string
}

var origLon float64
var origLat float64

const earthRadius = float64(6371)

// OrderData takes a JsonData object, iterates through it calculates the distance from the given starting point
// and returns an ordered list
func OrderData(airports httpgrabber.JsonData, starting inputhandler.LonLatRad) OrderedData {
    var response OrderedData
    origLon, _ = strconv.ParseFloat(starting.Lon, 64)
    origLat, _ = strconv.ParseFloat(starting.Lat, 64)

    for _, row := range airports.Rows {
        airport := Airport{
            Name:     row.Fields.Name,
            Distance: getdistance(row.Fields.Lon, row.Fields.Lat),
        }
        response.Airports = append(response.Airports, airport)
    }

    sort.Slice(response.Airports, func(i, j int) bool {
        return response.Airports[i].Distance < response.Airports[j].Distance
    })

    return response
}

// The haversine formula will calculate the spherical distance between lat and lon for two given points in km
func getdistance(lon float64, lat float64) string {
    deltaLon := (lon - origLon) * (math.Pi / 180)
    deltaLat := (lat - origLat) * (math.Pi / 180)

    a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
        math.Cos(origLat*(math.Pi/180))*math.Cos(lat*(math.Pi/180))*
            math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    return strconv.FormatFloat(earthRadius*c, 'f', 2, 64)
}
