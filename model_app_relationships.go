/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
)

// AppRelationships struct for AppRelationships
type AppRelationships struct {
	BetaGroups *AppRelationshipsBetaGroups `json:"betaGroups,omitempty"`
	AppStoreVersions *AppRelationshipsAppStoreVersions `json:"appStoreVersions,omitempty"`
	PreReleaseVersions *AppRelationshipsPreReleaseVersions `json:"preReleaseVersions,omitempty"`
	BetaAppLocalizations *AppRelationshipsBetaAppLocalizations `json:"betaAppLocalizations,omitempty"`
	Builds *AppRelationshipsBuilds `json:"builds,omitempty"`
	BetaLicenseAgreement *AppRelationshipsBetaLicenseAgreement `json:"betaLicenseAgreement,omitempty"`
	BetaAppReviewDetail *AppRelationshipsBetaAppReviewDetail `json:"betaAppReviewDetail,omitempty"`
	AppInfos *AppRelationshipsAppInfos `json:"appInfos,omitempty"`
	EndUserLicenseAgreement *AppRelationshipsEndUserLicenseAgreement `json:"endUserLicenseAgreement,omitempty"`
	PreOrder *AppRelationshipsPreOrder `json:"preOrder,omitempty"`
	Prices *AppRelationshipsPrices `json:"prices,omitempty"`
	AvailableTerritories *AppRelationshipsAvailableTerritories `json:"availableTerritories,omitempty"`
	InAppPurchases *AppRelationshipsInAppPurchases `json:"inAppPurchases,omitempty"`
	GameCenterEnabledVersions *AppRelationshipsGameCenterEnabledVersions `json:"gameCenterEnabledVersions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppRelationships AppRelationships

// NewAppRelationships instantiates a new AppRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationships() *AppRelationships {
	this := AppRelationships{}
	return &this
}

// NewAppRelationshipsWithDefaults instantiates a new AppRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsWithDefaults() *AppRelationships {
	this := AppRelationships{}
	return &this
}

// GetBetaGroups returns the BetaGroups field value if set, zero value otherwise.
func (o *AppRelationships) GetBetaGroups() AppRelationshipsBetaGroups {
	if o == nil || o.BetaGroups == nil {
		var ret AppRelationshipsBetaGroups
		return ret
	}
	return *o.BetaGroups
}

// GetBetaGroupsOk returns a tuple with the BetaGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool) {
	if o == nil || o.BetaGroups == nil {
		return nil, false
	}
	return o.BetaGroups, true
}

// HasBetaGroups returns a boolean if a field has been set.
func (o *AppRelationships) HasBetaGroups() bool {
	if o != nil && o.BetaGroups != nil {
		return true
	}

	return false
}

// SetBetaGroups gets a reference to the given AppRelationshipsBetaGroups and assigns it to the BetaGroups field.
func (o *AppRelationships) SetBetaGroups(v AppRelationshipsBetaGroups) {
	o.BetaGroups = &v
}

// GetAppStoreVersions returns the AppStoreVersions field value if set, zero value otherwise.
func (o *AppRelationships) GetAppStoreVersions() AppRelationshipsAppStoreVersions {
	if o == nil || o.AppStoreVersions == nil {
		var ret AppRelationshipsAppStoreVersions
		return ret
	}
	return *o.AppStoreVersions
}

// GetAppStoreVersionsOk returns a tuple with the AppStoreVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetAppStoreVersionsOk() (*AppRelationshipsAppStoreVersions, bool) {
	if o == nil || o.AppStoreVersions == nil {
		return nil, false
	}
	return o.AppStoreVersions, true
}

// HasAppStoreVersions returns a boolean if a field has been set.
func (o *AppRelationships) HasAppStoreVersions() bool {
	if o != nil && o.AppStoreVersions != nil {
		return true
	}

	return false
}

// SetAppStoreVersions gets a reference to the given AppRelationshipsAppStoreVersions and assigns it to the AppStoreVersions field.
func (o *AppRelationships) SetAppStoreVersions(v AppRelationshipsAppStoreVersions) {
	o.AppStoreVersions = &v
}

// GetPreReleaseVersions returns the PreReleaseVersions field value if set, zero value otherwise.
func (o *AppRelationships) GetPreReleaseVersions() AppRelationshipsPreReleaseVersions {
	if o == nil || o.PreReleaseVersions == nil {
		var ret AppRelationshipsPreReleaseVersions
		return ret
	}
	return *o.PreReleaseVersions
}

// GetPreReleaseVersionsOk returns a tuple with the PreReleaseVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetPreReleaseVersionsOk() (*AppRelationshipsPreReleaseVersions, bool) {
	if o == nil || o.PreReleaseVersions == nil {
		return nil, false
	}
	return o.PreReleaseVersions, true
}

// HasPreReleaseVersions returns a boolean if a field has been set.
func (o *AppRelationships) HasPreReleaseVersions() bool {
	if o != nil && o.PreReleaseVersions != nil {
		return true
	}

	return false
}

// SetPreReleaseVersions gets a reference to the given AppRelationshipsPreReleaseVersions and assigns it to the PreReleaseVersions field.
func (o *AppRelationships) SetPreReleaseVersions(v AppRelationshipsPreReleaseVersions) {
	o.PreReleaseVersions = &v
}

// GetBetaAppLocalizations returns the BetaAppLocalizations field value if set, zero value otherwise.
func (o *AppRelationships) GetBetaAppLocalizations() AppRelationshipsBetaAppLocalizations {
	if o == nil || o.BetaAppLocalizations == nil {
		var ret AppRelationshipsBetaAppLocalizations
		return ret
	}
	return *o.BetaAppLocalizations
}

// GetBetaAppLocalizationsOk returns a tuple with the BetaAppLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetBetaAppLocalizationsOk() (*AppRelationshipsBetaAppLocalizations, bool) {
	if o == nil || o.BetaAppLocalizations == nil {
		return nil, false
	}
	return o.BetaAppLocalizations, true
}

// HasBetaAppLocalizations returns a boolean if a field has been set.
func (o *AppRelationships) HasBetaAppLocalizations() bool {
	if o != nil && o.BetaAppLocalizations != nil {
		return true
	}

	return false
}

// SetBetaAppLocalizations gets a reference to the given AppRelationshipsBetaAppLocalizations and assigns it to the BetaAppLocalizations field.
func (o *AppRelationships) SetBetaAppLocalizations(v AppRelationshipsBetaAppLocalizations) {
	o.BetaAppLocalizations = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *AppRelationships) GetBuilds() AppRelationshipsBuilds {
	if o == nil || o.Builds == nil {
		var ret AppRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool) {
	if o == nil || o.Builds == nil {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *AppRelationships) HasBuilds() bool {
	if o != nil && o.Builds != nil {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppRelationshipsBuilds and assigns it to the Builds field.
func (o *AppRelationships) SetBuilds(v AppRelationshipsBuilds) {
	o.Builds = &v
}

// GetBetaLicenseAgreement returns the BetaLicenseAgreement field value if set, zero value otherwise.
func (o *AppRelationships) GetBetaLicenseAgreement() AppRelationshipsBetaLicenseAgreement {
	if o == nil || o.BetaLicenseAgreement == nil {
		var ret AppRelationshipsBetaLicenseAgreement
		return ret
	}
	return *o.BetaLicenseAgreement
}

// GetBetaLicenseAgreementOk returns a tuple with the BetaLicenseAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetBetaLicenseAgreementOk() (*AppRelationshipsBetaLicenseAgreement, bool) {
	if o == nil || o.BetaLicenseAgreement == nil {
		return nil, false
	}
	return o.BetaLicenseAgreement, true
}

// HasBetaLicenseAgreement returns a boolean if a field has been set.
func (o *AppRelationships) HasBetaLicenseAgreement() bool {
	if o != nil && o.BetaLicenseAgreement != nil {
		return true
	}

	return false
}

// SetBetaLicenseAgreement gets a reference to the given AppRelationshipsBetaLicenseAgreement and assigns it to the BetaLicenseAgreement field.
func (o *AppRelationships) SetBetaLicenseAgreement(v AppRelationshipsBetaLicenseAgreement) {
	o.BetaLicenseAgreement = &v
}

// GetBetaAppReviewDetail returns the BetaAppReviewDetail field value if set, zero value otherwise.
func (o *AppRelationships) GetBetaAppReviewDetail() AppRelationshipsBetaAppReviewDetail {
	if o == nil || o.BetaAppReviewDetail == nil {
		var ret AppRelationshipsBetaAppReviewDetail
		return ret
	}
	return *o.BetaAppReviewDetail
}

// GetBetaAppReviewDetailOk returns a tuple with the BetaAppReviewDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetBetaAppReviewDetailOk() (*AppRelationshipsBetaAppReviewDetail, bool) {
	if o == nil || o.BetaAppReviewDetail == nil {
		return nil, false
	}
	return o.BetaAppReviewDetail, true
}

// HasBetaAppReviewDetail returns a boolean if a field has been set.
func (o *AppRelationships) HasBetaAppReviewDetail() bool {
	if o != nil && o.BetaAppReviewDetail != nil {
		return true
	}

	return false
}

// SetBetaAppReviewDetail gets a reference to the given AppRelationshipsBetaAppReviewDetail and assigns it to the BetaAppReviewDetail field.
func (o *AppRelationships) SetBetaAppReviewDetail(v AppRelationshipsBetaAppReviewDetail) {
	o.BetaAppReviewDetail = &v
}

// GetAppInfos returns the AppInfos field value if set, zero value otherwise.
func (o *AppRelationships) GetAppInfos() AppRelationshipsAppInfos {
	if o == nil || o.AppInfos == nil {
		var ret AppRelationshipsAppInfos
		return ret
	}
	return *o.AppInfos
}

// GetAppInfosOk returns a tuple with the AppInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetAppInfosOk() (*AppRelationshipsAppInfos, bool) {
	if o == nil || o.AppInfos == nil {
		return nil, false
	}
	return o.AppInfos, true
}

// HasAppInfos returns a boolean if a field has been set.
func (o *AppRelationships) HasAppInfos() bool {
	if o != nil && o.AppInfos != nil {
		return true
	}

	return false
}

// SetAppInfos gets a reference to the given AppRelationshipsAppInfos and assigns it to the AppInfos field.
func (o *AppRelationships) SetAppInfos(v AppRelationshipsAppInfos) {
	o.AppInfos = &v
}

// GetEndUserLicenseAgreement returns the EndUserLicenseAgreement field value if set, zero value otherwise.
func (o *AppRelationships) GetEndUserLicenseAgreement() AppRelationshipsEndUserLicenseAgreement {
	if o == nil || o.EndUserLicenseAgreement == nil {
		var ret AppRelationshipsEndUserLicenseAgreement
		return ret
	}
	return *o.EndUserLicenseAgreement
}

// GetEndUserLicenseAgreementOk returns a tuple with the EndUserLicenseAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetEndUserLicenseAgreementOk() (*AppRelationshipsEndUserLicenseAgreement, bool) {
	if o == nil || o.EndUserLicenseAgreement == nil {
		return nil, false
	}
	return o.EndUserLicenseAgreement, true
}

// HasEndUserLicenseAgreement returns a boolean if a field has been set.
func (o *AppRelationships) HasEndUserLicenseAgreement() bool {
	if o != nil && o.EndUserLicenseAgreement != nil {
		return true
	}

	return false
}

// SetEndUserLicenseAgreement gets a reference to the given AppRelationshipsEndUserLicenseAgreement and assigns it to the EndUserLicenseAgreement field.
func (o *AppRelationships) SetEndUserLicenseAgreement(v AppRelationshipsEndUserLicenseAgreement) {
	o.EndUserLicenseAgreement = &v
}

// GetPreOrder returns the PreOrder field value if set, zero value otherwise.
func (o *AppRelationships) GetPreOrder() AppRelationshipsPreOrder {
	if o == nil || o.PreOrder == nil {
		var ret AppRelationshipsPreOrder
		return ret
	}
	return *o.PreOrder
}

// GetPreOrderOk returns a tuple with the PreOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetPreOrderOk() (*AppRelationshipsPreOrder, bool) {
	if o == nil || o.PreOrder == nil {
		return nil, false
	}
	return o.PreOrder, true
}

// HasPreOrder returns a boolean if a field has been set.
func (o *AppRelationships) HasPreOrder() bool {
	if o != nil && o.PreOrder != nil {
		return true
	}

	return false
}

// SetPreOrder gets a reference to the given AppRelationshipsPreOrder and assigns it to the PreOrder field.
func (o *AppRelationships) SetPreOrder(v AppRelationshipsPreOrder) {
	o.PreOrder = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AppRelationships) GetPrices() AppRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret AppRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetPricesOk() (*AppRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AppRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given AppRelationshipsPrices and assigns it to the Prices field.
func (o *AppRelationships) SetPrices(v AppRelationshipsPrices) {
	o.Prices = &v
}

// GetAvailableTerritories returns the AvailableTerritories field value if set, zero value otherwise.
func (o *AppRelationships) GetAvailableTerritories() AppRelationshipsAvailableTerritories {
	if o == nil || o.AvailableTerritories == nil {
		var ret AppRelationshipsAvailableTerritories
		return ret
	}
	return *o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetAvailableTerritoriesOk() (*AppRelationshipsAvailableTerritories, bool) {
	if o == nil || o.AvailableTerritories == nil {
		return nil, false
	}
	return o.AvailableTerritories, true
}

// HasAvailableTerritories returns a boolean if a field has been set.
func (o *AppRelationships) HasAvailableTerritories() bool {
	if o != nil && o.AvailableTerritories != nil {
		return true
	}

	return false
}

// SetAvailableTerritories gets a reference to the given AppRelationshipsAvailableTerritories and assigns it to the AvailableTerritories field.
func (o *AppRelationships) SetAvailableTerritories(v AppRelationshipsAvailableTerritories) {
	o.AvailableTerritories = &v
}

// GetInAppPurchases returns the InAppPurchases field value if set, zero value otherwise.
func (o *AppRelationships) GetInAppPurchases() AppRelationshipsInAppPurchases {
	if o == nil || o.InAppPurchases == nil {
		var ret AppRelationshipsInAppPurchases
		return ret
	}
	return *o.InAppPurchases
}

// GetInAppPurchasesOk returns a tuple with the InAppPurchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetInAppPurchasesOk() (*AppRelationshipsInAppPurchases, bool) {
	if o == nil || o.InAppPurchases == nil {
		return nil, false
	}
	return o.InAppPurchases, true
}

// HasInAppPurchases returns a boolean if a field has been set.
func (o *AppRelationships) HasInAppPurchases() bool {
	if o != nil && o.InAppPurchases != nil {
		return true
	}

	return false
}

// SetInAppPurchases gets a reference to the given AppRelationshipsInAppPurchases and assigns it to the InAppPurchases field.
func (o *AppRelationships) SetInAppPurchases(v AppRelationshipsInAppPurchases) {
	o.InAppPurchases = &v
}

// GetGameCenterEnabledVersions returns the GameCenterEnabledVersions field value if set, zero value otherwise.
func (o *AppRelationships) GetGameCenterEnabledVersions() AppRelationshipsGameCenterEnabledVersions {
	if o == nil || o.GameCenterEnabledVersions == nil {
		var ret AppRelationshipsGameCenterEnabledVersions
		return ret
	}
	return *o.GameCenterEnabledVersions
}

// GetGameCenterEnabledVersionsOk returns a tuple with the GameCenterEnabledVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationships) GetGameCenterEnabledVersionsOk() (*AppRelationshipsGameCenterEnabledVersions, bool) {
	if o == nil || o.GameCenterEnabledVersions == nil {
		return nil, false
	}
	return o.GameCenterEnabledVersions, true
}

// HasGameCenterEnabledVersions returns a boolean if a field has been set.
func (o *AppRelationships) HasGameCenterEnabledVersions() bool {
	if o != nil && o.GameCenterEnabledVersions != nil {
		return true
	}

	return false
}

// SetGameCenterEnabledVersions gets a reference to the given AppRelationshipsGameCenterEnabledVersions and assigns it to the GameCenterEnabledVersions field.
func (o *AppRelationships) SetGameCenterEnabledVersions(v AppRelationshipsGameCenterEnabledVersions) {
	o.GameCenterEnabledVersions = &v
}

func (o AppRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BetaGroups != nil {
		toSerialize["betaGroups"] = o.BetaGroups
	}
	if o.AppStoreVersions != nil {
		toSerialize["appStoreVersions"] = o.AppStoreVersions
	}
	if o.PreReleaseVersions != nil {
		toSerialize["preReleaseVersions"] = o.PreReleaseVersions
	}
	if o.BetaAppLocalizations != nil {
		toSerialize["betaAppLocalizations"] = o.BetaAppLocalizations
	}
	if o.Builds != nil {
		toSerialize["builds"] = o.Builds
	}
	if o.BetaLicenseAgreement != nil {
		toSerialize["betaLicenseAgreement"] = o.BetaLicenseAgreement
	}
	if o.BetaAppReviewDetail != nil {
		toSerialize["betaAppReviewDetail"] = o.BetaAppReviewDetail
	}
	if o.AppInfos != nil {
		toSerialize["appInfos"] = o.AppInfos
	}
	if o.EndUserLicenseAgreement != nil {
		toSerialize["endUserLicenseAgreement"] = o.EndUserLicenseAgreement
	}
	if o.PreOrder != nil {
		toSerialize["preOrder"] = o.PreOrder
	}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	if o.AvailableTerritories != nil {
		toSerialize["availableTerritories"] = o.AvailableTerritories
	}
	if o.InAppPurchases != nil {
		toSerialize["inAppPurchases"] = o.InAppPurchases
	}
	if o.GameCenterEnabledVersions != nil {
		toSerialize["gameCenterEnabledVersions"] = o.GameCenterEnabledVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppRelationships := _AppRelationships{}

	if err = json.Unmarshal(bytes, &varAppRelationships); err == nil {
		*o = AppRelationships(varAppRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "betaGroups")
		delete(additionalProperties, "appStoreVersions")
		delete(additionalProperties, "preReleaseVersions")
		delete(additionalProperties, "betaAppLocalizations")
		delete(additionalProperties, "builds")
		delete(additionalProperties, "betaLicenseAgreement")
		delete(additionalProperties, "betaAppReviewDetail")
		delete(additionalProperties, "appInfos")
		delete(additionalProperties, "endUserLicenseAgreement")
		delete(additionalProperties, "preOrder")
		delete(additionalProperties, "prices")
		delete(additionalProperties, "availableTerritories")
		delete(additionalProperties, "inAppPurchases")
		delete(additionalProperties, "gameCenterEnabledVersions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppRelationships struct {
	value *AppRelationships
	isSet bool
}

func (v NullableAppRelationships) Get() *AppRelationships {
	return v.value
}

func (v *NullableAppRelationships) Set(val *AppRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationships(val *AppRelationships) *NullableAppRelationships {
	return &NullableAppRelationships{value: val, isSet: true}
}

func (v NullableAppRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

