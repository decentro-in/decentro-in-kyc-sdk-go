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

// DigilockerFileDataResponseData struct for DigilockerFileDataResponseData
type DigilockerFileDataResponseData struct {
	DocumentIssuer *string `json:"documentIssuer,omitempty"`
	DocumentName *string `json:"documentName,omitempty"`
	DocumentType *string `json:"documentType,omitempty"`
	IdNumber *string `json:"idNumber,omitempty"`
	UserName *string `json:"userName,omitempty"`
	UserDateOfBirth *string `json:"userDateOfBirth,omitempty"`
	UserGender *string `json:"userGender,omitempty"`
	DocumentVerifiedOn *string `json:"documentVerifiedOn,omitempty"`
	DocumentStatus *string `json:"documentStatus,omitempty"`
}

// NewDigilockerFileDataResponseData instantiates a new DigilockerFileDataResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerFileDataResponseData() *DigilockerFileDataResponseData {
	this := DigilockerFileDataResponseData{}
	return &this
}

// NewDigilockerFileDataResponseDataWithDefaults instantiates a new DigilockerFileDataResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerFileDataResponseDataWithDefaults() *DigilockerFileDataResponseData {
	this := DigilockerFileDataResponseData{}
	return &this
}

// GetDocumentIssuer returns the DocumentIssuer field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetDocumentIssuer() string {
	if o == nil || isNil(o.DocumentIssuer) {
		var ret string
		return ret
	}
	return *o.DocumentIssuer
}

// GetDocumentIssuerOk returns a tuple with the DocumentIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetDocumentIssuerOk() (*string, bool) {
	if o == nil || isNil(o.DocumentIssuer) {
    return nil, false
	}
	return o.DocumentIssuer, true
}

// HasDocumentIssuer returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasDocumentIssuer() bool {
	if o != nil && !isNil(o.DocumentIssuer) {
		return true
	}

	return false
}

