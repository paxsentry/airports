package parambuilder

import (
    "inputhandler"
    "strconv"
)

func GetParamStringFromUserInput(data inputhandler.LonLatRad) string {
    paramResult := "lon:[" + data.Lon + " TO " + calculateRange(data.Lon, data.Rad, "lon")
    paramResult += "] AND lat:[" + data.Lat + " TO " + calculateRange(data.Lat, data.Rad, "lat") + "]"

    return paramResult
}

func calculateRange(dot string, rng string, property string) string {
    var dotInt, _ = strconv.Atoi(dot)
    var rngInt, _ = strconv.Atoi(rng)
    sum := dotInt + rngInt

    switch property {
    case "lon": // -180 => 180
        if sum > 180 {
            return "180"
        }
    case "lat": // -90 => 90
        if sum > 90 {
            return "90"
        }
    }
    return strconv.Itoa(sum)
}
