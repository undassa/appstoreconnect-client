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

// BetaGroupBuildsLinkagesResponse struct for BetaGroupBuildsLinkagesResponse
type BetaGroupBuildsLinkagesResponse struct {
	Data []AppStoreVersionRelationshipsBuildData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupBuildsLinkagesResponse BetaGroupBuildsLinkagesResponse

// NewBetaGroupBuildsLinkagesResponse instantiates a new BetaGroupBuildsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupBuildsLinkagesResponse(data []AppStoreVersionRelationshipsBuildData, links PagedDocumentLinks, ) *BetaGroupBuildsLinkagesResponse {
	this := BetaGroupBuildsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaGroupBuildsLinkagesResponseWithDefaults instantiates a new BetaGroupBuildsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupBuildsLinkagesResponseWithDefaults() *BetaGroupBuildsLinkagesResponse {
	this := BetaGroupBuildsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupBuildsLinkagesResponse) GetData() []AppStoreVersionRelationshipsBuildData {
	if o == nil  {
		var ret []AppStoreVersionRelationshipsBuildData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBuildsLinkagesResponse) GetDataOk() (*[]AppStoreVersionRelationshipsBuildData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaGroupBuildsLinkagesResponse) SetData(v []AppStoreVersionRelationshipsBuildData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaGroupBuildsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBuildsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroupBuildsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaGroupBuildsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupBuildsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaGroupBuildsLinkagesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaGroupBuildsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaGroupBuildsLinkagesResponse) MarshalJSON() ([]byte, error) {
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

func (o *BetaGroupBuildsLinkagesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupBuildsLinkagesResponse := _BetaGroupBuildsLinkagesResponse{}

	if err = json.Unmarshal(bytes, &varBetaGroupBuildsLinkagesResponse); err == nil {
		*o = BetaGroupBuildsLinkagesResponse(varBetaGroupBuildsLinkagesResponse)
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

type NullableBetaGroupBuildsLinkagesResponse struct {
	value *BetaGroupBuildsLinkagesResponse
	isSet bool
}

func (v NullableBetaGroupBuildsLinkagesResponse) Get() *BetaGroupBuildsLinkagesResponse {
	return v.value
}

func (v *NullableBetaGroupBuildsLinkagesResponse) Set(val *BetaGroupBuildsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupBuildsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupBuildsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupBuildsLinkagesResponse(val *BetaGroupBuildsLinkagesResponse) *NullableBetaGroupBuildsLinkagesResponse {
	return &NullableBetaGroupBuildsLinkagesResponse{value: val, isSet: true}
}

func (v NullableBetaGroupBuildsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupBuildsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


