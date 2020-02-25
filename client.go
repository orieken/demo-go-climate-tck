package client

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// ClimateAPI used to get percipitation data for a country
type ClimateAPI struct {
	CountryCode string
	ToYear      string
	FromYear    string
	Total       float64
	Average     float64
}

// CalculateTotalPrecipitation sums all yearly average precipitation
func (r *ClimateAPI) CalculateTotalPrecipitation(v AnnualGcmDatum) {
	r.Total = 0.0
	for _, year := range v.AnnualGcmDatums {
		r.Total += year.AnnualData
	}
}

// GetAveragePrecipitation calculates the average precipitation
func (r *ClimateAPI) GetAveragePrecipitation(v AnnualGcmDatum) {
	r.Average = r.Total / float64(len(v.AnnualGcmDatums))
}

// GetPrecipitationBetweenDates gets the yearly average precipitation for a country
func (r *ClimateAPI) GetPrecipitationBetweenDates(v AnnualGcmDatum) (AnnualGcmDatum, error) {
	url, urlError := buildURL(*r)
	if urlError != nil {
		return AnnualGcmDatum{}, urlError
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	contents, readingError := ioutil.ReadAll(res.Body)
	if readingError != nil {
		log.Fatal(readingError)
	}

	if string(contents) == "Invalid country code. Three letters are required" {
		return AnnualGcmDatum{}, errors.New(string(contents))
	}

	xmlErr := xml.Unmarshal([]byte(contents), &v)
	if xmlErr != nil {
		return AnnualGcmDatum{}, xmlErr
	}
	return v, nil
}

func buildURL(r ClimateAPI) (string, error) {
	u, err := url.Parse("http://climatedataapi.worldbank.org/climateweb/rest/v1/country/annualavg/pr/")
	if err != nil {
		return "", errors.New("Invalid url")
	}
	u.Path = path.Join(u.Path, r.FromYear, r.ToYear, r.CountryCode+".xml")
	return u.String(), nil
}
