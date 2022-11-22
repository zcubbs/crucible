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

// ScheduleRequest struct for ScheduleRequest
type ScheduleRequest struct {
	Id         *int32  `json:"id,omitempty"`
	CronFormat *string `json:"cron_format,omitempty"`
	ProjectId  *int32  `json:"project_id,omitempty"`
	TemplateId *int32  `json:"template_id,omitempty"`
}

// NewScheduleRequest instantiates a new ScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleRequest() *ScheduleRequest {
	this := ScheduleRequest{}
	return &this
}

// NewScheduleRequestWithDefaults instantiates a new ScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleRequestWithDefaults() *ScheduleRequest {
	this := ScheduleRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleRequest) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRequest) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ScheduleRequest) SetId(v int32) {
	o.Id = &v
}

// GetCronFormat returns the CronFormat field value if set, zero value otherwise.
func (o *ScheduleRequest) GetCronFormat() string {
	if o == nil || o.CronFormat == nil {
		var ret string
		return ret
	}
	return *o.CronFormat
}

// GetCronFormatOk returns a tuple with the CronFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRequest) GetCronFormatOk() (*string, bool) {
	if o == nil || o.CronFormat == nil {
		return nil, false
	}
	return o.CronFormat, true
}

// HasCronFormat returns a boolean if a field has been set.
func (o *ScheduleRequest) HasCronFormat() bool {
	if o != nil && o.CronFormat != nil {
		return true
	}

	return false
}

// SetCronFormat gets a reference to the given string and assigns it to the CronFormat field.
func (o *ScheduleRequest) SetCronFormat(v string) {
	o.CronFormat = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ScheduleRequest) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ScheduleRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ScheduleRequest) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ScheduleRequest) GetTemplateId() int32 {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRequest) GetTemplateIdOk() (*int32, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ScheduleRequest) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ScheduleRequest) SetTemplateId(v int32) {
	o.TemplateId = &v
}

func (o ScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CronFormat != nil {
		toSerialize["cron_format"] = o.CronFormat
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.TemplateId != nil {
		toSerialize["template_id"] = o.TemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleRequest struct {
	value *ScheduleRequest
	isSet bool
}

func (v NullableScheduleRequest) Get() *ScheduleRequest {
	return v.value
}

func (v *NullableScheduleRequest) Set(val *ScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleRequest(val *ScheduleRequest) *NullableScheduleRequest {
	return &NullableScheduleRequest{value: val, isSet: true}
}

func (v NullableScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
