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

// BetaGroupUpdateRequestDataAttributes struct for BetaGroupUpdateRequestDataAttributes
type BetaGroupUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	PublicLinkEnabled *bool `json:"publicLinkEnabled,omitempty"`
	PublicLinkLimitEnabled *bool `json:"publicLinkLimitEnabled,omitempty"`
	PublicLinkLimit *int32 `json:"publicLinkLimit,omitempty"`
	FeedbackEnabled *bool `json:"feedbackEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupUpdateRequestDataAttributes BetaGroupUpdateRequestDataAttributes

// NewBetaGroupUpdateRequestDataAttributes instantiates a new BetaGroupUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupUpdateRequestDataAttributes() *BetaGroupUpdateRequestDataAttributes {
	this := BetaGroupUpdateRequestDataAttributes{}
	return &this
}

// NewBetaGroupUpdateRequestDataAttributesWithDefaults instantiates a new BetaGroupUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupUpdateRequestDataAttributesWithDefaults() *BetaGroupUpdateRequestDataAttributes {
	this := BetaGroupUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BetaGroupUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublicLinkEnabled returns the PublicLinkEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkEnabled() bool {
	if o == nil || o.PublicLinkEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PublicLinkEnabled
}

// GetPublicLinkEnabledOk returns a tuple with the PublicLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkEnabledOk() (*bool, bool) {
	if o == nil || o.PublicLinkEnabled == nil {
		return nil, false
	}
	return o.PublicLinkEnabled, true
}

// HasPublicLinkEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkEnabled() bool {
	if o != nil && o.PublicLinkEnabled != nil {
		return true
	}

	return false
}

// SetPublicLinkEnabled gets a reference to the given bool and assigns it to the PublicLinkEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkEnabled(v bool) {
	o.PublicLinkEnabled = &v
}

// GetPublicLinkLimitEnabled returns the PublicLinkLimitEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitEnabled() bool {
	if o == nil || o.PublicLinkLimitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PublicLinkLimitEnabled
}

// GetPublicLinkLimitEnabledOk returns a tuple with the PublicLinkLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitEnabledOk() (*bool, bool) {
	if o == nil || o.PublicLinkLimitEnabled == nil {
		return nil, false
	}
	return o.PublicLinkLimitEnabled, true
}

// HasPublicLinkLimitEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkLimitEnabled() bool {
	if o != nil && o.PublicLinkLimitEnabled != nil {
		return true
	}

	return false
}

// SetPublicLinkLimitEnabled gets a reference to the given bool and assigns it to the PublicLinkLimitEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkLimitEnabled(v bool) {
	o.PublicLinkLimitEnabled = &v
}

// GetPublicLinkLimit returns the PublicLinkLimit field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimit() int32 {
	if o == nil || o.PublicLinkLimit == nil {
		var ret int32
		return ret
	}
	return *o.PublicLinkLimit
}

// GetPublicLinkLimitOk returns a tuple with the PublicLinkLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitOk() (*int32, bool) {
	if o == nil || o.PublicLinkLimit == nil {
		return nil, false
	}
	return o.PublicLinkLimit, true
}

// HasPublicLinkLimit returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkLimit() bool {
	if o != nil && o.PublicLinkLimit != nil {
		return true
	}

	return false
}

// SetPublicLinkLimit gets a reference to the given int32 and assigns it to the PublicLinkLimit field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkLimit(v int32) {
	o.PublicLinkLimit = &v
}

// GetFeedbackEnabled returns the FeedbackEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetFeedbackEnabled() bool {
	if o == nil || o.FeedbackEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FeedbackEnabled
}

// GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetFeedbackEnabledOk() (*bool, bool) {
	if o == nil || o.FeedbackEnabled == nil {
		return nil, false
	}
	return o.FeedbackEnabled, true
}

// HasFeedbackEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasFeedbackEnabled() bool {
	if o != nil && o.FeedbackEnabled != nil {
		return true
	}

	return false
}

// SetFeedbackEnabled gets a reference to the given bool and assigns it to the FeedbackEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetFeedbackEnabled(v bool) {
	o.FeedbackEnabled = &v
}

func (o BetaGroupUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublicLinkEnabled != nil {
		toSerialize["publicLinkEnabled"] = o.PublicLinkEnabled
	}
	if o.PublicLinkLimitEnabled != nil {
		toSerialize["publicLinkLimitEnabled"] = o.PublicLinkLimitEnabled
	}
	if o.PublicLinkLimit != nil {
		toSerialize["publicLinkLimit"] = o.PublicLinkLimit
	}
	if o.FeedbackEnabled != nil {
		toSerialize["feedbackEnabled"] = o.FeedbackEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaGroupUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupUpdateRequestDataAttributes := _BetaGroupUpdateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varBetaGroupUpdateRequestDataAttributes); err == nil {
		*o = BetaGroupUpdateRequestDataAttributes(varBetaGroupUpdateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "publicLinkEnabled")
		delete(additionalProperties, "publicLinkLimitEnabled")
		delete(additionalProperties, "publicLinkLimit")
		delete(additionalProperties, "feedbackEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaGroupUpdateRequestDataAttributes struct {
	value *BetaGroupUpdateRequestDataAttributes
	isSet bool
}

func (v NullableBetaGroupUpdateRequestDataAttributes) Get() *BetaGroupUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) Set(val *BetaGroupUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupUpdateRequestDataAttributes(val *BetaGroupUpdateRequestDataAttributes) *NullableBetaGroupUpdateRequestDataAttributes {
	return &NullableBetaGroupUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaGroupUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


