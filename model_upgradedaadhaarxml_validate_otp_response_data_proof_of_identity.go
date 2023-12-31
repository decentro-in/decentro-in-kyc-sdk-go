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

// UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity struct for UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity
type UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity struct {
	Dob *string `json:"dob,omitempty"`
	HashedEmail *string `json:"hashedEmail,omitempty"`
	Gender *string `json:"gender,omitempty"`
	HashedMobileNumber *string `json:"hashedMobileNumber,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity instantiates a new UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity() *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity {
	this := UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity{}
	return &this
}

// NewUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentityWithDefaults instantiates a new UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentityWithDefaults() *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity {
	this := UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) SetDob(v string) {
	o.Dob = &v
}

// GetHashedEmail returns the HashedEmail field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetHashedEmail() string {
	if o == nil || isNil(o.HashedEmail) {
		var ret string
		return ret
	}
	return *o.HashedEmail
}

// GetHashedEmailOk returns a tuple with the HashedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetHashedEmailOk() (*string, bool) {
	if o == nil || isNil(o.HashedEmail) {
    return nil, false
	}
	return o.HashedEmail, true
}

// HasHashedEmail returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) HasHashedEmail() bool {
	if o != nil && !isNil(o.HashedEmail) {
		return true
	}

	return false
}

// SetHashedEmail gets a reference to the given string and assigns it to the HashedEmail field.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) SetHashedEmail(v string) {
	o.HashedEmail = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) SetGender(v string) {
	o.Gender = &v
}

// GetHashedMobileNumber returns the HashedMobileNumber field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetHashedMobileNumber() string {
	if o == nil || isNil(o.HashedMobileNumber) {
		var ret string
		return ret
	}
	return *o.HashedMobileNumber
}

// GetHashedMobileNumberOk returns a tuple with the HashedMobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetHashedMobileNumberOk() (*string, bool) {
	if o == nil || isNil(o.HashedMobileNumber) {
    return nil, false
	}
	return o.HashedMobileNumber, true
}

// HasHashedMobileNumber returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) HasHashedMobileNumber() bool {
	if o != nil && !isNil(o.HashedMobileNumber) {
		return true
	}

	return false
}

// SetHashedMobileNumber gets a reference to the given string and assigns it to the HashedMobileNumber field.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) SetHashedMobileNumber(v string) {
	o.HashedMobileNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) SetName(v string) {
	o.Name = &v
}

func (o UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
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

type NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity struct {
	value *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity
	isSet bool
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) Get() *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) Set(val *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity(val *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) *NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity {
	return &NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


