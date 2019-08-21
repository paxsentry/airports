// inputhandler user interaction to get input data and validate it
package inputhandler

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    . "strings"
)

// LonLatRad store user provided data to work with
type LonLatRad struct {
    Lon string
    Lat string
    Rng string
}

type validationResult struct {
    IsValid  bool
    ErrorMsg string
}

// GetUserInput to get input from the user and validate it, then return a struct to caller
func GetUserInput() LonLatRad {
    var response LonLatRad
    var answer validationResult
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Please enter Longitude:\n(no minutes or seconds, eg.: -86 or +154)")
        lon, _ := reader.ReadString('\n')
        answer = convertAndValidate(&lon, "lon")
        if answer.IsValid {
            lon = TrimPrefix(lon, "+")
            response.Lon = TrimSuffix(lon, "\n")
            break
        }
        fmt.Println(answer.ErrorMsg)
    }

    for {
        fmt.Println("Please enter Latitude:\n(no minutes or seconds, eg.: -49 or +10)")
        lat, _ := reader.ReadString('\n')
        answer = convertAndValidate(&lat, "lat")
        if answer.IsValid {
            lat = TrimPrefix(lat, "+")
            response.Lat = TrimSuffix(lat, "\n")
            break
        }
        fmt.Println(answer.ErrorMsg)
    }

    for {
        fmt.Println("Please enter range around the center point:")
        rng, _ := reader.ReadString('\n')
        answer = convertAndValidate(&rng, "rad")
        if answer.IsValid {
            response.Rng = TrimSuffix(rng, "\n")
            break
        }
        fmt.Println(answer.ErrorMsg)
    }

    return response
}

func convertAndValidate(s *string, property string) validationResult {
    var result validationResult
    result.IsValid = false

    sLen := len(TrimSpace(*s))
    if sLen == 0 || sLen > 4 {
        result.ErrorMsg = "Input length is invalid"
        return result
    }

    switch property {
    case "lon", "lat":
        if *s == "0\n" {
            break
        }

        if !HasPrefix(*s, "+") && !HasPrefix(*s, "-") {
            result.ErrorMsg = "Invalid prefix for coordinate, please use +/-"
            return result
        }
    case "rad":
        if HasPrefix(*s, "-") {
            result.ErrorMsg = "Radius must be positive number"
            return result
        }
    }

    sInt, err := strconv.Atoi(TrimSuffix(*s, "\n"))
    if err == nil {
        switch property {
        case "lon":
            if sInt < -180 || sInt > 180 {
                result.ErrorMsg = "Longitude must be within -180 and +180"
                return result
            }
        case "lat":
            if sInt < -90 || sInt > 90 {
                result.ErrorMsg = "Latitude must be within -90 and +90"
                return result
            }
        case "rad":
            if sInt <= 0 || sInt > 360 {
                result.ErrorMsg = "The possible Radius values are between 1 and 360"
                return result
            }
        default:
            result.ErrorMsg = "Invalid property"
        }
    }
    result.IsValid = true

    return result
}
