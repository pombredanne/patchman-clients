/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PackagesResponsePackageList struct for PackagesResponsePackageList
type PackagesResponsePackageList struct {
	Summary *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	SourcePackage *string `json:"source_package,omitempty"`
	PackageList *[]string `json:"package_list,omitempty"`
	Repositories *[]PackagesResponseRepositories `json:"repositories,omitempty"`
}

// NewPackagesResponsePackageList instantiates a new PackagesResponsePackageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesResponsePackageList() *PackagesResponsePackageList {
	this := PackagesResponsePackageList{}
	return &this
}

// NewPackagesResponsePackageListWithDefaults instantiates a new PackagesResponsePackageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesResponsePackageListWithDefaults() *PackagesResponsePackageList {
	this := PackagesResponsePackageList{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PackagesResponsePackageList) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponsePackageList) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PackagesResponsePackageList) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PackagesResponsePackageList) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackagesResponsePackageList) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponsePackageList) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackagesResponsePackageList) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackagesResponsePackageList) SetDescription(v string) {
	o.Description = &v
}

// GetSourcePackage returns the SourcePackage field value if set, zero value otherwise.
func (o *PackagesResponsePackageList) GetSourcePackage() string {
	if o == nil || o.SourcePackage == nil {
		var ret string
		return ret
	}
	return *o.SourcePackage
}

// GetSourcePackageOk returns a tuple with the SourcePackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponsePackageList) GetSourcePackageOk() (*string, bool) {
	if o == nil || o.SourcePackage == nil {
		return nil, false
	}
	return o.SourcePackage, true
}

// HasSourcePackage returns a boolean if a field has been set.
func (o *PackagesResponsePackageList) HasSourcePackage() bool {
	if o != nil && o.SourcePackage != nil {
		return true
	}

	return false
}

// SetSourcePackage gets a reference to the given string and assigns it to the SourcePackage field.
func (o *PackagesResponsePackageList) SetSourcePackage(v string) {
	o.SourcePackage = &v
}

// GetPackageList returns the PackageList field value if set, zero value otherwise.
func (o *PackagesResponsePackageList) GetPackageList() []string {
	if o == nil || o.PackageList == nil {
		var ret []string
		return ret
	}
	return *o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponsePackageList) GetPackageListOk() (*[]string, bool) {
	if o == nil || o.PackageList == nil {
		return nil, false
	}
	return o.PackageList, true
}

// HasPackageList returns a boolean if a field has been set.
func (o *PackagesResponsePackageList) HasPackageList() bool {
	if o != nil && o.PackageList != nil {
		return true
	}

	return false
}

// SetPackageList gets a reference to the given []string and assigns it to the PackageList field.
func (o *PackagesResponsePackageList) SetPackageList(v []string) {
	o.PackageList = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *PackagesResponsePackageList) GetRepositories() []PackagesResponseRepositories {
	if o == nil || o.Repositories == nil {
		var ret []PackagesResponseRepositories
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponsePackageList) GetRepositoriesOk() (*[]PackagesResponseRepositories, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *PackagesResponsePackageList) HasRepositories() bool {
	if o != nil && o.Repositories != nil {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []PackagesResponseRepositories and assigns it to the Repositories field.
func (o *PackagesResponsePackageList) SetRepositories(v []PackagesResponseRepositories) {
	o.Repositories = &v
}

func (o PackagesResponsePackageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.SourcePackage != nil {
		toSerialize["source_package"] = o.SourcePackage
	}
	if o.PackageList != nil {
		toSerialize["package_list"] = o.PackageList
	}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesResponsePackageList struct {
	value *PackagesResponsePackageList
	isSet bool
}

func (v NullablePackagesResponsePackageList) Get() *PackagesResponsePackageList {
	return v.value
}

func (v *NullablePackagesResponsePackageList) Set(val *PackagesResponsePackageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesResponsePackageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesResponsePackageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesResponsePackageList(val *PackagesResponsePackageList) *NullablePackagesResponsePackageList {
	return &NullablePackagesResponsePackageList{value: val, isSet: true}
}

func (v NullablePackagesResponsePackageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesResponsePackageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


