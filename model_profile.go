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

// Profile struct for Profile
type Profile struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *ProfileAttributes `json:"attributes,omitempty"`
	Relationships *ProfileRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _Profile Profile

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile(type_ string, id string, links ResourceLinks, ) *Profile {
	this := Profile{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetType returns the Type field value
func (o *Profile) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Profile) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Profile) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *Profile) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Profile) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Profile) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Profile) GetAttributes() ProfileAttributes {
	if o == nil || o.Attributes == nil {
		var ret ProfileAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetAttributesOk() (*ProfileAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Profile) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ProfileAttributes and assigns it to the Attributes field.
func (o *Profile) SetAttributes(v ProfileAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Profile) GetRelationships() ProfileRelationships {
	if o == nil || o.Relationships == nil {
		var ret ProfileRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetRelationshipsOk() (*ProfileRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Profile) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ProfileRelationships and assigns it to the Relationships field.
func (o *Profile) SetRelationships(v ProfileRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *Profile) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Profile) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *Profile) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o Profile) MarshalJSON() ([]byte, error) {
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

func (o *Profile) UnmarshalJSON(bytes []byte) (err error) {
	varProfile := _Profile{}

	if err = json.Unmarshal(bytes, &varProfile); err == nil {
		*o = Profile(varProfile)
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

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


