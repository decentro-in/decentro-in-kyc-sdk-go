# ValidateResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ResponseCode** | Pointer to **string** |  | [optional] 

## Methods

### NewValidateResponseError

`func NewValidateResponseError() *ValidateResponseError`

NewValidateResponseError instantiates a new ValidateResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateResponseErrorWithDefaults

`func NewValidateResponseErrorWithDefaults() *ValidateResponseError`

NewValidateResponseErrorWithDefaults instantiates a new ValidateResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidateResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidateResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidateResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidateResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResponseCode

`func (o *ValidateResponseError) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *ValidateResponseError) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *ValidateResponseError) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *ValidateResponseError) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


