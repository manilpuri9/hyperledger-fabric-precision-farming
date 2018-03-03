package main

import "fmt"

//Crop struct is the defination of datastructure of crop
type Crop struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Quantity int    `json:"quantity"`
	FarmInfo struct {
		GeoLocation struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geo_location"`
		SoilType string `json:"soil_type"`
	} `json:"farm_info"`
	Weather struct {
		Temperature struct {
			Celcius int `json:"celcius"`
		} `json:"temperature"`
		Pressure struct {
			Pascal int `json:"pascal"`
		} `json:"pressure"`
		Humidity struct {
			CubicMeter float64 `json:"cubic_meter"`
		} `json:"humidity"`
		Radiation struct {
			Rem float64 `json:"rem"`
		} `json:"radiation"`
	} `json:"weather"`
	SoilCondition struct {
		Moisture struct {
			CubicMeter float64 `json:"cubic meter"`
		} `json:"moisture"`
		Ph       int `json:"ph"`
		Nitrogen struct {
			Percentage float64 `json:"percentage"`
		} `json:"nitrogen"`
		Phosphorus struct {
			Percentage float64 `json:"percentage"`
		} `json:"phosphorus"`
	} `json:"soil_condition"`
	Image              string `json:"image"`
	Cghc               int    `json:"cghc"`
	Irrigation         bool   `json:"irrigation"`
	FertilizerAddition bool   `json:"fertilizer_addition"`
	ApplyPesticide     bool   `json:"apply_pesticide"`
	Harvesting         bool   `json:"harvesting"`
}

func f1()

func main() {
	fmt.Println("hello world")
	var n = "rice"
	cropJson := Crop{
		Name:     n,
		Owner:    "manil puri",
		Quantity: 400,
		FarmInfo: {
			GeoLocation: {
				Latitude:  43.2,
				Longitude: 21.3,
			},
			SoilType: "clay",
		},
		Weather: {
			Temperature: {
				Celcius: 34,
			},
			Pressure: {
				Pascal: 4,
			},
			Humidity: {
				CubicMeter: 434,
			},
			Radiation: {
				Rem: 10.3,
			},
		},
		SoilCondition: {
			Moisture: {
				CubicMeter: 32,
			},
			Ph: 3,
			Nitrogen: {
				Percentage: 1.2,
			},
			Phosphorus: {
				Percentage: 3.4,
			},
		},
		Image:              "www.ncc.com/precisionfarming/crop/image/000",
		Cghc:               4,
		Irrigation:         true,
		FertilizerAddition: true,
		ApplyPesticide:     true,
		Harvesting:         false,
	}

	fmt.Println(cropJson)

}
