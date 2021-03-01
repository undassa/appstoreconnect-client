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

// RoutingAppCoverage struct for RoutingAppCoverage
type RoutingAppCoverage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreReviewAttachmentAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionSubmissionRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _RoutingAppCoverage RoutingAppCoverage

// NewRoutingAppCoverage instantiates a new RoutingAppCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAppCoverage(type_ string, id string, links ResourceLinks, ) *RoutingAppCoverage {
	this := RoutingAppCoverage{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewRoutingAppCoverageWithDefaults instantiates a new RoutingAppCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAppCoverageWithDefaults() *RoutingAppCoverage {
	this := RoutingAppCoverage{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingAppCoverage) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverage) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingAppCoverage) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RoutingAppCoverage) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverage) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoutingAppCoverage) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RoutingAppCoverage) GetAttributes() AppStoreReviewAttachmentAttributes {
	if o == nil || o.Attributes == nil {
		var ret AppStoreReviewAttachmentAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverage) GetAttributesOk() (*AppStoreReviewAttachmentAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RoutingAppCoverage) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreReviewAttachmentAttributes and assigns it to the Attributes field.
func (o *RoutingAppCoverage) SetAttributes(v AppStoreReviewAttachmentAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RoutingAppCoverage) GetRelationships() AppStoreVersionSubmissionRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppStoreVersionSubmissionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverage) GetRelationshipsOk() (*AppStoreVersionSubmissionRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RoutingAppCoverage) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionSubmissionRelationships and assigns it to the Relationships field.
func (o *RoutingAppCoverage) SetRelationships(v AppStoreVersionSubmissionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *RoutingAppCoverage) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverage) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RoutingAppCoverage) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o RoutingAppCoverage) MarshalJSON() ([]byte, error) {
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

func (o *RoutingAppCoverage) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingAppCoverage := _RoutingAppCoverage{}

	if err = json.Unmarshal(bytes, &varRoutingAppCoverage); err == nil {
		*o = RoutingAppCoverage(varRoutingAppCoverage)
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

type NullableRoutingAppCoverage struct {
	value *RoutingAppCoverage
	isSet bool
}

func (v NullableRoutingAppCoverage) Get() *RoutingAppCoverage {
	return v.value
}

func (v *NullableRoutingAppCoverage) Set(val *RoutingAppCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAppCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAppCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAppCoverage(val *RoutingAppCoverage) *NullableRoutingAppCoverage {
	return &NullableRoutingAppCoverage{value: val, isSet: true}
}

func (v NullableRoutingAppCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAppCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


