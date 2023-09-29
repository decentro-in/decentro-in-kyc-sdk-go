/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"encoding/json"
	"fmt"
)

// ValidateResponseKycResultAddressState - struct for ValidateResponseKycResultAddressState
type ValidateResponseKycResultAddressState struct {
	ArrayOfArrayOfString *[][]string
	String *string
}

// [][]stringAsValidateResponseKycResultAddressState is a convenience function that returns [][]string wrapped in ValidateResponseKycResultAddressState
func ArrayOfArrayOfStringAsValidateResponseKycResultAddressState(v *[][]string) ValidateResponseKycResultAddressState {
	return ValidateResponseKycResultAddressState{
		ArrayOfArrayOfString: v,
	}
}

// stringAsValidateResponseKycResultAddressState is a convenience function that returns string wrapped in ValidateResponseKycResultAddressState
func StringAsValidateResponseKycResultAddressState(v *string) ValidateResponseKycResultAddressState {
	return ValidateResponseKycResultAddressState{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ValidateResponseKycResultAddressState) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfArrayOfString)
	if err == nil {
		jsonArrayOfArrayOfString, _ := json.Marshal(dst.ArrayOfArrayOfString)
		if string(jsonArrayOfArrayOfString) == "{}" { // empty struct
			dst.ArrayOfArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ValidateResponseKycResultAddressState)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ValidateResponseKycResultAddressState)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ValidateResponseKycResultAddressState) MarshalJSON() ([]byte, error) {
	if src.ArrayOfArrayOfString != nil {
		return json.Marshal(&src.ArrayOfArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ValidateResponseKycResultAddressState) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfArrayOfString != nil {
		return obj.ArrayOfArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableValidateResponseKycResultAddressState struct {
	value *ValidateResponseKycResultAddressState
	isSet bool
}

func (v NullableValidateResponseKycResultAddressState) Get() *ValidateResponseKycResultAddressState {
	return v.value
}

func (v *NullableValidateResponseKycResultAddressState) Set(val *ValidateResponseKycResultAddressState) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultAddressState) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultAddressState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultAddressState(val *ValidateResponseKycResultAddressState) *NullableValidateResponseKycResultAddressState {
	return &NullableValidateResponseKycResultAddressState{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultAddressState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultAddressState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


