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

// AppStoreVersionLocalizationUpdateRequest struct for AppStoreVersionLocalizationUpdateRequest
type AppStoreVersionLocalizationUpdateRequest struct {
	Data AppStoreVersionLocalizationUpdateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionLocalizationUpdateRequest AppStoreVersionLocalizationUpdateRequest

// NewAppStoreVersionLocalizationUpdateRequest instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationUpdateRequest(data AppStoreVersionLocalizationUpdateRequestData, ) *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionLocalizationUpdateRequestWithDefaults instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationUpdateRequestWithDefaults() *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationUpdateRequest) GetData() AppStoreVersionLocalizationUpdateRequestData {
	if o == nil  {
		var ret AppStoreVersionLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationUpdateRequest) GetDataOk() (*AppStoreVersionLocalizationUpdateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationUpdateRequest) SetData(v AppStoreVersionLocalizationUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionLocalizationUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionLocalizationUpdateRequest := _AppStoreVersionLocalizationUpdateRequest{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionLocalizationUpdateRequest); err == nil {
		*o = AppStoreVersionLocalizationUpdateRequest(varAppStoreVersionLocalizationUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionLocalizationUpdateRequest struct {
	value *AppStoreVersionLocalizationUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) Get() *AppStoreVersionLocalizationUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Set(val *AppStoreVersionLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationUpdateRequest(val *AppStoreVersionLocalizationUpdateRequest) *NullableAppStoreVersionLocalizationUpdateRequest {
	return &NullableAppStoreVersionLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

