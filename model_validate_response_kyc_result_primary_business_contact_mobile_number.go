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

// ValidateResponseKycResultPrimaryBusinessContactMobileNumber - struct for ValidateResponseKycResultPrimaryBusinessContactMobileNumber
type ValidateResponseKycResultPrimaryBusinessContactMobileNumber struct {
	Float32 *float32
	String *string
}

// float32AsValidateResponseKycResultPrimaryBusinessContactMobileNumber is a convenience function that returns float32 wrapped in ValidateResponseKycResultPrimaryBusinessContactMobileNumber
func Float32AsValidateResponseKycResultPrimaryBusinessContactMobileNumber(v *float32) ValidateResponseKycResultPrimaryBusinessContactMobileNumber {
	return ValidateResponseKycResultPrimaryBusinessContactMobileNumber{
		Float32: v,
	}
}

// stringAsValidateResponseKycResultPrimaryBusinessContactMobileNumber is a convenience function that returns string wrapped in ValidateResponseKycResultPrimaryBusinessContactMobileNumber
func StringAsValidateResponseKycResultPrimaryBusinessContactMobileNumber(v *string) ValidateResponseKycResultPrimaryBusinessContactMobileNumber {
	return ValidateResponseKycResultPrimaryBusinessContactMobileNumber{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ValidateResponseKycResultPrimaryBusinessContactMobileNumber) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
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
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ValidateResponseKycResultPrimaryBusinessContactMobileNumber)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ValidateResponseKycResultPrimaryBusinessContactMobileNumber)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ValidateResponseKycResultPrimaryBusinessContactMobileNumber) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ValidateResponseKycResultPrimaryBusinessContactMobileNumber) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber struct {
	value *ValidateResponseKycResultPrimaryBusinessContactMobileNumber
	isSet bool
}

func (v NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) Get() *ValidateResponseKycResultPrimaryBusinessContactMobileNumber {
	return v.value
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) Set(val *ValidateResponseKycResultPrimaryBusinessContactMobileNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultPrimaryBusinessContactMobileNumber(val *ValidateResponseKycResultPrimaryBusinessContactMobileNumber) *NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber {
	return &NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContactMobileNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


