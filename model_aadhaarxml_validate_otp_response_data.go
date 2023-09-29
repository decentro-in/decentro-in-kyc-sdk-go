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

// AadhaarxmlValidateOtpResponseData struct for AadhaarxmlValidateOtpResponseData
type AadhaarxmlValidateOtpResponseData struct {
	AadhaarReferenceNumber *string `json:"aadhaarReferenceNumber,omitempty"`
	ProofOfIdentity *AadhaarxmlValidateOtpResponseDataProofOfIdentity `json:"proofOfIdentity,omitempty"`
	ProofOfAddress *AadhaarxmlValidateOtpResponseDataProofOfAddress `json:"proofOfAddress,omitempty"`
	Image *string `json:"image,omitempty"`
	AadhaarFile *AadhaarxmlValidateOtpResponseDataAadhaarFile `json:"aadhaarFile,omitempty"`
}

// NewAadhaarxmlValidateOtpResponseData instantiates a new AadhaarxmlValidateOtpResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtpResponseData() *AadhaarxmlValidateOtpResponseData {
	this := AadhaarxmlValidateOtpResponseData{}
	return &this
}

// NewAadhaarxmlValidateOtpResponseDataWithDefaults instantiates a new AadhaarxmlValidateOtpResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtpResponseDataWithDefaults() *AadhaarxmlValidateOtpResponseData {
	this := AadhaarxmlValidateOtpResponseData{}
	return &this
}

// GetAadhaarReferenceNumber returns the AadhaarReferenceNumber field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseData) GetAadhaarReferenceNumber() string {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarReferenceNumber
}

// GetAadhaarReferenceNumberOk returns a tuple with the AadhaarReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseData) GetAadhaarReferenceNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
    return nil, false
	}
	return o.AadhaarReferenceNumber, true
}

// HasAadhaarReferenceNumber returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseData) HasAadhaarReferenceNumber() bool {
	if o != nil && !isNil(o.AadhaarReferenceNumber) {
		return true
	}

	return false
}

// SetAadhaarReferenceNumber gets a reference to the given string and assigns it to the AadhaarReferenceNumber field.
func (o *AadhaarxmlValidateOtpResponseData) SetAadhaarReferenceNumber(v string) {
	o.AadhaarReferenceNumber = &v
}

// GetProofOfIdentity returns the ProofOfIdentity field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseData) GetProofOfIdentity() AadhaarxmlValidateOtpResponseDataProofOfIdentity {
	if o == nil || isNil(o.ProofOfIdentity) {
		var ret AadhaarxmlValidateOtpResponseDataProofOfIdentity
		return ret
	}
	return *o.ProofOfIdentity
}

// GetProofOfIdentityOk returns a tuple with the ProofOfIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseData) GetProofOfIdentityOk() (*AadhaarxmlValidateOtpResponseDataProofOfIdentity, bool) {
	if o == nil || isNil(o.ProofOfIdentity) {
    return nil, false
	}
	return o.ProofOfIdentity, true
}

// HasProofOfIdentity returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseData) HasProofOfIdentity() bool {
	if o != nil && !isNil(o.ProofOfIdentity) {
		return true
	}

	return false
}

// SetProofOfIdentity gets a reference to the given AadhaarxmlValidateOtpResponseDataProofOfIdentity and assigns it to the ProofOfIdentity field.
func (o *AadhaarxmlValidateOtpResponseData) SetProofOfIdentity(v AadhaarxmlValidateOtpResponseDataProofOfIdentity) {
	o.ProofOfIdentity = &v
}

// GetProofOfAddress returns the ProofOfAddress field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseData) GetProofOfAddress() AadhaarxmlValidateOtpResponseDataProofOfAddress {
	if o == nil || isNil(o.ProofOfAddress) {
		var ret AadhaarxmlValidateOtpResponseDataProofOfAddress
		return ret
	}
	return *o.ProofOfAddress
}

// GetProofOfAddressOk returns a tuple with the ProofOfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseData) GetProofOfAddressOk() (*AadhaarxmlValidateOtpResponseDataProofOfAddress, bool) {
	if o == nil || isNil(o.ProofOfAddress) {
    return nil, false
	}
	return o.ProofOfAddress, true
}

// HasProofOfAddress returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseData) HasProofOfAddress() bool {
	if o != nil && !isNil(o.ProofOfAddress) {
		return true
	}

	return false
}

// SetProofOfAddress gets a reference to the given AadhaarxmlValidateOtpResponseDataProofOfAddress and assigns it to the ProofOfAddress field.
func (o *AadhaarxmlValidateOtpResponseData) SetProofOfAddress(v AadhaarxmlValidateOtpResponseDataProofOfAddress) {
	o.ProofOfAddress = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseData) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseData) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseData) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AadhaarxmlValidateOtpResponseData) SetImage(v string) {
	o.Image = &v
}

// GetAadhaarFile returns the AadhaarFile field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseData) GetAadhaarFile() AadhaarxmlValidateOtpResponseDataAadhaarFile {
	if o == nil || isNil(o.AadhaarFile) {
		var ret AadhaarxmlValidateOtpResponseDataAadhaarFile
		return ret
	}
	return *o.AadhaarFile
}

// GetAadhaarFileOk returns a tuple with the AadhaarFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseData) GetAadhaarFileOk() (*AadhaarxmlValidateOtpResponseDataAadhaarFile, bool) {
	if o == nil || isNil(o.AadhaarFile) {
    return nil, false
	}
	return o.AadhaarFile, true
}

// HasAadhaarFile returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseData) HasAadhaarFile() bool {
	if o != nil && !isNil(o.AadhaarFile) {
		return true
	}

	return false
}

// SetAadhaarFile gets a reference to the given AadhaarxmlValidateOtpResponseDataAadhaarFile and assigns it to the AadhaarFile field.
func (o *AadhaarxmlValidateOtpResponseData) SetAadhaarFile(v AadhaarxmlValidateOtpResponseDataAadhaarFile) {
	o.AadhaarFile = &v
}

func (o AadhaarxmlValidateOtpResponseData) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AadhaarFile) {
		toSerialize["aadhaarFile"] = o.AadhaarFile
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlValidateOtpResponseData struct {
	value *AadhaarxmlValidateOtpResponseData
	isSet bool
}

func (v NullableAadhaarxmlValidateOtpResponseData) Get() *AadhaarxmlValidateOtpResponseData {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtpResponseData) Set(val *AadhaarxmlValidateOtpResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtpResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtpResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtpResponseData(val *AadhaarxmlValidateOtpResponseData) *NullableAadhaarxmlValidateOtpResponseData {
	return &NullableAadhaarxmlValidateOtpResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtpResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtpResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


