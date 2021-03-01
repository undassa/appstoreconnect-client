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

// UsersResponse struct for UsersResponse
type UsersResponse struct {
	Data []User `json:"data"`
	Included *[]App `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersResponse UsersResponse

// NewUsersResponse instantiates a new UsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersResponse(data []User, links PagedDocumentLinks, ) *UsersResponse {
	this := UsersResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewUsersResponseWithDefaults instantiates a new UsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersResponseWithDefaults() *UsersResponse {
	this := UsersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UsersResponse) GetData() []User {
	if o == nil  {
		var ret []User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetDataOk() (*[]User, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UsersResponse) SetData(v []User) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *UsersResponse) GetIncluded() []App {
	if o == nil || o.Included == nil {
		var ret []App
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetIncludedOk() (*[]App, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *UsersResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []App and assigns it to the Included field.
func (o *UsersResponse) SetIncluded(v []App) {
	o.Included = &v
}

// GetLinks returns the Links field value
func (o *UsersResponse) GetLinks() PagedDocumentLinks {
	if o == nil  {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UsersResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsersResponse) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsersResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *UsersResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o UsersResponse) MarshalJSON() ([]byte, error) {
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

func (o *UsersResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUsersResponse := _UsersResponse{}

	if err = json.Unmarshal(bytes, &varUsersResponse); err == nil {
		*o = UsersResponse(varUsersResponse)
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

type NullableUsersResponse struct {
	value *UsersResponse
	isSet bool
}

func (v NullableUsersResponse) Get() *UsersResponse {
	return v.value
}

func (v *NullableUsersResponse) Set(val *UsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersResponse(val *UsersResponse) *NullableUsersResponse {
	return &NullableUsersResponse{value: val, isSet: true}
}

func (v NullableUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


