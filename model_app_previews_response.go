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

// AppPreviewsResponse struct for AppPreviewsResponse
type AppPreviewsResponse struct {
	Data []AppPreview `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewsResponse AppPreviewsResponse

// NewAppPreviewsResponse instantiates a new AppPreviewsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewsResponse(data []AppPreview, links PagedDocumentLinks, ) *AppPreviewsResponse {
	this := AppPreviewsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPreviewsResponseWithDefaults instantiates a new AppPreviewsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewsResponseWithDefaults() *AppPreviewsResponse {
	this := AppPreviewsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewsResponse) GetData() []AppPreview {
	if o == nil  {
		var ret []AppPreview
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewsResponse) GetDataOk() (*[]AppPreview, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewsResponse) SetData(v []AppPreview) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppPreviewsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPreviewsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPreviewsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppPreviewsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppPreviewsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppPreviewsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppPreviewsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
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

func (o *AppPreviewsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewsResponse := _AppPreviewsResponse{}

	if err = json.Unmarshal(bytes, &varAppPreviewsResponse); err == nil {
		*o = AppPreviewsResponse(varAppPreviewsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewsResponse struct {
	value *AppPreviewsResponse
	isSet bool
}

func (v NullableAppPreviewsResponse) Get() *AppPreviewsResponse {
	return v.value
}

func (v *NullableAppPreviewsResponse) Set(val *AppPreviewsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewsResponse(val *AppPreviewsResponse) *NullableAppPreviewsResponse {
	return &NullableAppPreviewsResponse{value: val, isSet: true}
}

func (v NullableAppPreviewsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


