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

// UpgradedaadhaarxmlValidateOtpResponseData struct for UpgradedaadhaarxmlValidateOtpResponseData
type UpgradedaadhaarxmlValidateOtpResponseData struct {
	AadhaarReferenceNumber *string `json:"aadhaarReferenceNumber,omitempty"`
	ProofOfIdentity *UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity `json:"proofOfIdentity,omitempty"`
	ProofOfAddress *UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress `json:"proofOfAddress,omitempty"`
	Image *string `json:"image,omitempty"`
}

// NewUpgradedaadhaarxmlValidateOtpResponseData instantiates a new UpgradedaadhaarxmlValidateOtpResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlValidateOtpResponseData() *UpgradedaadhaarxmlValidateOtpResponseData {
	this := UpgradedaadhaarxmlValidateOtpResponseData{}
	return &this
}

// NewUpgradedaadhaarxmlValidateOtpResponseDataWithDefaults instantiates a new UpgradedaadhaarxmlValidateOtpResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlValidateOtpResponseDataWithDefaults() *UpgradedaadhaarxmlValidateOtpResponseData {
	this := UpgradedaadhaarxmlValidateOtpResponseData{}
	return &this
}

// GetAadhaarReferenceNumber returns the AadhaarReferenceNumber field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetAadhaarReferenceNumber() string {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarReferenceNumber
}

// GetAadhaarReferenceNumberOk returns a tuple with the AadhaarReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetAadhaarReferenceNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
    return nil, false
	}
	return o.AadhaarReferenceNumber, true
}

// HasAadhaarReferenceNumber returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) HasAadhaarReferenceNumber() bool {
	if o != nil && !isNil(o.AadhaarReferenceNumber) {
		return true
	}

	return false
}

// SetAadhaarReferenceNumber gets a reference to the given string and assigns it to the AadhaarReferenceNumber field.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) SetAadhaarReferenceNumber(v string) {
	o.AadhaarReferenceNumber = &v
}

// GetProofOfIdentity returns the ProofOfIdentity field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetProofOfIdentity() UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity {
	if o == nil || isNil(o.ProofOfIdentity) {
		var ret UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity
		return ret
	}
	return *o.ProofOfIdentity
}

// GetProofOfIdentityOk returns a tuple with the ProofOfIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetProofOfIdentityOk() (*UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity, bool) {
	if o == nil || isNil(o.ProofOfIdentity) {
    return nil, false
	}
	return o.ProofOfIdentity, true
}

// HasProofOfIdentity returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) HasProofOfIdentity() bool {
	if o != nil && !isNil(o.ProofOfIdentity) {
		return true
	}

	return false
}

// SetProofOfIdentity gets a reference to the given UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity and assigns it to the ProofOfIdentity field.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) SetProofOfIdentity(v UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity) {
	o.ProofOfIdentity = &v
}

// GetProofOfAddress returns the ProofOfAddress field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetProofOfAddress() UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress {
	if o == nil || isNil(o.ProofOfAddress) {
		var ret UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress
		return ret
	}
	return *o.ProofOfAddress
}

// GetProofOfAddressOk returns a tuple with the ProofOfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetProofOfAddressOk() (*UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress, bool) {
	if o == nil || isNil(o.ProofOfAddress) {
    return nil, false
	}
	return o.ProofOfAddress, true
}

// HasProofOfAddress returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) HasProofOfAddress() bool {
	if o != nil && !isNil(o.ProofOfAddress) {
		return true
	}

	return false
}

// SetProofOfAddress gets a reference to the given UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress and assigns it to the ProofOfAddress field.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) SetProofOfAddress(v UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress) {
	o.ProofOfAddress = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *UpgradedaadhaarxmlValidateOtpResponseData) SetImage(v string) {
	o.Image = &v
}

func (o UpgradedaadhaarxmlValidateOtpResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarReferenceNumber) {
		toSerialize["aadhaarReferenceNumber"] = o.AadhaarReferenceNumber
	}
	if !isNil(o.ProofOfIdentity) {
		toSerialize["proofOfIdentity"] = o.ProofOfIdentity
	}
	if !isNil(o.ProofOfAddress) {
		toSerialize["proofOfAddress"] = o.ProofOfAddress
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradedaadhaarxmlValidateOtpResponseData struct {
	value *UpgradedaadhaarxmlValidateOtpResponseData
	isSet bool
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseData) Get() *UpgradedaadhaarxmlValidateOtpResponseData {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseData) Set(val *UpgradedaadhaarxmlValidateOtpResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlValidateOtpResponseData(val *UpgradedaadhaarxmlValidateOtpResponseData) *NullableUpgradedaadhaarxmlValidateOtpResponseData {
	return &NullableUpgradedaadhaarxmlValidateOtpResponseData{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlValidateOtpResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlValidateOtpResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


