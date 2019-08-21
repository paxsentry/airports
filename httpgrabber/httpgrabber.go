// httpgrabber a package to retrieve data from the internet
package httpgrabber

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

//JsonData the parsed response from the database
type JsonData struct {
    Rows []Row `json:rows`
}

type Row struct {
    Fields Field `json:"fields"`
}

type Field struct {
    Lat  float64 `json:"lat"`
    Lon  float64 `json:"lon"`
    Name string  `json:"name"`
}

const basePath = "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo"

// GetDataFromDatabase grabs the data from an internet database
func GetDataFromDatabase(queryparam string) JsonData {
    var response JsonData

    fullpath, _ := url.Parse(basePath)
    params := url.Values{}
    params.Add("q", queryparam)
    fullpath.RawQuery = params.Encode()

    resp, err := http.Get(fullpath.String())
    if err != nil {
        fmt.Println(err)
    }
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    json.Unmarshal(bytes, &response)

    return response
}
