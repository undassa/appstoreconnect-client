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

// CertificateCreateRequestData struct for CertificateCreateRequestData
type CertificateCreateRequestData struct {
	Type string `json:"type"`
	Attributes CertificateCreateRequestDataAttributes `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _CertificateCreateRequestData CertificateCreateRequestData

// NewCertificateCreateRequestData instantiates a new CertificateCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateCreateRequestData(type_ string, attributes CertificateCreateRequestDataAttributes, ) *CertificateCreateRequestData {
	this := CertificateCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCertificateCreateRequestDataWithDefaults instantiates a new CertificateCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateCreateRequestDataWithDefaults() *CertificateCreateRequestData {
	this := CertificateCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *CertificateCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CertificateCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CertificateCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CertificateCreateRequestData) GetAttributes() CertificateCreateRequestDataAttributes {
	if o == nil  {
		var ret CertificateCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CertificateCreateRequestData) GetAttributesOk() (*CertificateCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CertificateCreateRequestData) SetAttributes(v CertificateCreateRequestDataAttributes) {
	o.Attributes = v
}

func (o CertificateCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificateCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varCertificateCreateRequestData := _CertificateCreateRequestData{}

	if err = json.Unmarshal(bytes, &varCertificateCreateRequestData); err == nil {
		*o = CertificateCreateRequestData(varCertificateCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateCreateRequestData struct {
	value *CertificateCreateRequestData
	isSet bool
}

func (v NullableCertificateCreateRequestData) Get() *CertificateCreateRequestData {
	return v.value
}

func (v *NullableCertificateCreateRequestData) Set(val *CertificateCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateCreateRequestData(val *CertificateCreateRequestData) *NullableCertificateCreateRequestData {
	return &NullableCertificateCreateRequestData{value: val, isSet: true}
}

func (v NullableCertificateCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


