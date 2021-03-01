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

// UserInvitationCreateRequestDataAttributes struct for UserInvitationCreateRequestDataAttributes
type UserInvitationCreateRequestDataAttributes struct {
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Roles []UserRole `json:"roles"`
	AllAppsVisible *bool `json:"allAppsVisible,omitempty"`
	ProvisioningAllowed *bool `json:"provisioningAllowed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserInvitationCreateRequestDataAttributes UserInvitationCreateRequestDataAttributes

// NewUserInvitationCreateRequestDataAttributes instantiates a new UserInvitationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequestDataAttributes(email string, firstName string, lastName string, roles []UserRole, ) *UserInvitationCreateRequestDataAttributes {
	this := UserInvitationCreateRequestDataAttributes{}
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.Roles = roles
	return &this
}

// NewUserInvitationCreateRequestDataAttributesWithDefaults instantiates a new UserInvitationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestDataAttributesWithDefaults() *UserInvitationCreateRequestDataAttributes {
	this := UserInvitationCreateRequestDataAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserInvitationCreateRequestDataAttributes) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserInvitationCreateRequestDataAttributes) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *UserInvitationCreateRequestDataAttributes) GetFirstName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserInvitationCreateRequestDataAttributes) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserInvitationCreateRequestDataAttributes) GetLastName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserInvitationCreateRequestDataAttributes) SetLastName(v string) {
	o.LastName = v
}

// GetRoles returns the Roles field value
func (o *UserInvitationCreateRequestDataAttributes) GetRoles() []UserRole {
	if o == nil  {
		var ret []UserRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetRolesOk() (*[]UserRole, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UserInvitationCreateRequestDataAttributes) SetRoles(v []UserRole) {
	o.Roles = v
}

// GetAllAppsVisible returns the AllAppsVisible field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestDataAttributes) GetAllAppsVisible() bool {
	if o == nil || o.AllAppsVisible == nil {
		var ret bool
		return ret
	}
	return *o.AllAppsVisible
}

// GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetAllAppsVisibleOk() (*bool, bool) {
	if o == nil || o.AllAppsVisible == nil {
		return nil, false
	}
	return o.AllAppsVisible, true
}

// HasAllAppsVisible returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestDataAttributes) HasAllAppsVisible() bool {
	if o != nil && o.AllAppsVisible != nil {
		return true
	}

	return false
}

// SetAllAppsVisible gets a reference to the given bool and assigns it to the AllAppsVisible field.
func (o *UserInvitationCreateRequestDataAttributes) SetAllAppsVisible(v bool) {
	o.AllAppsVisible = &v
}

// GetProvisioningAllowed returns the ProvisioningAllowed field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestDataAttributes) GetProvisioningAllowed() bool {
	if o == nil || o.ProvisioningAllowed == nil {
		var ret bool
		return ret
	}
	return *o.ProvisioningAllowed
}

// GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataAttributes) GetProvisioningAllowedOk() (*bool, bool) {
	if o == nil || o.ProvisioningAllowed == nil {
		return nil, false
	}
	return o.ProvisioningAllowed, true
}

// HasProvisioningAllowed returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestDataAttributes) HasProvisioningAllowed() bool {
	if o != nil && o.ProvisioningAllowed != nil {
		return true
	}

	return false
}

// SetProvisioningAllowed gets a reference to the given bool and assigns it to the ProvisioningAllowed field.
func (o *UserInvitationCreateRequestDataAttributes) SetProvisioningAllowed(v bool) {
	o.ProvisioningAllowed = &v
}

func (o UserInvitationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if o.AllAppsVisible != nil {
		toSerialize["allAppsVisible"] = o.AllAppsVisible
	}
	if o.ProvisioningAllowed != nil {
		toSerialize["provisioningAllowed"] = o.ProvisioningAllowed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserInvitationCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varUserInvitationCreateRequestDataAttributes := _UserInvitationCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varUserInvitationCreateRequestDataAttributes); err == nil {
		*o = UserInvitationCreateRequestDataAttributes(varUserInvitationCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "allAppsVisible")
		delete(additionalProperties, "provisioningAllowed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInvitationCreateRequestDataAttributes struct {
	value *UserInvitationCreateRequestDataAttributes
	isSet bool
}

func (v NullableUserInvitationCreateRequestDataAttributes) Get() *UserInvitationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableUserInvitationCreateRequestDataAttributes) Set(val *UserInvitationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequestDataAttributes(val *UserInvitationCreateRequestDataAttributes) *NullableUserInvitationCreateRequestDataAttributes {
	return &NullableUserInvitationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

