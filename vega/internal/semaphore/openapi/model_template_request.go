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

// TemplateRequest struct for TemplateRequest
type TemplateRequest struct {
	ProjectId     *int32  `json:"project_id,omitempty"`
	InventoryId   *int32  `json:"inventory_id,omitempty"`
	RepositoryId  *int32  `json:"repository_id,omitempty"`
	EnvironmentId *int32  `json:"environment_id,omitempty"`
	ViewId        *int32  `json:"view_id,omitempty"`
	Alias         *string `json:"alias,omitempty"`
	Playbook      *string `json:"playbook,omitempty"`
	Arguments     *string `json:"arguments,omitempty"`
	Description   *string `json:"description,omitempty"`
	OverrideArgs  *bool   `json:"override_args,omitempty"`
}

// NewTemplateRequest instantiates a new TemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateRequest() *TemplateRequest {
	this := TemplateRequest{}
	return &this
}

// NewTemplateRequestWithDefaults instantiates a new TemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateRequestWithDefaults() *TemplateRequest {
	this := TemplateRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *TemplateRequest) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TemplateRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *TemplateRequest) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetInventoryId returns the InventoryId field value if set, zero value otherwise.
func (o *TemplateRequest) GetInventoryId() int32 {
	if o == nil || o.InventoryId == nil {
		var ret int32
		return ret
	}
	return *o.InventoryId
}

// GetInventoryIdOk returns a tuple with the InventoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetInventoryIdOk() (*int32, bool) {
	if o == nil || o.InventoryId == nil {
		return nil, false
	}
	return o.InventoryId, true
}

// HasInventoryId returns a boolean if a field has been set.
func (o *TemplateRequest) HasInventoryId() bool {
	if o != nil && o.InventoryId != nil {
		return true
	}

	return false
}

// SetInventoryId gets a reference to the given int32 and assigns it to the InventoryId field.
func (o *TemplateRequest) SetInventoryId(v int32) {
	o.InventoryId = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *TemplateRequest) GetRepositoryId() int32 {
	if o == nil || o.RepositoryId == nil {
		var ret int32
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetRepositoryIdOk() (*int32, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *TemplateRequest) HasRepositoryId() bool {
	if o != nil && o.RepositoryId != nil {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given int32 and assigns it to the RepositoryId field.
func (o *TemplateRequest) SetRepositoryId(v int32) {
	o.RepositoryId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *TemplateRequest) GetEnvironmentId() int32 {
	if o == nil || o.EnvironmentId == nil {
		var ret int32
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetEnvironmentIdOk() (*int32, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *TemplateRequest) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given int32 and assigns it to the EnvironmentId field.
func (o *TemplateRequest) SetEnvironmentId(v int32) {
	o.EnvironmentId = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *TemplateRequest) GetViewId() int32 {
	if o == nil || o.ViewId == nil {
		var ret int32
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetViewIdOk() (*int32, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *TemplateRequest) HasViewId() bool {
	if o != nil && o.ViewId != nil {
		return true
	}

	return false
}

// SetViewId gets a reference to the given int32 and assigns it to the ViewId field.
func (o *TemplateRequest) SetViewId(v int32) {
	o.ViewId = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *TemplateRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *TemplateRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *TemplateRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *TemplateRequest) GetPlaybook() string {
	if o == nil || o.Playbook == nil {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetPlaybookOk() (*string, bool) {
	if o == nil || o.Playbook == nil {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *TemplateRequest) HasPlaybook() bool {
	if o != nil && o.Playbook != nil {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *TemplateRequest) SetPlaybook(v string) {
	o.Playbook = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *TemplateRequest) GetArguments() string {
	if o == nil || o.Arguments == nil {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetArgumentsOk() (*string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *TemplateRequest) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *TemplateRequest) SetArguments(v string) {
	o.Arguments = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOverrideArgs returns the OverrideArgs field value if set, zero value otherwise.
func (o *TemplateRequest) GetOverrideArgs() bool {
	if o == nil || o.OverrideArgs == nil {
		var ret bool
		return ret
	}
	return *o.OverrideArgs
}

// GetOverrideArgsOk returns a tuple with the OverrideArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateRequest) GetOverrideArgsOk() (*bool, bool) {
	if o == nil || o.OverrideArgs == nil {
		return nil, false
	}
	return o.OverrideArgs, true
}

// HasOverrideArgs returns a boolean if a field has been set.
func (o *TemplateRequest) HasOverrideArgs() bool {
	if o != nil && o.OverrideArgs != nil {
		return true
	}

	return false
}

// SetOverrideArgs gets a reference to the given bool and assigns it to the OverrideArgs field.
func (o *TemplateRequest) SetOverrideArgs(v bool) {
	o.OverrideArgs = &v
}

func (o TemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.InventoryId != nil {
		toSerialize["inventory_id"] = o.InventoryId
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Playbook != nil {
		toSerialize["playbook"] = o.Playbook
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OverrideArgs != nil {
		toSerialize["override_args"] = o.OverrideArgs
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateRequest struct {
	value *TemplateRequest
	isSet bool
}

func (v NullableTemplateRequest) Get() *TemplateRequest {
	return v.value
}

func (v *NullableTemplateRequest) Set(val *TemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateRequest(val *TemplateRequest) *NullableTemplateRequest {
	return &NullableTemplateRequest{value: val, isSet: true}
}

func (v NullableTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
