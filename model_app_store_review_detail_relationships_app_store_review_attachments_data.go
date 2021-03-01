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

// AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData struct for AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData
type AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData

// NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData instantiates a new AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData(type_ string, id string, ) *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData {
	this := AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataWithDefaults instantiates a new AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataWithDefaults() *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData {
	this := AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) SetId(v string) {
	o.Id = v
}

func (o AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData := _AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData); err == nil {
		*o = AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData(varAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData struct {
	value *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData
	isSet bool
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) Get() *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData {
	return v.value
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) Set(val *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData(val *AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData {
	return &NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


