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

// ValidateResponseKycResultForeignCompanyMasterData struct for ValidateResponseKycResultForeignCompanyMasterData
type ValidateResponseKycResultForeignCompanyMasterData struct {
	EmailId *string `json:"emailId,omitempty"`
	ForeignCompanyWithShareCapital *string `json:"foreignCompanyWithShareCapital,omitempty"`
	RegisteredAddress *string `json:"registeredAddress,omitempty"`
	TypeOfOffice *string `json:"typeOfOffice,omitempty"`
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	CountryOfIncorporation *string `json:"countryOfIncorporation,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	CompanyStatus *string `json:"companyStatus,omitempty"`
	Details *string `json:"details,omitempty"`
	Fcrn *string `json:"fcrn ,omitempty"`
	DescriptionOfMainDivision *string `json:"descriptionOfMainDivision,omitempty"`
	MainDivisionOfBusinessActivityToBeCarriedOutInIndia *string `json:"mainDivisionOfBusinessActivityToBeCarriedOutInIndia,omitempty"`
}

// NewValidateResponseKycResultForeignCompanyMasterData instantiates a new ValidateResponseKycResultForeignCompanyMasterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultForeignCompanyMasterData() *ValidateResponseKycResultForeignCompanyMasterData {
	this := ValidateResponseKycResultForeignCompanyMasterData{}
	return &this
}

// NewValidateResponseKycResultForeignCompanyMasterDataWithDefaults instantiates a new ValidateResponseKycResultForeignCompanyMasterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultForeignCompanyMasterDataWithDefaults() *ValidateResponseKycResultForeignCompanyMasterData {
	this := ValidateResponseKycResultForeignCompanyMasterData{}
	return &this
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetEmailId() string {
	if o == nil || isNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetEmailIdOk() (*string, bool) {
	if o == nil || isNil(o.EmailId) {
    return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasEmailId() bool {
	if o != nil && !isNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetEmailId(v string) {
	o.EmailId = &v
}

// GetForeignCompanyWithShareCapital returns the ForeignCompanyWithShareCapital field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetForeignCompanyWithShareCapital() string {
	if o == nil || isNil(o.ForeignCompanyWithShareCapital) {
		var ret string
		return ret
	}
	return *o.ForeignCompanyWithShareCapital
}

// GetForeignCompanyWithShareCapitalOk returns a tuple with the ForeignCompanyWithShareCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetForeignCompanyWithShareCapitalOk() (*string, bool) {
	if o == nil || isNil(o.ForeignCompanyWithShareCapital) {
    return nil, false
	}
	return o.ForeignCompanyWithShareCapital, true
}

// HasForeignCompanyWithShareCapital returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasForeignCompanyWithShareCapital() bool {
	if o != nil && !isNil(o.ForeignCompanyWithShareCapital) {
		return true
	}

	return false
}

// SetForeignCompanyWithShareCapital gets a reference to the given string and assigns it to the ForeignCompanyWithShareCapital field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetForeignCompanyWithShareCapital(v string) {
	o.ForeignCompanyWithShareCapital = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetRegisteredAddress() string {
	if o == nil || isNil(o.RegisteredAddress) {
		var ret string
		return ret
	}
	return *o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetRegisteredAddressOk() (*string, bool) {
	if o == nil || isNil(o.RegisteredAddress) {
    return nil, false
	}
	return o.RegisteredAddress, true
}

// HasRegisteredAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasRegisteredAddress() bool {
	if o != nil && !isNil(o.RegisteredAddress) {
		return true
	}

	return false
}

// SetRegisteredAddress gets a reference to the given string and assigns it to the RegisteredAddress field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetRegisteredAddress(v string) {
	o.RegisteredAddress = &v
}

// GetTypeOfOffice returns the TypeOfOffice field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetTypeOfOffice() string {
	if o == nil || isNil(o.TypeOfOffice) {
		var ret string
		return ret
	}
	return *o.TypeOfOffice
}

// GetTypeOfOfficeOk returns a tuple with the TypeOfOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetTypeOfOfficeOk() (*string, bool) {
	if o == nil || isNil(o.TypeOfOffice) {
    return nil, false
	}
	return o.TypeOfOffice, true
}

// HasTypeOfOffice returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasTypeOfOffice() bool {
	if o != nil && !isNil(o.TypeOfOffice) {
		return true
	}

	return false
}

// SetTypeOfOffice gets a reference to the given string and assigns it to the TypeOfOffice field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetTypeOfOffice(v string) {
	o.TypeOfOffice = &v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDateOfIncorporation() string {
	if o == nil || isNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || isNil(o.DateOfIncorporation) {
    return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasDateOfIncorporation() bool {
	if o != nil && !isNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetCountryOfIncorporation returns the CountryOfIncorporation field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCountryOfIncorporation() string {
	if o == nil || isNil(o.CountryOfIncorporation) {
		var ret string
		return ret
	}
	return *o.CountryOfIncorporation
}

// GetCountryOfIncorporationOk returns a tuple with the CountryOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCountryOfIncorporationOk() (*string, bool) {
	if o == nil || isNil(o.CountryOfIncorporation) {
    return nil, false
	}
	return o.CountryOfIncorporation, true
}

// HasCountryOfIncorporation returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasCountryOfIncorporation() bool {
	if o != nil && !isNil(o.CountryOfIncorporation) {
		return true
	}

	return false
}

// SetCountryOfIncorporation gets a reference to the given string and assigns it to the CountryOfIncorporation field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetCountryOfIncorporation(v string) {
	o.CountryOfIncorporation = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCompanyName() string {
	if o == nil || isNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCompanyNameOk() (*string, bool) {
	if o == nil || isNil(o.CompanyName) {
    return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasCompanyName() bool {
	if o != nil && !isNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCompanyStatus returns the CompanyStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCompanyStatus() string {
	if o == nil || isNil(o.CompanyStatus) {
		var ret string
		return ret
	}
	return *o.CompanyStatus
}

// GetCompanyStatusOk returns a tuple with the CompanyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetCompanyStatusOk() (*string, bool) {
	if o == nil || isNil(o.CompanyStatus) {
    return nil, false
	}
	return o.CompanyStatus, true
}

// HasCompanyStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasCompanyStatus() bool {
	if o != nil && !isNil(o.CompanyStatus) {
		return true
	}

	return false
}

// SetCompanyStatus gets a reference to the given string and assigns it to the CompanyStatus field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetCompanyStatus(v string) {
	o.CompanyStatus = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetDetails(v string) {
	o.Details = &v
}

// GetFcrn returns the Fcrn field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetFcrn() string {
	if o == nil || isNil(o.Fcrn) {
		var ret string
		return ret
	}
	return *o.Fcrn
}

// GetFcrnOk returns a tuple with the Fcrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetFcrnOk() (*string, bool) {
	if o == nil || isNil(o.Fcrn) {
    return nil, false
	}
	return o.Fcrn, true
}

// HasFcrn returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasFcrn() bool {
	if o != nil && !isNil(o.Fcrn) {
		return true
	}

	return false
}

// SetFcrn gets a reference to the given string and assigns it to the Fcrn field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetFcrn(v string) {
	o.Fcrn = &v
}

// GetDescriptionOfMainDivision returns the DescriptionOfMainDivision field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDescriptionOfMainDivision() string {
	if o == nil || isNil(o.DescriptionOfMainDivision) {
		var ret string
		return ret
	}
	return *o.DescriptionOfMainDivision
}

// GetDescriptionOfMainDivisionOk returns a tuple with the DescriptionOfMainDivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetDescriptionOfMainDivisionOk() (*string, bool) {
	if o == nil || isNil(o.DescriptionOfMainDivision) {
    return nil, false
	}
	return o.DescriptionOfMainDivision, true
}

// HasDescriptionOfMainDivision returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasDescriptionOfMainDivision() bool {
	if o != nil && !isNil(o.DescriptionOfMainDivision) {
		return true
	}

	return false
}

// SetDescriptionOfMainDivision gets a reference to the given string and assigns it to the DescriptionOfMainDivision field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetDescriptionOfMainDivision(v string) {
	o.DescriptionOfMainDivision = &v
}

// GetMainDivisionOfBusinessActivityToBeCarriedOutInIndia returns the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field value if set, zero value otherwise.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetMainDivisionOfBusinessActivityToBeCarriedOutInIndia() string {
	if o == nil || isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		var ret string
		return ret
	}
	return *o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia
}

// GetMainDivisionOfBusinessActivityToBeCarriedOutInIndiaOk returns a tuple with the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) GetMainDivisionOfBusinessActivityToBeCarriedOutInIndiaOk() (*string, bool) {
	if o == nil || isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
    return nil, false
	}
	return o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia, true
}

// HasMainDivisionOfBusinessActivityToBeCarriedOutInIndia returns a boolean if a field has been set.
func (o *ValidateResponseKycResultForeignCompanyMasterData) HasMainDivisionOfBusinessActivityToBeCarriedOutInIndia() bool {
	if o != nil && !isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		return true
	}

	return false
}

// SetMainDivisionOfBusinessActivityToBeCarriedOutInIndia gets a reference to the given string and assigns it to the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field.
func (o *ValidateResponseKycResultForeignCompanyMasterData) SetMainDivisionOfBusinessActivityToBeCarriedOutInIndia(v string) {
	o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia = &v
}

func (o ValidateResponseKycResultForeignCompanyMasterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !isNil(o.ForeignCompanyWithShareCapital) {
		toSerialize["foreignCompanyWithShareCapital"] = o.ForeignCompanyWithShareCapital
	}
	if !isNil(o.RegisteredAddress) {
		toSerialize["registeredAddress"] = o.RegisteredAddress
	}
	if !isNil(o.TypeOfOffice) {
		toSerialize["typeOfOffice"] = o.TypeOfOffice
	}
	if !isNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !isNil(o.CountryOfIncorporation) {
		toSerialize["countryOfIncorporation"] = o.CountryOfIncorporation
	}
	if !isNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !isNil(o.CompanyStatus) {
		toSerialize["companyStatus"] = o.CompanyStatus
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.Fcrn) {
		toSerialize["fcrn "] = o.Fcrn
	}
	if !isNil(o.DescriptionOfMainDivision) {
		toSerialize["descriptionOfMainDivision"] = o.DescriptionOfMainDivision
	}
	if !isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		toSerialize["mainDivisionOfBusinessActivityToBeCarriedOutInIndia"] = o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultForeignCompanyMasterData struct {
	value *ValidateResponseKycResultForeignCompanyMasterData
	isSet bool
}

func (v NullableValidateResponseKycResultForeignCompanyMasterData) Get() *ValidateResponseKycResultForeignCompanyMasterData {
	return v.value
}

func (v *NullableValidateResponseKycResultForeignCompanyMasterData) Set(val *ValidateResponseKycResultForeignCompanyMasterData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultForeignCompanyMasterData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultForeignCompanyMasterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultForeignCompanyMasterData(val *ValidateResponseKycResultForeignCompanyMasterData) *NullableValidateResponseKycResultForeignCompanyMasterData {
	return &NullableValidateResponseKycResultForeignCompanyMasterData{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultForeignCompanyMasterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultForeignCompanyMasterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