// SetDocumentIssuer gets a reference to the given string and assigns it to the DocumentIssuer field.
func (o *DigilockerFileDataResponseData) SetDocumentIssuer(v string) {
	o.DocumentIssuer = &v
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetDocumentName() string {
	if o == nil || isNil(o.DocumentName) {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetDocumentNameOk() (*string, bool) {
	if o == nil || isNil(o.DocumentName) {
    return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasDocumentName() bool {
	if o != nil && !isNil(o.DocumentName) {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
func (o *DigilockerFileDataResponseData) SetDocumentName(v string) {
	o.DocumentName = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetDocumentType() string {
	if o == nil || isNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetDocumentTypeOk() (*string, bool) {
	if o == nil || isNil(o.DocumentType) {
    return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasDocumentType() bool {
	if o != nil && !isNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *DigilockerFileDataResponseData) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetIdNumber() string {
	if o == nil || isNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetIdNumberOk() (*string, bool) {
	if o == nil || isNil(o.IdNumber) {
    return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasIdNumber() bool {
	if o != nil && !isNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *DigilockerFileDataResponseData) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
    return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DigilockerFileDataResponseData) SetUserName(v string) {
	o.UserName = &v
}

// GetUserDateOfBirth returns the UserDateOfBirth field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetUserDateOfBirth() string {
	if o == nil || isNil(o.UserDateOfBirth) {
		var ret string
		return ret
	}
	return *o.UserDateOfBirth
}

// GetUserDateOfBirthOk returns a tuple with the UserDateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetUserDateOfBirthOk() (*string, bool) {
	if o == nil || isNil(o.UserDateOfBirth) {
    return nil, false
	}
	return o.UserDateOfBirth, true
}

// HasUserDateOfBirth returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasUserDateOfBirth() bool {
	if o != nil && !isNil(o.UserDateOfBirth) {
		return true
	}

	return false
}

// SetUserDateOfBirth gets a reference to the given string and assigns it to the UserDateOfBirth field.
func (o *DigilockerFileDataResponseData) SetUserDateOfBirth(v string) {
	o.UserDateOfBirth = &v
}

// GetUserGender returns the UserGender field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetUserGender() string {
	if o == nil || isNil(o.UserGender) {
		var ret string
		return ret
	}
	return *o.UserGender
}

// GetUserGenderOk returns a tuple with the UserGender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetUserGenderOk() (*string, bool) {
	if o == nil || isNil(o.UserGender) {
    return nil, false
	}
	return o.UserGender, true
}

// HasUserGender returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasUserGender() bool {
	if o != nil && !isNil(o.UserGender) {
		return true
	}

	return false
}

// SetUserGender gets a reference to the given string and assigns it to the UserGender field.
func (o *DigilockerFileDataResponseData) SetUserGender(v string) {
	o.UserGender = &v
}

// GetDocumentVerifiedOn returns the DocumentVerifiedOn field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetDocumentVerifiedOn() string {
	if o == nil || isNil(o.DocumentVerifiedOn) {
		var ret string
		return ret
	}
	return *o.DocumentVerifiedOn
}

// GetDocumentVerifiedOnOk returns a tuple with the DocumentVerifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetDocumentVerifiedOnOk() (*string, bool) {
	if o == nil || isNil(o.DocumentVerifiedOn) {
    return nil, false
	}
	return o.DocumentVerifiedOn, true
}

// HasDocumentVerifiedOn returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasDocumentVerifiedOn() bool {
	if o != nil && !isNil(o.DocumentVerifiedOn) {
		return true
	}

	return false
}

// SetDocumentVerifiedOn gets a reference to the given string and assigns it to the DocumentVerifiedOn field.
func (o *DigilockerFileDataResponseData) SetDocumentVerifiedOn(v string) {
	o.DocumentVerifiedOn = &v
}

// GetDocumentStatus returns the DocumentStatus field value if set, zero value otherwise.
func (o *DigilockerFileDataResponseData) GetDocumentStatus() string {
	if o == nil || isNil(o.DocumentStatus) {
		var ret string
		return ret
	}
	return *o.DocumentStatus
}

// GetDocumentStatusOk returns a tuple with the DocumentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileDataResponseData) GetDocumentStatusOk() (*string, bool) {
	if o == nil || isNil(o.DocumentStatus) {
    return nil, false
	}
	return o.DocumentStatus, true
}

// HasDocumentStatus returns a boolean if a field has been set.
func (o *DigilockerFileDataResponseData) HasDocumentStatus() bool {
	if o != nil && !isNil(o.DocumentStatus) {
		return true
	}

	return false
}

// SetDocumentStatus gets a reference to the given string and assigns it to the DocumentStatus field.
func (o *DigilockerFileDataResponseData) SetDocumentStatus(v string) {
	o.DocumentStatus = &v
}

func (o DigilockerFileDataResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DocumentIssuer) {
		toSerialize["documentIssuer"] = o.DocumentIssuer
	}
	if !isNil(o.DocumentName) {
		toSerialize["documentName"] = o.DocumentName
	}
	if !isNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !isNil(o.IdNumber) {
		toSerialize["idNumber"] = o.IdNumber
	}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !isNil(o.UserDateOfBirth) {
		toSerialize["userDateOfBirth"] = o.UserDateOfBirth
	}
	if !isNil(o.UserGender) {
		toSerialize["userGender"] = o.UserGender
	}
	if !isNil(o.DocumentVerifiedOn) {
		toSerialize["documentVerifiedOn"] = o.DocumentVerifiedOn
	}
	if !isNil(o.DocumentStatus) {
		toSerialize["documentStatus"] = o.DocumentStatus
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerFileDataResponseData struct {
	value *DigilockerFileDataResponseData
	isSet bool
}

func (v NullableDigilockerFileDataResponseData) Get() *DigilockerFileDataResponseData {
	return v.value
}

func (v *NullableDigilockerFileDataResponseData) Set(val *DigilockerFileDataResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerFileDataResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerFileDataResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerFileDataResponseData(val *DigilockerFileDataResponseData) *NullableDigilockerFileDataResponseData {
	return &NullableDigilockerFileDataResponseData{value: val, isSet: true}
}

func (v NullableDigilockerFileDataResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerFileDataResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


