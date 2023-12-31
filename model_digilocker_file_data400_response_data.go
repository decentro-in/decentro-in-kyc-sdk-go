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

// DigilockerFileData400ResponseData struct for DigilockerFileData400ResponseData
type DigilockerFileData400ResponseData struct {
	DocumentIssuer *string `json:"documentIssuer,omitempty"`
	DocumentName *string `json:"documentName,omitempty"`
	DocumentStatus *string `json:"documentStatus,omitempty"`
	DocumentType *string `json:"documentType,omitempty"`
	DocumentVerifiedOn *string `json:"documentVerifiedOn,omitempty"`
	IdNumber *string `json:"idNumber,omitempty"`
	UserDateOfBirth *string `json:"userDateOfBirth,omitempty"`
	UserGender *string `json:"userGender,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// NewDigilockerFileData400ResponseData instantiates a new DigilockerFileData400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerFileData400ResponseData() *DigilockerFileData400ResponseData {
	this := DigilockerFileData400ResponseData{}
	return &this
}

// NewDigilockerFileData400ResponseDataWithDefaults instantiates a new DigilockerFileData400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerFileData400ResponseDataWithDefaults() *DigilockerFileData400ResponseData {
	this := DigilockerFileData400ResponseData{}
	return &this
}

// GetDocumentIssuer returns the DocumentIssuer field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetDocumentIssuer() string {
	if o == nil || isNil(o.DocumentIssuer) {
		var ret string
		return ret
	}
	return *o.DocumentIssuer
}

// GetDocumentIssuerOk returns a tuple with the DocumentIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetDocumentIssuerOk() (*string, bool) {
	if o == nil || isNil(o.DocumentIssuer) {
    return nil, false
	}
	return o.DocumentIssuer, true
}

// HasDocumentIssuer returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasDocumentIssuer() bool {
	if o != nil && !isNil(o.DocumentIssuer) {
		return true
	}

	return false
}

// SetDocumentIssuer gets a reference to the given string and assigns it to the DocumentIssuer field.
func (o *DigilockerFileData400ResponseData) SetDocumentIssuer(v string) {
	o.DocumentIssuer = &v
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetDocumentName() string {
	if o == nil || isNil(o.DocumentName) {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetDocumentNameOk() (*string, bool) {
	if o == nil || isNil(o.DocumentName) {
    return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasDocumentName() bool {
	if o != nil && !isNil(o.DocumentName) {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
func (o *DigilockerFileData400ResponseData) SetDocumentName(v string) {
	o.DocumentName = &v
}

// GetDocumentStatus returns the DocumentStatus field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetDocumentStatus() string {
	if o == nil || isNil(o.DocumentStatus) {
		var ret string
		return ret
	}
	return *o.DocumentStatus
}

// GetDocumentStatusOk returns a tuple with the DocumentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetDocumentStatusOk() (*string, bool) {
	if o == nil || isNil(o.DocumentStatus) {
    return nil, false
	}
	return o.DocumentStatus, true
}

// HasDocumentStatus returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasDocumentStatus() bool {
	if o != nil && !isNil(o.DocumentStatus) {
		return true
	}

	return false
}

// SetDocumentStatus gets a reference to the given string and assigns it to the DocumentStatus field.
func (o *DigilockerFileData400ResponseData) SetDocumentStatus(v string) {
	o.DocumentStatus = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetDocumentType() string {
	if o == nil || isNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetDocumentTypeOk() (*string, bool) {
	if o == nil || isNil(o.DocumentType) {
    return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasDocumentType() bool {
	if o != nil && !isNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *DigilockerFileData400ResponseData) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetDocumentVerifiedOn returns the DocumentVerifiedOn field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetDocumentVerifiedOn() string {
	if o == nil || isNil(o.DocumentVerifiedOn) {
		var ret string
		return ret
	}
	return *o.DocumentVerifiedOn
}

// GetDocumentVerifiedOnOk returns a tuple with the DocumentVerifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetDocumentVerifiedOnOk() (*string, bool) {
	if o == nil || isNil(o.DocumentVerifiedOn) {
    return nil, false
	}
	return o.DocumentVerifiedOn, true
}

// HasDocumentVerifiedOn returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasDocumentVerifiedOn() bool {
	if o != nil && !isNil(o.DocumentVerifiedOn) {
		return true
	}

	return false
}

// SetDocumentVerifiedOn gets a reference to the given string and assigns it to the DocumentVerifiedOn field.
func (o *DigilockerFileData400ResponseData) SetDocumentVerifiedOn(v string) {
	o.DocumentVerifiedOn = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetIdNumber() string {
	if o == nil || isNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetIdNumberOk() (*string, bool) {
	if o == nil || isNil(o.IdNumber) {
    return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasIdNumber() bool {
	if o != nil && !isNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *DigilockerFileData400ResponseData) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetUserDateOfBirth returns the UserDateOfBirth field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetUserDateOfBirth() string {
	if o == nil || isNil(o.UserDateOfBirth) {
		var ret string
		return ret
	}
	return *o.UserDateOfBirth
}

// GetUserDateOfBirthOk returns a tuple with the UserDateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetUserDateOfBirthOk() (*string, bool) {
	if o == nil || isNil(o.UserDateOfBirth) {
    return nil, false
	}
	return o.UserDateOfBirth, true
}

// HasUserDateOfBirth returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasUserDateOfBirth() bool {
	if o != nil && !isNil(o.UserDateOfBirth) {
		return true
	}

	return false
}

// SetUserDateOfBirth gets a reference to the given string and assigns it to the UserDateOfBirth field.
func (o *DigilockerFileData400ResponseData) SetUserDateOfBirth(v string) {
	o.UserDateOfBirth = &v
}

// GetUserGender returns the UserGender field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetUserGender() string {
	if o == nil || isNil(o.UserGender) {
		var ret string
		return ret
	}
	return *o.UserGender
}

// GetUserGenderOk returns a tuple with the UserGender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetUserGenderOk() (*string, bool) {
	if o == nil || isNil(o.UserGender) {
    return nil, false
	}
	return o.UserGender, true
}

// HasUserGender returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasUserGender() bool {
	if o != nil && !isNil(o.UserGender) {
		return true
	}

	return false
}

// SetUserGender gets a reference to the given string and assigns it to the UserGender field.
func (o *DigilockerFileData400ResponseData) SetUserGender(v string) {
	o.UserGender = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DigilockerFileData400ResponseData) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerFileData400ResponseData) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
    return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DigilockerFileData400ResponseData) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DigilockerFileData400ResponseData) SetUserName(v string) {
	o.UserName = &v
}

func (o DigilockerFileData400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DocumentIssuer) {
		toSerialize["documentIssuer"] = o.DocumentIssuer
	}
	if !isNil(o.DocumentName) {
		toSerialize["documentName"] = o.DocumentName
	}
	if !isNil(o.DocumentStatus) {
		toSerialize["documentStatus"] = o.DocumentStatus
	}
	if !isNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !isNil(o.DocumentVerifiedOn) {
		toSerialize["documentVerifiedOn"] = o.DocumentVerifiedOn
	}
	if !isNil(o.IdNumber) {
		toSerialize["idNumber"] = o.IdNumber
	}
	if !isNil(o.UserDateOfBirth) {
		toSerialize["userDateOfBirth"] = o.UserDateOfBirth
	}
	if !isNil(o.UserGender) {
		toSerialize["userGender"] = o.UserGender
	}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerFileData400ResponseData struct {
	value *DigilockerFileData400ResponseData
	isSet bool
}

func (v NullableDigilockerFileData400ResponseData) Get() *DigilockerFileData400ResponseData {
	return v.value
}

func (v *NullableDigilockerFileData400ResponseData) Set(val *DigilockerFileData400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerFileData400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerFileData400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerFileData400ResponseData(val *DigilockerFileData400ResponseData) *NullableDigilockerFileData400ResponseData {
	return &NullableDigilockerFileData400ResponseData{value: val, isSet: true}
}

func (v NullableDigilockerFileData400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerFileData400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


