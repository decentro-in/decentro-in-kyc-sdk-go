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

// AadhaarxmlValidateOtp400ResponseData struct for AadhaarxmlValidateOtp400ResponseData
type AadhaarxmlValidateOtp400ResponseData struct {
	AadhaarFile *AadhaarxmlValidateOtp400ResponseDataAadhaarFile `json:"aadhaarFile,omitempty"`
	AadhaarReferenceNumber *string `json:"aadhaarReferenceNumber,omitempty"`
	Image *string `json:"image,omitempty"`
	ProofOfAddress *AadhaarxmlValidateOtp400ResponseDataProofOfAddress `json:"proofOfAddress,omitempty"`
	ProofOfIdentity *AadhaarxmlValidateOtp400ResponseDataProofOfIdentity `json:"proofOfIdentity,omitempty"`
}

// NewAadhaarxmlValidateOtp400ResponseData instantiates a new AadhaarxmlValidateOtp400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtp400ResponseData() *AadhaarxmlValidateOtp400ResponseData {
	this := AadhaarxmlValidateOtp400ResponseData{}
	return &this
}

// NewAadhaarxmlValidateOtp400ResponseDataWithDefaults instantiates a new AadhaarxmlValidateOtp400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtp400ResponseDataWithDefaults() *AadhaarxmlValidateOtp400ResponseData {
	this := AadhaarxmlValidateOtp400ResponseData{}
	return &this
}

// GetAadhaarFile returns the AadhaarFile field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseData) GetAadhaarFile() AadhaarxmlValidateOtp400ResponseDataAadhaarFile {
	if o == nil || isNil(o.AadhaarFile) {
		var ret AadhaarxmlValidateOtp400ResponseDataAadhaarFile
		return ret
	}
	return *o.AadhaarFile
}

// GetAadhaarFileOk returns a tuple with the AadhaarFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) GetAadhaarFileOk() (*AadhaarxmlValidateOtp400ResponseDataAadhaarFile, bool) {
	if o == nil || isNil(o.AadhaarFile) {
    return nil, false
	}
	return o.AadhaarFile, true
}

// HasAadhaarFile returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) HasAadhaarFile() bool {
	if o != nil && !isNil(o.AadhaarFile) {
		return true
	}

	return false
}

// SetAadhaarFile gets a reference to the given AadhaarxmlValidateOtp400ResponseDataAadhaarFile and assigns it to the AadhaarFile field.
func (o *AadhaarxmlValidateOtp400ResponseData) SetAadhaarFile(v AadhaarxmlValidateOtp400ResponseDataAadhaarFile) {
	o.AadhaarFile = &v
}

// GetAadhaarReferenceNumber returns the AadhaarReferenceNumber field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseData) GetAadhaarReferenceNumber() string {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarReferenceNumber
}

// GetAadhaarReferenceNumberOk returns a tuple with the AadhaarReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) GetAadhaarReferenceNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
    return nil, false
	}
	return o.AadhaarReferenceNumber, true
}

// HasAadhaarReferenceNumber returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) HasAadhaarReferenceNumber() bool {
	if o != nil && !isNil(o.AadhaarReferenceNumber) {
		return true
	}

	return false
}

// SetAadhaarReferenceNumber gets a reference to the given string and assigns it to the AadhaarReferenceNumber field.
func (o *AadhaarxmlValidateOtp400ResponseData) SetAadhaarReferenceNumber(v string) {
	o.AadhaarReferenceNumber = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseData) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AadhaarxmlValidateOtp400ResponseData) SetImage(v string) {
	o.Image = &v
}

// GetProofOfAddress returns the ProofOfAddress field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseData) GetProofOfAddress() AadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	if o == nil || isNil(o.ProofOfAddress) {
		var ret AadhaarxmlValidateOtp400ResponseDataProofOfAddress
		return ret
	}
	return *o.ProofOfAddress
}

// GetProofOfAddressOk returns a tuple with the ProofOfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) GetProofOfAddressOk() (*AadhaarxmlValidateOtp400ResponseDataProofOfAddress, bool) {
	if o == nil || isNil(o.ProofOfAddress) {
    return nil, false
	}
	return o.ProofOfAddress, true
}

// HasProofOfAddress returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) HasProofOfAddress() bool {
	if o != nil && !isNil(o.ProofOfAddress) {
		return true
	}

	return false
}

// SetProofOfAddress gets a reference to the given AadhaarxmlValidateOtp400ResponseDataProofOfAddress and assigns it to the ProofOfAddress field.
func (o *AadhaarxmlValidateOtp400ResponseData) SetProofOfAddress(v AadhaarxmlValidateOtp400ResponseDataProofOfAddress) {
	o.ProofOfAddress = &v
}

// GetProofOfIdentity returns the ProofOfIdentity field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseData) GetProofOfIdentity() AadhaarxmlValidateOtp400ResponseDataProofOfIdentity {
	if o == nil || isNil(o.ProofOfIdentity) {
		var ret AadhaarxmlValidateOtp400ResponseDataProofOfIdentity
		return ret
	}
	return *o.ProofOfIdentity
}

// GetProofOfIdentityOk returns a tuple with the ProofOfIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) GetProofOfIdentityOk() (*AadhaarxmlValidateOtp400ResponseDataProofOfIdentity, bool) {
	if o == nil || isNil(o.ProofOfIdentity) {
    return nil, false
	}
	return o.ProofOfIdentity, true
}

// HasProofOfIdentity returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseData) HasProofOfIdentity() bool {
	if o != nil && !isNil(o.ProofOfIdentity) {
		return true
	}

	return false
}

// SetProofOfIdentity gets a reference to the given AadhaarxmlValidateOtp400ResponseDataProofOfIdentity and assigns it to the ProofOfIdentity field.
func (o *AadhaarxmlValidateOtp400ResponseData) SetProofOfIdentity(v AadhaarxmlValidateOtp400ResponseDataProofOfIdentity) {
	o.ProofOfIdentity = &v
}

func (o AadhaarxmlValidateOtp400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarFile) {
		toSerialize["aadhaarFile"] = o.AadhaarFile
	}
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

type NullableAadhaarxmlValidateOtp400ResponseData struct {
	value *AadhaarxmlValidateOtp400ResponseData
	isSet bool
}

func (v NullableAadhaarxmlValidateOtp400ResponseData) Get() *AadhaarxmlValidateOtp400ResponseData {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtp400ResponseData) Set(val *AadhaarxmlValidateOtp400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtp400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtp400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtp400ResponseData(val *AadhaarxmlValidateOtp400ResponseData) *NullableAadhaarxmlValidateOtp400ResponseData {
	return &NullableAadhaarxmlValidateOtp400ResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtp400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtp400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


