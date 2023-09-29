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

// ValidateResponseKycResult struct for ValidateResponseKycResult
type ValidateResponseKycResult struct {
	IdNumber *string `json:"idNumber,omitempty"`
	IdStatus *string `json:"idStatus,omitempty"`
	Name *string `json:"name,omitempty"`
	LicenseType *string `json:"licenseType,omitempty"`
	EntityName *string `json:"entityName,omitempty"`
	Status *string `json:"status,omitempty"`
	PremissesAddress *ValidateResponseKycResultPremissesAddress `json:"premissesAddress,omitempty"`
	Products []ValidateResponseKycResultProductsInner `json:"products,omitempty"`
	Address *ValidateResponseKycResultAddress `json:"address,omitempty"`
	UamNumber *string `json:"uamNumber,omitempty"`
	EnterpriseName *string `json:"enterpriseName,omitempty"`
	MajorActivity *string `json:"majorActivity,omitempty"`
	SocialCategory *string `json:"socialCategory,omitempty"`
	EnterpriseType *string `json:"enterpriseType,omitempty"`
	DateOfCommencement *string `json:"dateOfCommencement,omitempty"`
	District *string `json:"district,omitempty"`
	State *string `json:"state,omitempty"`
	AppliedDate *string `json:"appliedDate,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	Nic2Digit *string `json:"nic_2Digit,omitempty"`
	Nic4Digit *string `json:"nic_4Digit,omitempty"`
	Nic5Digit *string `json:"nic_5Digit,omitempty"`
	PanStatus *string `json:"panStatus,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	IdHolderTitle *string `json:"idHolderTitle,omitempty"`
	IdLastUpdated *string `json:"idLastUpdated,omitempty"`
	AadhaarSeedingStatus *string `json:"aadhaarSeedingStatus,omitempty"`
	Addresses []ValidateResponseKycResultAddressesInner `json:"addresses,omitempty"`
	AllClassOfVehicle []ValidateResponseKycResultAllClassOfVehicleInner `json:"allClassOfVehicle,omitempty"`
	DrivingLicenseNumber *string `json:"drivingLicenseNumber,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	EndorseDate *string `json:"endorseDate,omitempty"`
	EndorseNumber *string `json:"endorseNumber,omitempty"`
	FatherOrHusbandName *string `json:"fatherOrHusbandName,omitempty"`
	ValidFrom *string `json:"validFrom,omitempty"`
	ValidTo *string `json:"validTo,omitempty"`
	EpicNo *string `json:"epicNo,omitempty"`
	NameInVernacular *string `json:"nameInVernacular,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Age *float32 `json:"age,omitempty"`
	RelativeName *string `json:"relativeName,omitempty"`
	RelativeNameInVernacular *string `json:"relativeNameInVernacular,omitempty"`
	RelativeRelationType *string `json:"relativeRelationType,omitempty"`
	HouseNumber *string `json:"houseNumber,omitempty"`
	PartOrLocationInConstituency *string `json:"partOrLocationInConstituency,omitempty"`
	PartNumberOrLocationNumberInConstituency *string `json:"partNumberOrLocationNumberInConstituency,omitempty"`
	ParliamentaryConstituency *string `json:"parliamentaryConstituency,omitempty"`
	AssemblyConstituency *string `json:"assemblyConstituency,omitempty"`
	SectionOfConstituencyPart *string `json:"sectionOfConstituencyPart,omitempty"`
	CardSerialNumberInPollingList *string `json:"cardSerialNumberInPollingList,omitempty"`
	LastUpdateDate *string `json:"lastUpdateDate,omitempty"`
	PollingBoothDetails *ValidateResponseKycResultPollingBoothDetails `json:"pollingBoothDetails,omitempty"`
	EmailId *string `json:"emailId,omitempty"`
	MobileNumber *string `json:"mobileNumber,omitempty"`
	StateCode *string `json:"stateCode,omitempty"`
	PollingBoothCoordinates *string `json:"pollingBoothCoordinates,omitempty"`
	PollingBoothAddress *string `json:"pollingBoothAddress,omitempty"`
	PollingBoothNumber *string `json:"pollingBoothNumber,omitempty"`
	Id *string `json:"id,omitempty"`
	BlacklistStatus *string `json:"blacklistStatus,omitempty"`
	RegistrationDate *string `json:"registrationDate,omitempty"`
	RegistrationLocation *string `json:"registrationLocation,omitempty"`
	Class *string `json:"class,omitempty"`
	Maker *string `json:"maker,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	ChassisNumber *string `json:"chassisNumber,omitempty"`
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	EngineNumber *string `json:"engineNumber,omitempty"`
	FuelType *string `json:"fuelType,omitempty"`
	FitUpto *string `json:"fitUpto,omitempty"`
	InsuranceUpto *string `json:"insuranceUpto,omitempty"`
	TaxUpto *string `json:"taxUpto,omitempty"`
	InsuranceDetails *string `json:"insuranceDetails,omitempty"`
	InsuranceValidity *string `json:"insuranceValidity,omitempty"`
	PermitType *string `json:"permitType,omitempty"`
	PermitValidUpto *string `json:"permitValidUpto,omitempty"`
	PollutionControlValidity *string `json:"pollutionControlValidity,omitempty"`
	PollutionNorms *string `json:"pollutionNorms,omitempty"`
	LicenseAddress *string `json:"licenseAddress,omitempty"`
	RegistrationAddress *string `json:"registrationAddress,omitempty"`
	OwnerFatherName *string `json:"ownerFatherName,omitempty"`
	OwnerPresentAddress *string `json:"ownerPresentAddress,omitempty"`
	BodyType *string `json:"bodyType,omitempty"`
	Category *string `json:"category,omitempty"`
	Color *string `json:"color,omitempty"`
	EngineCubicCapacity *string `json:"engineCubicCapacity,omitempty"`
	NumberCylinders *string `json:"numberCylinders,omitempty"`
	UnladenWeight *string `json:"unladenWeight,omitempty"`
	GrossWeight *string `json:"grossWeight,omitempty"`
	Wheelbase *string `json:"wheelbase,omitempty"`
	ManufacturedMonthYear *string `json:"manufacturedMonthYear,omitempty"`
	MakerDescription *string `json:"makerDescription,omitempty"`
	NocDetails *string `json:"nocDetails,omitempty"`
	NormsDescription *string `json:"normsDescription,omitempty"`
	Financier *string `json:"financier,omitempty"`
	PermitIssueDate *string `json:"permitIssueDate,omitempty"`
	PermitNumber *string `json:"permitNumber,omitempty"`
	PermitValidFrom *string `json:"permitValidFrom,omitempty"`
	SeatingCapacity *string `json:"seatingCapacity,omitempty"`
	SleepingCapacity *string `json:"sleepingCapacity,omitempty"`
	StandingCapacity *string `json:"standingCapacity,omitempty"`
	StatusAsOn *string `json:"statusAsOn,omitempty"`
	PrimaryBusinessContact *ValidateResponseKycResultPrimaryBusinessContact `json:"primaryBusinessContact,omitempty"`
	StateJurisdiction *string `json:"stateJurisdiction,omitempty"`
	StateJurisdictionCode *string `json:"stateJurisdictionCode,omitempty"`
	TaxpayerType *string `json:"taxpayerType,omitempty"`
	ConstitutionOfBusiness *string `json:"constitutionOfBusiness,omitempty"`
	GstnStatus *string `json:"gstnStatus,omitempty"`
	TradeName *string `json:"tradeName,omitempty"`
	Gstin *string `json:"gstin,omitempty"`
	AdditionalPlacesOfBusinessInState []string `json:"additionalPlacesOfBusinessInState,omitempty"`
	LegalName *string `json:"legalName,omitempty"`
	NatureOfBusiness []string `json:"natureOfBusiness,omitempty"`
	CentralJurisdiction *string `json:"centralJurisdiction,omitempty"`
	CentralJurisdictionCode *string `json:"centralJurisdictionCode,omitempty"`
	Pan *string `json:"pan,omitempty"`
	AuthorizedSignatories *string `json:"authorizedSignatories,omitempty"`
	ComplianceRating *string `json:"complianceRating,omitempty"`
	Cxdt *string `json:"cxdt,omitempty"`
	BusinessDetails []ValidateResponseKycResultBusinessDetailsInner `json:"businessDetails,omitempty"`
	AnnualAggregateTurnover *string `json:"annualAggregateTurnover,omitempty"`
	MandatoryEInvoicing *string `json:"mandatoryEInvoicing,omitempty"`
	GrossTotalIncome *string `json:"grossTotalIncome,omitempty"`
	GrossTotalIncomeFinancialYear *string `json:"grossTotalIncomeFinancialYear,omitempty"`
	IsFieldVisitConducted *string `json:"isFieldVisitConducted,omitempty"`
	FilingStatus []ValidateResponseKycResultFilingStatusInner `json:"filingStatus,omitempty"`
	Directors *ValidateResponseKycResultDirectors `json:"directors,omitempty"`
	CompanyMasterData *ValidateResponseKycResultCompanyMasterData `json:"companyMasterData,omitempty"`
	Charges []string `json:"charges,omitempty"`
	LlpData []string `json:"llpData,omitempty"`
	CompanyData []string `json:"companyData,omitempty"`
	DirectorData *ValidateResponseKycResultDirectorData `json:"directorData,omitempty"`
	LlpMasterData *ValidateResponseKycResultLlpMasterData `json:"llpMasterData,omitempty"`
	ForeignCompanyMasterData *ValidateResponseKycResultForeignCompanyMasterData `json:"foreignCompanyMasterData,omitempty"`
}

// NewValidateResponseKycResult instantiates a new ValidateResponseKycResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResult() *ValidateResponseKycResult {
	this := ValidateResponseKycResult{}
	return &this
}

// NewValidateResponseKycResultWithDefaults instantiates a new ValidateResponseKycResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultWithDefaults() *ValidateResponseKycResult {
	this := ValidateResponseKycResult{}
	return &this
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetIdNumber() string {
	if o == nil || isNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIdNumberOk() (*string, bool) {
	if o == nil || isNil(o.IdNumber) {
    return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasIdNumber() bool {
	if o != nil && !isNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *ValidateResponseKycResult) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetIdStatus returns the IdStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetIdStatus() string {
	if o == nil || isNil(o.IdStatus) {
		var ret string
		return ret
	}
	return *o.IdStatus
}

// GetIdStatusOk returns a tuple with the IdStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIdStatusOk() (*string, bool) {
	if o == nil || isNil(o.IdStatus) {
    return nil, false
	}
	return o.IdStatus, true
}

// HasIdStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasIdStatus() bool {
	if o != nil && !isNil(o.IdStatus) {
		return true
	}

	return false
}

// SetIdStatus gets a reference to the given string and assigns it to the IdStatus field.
func (o *ValidateResponseKycResult) SetIdStatus(v string) {
	o.IdStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ValidateResponseKycResult) SetName(v string) {
	o.Name = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLicenseType() string {
	if o == nil || isNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLicenseTypeOk() (*string, bool) {
	if o == nil || isNil(o.LicenseType) {
    return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLicenseType() bool {
	if o != nil && !isNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *ValidateResponseKycResult) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEntityName() string {
	if o == nil || isNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEntityNameOk() (*string, bool) {
	if o == nil || isNil(o.EntityName) {
    return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEntityName() bool {
	if o != nil && !isNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *ValidateResponseKycResult) SetEntityName(v string) {
	o.EntityName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ValidateResponseKycResult) SetStatus(v string) {
	o.Status = &v
}

// GetPremissesAddress returns the PremissesAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPremissesAddress() ValidateResponseKycResultPremissesAddress {
	if o == nil || isNil(o.PremissesAddress) {
		var ret ValidateResponseKycResultPremissesAddress
		return ret
	}
	return *o.PremissesAddress
}

// GetPremissesAddressOk returns a tuple with the PremissesAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPremissesAddressOk() (*ValidateResponseKycResultPremissesAddress, bool) {
	if o == nil || isNil(o.PremissesAddress) {
    return nil, false
	}
	return o.PremissesAddress, true
}

// HasPremissesAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPremissesAddress() bool {
	if o != nil && !isNil(o.PremissesAddress) {
		return true
	}

	return false
}

// SetPremissesAddress gets a reference to the given ValidateResponseKycResultPremissesAddress and assigns it to the PremissesAddress field.
func (o *ValidateResponseKycResult) SetPremissesAddress(v ValidateResponseKycResultPremissesAddress) {
	o.PremissesAddress = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetProducts() []ValidateResponseKycResultProductsInner {
	if o == nil || isNil(o.Products) {
		var ret []ValidateResponseKycResultProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetProductsOk() ([]ValidateResponseKycResultProductsInner, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ValidateResponseKycResultProductsInner and assigns it to the Products field.
func (o *ValidateResponseKycResult) SetProducts(v []ValidateResponseKycResultProductsInner) {
	o.Products = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAddress() ValidateResponseKycResultAddress {
	if o == nil || isNil(o.Address) {
		var ret ValidateResponseKycResultAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAddressOk() (*ValidateResponseKycResultAddress, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ValidateResponseKycResultAddress and assigns it to the Address field.
func (o *ValidateResponseKycResult) SetAddress(v ValidateResponseKycResultAddress) {
	o.Address = &v
}

// GetUamNumber returns the UamNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetUamNumber() string {
	if o == nil || isNil(o.UamNumber) {
		var ret string
		return ret
	}
	return *o.UamNumber
}

// GetUamNumberOk returns a tuple with the UamNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetUamNumberOk() (*string, bool) {
	if o == nil || isNil(o.UamNumber) {
    return nil, false
	}
	return o.UamNumber, true
}

// HasUamNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasUamNumber() bool {
	if o != nil && !isNil(o.UamNumber) {
		return true
	}

	return false
}

// SetUamNumber gets a reference to the given string and assigns it to the UamNumber field.
func (o *ValidateResponseKycResult) SetUamNumber(v string) {
	o.UamNumber = &v
}

// GetEnterpriseName returns the EnterpriseName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEnterpriseName() string {
	if o == nil || isNil(o.EnterpriseName) {
		var ret string
		return ret
	}
	return *o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEnterpriseNameOk() (*string, bool) {
	if o == nil || isNil(o.EnterpriseName) {
    return nil, false
	}
	return o.EnterpriseName, true
}

// HasEnterpriseName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEnterpriseName() bool {
	if o != nil && !isNil(o.EnterpriseName) {
		return true
	}

	return false
}

// SetEnterpriseName gets a reference to the given string and assigns it to the EnterpriseName field.
func (o *ValidateResponseKycResult) SetEnterpriseName(v string) {
	o.EnterpriseName = &v
}

// GetMajorActivity returns the MajorActivity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetMajorActivity() string {
	if o == nil || isNil(o.MajorActivity) {
		var ret string
		return ret
	}
	return *o.MajorActivity
}

// GetMajorActivityOk returns a tuple with the MajorActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetMajorActivityOk() (*string, bool) {
	if o == nil || isNil(o.MajorActivity) {
    return nil, false
	}
	return o.MajorActivity, true
}

// HasMajorActivity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasMajorActivity() bool {
	if o != nil && !isNil(o.MajorActivity) {
		return true
	}

	return false
}

// SetMajorActivity gets a reference to the given string and assigns it to the MajorActivity field.
func (o *ValidateResponseKycResult) SetMajorActivity(v string) {
	o.MajorActivity = &v
}

// GetSocialCategory returns the SocialCategory field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetSocialCategory() string {
	if o == nil || isNil(o.SocialCategory) {
		var ret string
		return ret
	}
	return *o.SocialCategory
}

// GetSocialCategoryOk returns a tuple with the SocialCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetSocialCategoryOk() (*string, bool) {
	if o == nil || isNil(o.SocialCategory) {
    return nil, false
	}
	return o.SocialCategory, true
}

// HasSocialCategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasSocialCategory() bool {
	if o != nil && !isNil(o.SocialCategory) {
		return true
	}

	return false
}

// SetSocialCategory gets a reference to the given string and assigns it to the SocialCategory field.
func (o *ValidateResponseKycResult) SetSocialCategory(v string) {
	o.SocialCategory = &v
}

// GetEnterpriseType returns the EnterpriseType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEnterpriseType() string {
	if o == nil || isNil(o.EnterpriseType) {
		var ret string
		return ret
	}
	return *o.EnterpriseType
}

// GetEnterpriseTypeOk returns a tuple with the EnterpriseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEnterpriseTypeOk() (*string, bool) {
	if o == nil || isNil(o.EnterpriseType) {
    return nil, false
	}
	return o.EnterpriseType, true
}

// HasEnterpriseType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEnterpriseType() bool {
	if o != nil && !isNil(o.EnterpriseType) {
		return true
	}

	return false
}

// SetEnterpriseType gets a reference to the given string and assigns it to the EnterpriseType field.
func (o *ValidateResponseKycResult) SetEnterpriseType(v string) {
	o.EnterpriseType = &v
}

// GetDateOfCommencement returns the DateOfCommencement field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDateOfCommencement() string {
	if o == nil || isNil(o.DateOfCommencement) {
		var ret string
		return ret
	}
	return *o.DateOfCommencement
}

// GetDateOfCommencementOk returns a tuple with the DateOfCommencement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDateOfCommencementOk() (*string, bool) {
	if o == nil || isNil(o.DateOfCommencement) {
    return nil, false
	}
	return o.DateOfCommencement, true
}

// HasDateOfCommencement returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDateOfCommencement() bool {
	if o != nil && !isNil(o.DateOfCommencement) {
		return true
	}

	return false
}

// SetDateOfCommencement gets a reference to the given string and assigns it to the DateOfCommencement field.
func (o *ValidateResponseKycResult) SetDateOfCommencement(v string) {
	o.DateOfCommencement = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDistrict() string {
	if o == nil || isNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDistrictOk() (*string, bool) {
	if o == nil || isNil(o.District) {
    return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDistrict() bool {
	if o != nil && !isNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *ValidateResponseKycResult) SetDistrict(v string) {
	o.District = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ValidateResponseKycResult) SetState(v string) {
	o.State = &v
}

// GetAppliedDate returns the AppliedDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAppliedDate() string {
	if o == nil || isNil(o.AppliedDate) {
		var ret string
		return ret
	}
	return *o.AppliedDate
}

// GetAppliedDateOk returns a tuple with the AppliedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAppliedDateOk() (*string, bool) {
	if o == nil || isNil(o.AppliedDate) {
    return nil, false
	}
	return o.AppliedDate, true
}

// HasAppliedDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAppliedDate() bool {
	if o != nil && !isNil(o.AppliedDate) {
		return true
	}

	return false
}

// SetAppliedDate gets a reference to the given string and assigns it to the AppliedDate field.
func (o *ValidateResponseKycResult) SetAppliedDate(v string) {
	o.AppliedDate = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetModifiedDate() string {
	if o == nil || isNil(o.ModifiedDate) {
		var ret string
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetModifiedDateOk() (*string, bool) {
	if o == nil || isNil(o.ModifiedDate) {
    return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasModifiedDate() bool {
	if o != nil && !isNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given string and assigns it to the ModifiedDate field.
func (o *ValidateResponseKycResult) SetModifiedDate(v string) {
	o.ModifiedDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetExpiryDate() string {
	if o == nil || isNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetExpiryDateOk() (*string, bool) {
	if o == nil || isNil(o.ExpiryDate) {
    return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasExpiryDate() bool {
	if o != nil && !isNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ValidateResponseKycResult) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetNic2Digit returns the Nic2Digit field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNic2Digit() string {
	if o == nil || isNil(o.Nic2Digit) {
		var ret string
		return ret
	}
	return *o.Nic2Digit
}

// GetNic2DigitOk returns a tuple with the Nic2Digit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNic2DigitOk() (*string, bool) {
	if o == nil || isNil(o.Nic2Digit) {
    return nil, false
	}
	return o.Nic2Digit, true
}

// HasNic2Digit returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNic2Digit() bool {
	if o != nil && !isNil(o.Nic2Digit) {
		return true
	}

	return false
}

// SetNic2Digit gets a reference to the given string and assigns it to the Nic2Digit field.
func (o *ValidateResponseKycResult) SetNic2Digit(v string) {
	o.Nic2Digit = &v
}

// GetNic4Digit returns the Nic4Digit field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNic4Digit() string {
	if o == nil || isNil(o.Nic4Digit) {
		var ret string
		return ret
	}
	return *o.Nic4Digit
}

// GetNic4DigitOk returns a tuple with the Nic4Digit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNic4DigitOk() (*string, bool) {
	if o == nil || isNil(o.Nic4Digit) {
    return nil, false
	}
	return o.Nic4Digit, true
}

// HasNic4Digit returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNic4Digit() bool {
	if o != nil && !isNil(o.Nic4Digit) {
		return true
	}

	return false
}

// SetNic4Digit gets a reference to the given string and assigns it to the Nic4Digit field.
func (o *ValidateResponseKycResult) SetNic4Digit(v string) {
	o.Nic4Digit = &v
}

// GetNic5Digit returns the Nic5Digit field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNic5Digit() string {
	if o == nil || isNil(o.Nic5Digit) {
		var ret string
		return ret
	}
	return *o.Nic5Digit
}

// GetNic5DigitOk returns a tuple with the Nic5Digit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNic5DigitOk() (*string, bool) {
	if o == nil || isNil(o.Nic5Digit) {
    return nil, false
	}
	return o.Nic5Digit, true
}

// HasNic5Digit returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNic5Digit() bool {
	if o != nil && !isNil(o.Nic5Digit) {
		return true
	}

	return false
}

// SetNic5Digit gets a reference to the given string and assigns it to the Nic5Digit field.
func (o *ValidateResponseKycResult) SetNic5Digit(v string) {
	o.Nic5Digit = &v
}

// GetPanStatus returns the PanStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPanStatus() string {
	if o == nil || isNil(o.PanStatus) {
		var ret string
		return ret
	}
	return *o.PanStatus
}

// GetPanStatusOk returns a tuple with the PanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPanStatusOk() (*string, bool) {
	if o == nil || isNil(o.PanStatus) {
    return nil, false
	}
	return o.PanStatus, true
}

// HasPanStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPanStatus() bool {
	if o != nil && !isNil(o.PanStatus) {
		return true
	}

	return false
}

// SetPanStatus gets a reference to the given string and assigns it to the PanStatus field.
func (o *ValidateResponseKycResult) SetPanStatus(v string) {
	o.PanStatus = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
    return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ValidateResponseKycResult) SetLastName(v string) {
	o.LastName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
    return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ValidateResponseKycResult) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFullName() string {
	if o == nil || isNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFullNameOk() (*string, bool) {
	if o == nil || isNil(o.FullName) {
    return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFullName() bool {
	if o != nil && !isNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ValidateResponseKycResult) SetFullName(v string) {
	o.FullName = &v
}

// GetIdHolderTitle returns the IdHolderTitle field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetIdHolderTitle() string {
	if o == nil || isNil(o.IdHolderTitle) {
		var ret string
		return ret
	}
	return *o.IdHolderTitle
}

// GetIdHolderTitleOk returns a tuple with the IdHolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIdHolderTitleOk() (*string, bool) {
	if o == nil || isNil(o.IdHolderTitle) {
    return nil, false
	}
	return o.IdHolderTitle, true
}

// HasIdHolderTitle returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasIdHolderTitle() bool {
	if o != nil && !isNil(o.IdHolderTitle) {
		return true
	}

	return false
}

// SetIdHolderTitle gets a reference to the given string and assigns it to the IdHolderTitle field.
func (o *ValidateResponseKycResult) SetIdHolderTitle(v string) {
	o.IdHolderTitle = &v
}

// GetIdLastUpdated returns the IdLastUpdated field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetIdLastUpdated() string {
	if o == nil || isNil(o.IdLastUpdated) {
		var ret string
		return ret
	}
	return *o.IdLastUpdated
}

// GetIdLastUpdatedOk returns a tuple with the IdLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIdLastUpdatedOk() (*string, bool) {
	if o == nil || isNil(o.IdLastUpdated) {
    return nil, false
	}
	return o.IdLastUpdated, true
}

// HasIdLastUpdated returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasIdLastUpdated() bool {
	if o != nil && !isNil(o.IdLastUpdated) {
		return true
	}

	return false
}

// SetIdLastUpdated gets a reference to the given string and assigns it to the IdLastUpdated field.
func (o *ValidateResponseKycResult) SetIdLastUpdated(v string) {
	o.IdLastUpdated = &v
}

// GetAadhaarSeedingStatus returns the AadhaarSeedingStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAadhaarSeedingStatus() string {
	if o == nil || isNil(o.AadhaarSeedingStatus) {
		var ret string
		return ret
	}
	return *o.AadhaarSeedingStatus
}

// GetAadhaarSeedingStatusOk returns a tuple with the AadhaarSeedingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAadhaarSeedingStatusOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarSeedingStatus) {
    return nil, false
	}
	return o.AadhaarSeedingStatus, true
}

// HasAadhaarSeedingStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAadhaarSeedingStatus() bool {
	if o != nil && !isNil(o.AadhaarSeedingStatus) {
		return true
	}

	return false
}

// SetAadhaarSeedingStatus gets a reference to the given string and assigns it to the AadhaarSeedingStatus field.
func (o *ValidateResponseKycResult) SetAadhaarSeedingStatus(v string) {
	o.AadhaarSeedingStatus = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAddresses() []ValidateResponseKycResultAddressesInner {
	if o == nil || isNil(o.Addresses) {
		var ret []ValidateResponseKycResultAddressesInner
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAddressesOk() ([]ValidateResponseKycResultAddressesInner, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []ValidateResponseKycResultAddressesInner and assigns it to the Addresses field.
func (o *ValidateResponseKycResult) SetAddresses(v []ValidateResponseKycResultAddressesInner) {
	o.Addresses = v
}

// GetAllClassOfVehicle returns the AllClassOfVehicle field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAllClassOfVehicle() []ValidateResponseKycResultAllClassOfVehicleInner {
	if o == nil || isNil(o.AllClassOfVehicle) {
		var ret []ValidateResponseKycResultAllClassOfVehicleInner
		return ret
	}
	return o.AllClassOfVehicle
}

// GetAllClassOfVehicleOk returns a tuple with the AllClassOfVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAllClassOfVehicleOk() ([]ValidateResponseKycResultAllClassOfVehicleInner, bool) {
	if o == nil || isNil(o.AllClassOfVehicle) {
    return nil, false
	}
	return o.AllClassOfVehicle, true
}

// HasAllClassOfVehicle returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAllClassOfVehicle() bool {
	if o != nil && !isNil(o.AllClassOfVehicle) {
		return true
	}

	return false
}

// SetAllClassOfVehicle gets a reference to the given []ValidateResponseKycResultAllClassOfVehicleInner and assigns it to the AllClassOfVehicle field.
func (o *ValidateResponseKycResult) SetAllClassOfVehicle(v []ValidateResponseKycResultAllClassOfVehicleInner) {
	o.AllClassOfVehicle = v
}

// GetDrivingLicenseNumber returns the DrivingLicenseNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDrivingLicenseNumber() string {
	if o == nil || isNil(o.DrivingLicenseNumber) {
		var ret string
		return ret
	}
	return *o.DrivingLicenseNumber
}

// GetDrivingLicenseNumberOk returns a tuple with the DrivingLicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDrivingLicenseNumberOk() (*string, bool) {
	if o == nil || isNil(o.DrivingLicenseNumber) {
    return nil, false
	}
	return o.DrivingLicenseNumber, true
}

// HasDrivingLicenseNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDrivingLicenseNumber() bool {
	if o != nil && !isNil(o.DrivingLicenseNumber) {
		return true
	}

	return false
}

// SetDrivingLicenseNumber gets a reference to the given string and assigns it to the DrivingLicenseNumber field.
func (o *ValidateResponseKycResult) SetDrivingLicenseNumber(v string) {
	o.DrivingLicenseNumber = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDateOfBirth() string {
	if o == nil || isNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDateOfBirthOk() (*string, bool) {
	if o == nil || isNil(o.DateOfBirth) {
    return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDateOfBirth() bool {
	if o != nil && !isNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *ValidateResponseKycResult) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetEndorseDate returns the EndorseDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEndorseDate() string {
	if o == nil || isNil(o.EndorseDate) {
		var ret string
		return ret
	}
	return *o.EndorseDate
}

// GetEndorseDateOk returns a tuple with the EndorseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEndorseDateOk() (*string, bool) {
	if o == nil || isNil(o.EndorseDate) {
    return nil, false
	}
	return o.EndorseDate, true
}

// HasEndorseDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEndorseDate() bool {
	if o != nil && !isNil(o.EndorseDate) {
		return true
	}

	return false
}

// SetEndorseDate gets a reference to the given string and assigns it to the EndorseDate field.
func (o *ValidateResponseKycResult) SetEndorseDate(v string) {
	o.EndorseDate = &v
}

// GetEndorseNumber returns the EndorseNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEndorseNumber() string {
	if o == nil || isNil(o.EndorseNumber) {
		var ret string
		return ret
	}
	return *o.EndorseNumber
}

// GetEndorseNumberOk returns a tuple with the EndorseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEndorseNumberOk() (*string, bool) {
	if o == nil || isNil(o.EndorseNumber) {
    return nil, false
	}
	return o.EndorseNumber, true
}

// HasEndorseNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEndorseNumber() bool {
	if o != nil && !isNil(o.EndorseNumber) {
		return true
	}

	return false
}

// SetEndorseNumber gets a reference to the given string and assigns it to the EndorseNumber field.
func (o *ValidateResponseKycResult) SetEndorseNumber(v string) {
	o.EndorseNumber = &v
}

// GetFatherOrHusbandName returns the FatherOrHusbandName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFatherOrHusbandName() string {
	if o == nil || isNil(o.FatherOrHusbandName) {
		var ret string
		return ret
	}
	return *o.FatherOrHusbandName
}

// GetFatherOrHusbandNameOk returns a tuple with the FatherOrHusbandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFatherOrHusbandNameOk() (*string, bool) {
	if o == nil || isNil(o.FatherOrHusbandName) {
    return nil, false
	}
	return o.FatherOrHusbandName, true
}

// HasFatherOrHusbandName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFatherOrHusbandName() bool {
	if o != nil && !isNil(o.FatherOrHusbandName) {
		return true
	}

	return false
}

// SetFatherOrHusbandName gets a reference to the given string and assigns it to the FatherOrHusbandName field.
func (o *ValidateResponseKycResult) SetFatherOrHusbandName(v string) {
	o.FatherOrHusbandName = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetValidFrom() string {
	if o == nil || isNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetValidFromOk() (*string, bool) {
	if o == nil || isNil(o.ValidFrom) {
    return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasValidFrom() bool {
	if o != nil && !isNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ValidateResponseKycResult) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetValidTo() string {
	if o == nil || isNil(o.ValidTo) {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetValidToOk() (*string, bool) {
	if o == nil || isNil(o.ValidTo) {
    return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasValidTo() bool {
	if o != nil && !isNil(o.ValidTo) {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *ValidateResponseKycResult) SetValidTo(v string) {
	o.ValidTo = &v
}

// GetEpicNo returns the EpicNo field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEpicNo() string {
	if o == nil || isNil(o.EpicNo) {
		var ret string
		return ret
	}
	return *o.EpicNo
}

// GetEpicNoOk returns a tuple with the EpicNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEpicNoOk() (*string, bool) {
	if o == nil || isNil(o.EpicNo) {
    return nil, false
	}
	return o.EpicNo, true
}

// HasEpicNo returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEpicNo() bool {
	if o != nil && !isNil(o.EpicNo) {
		return true
	}

	return false
}

// SetEpicNo gets a reference to the given string and assigns it to the EpicNo field.
func (o *ValidateResponseKycResult) SetEpicNo(v string) {
	o.EpicNo = &v
}

// GetNameInVernacular returns the NameInVernacular field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNameInVernacular() string {
	if o == nil || isNil(o.NameInVernacular) {
		var ret string
		return ret
	}
	return *o.NameInVernacular
}

// GetNameInVernacularOk returns a tuple with the NameInVernacular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNameInVernacularOk() (*string, bool) {
	if o == nil || isNil(o.NameInVernacular) {
    return nil, false
	}
	return o.NameInVernacular, true
}

// HasNameInVernacular returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNameInVernacular() bool {
	if o != nil && !isNil(o.NameInVernacular) {
		return true
	}

	return false
}

// SetNameInVernacular gets a reference to the given string and assigns it to the NameInVernacular field.
func (o *ValidateResponseKycResult) SetNameInVernacular(v string) {
	o.NameInVernacular = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *ValidateResponseKycResult) SetGender(v string) {
	o.Gender = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAge() float32 {
	if o == nil || isNil(o.Age) {
		var ret float32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAgeOk() (*float32, bool) {
	if o == nil || isNil(o.Age) {
    return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAge() bool {
	if o != nil && !isNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given float32 and assigns it to the Age field.
func (o *ValidateResponseKycResult) SetAge(v float32) {
	o.Age = &v
}

// GetRelativeName returns the RelativeName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRelativeName() string {
	if o == nil || isNil(o.RelativeName) {
		var ret string
		return ret
	}
	return *o.RelativeName
}

// GetRelativeNameOk returns a tuple with the RelativeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRelativeNameOk() (*string, bool) {
	if o == nil || isNil(o.RelativeName) {
    return nil, false
	}
	return o.RelativeName, true
}

// HasRelativeName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRelativeName() bool {
	if o != nil && !isNil(o.RelativeName) {
		return true
	}

	return false
}

// SetRelativeName gets a reference to the given string and assigns it to the RelativeName field.
func (o *ValidateResponseKycResult) SetRelativeName(v string) {
	o.RelativeName = &v
}

// GetRelativeNameInVernacular returns the RelativeNameInVernacular field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRelativeNameInVernacular() string {
	if o == nil || isNil(o.RelativeNameInVernacular) {
		var ret string
		return ret
	}
	return *o.RelativeNameInVernacular
}

// GetRelativeNameInVernacularOk returns a tuple with the RelativeNameInVernacular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRelativeNameInVernacularOk() (*string, bool) {
	if o == nil || isNil(o.RelativeNameInVernacular) {
    return nil, false
	}
	return o.RelativeNameInVernacular, true
}

// HasRelativeNameInVernacular returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRelativeNameInVernacular() bool {
	if o != nil && !isNil(o.RelativeNameInVernacular) {
		return true
	}

	return false
}

// SetRelativeNameInVernacular gets a reference to the given string and assigns it to the RelativeNameInVernacular field.
func (o *ValidateResponseKycResult) SetRelativeNameInVernacular(v string) {
	o.RelativeNameInVernacular = &v
}

// GetRelativeRelationType returns the RelativeRelationType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRelativeRelationType() string {
	if o == nil || isNil(o.RelativeRelationType) {
		var ret string
		return ret
	}
	return *o.RelativeRelationType
}

// GetRelativeRelationTypeOk returns a tuple with the RelativeRelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRelativeRelationTypeOk() (*string, bool) {
	if o == nil || isNil(o.RelativeRelationType) {
    return nil, false
	}
	return o.RelativeRelationType, true
}

// HasRelativeRelationType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRelativeRelationType() bool {
	if o != nil && !isNil(o.RelativeRelationType) {
		return true
	}

	return false
}

// SetRelativeRelationType gets a reference to the given string and assigns it to the RelativeRelationType field.
func (o *ValidateResponseKycResult) SetRelativeRelationType(v string) {
	o.RelativeRelationType = &v
}

// GetHouseNumber returns the HouseNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetHouseNumber() string {
	if o == nil || isNil(o.HouseNumber) {
		var ret string
		return ret
	}
	return *o.HouseNumber
}

// GetHouseNumberOk returns a tuple with the HouseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetHouseNumberOk() (*string, bool) {
	if o == nil || isNil(o.HouseNumber) {
    return nil, false
	}
	return o.HouseNumber, true
}

// HasHouseNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasHouseNumber() bool {
	if o != nil && !isNil(o.HouseNumber) {
		return true
	}

	return false
}

// SetHouseNumber gets a reference to the given string and assigns it to the HouseNumber field.
func (o *ValidateResponseKycResult) SetHouseNumber(v string) {
	o.HouseNumber = &v
}

// GetPartOrLocationInConstituency returns the PartOrLocationInConstituency field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPartOrLocationInConstituency() string {
	if o == nil || isNil(o.PartOrLocationInConstituency) {
		var ret string
		return ret
	}
	return *o.PartOrLocationInConstituency
}

// GetPartOrLocationInConstituencyOk returns a tuple with the PartOrLocationInConstituency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPartOrLocationInConstituencyOk() (*string, bool) {
	if o == nil || isNil(o.PartOrLocationInConstituency) {
    return nil, false
	}
	return o.PartOrLocationInConstituency, true
}

// HasPartOrLocationInConstituency returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPartOrLocationInConstituency() bool {
	if o != nil && !isNil(o.PartOrLocationInConstituency) {
		return true
	}

	return false
}

// SetPartOrLocationInConstituency gets a reference to the given string and assigns it to the PartOrLocationInConstituency field.
func (o *ValidateResponseKycResult) SetPartOrLocationInConstituency(v string) {
	o.PartOrLocationInConstituency = &v
}

// GetPartNumberOrLocationNumberInConstituency returns the PartNumberOrLocationNumberInConstituency field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPartNumberOrLocationNumberInConstituency() string {
	if o == nil || isNil(o.PartNumberOrLocationNumberInConstituency) {
		var ret string
		return ret
	}
	return *o.PartNumberOrLocationNumberInConstituency
}

// GetPartNumberOrLocationNumberInConstituencyOk returns a tuple with the PartNumberOrLocationNumberInConstituency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPartNumberOrLocationNumberInConstituencyOk() (*string, bool) {
	if o == nil || isNil(o.PartNumberOrLocationNumberInConstituency) {
    return nil, false
	}
	return o.PartNumberOrLocationNumberInConstituency, true
}

// HasPartNumberOrLocationNumberInConstituency returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPartNumberOrLocationNumberInConstituency() bool {
	if o != nil && !isNil(o.PartNumberOrLocationNumberInConstituency) {
		return true
	}

	return false
}

// SetPartNumberOrLocationNumberInConstituency gets a reference to the given string and assigns it to the PartNumberOrLocationNumberInConstituency field.
func (o *ValidateResponseKycResult) SetPartNumberOrLocationNumberInConstituency(v string) {
	o.PartNumberOrLocationNumberInConstituency = &v
}

// GetParliamentaryConstituency returns the ParliamentaryConstituency field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetParliamentaryConstituency() string {
	if o == nil || isNil(o.ParliamentaryConstituency) {
		var ret string
		return ret
	}
	return *o.ParliamentaryConstituency
}

// GetParliamentaryConstituencyOk returns a tuple with the ParliamentaryConstituency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetParliamentaryConstituencyOk() (*string, bool) {
	if o == nil || isNil(o.ParliamentaryConstituency) {
    return nil, false
	}
	return o.ParliamentaryConstituency, true
}

// HasParliamentaryConstituency returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasParliamentaryConstituency() bool {
	if o != nil && !isNil(o.ParliamentaryConstituency) {
		return true
	}

	return false
}

// SetParliamentaryConstituency gets a reference to the given string and assigns it to the ParliamentaryConstituency field.
func (o *ValidateResponseKycResult) SetParliamentaryConstituency(v string) {
	o.ParliamentaryConstituency = &v
}

// GetAssemblyConstituency returns the AssemblyConstituency field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAssemblyConstituency() string {
	if o == nil || isNil(o.AssemblyConstituency) {
		var ret string
		return ret
	}
	return *o.AssemblyConstituency
}

// GetAssemblyConstituencyOk returns a tuple with the AssemblyConstituency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAssemblyConstituencyOk() (*string, bool) {
	if o == nil || isNil(o.AssemblyConstituency) {
    return nil, false
	}
	return o.AssemblyConstituency, true
}

// HasAssemblyConstituency returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAssemblyConstituency() bool {
	if o != nil && !isNil(o.AssemblyConstituency) {
		return true
	}

	return false
}

// SetAssemblyConstituency gets a reference to the given string and assigns it to the AssemblyConstituency field.
func (o *ValidateResponseKycResult) SetAssemblyConstituency(v string) {
	o.AssemblyConstituency = &v
}

// GetSectionOfConstituencyPart returns the SectionOfConstituencyPart field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetSectionOfConstituencyPart() string {
	if o == nil || isNil(o.SectionOfConstituencyPart) {
		var ret string
		return ret
	}
	return *o.SectionOfConstituencyPart
}

// GetSectionOfConstituencyPartOk returns a tuple with the SectionOfConstituencyPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetSectionOfConstituencyPartOk() (*string, bool) {
	if o == nil || isNil(o.SectionOfConstituencyPart) {
    return nil, false
	}
	return o.SectionOfConstituencyPart, true
}

// HasSectionOfConstituencyPart returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasSectionOfConstituencyPart() bool {
	if o != nil && !isNil(o.SectionOfConstituencyPart) {
		return true
	}

	return false
}

// SetSectionOfConstituencyPart gets a reference to the given string and assigns it to the SectionOfConstituencyPart field.
func (o *ValidateResponseKycResult) SetSectionOfConstituencyPart(v string) {
	o.SectionOfConstituencyPart = &v
}

// GetCardSerialNumberInPollingList returns the CardSerialNumberInPollingList field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCardSerialNumberInPollingList() string {
	if o == nil || isNil(o.CardSerialNumberInPollingList) {
		var ret string
		return ret
	}
	return *o.CardSerialNumberInPollingList
}

// GetCardSerialNumberInPollingListOk returns a tuple with the CardSerialNumberInPollingList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCardSerialNumberInPollingListOk() (*string, bool) {
	if o == nil || isNil(o.CardSerialNumberInPollingList) {
    return nil, false
	}
	return o.CardSerialNumberInPollingList, true
}

// HasCardSerialNumberInPollingList returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCardSerialNumberInPollingList() bool {
	if o != nil && !isNil(o.CardSerialNumberInPollingList) {
		return true
	}

	return false
}

// SetCardSerialNumberInPollingList gets a reference to the given string and assigns it to the CardSerialNumberInPollingList field.
func (o *ValidateResponseKycResult) SetCardSerialNumberInPollingList(v string) {
	o.CardSerialNumberInPollingList = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLastUpdateDate() string {
	if o == nil || isNil(o.LastUpdateDate) {
		var ret string
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLastUpdateDateOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdateDate) {
    return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLastUpdateDate() bool {
	if o != nil && !isNil(o.LastUpdateDate) {
		return true
	}

	return false
}

// SetLastUpdateDate gets a reference to the given string and assigns it to the LastUpdateDate field.
func (o *ValidateResponseKycResult) SetLastUpdateDate(v string) {
	o.LastUpdateDate = &v
}

// GetPollingBoothDetails returns the PollingBoothDetails field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollingBoothDetails() ValidateResponseKycResultPollingBoothDetails {
	if o == nil || isNil(o.PollingBoothDetails) {
		var ret ValidateResponseKycResultPollingBoothDetails
		return ret
	}
	return *o.PollingBoothDetails
}

// GetPollingBoothDetailsOk returns a tuple with the PollingBoothDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollingBoothDetailsOk() (*ValidateResponseKycResultPollingBoothDetails, bool) {
	if o == nil || isNil(o.PollingBoothDetails) {
    return nil, false
	}
	return o.PollingBoothDetails, true
}

// HasPollingBoothDetails returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollingBoothDetails() bool {
	if o != nil && !isNil(o.PollingBoothDetails) {
		return true
	}

	return false
}

// SetPollingBoothDetails gets a reference to the given ValidateResponseKycResultPollingBoothDetails and assigns it to the PollingBoothDetails field.
func (o *ValidateResponseKycResult) SetPollingBoothDetails(v ValidateResponseKycResultPollingBoothDetails) {
	o.PollingBoothDetails = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEmailId() string {
	if o == nil || isNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEmailIdOk() (*string, bool) {
	if o == nil || isNil(o.EmailId) {
    return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEmailId() bool {
	if o != nil && !isNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ValidateResponseKycResult) SetEmailId(v string) {
	o.EmailId = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetMobileNumber() string {
	if o == nil || isNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetMobileNumberOk() (*string, bool) {
	if o == nil || isNil(o.MobileNumber) {
    return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasMobileNumber() bool {
	if o != nil && !isNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *ValidateResponseKycResult) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStateCode() string {
	if o == nil || isNil(o.StateCode) {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStateCodeOk() (*string, bool) {
	if o == nil || isNil(o.StateCode) {
    return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStateCode() bool {
	if o != nil && !isNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *ValidateResponseKycResult) SetStateCode(v string) {
	o.StateCode = &v
}

// GetPollingBoothCoordinates returns the PollingBoothCoordinates field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollingBoothCoordinates() string {
	if o == nil || isNil(o.PollingBoothCoordinates) {
		var ret string
		return ret
	}
	return *o.PollingBoothCoordinates
}

// GetPollingBoothCoordinatesOk returns a tuple with the PollingBoothCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollingBoothCoordinatesOk() (*string, bool) {
	if o == nil || isNil(o.PollingBoothCoordinates) {
    return nil, false
	}
	return o.PollingBoothCoordinates, true
}

// HasPollingBoothCoordinates returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollingBoothCoordinates() bool {
	if o != nil && !isNil(o.PollingBoothCoordinates) {
		return true
	}

	return false
}

// SetPollingBoothCoordinates gets a reference to the given string and assigns it to the PollingBoothCoordinates field.
func (o *ValidateResponseKycResult) SetPollingBoothCoordinates(v string) {
	o.PollingBoothCoordinates = &v
}

// GetPollingBoothAddress returns the PollingBoothAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollingBoothAddress() string {
	if o == nil || isNil(o.PollingBoothAddress) {
		var ret string
		return ret
	}
	return *o.PollingBoothAddress
}

// GetPollingBoothAddressOk returns a tuple with the PollingBoothAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollingBoothAddressOk() (*string, bool) {
	if o == nil || isNil(o.PollingBoothAddress) {
    return nil, false
	}
	return o.PollingBoothAddress, true
}

// HasPollingBoothAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollingBoothAddress() bool {
	if o != nil && !isNil(o.PollingBoothAddress) {
		return true
	}

	return false
}

// SetPollingBoothAddress gets a reference to the given string and assigns it to the PollingBoothAddress field.
func (o *ValidateResponseKycResult) SetPollingBoothAddress(v string) {
	o.PollingBoothAddress = &v
}

// GetPollingBoothNumber returns the PollingBoothNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollingBoothNumber() string {
	if o == nil || isNil(o.PollingBoothNumber) {
		var ret string
		return ret
	}
	return *o.PollingBoothNumber
}

// GetPollingBoothNumberOk returns a tuple with the PollingBoothNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollingBoothNumberOk() (*string, bool) {
	if o == nil || isNil(o.PollingBoothNumber) {
    return nil, false
	}
	return o.PollingBoothNumber, true
}

// HasPollingBoothNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollingBoothNumber() bool {
	if o != nil && !isNil(o.PollingBoothNumber) {
		return true
	}

	return false
}

// SetPollingBoothNumber gets a reference to the given string and assigns it to the PollingBoothNumber field.
func (o *ValidateResponseKycResult) SetPollingBoothNumber(v string) {
	o.PollingBoothNumber = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ValidateResponseKycResult) SetId(v string) {
	o.Id = &v
}

// GetBlacklistStatus returns the BlacklistStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetBlacklistStatus() string {
	if o == nil || isNil(o.BlacklistStatus) {
		var ret string
		return ret
	}
	return *o.BlacklistStatus
}

// GetBlacklistStatusOk returns a tuple with the BlacklistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetBlacklistStatusOk() (*string, bool) {
	if o == nil || isNil(o.BlacklistStatus) {
    return nil, false
	}
	return o.BlacklistStatus, true
}

// HasBlacklistStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasBlacklistStatus() bool {
	if o != nil && !isNil(o.BlacklistStatus) {
		return true
	}

	return false
}

// SetBlacklistStatus gets a reference to the given string and assigns it to the BlacklistStatus field.
func (o *ValidateResponseKycResult) SetBlacklistStatus(v string) {
	o.BlacklistStatus = &v
}

// GetRegistrationDate returns the RegistrationDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRegistrationDate() string {
	if o == nil || isNil(o.RegistrationDate) {
		var ret string
		return ret
	}
	return *o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRegistrationDateOk() (*string, bool) {
	if o == nil || isNil(o.RegistrationDate) {
    return nil, false
	}
	return o.RegistrationDate, true
}

// HasRegistrationDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRegistrationDate() bool {
	if o != nil && !isNil(o.RegistrationDate) {
		return true
	}

	return false
}

// SetRegistrationDate gets a reference to the given string and assigns it to the RegistrationDate field.
func (o *ValidateResponseKycResult) SetRegistrationDate(v string) {
	o.RegistrationDate = &v
}

// GetRegistrationLocation returns the RegistrationLocation field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRegistrationLocation() string {
	if o == nil || isNil(o.RegistrationLocation) {
		var ret string
		return ret
	}
	return *o.RegistrationLocation
}

// GetRegistrationLocationOk returns a tuple with the RegistrationLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRegistrationLocationOk() (*string, bool) {
	if o == nil || isNil(o.RegistrationLocation) {
    return nil, false
	}
	return o.RegistrationLocation, true
}

// HasRegistrationLocation returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRegistrationLocation() bool {
	if o != nil && !isNil(o.RegistrationLocation) {
		return true
	}

	return false
}

// SetRegistrationLocation gets a reference to the given string and assigns it to the RegistrationLocation field.
func (o *ValidateResponseKycResult) SetRegistrationLocation(v string) {
	o.RegistrationLocation = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetClass() string {
	if o == nil || isNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetClassOk() (*string, bool) {
	if o == nil || isNil(o.Class) {
    return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasClass() bool {
	if o != nil && !isNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *ValidateResponseKycResult) SetClass(v string) {
	o.Class = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetMaker() string {
	if o == nil || isNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetMakerOk() (*string, bool) {
	if o == nil || isNil(o.Maker) {
    return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasMaker() bool {
	if o != nil && !isNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *ValidateResponseKycResult) SetMaker(v string) {
	o.Maker = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetOwnerName() string {
	if o == nil || isNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetOwnerNameOk() (*string, bool) {
	if o == nil || isNil(o.OwnerName) {
    return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasOwnerName() bool {
	if o != nil && !isNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *ValidateResponseKycResult) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetChassisNumber returns the ChassisNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetChassisNumber() string {
	if o == nil || isNil(o.ChassisNumber) {
		var ret string
		return ret
	}
	return *o.ChassisNumber
}

// GetChassisNumberOk returns a tuple with the ChassisNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetChassisNumberOk() (*string, bool) {
	if o == nil || isNil(o.ChassisNumber) {
    return nil, false
	}
	return o.ChassisNumber, true
}

// HasChassisNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasChassisNumber() bool {
	if o != nil && !isNil(o.ChassisNumber) {
		return true
	}

	return false
}

// SetChassisNumber gets a reference to the given string and assigns it to the ChassisNumber field.
func (o *ValidateResponseKycResult) SetChassisNumber(v string) {
	o.ChassisNumber = &v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRegistrationNumber() string {
	if o == nil || isNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || isNil(o.RegistrationNumber) {
    return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRegistrationNumber() bool {
	if o != nil && !isNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *ValidateResponseKycResult) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetEngineNumber returns the EngineNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEngineNumber() string {
	if o == nil || isNil(o.EngineNumber) {
		var ret string
		return ret
	}
	return *o.EngineNumber
}

// GetEngineNumberOk returns a tuple with the EngineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEngineNumberOk() (*string, bool) {
	if o == nil || isNil(o.EngineNumber) {
    return nil, false
	}
	return o.EngineNumber, true
}

// HasEngineNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEngineNumber() bool {
	if o != nil && !isNil(o.EngineNumber) {
		return true
	}

	return false
}

// SetEngineNumber gets a reference to the given string and assigns it to the EngineNumber field.
func (o *ValidateResponseKycResult) SetEngineNumber(v string) {
	o.EngineNumber = &v
}

// GetFuelType returns the FuelType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFuelType() string {
	if o == nil || isNil(o.FuelType) {
		var ret string
		return ret
	}
	return *o.FuelType
}

// GetFuelTypeOk returns a tuple with the FuelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFuelTypeOk() (*string, bool) {
	if o == nil || isNil(o.FuelType) {
    return nil, false
	}
	return o.FuelType, true
}

// HasFuelType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFuelType() bool {
	if o != nil && !isNil(o.FuelType) {
		return true
	}

	return false
}

// SetFuelType gets a reference to the given string and assigns it to the FuelType field.
func (o *ValidateResponseKycResult) SetFuelType(v string) {
	o.FuelType = &v
}

// GetFitUpto returns the FitUpto field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFitUpto() string {
	if o == nil || isNil(o.FitUpto) {
		var ret string
		return ret
	}
	return *o.FitUpto
}

// GetFitUptoOk returns a tuple with the FitUpto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFitUptoOk() (*string, bool) {
	if o == nil || isNil(o.FitUpto) {
    return nil, false
	}
	return o.FitUpto, true
}

// HasFitUpto returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFitUpto() bool {
	if o != nil && !isNil(o.FitUpto) {
		return true
	}

	return false
}

// SetFitUpto gets a reference to the given string and assigns it to the FitUpto field.
func (o *ValidateResponseKycResult) SetFitUpto(v string) {
	o.FitUpto = &v
}

// GetInsuranceUpto returns the InsuranceUpto field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetInsuranceUpto() string {
	if o == nil || isNil(o.InsuranceUpto) {
		var ret string
		return ret
	}
	return *o.InsuranceUpto
}

// GetInsuranceUptoOk returns a tuple with the InsuranceUpto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetInsuranceUptoOk() (*string, bool) {
	if o == nil || isNil(o.InsuranceUpto) {
    return nil, false
	}
	return o.InsuranceUpto, true
}

// HasInsuranceUpto returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasInsuranceUpto() bool {
	if o != nil && !isNil(o.InsuranceUpto) {
		return true
	}

	return false
}

// SetInsuranceUpto gets a reference to the given string and assigns it to the InsuranceUpto field.
func (o *ValidateResponseKycResult) SetInsuranceUpto(v string) {
	o.InsuranceUpto = &v
}

// GetTaxUpto returns the TaxUpto field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetTaxUpto() string {
	if o == nil || isNil(o.TaxUpto) {
		var ret string
		return ret
	}
	return *o.TaxUpto
}

// GetTaxUptoOk returns a tuple with the TaxUpto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetTaxUptoOk() (*string, bool) {
	if o == nil || isNil(o.TaxUpto) {
    return nil, false
	}
	return o.TaxUpto, true
}

// HasTaxUpto returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasTaxUpto() bool {
	if o != nil && !isNil(o.TaxUpto) {
		return true
	}

	return false
}

// SetTaxUpto gets a reference to the given string and assigns it to the TaxUpto field.
func (o *ValidateResponseKycResult) SetTaxUpto(v string) {
	o.TaxUpto = &v
}

// GetInsuranceDetails returns the InsuranceDetails field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetInsuranceDetails() string {
	if o == nil || isNil(o.InsuranceDetails) {
		var ret string
		return ret
	}
	return *o.InsuranceDetails
}

// GetInsuranceDetailsOk returns a tuple with the InsuranceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetInsuranceDetailsOk() (*string, bool) {
	if o == nil || isNil(o.InsuranceDetails) {
    return nil, false
	}
	return o.InsuranceDetails, true
}

// HasInsuranceDetails returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasInsuranceDetails() bool {
	if o != nil && !isNil(o.InsuranceDetails) {
		return true
	}

	return false
}

// SetInsuranceDetails gets a reference to the given string and assigns it to the InsuranceDetails field.
func (o *ValidateResponseKycResult) SetInsuranceDetails(v string) {
	o.InsuranceDetails = &v
}

// GetInsuranceValidity returns the InsuranceValidity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetInsuranceValidity() string {
	if o == nil || isNil(o.InsuranceValidity) {
		var ret string
		return ret
	}
	return *o.InsuranceValidity
}

// GetInsuranceValidityOk returns a tuple with the InsuranceValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetInsuranceValidityOk() (*string, bool) {
	if o == nil || isNil(o.InsuranceValidity) {
    return nil, false
	}
	return o.InsuranceValidity, true
}

// HasInsuranceValidity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasInsuranceValidity() bool {
	if o != nil && !isNil(o.InsuranceValidity) {
		return true
	}

	return false
}

// SetInsuranceValidity gets a reference to the given string and assigns it to the InsuranceValidity field.
func (o *ValidateResponseKycResult) SetInsuranceValidity(v string) {
	o.InsuranceValidity = &v
}

// GetPermitType returns the PermitType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPermitType() string {
	if o == nil || isNil(o.PermitType) {
		var ret string
		return ret
	}
	return *o.PermitType
}

// GetPermitTypeOk returns a tuple with the PermitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPermitTypeOk() (*string, bool) {
	if o == nil || isNil(o.PermitType) {
    return nil, false
	}
	return o.PermitType, true
}

// HasPermitType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPermitType() bool {
	if o != nil && !isNil(o.PermitType) {
		return true
	}

	return false
}

// SetPermitType gets a reference to the given string and assigns it to the PermitType field.
func (o *ValidateResponseKycResult) SetPermitType(v string) {
	o.PermitType = &v
}

// GetPermitValidUpto returns the PermitValidUpto field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPermitValidUpto() string {
	if o == nil || isNil(o.PermitValidUpto) {
		var ret string
		return ret
	}
	return *o.PermitValidUpto
}

// GetPermitValidUptoOk returns a tuple with the PermitValidUpto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPermitValidUptoOk() (*string, bool) {
	if o == nil || isNil(o.PermitValidUpto) {
    return nil, false
	}
	return o.PermitValidUpto, true
}

// HasPermitValidUpto returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPermitValidUpto() bool {
	if o != nil && !isNil(o.PermitValidUpto) {
		return true
	}

	return false
}

// SetPermitValidUpto gets a reference to the given string and assigns it to the PermitValidUpto field.
func (o *ValidateResponseKycResult) SetPermitValidUpto(v string) {
	o.PermitValidUpto = &v
}

// GetPollutionControlValidity returns the PollutionControlValidity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollutionControlValidity() string {
	if o == nil || isNil(o.PollutionControlValidity) {
		var ret string
		return ret
	}
	return *o.PollutionControlValidity
}

// GetPollutionControlValidityOk returns a tuple with the PollutionControlValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollutionControlValidityOk() (*string, bool) {
	if o == nil || isNil(o.PollutionControlValidity) {
    return nil, false
	}
	return o.PollutionControlValidity, true
}

// HasPollutionControlValidity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollutionControlValidity() bool {
	if o != nil && !isNil(o.PollutionControlValidity) {
		return true
	}

	return false
}

// SetPollutionControlValidity gets a reference to the given string and assigns it to the PollutionControlValidity field.
func (o *ValidateResponseKycResult) SetPollutionControlValidity(v string) {
	o.PollutionControlValidity = &v
}

// GetPollutionNorms returns the PollutionNorms field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPollutionNorms() string {
	if o == nil || isNil(o.PollutionNorms) {
		var ret string
		return ret
	}
	return *o.PollutionNorms
}

// GetPollutionNormsOk returns a tuple with the PollutionNorms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPollutionNormsOk() (*string, bool) {
	if o == nil || isNil(o.PollutionNorms) {
    return nil, false
	}
	return o.PollutionNorms, true
}

// HasPollutionNorms returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPollutionNorms() bool {
	if o != nil && !isNil(o.PollutionNorms) {
		return true
	}

	return false
}

// SetPollutionNorms gets a reference to the given string and assigns it to the PollutionNorms field.
func (o *ValidateResponseKycResult) SetPollutionNorms(v string) {
	o.PollutionNorms = &v
}

// GetLicenseAddress returns the LicenseAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLicenseAddress() string {
	if o == nil || isNil(o.LicenseAddress) {
		var ret string
		return ret
	}
	return *o.LicenseAddress
}

// GetLicenseAddressOk returns a tuple with the LicenseAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLicenseAddressOk() (*string, bool) {
	if o == nil || isNil(o.LicenseAddress) {
    return nil, false
	}
	return o.LicenseAddress, true
}

// HasLicenseAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLicenseAddress() bool {
	if o != nil && !isNil(o.LicenseAddress) {
		return true
	}

	return false
}

// SetLicenseAddress gets a reference to the given string and assigns it to the LicenseAddress field.
func (o *ValidateResponseKycResult) SetLicenseAddress(v string) {
	o.LicenseAddress = &v
}

// GetRegistrationAddress returns the RegistrationAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetRegistrationAddress() string {
	if o == nil || isNil(o.RegistrationAddress) {
		var ret string
		return ret
	}
	return *o.RegistrationAddress
}

// GetRegistrationAddressOk returns a tuple with the RegistrationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetRegistrationAddressOk() (*string, bool) {
	if o == nil || isNil(o.RegistrationAddress) {
    return nil, false
	}
	return o.RegistrationAddress, true
}

// HasRegistrationAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasRegistrationAddress() bool {
	if o != nil && !isNil(o.RegistrationAddress) {
		return true
	}

	return false
}

// SetRegistrationAddress gets a reference to the given string and assigns it to the RegistrationAddress field.
func (o *ValidateResponseKycResult) SetRegistrationAddress(v string) {
	o.RegistrationAddress = &v
}

// GetOwnerFatherName returns the OwnerFatherName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetOwnerFatherName() string {
	if o == nil || isNil(o.OwnerFatherName) {
		var ret string
		return ret
	}
	return *o.OwnerFatherName
}

// GetOwnerFatherNameOk returns a tuple with the OwnerFatherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetOwnerFatherNameOk() (*string, bool) {
	if o == nil || isNil(o.OwnerFatherName) {
    return nil, false
	}
	return o.OwnerFatherName, true
}

// HasOwnerFatherName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasOwnerFatherName() bool {
	if o != nil && !isNil(o.OwnerFatherName) {
		return true
	}

	return false
}

// SetOwnerFatherName gets a reference to the given string and assigns it to the OwnerFatherName field.
func (o *ValidateResponseKycResult) SetOwnerFatherName(v string) {
	o.OwnerFatherName = &v
}

// GetOwnerPresentAddress returns the OwnerPresentAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetOwnerPresentAddress() string {
	if o == nil || isNil(o.OwnerPresentAddress) {
		var ret string
		return ret
	}
	return *o.OwnerPresentAddress
}

// GetOwnerPresentAddressOk returns a tuple with the OwnerPresentAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetOwnerPresentAddressOk() (*string, bool) {
	if o == nil || isNil(o.OwnerPresentAddress) {
    return nil, false
	}
	return o.OwnerPresentAddress, true
}

// HasOwnerPresentAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasOwnerPresentAddress() bool {
	if o != nil && !isNil(o.OwnerPresentAddress) {
		return true
	}

	return false
}

// SetOwnerPresentAddress gets a reference to the given string and assigns it to the OwnerPresentAddress field.
func (o *ValidateResponseKycResult) SetOwnerPresentAddress(v string) {
	o.OwnerPresentAddress = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetBodyType() string {
	if o == nil || isNil(o.BodyType) {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetBodyTypeOk() (*string, bool) {
	if o == nil || isNil(o.BodyType) {
    return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasBodyType() bool {
	if o != nil && !isNil(o.BodyType) {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *ValidateResponseKycResult) SetBodyType(v string) {
	o.BodyType = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ValidateResponseKycResult) SetCategory(v string) {
	o.Category = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
    return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ValidateResponseKycResult) SetColor(v string) {
	o.Color = &v
}

// GetEngineCubicCapacity returns the EngineCubicCapacity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetEngineCubicCapacity() string {
	if o == nil || isNil(o.EngineCubicCapacity) {
		var ret string
		return ret
	}
	return *o.EngineCubicCapacity
}

// GetEngineCubicCapacityOk returns a tuple with the EngineCubicCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetEngineCubicCapacityOk() (*string, bool) {
	if o == nil || isNil(o.EngineCubicCapacity) {
    return nil, false
	}
	return o.EngineCubicCapacity, true
}

// HasEngineCubicCapacity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasEngineCubicCapacity() bool {
	if o != nil && !isNil(o.EngineCubicCapacity) {
		return true
	}

	return false
}

// SetEngineCubicCapacity gets a reference to the given string and assigns it to the EngineCubicCapacity field.
func (o *ValidateResponseKycResult) SetEngineCubicCapacity(v string) {
	o.EngineCubicCapacity = &v
}

// GetNumberCylinders returns the NumberCylinders field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNumberCylinders() string {
	if o == nil || isNil(o.NumberCylinders) {
		var ret string
		return ret
	}
	return *o.NumberCylinders
}

// GetNumberCylindersOk returns a tuple with the NumberCylinders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNumberCylindersOk() (*string, bool) {
	if o == nil || isNil(o.NumberCylinders) {
    return nil, false
	}
	return o.NumberCylinders, true
}

// HasNumberCylinders returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNumberCylinders() bool {
	if o != nil && !isNil(o.NumberCylinders) {
		return true
	}

	return false
}

// SetNumberCylinders gets a reference to the given string and assigns it to the NumberCylinders field.
func (o *ValidateResponseKycResult) SetNumberCylinders(v string) {
	o.NumberCylinders = &v
}

// GetUnladenWeight returns the UnladenWeight field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetUnladenWeight() string {
	if o == nil || isNil(o.UnladenWeight) {
		var ret string
		return ret
	}
	return *o.UnladenWeight
}

// GetUnladenWeightOk returns a tuple with the UnladenWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetUnladenWeightOk() (*string, bool) {
	if o == nil || isNil(o.UnladenWeight) {
    return nil, false
	}
	return o.UnladenWeight, true
}

// HasUnladenWeight returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasUnladenWeight() bool {
	if o != nil && !isNil(o.UnladenWeight) {
		return true
	}

	return false
}

// SetUnladenWeight gets a reference to the given string and assigns it to the UnladenWeight field.
func (o *ValidateResponseKycResult) SetUnladenWeight(v string) {
	o.UnladenWeight = &v
}

// GetGrossWeight returns the GrossWeight field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGrossWeight() string {
	if o == nil || isNil(o.GrossWeight) {
		var ret string
		return ret
	}
	return *o.GrossWeight
}

// GetGrossWeightOk returns a tuple with the GrossWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGrossWeightOk() (*string, bool) {
	if o == nil || isNil(o.GrossWeight) {
    return nil, false
	}
	return o.GrossWeight, true
}

// HasGrossWeight returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGrossWeight() bool {
	if o != nil && !isNil(o.GrossWeight) {
		return true
	}

	return false
}

// SetGrossWeight gets a reference to the given string and assigns it to the GrossWeight field.
func (o *ValidateResponseKycResult) SetGrossWeight(v string) {
	o.GrossWeight = &v
}

// GetWheelbase returns the Wheelbase field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetWheelbase() string {
	if o == nil || isNil(o.Wheelbase) {
		var ret string
		return ret
	}
	return *o.Wheelbase
}

// GetWheelbaseOk returns a tuple with the Wheelbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetWheelbaseOk() (*string, bool) {
	if o == nil || isNil(o.Wheelbase) {
    return nil, false
	}
	return o.Wheelbase, true
}

// HasWheelbase returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasWheelbase() bool {
	if o != nil && !isNil(o.Wheelbase) {
		return true
	}

	return false
}

// SetWheelbase gets a reference to the given string and assigns it to the Wheelbase field.
func (o *ValidateResponseKycResult) SetWheelbase(v string) {
	o.Wheelbase = &v
}

// GetManufacturedMonthYear returns the ManufacturedMonthYear field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetManufacturedMonthYear() string {
	if o == nil || isNil(o.ManufacturedMonthYear) {
		var ret string
		return ret
	}
	return *o.ManufacturedMonthYear
}

// GetManufacturedMonthYearOk returns a tuple with the ManufacturedMonthYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetManufacturedMonthYearOk() (*string, bool) {
	if o == nil || isNil(o.ManufacturedMonthYear) {
    return nil, false
	}
	return o.ManufacturedMonthYear, true
}

// HasManufacturedMonthYear returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasManufacturedMonthYear() bool {
	if o != nil && !isNil(o.ManufacturedMonthYear) {
		return true
	}

	return false
}

// SetManufacturedMonthYear gets a reference to the given string and assigns it to the ManufacturedMonthYear field.
func (o *ValidateResponseKycResult) SetManufacturedMonthYear(v string) {
	o.ManufacturedMonthYear = &v
}

// GetMakerDescription returns the MakerDescription field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetMakerDescription() string {
	if o == nil || isNil(o.MakerDescription) {
		var ret string
		return ret
	}
	return *o.MakerDescription
}

// GetMakerDescriptionOk returns a tuple with the MakerDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetMakerDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.MakerDescription) {
    return nil, false
	}
	return o.MakerDescription, true
}

// HasMakerDescription returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasMakerDescription() bool {
	if o != nil && !isNil(o.MakerDescription) {
		return true
	}

	return false
}

// SetMakerDescription gets a reference to the given string and assigns it to the MakerDescription field.
func (o *ValidateResponseKycResult) SetMakerDescription(v string) {
	o.MakerDescription = &v
}

// GetNocDetails returns the NocDetails field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNocDetails() string {
	if o == nil || isNil(o.NocDetails) {
		var ret string
		return ret
	}
	return *o.NocDetails
}

// GetNocDetailsOk returns a tuple with the NocDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNocDetailsOk() (*string, bool) {
	if o == nil || isNil(o.NocDetails) {
    return nil, false
	}
	return o.NocDetails, true
}

// HasNocDetails returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNocDetails() bool {
	if o != nil && !isNil(o.NocDetails) {
		return true
	}

	return false
}

// SetNocDetails gets a reference to the given string and assigns it to the NocDetails field.
func (o *ValidateResponseKycResult) SetNocDetails(v string) {
	o.NocDetails = &v
}

// GetNormsDescription returns the NormsDescription field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNormsDescription() string {
	if o == nil || isNil(o.NormsDescription) {
		var ret string
		return ret
	}
	return *o.NormsDescription
}

// GetNormsDescriptionOk returns a tuple with the NormsDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNormsDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.NormsDescription) {
    return nil, false
	}
	return o.NormsDescription, true
}

// HasNormsDescription returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNormsDescription() bool {
	if o != nil && !isNil(o.NormsDescription) {
		return true
	}

	return false
}

// SetNormsDescription gets a reference to the given string and assigns it to the NormsDescription field.
func (o *ValidateResponseKycResult) SetNormsDescription(v string) {
	o.NormsDescription = &v
}

// GetFinancier returns the Financier field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFinancier() string {
	if o == nil || isNil(o.Financier) {
		var ret string
		return ret
	}
	return *o.Financier
}

// GetFinancierOk returns a tuple with the Financier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFinancierOk() (*string, bool) {
	if o == nil || isNil(o.Financier) {
    return nil, false
	}
	return o.Financier, true
}

// HasFinancier returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFinancier() bool {
	if o != nil && !isNil(o.Financier) {
		return true
	}

	return false
}

// SetFinancier gets a reference to the given string and assigns it to the Financier field.
func (o *ValidateResponseKycResult) SetFinancier(v string) {
	o.Financier = &v
}

// GetPermitIssueDate returns the PermitIssueDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPermitIssueDate() string {
	if o == nil || isNil(o.PermitIssueDate) {
		var ret string
		return ret
	}
	return *o.PermitIssueDate
}

// GetPermitIssueDateOk returns a tuple with the PermitIssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPermitIssueDateOk() (*string, bool) {
	if o == nil || isNil(o.PermitIssueDate) {
    return nil, false
	}
	return o.PermitIssueDate, true
}

// HasPermitIssueDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPermitIssueDate() bool {
	if o != nil && !isNil(o.PermitIssueDate) {
		return true
	}

	return false
}

// SetPermitIssueDate gets a reference to the given string and assigns it to the PermitIssueDate field.
func (o *ValidateResponseKycResult) SetPermitIssueDate(v string) {
	o.PermitIssueDate = &v
}

// GetPermitNumber returns the PermitNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPermitNumber() string {
	if o == nil || isNil(o.PermitNumber) {
		var ret string
		return ret
	}
	return *o.PermitNumber
}

// GetPermitNumberOk returns a tuple with the PermitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPermitNumberOk() (*string, bool) {
	if o == nil || isNil(o.PermitNumber) {
    return nil, false
	}
	return o.PermitNumber, true
}

// HasPermitNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPermitNumber() bool {
	if o != nil && !isNil(o.PermitNumber) {
		return true
	}

	return false
}

// SetPermitNumber gets a reference to the given string and assigns it to the PermitNumber field.
func (o *ValidateResponseKycResult) SetPermitNumber(v string) {
	o.PermitNumber = &v
}

// GetPermitValidFrom returns the PermitValidFrom field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPermitValidFrom() string {
	if o == nil || isNil(o.PermitValidFrom) {
		var ret string
		return ret
	}
	return *o.PermitValidFrom
}

// GetPermitValidFromOk returns a tuple with the PermitValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPermitValidFromOk() (*string, bool) {
	if o == nil || isNil(o.PermitValidFrom) {
    return nil, false
	}
	return o.PermitValidFrom, true
}

// HasPermitValidFrom returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPermitValidFrom() bool {
	if o != nil && !isNil(o.PermitValidFrom) {
		return true
	}

	return false
}

// SetPermitValidFrom gets a reference to the given string and assigns it to the PermitValidFrom field.
func (o *ValidateResponseKycResult) SetPermitValidFrom(v string) {
	o.PermitValidFrom = &v
}

// GetSeatingCapacity returns the SeatingCapacity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetSeatingCapacity() string {
	if o == nil || isNil(o.SeatingCapacity) {
		var ret string
		return ret
	}
	return *o.SeatingCapacity
}

// GetSeatingCapacityOk returns a tuple with the SeatingCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetSeatingCapacityOk() (*string, bool) {
	if o == nil || isNil(o.SeatingCapacity) {
    return nil, false
	}
	return o.SeatingCapacity, true
}

// HasSeatingCapacity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasSeatingCapacity() bool {
	if o != nil && !isNil(o.SeatingCapacity) {
		return true
	}

	return false
}

// SetSeatingCapacity gets a reference to the given string and assigns it to the SeatingCapacity field.
func (o *ValidateResponseKycResult) SetSeatingCapacity(v string) {
	o.SeatingCapacity = &v
}

// GetSleepingCapacity returns the SleepingCapacity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetSleepingCapacity() string {
	if o == nil || isNil(o.SleepingCapacity) {
		var ret string
		return ret
	}
	return *o.SleepingCapacity
}

// GetSleepingCapacityOk returns a tuple with the SleepingCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetSleepingCapacityOk() (*string, bool) {
	if o == nil || isNil(o.SleepingCapacity) {
    return nil, false
	}
	return o.SleepingCapacity, true
}

// HasSleepingCapacity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasSleepingCapacity() bool {
	if o != nil && !isNil(o.SleepingCapacity) {
		return true
	}

	return false
}

// SetSleepingCapacity gets a reference to the given string and assigns it to the SleepingCapacity field.
func (o *ValidateResponseKycResult) SetSleepingCapacity(v string) {
	o.SleepingCapacity = &v
}

// GetStandingCapacity returns the StandingCapacity field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStandingCapacity() string {
	if o == nil || isNil(o.StandingCapacity) {
		var ret string
		return ret
	}
	return *o.StandingCapacity
}

// GetStandingCapacityOk returns a tuple with the StandingCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStandingCapacityOk() (*string, bool) {
	if o == nil || isNil(o.StandingCapacity) {
    return nil, false
	}
	return o.StandingCapacity, true
}

// HasStandingCapacity returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStandingCapacity() bool {
	if o != nil && !isNil(o.StandingCapacity) {
		return true
	}

	return false
}

// SetStandingCapacity gets a reference to the given string and assigns it to the StandingCapacity field.
func (o *ValidateResponseKycResult) SetStandingCapacity(v string) {
	o.StandingCapacity = &v
}

// GetStatusAsOn returns the StatusAsOn field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStatusAsOn() string {
	if o == nil || isNil(o.StatusAsOn) {
		var ret string
		return ret
	}
	return *o.StatusAsOn
}

// GetStatusAsOnOk returns a tuple with the StatusAsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStatusAsOnOk() (*string, bool) {
	if o == nil || isNil(o.StatusAsOn) {
    return nil, false
	}
	return o.StatusAsOn, true
}

// HasStatusAsOn returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStatusAsOn() bool {
	if o != nil && !isNil(o.StatusAsOn) {
		return true
	}

	return false
}

// SetStatusAsOn gets a reference to the given string and assigns it to the StatusAsOn field.
func (o *ValidateResponseKycResult) SetStatusAsOn(v string) {
	o.StatusAsOn = &v
}

// GetPrimaryBusinessContact returns the PrimaryBusinessContact field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPrimaryBusinessContact() ValidateResponseKycResultPrimaryBusinessContact {
	if o == nil || isNil(o.PrimaryBusinessContact) {
		var ret ValidateResponseKycResultPrimaryBusinessContact
		return ret
	}
	return *o.PrimaryBusinessContact
}

// GetPrimaryBusinessContactOk returns a tuple with the PrimaryBusinessContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPrimaryBusinessContactOk() (*ValidateResponseKycResultPrimaryBusinessContact, bool) {
	if o == nil || isNil(o.PrimaryBusinessContact) {
    return nil, false
	}
	return o.PrimaryBusinessContact, true
}

// HasPrimaryBusinessContact returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPrimaryBusinessContact() bool {
	if o != nil && !isNil(o.PrimaryBusinessContact) {
		return true
	}

	return false
}

// SetPrimaryBusinessContact gets a reference to the given ValidateResponseKycResultPrimaryBusinessContact and assigns it to the PrimaryBusinessContact field.
func (o *ValidateResponseKycResult) SetPrimaryBusinessContact(v ValidateResponseKycResultPrimaryBusinessContact) {
	o.PrimaryBusinessContact = &v
}

// GetStateJurisdiction returns the StateJurisdiction field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStateJurisdiction() string {
	if o == nil || isNil(o.StateJurisdiction) {
		var ret string
		return ret
	}
	return *o.StateJurisdiction
}

// GetStateJurisdictionOk returns a tuple with the StateJurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStateJurisdictionOk() (*string, bool) {
	if o == nil || isNil(o.StateJurisdiction) {
    return nil, false
	}
	return o.StateJurisdiction, true
}

// HasStateJurisdiction returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStateJurisdiction() bool {
	if o != nil && !isNil(o.StateJurisdiction) {
		return true
	}

	return false
}

// SetStateJurisdiction gets a reference to the given string and assigns it to the StateJurisdiction field.
func (o *ValidateResponseKycResult) SetStateJurisdiction(v string) {
	o.StateJurisdiction = &v
}

// GetStateJurisdictionCode returns the StateJurisdictionCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetStateJurisdictionCode() string {
	if o == nil || isNil(o.StateJurisdictionCode) {
		var ret string
		return ret
	}
	return *o.StateJurisdictionCode
}

// GetStateJurisdictionCodeOk returns a tuple with the StateJurisdictionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetStateJurisdictionCodeOk() (*string, bool) {
	if o == nil || isNil(o.StateJurisdictionCode) {
    return nil, false
	}
	return o.StateJurisdictionCode, true
}

// HasStateJurisdictionCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasStateJurisdictionCode() bool {
	if o != nil && !isNil(o.StateJurisdictionCode) {
		return true
	}

	return false
}

// SetStateJurisdictionCode gets a reference to the given string and assigns it to the StateJurisdictionCode field.
func (o *ValidateResponseKycResult) SetStateJurisdictionCode(v string) {
	o.StateJurisdictionCode = &v
}

// GetTaxpayerType returns the TaxpayerType field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetTaxpayerType() string {
	if o == nil || isNil(o.TaxpayerType) {
		var ret string
		return ret
	}
	return *o.TaxpayerType
}

// GetTaxpayerTypeOk returns a tuple with the TaxpayerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetTaxpayerTypeOk() (*string, bool) {
	if o == nil || isNil(o.TaxpayerType) {
    return nil, false
	}
	return o.TaxpayerType, true
}

// HasTaxpayerType returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasTaxpayerType() bool {
	if o != nil && !isNil(o.TaxpayerType) {
		return true
	}

	return false
}

// SetTaxpayerType gets a reference to the given string and assigns it to the TaxpayerType field.
func (o *ValidateResponseKycResult) SetTaxpayerType(v string) {
	o.TaxpayerType = &v
}

// GetConstitutionOfBusiness returns the ConstitutionOfBusiness field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetConstitutionOfBusiness() string {
	if o == nil || isNil(o.ConstitutionOfBusiness) {
		var ret string
		return ret
	}
	return *o.ConstitutionOfBusiness
}

// GetConstitutionOfBusinessOk returns a tuple with the ConstitutionOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetConstitutionOfBusinessOk() (*string, bool) {
	if o == nil || isNil(o.ConstitutionOfBusiness) {
    return nil, false
	}
	return o.ConstitutionOfBusiness, true
}

// HasConstitutionOfBusiness returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasConstitutionOfBusiness() bool {
	if o != nil && !isNil(o.ConstitutionOfBusiness) {
		return true
	}

	return false
}

// SetConstitutionOfBusiness gets a reference to the given string and assigns it to the ConstitutionOfBusiness field.
func (o *ValidateResponseKycResult) SetConstitutionOfBusiness(v string) {
	o.ConstitutionOfBusiness = &v
}

// GetGstnStatus returns the GstnStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGstnStatus() string {
	if o == nil || isNil(o.GstnStatus) {
		var ret string
		return ret
	}
	return *o.GstnStatus
}

// GetGstnStatusOk returns a tuple with the GstnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGstnStatusOk() (*string, bool) {
	if o == nil || isNil(o.GstnStatus) {
    return nil, false
	}
	return o.GstnStatus, true
}

// HasGstnStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGstnStatus() bool {
	if o != nil && !isNil(o.GstnStatus) {
		return true
	}

	return false
}

// SetGstnStatus gets a reference to the given string and assigns it to the GstnStatus field.
func (o *ValidateResponseKycResult) SetGstnStatus(v string) {
	o.GstnStatus = &v
}

// GetTradeName returns the TradeName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetTradeName() string {
	if o == nil || isNil(o.TradeName) {
		var ret string
		return ret
	}
	return *o.TradeName
}

// GetTradeNameOk returns a tuple with the TradeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetTradeNameOk() (*string, bool) {
	if o == nil || isNil(o.TradeName) {
    return nil, false
	}
	return o.TradeName, true
}

// HasTradeName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasTradeName() bool {
	if o != nil && !isNil(o.TradeName) {
		return true
	}

	return false
}

// SetTradeName gets a reference to the given string and assigns it to the TradeName field.
func (o *ValidateResponseKycResult) SetTradeName(v string) {
	o.TradeName = &v
}

// GetGstin returns the Gstin field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGstin() string {
	if o == nil || isNil(o.Gstin) {
		var ret string
		return ret
	}
	return *o.Gstin
}

// GetGstinOk returns a tuple with the Gstin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGstinOk() (*string, bool) {
	if o == nil || isNil(o.Gstin) {
    return nil, false
	}
	return o.Gstin, true
}

// HasGstin returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGstin() bool {
	if o != nil && !isNil(o.Gstin) {
		return true
	}

	return false
}

// SetGstin gets a reference to the given string and assigns it to the Gstin field.
func (o *ValidateResponseKycResult) SetGstin(v string) {
	o.Gstin = &v
}

// GetAdditionalPlacesOfBusinessInState returns the AdditionalPlacesOfBusinessInState field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAdditionalPlacesOfBusinessInState() []string {
	if o == nil || isNil(o.AdditionalPlacesOfBusinessInState) {
		var ret []string
		return ret
	}
	return o.AdditionalPlacesOfBusinessInState
}

// GetAdditionalPlacesOfBusinessInStateOk returns a tuple with the AdditionalPlacesOfBusinessInState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAdditionalPlacesOfBusinessInStateOk() ([]string, bool) {
	if o == nil || isNil(o.AdditionalPlacesOfBusinessInState) {
    return nil, false
	}
	return o.AdditionalPlacesOfBusinessInState, true
}

// HasAdditionalPlacesOfBusinessInState returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAdditionalPlacesOfBusinessInState() bool {
	if o != nil && !isNil(o.AdditionalPlacesOfBusinessInState) {
		return true
	}

	return false
}

// SetAdditionalPlacesOfBusinessInState gets a reference to the given []string and assigns it to the AdditionalPlacesOfBusinessInState field.
func (o *ValidateResponseKycResult) SetAdditionalPlacesOfBusinessInState(v []string) {
	o.AdditionalPlacesOfBusinessInState = v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLegalName() string {
	if o == nil || isNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLegalNameOk() (*string, bool) {
	if o == nil || isNil(o.LegalName) {
    return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLegalName() bool {
	if o != nil && !isNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *ValidateResponseKycResult) SetLegalName(v string) {
	o.LegalName = &v
}

// GetNatureOfBusiness returns the NatureOfBusiness field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetNatureOfBusiness() []string {
	if o == nil || isNil(o.NatureOfBusiness) {
		var ret []string
		return ret
	}
	return o.NatureOfBusiness
}

// GetNatureOfBusinessOk returns a tuple with the NatureOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetNatureOfBusinessOk() ([]string, bool) {
	if o == nil || isNil(o.NatureOfBusiness) {
    return nil, false
	}
	return o.NatureOfBusiness, true
}

// HasNatureOfBusiness returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasNatureOfBusiness() bool {
	if o != nil && !isNil(o.NatureOfBusiness) {
		return true
	}

	return false
}

// SetNatureOfBusiness gets a reference to the given []string and assigns it to the NatureOfBusiness field.
func (o *ValidateResponseKycResult) SetNatureOfBusiness(v []string) {
	o.NatureOfBusiness = v
}

// GetCentralJurisdiction returns the CentralJurisdiction field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCentralJurisdiction() string {
	if o == nil || isNil(o.CentralJurisdiction) {
		var ret string
		return ret
	}
	return *o.CentralJurisdiction
}

// GetCentralJurisdictionOk returns a tuple with the CentralJurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCentralJurisdictionOk() (*string, bool) {
	if o == nil || isNil(o.CentralJurisdiction) {
    return nil, false
	}
	return o.CentralJurisdiction, true
}

// HasCentralJurisdiction returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCentralJurisdiction() bool {
	if o != nil && !isNil(o.CentralJurisdiction) {
		return true
	}

	return false
}

// SetCentralJurisdiction gets a reference to the given string and assigns it to the CentralJurisdiction field.
func (o *ValidateResponseKycResult) SetCentralJurisdiction(v string) {
	o.CentralJurisdiction = &v
}

// GetCentralJurisdictionCode returns the CentralJurisdictionCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCentralJurisdictionCode() string {
	if o == nil || isNil(o.CentralJurisdictionCode) {
		var ret string
		return ret
	}
	return *o.CentralJurisdictionCode
}

// GetCentralJurisdictionCodeOk returns a tuple with the CentralJurisdictionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCentralJurisdictionCodeOk() (*string, bool) {
	if o == nil || isNil(o.CentralJurisdictionCode) {
    return nil, false
	}
	return o.CentralJurisdictionCode, true
}

// HasCentralJurisdictionCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCentralJurisdictionCode() bool {
	if o != nil && !isNil(o.CentralJurisdictionCode) {
		return true
	}

	return false
}

// SetCentralJurisdictionCode gets a reference to the given string and assigns it to the CentralJurisdictionCode field.
func (o *ValidateResponseKycResult) SetCentralJurisdictionCode(v string) {
	o.CentralJurisdictionCode = &v
}

// GetPan returns the Pan field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetPan() string {
	if o == nil || isNil(o.Pan) {
		var ret string
		return ret
	}
	return *o.Pan
}

// GetPanOk returns a tuple with the Pan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetPanOk() (*string, bool) {
	if o == nil || isNil(o.Pan) {
    return nil, false
	}
	return o.Pan, true
}

// HasPan returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasPan() bool {
	if o != nil && !isNil(o.Pan) {
		return true
	}

	return false
}

// SetPan gets a reference to the given string and assigns it to the Pan field.
func (o *ValidateResponseKycResult) SetPan(v string) {
	o.Pan = &v
}

// GetAuthorizedSignatories returns the AuthorizedSignatories field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAuthorizedSignatories() string {
	if o == nil || isNil(o.AuthorizedSignatories) {
		var ret string
		return ret
	}
	return *o.AuthorizedSignatories
}

// GetAuthorizedSignatoriesOk returns a tuple with the AuthorizedSignatories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAuthorizedSignatoriesOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedSignatories) {
    return nil, false
	}
	return o.AuthorizedSignatories, true
}

// HasAuthorizedSignatories returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAuthorizedSignatories() bool {
	if o != nil && !isNil(o.AuthorizedSignatories) {
		return true
	}

	return false
}

// SetAuthorizedSignatories gets a reference to the given string and assigns it to the AuthorizedSignatories field.
func (o *ValidateResponseKycResult) SetAuthorizedSignatories(v string) {
	o.AuthorizedSignatories = &v
}

// GetComplianceRating returns the ComplianceRating field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetComplianceRating() string {
	if o == nil || isNil(o.ComplianceRating) {
		var ret string
		return ret
	}
	return *o.ComplianceRating
}

// GetComplianceRatingOk returns a tuple with the ComplianceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetComplianceRatingOk() (*string, bool) {
	if o == nil || isNil(o.ComplianceRating) {
    return nil, false
	}
	return o.ComplianceRating, true
}

// HasComplianceRating returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasComplianceRating() bool {
	if o != nil && !isNil(o.ComplianceRating) {
		return true
	}

	return false
}

// SetComplianceRating gets a reference to the given string and assigns it to the ComplianceRating field.
func (o *ValidateResponseKycResult) SetComplianceRating(v string) {
	o.ComplianceRating = &v
}

// GetCxdt returns the Cxdt field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCxdt() string {
	if o == nil || isNil(o.Cxdt) {
		var ret string
		return ret
	}
	return *o.Cxdt
}

// GetCxdtOk returns a tuple with the Cxdt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCxdtOk() (*string, bool) {
	if o == nil || isNil(o.Cxdt) {
    return nil, false
	}
	return o.Cxdt, true
}

// HasCxdt returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCxdt() bool {
	if o != nil && !isNil(o.Cxdt) {
		return true
	}

	return false
}

// SetCxdt gets a reference to the given string and assigns it to the Cxdt field.
func (o *ValidateResponseKycResult) SetCxdt(v string) {
	o.Cxdt = &v
}

// GetBusinessDetails returns the BusinessDetails field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetBusinessDetails() []ValidateResponseKycResultBusinessDetailsInner {
	if o == nil || isNil(o.BusinessDetails) {
		var ret []ValidateResponseKycResultBusinessDetailsInner
		return ret
	}
	return o.BusinessDetails
}

// GetBusinessDetailsOk returns a tuple with the BusinessDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetBusinessDetailsOk() ([]ValidateResponseKycResultBusinessDetailsInner, bool) {
	if o == nil || isNil(o.BusinessDetails) {
    return nil, false
	}
	return o.BusinessDetails, true
}

// HasBusinessDetails returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasBusinessDetails() bool {
	if o != nil && !isNil(o.BusinessDetails) {
		return true
	}

	return false
}

// SetBusinessDetails gets a reference to the given []ValidateResponseKycResultBusinessDetailsInner and assigns it to the BusinessDetails field.
func (o *ValidateResponseKycResult) SetBusinessDetails(v []ValidateResponseKycResultBusinessDetailsInner) {
	o.BusinessDetails = v
}

// GetAnnualAggregateTurnover returns the AnnualAggregateTurnover field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetAnnualAggregateTurnover() string {
	if o == nil || isNil(o.AnnualAggregateTurnover) {
		var ret string
		return ret
	}
	return *o.AnnualAggregateTurnover
}

// GetAnnualAggregateTurnoverOk returns a tuple with the AnnualAggregateTurnover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetAnnualAggregateTurnoverOk() (*string, bool) {
	if o == nil || isNil(o.AnnualAggregateTurnover) {
    return nil, false
	}
	return o.AnnualAggregateTurnover, true
}

// HasAnnualAggregateTurnover returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasAnnualAggregateTurnover() bool {
	if o != nil && !isNil(o.AnnualAggregateTurnover) {
		return true
	}

	return false
}

// SetAnnualAggregateTurnover gets a reference to the given string and assigns it to the AnnualAggregateTurnover field.
func (o *ValidateResponseKycResult) SetAnnualAggregateTurnover(v string) {
	o.AnnualAggregateTurnover = &v
}

// GetMandatoryEInvoicing returns the MandatoryEInvoicing field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetMandatoryEInvoicing() string {
	if o == nil || isNil(o.MandatoryEInvoicing) {
		var ret string
		return ret
	}
	return *o.MandatoryEInvoicing
}

// GetMandatoryEInvoicingOk returns a tuple with the MandatoryEInvoicing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetMandatoryEInvoicingOk() (*string, bool) {
	if o == nil || isNil(o.MandatoryEInvoicing) {
    return nil, false
	}
	return o.MandatoryEInvoicing, true
}

// HasMandatoryEInvoicing returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasMandatoryEInvoicing() bool {
	if o != nil && !isNil(o.MandatoryEInvoicing) {
		return true
	}

	return false
}

// SetMandatoryEInvoicing gets a reference to the given string and assigns it to the MandatoryEInvoicing field.
func (o *ValidateResponseKycResult) SetMandatoryEInvoicing(v string) {
	o.MandatoryEInvoicing = &v
}

// GetGrossTotalIncome returns the GrossTotalIncome field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGrossTotalIncome() string {
	if o == nil || isNil(o.GrossTotalIncome) {
		var ret string
		return ret
	}
	return *o.GrossTotalIncome
}

// GetGrossTotalIncomeOk returns a tuple with the GrossTotalIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGrossTotalIncomeOk() (*string, bool) {
	if o == nil || isNil(o.GrossTotalIncome) {
    return nil, false
	}
	return o.GrossTotalIncome, true
}

// HasGrossTotalIncome returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGrossTotalIncome() bool {
	if o != nil && !isNil(o.GrossTotalIncome) {
		return true
	}

	return false
}

// SetGrossTotalIncome gets a reference to the given string and assigns it to the GrossTotalIncome field.
func (o *ValidateResponseKycResult) SetGrossTotalIncome(v string) {
	o.GrossTotalIncome = &v
}

// GetGrossTotalIncomeFinancialYear returns the GrossTotalIncomeFinancialYear field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetGrossTotalIncomeFinancialYear() string {
	if o == nil || isNil(o.GrossTotalIncomeFinancialYear) {
		var ret string
		return ret
	}
	return *o.GrossTotalIncomeFinancialYear
}

// GetGrossTotalIncomeFinancialYearOk returns a tuple with the GrossTotalIncomeFinancialYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetGrossTotalIncomeFinancialYearOk() (*string, bool) {
	if o == nil || isNil(o.GrossTotalIncomeFinancialYear) {
    return nil, false
	}
	return o.GrossTotalIncomeFinancialYear, true
}

// HasGrossTotalIncomeFinancialYear returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasGrossTotalIncomeFinancialYear() bool {
	if o != nil && !isNil(o.GrossTotalIncomeFinancialYear) {
		return true
	}

	return false
}

// SetGrossTotalIncomeFinancialYear gets a reference to the given string and assigns it to the GrossTotalIncomeFinancialYear field.
func (o *ValidateResponseKycResult) SetGrossTotalIncomeFinancialYear(v string) {
	o.GrossTotalIncomeFinancialYear = &v
}

// GetIsFieldVisitConducted returns the IsFieldVisitConducted field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetIsFieldVisitConducted() string {
	if o == nil || isNil(o.IsFieldVisitConducted) {
		var ret string
		return ret
	}
	return *o.IsFieldVisitConducted
}

// GetIsFieldVisitConductedOk returns a tuple with the IsFieldVisitConducted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetIsFieldVisitConductedOk() (*string, bool) {
	if o == nil || isNil(o.IsFieldVisitConducted) {
    return nil, false
	}
	return o.IsFieldVisitConducted, true
}

// HasIsFieldVisitConducted returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasIsFieldVisitConducted() bool {
	if o != nil && !isNil(o.IsFieldVisitConducted) {
		return true
	}

	return false
}

// SetIsFieldVisitConducted gets a reference to the given string and assigns it to the IsFieldVisitConducted field.
func (o *ValidateResponseKycResult) SetIsFieldVisitConducted(v string) {
	o.IsFieldVisitConducted = &v
}

// GetFilingStatus returns the FilingStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetFilingStatus() []ValidateResponseKycResultFilingStatusInner {
	if o == nil || isNil(o.FilingStatus) {
		var ret []ValidateResponseKycResultFilingStatusInner
		return ret
	}
	return o.FilingStatus
}

// GetFilingStatusOk returns a tuple with the FilingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetFilingStatusOk() ([]ValidateResponseKycResultFilingStatusInner, bool) {
	if o == nil || isNil(o.FilingStatus) {
    return nil, false
	}
	return o.FilingStatus, true
}

// HasFilingStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasFilingStatus() bool {
	if o != nil && !isNil(o.FilingStatus) {
		return true
	}

	return false
}

// SetFilingStatus gets a reference to the given []ValidateResponseKycResultFilingStatusInner and assigns it to the FilingStatus field.
func (o *ValidateResponseKycResult) SetFilingStatus(v []ValidateResponseKycResultFilingStatusInner) {
	o.FilingStatus = v
}

// GetDirectors returns the Directors field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDirectors() ValidateResponseKycResultDirectors {
	if o == nil || isNil(o.Directors) {
		var ret ValidateResponseKycResultDirectors
		return ret
	}
	return *o.Directors
}

// GetDirectorsOk returns a tuple with the Directors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDirectorsOk() (*ValidateResponseKycResultDirectors, bool) {
	if o == nil || isNil(o.Directors) {
    return nil, false
	}
	return o.Directors, true
}

// HasDirectors returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDirectors() bool {
	if o != nil && !isNil(o.Directors) {
		return true
	}

	return false
}

// SetDirectors gets a reference to the given ValidateResponseKycResultDirectors and assigns it to the Directors field.
func (o *ValidateResponseKycResult) SetDirectors(v ValidateResponseKycResultDirectors) {
	o.Directors = &v
}

// GetCompanyMasterData returns the CompanyMasterData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCompanyMasterData() ValidateResponseKycResultCompanyMasterData {
	if o == nil || isNil(o.CompanyMasterData) {
		var ret ValidateResponseKycResultCompanyMasterData
		return ret
	}
	return *o.CompanyMasterData
}

// GetCompanyMasterDataOk returns a tuple with the CompanyMasterData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCompanyMasterDataOk() (*ValidateResponseKycResultCompanyMasterData, bool) {
	if o == nil || isNil(o.CompanyMasterData) {
    return nil, false
	}
	return o.CompanyMasterData, true
}

// HasCompanyMasterData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCompanyMasterData() bool {
	if o != nil && !isNil(o.CompanyMasterData) {
		return true
	}

	return false
}

// SetCompanyMasterData gets a reference to the given ValidateResponseKycResultCompanyMasterData and assigns it to the CompanyMasterData field.
func (o *ValidateResponseKycResult) SetCompanyMasterData(v ValidateResponseKycResultCompanyMasterData) {
	o.CompanyMasterData = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCharges() []string {
	if o == nil || isNil(o.Charges) {
		var ret []string
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetChargesOk() ([]string, bool) {
	if o == nil || isNil(o.Charges) {
    return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCharges() bool {
	if o != nil && !isNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []string and assigns it to the Charges field.
func (o *ValidateResponseKycResult) SetCharges(v []string) {
	o.Charges = v
}

// GetLlpData returns the LlpData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLlpData() []string {
	if o == nil || isNil(o.LlpData) {
		var ret []string
		return ret
	}
	return o.LlpData
}

// GetLlpDataOk returns a tuple with the LlpData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLlpDataOk() ([]string, bool) {
	if o == nil || isNil(o.LlpData) {
    return nil, false
	}
	return o.LlpData, true
}

// HasLlpData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLlpData() bool {
	if o != nil && !isNil(o.LlpData) {
		return true
	}

	return false
}

// SetLlpData gets a reference to the given []string and assigns it to the LlpData field.
func (o *ValidateResponseKycResult) SetLlpData(v []string) {
	o.LlpData = v
}

// GetCompanyData returns the CompanyData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetCompanyData() []string {
	if o == nil || isNil(o.CompanyData) {
		var ret []string
		return ret
	}
	return o.CompanyData
}

// GetCompanyDataOk returns a tuple with the CompanyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetCompanyDataOk() ([]string, bool) {
	if o == nil || isNil(o.CompanyData) {
    return nil, false
	}
	return o.CompanyData, true
}

// HasCompanyData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasCompanyData() bool {
	if o != nil && !isNil(o.CompanyData) {
		return true
	}

	return false
}

// SetCompanyData gets a reference to the given []string and assigns it to the CompanyData field.
func (o *ValidateResponseKycResult) SetCompanyData(v []string) {
	o.CompanyData = v
}

// GetDirectorData returns the DirectorData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetDirectorData() ValidateResponseKycResultDirectorData {
	if o == nil || isNil(o.DirectorData) {
		var ret ValidateResponseKycResultDirectorData
		return ret
	}
	return *o.DirectorData
}

// GetDirectorDataOk returns a tuple with the DirectorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetDirectorDataOk() (*ValidateResponseKycResultDirectorData, bool) {
	if o == nil || isNil(o.DirectorData) {
    return nil, false
	}
	return o.DirectorData, true
}

// HasDirectorData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasDirectorData() bool {
	if o != nil && !isNil(o.DirectorData) {
		return true
	}

	return false
}

// SetDirectorData gets a reference to the given ValidateResponseKycResultDirectorData and assigns it to the DirectorData field.
func (o *ValidateResponseKycResult) SetDirectorData(v ValidateResponseKycResultDirectorData) {
	o.DirectorData = &v
}

// GetLlpMasterData returns the LlpMasterData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetLlpMasterData() ValidateResponseKycResultLlpMasterData {
	if o == nil || isNil(o.LlpMasterData) {
		var ret ValidateResponseKycResultLlpMasterData
		return ret
	}
	return *o.LlpMasterData
}

// GetLlpMasterDataOk returns a tuple with the LlpMasterData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetLlpMasterDataOk() (*ValidateResponseKycResultLlpMasterData, bool) {
	if o == nil || isNil(o.LlpMasterData) {
    return nil, false
	}
	return o.LlpMasterData, true
}

// HasLlpMasterData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasLlpMasterData() bool {
	if o != nil && !isNil(o.LlpMasterData) {
		return true
	}

	return false
}

// SetLlpMasterData gets a reference to the given ValidateResponseKycResultLlpMasterData and assigns it to the LlpMasterData field.
func (o *ValidateResponseKycResult) SetLlpMasterData(v ValidateResponseKycResultLlpMasterData) {
	o.LlpMasterData = &v
}

// GetForeignCompanyMasterData returns the ForeignCompanyMasterData field value if set, zero value otherwise.
func (o *ValidateResponseKycResult) GetForeignCompanyMasterData() ValidateResponseKycResultForeignCompanyMasterData {
	if o == nil || isNil(o.ForeignCompanyMasterData) {
		var ret ValidateResponseKycResultForeignCompanyMasterData
		return ret
	}
	return *o.ForeignCompanyMasterData
}

// GetForeignCompanyMasterDataOk returns a tuple with the ForeignCompanyMasterData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResult) GetForeignCompanyMasterDataOk() (*ValidateResponseKycResultForeignCompanyMasterData, bool) {
	if o == nil || isNil(o.ForeignCompanyMasterData) {
    return nil, false
	}
	return o.ForeignCompanyMasterData, true
}

// HasForeignCompanyMasterData returns a boolean if a field has been set.
func (o *ValidateResponseKycResult) HasForeignCompanyMasterData() bool {
	if o != nil && !isNil(o.ForeignCompanyMasterData) {
		return true
	}

	return false
}

// SetForeignCompanyMasterData gets a reference to the given ValidateResponseKycResultForeignCompanyMasterData and assigns it to the ForeignCompanyMasterData field.
func (o *ValidateResponseKycResult) SetForeignCompanyMasterData(v ValidateResponseKycResultForeignCompanyMasterData) {
	o.ForeignCompanyMasterData = &v
}

func (o ValidateResponseKycResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdNumber) {
		toSerialize["idNumber"] = o.IdNumber
	}
	if !isNil(o.IdStatus) {
		toSerialize["idStatus"] = o.IdStatus
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !isNil(o.EntityName) {
		toSerialize["entityName"] = o.EntityName
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.PremissesAddress) {
		toSerialize["premissesAddress"] = o.PremissesAddress
	}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.UamNumber) {
		toSerialize["uamNumber"] = o.UamNumber
	}
	if !isNil(o.EnterpriseName) {
		toSerialize["enterpriseName"] = o.EnterpriseName
	}
	if !isNil(o.MajorActivity) {
		toSerialize["majorActivity"] = o.MajorActivity
	}
	if !isNil(o.SocialCategory) {
		toSerialize["socialCategory"] = o.SocialCategory
	}
	if !isNil(o.EnterpriseType) {
		toSerialize["enterpriseType"] = o.EnterpriseType
	}
	if !isNil(o.DateOfCommencement) {
		toSerialize["dateOfCommencement"] = o.DateOfCommencement
	}
	if !isNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.AppliedDate) {
		toSerialize["appliedDate"] = o.AppliedDate
	}
	if !isNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !isNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !isNil(o.Nic2Digit) {
		toSerialize["nic_2Digit"] = o.Nic2Digit
	}
	if !isNil(o.Nic4Digit) {
		toSerialize["nic_4Digit"] = o.Nic4Digit
	}
	if !isNil(o.Nic5Digit) {
		toSerialize["nic_5Digit"] = o.Nic5Digit
	}
	if !isNil(o.PanStatus) {
		toSerialize["panStatus"] = o.PanStatus
	}
	if !isNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !isNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !isNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !isNil(o.IdHolderTitle) {
		toSerialize["idHolderTitle"] = o.IdHolderTitle
	}
	if !isNil(o.IdLastUpdated) {
		toSerialize["idLastUpdated"] = o.IdLastUpdated
	}
	if !isNil(o.AadhaarSeedingStatus) {
		toSerialize["aadhaarSeedingStatus"] = o.AadhaarSeedingStatus
	}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !isNil(o.AllClassOfVehicle) {
		toSerialize["allClassOfVehicle"] = o.AllClassOfVehicle
	}
	if !isNil(o.DrivingLicenseNumber) {
		toSerialize["drivingLicenseNumber"] = o.DrivingLicenseNumber
	}
	if !isNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !isNil(o.EndorseDate) {
		toSerialize["endorseDate"] = o.EndorseDate
	}
	if !isNil(o.EndorseNumber) {
		toSerialize["endorseNumber"] = o.EndorseNumber
	}
	if !isNil(o.FatherOrHusbandName) {
		toSerialize["fatherOrHusbandName"] = o.FatherOrHusbandName
	}
	if !isNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !isNil(o.ValidTo) {
		toSerialize["validTo"] = o.ValidTo
	}
	if !isNil(o.EpicNo) {
		toSerialize["epicNo"] = o.EpicNo
	}
	if !isNil(o.NameInVernacular) {
		toSerialize["nameInVernacular"] = o.NameInVernacular
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !isNil(o.RelativeName) {
		toSerialize["relativeName"] = o.RelativeName
	}
	if !isNil(o.RelativeNameInVernacular) {
		toSerialize["relativeNameInVernacular"] = o.RelativeNameInVernacular
	}
	if !isNil(o.RelativeRelationType) {
		toSerialize["relativeRelationType"] = o.RelativeRelationType
	}
	if !isNil(o.HouseNumber) {
		toSerialize["houseNumber"] = o.HouseNumber
	}
	if !isNil(o.PartOrLocationInConstituency) {
		toSerialize["partOrLocationInConstituency"] = o.PartOrLocationInConstituency
	}
	if !isNil(o.PartNumberOrLocationNumberInConstituency) {
		toSerialize["partNumberOrLocationNumberInConstituency"] = o.PartNumberOrLocationNumberInConstituency
	}
	if !isNil(o.ParliamentaryConstituency) {
		toSerialize["parliamentaryConstituency"] = o.ParliamentaryConstituency
	}
	if !isNil(o.AssemblyConstituency) {
		toSerialize["assemblyConstituency"] = o.AssemblyConstituency
	}
	if !isNil(o.SectionOfConstituencyPart) {
		toSerialize["sectionOfConstituencyPart"] = o.SectionOfConstituencyPart
	}
	if !isNil(o.CardSerialNumberInPollingList) {
		toSerialize["cardSerialNumberInPollingList"] = o.CardSerialNumberInPollingList
	}
	if !isNil(o.LastUpdateDate) {
		toSerialize["lastUpdateDate"] = o.LastUpdateDate
	}
	if !isNil(o.PollingBoothDetails) {
		toSerialize["pollingBoothDetails"] = o.PollingBoothDetails
	}
	if !isNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !isNil(o.MobileNumber) {
		toSerialize["mobileNumber"] = o.MobileNumber
	}
	if !isNil(o.StateCode) {
		toSerialize["stateCode"] = o.StateCode
	}
	if !isNil(o.PollingBoothCoordinates) {
		toSerialize["pollingBoothCoordinates"] = o.PollingBoothCoordinates
	}
	if !isNil(o.PollingBoothAddress) {
		toSerialize["pollingBoothAddress"] = o.PollingBoothAddress
	}
	if !isNil(o.PollingBoothNumber) {
		toSerialize["pollingBoothNumber"] = o.PollingBoothNumber
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BlacklistStatus) {
		toSerialize["blacklistStatus"] = o.BlacklistStatus
	}
	if !isNil(o.RegistrationDate) {
		toSerialize["registrationDate"] = o.RegistrationDate
	}
	if !isNil(o.RegistrationLocation) {
		toSerialize["registrationLocation"] = o.RegistrationLocation
	}
	if !isNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !isNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !isNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !isNil(o.ChassisNumber) {
		toSerialize["chassisNumber"] = o.ChassisNumber
	}
	if !isNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !isNil(o.EngineNumber) {
		toSerialize["engineNumber"] = o.EngineNumber
	}
	if !isNil(o.FuelType) {
		toSerialize["fuelType"] = o.FuelType
	}
	if !isNil(o.FitUpto) {
		toSerialize["fitUpto"] = o.FitUpto
	}
	if !isNil(o.InsuranceUpto) {
		toSerialize["insuranceUpto"] = o.InsuranceUpto
	}
	if !isNil(o.TaxUpto) {
		toSerialize["taxUpto"] = o.TaxUpto
	}
	if !isNil(o.InsuranceDetails) {
		toSerialize["insuranceDetails"] = o.InsuranceDetails
	}
	if !isNil(o.InsuranceValidity) {
		toSerialize["insuranceValidity"] = o.InsuranceValidity
	}
	if !isNil(o.PermitType) {
		toSerialize["permitType"] = o.PermitType
	}
	if !isNil(o.PermitValidUpto) {
		toSerialize["permitValidUpto"] = o.PermitValidUpto
	}
	if !isNil(o.PollutionControlValidity) {
		toSerialize["pollutionControlValidity"] = o.PollutionControlValidity
	}
	if !isNil(o.PollutionNorms) {
		toSerialize["pollutionNorms"] = o.PollutionNorms
	}
	if !isNil(o.LicenseAddress) {
		toSerialize["licenseAddress"] = o.LicenseAddress
	}
	if !isNil(o.RegistrationAddress) {
		toSerialize["registrationAddress"] = o.RegistrationAddress
	}
	if !isNil(o.OwnerFatherName) {
		toSerialize["ownerFatherName"] = o.OwnerFatherName
	}
	if !isNil(o.OwnerPresentAddress) {
		toSerialize["ownerPresentAddress"] = o.OwnerPresentAddress
	}
	if !isNil(o.BodyType) {
		toSerialize["bodyType"] = o.BodyType
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !isNil(o.EngineCubicCapacity) {
		toSerialize["engineCubicCapacity"] = o.EngineCubicCapacity
	}
	if !isNil(o.NumberCylinders) {
		toSerialize["numberCylinders"] = o.NumberCylinders
	}
	if !isNil(o.UnladenWeight) {
		toSerialize["unladenWeight"] = o.UnladenWeight
	}
	if !isNil(o.GrossWeight) {
		toSerialize["grossWeight"] = o.GrossWeight
	}
	if !isNil(o.Wheelbase) {
		toSerialize["wheelbase"] = o.Wheelbase
	}
	if !isNil(o.ManufacturedMonthYear) {
		toSerialize["manufacturedMonthYear"] = o.ManufacturedMonthYear
	}
	if !isNil(o.MakerDescription) {
		toSerialize["makerDescription"] = o.MakerDescription
	}
	if !isNil(o.NocDetails) {
		toSerialize["nocDetails"] = o.NocDetails
	}
	if !isNil(o.NormsDescription) {
		toSerialize["normsDescription"] = o.NormsDescription
	}
	if !isNil(o.Financier) {
		toSerialize["financier"] = o.Financier
	}
	if !isNil(o.PermitIssueDate) {
		toSerialize["permitIssueDate"] = o.PermitIssueDate
	}
	if !isNil(o.PermitNumber) {
		toSerialize["permitNumber"] = o.PermitNumber
	}
	if !isNil(o.PermitValidFrom) {
		toSerialize["permitValidFrom"] = o.PermitValidFrom
	}
	if !isNil(o.SeatingCapacity) {
		toSerialize["seatingCapacity"] = o.SeatingCapacity
	}
	if !isNil(o.SleepingCapacity) {
		toSerialize["sleepingCapacity"] = o.SleepingCapacity
	}
	if !isNil(o.StandingCapacity) {
		toSerialize["standingCapacity"] = o.StandingCapacity
	}
	if !isNil(o.StatusAsOn) {
		toSerialize["statusAsOn"] = o.StatusAsOn
	}
	if !isNil(o.PrimaryBusinessContact) {
		toSerialize["primaryBusinessContact"] = o.PrimaryBusinessContact
	}
	if !isNil(o.StateJurisdiction) {
		toSerialize["stateJurisdiction"] = o.StateJurisdiction
	}
	if !isNil(o.StateJurisdictionCode) {
		toSerialize["stateJurisdictionCode"] = o.StateJurisdictionCode
	}
	if !isNil(o.TaxpayerType) {
		toSerialize["taxpayerType"] = o.TaxpayerType
	}
	if !isNil(o.ConstitutionOfBusiness) {
		toSerialize["constitutionOfBusiness"] = o.ConstitutionOfBusiness
	}
	if !isNil(o.GstnStatus) {
		toSerialize["gstnStatus"] = o.GstnStatus
	}
	if !isNil(o.TradeName) {
		toSerialize["tradeName"] = o.TradeName
	}
	if !isNil(o.Gstin) {
		toSerialize["gstin"] = o.Gstin
	}
	if !isNil(o.AdditionalPlacesOfBusinessInState) {
		toSerialize["additionalPlacesOfBusinessInState"] = o.AdditionalPlacesOfBusinessInState
	}
	if !isNil(o.LegalName) {
		toSerialize["legalName"] = o.LegalName
	}
	if !isNil(o.NatureOfBusiness) {
		toSerialize["natureOfBusiness"] = o.NatureOfBusiness
	}
	if !isNil(o.CentralJurisdiction) {
		toSerialize["centralJurisdiction"] = o.CentralJurisdiction
	}
	if !isNil(o.CentralJurisdictionCode) {
		toSerialize["centralJurisdictionCode"] = o.CentralJurisdictionCode
	}
	if !isNil(o.Pan) {
		toSerialize["pan"] = o.Pan
	}
	if !isNil(o.AuthorizedSignatories) {
		toSerialize["authorizedSignatories"] = o.AuthorizedSignatories
	}
	if !isNil(o.ComplianceRating) {
		toSerialize["complianceRating"] = o.ComplianceRating
	}
	if !isNil(o.Cxdt) {
		toSerialize["cxdt"] = o.Cxdt
	}
	if !isNil(o.BusinessDetails) {
		toSerialize["businessDetails"] = o.BusinessDetails
	}
	if !isNil(o.AnnualAggregateTurnover) {
		toSerialize["annualAggregateTurnover"] = o.AnnualAggregateTurnover
	}
	if !isNil(o.MandatoryEInvoicing) {
		toSerialize["mandatoryEInvoicing"] = o.MandatoryEInvoicing
	}
	if !isNil(o.GrossTotalIncome) {
		toSerialize["grossTotalIncome"] = o.GrossTotalIncome
	}
	if !isNil(o.GrossTotalIncomeFinancialYear) {
		toSerialize["grossTotalIncomeFinancialYear"] = o.GrossTotalIncomeFinancialYear
	}
	if !isNil(o.IsFieldVisitConducted) {
		toSerialize["isFieldVisitConducted"] = o.IsFieldVisitConducted
	}
	if !isNil(o.FilingStatus) {
		toSerialize["filingStatus"] = o.FilingStatus
	}
	if !isNil(o.Directors) {
		toSerialize["directors"] = o.Directors
	}
	if !isNil(o.CompanyMasterData) {
		toSerialize["companyMasterData"] = o.CompanyMasterData
	}
	if !isNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !isNil(o.LlpData) {
		toSerialize["llpData"] = o.LlpData
	}
	if !isNil(o.CompanyData) {
		toSerialize["companyData"] = o.CompanyData
	}
	if !isNil(o.DirectorData) {
		toSerialize["directorData"] = o.DirectorData
	}
	if !isNil(o.LlpMasterData) {
		toSerialize["llpMasterData"] = o.LlpMasterData
	}
	if !isNil(o.ForeignCompanyMasterData) {
		toSerialize["foreignCompanyMasterData"] = o.ForeignCompanyMasterData
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResult struct {
	value *ValidateResponseKycResult
	isSet bool
}

func (v NullableValidateResponseKycResult) Get() *ValidateResponseKycResult {
	return v.value
}

func (v *NullableValidateResponseKycResult) Set(val *ValidateResponseKycResult) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResult) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResult(val *ValidateResponseKycResult) *NullableValidateResponseKycResult {
	return &NullableValidateResponseKycResult{value: val, isSet: true}
}

func (v NullableValidateResponseKycResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


