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

// BuildAppEncryptionDeclarationLinkageRequest struct for BuildAppEncryptionDeclarationLinkageRequest
type BuildAppEncryptionDeclarationLinkageRequest struct {
	Data BuildRelationshipsAppEncryptionDeclarationData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BuildAppEncryptionDeclarationLinkageRequest BuildAppEncryptionDeclarationLinkageRequest

// NewBuildAppEncryptionDeclarationLinkageRequest instantiates a new BuildAppEncryptionDeclarationLinkageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildAppEncryptionDeclarationLinkageRequest(data BuildRelationshipsAppEncryptionDeclarationData, ) *BuildAppEncryptionDeclarationLinkageRequest {
	this := BuildAppEncryptionDeclarationLinkageRequest{}
	this.Data = data
	return &this
}

// NewBuildAppEncryptionDeclarationLinkageRequestWithDefaults instantiates a new BuildAppEncryptionDeclarationLinkageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildAppEncryptionDeclarationLinkageRequestWithDefaults() *BuildAppEncryptionDeclarationLinkageRequest {
	this := BuildAppEncryptionDeclarationLinkageRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BuildAppEncryptionDeclarationLinkageRequest) GetData() BuildRelationshipsAppEncryptionDeclarationData {
	if o == nil  {
		var ret BuildRelationshipsAppEncryptionDeclarationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildAppEncryptionDeclarationLinkageRequest) GetDataOk() (*BuildRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildAppEncryptionDeclarationLinkageRequest) SetData(v BuildRelationshipsAppEncryptionDeclarationData) {
	o.Data = v
}

func (o BuildAppEncryptionDeclarationLinkageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildAppEncryptionDeclarationLinkageRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBuildAppEncryptionDeclarationLinkageRequest := _BuildAppEncryptionDeclarationLinkageRequest{}

	if err = json.Unmarshal(bytes, &varBuildAppEncryptionDeclarationLinkageRequest); err == nil {
		*o = BuildAppEncryptionDeclarationLinkageRequest(varBuildAppEncryptionDeclarationLinkageRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildAppEncryptionDeclarationLinkageRequest struct {
	value *BuildAppEncryptionDeclarationLinkageRequest
	isSet bool
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) Get() *BuildAppEncryptionDeclarationLinkageRequest {
	return v.value
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) Set(val *BuildAppEncryptionDeclarationLinkageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildAppEncryptionDeclarationLinkageRequest(val *BuildAppEncryptionDeclarationLinkageRequest) *NullableBuildAppEncryptionDeclarationLinkageRequest {
	return &NullableBuildAppEncryptionDeclarationLinkageRequest{value: val, isSet: true}
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


