/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

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
	Name           string            `json:"name"`
	Owner          string            `json:"owner"`
	Quantity       int               `json:"quantity"`
	FarmInfo       FarmInfoType      `json:"farm_info"`
	Weather        WeatherType       `json:"weather"`
	SoilCondition  SoilConditionType `json:"soil_condition"`
	Image          string            `json:"image"`
	Cghc           int               `json:"cghc"`
	Irrigation     bool              `json:"irrigation"`
	AddFertilizer  bool              `json:"fertilizer_addition"`
	ApplyPesticide bool              `json:"apply_pesticide"`
	Harvesting     bool              `json:"harvesting"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initCrop" { //create a new Crop
		return t.initCrop(stub, args)
	} else if function == "queryCrop" { //find Crop based on an ad hoc rich query
		return t.queryCrop(stub, args)
	}else if function == "updateCrop" { //update Crop based on an ad hoc rich query
		return t.updateCrop(stub, args)
	} else if function == "historyOfCrop" { //find Crop based on an ad hoc rich query
		return t.getHistoryForCrop(stub, args)
	} else if function == "readCrop" { //find Crop based on an ad hoc rich query
		return t.readCrop(stub, args)
	} else if function == "deleteCrop" { //find Crop based on an ad hoc rich query
		return t.delete(stub, args)
	} else if function == "irrigationCrop" { //find Crop based on an ad hoc rich query
		return t.irrigation(stub, args)
	} else if function == "addFertilizerCrop" { //find Crop based on an ad hoc rich query
		return t.addFertilizer(stub, args)
	} else if function == "applyPesticideCrop" { //find Crop based on an ad hoc rich query
		return t.applyPesticide(stub, args)
	} else if function == "harvestCrop" { //find Crop based on an ad hoc rich query
		return t.harvest(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initCrop - create a new Crop, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 20 {
		return shim.Error("Incorrect number of arguments. Expecting 20")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init crop")

	cropnamev := args[0]
	ownerv := args[1]
	quantityv, err := strconv.Atoi(args[2])
	lativ, err := strconv.ParseFloat(args[3], 64)
	longiv, err := strconv.ParseFloat(args[4], 64)
	celv, err := strconv.ParseFloat(args[6], 64)
	pasv, err := strconv.ParseFloat(args[7], 64)
	humv, err := strconv.ParseFloat(args[8], 64)
	radv, err := strconv.ParseFloat(args[9], 64)
	moistv, err := strconv.ParseFloat(args[10], 64)
	phv, err := strconv.Atoi(args[11])
	nitrov, err := strconv.ParseFloat(args[12], 64)
	phosv, err := strconv.ParseFloat(args[13], 64)
	imagev := args[14]
	cgphv, err := strconv.Atoi(args[15])
	irrv, err := strconv.ParseBool(args[16])
	ferv, err := strconv.ParseBool(args[17])
	appv, err := strconv.ParseBool(args[18])
	harv, err := strconv.ParseBool(args[19])

	if err != nil {
		return shim.Error(err.Error())
	}
	crop := Crop{
		Name:     cropnamev,
		Owner:    ownerv,
		Quantity: quantityv,
		FarmInfo: FarmInfoType{
			GeoLocation: GeoLocationType{
				Latitude:  lativ,
				Longitude: longiv,
			},
			SoilType: strings.ToLower(args[5]),
		},
		Weather: WeatherType{
			Temperature: TemperatureType{
				Celcius: celv,
			},
			Pressure: PressureType{
				Pascal: pasv,
			},
			Humidity: HumidityType{
				CubicMeter: humv,
			},
			Radiation: RadiationType{
				Rem: radv,
			},
		},
		//soil condition
		SoilCondition: SoilConditionType{
			Moisture: MoistureType{
				CubicMeter: moistv,
			},
			Ph: phv,
			Nitrogen: NitrogenType{
				Percentage: nitrov,
			},
			Phosphorus: PhosphorusType{
				Percentage: phosv,
			},
		},
		Image:          imagev,
		Cghc:           cgphv,
		Irrigation:     irrv,
		AddFertilizer:  ferv,
		ApplyPesticide: appv,
		Harvesting:     harv,
	}

	// ==== Check if crop already exists ====
	gotCropAsBytes, err := stub.GetState(cropnamev)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	} else if gotCropAsBytes != nil {
		fmt.Println("This marble already exists: " + cropnamev)
		return shim.Error("This marble already exists: " + cropnamev)
	}

	// ==== Create crop object and crop to JSON ====
	cropJsonBytes, err := json.Marshal(crop)
	if err != nil {
		return shim.Error(err.Error())
	}

	

	// === Save crop to state ===
	err = stub.PutState(cropnamev, cropJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	///////////////////////////////////////////////////////////////////////
	//  ==== Index the crop to enable owner-based range queries, e.g. return all crops with same owner ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~owner~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~owner~*
	indexName := "owner~name"
	ownerNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{crop.Owner, crop.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(ownerNameIndexKey, value)

	// ==== Crop saved and indexed. Return success ====
	fmt.Println("- end init crop successful")
	return shim.Success(nil)
}

// ============================================================
// UpdateCrop - updates info of Crop, store into chaincode state
// ============================================================
func (t *SimpleChaincode) updateCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var cropJSON Crop

	if len(args) != 16 {
		return shim.Error("Incorrect number of arguments. Expecting 20")
	}

	// ==== Input sanitation ====
	fmt.Println("- start update crop")

	cropnamev := args[0]

	// ==== Check if crop already exists ====
	gotCropAsBytes, err := stub.GetState(cropnamev)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	}
	err = json.Unmarshal([]byte(gotCropAsBytes), &cropJSON)
	if err != nil {
		return shim.Error("Failed to unmarshal crop to json format " + err.Error())
	}

	ownerv := cropJSON.Owner
	quantityv := cropJSON.Quantity
	lativ := cropJSON.FarmInfo.GeoLocation.Latitude
	longiv := cropJSON.FarmInfo.GeoLocation.Longitude
	soilv := cropJSON.FarmInfo.SoilType

	celv, err := strconv.ParseFloat(args[6], 64)
	pasv, err := strconv.ParseFloat(args[7], 64)
	humv, err := strconv.ParseFloat(args[8], 64)
	radv, err := strconv.ParseFloat(args[9], 64)
	moistv, err := strconv.ParseFloat(args[10], 64)
	phv, err := strconv.Atoi(args[11])
	nitrov, err := strconv.ParseFloat(args[12], 64)
	phosv, err := strconv.ParseFloat(args[13], 64)
	imagev := args[14]
	cgphv, err := strconv.Atoi(args[15])
	irrv := cropJSON.Irrigation
	ferv := cropJSON.AddFertilizer
	appv := cropJSON.ApplyPesticide
	harv := cropJSON.Harvesting

	if err != nil {
		return shim.Error(err.Error())
	}
	crop := Crop{
		Name:     cropnamev,
		Owner:    ownerv,
		Quantity: quantityv,
		FarmInfo: FarmInfoType{
			GeoLocation: GeoLocationType{
				Latitude:  lativ,
				Longitude: longiv,
			},
			SoilType: strings.ToLower(soilv),
		},
		Weather: WeatherType{
			Temperature: TemperatureType{
				Celcius: celv,
			},
			Pressure: PressureType{
				Pascal: pasv,
			},
			Humidity: HumidityType{
				CubicMeter: humv,
			},
			Radiation: RadiationType{
				Rem: radv,
			},
		},
		//soil condition
		SoilCondition: SoilConditionType{
			Moisture: MoistureType{
				CubicMeter: moistv,
			},
			Ph: phv,
			Nitrogen: NitrogenType{
				Percentage: nitrov,
			},
			Phosphorus: PhosphorusType{
				Percentage: phosv,
			},
		},
		Image:          imagev,
		Cghc:           cgphv,
		Irrigation:     irrv,
		AddFertilizer:  ferv,
		ApplyPesticide: appv,
		Harvesting:     harv,
	}

	// ==== Create crop object and crop to JSON ====
	cropJsonBytes, err := json.Marshal(crop)
	if err != nil {
		return shim.Error(err.Error())
	}

	

	// === Save marble to state ===
	err = stub.PutState(cropnamev, cropJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Crop update done Return success ====
	fmt.Println("- end update crop successful")
	return shim.Success(nil)
}

// Quary Crop
// =========================================================================================
func (t *SimpleChaincode) queryCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *SimpleChaincode) getHistoryForCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cropName := args[0]

	fmt.Printf("- start getHistoryForCrop: %s\n", cropName)

	resultsIterator, err := stub.GetHistoryForKey(cropName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the crop
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON crop)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForCrop returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===============================================
// readCrop - read a Crop from chaincode state
// ===============================================
func (t *SimpleChaincode) readCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Crop to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the crop from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove a crop key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var cropJSON Crop
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	cropName := args[0]

	// to maintain the owner~name index, we need to read the crop first and get its owner
	valAsbytes, err := stub.GetState(cropName) //get the crop from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + cropName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"crop does not exist: " + cropName + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &cropJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + cropName + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(cropName) //remove the crop from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "owner~name"
	ownerNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{cropJSON.Owner, cropJSON.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	//  Delete index entry to state.
	err = stub.DelState(ownerNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// ===========================================================
// irrigation
// ===========================================================
func (t *SimpleChaincode) irrigation(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cropName := args[0]
	newIrrigationValue, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("Unable to parse boolean")
	}
	fmt.Println("- start irrigation value update", cropName, newIrrigationValue)

	cropAsBytes, err := stub.GetState(cropName)
	if err != nil {
		return shim.Error("Failed to get crop:" + err.Error())
	} else if cropAsBytes == nil {
		return shim.Error("crop does not exist")
	}

	cropIrrigation := Crop{}
	err = json.Unmarshal(cropAsBytes, &cropIrrigation) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	cropIrrigation.Irrigation = newIrrigationValue //change the owner

	cropJSONasBytes, _ := json.Marshal(cropIrrigation)
	err = stub.PutState(cropName, cropJSONasBytes) //rewrite the crop
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end irrigation value update(successful)")
	return shim.Success(nil)
}

// ===========================================================
// add fertilizer
// ===========================================================
func (t *SimpleChaincode) addFertilizer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cropName := args[0]
	newFertilizerValue, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("Unable to parse boolean")
	}
	fmt.Println("- start fertilization value update ", cropName, newFertilizerValue)

	cropAsBytes, err := stub.GetState(cropName)
	if err != nil {
		return shim.Error("Failed to get crop:" + err.Error())
	} else if cropAsBytes == nil {
		return shim.Error("crop does not exist")
	}

	cropFertilization := Crop{}
	err = json.Unmarshal(cropAsBytes, &cropFertilization) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	cropFertilization.AddFertilizer = newFertilizerValue //change the fertilizer value

	cropJSONasBytes, _ := json.Marshal(cropFertilization)
	err = stub.PutState(cropName, cropJSONasBytes) //rewrite the crop
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end addFertilizer value update(successful)")
	return shim.Success(nil)
}

// ===========================================================
// Add pesticide
// ===========================================================
func (t *SimpleChaincode) applyPesticide(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cropName := args[0]
	newPesticideValue, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("Unable to parse boolean")
	}
	fmt.Println("- start applyPesticide value update ", cropName, newPesticideValue)

	cropAsBytes, err := stub.GetState(cropName)
	if err != nil {
		return shim.Error("Failed to get crop:" + err.Error())
	} else if cropAsBytes == nil {
		return shim.Error("crop does not exist")
	}

	cropPesticideAddition := Crop{}
	err = json.Unmarshal(cropAsBytes, &cropPesticideAddition) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	cropPesticideAddition.ApplyPesticide = newPesticideValue //change the ApplyPesticide value

	cropJSONasBytes, _ := json.Marshal(cropPesticideAddition)
	err = stub.PutState(cropName, cropJSONasBytes) //rewrite the crop
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end applyPesticide value update(successful)")
	return shim.Success(nil)
}

// ===========================================================
// Harvesting
// ===========================================================
func (t *SimpleChaincode) harvest(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cropName := args[0]
	newHarvestValue, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("Unable to parse boolean")
	}
	fmt.Println("- start harvest value update ", cropName, newHarvestValue)

	cropAsBytes, err := stub.GetState(cropName)
	if err != nil {
		return shim.Error("Failed to get crop:" + err.Error())
	} else if cropAsBytes == nil {
		return shim.Error("crop does not exist")
	}

	cropHarvest := Crop{}
	err = json.Unmarshal(cropAsBytes, &cropHarvest) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	cropHarvest.Harvesting = newHarvestValue //change the harvest value

	cropJSONasBytes, _ := json.Marshal(cropHarvest)
	err = stub.PutState(cropName, cropJSONasBytes) //rewrite the crop
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end Harvest value update(successful)")
	return shim.Success(nil)
}
