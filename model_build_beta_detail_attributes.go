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

// BuildBetaDetailAttributes struct for BuildBetaDetailAttributes
type BuildBetaDetailAttributes struct {
	AutoNotifyEnabled *bool `json:"autoNotifyEnabled,omitempty"`
	InternalBuildState *InternalBetaState `json:"internalBuildState,omitempty"`
	ExternalBuildState *ExternalBetaState `json:"externalBuildState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildBetaDetailAttributes BuildBetaDetailAttributes

// NewBuildBetaDetailAttributes instantiates a new BuildBetaDetailAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailAttributes() *BuildBetaDetailAttributes {
	this := BuildBetaDetailAttributes{}
	return &this
}

// NewBuildBetaDetailAttributesWithDefaults instantiates a new BuildBetaDetailAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailAttributesWithDefaults() *BuildBetaDetailAttributes {
	this := BuildBetaDetailAttributes{}
	return &this
}

// GetAutoNotifyEnabled returns the AutoNotifyEnabled field value if set, zero value otherwise.
func (o *BuildBetaDetailAttributes) GetAutoNotifyEnabled() bool {
	if o == nil || o.AutoNotifyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoNotifyEnabled
}

// GetAutoNotifyEnabledOk returns a tuple with the AutoNotifyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailAttributes) GetAutoNotifyEnabledOk() (*bool, bool) {
	if o == nil || o.AutoNotifyEnabled == nil {
		return nil, false
	}
	return o.AutoNotifyEnabled, true
}

// HasAutoNotifyEnabled returns a boolean if a field has been set.
func (o *BuildBetaDetailAttributes) HasAutoNotifyEnabled() bool {
	if o != nil && o.AutoNotifyEnabled != nil {
		return true
	}

	return false
}

// SetAutoNotifyEnabled gets a reference to the given bool and assigns it to the AutoNotifyEnabled field.
func (o *BuildBetaDetailAttributes) SetAutoNotifyEnabled(v bool) {
	o.AutoNotifyEnabled = &v
}

// GetInternalBuildState returns the InternalBuildState field value if set, zero value otherwise.
func (o *BuildBetaDetailAttributes) GetInternalBuildState() InternalBetaState {
	if o == nil || o.InternalBuildState == nil {
		var ret InternalBetaState
		return ret
	}
	return *o.InternalBuildState
}

// GetInternalBuildStateOk returns a tuple with the InternalBuildState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailAttributes) GetInternalBuildStateOk() (*InternalBetaState, bool) {
	if o == nil || o.InternalBuildState == nil {
		return nil, false
	}
	return o.InternalBuildState, true
}

// HasInternalBuildState returns a boolean if a field has been set.
func (o *BuildBetaDetailAttributes) HasInternalBuildState() bool {
	if o != nil && o.InternalBuildState != nil {
		return true
	}

	return false
}

// SetInternalBuildState gets a reference to the given InternalBetaState and assigns it to the InternalBuildState field.
func (o *BuildBetaDetailAttributes) SetInternalBuildState(v InternalBetaState) {
	o.InternalBuildState = &v
}

// GetExternalBuildState returns the ExternalBuildState field value if set, zero value otherwise.
func (o *BuildBetaDetailAttributes) GetExternalBuildState() ExternalBetaState {
	if o == nil || o.ExternalBuildState == nil {
		var ret ExternalBetaState
		return ret
	}
	return *o.ExternalBuildState
}

// GetExternalBuildStateOk returns a tuple with the ExternalBuildState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailAttributes) GetExternalBuildStateOk() (*ExternalBetaState, bool) {
	if o == nil || o.ExternalBuildState == nil {
		return nil, false
	}
	return o.ExternalBuildState, true
}

// HasExternalBuildState returns a boolean if a field has been set.
func (o *BuildBetaDetailAttributes) HasExternalBuildState() bool {
	if o != nil && o.ExternalBuildState != nil {
		return true
	}

	return false
}

// SetExternalBuildState gets a reference to the given ExternalBetaState and assigns it to the ExternalBuildState field.
func (o *BuildBetaDetailAttributes) SetExternalBuildState(v ExternalBetaState) {
	o.ExternalBuildState = &v
}

func (o BuildBetaDetailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoNotifyEnabled != nil {
		toSerialize["autoNotifyEnabled"] = o.AutoNotifyEnabled
	}
	if o.InternalBuildState != nil {
		toSerialize["internalBuildState"] = o.InternalBuildState
	}
	if o.ExternalBuildState != nil {
		toSerialize["externalBuildState"] = o.ExternalBuildState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BuildBetaDetailAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBuildBetaDetailAttributes := _BuildBetaDetailAttributes{}

	if err = json.Unmarshal(bytes, &varBuildBetaDetailAttributes); err == nil {
		*o = BuildBetaDetailAttributes(varBuildBetaDetailAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "autoNotifyEnabled")
		delete(additionalProperties, "internalBuildState")
		delete(additionalProperties, "externalBuildState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildBetaDetailAttributes struct {
	value *BuildBetaDetailAttributes
	isSet bool
}

func (v NullableBuildBetaDetailAttributes) Get() *BuildBetaDetailAttributes {
	return v.value
}

func (v *NullableBuildBetaDetailAttributes) Set(val *BuildBetaDetailAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailAttributes(val *BuildBetaDetailAttributes) *NullableBuildBetaDetailAttributes {
	return &NullableBuildBetaDetailAttributes{value: val, isSet: true}
}

func (v NullableBuildBetaDetailAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

