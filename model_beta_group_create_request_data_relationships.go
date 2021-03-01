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

// BetaGroupCreateRequestDataRelationships struct for BetaGroupCreateRequestDataRelationships
type BetaGroupCreateRequestDataRelationships struct {
	App AppPreOrderCreateRequestDataRelationshipsApp `json:"app"`
	Builds *BetaGroupCreateRequestDataRelationshipsBuilds `json:"builds,omitempty"`
	BetaTesters *BetaGroupCreateRequestDataRelationshipsBetaTesters `json:"betaTesters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BetaGroupCreateRequestDataRelationships BetaGroupCreateRequestDataRelationships

// NewBetaGroupCreateRequestDataRelationships instantiates a new BetaGroupCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupCreateRequestDataRelationships(app AppPreOrderCreateRequestDataRelationshipsApp, ) *BetaGroupCreateRequestDataRelationships {
	this := BetaGroupCreateRequestDataRelationships{}
	this.App = app
	return &this
}

// NewBetaGroupCreateRequestDataRelationshipsWithDefaults instantiates a new BetaGroupCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupCreateRequestDataRelationshipsWithDefaults() *BetaGroupCreateRequestDataRelationships {
	this := BetaGroupCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *BetaGroupCreateRequestDataRelationships) GetApp() AppPreOrderCreateRequestDataRelationshipsApp {
	if o == nil  {
		var ret AppPreOrderCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetAppOk() (*AppPreOrderCreateRequestDataRelationshipsApp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BetaGroupCreateRequestDataRelationships) SetApp(v AppPreOrderCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataRelationships) GetBuilds() BetaGroupCreateRequestDataRelationshipsBuilds {
	if o == nil || o.Builds == nil {
		var ret BetaGroupCreateRequestDataRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetBuildsOk() (*BetaGroupCreateRequestDataRelationshipsBuilds, bool) {
	if o == nil || o.Builds == nil {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataRelationships) HasBuilds() bool {
	if o != nil && o.Builds != nil {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given BetaGroupCreateRequestDataRelationshipsBuilds and assigns it to the Builds field.
func (o *BetaGroupCreateRequestDataRelationships) SetBuilds(v BetaGroupCreateRequestDataRelationshipsBuilds) {
	o.Builds = &v
}

// GetBetaTesters returns the BetaTesters field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataRelationships) GetBetaTesters() BetaGroupCreateRequestDataRelationshipsBetaTesters {
	if o == nil || o.BetaTesters == nil {
		var ret BetaGroupCreateRequestDataRelationshipsBetaTesters
		return ret
	}
	return *o.BetaTesters
}

// GetBetaTestersOk returns a tuple with the BetaTesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetBetaTestersOk() (*BetaGroupCreateRequestDataRelationshipsBetaTesters, bool) {
	if o == nil || o.BetaTesters == nil {
		return nil, false
	}
	return o.BetaTesters, true
}

// HasBetaTesters returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataRelationships) HasBetaTesters() bool {
	if o != nil && o.BetaTesters != nil {
		return true
	}

	return false
}

// SetBetaTesters gets a reference to the given BetaGroupCreateRequestDataRelationshipsBetaTesters and assigns it to the BetaTesters field.
func (o *BetaGroupCreateRequestDataRelationships) SetBetaTesters(v BetaGroupCreateRequestDataRelationshipsBetaTesters) {
	o.BetaTesters = &v
}

func (o BetaGroupCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app"] = o.App
	}
	if o.Builds != nil {
		toSerialize["builds"] = o.Builds
	}
	if o.BetaTesters != nil {
		toSerialize["betaTesters"] = o.BetaTesters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaGroupCreateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varBetaGroupCreateRequestDataRelationships := _BetaGroupCreateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varBetaGroupCreateRequestDataRelationships); err == nil {
		*o = BetaGroupCreateRequestDataRelationships(varBetaGroupCreateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "builds")
		delete(additionalProperties, "betaTesters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaGroupCreateRequestDataRelationships struct {
	value *BetaGroupCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaGroupCreateRequestDataRelationships) Get() *BetaGroupCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaGroupCreateRequestDataRelationships) Set(val *BetaGroupCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupCreateRequestDataRelationships(val *BetaGroupCreateRequestDataRelationships) *NullableBetaGroupCreateRequestDataRelationships {
	return &NullableBetaGroupCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaGroupCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


