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

// AppStoreVersionSubmissionResponse struct for AppStoreVersionSubmissionResponse
type AppStoreVersionSubmissionResponse struct {
	Data AppStoreVersionSubmission `json:"data"`
	Links DocumentLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionSubmissionResponse AppStoreVersionSubmissionResponse

// NewAppStoreVersionSubmissionResponse instantiates a new AppStoreVersionSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionSubmissionResponse(data AppStoreVersionSubmission, links DocumentLinks, ) *AppStoreVersionSubmissionResponse {
	this := AppStoreVersionSubmissionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionSubmissionResponseWithDefaults instantiates a new AppStoreVersionSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionSubmissionResponseWithDefaults() *AppStoreVersionSubmissionResponse {
	this := AppStoreVersionSubmissionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionSubmissionResponse) GetData() AppStoreVersionSubmission {
	if o == nil  {
		var ret AppStoreVersionSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionResponse) GetDataOk() (*AppStoreVersionSubmission, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionSubmissionResponse) SetData(v AppStoreVersionSubmission) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionSubmissionResponse) GetLinks() DocumentLinks {
	if o == nil  {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionSubmissionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionSubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionSubmissionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionSubmissionResponse := _AppStoreVersionSubmissionResponse{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionSubmissionResponse); err == nil {
		*o = AppStoreVersionSubmissionResponse(varAppStoreVersionSubmissionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionSubmissionResponse struct {
	value *AppStoreVersionSubmissionResponse
	isSet bool
}

func (v NullableAppStoreVersionSubmissionResponse) Get() *AppStoreVersionSubmissionResponse {
	return v.value
}

func (v *NullableAppStoreVersionSubmissionResponse) Set(val *AppStoreVersionSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionSubmissionResponse(val *AppStoreVersionSubmissionResponse) *NullableAppStoreVersionSubmissionResponse {
	return &NullableAppStoreVersionSubmissionResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

