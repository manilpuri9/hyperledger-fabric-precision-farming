package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//Crop struct is the defination of datastructure of crop
type Crop struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Quantity int    `json:"quantity"`
	FarmInfo struct {
		GeoLocation struct {
			Latitude  []interface{} `json:"latitude"`
			Longitude []interface{} `json:"longitude"`
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
	if function == "initCrop" { //create a new marble
		return t.initCrop(stub, args)
	} /*
		else if function == "transferCrop" { //change owner of a crop marble
			return t.transferCrop(stub, args)
		} else if function == "delete" { //delete a crop
			return t.delete(stub, args)
		} else if function == "readCrop" { //read a crop
			return t.readCrop(stub, args)
		} else if function == "queryCropByOwner" { //find crop for owner X using rich query
			return t.queryCropByOwner(stub, args)
		} else if function == "queryCrops" { //find crops based on an ad hoc rich query
			return t.queryCrops(stub, args)
		} else if function == "getHistoryForCrop" { //get history of values for a crop
			return t.getHistoryForCrop(stub, args)
		} else if function == "getCropsByRange" { //get marbles based on range query
			return t.getCropsByRange(stub, args)
		}
	*/

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init crop")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	cropName := args[0]
	owner := args[1]

	// ==== Check if crop already exists ====
	cropAsBytes, err := stub.GetState(cropName)
	if err != nil {
		return shim.Error("Failed to get crop: " + err.Error())
	} else if cropAsBytes != nil {
		fmt.Println("This crop already exists: " + cropName)
		return shim.Error("This crop already exists: " + cropName)
	}

	// ==== Create crop object and marshal to JSON ====
	//Alternatively, build the crop json string manually if you don't want to use struct marshalling
	cropJSONasString := `{"docType":"Crop",  "Name": "` + cropName + `",  "owner": "` + owner + `"}`
	cropJSONasBytes := []byte(cropJSONasString)

	// === Save crop to state ===
	err = stub.PutState(cropName, cropJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init marble")
	return shim.Success(nil)
}
