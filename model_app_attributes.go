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

// AppAttributes struct for AppAttributes
type AppAttributes struct {
	Name *string `json:"name,omitempty"`
	BundleId *string `json:"bundleId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	PrimaryLocale *string `json:"primaryLocale,omitempty"`
	IsOrEverWasMadeForKids *bool `json:"isOrEverWasMadeForKids,omitempty"`
	AvailableInNewTerritories *bool `json:"availableInNewTerritories,omitempty"`
	ContentRightsDeclaration *string `json:"contentRightsDeclaration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAttributes AppAttributes

// NewAppAttributes instantiates a new AppAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAttributes() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// NewAppAttributesWithDefaults instantiates a new AppAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAttributesWithDefaults() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppAttributes) SetName(v string) {
	o.Name = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AppAttributes) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil || o.BundleId == nil {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AppAttributes) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AppAttributes) SetBundleId(v string) {
	o.BundleId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AppAttributes) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AppAttributes) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AppAttributes) SetSku(v string) {
	o.Sku = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *AppAttributes) GetPrimaryLocale() string {
	if o == nil || o.PrimaryLocale == nil {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || o.PrimaryLocale == nil {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// HasPrimaryLocale returns a boolean if a field has been set.
func (o *AppAttributes) HasPrimaryLocale() bool {
	if o != nil && o.PrimaryLocale != nil {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *AppAttributes) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

// GetIsOrEverWasMadeForKids returns the IsOrEverWasMadeForKids field value if set, zero value otherwise.
func (o *AppAttributes) GetIsOrEverWasMadeForKids() bool {
	if o == nil || o.IsOrEverWasMadeForKids == nil {
		var ret bool
		return ret
	}
	return *o.IsOrEverWasMadeForKids
}

// GetIsOrEverWasMadeForKidsOk returns a tuple with the IsOrEverWasMadeForKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetIsOrEverWasMadeForKidsOk() (*bool, bool) {
	if o == nil || o.IsOrEverWasMadeForKids == nil {
		return nil, false
	}
	return o.IsOrEverWasMadeForKids, true
}

// HasIsOrEverWasMadeForKids returns a boolean if a field has been set.
func (o *AppAttributes) HasIsOrEverWasMadeForKids() bool {
	if o != nil && o.IsOrEverWasMadeForKids != nil {
		return true
	}

	return false
}

// SetIsOrEverWasMadeForKids gets a reference to the given bool and assigns it to the IsOrEverWasMadeForKids field.
func (o *AppAttributes) SetIsOrEverWasMadeForKids(v bool) {
	o.IsOrEverWasMadeForKids = &v
}

// GetAvailableInNewTerritories returns the AvailableInNewTerritories field value if set, zero value otherwise.
func (o *AppAttributes) GetAvailableInNewTerritories() bool {
	if o == nil || o.AvailableInNewTerritories == nil {
		var ret bool
		return ret
	}
	return *o.AvailableInNewTerritories
}

// GetAvailableInNewTerritoriesOk returns a tuple with the AvailableInNewTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetAvailableInNewTerritoriesOk() (*bool, bool) {
	if o == nil || o.AvailableInNewTerritories == nil {
		return nil, false
	}
	return o.AvailableInNewTerritories, true
}

// HasAvailableInNewTerritories returns a boolean if a field has been set.
func (o *AppAttributes) HasAvailableInNewTerritories() bool {
	if o != nil && o.AvailableInNewTerritories != nil {
		return true
	}

	return false
}

// SetAvailableInNewTerritories gets a reference to the given bool and assigns it to the AvailableInNewTerritories field.
func (o *AppAttributes) SetAvailableInNewTerritories(v bool) {
	o.AvailableInNewTerritories = &v
}

// GetContentRightsDeclaration returns the ContentRightsDeclaration field value if set, zero value otherwise.
func (o *AppAttributes) GetContentRightsDeclaration() string {
	if o == nil || o.ContentRightsDeclaration == nil {
		var ret string
		return ret
	}
	return *o.ContentRightsDeclaration
}

// GetContentRightsDeclarationOk returns a tuple with the ContentRightsDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetContentRightsDeclarationOk() (*string, bool) {
	if o == nil || o.ContentRightsDeclaration == nil {
		return nil, false
	}
	return o.ContentRightsDeclaration, true
}

// HasContentRightsDeclaration returns a boolean if a field has been set.
func (o *AppAttributes) HasContentRightsDeclaration() bool {
	if o != nil && o.ContentRightsDeclaration != nil {
		return true
	}

	return false
}

// SetContentRightsDeclaration gets a reference to the given string and assigns it to the ContentRightsDeclaration field.
func (o *AppAttributes) SetContentRightsDeclaration(v string) {
	o.ContentRightsDeclaration = &v
}

func (o AppAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BundleId != nil {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.PrimaryLocale != nil {
		toSerialize["primaryLocale"] = o.PrimaryLocale
	}
	if o.IsOrEverWasMadeForKids != nil {
		toSerialize["isOrEverWasMadeForKids"] = o.IsOrEverWasMadeForKids
	}
	if o.AvailableInNewTerritories != nil {
		toSerialize["availableInNewTerritories"] = o.AvailableInNewTerritories
	}
	if o.ContentRightsDeclaration != nil {
		toSerialize["contentRightsDeclaration"] = o.ContentRightsDeclaration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppAttributes := _AppAttributes{}

	if err = json.Unmarshal(bytes, &varAppAttributes); err == nil {
		*o = AppAttributes(varAppAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "bundleId")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "primaryLocale")
		delete(additionalProperties, "isOrEverWasMadeForKids")
		delete(additionalProperties, "availableInNewTerritories")
		delete(additionalProperties, "contentRightsDeclaration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAttributes struct {
	value *AppAttributes
	isSet bool
}

func (v NullableAppAttributes) Get() *AppAttributes {
	return v.value
}

func (v *NullableAppAttributes) Set(val *AppAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAttributes(val *AppAttributes) *NullableAppAttributes {
	return &NullableAppAttributes{value: val, isSet: true}
}

func (v NullableAppAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


