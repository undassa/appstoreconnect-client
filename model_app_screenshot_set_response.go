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

// AppScreenshotSetResponse struct for AppScreenshotSetResponse
type AppScreenshotSetResponse struct {
	Data AppScreenshotSet `json:"data"`
	Included *[]AppScreenshot `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotSetResponse AppScreenshotSetResponse

// NewAppScreenshotSetResponse instantiates a new AppScreenshotSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetResponse(data AppScreenshotSet, links DocumentLinks, ) *AppScreenshotSetResponse {
	this := AppScreenshotSetResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppScreenshotSetResponseWithDefaults instantiates a new AppScreenshotSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetResponseWithDefaults() *AppScreenshotSetResponse {
	this := AppScreenshotSetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotSetResponse) GetData() AppScreenshotSet {
	if o == nil  {
		var ret AppScreenshotSet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetResponse) GetDataOk() (*AppScreenshotSet, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotSetResponse) SetData(v AppScreenshotSet) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppScreenshotSetResponse) GetIncluded() []AppScreenshot {
	if o == nil || o.Included == nil {
		var ret []AppScreenshot
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetResponse) GetIncludedOk() (*[]AppScreenshot, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppScreenshotSetResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppScreenshot and assigns it to the Included field.
func (o *AppScreenshotSetResponse) SetIncluded(v []AppScreenshot) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *AppScreenshotSetResponse) GetLinks() DocumentLinks {
	if o == nil  {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppScreenshotSetResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppScreenshotSetResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppScreenshotSetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotSetResponse := _AppScreenshotSetResponse{}

	if err = json.Unmarshal(bytes, &varAppScreenshotSetResponse); err == nil {
		*o = AppScreenshotSetResponse(varAppScreenshotSetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppScreenshotSetResponse struct {
	value *AppScreenshotSetResponse
	isSet bool
}

func (v NullableAppScreenshotSetResponse) Get() *AppScreenshotSetResponse {
	return v.value
}

func (v *NullableAppScreenshotSetResponse) Set(val *AppScreenshotSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetResponse(val *AppScreenshotSetResponse) *NullableAppScreenshotSetResponse {
	return &NullableAppScreenshotSetResponse{value: val, isSet: true}
}

func (v NullableAppScreenshotSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


