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

// AppInfoUpdateRequest struct for AppInfoUpdateRequest
type AppInfoUpdateRequest struct {
	Data AppInfoUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppInfoUpdateRequest AppInfoUpdateRequest

// NewAppInfoUpdateRequest instantiates a new AppInfoUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoUpdateRequest(data AppInfoUpdateRequestData, ) *AppInfoUpdateRequest {
	this := AppInfoUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppInfoUpdateRequestWithDefaults instantiates a new AppInfoUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoUpdateRequestWithDefaults() *AppInfoUpdateRequest {
	this := AppInfoUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppInfoUpdateRequest) GetData() AppInfoUpdateRequestData {
	if o == nil  {
		var ret AppInfoUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppInfoUpdateRequest) GetDataOk() (*AppInfoUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppInfoUpdateRequest) SetData(v AppInfoUpdateRequestData) {
	o.Data = v
}

func (o AppInfoUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfoUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfoUpdateRequest := _AppInfoUpdateRequest{}

	if err = json.Unmarshal(bytes, &varAppInfoUpdateRequest); err == nil {
		*o = AppInfoUpdateRequest(varAppInfoUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfoUpdateRequest struct {
	value *AppInfoUpdateRequest
	isSet bool
}

func (v NullableAppInfoUpdateRequest) Get() *AppInfoUpdateRequest {
	return v.value
}

func (v *NullableAppInfoUpdateRequest) Set(val *AppInfoUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoUpdateRequest(val *AppInfoUpdateRequest) *NullableAppInfoUpdateRequest {
	return &NullableAppInfoUpdateRequest{value: val, isSet: true}
}

func (v NullableAppInfoUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


