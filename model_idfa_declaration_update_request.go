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

// IdfaDeclarationUpdateRequest struct for IdfaDeclarationUpdateRequest
type IdfaDeclarationUpdateRequest struct {
	Data IdfaDeclarationUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _IdfaDeclarationUpdateRequest IdfaDeclarationUpdateRequest

// NewIdfaDeclarationUpdateRequest instantiates a new IdfaDeclarationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdfaDeclarationUpdateRequest(data IdfaDeclarationUpdateRequestData, ) *IdfaDeclarationUpdateRequest {
	this := IdfaDeclarationUpdateRequest{}
	this.Data = data
	return &this
}

// NewIdfaDeclarationUpdateRequestWithDefaults instantiates a new IdfaDeclarationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdfaDeclarationUpdateRequestWithDefaults() *IdfaDeclarationUpdateRequest {
	this := IdfaDeclarationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *IdfaDeclarationUpdateRequest) GetData() IdfaDeclarationUpdateRequestData {
	if o == nil  {
		var ret IdfaDeclarationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IdfaDeclarationUpdateRequest) GetDataOk() (*IdfaDeclarationUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IdfaDeclarationUpdateRequest) SetData(v IdfaDeclarationUpdateRequestData) {
	o.Data = v
}

func (o IdfaDeclarationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdfaDeclarationUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varIdfaDeclarationUpdateRequest := _IdfaDeclarationUpdateRequest{}

	if err = json.Unmarshal(bytes, &varIdfaDeclarationUpdateRequest); err == nil {
		*o = IdfaDeclarationUpdateRequest(varIdfaDeclarationUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdfaDeclarationUpdateRequest struct {
	value *IdfaDeclarationUpdateRequest
	isSet bool
}

func (v NullableIdfaDeclarationUpdateRequest) Get() *IdfaDeclarationUpdateRequest {
	return v.value
}

func (v *NullableIdfaDeclarationUpdateRequest) Set(val *IdfaDeclarationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdfaDeclarationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdfaDeclarationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdfaDeclarationUpdateRequest(val *IdfaDeclarationUpdateRequest) *NullableIdfaDeclarationUpdateRequest {
	return &NullableIdfaDeclarationUpdateRequest{value: val, isSet: true}
}

func (v NullableIdfaDeclarationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdfaDeclarationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


