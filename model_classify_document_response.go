/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package decentroinkyc

import (
	"encoding/json"
)

// ClassifyDocumentResponse struct for ClassifyDocumentResponse
type ClassifyDocumentResponse struct {
	Data *ClassifyDocumentResponseData `json:"data,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewClassifyDocumentResponse instantiates a new ClassifyDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifyDocumentResponse() *ClassifyDocumentResponse {
	this := ClassifyDocumentResponse{}
	return &this
}

// NewClassifyDocumentResponseWithDefaults instantiates a new ClassifyDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifyDocumentResponseWithDefaults() *ClassifyDocumentResponse {
	this := ClassifyDocumentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ClassifyDocumentResponse) GetData() ClassifyDocumentResponseData {
	if o == nil || isNil(o.Data) {
		var ret ClassifyDocumentResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponse) GetDataOk() (*ClassifyDocumentResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ClassifyDocumentResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ClassifyDocumentResponseData and assigns it to the Data field.
func (o *ClassifyDocumentResponse) SetData(v ClassifyDocumentResponseData) {
	o.Data = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *ClassifyDocumentResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *ClassifyDocumentResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *ClassifyDocumentResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClassifyDocumentResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClassifyDocumentResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClassifyDocumentResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *ClassifyDocumentResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *ClassifyDocumentResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *ClassifyDocumentResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClassifyDocumentResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClassifyDocumentResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClassifyDocumentResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ClassifyDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableClassifyDocumentResponse struct {
	value *ClassifyDocumentResponse
	isSet bool
}

func (v NullableClassifyDocumentResponse) Get() *ClassifyDocumentResponse {
	return v.value
}

func (v *NullableClassifyDocumentResponse) Set(val *ClassifyDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifyDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifyDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifyDocumentResponse(val *ClassifyDocumentResponse) *NullableClassifyDocumentResponse {
	return &NullableClassifyDocumentResponse{value: val, isSet: true}
}

func (v NullableClassifyDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifyDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


