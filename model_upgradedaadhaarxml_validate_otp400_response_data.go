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

// UpgradedaadhaarxmlValidateOtp400ResponseData struct for UpgradedaadhaarxmlValidateOtp400ResponseData
type UpgradedaadhaarxmlValidateOtp400ResponseData struct {
	AadhaarReferenceNumber *string `json:"aadhaarReferenceNumber,omitempty"`
	Image *string `json:"image,omitempty"`
	ProofOfAddress *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress `json:"proofOfAddress,omitempty"`
	ProofOfIdentity *UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity `json:"proofOfIdentity,omitempty"`
}

// NewUpgradedaadhaarxmlValidateOtp400ResponseData instantiates a new UpgradedaadhaarxmlValidateOtp400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlValidateOtp400ResponseData() *UpgradedaadhaarxmlValidateOtp400ResponseData {
	this := UpgradedaadhaarxmlValidateOtp400ResponseData{}
	return &this
}

// NewUpgradedaadhaarxmlValidateOtp400ResponseDataWithDefaults instantiates a new UpgradedaadhaarxmlValidateOtp400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlValidateOtp400ResponseDataWithDefaults() *UpgradedaadhaarxmlValidateOtp400ResponseData {
	this := UpgradedaadhaarxmlValidateOtp400ResponseData{}
	return &this
}

// GetAadhaarReferenceNumber returns the AadhaarReferenceNumber field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetAadhaarReferenceNumber() string {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarReferenceNumber
}

// GetAadhaarReferenceNumberOk returns a tuple with the AadhaarReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetAadhaarReferenceNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
    return nil, false
	}
	return o.AadhaarReferenceNumber, true
}

// HasAadhaarReferenceNumber returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) HasAadhaarReferenceNumber() bool {
	if o != nil && !isNil(o.AadhaarReferenceNumber) {
		return true
	}

	return false
}

// SetAadhaarReferenceNumber gets a reference to the given string and assigns it to the AadhaarReferenceNumber field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) SetAadhaarReferenceNumber(v string) {
	o.AadhaarReferenceNumber = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) SetImage(v string) {
	o.Image = &v
}

// GetProofOfAddress returns the ProofOfAddress field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetProofOfAddress() UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	if o == nil || isNil(o.ProofOfAddress) {
		var ret UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress
		return ret
	}
	return *o.ProofOfAddress
}

// GetProofOfAddressOk returns a tuple with the ProofOfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetProofOfAddressOk() (*UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress, bool) {
	if o == nil || isNil(o.ProofOfAddress) {
    return nil, false
	}
	return o.ProofOfAddress, true
}

// HasProofOfAddress returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) HasProofOfAddress() bool {
	if o != nil && !isNil(o.ProofOfAddress) {
		return true
	}

	return false
}

// SetProofOfAddress gets a reference to the given UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress and assigns it to the ProofOfAddress field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) SetProofOfAddress(v UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress) {
	o.ProofOfAddress = &v
}

// GetProofOfIdentity returns the ProofOfIdentity field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetProofOfIdentity() UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	if o == nil || isNil(o.ProofOfIdentity) {
		var ret UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity
		return ret
	}
	return *o.ProofOfIdentity
}

// GetProofOfIdentityOk returns a tuple with the ProofOfIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) GetProofOfIdentityOk() (*UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity, bool) {
	if o == nil || isNil(o.ProofOfIdentity) {
    return nil, false
	}
	return o.ProofOfIdentity, true
}

// HasProofOfIdentity returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) HasProofOfIdentity() bool {
	if o != nil && !isNil(o.ProofOfIdentity) {
		return true
	}

	return false
}

// SetProofOfIdentity gets a reference to the given UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity and assigns it to the ProofOfIdentity field.
func (o *UpgradedaadhaarxmlValidateOtp400ResponseData) SetProofOfIdentity(v UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity) {
	o.ProofOfIdentity = &v
}

func (o UpgradedaadhaarxmlValidateOtp400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarReferenceNumber) {
		toSerialize["aadhaarReferenceNumber"] = o.AadhaarReferenceNumber
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.ProofOfAddress) {
		toSerialize["proofOfAddress"] = o.ProofOfAddress
	}
	if !isNil(o.ProofOfIdentity) {
		toSerialize["proofOfIdentity"] = o.ProofOfIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradedaadhaarxmlValidateOtp400ResponseData struct {
	value *UpgradedaadhaarxmlValidateOtp400ResponseData
	isSet bool
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseData) Get() *UpgradedaadhaarxmlValidateOtp400ResponseData {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseData) Set(val *UpgradedaadhaarxmlValidateOtp400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlValidateOtp400ResponseData(val *UpgradedaadhaarxmlValidateOtp400ResponseData) *NullableUpgradedaadhaarxmlValidateOtp400ResponseData {
	return &NullableUpgradedaadhaarxmlValidateOtp400ResponseData{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlValidateOtp400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlValidateOtp400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


