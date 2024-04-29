/*
EmaoPay API

EmaoPay API.

API version: 1.0
Contact: support@emaopay.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emaopay

import (
	"encoding/json"
	"time"
)

// checks if the EmaopayMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmaopayMerchant{}

// EmaopayMerchant struct for EmaopayMerchant
type EmaopayMerchant struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewEmaopayMerchant instantiates a new EmaopayMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmaopayMerchant() *EmaopayMerchant {
	this := EmaopayMerchant{}
	return &this
}

// NewEmaopayMerchantWithDefaults instantiates a new EmaopayMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmaopayMerchantWithDefaults() *EmaopayMerchant {
	this := EmaopayMerchant{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EmaopayMerchant) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *EmaopayMerchant) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmaopayMerchant) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EmaopayMerchant) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmaopayMerchant) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EmaopayMerchant) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EmaopayMerchant) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayMerchant) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EmaopayMerchant) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EmaopayMerchant) SetUserId(v string) {
	o.UserId = &v
}

func (o EmaopayMerchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmaopayMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableEmaopayMerchant struct {
	value *EmaopayMerchant
	isSet bool
}

func (v NullableEmaopayMerchant) Get() *EmaopayMerchant {
	return v.value
}

func (v *NullableEmaopayMerchant) Set(val *EmaopayMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableEmaopayMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableEmaopayMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmaopayMerchant(val *EmaopayMerchant) *NullableEmaopayMerchant {
	return &NullableEmaopayMerchant{value: val, isSet: true}
}

func (v NullableEmaopayMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmaopayMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


