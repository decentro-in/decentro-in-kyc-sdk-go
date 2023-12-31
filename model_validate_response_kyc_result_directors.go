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

// ValidateResponseKycResultDirectors - struct for ValidateResponseKycResultDirectors
type ValidateResponseKycResultDirectors struct {
	ArrayOfValidateResponseKycResultDirectorsOneOfInner *[]ValidateResponseKycResultDirectorsOneOfInner
	String *string
}

// []ValidateResponseKycResultDirectorsOneOfInnerAsValidateResponseKycResultDirectors is a convenience function that returns []ValidateResponseKycResultDirectorsOneOfInner wrapped in ValidateResponseKycResultDirectors
func ArrayOfValidateResponseKycResultDirectorsOneOfInnerAsValidateResponseKycResultDirectors(v *[]ValidateResponseKycResultDirectorsOneOfInner) ValidateResponseKycResultDirectors {
	return ValidateResponseKycResultDirectors{
		ArrayOfValidateResponseKycResultDirectorsOneOfInner: v,
	}
}

// stringAsValidateResponseKycResultDirectors is a convenience function that returns string wrapped in ValidateResponseKycResultDirectors
func StringAsValidateResponseKycResultDirectors(v *string) ValidateResponseKycResultDirectors {
	return ValidateResponseKycResultDirectors{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ValidateResponseKycResultDirectors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfValidateResponseKycResultDirectorsOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfValidateResponseKycResultDirectorsOneOfInner)
	if err == nil {
		jsonArrayOfValidateResponseKycResultDirectorsOneOfInner, _ := json.Marshal(dst.ArrayOfValidateResponseKycResultDirectorsOneOfInner)
		if string(jsonArrayOfValidateResponseKycResultDirectorsOneOfInner) == "{}" { // empty struct
			dst.ArrayOfValidateResponseKycResultDirectorsOneOfInner = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfValidateResponseKycResultDirectorsOneOfInner = nil
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
		dst.ArrayOfValidateResponseKycResultDirectorsOneOfInner = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ValidateResponseKycResultDirectors)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ValidateResponseKycResultDirectors)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ValidateResponseKycResultDirectors) MarshalJSON() ([]byte, error) {
	if src.ArrayOfValidateResponseKycResultDirectorsOneOfInner != nil {
		return json.Marshal(&src.ArrayOfValidateResponseKycResultDirectorsOneOfInner)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ValidateResponseKycResultDirectors) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfValidateResponseKycResultDirectorsOneOfInner != nil {
		return obj.ArrayOfValidateResponseKycResultDirectorsOneOfInner
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableValidateResponseKycResultDirectors struct {
	value *ValidateResponseKycResultDirectors
	isSet bool
}

func (v NullableValidateResponseKycResultDirectors) Get() *ValidateResponseKycResultDirectors {
	return v.value
}

func (v *NullableValidateResponseKycResultDirectors) Set(val *ValidateResponseKycResultDirectors) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultDirectors) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultDirectors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultDirectors(val *ValidateResponseKycResultDirectors) *NullableValidateResponseKycResultDirectors {
	return &NullableValidateResponseKycResultDirectors{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultDirectors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultDirectors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


