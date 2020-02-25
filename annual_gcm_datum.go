package client

import "encoding/xml"

// GCM represents yearly average precipitation
type GCM struct {
	XMLName       xml.Name `xml:"domain.web.AnnualGcmDatum"`
	Name          string   `xml:"gcm"`
	Precipitation string   `xml:"variable"`
	FromYear      string   `xml:"fromYear"`
	ToYear        string   `xml:"toYear"`
	AnnualData    float64  `xml:"annualData>double"`
}

// AnnualGcmDatum represents a collection of yearly precipitation for a country
//    <list>
//    	<domain.web.AnnualGcmDatum>
// 		<gcm>ukmo_hadgem1</gcm>
// 		<variable>pr</variable>
// 		<fromYear>1980</fromYear>
// 		<toYear>1999</toYear>
// 		<annualData>
// 			<double>1067.7044099748487</double>
// 		</annualData>
//    	</domain.web.AnnualGcmDatum>
//    </list>
type AnnualGcmDatum struct {
	XMLName         xml.Name `xml:"list"`
	AnnualGcmDatums []GCM    `xml:"domain.web.AnnualGcmDatum"`
}
