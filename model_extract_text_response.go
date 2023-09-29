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

// ExtractTextResponse struct for ExtractTextResponse
type ExtractTextResponse struct {
	OcrStatus *string `json:"ocrStatus,omitempty"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	OcrResult *ExtractTextResponseOcrResult `json:"ocrResult,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	RequestTimestamp *string `json:"requestTimestamp,omitempty"`
	ResponseTimestamp *string `json:"responseTimestamp,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
}

// NewExtractTextResponse instantiates a new ExtractTextResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractTextResponse() *ExtractTextResponse {
	this := ExtractTextResponse{}
	return &this
}

// NewExtractTextResponseWithDefaults instantiates a new ExtractTextResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractTextResponseWithDefaults() *ExtractTextResponse {
	this := ExtractTextResponse{}
	return &this
}

// GetOcrStatus returns the OcrStatus field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetOcrStatus() string {
	if o == nil || isNil(o.OcrStatus) {
		var ret string
		return ret
	}
	return *o.OcrStatus
}

// GetOcrStatusOk returns a tuple with the OcrStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetOcrStatusOk() (*string, bool) {
	if o == nil || isNil(o.OcrStatus) {
    return nil, false
	}
	return o.OcrStatus, true
}

// HasOcrStatus returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasOcrStatus() bool {
	if o != nil && !isNil(o.OcrStatus) {
		return true
	}

	return false
}

// SetOcrStatus gets a reference to the given string and assigns it to the OcrStatus field.
func (o *ExtractTextResponse) SetOcrStatus(v string) {
	o.OcrStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExtractTextResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ExtractTextResponse) SetMessage(v string) {
	o.Message = &v
}

// GetOcrResult returns the OcrResult field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetOcrResult() ExtractTextResponseOcrResult {
	if o == nil || isNil(o.OcrResult) {
		var ret ExtractTextResponseOcrResult
		return ret
	}
	return *o.OcrResult
}

// GetOcrResultOk returns a tuple with the OcrResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetOcrResultOk() (*ExtractTextResponseOcrResult, bool) {
	if o == nil || isNil(o.OcrResult) {
    return nil, false
	}
	return o.OcrResult, true
}

// HasOcrResult returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasOcrResult() bool {
	if o != nil && !isNil(o.OcrResult) {
		return true
	}

	return false
}

// SetOcrResult gets a reference to the given ExtractTextResponseOcrResult and assigns it to the OcrResult field.
func (o *ExtractTextResponse) SetOcrResult(v ExtractTextResponseOcrResult) {
	o.OcrResult = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *ExtractTextResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetRequestTimestamp() string {
	if o == nil || isNil(o.RequestTimestamp) {
		var ret string
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetRequestTimestampOk() (*string, bool) {
	if o == nil || isNil(o.RequestTimestamp) {
    return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasRequestTimestamp() bool {
	if o != nil && !isNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given string and assigns it to the RequestTimestamp field.
func (o *ExtractTextResponse) SetRequestTimestamp(v string) {
	o.RequestTimestamp = &v
}

// GetResponseTimestamp returns the ResponseTimestamp field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetResponseTimestamp() string {
	if o == nil || isNil(o.ResponseTimestamp) {
		var ret string
		return ret
	}
	return *o.ResponseTimestamp
}

// GetResponseTimestampOk returns a tuple with the ResponseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetResponseTimestampOk() (*string, bool) {
	if o == nil || isNil(o.ResponseTimestamp) {
    return nil, false
	}
	return o.ResponseTimestamp, true
}

// HasResponseTimestamp returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasResponseTimestamp() bool {
	if o != nil && !isNil(o.ResponseTimestamp) {
		return true
	}

	return false
}

// SetResponseTimestamp gets a reference to the given string and assigns it to the ResponseTimestamp field.
func (o *ExtractTextResponse) SetResponseTimestamp(v string) {
	o.ResponseTimestamp = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *ExtractTextResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *ExtractTextResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *ExtractTextResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

func (o ExtractTextResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OcrStatus) {
		toSerialize["ocrStatus"] = o.OcrStatus
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.OcrResult) {
		toSerialize["ocrResult"] = o.OcrResult
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.RequestTimestamp) {
		toSerialize["requestTimestamp"] = o.RequestTimestamp
	}
	if !isNil(o.ResponseTimestamp) {
		toSerialize["responseTimestamp"] = o.ResponseTimestamp
	}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	return json.Marshal(toSerialize)
}

type NullableExtractTextResponse struct {
	value *ExtractTextResponse
	isSet bool
}

func (v NullableExtractTextResponse) Get() *ExtractTextResponse {
	return v.value
}

func (v *NullableExtractTextResponse) Set(val *ExtractTextResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractTextResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractTextResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractTextResponse(val *ExtractTextResponse) *NullableExtractTextResponse {
	return &NullableExtractTextResponse{value: val, isSet: true}
}

func (v NullableExtractTextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractTextResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


