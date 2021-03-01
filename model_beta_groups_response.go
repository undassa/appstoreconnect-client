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

// BetaGroupsResponse struct for BetaGroupsResponse
type BetaGroupsResponse struct {
	Data []BetaGroup `json:"data"`
	Included *[]OneOfAppBuildBetaTester `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupsResponse BetaGroupsResponse

// NewBetaGroupsResponse instantiates a new BetaGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupsResponse(data []BetaGroup, links PagedDocumentLinks, ) *BetaGroupsResponse {
	this := BetaGroupsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaGroupsResponseWithDefaults instantiates a new BetaGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupsResponseWithDefaults() *BetaGroupsResponse {
	this := BetaGroupsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupsResponse) GetData() []BetaGroup {
	if o == nil  {
		var ret []BetaGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupsResponse) GetDataOk() (*[]BetaGroup, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaGroupsResponse) SetData(v []BetaGroup) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaGroupsResponse) GetIncluded() []OneOfAppBuildBetaTester {
	if o == nil || o.Included == nil {
		var ret []OneOfAppBuildBetaTester
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupsResponse) GetIncludedOk() (*[]OneOfAppBuildBetaTester, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaGroupsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []OneOfAppBuildBetaTester and assigns it to the Included field.
func (o *BetaGroupsResponse) SetIncluded(v []OneOfAppBuildBetaTester) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *BetaGroupsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroupsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroupsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaGroupsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaGroupsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaGroupsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaGroupsResponse) MarshalJSON() ([]byte, error) {
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

func (o *BetaGroupsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupsResponse := _BetaGroupsResponse{}

	if err = json.Unmarshal(bytes, &varBetaGroupsResponse); err == nil {
		*o = BetaGroupsResponse(varBetaGroupsResponse)
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

type NullableBetaGroupsResponse struct {
	value *BetaGroupsResponse
	isSet bool
}

func (v NullableBetaGroupsResponse) Get() *BetaGroupsResponse {
	return v.value
}

func (v *NullableBetaGroupsResponse) Set(val *BetaGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupsResponse(val *BetaGroupsResponse) *NullableBetaGroupsResponse {
	return &NullableBetaGroupsResponse{value: val, isSet: true}
}

func (v NullableBetaGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

