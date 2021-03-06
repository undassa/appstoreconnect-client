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

// UserInvitation struct for UserInvitation
type UserInvitation struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *UserInvitationAttributes `json:"attributes,omitempty"`
	Relationships *UserInvitationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitation UserInvitation

// NewUserInvitation instantiates a new UserInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitation(type_ string, id string, links ResourceLinks, ) *UserInvitation {
	this := UserInvitation{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewUserInvitationWithDefaults instantiates a new UserInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationWithDefaults() *UserInvitation {
	this := UserInvitation{}
	return &this
}

// GetType returns the Type field value
func (o *UserInvitation) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserInvitation) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *UserInvitation) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserInvitation) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserInvitation) GetAttributes() UserInvitationAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserInvitationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetAttributesOk() (*UserInvitationAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserInvitation) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UserInvitationAttributes and assigns it to the Attributes field.
func (o *UserInvitation) SetAttributes(v UserInvitationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserInvitation) GetRelationships() UserInvitationRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserInvitationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetRelationshipsOk() (*UserInvitationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserInvitation) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given UserInvitationRelationships and assigns it to the Relationships field.
func (o *UserInvitation) SetRelationships(v UserInvitationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *UserInvitation) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserInvitation) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o UserInvitation) MarshalJSON() ([]byte, error) {
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

func (o *UserInvitation) UnmarshalJSON(bytes []byte) (err error) {
	varUserInvitation := _UserInvitation{}

	if err = json.Unmarshal(bytes, &varUserInvitation); err == nil {
		*o = UserInvitation(varUserInvitation)
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

type NullableUserInvitation struct {
	value *UserInvitation
	isSet bool
}

func (v NullableUserInvitation) Get() *UserInvitation {
	return v.value
}

func (v *NullableUserInvitation) Set(val *UserInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitation(val *UserInvitation) *NullableUserInvitation {
	return &NullableUserInvitation{value: val, isSet: true}
}

func (v NullableUserInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


