/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"encoding/json"
)

// DigilockerEAadhaarResponseDataProofOfIdentity struct for DigilockerEAadhaarResponseDataProofOfIdentity
type DigilockerEAadhaarResponseDataProofOfIdentity struct {
	Dob *string `json:"dob,omitempty"`
	HashedEmail *string `json:"hashedEmail,omitempty"`
	Gender *string `json:"gender,omitempty"`
	HashedMobileNumber *string `json:"hashedMobileNumber,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewDigilockerEAadhaarResponseDataProofOfIdentity instantiates a new DigilockerEAadhaarResponseDataProofOfIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerEAadhaarResponseDataProofOfIdentity() *DigilockerEAadhaarResponseDataProofOfIdentity {
	this := DigilockerEAadhaarResponseDataProofOfIdentity{}
	return &this
}

// NewDigilockerEAadhaarResponseDataProofOfIdentityWithDefaults instantiates a new DigilockerEAadhaarResponseDataProofOfIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerEAadhaarResponseDataProofOfIdentityWithDefaults() *DigilockerEAadhaarResponseDataProofOfIdentity {
	this := DigilockerEAadhaarResponseDataProofOfIdentity{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) SetDob(v string) {
	o.Dob = &v
}

// GetHashedEmail returns the HashedEmail field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetHashedEmail() string {
	if o == nil || isNil(o.HashedEmail) {
		var ret string
		return ret
	}
	return *o.HashedEmail
}

// GetHashedEmailOk returns a tuple with the HashedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetHashedEmailOk() (*string, bool) {
	if o == nil || isNil(o.HashedEmail) {
    return nil, false
	}
	return o.HashedEmail, true
}

// HasHashedEmail returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) HasHashedEmail() bool {
	if o != nil && !isNil(o.HashedEmail) {
		return true
	}

	return false
}

// SetHashedEmail gets a reference to the given string and assigns it to the HashedEmail field.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) SetHashedEmail(v string) {
	o.HashedEmail = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) SetGender(v string) {
	o.Gender = &v
}

// GetHashedMobileNumber returns the HashedMobileNumber field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetHashedMobileNumber() string {
	if o == nil || isNil(o.HashedMobileNumber) {
		var ret string
		return ret
	}
	return *o.HashedMobileNumber
}

// GetHashedMobileNumberOk returns a tuple with the HashedMobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetHashedMobileNumberOk() (*string, bool) {
	if o == nil || isNil(o.HashedMobileNumber) {
    return nil, false
	}
	return o.HashedMobileNumber, true
}

// HasHashedMobileNumber returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) HasHashedMobileNumber() bool {
	if o != nil && !isNil(o.HashedMobileNumber) {
		return true
	}

	return false
}

// SetHashedMobileNumber gets a reference to the given string and assigns it to the HashedMobileNumber field.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) SetHashedMobileNumber(v string) {
	o.HashedMobileNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DigilockerEAadhaarResponseDataProofOfIdentity) SetName(v string) {
	o.Name = &v
}

func (o DigilockerEAadhaarResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !isNil(o.HashedEmail) {
		toSerialize["hashedEmail"] = o.HashedEmail
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.HashedMobileNumber) {
		toSerialize["hashedMobileNumber"] = o.HashedMobileNumber
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerEAadhaarResponseDataProofOfIdentity struct {
	value *DigilockerEAadhaarResponseDataProofOfIdentity
	isSet bool
}

func (v NullableDigilockerEAadhaarResponseDataProofOfIdentity) Get() *DigilockerEAadhaarResponseDataProofOfIdentity {
	return v.value
}

func (v *NullableDigilockerEAadhaarResponseDataProofOfIdentity) Set(val *DigilockerEAadhaarResponseDataProofOfIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerEAadhaarResponseDataProofOfIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerEAadhaarResponseDataProofOfIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerEAadhaarResponseDataProofOfIdentity(val *DigilockerEAadhaarResponseDataProofOfIdentity) *NullableDigilockerEAadhaarResponseDataProofOfIdentity {
	return &NullableDigilockerEAadhaarResponseDataProofOfIdentity{value: val, isSet: true}
}

func (v NullableDigilockerEAadhaarResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerEAadhaarResponseDataProofOfIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


