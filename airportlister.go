package main

import (
    "distancecalc"
    "fmt"
    "httpgrabber"
    "inputhandler"
    "parambuilder"
)

func main() {
    inputResults := inputhandler.GetUserInput()

    urlparams := parambuilder.GetParamStringFromUserInput(inputResults)

    netresponse := httpgrabber.GetDataFromDatabase(urlparams)

    for _, answer := range distancecalc.OrderData(netresponse, inputResults).Airports {
        fmt.Println(answer.Name + " distance: " + answer.Distance + "km")
    }
}
