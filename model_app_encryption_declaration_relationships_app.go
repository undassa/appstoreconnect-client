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

// AppEncryptionDeclarationRelationshipsApp struct for AppEncryptionDeclarationRelationshipsApp
type AppEncryptionDeclarationRelationshipsApp struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppEncryptionDeclarationRelationshipsAppData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppEncryptionDeclarationRelationshipsApp AppEncryptionDeclarationRelationshipsApp

// NewAppEncryptionDeclarationRelationshipsApp instantiates a new AppEncryptionDeclarationRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationRelationshipsApp() *AppEncryptionDeclarationRelationshipsApp {
	this := AppEncryptionDeclarationRelationshipsApp{}
	return &this
}

// NewAppEncryptionDeclarationRelationshipsAppWithDefaults instantiates a new AppEncryptionDeclarationRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationRelationshipsAppWithDefaults() *AppEncryptionDeclarationRelationshipsApp {
	this := AppEncryptionDeclarationRelationshipsApp{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationRelationshipsApp) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsApp) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationRelationshipsApp) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppEncryptionDeclarationRelationshipsApp) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationRelationshipsApp) GetData() AppEncryptionDeclarationRelationshipsAppData {
	if o == nil || o.Data == nil {
		var ret AppEncryptionDeclarationRelationshipsAppData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsApp) GetDataOk() (*AppEncryptionDeclarationRelationshipsAppData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationRelationshipsApp) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEncryptionDeclarationRelationshipsAppData and assigns it to the Data field.
func (o *AppEncryptionDeclarationRelationshipsApp) SetData(v AppEncryptionDeclarationRelationshipsAppData) {
	o.Data = &v
}

func (o AppEncryptionDeclarationRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppEncryptionDeclarationRelationshipsApp) UnmarshalJSON(bytes []byte) (err error) {
	varAppEncryptionDeclarationRelationshipsApp := _AppEncryptionDeclarationRelationshipsApp{}

	if err = json.Unmarshal(bytes, &varAppEncryptionDeclarationRelationshipsApp); err == nil {
		*o = AppEncryptionDeclarationRelationshipsApp(varAppEncryptionDeclarationRelationshipsApp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppEncryptionDeclarationRelationshipsApp struct {
	value *AppEncryptionDeclarationRelationshipsApp
	isSet bool
}

func (v NullableAppEncryptionDeclarationRelationshipsApp) Get() *AppEncryptionDeclarationRelationshipsApp {
	return v.value
}

func (v *NullableAppEncryptionDeclarationRelationshipsApp) Set(val *AppEncryptionDeclarationRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationRelationshipsApp(val *AppEncryptionDeclarationRelationshipsApp) *NullableAppEncryptionDeclarationRelationshipsApp {
	return &NullableAppEncryptionDeclarationRelationshipsApp{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


