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

// AppPrice struct for AppPrice
type AppPrice struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Relationships *AppPriceRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppPrice AppPrice

// NewAppPrice instantiates a new AppPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPrice(type_ string, id string, links ResourceLinks, ) *AppPrice {
	this := AppPrice{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppPriceWithDefaults instantiates a new AppPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceWithDefaults() *AppPrice {
	this := AppPrice{}
	return &this
}

// GetType returns the Type field value
func (o *AppPrice) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPrice) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPrice) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPrice) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPrice) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPrice) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppPrice) GetRelationships() AppPriceRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppPriceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrice) GetRelationshipsOk() (*AppPriceRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppPrice) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPriceRelationships and assigns it to the Relationships field.
func (o *AppPrice) SetRelationships(v AppPriceRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppPrice) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPrice) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPrice) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppPrice) MarshalJSON() ([]byte, error) {
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

func (o *AppPrice) UnmarshalJSON(bytes []byte) (err error) {
	varAppPrice := _AppPrice{}

	if err = json.Unmarshal(bytes, &varAppPrice); err == nil {
		*o = AppPrice(varAppPrice)
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

type NullableAppPrice struct {
	value *AppPrice
	isSet bool
}

func (v NullableAppPrice) Get() *AppPrice {
	return v.value
}

func (v *NullableAppPrice) Set(val *AppPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPrice(val *AppPrice) *NullableAppPrice {
	return &NullableAppPrice{value: val, isSet: true}
}

func (v NullableAppPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


