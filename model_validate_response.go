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

// ValidateResponse struct for ValidateResponse
type ValidateResponse struct {
	KycStatus *string `json:"kycStatus,omitempty"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	KycResult *ValidateResponseKycResult `json:"kycResult,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	RequestTimestamp *string `json:"requestTimestamp,omitempty"`
	ResponseTimestamp *string `json:"responseTimestamp,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Error *ValidateResponseError `json:"error,omitempty"`
}

// NewValidateResponse instantiates a new ValidateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponse() *ValidateResponse {
	this := ValidateResponse{}
	return &this
}

// NewValidateResponseWithDefaults instantiates a new ValidateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseWithDefaults() *ValidateResponse {
	this := ValidateResponse{}
	return &this
}

// GetKycStatus returns the KycStatus field value if set, zero value otherwise.
func (o *ValidateResponse) GetKycStatus() string {
	if o == nil || isNil(o.KycStatus) {
		var ret string
		return ret
	}
	return *o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetKycStatusOk() (*string, bool) {
	if o == nil || isNil(o.KycStatus) {
    return nil, false
	}
	return o.KycStatus, true
}

// HasKycStatus returns a boolean if a field has been set.
func (o *ValidateResponse) HasKycStatus() bool {
	if o != nil && !isNil(o.KycStatus) {
		return true
	}

	return false
}

// SetKycStatus gets a reference to the given string and assigns it to the KycStatus field.
func (o *ValidateResponse) SetKycStatus(v string) {
	o.KycStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ValidateResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ValidateResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ValidateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidateResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidateResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidateResponse) SetMessage(v string) {
	o.Message = &v
}

// GetKycResult returns the KycResult field value if set, zero value otherwise.
func (o *ValidateResponse) GetKycResult() ValidateResponseKycResult {
	if o == nil || isNil(o.KycResult) {
		var ret ValidateResponseKycResult
		return ret
	}
	return *o.KycResult
}

// GetKycResultOk returns a tuple with the KycResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetKycResultOk() (*ValidateResponseKycResult, bool) {
	if o == nil || isNil(o.KycResult) {
    return nil, false
	}
	return o.KycResult, true
}

// HasKycResult returns a boolean if a field has been set.
func (o *ValidateResponse) HasKycResult() bool {
	if o != nil && !isNil(o.KycResult) {
		return true
	}

	return false
}

// SetKycResult gets a reference to the given ValidateResponseKycResult and assigns it to the KycResult field.
func (o *ValidateResponse) SetKycResult(v ValidateResponseKycResult) {
	o.KycResult = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *ValidateResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *ValidateResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *ValidateResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *ValidateResponse) GetRequestTimestamp() string {
	if o == nil || isNil(o.RequestTimestamp) {
		var ret string
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetRequestTimestampOk() (*string, bool) {
	if o == nil || isNil(o.RequestTimestamp) {
    return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *ValidateResponse) HasRequestTimestamp() bool {
	if o != nil && !isNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given string and assigns it to the RequestTimestamp field.
func (o *ValidateResponse) SetRequestTimestamp(v string) {
	o.RequestTimestamp = &v
}

// GetResponseTimestamp returns the ResponseTimestamp field value if set, zero value otherwise.
func (o *ValidateResponse) GetResponseTimestamp() string {
	if o == nil || isNil(o.ResponseTimestamp) {
		var ret string
		return ret
	}
	return *o.ResponseTimestamp
}

// GetResponseTimestampOk returns a tuple with the ResponseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetResponseTimestampOk() (*string, bool) {
	if o == nil || isNil(o.ResponseTimestamp) {
    return nil, false
	}
	return o.ResponseTimestamp, true
}

// HasResponseTimestamp returns a boolean if a field has been set.
func (o *ValidateResponse) HasResponseTimestamp() bool {
	if o != nil && !isNil(o.ResponseTimestamp) {
		return true
	}

	return false
}

// SetResponseTimestamp gets a reference to the given string and assigns it to the ResponseTimestamp field.
func (o *ValidateResponse) SetResponseTimestamp(v string) {
	o.ResponseTimestamp = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *ValidateResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *ValidateResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *ValidateResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ValidateResponse) GetError() ValidateResponseError {
	if o == nil || isNil(o.Error) {
		var ret ValidateResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponse) GetErrorOk() (*ValidateResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ValidateResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ValidateResponseError and assigns it to the Error field.
func (o *ValidateResponse) SetError(v ValidateResponseError) {
	o.Error = &v
}

func (o ValidateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.KycStatus) {
		toSerialize["kycStatus"] = o.KycStatus
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.KycResult) {
		toSerialize["kycResult"] = o.KycResult
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
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponse struct {
	value *ValidateResponse
	isSet bool
}

func (v NullableValidateResponse) Get() *ValidateResponse {
	return v.value
}

func (v *NullableValidateResponse) Set(val *ValidateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponse(val *ValidateResponse) *NullableValidateResponse {
	return &NullableValidateResponse{value: val, isSet: true}
}

func (v NullableValidateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


