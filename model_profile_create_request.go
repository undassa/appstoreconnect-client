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

// ProfileCreateRequest struct for ProfileCreateRequest
type ProfileCreateRequest struct {
	Data ProfileCreateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _ProfileCreateRequest ProfileCreateRequest

// NewProfileCreateRequest instantiates a new ProfileCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCreateRequest(data ProfileCreateRequestData, ) *ProfileCreateRequest {
	this := ProfileCreateRequest{}
	this.Data = data
	return &this
}

// NewProfileCreateRequestWithDefaults instantiates a new ProfileCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileCreateRequestWithDefaults() *ProfileCreateRequest {
	this := ProfileCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ProfileCreateRequest) GetData() ProfileCreateRequestData {
	if o == nil  {
		var ret ProfileCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProfileCreateRequest) GetDataOk() (*ProfileCreateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProfileCreateRequest) SetData(v ProfileCreateRequestData) {
	o.Data = v
}

func (o ProfileCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varProfileCreateRequest := _ProfileCreateRequest{}

	if err = json.Unmarshal(bytes, &varProfileCreateRequest); err == nil {
		*o = ProfileCreateRequest(varProfileCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileCreateRequest struct {
	value *ProfileCreateRequest
	isSet bool
}

func (v NullableProfileCreateRequest) Get() *ProfileCreateRequest {
	return v.value
}

func (v *NullableProfileCreateRequest) Set(val *ProfileCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCreateRequest(val *ProfileCreateRequest) *NullableProfileCreateRequest {
	return &NullableProfileCreateRequest{value: val, isSet: true}
}

func (v NullableProfileCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


