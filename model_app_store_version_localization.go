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

// AppStoreVersionLocalization struct for AppStoreVersionLocalization
type AppStoreVersionLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionLocalizationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionLocalization AppStoreVersionLocalization

// NewAppStoreVersionLocalization instantiates a new AppStoreVersionLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalization(type_ string, id string, links ResourceLinks, ) *AppStoreVersionLocalization {
	this := AppStoreVersionLocalization{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppStoreVersionLocalizationWithDefaults instantiates a new AppStoreVersionLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationWithDefaults() *AppStoreVersionLocalization {
	this := AppStoreVersionLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionLocalization) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionLocalization) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionLocalization) GetAttributes() AppStoreVersionLocalizationAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppStoreVersionLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetAttributesOk() (*AppStoreVersionLocalizationAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionLocalization) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionLocalizationAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionLocalization) SetAttributes(v AppStoreVersionLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersionLocalization) GetRelationships() AppStoreVersionLocalizationRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppStoreVersionLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetRelationshipsOk() (*AppStoreVersionLocalizationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersionLocalization) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionLocalizationRelationships and assigns it to the Relationships field.
func (o *AppStoreVersionLocalization) SetRelationships(v AppStoreVersionLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionLocalization) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionLocalization) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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

func (o *AppStoreVersionLocalization) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionLocalization := _AppStoreVersionLocalization{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionLocalization); err == nil {
		*o = AppStoreVersionLocalization(varAppStoreVersionLocalization)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionLocalization struct {
	value *AppStoreVersionLocalization
	isSet bool
}

func (v NullableAppStoreVersionLocalization) Get() *AppStoreVersionLocalization {
	return v.value
}

func (v *NullableAppStoreVersionLocalization) Set(val *AppStoreVersionLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalization(val *AppStoreVersionLocalization) *NullableAppStoreVersionLocalization {
	return &NullableAppStoreVersionLocalization{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

