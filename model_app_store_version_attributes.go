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
	"time"
)

// AppStoreVersionAttributes struct for AppStoreVersionAttributes
type AppStoreVersionAttributes struct {
	Platform *Platform `json:"platform,omitempty"`
	VersionString *string `json:"versionString,omitempty"`
	AppStoreState *AppStoreVersionState `json:"appStoreState,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	ReleaseType *string `json:"releaseType,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	UsesIdfa *bool `json:"usesIdfa,omitempty"`
	Downloadable *bool `json:"downloadable,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionAttributes AppStoreVersionAttributes

// NewAppStoreVersionAttributes instantiates a new AppStoreVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionAttributes() *AppStoreVersionAttributes {
	this := AppStoreVersionAttributes{}
	return &this
}

// NewAppStoreVersionAttributesWithDefaults instantiates a new AppStoreVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionAttributesWithDefaults() *AppStoreVersionAttributes {
	this := AppStoreVersionAttributes{}
	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetPlatform() Platform {
	if o == nil || o.Platform == nil {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *AppStoreVersionAttributes) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetVersionString returns the VersionString field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetVersionString() string {
	if o == nil || o.VersionString == nil {
		var ret string
		return ret
	}
	return *o.VersionString
}

// GetVersionStringOk returns a tuple with the VersionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetVersionStringOk() (*string, bool) {
	if o == nil || o.VersionString == nil {
		return nil, false
	}
	return o.VersionString, true
}

// HasVersionString returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasVersionString() bool {
	if o != nil && o.VersionString != nil {
		return true
	}

	return false
}

// SetVersionString gets a reference to the given string and assigns it to the VersionString field.
func (o *AppStoreVersionAttributes) SetVersionString(v string) {
	o.VersionString = &v
}

// GetAppStoreState returns the AppStoreState field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetAppStoreState() AppStoreVersionState {
	if o == nil || o.AppStoreState == nil {
		var ret AppStoreVersionState
		return ret
	}
	return *o.AppStoreState
}

// GetAppStoreStateOk returns a tuple with the AppStoreState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetAppStoreStateOk() (*AppStoreVersionState, bool) {
	if o == nil || o.AppStoreState == nil {
		return nil, false
	}
	return o.AppStoreState, true
}

// HasAppStoreState returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasAppStoreState() bool {
	if o != nil && o.AppStoreState != nil {
		return true
	}

	return false
}

// SetAppStoreState gets a reference to the given AppStoreVersionState and assigns it to the AppStoreState field.
func (o *AppStoreVersionAttributes) SetAppStoreState(v AppStoreVersionState) {
	o.AppStoreState = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetCopyright() string {
	if o == nil || o.Copyright == nil {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetCopyrightOk() (*string, bool) {
	if o == nil || o.Copyright == nil {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasCopyright() bool {
	if o != nil && o.Copyright != nil {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *AppStoreVersionAttributes) SetCopyright(v string) {
	o.Copyright = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetReleaseType() string {
	if o == nil || o.ReleaseType == nil {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetReleaseTypeOk() (*string, bool) {
	if o == nil || o.ReleaseType == nil {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasReleaseType() bool {
	if o != nil && o.ReleaseType != nil {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *AppStoreVersionAttributes) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetEarliestReleaseDate returns the EarliestReleaseDate field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetEarliestReleaseDate() time.Time {
	if o == nil || o.EarliestReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EarliestReleaseDate
}

// GetEarliestReleaseDateOk returns a tuple with the EarliestReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetEarliestReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.EarliestReleaseDate == nil {
		return nil, false
	}
	return o.EarliestReleaseDate, true
}

// HasEarliestReleaseDate returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasEarliestReleaseDate() bool {
	if o != nil && o.EarliestReleaseDate != nil {
		return true
	}

	return false
}

// SetEarliestReleaseDate gets a reference to the given time.Time and assigns it to the EarliestReleaseDate field.
func (o *AppStoreVersionAttributes) SetEarliestReleaseDate(v time.Time) {
	o.EarliestReleaseDate = &v
}

// GetUsesIdfa returns the UsesIdfa field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetUsesIdfa() bool {
	if o == nil || o.UsesIdfa == nil {
		var ret bool
		return ret
	}
	return *o.UsesIdfa
}

// GetUsesIdfaOk returns a tuple with the UsesIdfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetUsesIdfaOk() (*bool, bool) {
	if o == nil || o.UsesIdfa == nil {
		return nil, false
	}
	return o.UsesIdfa, true
}

// HasUsesIdfa returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasUsesIdfa() bool {
	if o != nil && o.UsesIdfa != nil {
		return true
	}

	return false
}

// SetUsesIdfa gets a reference to the given bool and assigns it to the UsesIdfa field.
func (o *AppStoreVersionAttributes) SetUsesIdfa(v bool) {
	o.UsesIdfa = &v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetDownloadable() bool {
	if o == nil || o.Downloadable == nil {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetDownloadableOk() (*bool, bool) {
	if o == nil || o.Downloadable == nil {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasDownloadable() bool {
	if o != nil && o.Downloadable != nil {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *AppStoreVersionAttributes) SetDownloadable(v bool) {
	o.Downloadable = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AppStoreVersionAttributes) GetCreatedDate() time.Time {
	if o == nil || o.CreatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AppStoreVersionAttributes) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *AppStoreVersionAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

func (o AppStoreVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.VersionString != nil {
		toSerialize["versionString"] = o.VersionString
	}
	if o.AppStoreState != nil {
		toSerialize["appStoreState"] = o.AppStoreState
	}
	if o.Copyright != nil {
		toSerialize["copyright"] = o.Copyright
	}
	if o.ReleaseType != nil {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if o.EarliestReleaseDate != nil {
		toSerialize["earliestReleaseDate"] = o.EarliestReleaseDate
	}
	if o.UsesIdfa != nil {
		toSerialize["usesIdfa"] = o.UsesIdfa
	}
	if o.Downloadable != nil {
		toSerialize["downloadable"] = o.Downloadable
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionAttributes := _AppStoreVersionAttributes{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionAttributes); err == nil {
		*o = AppStoreVersionAttributes(varAppStoreVersionAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "platform")
		delete(additionalProperties, "versionString")
		delete(additionalProperties, "appStoreState")
		delete(additionalProperties, "copyright")
		delete(additionalProperties, "releaseType")
		delete(additionalProperties, "earliestReleaseDate")
		delete(additionalProperties, "usesIdfa")
		delete(additionalProperties, "downloadable")
		delete(additionalProperties, "createdDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionAttributes struct {
	value *AppStoreVersionAttributes
	isSet bool
}

func (v NullableAppStoreVersionAttributes) Get() *AppStoreVersionAttributes {
	return v.value
}

func (v *NullableAppStoreVersionAttributes) Set(val *AppStoreVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionAttributes(val *AppStoreVersionAttributes) *NullableAppStoreVersionAttributes {
	return &NullableAppStoreVersionAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


