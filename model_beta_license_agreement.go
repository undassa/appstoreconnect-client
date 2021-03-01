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

// BetaLicenseAgreement struct for BetaLicenseAgreement
type BetaLicenseAgreement struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaLicenseAgreementAttributes `json:"attributes,omitempty"`
	Relationships *AppEncryptionDeclarationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _BetaLicenseAgreement BetaLicenseAgreement

// NewBetaLicenseAgreement instantiates a new BetaLicenseAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreement(type_ string, id string, links ResourceLinks, ) *BetaLicenseAgreement {
	this := BetaLicenseAgreement{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewBetaLicenseAgreementWithDefaults instantiates a new BetaLicenseAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementWithDefaults() *BetaLicenseAgreement {
	this := BetaLicenseAgreement{}
	return &this
}

// GetType returns the Type field value
func (o *BetaLicenseAgreement) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaLicenseAgreement) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaLicenseAgreement) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaLicenseAgreement) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaLicenseAgreement) GetAttributes() BetaLicenseAgreementAttributes {
	if o == nil || o.Attributes == nil {
		var ret BetaLicenseAgreementAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaLicenseAgreement) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaLicenseAgreementAttributes and assigns it to the Attributes field.
func (o *BetaLicenseAgreement) SetAttributes(v BetaLicenseAgreementAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaLicenseAgreement) GetRelationships() AppEncryptionDeclarationRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppEncryptionDeclarationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetRelationshipsOk() (*AppEncryptionDeclarationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaLicenseAgreement) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppEncryptionDeclarationRelationships and assigns it to the Relationships field.
func (o *BetaLicenseAgreement) SetRelationships(v AppEncryptionDeclarationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *BetaLicenseAgreement) GetLinks() ResourceLinks {
	if o == nil  {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaLicenseAgreement) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o BetaLicenseAgreement) MarshalJSON() ([]byte, error) {
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

func (o *BetaLicenseAgreement) UnmarshalJSON(bytes []byte) (err error) {
	varBetaLicenseAgreement := _BetaLicenseAgreement{}

	if err = json.Unmarshal(bytes, &varBetaLicenseAgreement); err == nil {
		*o = BetaLicenseAgreement(varBetaLicenseAgreement)
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

type NullableBetaLicenseAgreement struct {
	value *BetaLicenseAgreement
	isSet bool
}

func (v NullableBetaLicenseAgreement) Get() *BetaLicenseAgreement {
	return v.value
}

func (v *NullableBetaLicenseAgreement) Set(val *BetaLicenseAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreement(val *BetaLicenseAgreement) *NullableBetaLicenseAgreement {
	return &NullableBetaLicenseAgreement{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


