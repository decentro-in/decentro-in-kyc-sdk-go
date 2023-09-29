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

// DigilockerEAadhaarResponseData struct for DigilockerEAadhaarResponseData
type DigilockerEAadhaarResponseData struct {
	AadhaarReferenceNumber *string `json:"aadhaarReferenceNumber,omitempty"`
	AadhaarUid *string `json:"aadhaarUid,omitempty"`
	ProofOfIdentity *DigilockerEAadhaarResponseDataProofOfIdentity `json:"proofOfIdentity,omitempty"`
	ProofOfAddress *DigilockerEAadhaarResponseDataProofOfAddress `json:"proofOfAddress,omitempty"`
	Image *string `json:"image,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	Xml *string `json:"xml,omitempty"`
}

// NewDigilockerEAadhaarResponseData instantiates a new DigilockerEAadhaarResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerEAadhaarResponseData() *DigilockerEAadhaarResponseData {
	this := DigilockerEAadhaarResponseData{}
	return &this
}

// NewDigilockerEAadhaarResponseDataWithDefaults instantiates a new DigilockerEAadhaarResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerEAadhaarResponseDataWithDefaults() *DigilockerEAadhaarResponseData {
	this := DigilockerEAadhaarResponseData{}
	return &this
}

// GetAadhaarReferenceNumber returns the AadhaarReferenceNumber field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetAadhaarReferenceNumber() string {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarReferenceNumber
}

// GetAadhaarReferenceNumberOk returns a tuple with the AadhaarReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetAadhaarReferenceNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarReferenceNumber) {
    return nil, false
	}
	return o.AadhaarReferenceNumber, true
}

// HasAadhaarReferenceNumber returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasAadhaarReferenceNumber() bool {
	if o != nil && !isNil(o.AadhaarReferenceNumber) {
		return true
	}

	return false
}

// SetAadhaarReferenceNumber gets a reference to the given string and assigns it to the AadhaarReferenceNumber field.
func (o *DigilockerEAadhaarResponseData) SetAadhaarReferenceNumber(v string) {
	o.AadhaarReferenceNumber = &v
}

// GetAadhaarUid returns the AadhaarUid field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetAadhaarUid() string {
	if o == nil || isNil(o.AadhaarUid) {
		var ret string
		return ret
	}
	return *o.AadhaarUid
}

// GetAadhaarUidOk returns a tuple with the AadhaarUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetAadhaarUidOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarUid) {
    return nil, false
	}
	return o.AadhaarUid, true
}

// HasAadhaarUid returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasAadhaarUid() bool {
	if o != nil && !isNil(o.AadhaarUid) {
		return true
	}

	return false
}

// SetAadhaarUid gets a reference to the given string and assigns it to the AadhaarUid field.
func (o *DigilockerEAadhaarResponseData) SetAadhaarUid(v string) {
	o.AadhaarUid = &v
}

// GetProofOfIdentity returns the ProofOfIdentity field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetProofOfIdentity() DigilockerEAadhaarResponseDataProofOfIdentity {
	if o == nil || isNil(o.ProofOfIdentity) {
		var ret DigilockerEAadhaarResponseDataProofOfIdentity
		return ret
	}
	return *o.ProofOfIdentity
}

// GetProofOfIdentityOk returns a tuple with the ProofOfIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetProofOfIdentityOk() (*DigilockerEAadhaarResponseDataProofOfIdentity, bool) {
	if o == nil || isNil(o.ProofOfIdentity) {
    return nil, false
	}
	return o.ProofOfIdentity, true
}

// HasProofOfIdentity returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasProofOfIdentity() bool {
	if o != nil && !isNil(o.ProofOfIdentity) {
		return true
	}

	return false
}

// SetProofOfIdentity gets a reference to the given DigilockerEAadhaarResponseDataProofOfIdentity and assigns it to the ProofOfIdentity field.
func (o *DigilockerEAadhaarResponseData) SetProofOfIdentity(v DigilockerEAadhaarResponseDataProofOfIdentity) {
	o.ProofOfIdentity = &v
}

// GetProofOfAddress returns the ProofOfAddress field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetProofOfAddress() DigilockerEAadhaarResponseDataProofOfAddress {
	if o == nil || isNil(o.ProofOfAddress) {
		var ret DigilockerEAadhaarResponseDataProofOfAddress
		return ret
	}
	return *o.ProofOfAddress
}

// GetProofOfAddressOk returns a tuple with the ProofOfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetProofOfAddressOk() (*DigilockerEAadhaarResponseDataProofOfAddress, bool) {
	if o == nil || isNil(o.ProofOfAddress) {
    return nil, false
	}
	return o.ProofOfAddress, true
}

// HasProofOfAddress returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasProofOfAddress() bool {
	if o != nil && !isNil(o.ProofOfAddress) {
		return true
	}

	return false
}

// SetProofOfAddress gets a reference to the given DigilockerEAadhaarResponseDataProofOfAddress and assigns it to the ProofOfAddress field.
func (o *DigilockerEAadhaarResponseData) SetProofOfAddress(v DigilockerEAadhaarResponseDataProofOfAddress) {
	o.ProofOfAddress = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DigilockerEAadhaarResponseData) SetImage(v string) {
	o.Image = &v
}

// GetPdf returns the Pdf field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetPdf() string {
	if o == nil || isNil(o.Pdf) {
		var ret string
		return ret
	}
	return *o.Pdf
}

// GetPdfOk returns a tuple with the Pdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetPdfOk() (*string, bool) {
	if o == nil || isNil(o.Pdf) {
    return nil, false
	}
	return o.Pdf, true
}

// HasPdf returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasPdf() bool {
	if o != nil && !isNil(o.Pdf) {
		return true
	}

	return false
}

// SetPdf gets a reference to the given string and assigns it to the Pdf field.
func (o *DigilockerEAadhaarResponseData) SetPdf(v string) {
	o.Pdf = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *DigilockerEAadhaarResponseData) GetXml() string {
	if o == nil || isNil(o.Xml) {
		var ret string
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerEAadhaarResponseData) GetXmlOk() (*string, bool) {
	if o == nil || isNil(o.Xml) {
    return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *DigilockerEAadhaarResponseData) HasXml() bool {
	if o != nil && !isNil(o.Xml) {
		return true
	}

	return false
}

// SetXml gets a reference to the given string and assigns it to the Xml field.
func (o *DigilockerEAadhaarResponseData) SetXml(v string) {
	o.Xml = &v
}

func (o DigilockerEAadhaarResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarReferenceNumber) {
		toSerialize["aadhaarReferenceNumber"] = o.AadhaarReferenceNumber
	}
	if !isNil(o.AadhaarUid) {
		toSerialize["aadhaarUid"] = o.AadhaarUid
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
	if !isNil(o.Pdf) {
		toSerialize["pdf"] = o.Pdf
	}
	if !isNil(o.Xml) {
		toSerialize["xml"] = o.Xml
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerEAadhaarResponseData struct {
	value *DigilockerEAadhaarResponseData
	isSet bool
}

func (v NullableDigilockerEAadhaarResponseData) Get() *DigilockerEAadhaarResponseData {
	return v.value
}

func (v *NullableDigilockerEAadhaarResponseData) Set(val *DigilockerEAadhaarResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerEAadhaarResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerEAadhaarResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerEAadhaarResponseData(val *DigilockerEAadhaarResponseData) *NullableDigilockerEAadhaarResponseData {
	return &NullableDigilockerEAadhaarResponseData{value: val, isSet: true}
}

func (v NullableDigilockerEAadhaarResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerEAadhaarResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


