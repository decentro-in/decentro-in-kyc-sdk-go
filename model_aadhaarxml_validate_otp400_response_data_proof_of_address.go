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

// AadhaarxmlValidateOtp400ResponseDataProofOfAddress struct for AadhaarxmlValidateOtp400ResponseDataProofOfAddress
type AadhaarxmlValidateOtp400ResponseDataProofOfAddress struct {
	CareOf *string `json:"careOf,omitempty"`
	Country *string `json:"country,omitempty"`
	District *string `json:"district,omitempty"`
	House *string `json:"house,omitempty"`
	Landmark *string `json:"landmark,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Pincode *string `json:"pincode,omitempty"`
	PostOffice *string `json:"postOffice,omitempty"`
	State *string `json:"state,omitempty"`
	Street *string `json:"street,omitempty"`
	SubDistrict *string `json:"subDistrict,omitempty"`
	Vtc *string `json:"vtc,omitempty"`
}

// NewAadhaarxmlValidateOtp400ResponseDataProofOfAddress instantiates a new AadhaarxmlValidateOtp400ResponseDataProofOfAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtp400ResponseDataProofOfAddress() *AadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	this := AadhaarxmlValidateOtp400ResponseDataProofOfAddress{}
	return &this
}

// NewAadhaarxmlValidateOtp400ResponseDataProofOfAddressWithDefaults instantiates a new AadhaarxmlValidateOtp400ResponseDataProofOfAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtp400ResponseDataProofOfAddressWithDefaults() *AadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	this := AadhaarxmlValidateOtp400ResponseDataProofOfAddress{}
	return &this
}

// GetCareOf returns the CareOf field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetCareOf() string {
	if o == nil || isNil(o.CareOf) {
		var ret string
		return ret
	}
	return *o.CareOf
}

// GetCareOfOk returns a tuple with the CareOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetCareOfOk() (*string, bool) {
	if o == nil || isNil(o.CareOf) {
    return nil, false
	}
	return o.CareOf, true
}

// HasCareOf returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasCareOf() bool {
	if o != nil && !isNil(o.CareOf) {
		return true
	}

	return false
}

// SetCareOf gets a reference to the given string and assigns it to the CareOf field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetCareOf(v string) {
	o.CareOf = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetCountry(v string) {
	o.Country = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetDistrict() string {
	if o == nil || isNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetDistrictOk() (*string, bool) {
	if o == nil || isNil(o.District) {
    return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasDistrict() bool {
	if o != nil && !isNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetDistrict(v string) {
	o.District = &v
}

// GetHouse returns the House field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetHouse() string {
	if o == nil || isNil(o.House) {
		var ret string
		return ret
	}
	return *o.House
}

// GetHouseOk returns a tuple with the House field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetHouseOk() (*string, bool) {
	if o == nil || isNil(o.House) {
    return nil, false
	}
	return o.House, true
}

// HasHouse returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasHouse() bool {
	if o != nil && !isNil(o.House) {
		return true
	}

	return false
}

// SetHouse gets a reference to the given string and assigns it to the House field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetHouse(v string) {
	o.House = &v
}

// GetLandmark returns the Landmark field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetLandmark() string {
	if o == nil || isNil(o.Landmark) {
		var ret string
		return ret
	}
	return *o.Landmark
}

// GetLandmarkOk returns a tuple with the Landmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetLandmarkOk() (*string, bool) {
	if o == nil || isNil(o.Landmark) {
    return nil, false
	}
	return o.Landmark, true
}

// HasLandmark returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasLandmark() bool {
	if o != nil && !isNil(o.Landmark) {
		return true
	}

	return false
}

// SetLandmark gets a reference to the given string and assigns it to the Landmark field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetLandmark(v string) {
	o.Landmark = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetLocality() string {
	if o == nil || isNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetLocalityOk() (*string, bool) {
	if o == nil || isNil(o.Locality) {
    return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasLocality() bool {
	if o != nil && !isNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetPincode returns the Pincode field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetPincode() string {
	if o == nil || isNil(o.Pincode) {
		var ret string
		return ret
	}
	return *o.Pincode
}

// GetPincodeOk returns a tuple with the Pincode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetPincodeOk() (*string, bool) {
	if o == nil || isNil(o.Pincode) {
    return nil, false
	}
	return o.Pincode, true
}

// HasPincode returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasPincode() bool {
	if o != nil && !isNil(o.Pincode) {
		return true
	}

	return false
}

// SetPincode gets a reference to the given string and assigns it to the Pincode field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetPincode(v string) {
	o.Pincode = &v
}

// GetPostOffice returns the PostOffice field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetPostOffice() string {
	if o == nil || isNil(o.PostOffice) {
		var ret string
		return ret
	}
	return *o.PostOffice
}

// GetPostOfficeOk returns a tuple with the PostOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetPostOfficeOk() (*string, bool) {
	if o == nil || isNil(o.PostOffice) {
    return nil, false
	}
	return o.PostOffice, true
}

// HasPostOffice returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasPostOffice() bool {
	if o != nil && !isNil(o.PostOffice) {
		return true
	}

	return false
}

// SetPostOffice gets a reference to the given string and assigns it to the PostOffice field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetPostOffice(v string) {
	o.PostOffice = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetState(v string) {
	o.State = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetStreet() string {
	if o == nil || isNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetStreetOk() (*string, bool) {
	if o == nil || isNil(o.Street) {
    return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasStreet() bool {
	if o != nil && !isNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetStreet(v string) {
	o.Street = &v
}

// GetSubDistrict returns the SubDistrict field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetSubDistrict() string {
	if o == nil || isNil(o.SubDistrict) {
		var ret string
		return ret
	}
	return *o.SubDistrict
}

// GetSubDistrictOk returns a tuple with the SubDistrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetSubDistrictOk() (*string, bool) {
	if o == nil || isNil(o.SubDistrict) {
    return nil, false
	}
	return o.SubDistrict, true
}

// HasSubDistrict returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasSubDistrict() bool {
	if o != nil && !isNil(o.SubDistrict) {
		return true
	}

	return false
}

// SetSubDistrict gets a reference to the given string and assigns it to the SubDistrict field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetSubDistrict(v string) {
	o.SubDistrict = &v
}

// GetVtc returns the Vtc field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetVtc() string {
	if o == nil || isNil(o.Vtc) {
		var ret string
		return ret
	}
	return *o.Vtc
}

// GetVtcOk returns a tuple with the Vtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) GetVtcOk() (*string, bool) {
	if o == nil || isNil(o.Vtc) {
    return nil, false
	}
	return o.Vtc, true
}

// HasVtc returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) HasVtc() bool {
	if o != nil && !isNil(o.Vtc) {
		return true
	}

	return false
}

// SetVtc gets a reference to the given string and assigns it to the Vtc field.
func (o *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) SetVtc(v string) {
	o.Vtc = &v
}

func (o AadhaarxmlValidateOtp400ResponseDataProofOfAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CareOf) {
		toSerialize["careOf"] = o.CareOf
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !isNil(o.House) {
		toSerialize["house"] = o.House
	}
	if !isNil(o.Landmark) {
		toSerialize["landmark"] = o.Landmark
	}
	if !isNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !isNil(o.Pincode) {
		toSerialize["pincode"] = o.Pincode
	}
	if !isNil(o.PostOffice) {
		toSerialize["postOffice"] = o.PostOffice
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !isNil(o.SubDistrict) {
		toSerialize["subDistrict"] = o.SubDistrict
	}
	if !isNil(o.Vtc) {
		toSerialize["vtc"] = o.Vtc
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress struct {
	value *AadhaarxmlValidateOtp400ResponseDataProofOfAddress
	isSet bool
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) Get() *AadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) Set(val *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress(val *AadhaarxmlValidateOtp400ResponseDataProofOfAddress) *NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress {
	return &NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataProofOfAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


