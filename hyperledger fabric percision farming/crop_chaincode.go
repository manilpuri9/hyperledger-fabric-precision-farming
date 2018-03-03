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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	} else if function == "readMarble" { //read a Crop
		return t.readCrop(stub, args)
	} else if function == "queryCrop" { //find Crop based on an ad hoc rich query
		return t.queryCrop(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0       1       2     3
	// "asdf", "blue", "35", "bob"
	if len(args) != 20 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init crop")
	// if len(args[0]) <= 0 {
	// 	return shim.Error("1st argument must be a non-empty string")
	// }
	// if len(args[1]) <= 0 {
	// 	return shim.Error("2nd argument must be a non-empty string")
	// }
	// if len(args[2]) <= 0 {
	// 	return shim.Error("3rd argument must be a non-empty string")
	// }
	// if len(args[3]) <= 0 {
	// 	return shim.Error("4th argument must be a non-empty string")
	// }
	marbleName := args[0]
	color := strings.ToLower(args[1])
	owner := strings.ToLower(args[3])
	size, err := strconv.Atoi(args[2])

	//////////////////////////////////////////////////////
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	// ==== Check if marble already exists ====
	marbleAsBytes, err := stub.GetState(marbleName)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	} else if marbleAsBytes != nil {
		fmt.Println("This marble already exists: " + marbleName)
		return shim.Error("This marble already exists: " + marbleName)
	}

	// ==== Create marble object and marshal to JSON ====
	objectType := "marble"
	marble := &marble{objectType, marbleName, color, size, owner}
	marbleJSONasBytes, err := json.Marshal(marble)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the marble json string manually if you don't want to use struct marshalling
	//marbleJSONasString := `{"docType":"Marble",  "name": "` + marbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//marbleJSONasBytes := []byte(str)

	// === Save marble to state ===
	err = stub.PutState(marbleName, marbleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the marble to enable color-based range queries, e.g. return all blue marbles ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "color~name"
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{marble.Color, marble.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(colorNameIndexKey, value)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init marble")
	return shim.Success(nil)
}
