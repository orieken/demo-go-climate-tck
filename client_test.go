package client

import (
	"log"
	"math/big"
	"testing"
)

func TestAverageRainfallForNigeriaFrom1980to1999(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "nga",
		FromYear:    "1980",
		ToYear:      "1999",
	}
	a := AnnualGcmDatum{}

	v, responseError := climateAPI.GetPrecipitationBetweenDates(a)
	if responseError != nil {
		log.Fatal(responseError)
	}

	climateAPI.CalculateTotalPrecipitation(v)
	climateAPI.GetAveragePrecipitation(v)

	expectedAverage := 1244.729934
	if big.NewFloat(climateAPI.Average).Text('f', 6) != big.NewFloat(expectedAverage).Text('f', 6) {
		t.Errorf("Expected %f to equal %f", climateAPI.Average, expectedAverage)
	}
}

func TestAverageRainfallForCamaroonFrom1980to1999(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "cmr",
		FromYear:    "1980",
		ToYear:      "1999",
	}
	a := AnnualGcmDatum{}

	v, responseError := climateAPI.GetPrecipitationBetweenDates(a)
	if responseError != nil {
		log.Fatal(responseError)
	}

	climateAPI.CalculateTotalPrecipitation(v)
	climateAPI.GetAveragePrecipitation(v)

	expectedAverage := 1592.049237
	if big.NewFloat(climateAPI.Average).Text('f', 6) != big.NewFloat(expectedAverage).Text('f', 6) {
		t.Errorf("Expected %f to equal %f", climateAPI.Average, expectedAverage)
	}
}

func TestAverageRainfallForHondurasFrom1980to1999(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "hnd",
		FromYear:    "1980",
		ToYear:      "1999",
	}
	a := AnnualGcmDatum{}

	v, responseError := climateAPI.GetPrecipitationBetweenDates(a)
	if responseError != nil {
		log.Fatal(responseError)
	}

	climateAPI.CalculateTotalPrecipitation(v)
	climateAPI.GetAveragePrecipitation(v)

	expectedAverage := 1136.964703
	if big.NewFloat(climateAPI.Average).Text('f', 6) != big.NewFloat(expectedAverage).Text('f', 6) {
		t.Errorf("Expected %f to equal %f", climateAPI.Average, expectedAverage)
	}
}

func TestTrowsErrorIfCountryISOisLessThanThreeCharacters(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "ng",
		FromYear:    "1980",
		ToYear:      "1999",
	}
	a := AnnualGcmDatum{}

	_, responseError := climateAPI.GetPrecipitationBetweenDates(a)
	expected := "Invalid country code. Three letters are required"
	if responseError.Error() != expected {
		t.Errorf("expected GetPrecipitationBetweenDates to throw %v but got %v", responseError.Error(), expected)
	}
}

func TestTrowsErrorIfCountryISOisMoreThanThreeCharacters(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "ngaa",
		FromYear:    "1980",
		ToYear:      "1999",
	}
	a := AnnualGcmDatum{}

	_, responseError := climateAPI.GetPrecipitationBetweenDates(a)
	expected := "Invalid country code. Three letters are required"
	if responseError.Error() != expected {
		t.Errorf("expected GetPrecipitationBetweenDates to throw %v but got %v", responseError.Error(), expected)
	}
}

func TestReturnsEmptyAnnualGcmDatumIfDateRangeIsInvalid(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "nga",
		FromYear:    "0000",
		ToYear:      "0000",
	}
	a := AnnualGcmDatum{}

	v, _ := climateAPI.GetPrecipitationBetweenDates(a)
	expected := 0
	if len(v.AnnualGcmDatums) != expected {
		t.Errorf("expected length %v to equal %v", len(v.AnnualGcmDatums), expected)
	}
}

func TestBuildUrlReturnsProperlyFormattedURL(t *testing.T) {
	climateAPI := ClimateAPI{
		CountryCode: "xxx",
		FromYear:    "xxxx",
		ToYear:      "oooo",
	}
	u, _ := buildURL(climateAPI)
	expectedURL := "http://climatedataapi.worldbank.org/climateweb/rest/v1/country/annualavg/pr/xxxx/oooo/xxx.xml"
	if u != expectedURL {
		t.Errorf("expected %v to equal %v", u, expectedURL)
	}
}
