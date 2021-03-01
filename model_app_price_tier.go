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

// AppPriceTier struct for AppPriceTier
type AppPriceTier struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Relationships *AppPriceTierRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppPriceTier AppPriceTier

// NewAppPriceTier instantiates a new AppPriceTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceTier(type_ string, id string, links ResourceLinks, ) *AppPriceTier {
	this := AppPriceTier{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppPriceTierWithDefaults instantiates a new AppPriceTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceTierWithDefaults() *AppPriceTier {
	this := AppPriceTier{}
	return &this
}

// GetType returns the Type field value
func (o *AppPriceTier) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPriceTier) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPriceTier) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPriceTier) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPriceTier) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPriceTier) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppPriceTier) GetRelationships() AppPriceTierRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppPriceTierRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceTier) GetRelationshipsOk() (*AppPriceTierRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppPriceTier) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPriceTierRelationships and assigns it to the Relationships field.
func (o *AppPriceTier) SetRelationships(v AppPriceTierRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppPriceTier) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPriceTier) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPriceTier) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppPriceTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPriceTier) UnmarshalJSON(bytes []byte) (err error) {
	varAppPriceTier := _AppPriceTier{}

	if err = json.Unmarshal(bytes, &varAppPriceTier); err == nil {
		*o = AppPriceTier(varAppPriceTier)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPriceTier struct {
	value *AppPriceTier
	isSet bool
}

func (v NullableAppPriceTier) Get() *AppPriceTier {
	return v.value
}

func (v *NullableAppPriceTier) Set(val *AppPriceTier) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceTier) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceTier(val *AppPriceTier) *NullableAppPriceTier {
	return &NullableAppPriceTier{value: val, isSet: true}
}

func (v NullableAppPriceTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


