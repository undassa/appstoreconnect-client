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

// EndUserLicenseAgreementRelationships struct for EndUserLicenseAgreementRelationships
type EndUserLicenseAgreementRelationships struct {
	App *AppEncryptionDeclarationRelationshipsApp `json:"app,omitempty"`
	Territories *AppRelationshipsAvailableTerritories `json:"territories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EndUserLicenseAgreementRelationships EndUserLicenseAgreementRelationships

// NewEndUserLicenseAgreementRelationships instantiates a new EndUserLicenseAgreementRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementRelationships() *EndUserLicenseAgreementRelationships {
	this := EndUserLicenseAgreementRelationships{}
	return &this
}

// NewEndUserLicenseAgreementRelationshipsWithDefaults instantiates a new EndUserLicenseAgreementRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementRelationshipsWithDefaults() *EndUserLicenseAgreementRelationships {
	this := EndUserLicenseAgreementRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementRelationships) GetApp() AppEncryptionDeclarationRelationshipsApp {
	if o == nil || o.App == nil {
		var ret AppEncryptionDeclarationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementRelationships) GetAppOk() (*AppEncryptionDeclarationRelationshipsApp, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementRelationships) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppEncryptionDeclarationRelationshipsApp and assigns it to the App field.
func (o *EndUserLicenseAgreementRelationships) SetApp(v AppEncryptionDeclarationRelationshipsApp) {
	o.App = &v
}

// GetTerritories returns the Territories field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementRelationships) GetTerritories() AppRelationshipsAvailableTerritories {
	if o == nil || o.Territories == nil {
		var ret AppRelationshipsAvailableTerritories
		return ret
	}
	return *o.Territories
}

// GetTerritoriesOk returns a tuple with the Territories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementRelationships) GetTerritoriesOk() (*AppRelationshipsAvailableTerritories, bool) {
	if o == nil || o.Territories == nil {
		return nil, false
	}
	return o.Territories, true
}

// HasTerritories returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementRelationships) HasTerritories() bool {
	if o != nil && o.Territories != nil {
		return true
	}

	return false
}

// SetTerritories gets a reference to the given AppRelationshipsAvailableTerritories and assigns it to the Territories field.
func (o *EndUserLicenseAgreementRelationships) SetTerritories(v AppRelationshipsAvailableTerritories) {
	o.Territories = &v
}

func (o EndUserLicenseAgreementRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Territories != nil {
		toSerialize["territories"] = o.Territories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EndUserLicenseAgreementRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varEndUserLicenseAgreementRelationships := _EndUserLicenseAgreementRelationships{}

	if err = json.Unmarshal(bytes, &varEndUserLicenseAgreementRelationships); err == nil {
		*o = EndUserLicenseAgreementRelationships(varEndUserLicenseAgreementRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "territories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndUserLicenseAgreementRelationships struct {
	value *EndUserLicenseAgreementRelationships
	isSet bool
}

func (v NullableEndUserLicenseAgreementRelationships) Get() *EndUserLicenseAgreementRelationships {
	return v.value
}

func (v *NullableEndUserLicenseAgreementRelationships) Set(val *EndUserLicenseAgreementRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementRelationships(val *EndUserLicenseAgreementRelationships) *NullableEndUserLicenseAgreementRelationships {
	return &NullableEndUserLicenseAgreementRelationships{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

