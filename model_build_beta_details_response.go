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

// BuildBetaDetailsResponse struct for BuildBetaDetailsResponse
type BuildBetaDetailsResponse struct {
	Data []BuildBetaDetail `json:"data"`
	Included *[]Build `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildBetaDetailsResponse BuildBetaDetailsResponse

// NewBuildBetaDetailsResponse instantiates a new BuildBetaDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailsResponse(data []BuildBetaDetail, links PagedDocumentLinks, ) *BuildBetaDetailsResponse {
	this := BuildBetaDetailsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildBetaDetailsResponseWithDefaults instantiates a new BuildBetaDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailsResponseWithDefaults() *BuildBetaDetailsResponse {
	this := BuildBetaDetailsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaDetailsResponse) GetData() []BuildBetaDetail {
	if o == nil  {
		var ret []BuildBetaDetail
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailsResponse) GetDataOk() (*[]BuildBetaDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildBetaDetailsResponse) SetData(v []BuildBetaDetail) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BuildBetaDetailsResponse) GetIncluded() []Build {
	if o == nil || o.Included == nil {
		var ret []Build
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailsResponse) GetIncludedOk() (*[]Build, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BuildBetaDetailsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Build and assigns it to the Included field.
func (o *BuildBetaDetailsResponse) SetIncluded(v []Build) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *BuildBetaDetailsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildBetaDetailsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildBetaDetailsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildBetaDetailsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildBetaDetailsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BuildBetaDetailsResponse) MarshalJSON() ([]byte, error) {
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

func (o *BuildBetaDetailsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBuildBetaDetailsResponse := _BuildBetaDetailsResponse{}

	if err = json.Unmarshal(bytes, &varBuildBetaDetailsResponse); err == nil {
		*o = BuildBetaDetailsResponse(varBuildBetaDetailsResponse)
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

type NullableBuildBetaDetailsResponse struct {
	value *BuildBetaDetailsResponse
	isSet bool
}

func (v NullableBuildBetaDetailsResponse) Get() *BuildBetaDetailsResponse {
	return v.value
}

func (v *NullableBuildBetaDetailsResponse) Set(val *BuildBetaDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailsResponse(val *BuildBetaDetailsResponse) *NullableBuildBetaDetailsResponse {
	return &NullableBuildBetaDetailsResponse{value: val, isSet: true}
}

func (v NullableBuildBetaDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


