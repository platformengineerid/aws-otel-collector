/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Locations struct for Locations
type Locations struct {
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// Array of items in the collection.
	Items *[]Location `json:"items,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
}

// NewLocations instantiates a new Locations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocations() *Locations {
	this := Locations{}

	return &this
}

// NewLocationsWithDefaults instantiates a new Locations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsWithDefaults() *Locations {
	this := Locations{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
func (o *Locations) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Locations) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Locations) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Locations) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *Locations) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Locations) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Locations) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Locations) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, nil is returned
func (o *Locations) GetItems() *[]Location {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Locations) GetItemsOk() (*[]Location, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *Locations) SetItems(v []Location) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *Locations) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *Locations) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Locations) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Locations) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Locations) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o Locations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableLocations struct {
	value *Locations
	isSet bool
}

func (v NullableLocations) Get() *Locations {
	return v.value
}

func (v *NullableLocations) Set(val *Locations) {
	v.value = val
	v.isSet = true
}

func (v NullableLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocations(val *Locations) *NullableLocations {
	return &NullableLocations{value: val, isSet: true}
}

func (v NullableLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
