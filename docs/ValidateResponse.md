# ValidateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**KycResult** | Pointer to [**ValidateResponseKycResult**](ValidateResponseKycResult.md) |  | [optional] 
**ResponseCode** | Pointer to **string** |  | [optional] 
**RequestTimestamp** | Pointer to **string** |  | [optional] 
**ResponseTimestamp** | Pointer to **string** |  | [optional] 
**DecentroTxnId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**ValidateResponseError**](ValidateResponseError.md) |  | [optional] 

## Methods

### NewValidateResponse

`func NewValidateResponse() *ValidateResponse`

NewValidateResponse instantiates a new ValidateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateResponseWithDefaults

`func NewValidateResponseWithDefaults() *ValidateResponse`

NewValidateResponseWithDefaults instantiates a new ValidateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycStatus

`func (o *ValidateResponse) GetKycStatus() string`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *ValidateResponse) GetKycStatusOk() (*string, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *ValidateResponse) SetKycStatus(v string)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *ValidateResponse) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetStatus

`func (o *ValidateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ValidateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetKycResult

`func (o *ValidateResponse) GetKycResult() ValidateResponseKycResult`

GetKycResult returns the KycResult field if non-nil, zero value otherwise.

### GetKycResultOk

`func (o *ValidateResponse) GetKycResultOk() (*ValidateResponseKycResult, bool)`

GetKycResultOk returns a tuple with the KycResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycResult

`func (o *ValidateResponse) SetKycResult(v ValidateResponseKycResult)`

SetKycResult sets KycResult field to given value.

### HasKycResult

`func (o *ValidateResponse) HasKycResult() bool`

HasKycResult returns a boolean if a field has been set.

### GetResponseCode

`func (o *ValidateResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *ValidateResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *ValidateResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *ValidateResponse) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *ValidateResponse) GetRequestTimestamp() string`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *ValidateResponse) GetRequestTimestampOk() (*string, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *ValidateResponse) SetRequestTimestamp(v string)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *ValidateResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetResponseTimestamp

`func (o *ValidateResponse) GetResponseTimestamp() string`

GetResponseTimestamp returns the ResponseTimestamp field if non-nil, zero value otherwise.

### GetResponseTimestampOk

`func (o *ValidateResponse) GetResponseTimestampOk() (*string, bool)`

GetResponseTimestampOk returns a tuple with the ResponseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimestamp

`func (o *ValidateResponse) SetResponseTimestamp(v string)`

SetResponseTimestamp sets ResponseTimestamp field to given value.

### HasResponseTimestamp

`func (o *ValidateResponse) HasResponseTimestamp() bool`

HasResponseTimestamp returns a boolean if a field has been set.

### GetDecentroTxnId

`func (o *ValidateResponse) GetDecentroTxnId() string`

GetDecentroTxnId returns the DecentroTxnId field if non-nil, zero value otherwise.

### GetDecentroTxnIdOk

`func (o *ValidateResponse) GetDecentroTxnIdOk() (*string, bool)`

GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentroTxnId

`func (o *ValidateResponse) SetDecentroTxnId(v string)`

SetDecentroTxnId sets DecentroTxnId field to given value.

### HasDecentroTxnId

`func (o *ValidateResponse) HasDecentroTxnId() bool`

HasDecentroTxnId returns a boolean if a field has been set.

### GetError

`func (o *ValidateResponse) GetError() ValidateResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidateResponse) GetErrorOk() (*ValidateResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidateResponse) SetError(v ValidateResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ValidateResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


