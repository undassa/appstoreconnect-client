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

// BetaAppLocalizationCreateRequest struct for BetaAppLocalizationCreateRequest
type BetaAppLocalizationCreateRequest struct {
	Data BetaAppLocalizationCreateRequestData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BetaAppLocalizationCreateRequest BetaAppLocalizationCreateRequest

// NewBetaAppLocalizationCreateRequest instantiates a new BetaAppLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationCreateRequest(data BetaAppLocalizationCreateRequestData, ) *BetaAppLocalizationCreateRequest {
	this := BetaAppLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppLocalizationCreateRequestWithDefaults instantiates a new BetaAppLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationCreateRequestWithDefaults() *BetaAppLocalizationCreateRequest {
	this := BetaAppLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppLocalizationCreateRequest) GetData() BetaAppLocalizationCreateRequestData {
	if o == nil  {
		var ret BetaAppLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequest) GetDataOk() (*BetaAppLocalizationCreateRequestData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppLocalizationCreateRequest) SetData(v BetaAppLocalizationCreateRequestData) {
	o.Data = v
}

func (o BetaAppLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaAppLocalizationCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBetaAppLocalizationCreateRequest := _BetaAppLocalizationCreateRequest{}

	if err = json.Unmarshal(bytes, &varBetaAppLocalizationCreateRequest); err == nil {
		*o = BetaAppLocalizationCreateRequest(varBetaAppLocalizationCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaAppLocalizationCreateRequest struct {
	value *BetaAppLocalizationCreateRequest
	isSet bool
}

func (v NullableBetaAppLocalizationCreateRequest) Get() *BetaAppLocalizationCreateRequest {
	return v.value
}

func (v *NullableBetaAppLocalizationCreateRequest) Set(val *BetaAppLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationCreateRequest(val *BetaAppLocalizationCreateRequest) *NullableBetaAppLocalizationCreateRequest {
	return &NullableBetaAppLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


