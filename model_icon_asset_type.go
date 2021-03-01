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

// IconAssetType the model 'IconAssetType'
type IconAssetType string

// List of IconAssetType
const (
	APP_STORE IconAssetType = "APP_STORE"
	MESSAGES_APP_STORE IconAssetType = "MESSAGES_APP_STORE"
	WATCH_APP_STORE IconAssetType = "WATCH_APP_STORE"
	TV_OS_HOME_SCREEN IconAssetType = "TV_OS_HOME_SCREEN"
	TV_OS_TOP_SHELF IconAssetType = "TV_OS_TOP_SHELF"
)

func (v *IconAssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IconAssetType(value)
	for _, existing := range []IconAssetType{ "APP_STORE", "MESSAGES_APP_STORE", "WATCH_APP_STORE", "TV_OS_HOME_SCREEN", "TV_OS_TOP_SHELF",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IconAssetType", value)
}

// Ptr returns reference to IconAssetType value
func (v IconAssetType) Ptr() *IconAssetType {
	return &v
}

type NullableIconAssetType struct {
	value *IconAssetType
	isSet bool
}

func (v NullableIconAssetType) Get() *IconAssetType {
	return v.value
}

func (v *NullableIconAssetType) Set(val *IconAssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableIconAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableIconAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconAssetType(val *IconAssetType) *NullableIconAssetType {
	return &NullableIconAssetType{value: val, isSet: true}
}

func (v NullableIconAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

