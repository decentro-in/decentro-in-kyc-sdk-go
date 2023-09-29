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

// Validate400Response struct for Validate400Response
type Validate400Response struct {
	Status *string `json:"status,omitempty"`
	KycStatus *string `json:"kycStatus,omitempty"`
	Error *Validate400ResponseError `json:"error,omitempty"`
	RequestTimestamp *string `json:"requestTimestamp,omitempty"`
	ResponseTimestamp *string `json:"responseTimestamp,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
}

// NewValidate400Response instantiates a new Validate400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidate400Response() *Validate400Response {
	this := Validate400Response{}
	return &this
}

// NewValidate400ResponseWithDefaults instantiates a new Validate400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidate400ResponseWithDefaults() *Validate400Response {
	this := Validate400Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Validate400Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Validate400Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Validate400Response) SetStatus(v string) {
	o.Status = &v
}

// GetKycStatus returns the KycStatus field value if set, zero value otherwise.
func (o *Validate400Response) GetKycStatus() string {
	if o == nil || isNil(o.KycStatus) {
		var ret string
		return ret
	}
	return *o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetKycStatusOk() (*string, bool) {
	if o == nil || isNil(o.KycStatus) {
    return nil, false
	}
	return o.KycStatus, true
}

// HasKycStatus returns a boolean if a field has been set.
func (o *Validate400Response) HasKycStatus() bool {
	if o != nil && !isNil(o.KycStatus) {
		return true
	}

	return false
}

// SetKycStatus gets a reference to the given string and assigns it to the KycStatus field.
func (o *Validate400Response) SetKycStatus(v string) {
	o.KycStatus = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Validate400Response) GetError() Validate400ResponseError {
	if o == nil || isNil(o.Error) {
		var ret Validate400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetErrorOk() (*Validate400ResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Validate400Response) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Validate400ResponseError and assigns it to the Error field.
func (o *Validate400Response) SetError(v Validate400ResponseError) {
	o.Error = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *Validate400Response) GetRequestTimestamp() string {
	if o == nil || isNil(o.RequestTimestamp) {
		var ret string
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetRequestTimestampOk() (*string, bool) {
	if o == nil || isNil(o.RequestTimestamp) {
    return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *Validate400Response) HasRequestTimestamp() bool {
	if o != nil && !isNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given string and assigns it to the RequestTimestamp field.
func (o *Validate400Response) SetRequestTimestamp(v string) {
	o.RequestTimestamp = &v
}

// GetResponseTimestamp returns the ResponseTimestamp field value if set, zero value otherwise.
func (o *Validate400Response) GetResponseTimestamp() string {
	if o == nil || isNil(o.ResponseTimestamp) {
		var ret string
		return ret
	}
	return *o.ResponseTimestamp
}

// GetResponseTimestampOk returns a tuple with the ResponseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetResponseTimestampOk() (*string, bool) {
	if o == nil || isNil(o.ResponseTimestamp) {
    return nil, false
	}
	return o.ResponseTimestamp, true
}

// HasResponseTimestamp returns a boolean if a field has been set.
func (o *Validate400Response) HasResponseTimestamp() bool {
	if o != nil && !isNil(o.ResponseTimestamp) {
		return true
	}

	return false
}

// SetResponseTimestamp gets a reference to the given string and assigns it to the ResponseTimestamp field.
func (o *Validate400Response) SetResponseTimestamp(v string) {
	o.ResponseTimestamp = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *Validate400Response) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400Response) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *Validate400Response) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *Validate400Response) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

func (o Validate400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.KycStatus) {
		toSerialize["kycStatus"] = o.KycStatus
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
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

type NullableValidate400Response struct {
	value *Validate400Response
	isSet bool
}

func (v NullableValidate400Response) Get() *Validate400Response {
	return v.value
}

func (v *NullableValidate400Response) Set(val *Validate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableValidate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableValidate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidate400Response(val *Validate400Response) *NullableValidate400Response {
	return &NullableValidate400Response{value: val, isSet: true}
}

func (v NullableValidate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


