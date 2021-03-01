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

// PerfPowerMetricsResponse struct for PerfPowerMetricsResponse
type PerfPowerMetricsResponse struct {
	Data []PerfPowerMetric `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PerfPowerMetricsResponse PerfPowerMetricsResponse

// NewPerfPowerMetricsResponse instantiates a new PerfPowerMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfPowerMetricsResponse(data []PerfPowerMetric, links PagedDocumentLinks, ) *PerfPowerMetricsResponse {
	this := PerfPowerMetricsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPerfPowerMetricsResponseWithDefaults instantiates a new PerfPowerMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfPowerMetricsResponseWithDefaults() *PerfPowerMetricsResponse {
	this := PerfPowerMetricsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PerfPowerMetricsResponse) GetData() []PerfPowerMetric {
	if o == nil  {
		var ret []PerfPowerMetric
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricsResponse) GetDataOk() (*[]PerfPowerMetric, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PerfPowerMetricsResponse) SetData(v []PerfPowerMetric) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *PerfPowerMetricsResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PerfPowerMetricsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PerfPowerMetricsResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PerfPowerMetricsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *PerfPowerMetricsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o PerfPowerMetricsResponse) MarshalJSON() ([]byte, error) {
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

func (o *PerfPowerMetricsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPerfPowerMetricsResponse := _PerfPowerMetricsResponse{}

	if err = json.Unmarshal(bytes, &varPerfPowerMetricsResponse); err == nil {
		*o = PerfPowerMetricsResponse(varPerfPowerMetricsResponse)
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

type NullablePerfPowerMetricsResponse struct {
	value *PerfPowerMetricsResponse
	isSet bool
}

func (v NullablePerfPowerMetricsResponse) Get() *PerfPowerMetricsResponse {
	return v.value
}

func (v *NullablePerfPowerMetricsResponse) Set(val *PerfPowerMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfPowerMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfPowerMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfPowerMetricsResponse(val *PerfPowerMetricsResponse) *NullablePerfPowerMetricsResponse {
	return &NullablePerfPowerMetricsResponse{value: val, isSet: true}
}

func (v NullablePerfPowerMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfPowerMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


