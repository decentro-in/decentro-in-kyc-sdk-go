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

// UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity struct for UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity
type UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity struct {
	Dob *string `json:"dob,omitempty"`
	Gender *string `json:"gender,omitempty"`
	HashedEmail *string `json:"hashedEmail,omitempty"`
	HashedMobileNumber *string `json:"hashedMobileNumber,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity instantiates a new UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity() *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	this := UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity{}
	return &this
}

// NewUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentityWithDefaults instantiates a new UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentityWithDefaults() *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	this := UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) SetDob(v string) {
	o.Dob = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) SetGender(v string) {
	o.Gender = &v
}

// GetHashedEmail returns the HashedEmail field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetHashedEmail() string {
	if o == nil || isNil(o.HashedEmail) {
		var ret string
		return ret
	}
	return *o.HashedEmail
}

// GetHashedEmailOk returns a tuple with the HashedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetHashedEmailOk() (*string, bool) {
	if o == nil || isNil(o.HashedEmail) {
    return nil, false
	}
	return o.HashedEmail, true
}

// HasHashedEmail returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) HasHashedEmail() bool {
	if o != nil && !isNil(o.HashedEmail) {
		return true
	}

	return false
}

// SetHashedEmail gets a reference to the given string and assigns it to the HashedEmail field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) SetHashedEmail(v string) {
	o.HashedEmail = &v
}

// GetHashedMobileNumber returns the HashedMobileNumber field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetHashedMobileNumber() string {
	if o == nil || isNil(o.HashedMobileNumber) {
		var ret string
		return ret
	}
	return *o.HashedMobileNumber
}

// GetHashedMobileNumberOk returns a tuple with the HashedMobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetHashedMobileNumberOk() (*string, bool) {
	if o == nil || isNil(o.HashedMobileNumber) {
    return nil, false
	}
	return o.HashedMobileNumber, true
}

// HasHashedMobileNumber returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) HasHashedMobileNumber() bool {
	if o != nil && !isNil(o.HashedMobileNumber) {
		return true
	}

	return false
}

// SetHashedMobileNumber gets a reference to the given string and assigns it to the HashedMobileNumber field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) SetHashedMobileNumber(v string) {
	o.HashedMobileNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) SetName(v string) {
	o.Name = &v
}

func (o UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.HashedEmail) {
		toSerialize["hashedEmail"] = o.HashedEmail
	}
	if !isNil(o.HashedMobileNumber) {
		toSerialize["hashedMobileNumber"] = o.HashedMobileNumber
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity struct {
	value *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity
	isSet bool
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) Get() *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) Set(val *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity(val *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) *NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	return &NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


