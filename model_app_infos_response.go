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

// AppInfosResponse struct for AppInfosResponse
type AppInfosResponse struct {
	Data []AppInfo `json:"data"`
	Included *[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppInfosResponse AppInfosResponse

// NewAppInfosResponse instantiates a new AppInfosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfosResponse(data []AppInfo, links PagedDocumentLinks, ) *AppInfosResponse {
	this := AppInfosResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppInfosResponseWithDefaults instantiates a new AppInfosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfosResponseWithDefaults() *AppInfosResponse {
	this := AppInfosResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppInfosResponse) GetData() []AppInfo {
	if o == nil  {
		var ret []AppInfo
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppInfosResponse) GetDataOk() (*[]AppInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppInfosResponse) SetData(v []AppInfo) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppInfosResponse) GetIncluded() []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory {
	if o == nil || o.Included == nil {
		var ret []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfosResponse) GetIncludedOk() (*[]OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppInfosResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory and assigns it to the Included field.
func (o *AppInfosResponse) SetIncluded(v []OneOfAppInfoLocalizationAppCategoryAppCategoryAppCategoryAppCategoryAppCategoryAppCategory) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *AppInfosResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppInfosResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppInfosResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppInfosResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfosResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppInfosResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppInfosResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppInfosResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfosResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfosResponse := _AppInfosResponse{}

	if err = json.Unmarshal(bytes, &varAppInfosResponse); err == nil {
		*o = AppInfosResponse(varAppInfosResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfosResponse struct {
	value *AppInfosResponse
	isSet bool
}

func (v NullableAppInfosResponse) Get() *AppInfosResponse {
	return v.value
}

func (v *NullableAppInfosResponse) Set(val *AppInfosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfosResponse(val *AppInfosResponse) *NullableAppInfosResponse {
	return &NullableAppInfosResponse{value: val, isSet: true}
}

func (v NullableAppInfosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


