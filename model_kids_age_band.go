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
	"fmt"
)

// KidsAgeBand the model 'KidsAgeBand'
type KidsAgeBand string

// List of KidsAgeBand
const (
	FIVE_AND_UNDER KidsAgeBand = "FIVE_AND_UNDER"
	SIX_TO_EIGHT KidsAgeBand = "SIX_TO_EIGHT"
	NINE_TO_ELEVEN KidsAgeBand = "NINE_TO_ELEVEN"
)

func (v *KidsAgeBand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KidsAgeBand(value)
	for _, existing := range []KidsAgeBand{ "FIVE_AND_UNDER", "SIX_TO_EIGHT", "NINE_TO_ELEVEN",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KidsAgeBand", value)
}

// Ptr returns reference to KidsAgeBand value
func (v KidsAgeBand) Ptr() *KidsAgeBand {
	return &v
}

type NullableKidsAgeBand struct {
	value *KidsAgeBand
	isSet bool
}

func (v NullableKidsAgeBand) Get() *KidsAgeBand {
	return v.value
}

func (v *NullableKidsAgeBand) Set(val *KidsAgeBand) {
	v.value = val
	v.isSet = true
}

func (v NullableKidsAgeBand) IsSet() bool {
	return v.isSet
}

func (v *NullableKidsAgeBand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKidsAgeBand(val *KidsAgeBand) *NullableKidsAgeBand {
	return &NullableKidsAgeBand{value: val, isSet: true}
}

func (v NullableKidsAgeBand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKidsAgeBand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

