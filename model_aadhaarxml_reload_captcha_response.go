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

// AadhaarxmlReloadCaptchaResponse struct for AadhaarxmlReloadCaptchaResponse
type AadhaarxmlReloadCaptchaResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data *AadhaarxmlReloadCaptchaResponseData `json:"data,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewAadhaarxmlReloadCaptchaResponse instantiates a new AadhaarxmlReloadCaptchaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlReloadCaptchaResponse() *AadhaarxmlReloadCaptchaResponse {
	this := AadhaarxmlReloadCaptchaResponse{}
	return &this
}

// NewAadhaarxmlReloadCaptchaResponseWithDefaults instantiates a new AadhaarxmlReloadCaptchaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlReloadCaptchaResponseWithDefaults() *AadhaarxmlReloadCaptchaResponse {
	this := AadhaarxmlReloadCaptchaResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *AadhaarxmlReloadCaptchaResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AadhaarxmlReloadCaptchaResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *AadhaarxmlReloadCaptchaResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AadhaarxmlReloadCaptchaResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetData() AadhaarxmlReloadCaptchaResponseData {
	if o == nil || isNil(o.Data) {
		var ret AadhaarxmlReloadCaptchaResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetDataOk() (*AadhaarxmlReloadCaptchaResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AadhaarxmlReloadCaptchaResponseData and assigns it to the Data field.
func (o *AadhaarxmlReloadCaptchaResponse) SetData(v AadhaarxmlReloadCaptchaResponseData) {
	o.Data = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *AadhaarxmlReloadCaptchaResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o AadhaarxmlReloadCaptchaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.ResponseKey) {
		toSerialize["responseKey"] = o.ResponseKey
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlReloadCaptchaResponse struct {
	value *AadhaarxmlReloadCaptchaResponse
	isSet bool
}

func (v NullableAadhaarxmlReloadCaptchaResponse) Get() *AadhaarxmlReloadCaptchaResponse {
	return v.value
}

func (v *NullableAadhaarxmlReloadCaptchaResponse) Set(val *AadhaarxmlReloadCaptchaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlReloadCaptchaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlReloadCaptchaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlReloadCaptchaResponse(val *AadhaarxmlReloadCaptchaResponse) *NullableAadhaarxmlReloadCaptchaResponse {
	return &NullableAadhaarxmlReloadCaptchaResponse{value: val, isSet: true}
}

func (v NullableAadhaarxmlReloadCaptchaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlReloadCaptchaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


