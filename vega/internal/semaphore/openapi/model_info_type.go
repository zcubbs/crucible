/*
 * API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.34
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InfoType struct for InfoType
type InfoType struct {
	Version    *string         `json:"version,omitempty"`
	UpdateBody *string         `json:"updateBody,omitempty"`
	Update     *InfoTypeUpdate `json:"update,omitempty"`
}

// NewInfoType instantiates a new InfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoType() *InfoType {
	this := InfoType{}
	return &this
}

// NewInfoTypeWithDefaults instantiates a new InfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoTypeWithDefaults() *InfoType {
	this := InfoType{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InfoType) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoType) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InfoType) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InfoType) SetVersion(v string) {
	o.Version = &v
}

// GetUpdateBody returns the UpdateBody field value if set, zero value otherwise.
func (o *InfoType) GetUpdateBody() string {
	if o == nil || o.UpdateBody == nil {
		var ret string
		return ret
	}
	return *o.UpdateBody
}

// GetUpdateBodyOk returns a tuple with the UpdateBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoType) GetUpdateBodyOk() (*string, bool) {
	if o == nil || o.UpdateBody == nil {
		return nil, false
	}
	return o.UpdateBody, true
}

// HasUpdateBody returns a boolean if a field has been set.
func (o *InfoType) HasUpdateBody() bool {
	if o != nil && o.UpdateBody != nil {
		return true
	}

	return false
}

// SetUpdateBody gets a reference to the given string and assigns it to the UpdateBody field.
func (o *InfoType) SetUpdateBody(v string) {
	o.UpdateBody = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *InfoType) GetUpdate() InfoTypeUpdate {
	if o == nil || o.Update == nil {
		var ret InfoTypeUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoType) GetUpdateOk() (*InfoTypeUpdate, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *InfoType) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given InfoTypeUpdate and assigns it to the Update field.
func (o *InfoType) SetUpdate(v InfoTypeUpdate) {
	o.Update = &v
}

func (o InfoType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.UpdateBody != nil {
		toSerialize["updateBody"] = o.UpdateBody
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	return json.Marshal(toSerialize)
}

type NullableInfoType struct {
	value *InfoType
	isSet bool
}

func (v NullableInfoType) Get() *InfoType {
	return v.value
}

func (v *NullableInfoType) Set(val *InfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoType(val *InfoType) *NullableInfoType {
	return &NullableInfoType{value: val, isSet: true}
}

func (v NullableInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
