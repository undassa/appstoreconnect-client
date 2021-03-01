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

// AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct for AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
type AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail

// NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail instantiates a new AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}
	return &this
}

// NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetailWithDefaults instantiates a new AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetailWithDefaults() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetData() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData {
	if o == nil || o.Data == nil {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetDataOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData and assigns it to the Data field.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) SetData(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData) {
	o.Data = &v
}

func (o AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail := _AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail); err == nil {
		*o = AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail(varAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct {
	value *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
	isSet bool
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Get() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Set(val *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail(val *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	return &NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


