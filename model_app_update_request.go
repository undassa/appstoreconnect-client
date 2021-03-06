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

// AppUpdateRequest struct for AppUpdateRequest
type AppUpdateRequest struct {
	Data AppUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppUpdateRequest AppUpdateRequest

// NewAppUpdateRequest instantiates a new AppUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUpdateRequest(data AppUpdateRequestData, ) *AppUpdateRequest {
	this := AppUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppUpdateRequestWithDefaults instantiates a new AppUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUpdateRequestWithDefaults() *AppUpdateRequest {
	this := AppUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppUpdateRequest) GetData() AppUpdateRequestData {
	if o == nil  {
		var ret AppUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppUpdateRequest) GetDataOk() (*AppUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppUpdateRequest) SetData(v AppUpdateRequestData) {
	o.Data = v
}

func (o AppUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAppUpdateRequest := _AppUpdateRequest{}

	if err = json.Unmarshal(bytes, &varAppUpdateRequest); err == nil {
		*o = AppUpdateRequest(varAppUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppUpdateRequest struct {
	value *AppUpdateRequest
	isSet bool
}

func (v NullableAppUpdateRequest) Get() *AppUpdateRequest {
	return v.value
}

func (v *NullableAppUpdateRequest) Set(val *AppUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUpdateRequest(val *AppUpdateRequest) *NullableAppUpdateRequest {
	return &NullableAppUpdateRequest{value: val, isSet: true}
}

func (v NullableAppUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


