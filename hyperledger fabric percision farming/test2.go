package main

import (
	"encoding/json"
	"fmt"
)

type GeoLocationType struct {
	Latitude  float64 `json:latitude`
	Longitude float64 `json:"longitude"`
}

type FarmInfoType struct {
	GeoLocation GeoLocationType `json:geo_location`
	SoilType    string          `json:"soil_type"`
}

type TemperatureType struct {
	Celcius float64 `json:"celcius"`
}
type PressureType struct {
	Pascal float64 `json:"pascal"`
}
type HumidityType struct {
	CubicMeter float64 `json:"cubic_meter"`
}
type RadiationType struct {
	Rem float64 `json:"rem"`
}
type MoistureType struct {
	CubicMeter float64 `json:"cubic meter"`
}
type NitrogenType struct {
	Percentage float64 `json:"percentage"`
}
type PhosphorusType struct {
	Percentage float64 `json:"percentage"`
}

type WeatherType struct {
	Temperature TemperatureType `json:"temperature"`
	Pressure    PressureType    `json:"pressure"`
	Humidity    HumidityType    `json:"humidity"`
	Radiation   RadiationType   `json:"radiation"`
}

type SoilConditionType struct {
	Moisture   MoistureType   `json:"moisture"`
	Ph         int            `json:"ph"`
	Nitrogen   NitrogenType   `json:"nitrogen"`
	Phosphorus PhosphorusType `json:"phosphorus"`
}

type Crop struct {
	Name               string            `json:"name"`
	Owner              string            `json:"owner"`
	Quantity           int               `json:"quantity"`
	FarmInfo           FarmInfoType      `json:"farm_info"`
	Weather            WeatherType       `json:"weather"`
	SoilCondition      SoilConditionType `json:"soil_condition"`
	Image              string            `json:"image"`
	Cghc               int               `json:"cghc"`
	Irrigation         bool              `json:"irrigation"`
	FertilizerAddition bool              `json:"fertilizer_addition"`
	ApplyPesticide     bool              `json:"apply_pesticide"`
	Harvesting         bool              `json:"harvesting"`
}

func main() {
	fmt.Println("hello world")
	crop := Crop{
		Name:     "rice",
		Owner:    "manil puri",
		Quantity: 400,
		FarmInfo: FarmInfoType{
			GeoLocation: GeoLocationType{
				Latitude:  43.2,
				Longitude: 21.3,
			},
			SoilType: "clay",
		},
		Weather: WeatherType{
			Temperature: TemperatureType{
				Celcius: 34,
			},
			Pressure: PressureType{
				Pascal: 4,
			},
			Humidity: HumidityType{
				CubicMeter: 434,
			},
			Radiation: RadiationType{
				Rem: 10.3,
			},
		},
		//soil condition
		SoilCondition: SoilConditionType{
			Moisture: MoistureType{
				CubicMeter: 32,
			},
			Ph: 3,
			Nitrogen: NitrogenType{
				Percentage: 1.2,
			},
			Phosphorus: PhosphorusType{
				Percentage: 3.4,
			},
		},
	}

	cropJsonBytes, err := json.Marshal(crop)
	//cropBytes := []byte(cropJson)
	if err != nil {
		fmt.Println("error converting go crop type to json format.")
	}
	var cropJson Crop
	err1 := json.Unmarshal(cropJsonBytes, &cropJson)
	if err1 != nil {
		fmt.Println("error converting go crop type to json format.")
	}
	fmt.Println(cropJson)
	fmt.Println(crop)
}
