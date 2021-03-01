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

// AppScreenshotSet struct for AppScreenshotSet
type AppScreenshotSet struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppScreenshotSetAttributes `json:"attributes,omitempty"`
	Relationships *AppScreenshotSetRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotSet AppScreenshotSet

// NewAppScreenshotSet instantiates a new AppScreenshotSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSet(type_ string, id string, links ResourceLinks, ) *AppScreenshotSet {
	this := AppScreenshotSet{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppScreenshotSetWithDefaults instantiates a new AppScreenshotSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetWithDefaults() *AppScreenshotSet {
	this := AppScreenshotSet{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshotSet) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSet) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshotSet) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppScreenshotSet) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSet) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppScreenshotSet) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppScreenshotSet) GetAttributes() AppScreenshotSetAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppScreenshotSetAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSet) GetAttributesOk() (*AppScreenshotSetAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppScreenshotSet) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppScreenshotSetAttributes and assigns it to the Attributes field.
func (o *AppScreenshotSet) SetAttributes(v AppScreenshotSetAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppScreenshotSet) GetRelationships() AppScreenshotSetRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppScreenshotSetRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSet) GetRelationshipsOk() (*AppScreenshotSetRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppScreenshotSet) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppScreenshotSetRelationships and assigns it to the Relationships field.
func (o *AppScreenshotSet) SetRelationships(v AppScreenshotSetRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppScreenshotSet) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSet) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppScreenshotSet) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppScreenshotSet) MarshalJSON() ([]byte, error) {
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

func (o *AppScreenshotSet) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotSet := _AppScreenshotSet{}

	if err = json.Unmarshal(bytes, &varAppScreenshotSet); err == nil {
		*o = AppScreenshotSet(varAppScreenshotSet)
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

type NullableAppScreenshotSet struct {
	value *AppScreenshotSet
	isSet bool
}

func (v NullableAppScreenshotSet) Get() *AppScreenshotSet {
	return v.value
}

func (v *NullableAppScreenshotSet) Set(val *AppScreenshotSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSet(val *AppScreenshotSet) *NullableAppScreenshotSet {
	return &NullableAppScreenshotSet{value: val, isSet: true}
}

func (v NullableAppScreenshotSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


