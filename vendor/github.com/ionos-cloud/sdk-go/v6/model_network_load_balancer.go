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

// NetworkLoadBalancer struct for NetworkLoadBalancer
type NetworkLoadBalancer struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                        `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata     `json:"metadata,omitempty"`
	Properties *NetworkLoadBalancerProperties `json:"properties"`
	Entities   *NetworkLoadBalancerEntities   `json:"entities,omitempty"`
}

// NewNetworkLoadBalancer instantiates a new NetworkLoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancer(properties NetworkLoadBalancerProperties) *NetworkLoadBalancer {
	this := NetworkLoadBalancer{}

	this.Properties = &properties

	return &this
}

// NewNetworkLoadBalancerWithDefaults instantiates a new NetworkLoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerWithDefaults() *NetworkLoadBalancer {
	this := NetworkLoadBalancer{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkLoadBalancer) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NetworkLoadBalancer) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *NetworkLoadBalancer) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NetworkLoadBalancer) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkLoadBalancer) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *NetworkLoadBalancer) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DatacenterElementMetadata will be returned
func (o *NetworkLoadBalancer) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkLoadBalancer) SetMetadata(v DatacenterElementMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for NetworkLoadBalancerProperties will be returned
func (o *NetworkLoadBalancer) GetProperties() *NetworkLoadBalancerProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetPropertiesOk() (*NetworkLoadBalancerProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *NetworkLoadBalancer) SetProperties(v NetworkLoadBalancerProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for NetworkLoadBalancerEntities will be returned
func (o *NetworkLoadBalancer) GetEntities() *NetworkLoadBalancerEntities {
	if o == nil {
		return nil
	}

	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancer) GetEntitiesOk() (*NetworkLoadBalancerEntities, bool) {
	if o == nil {
		return nil, false
	}

	return o.Entities, true
}

// SetEntities sets field value
func (o *NetworkLoadBalancer) SetEntities(v NetworkLoadBalancerEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *NetworkLoadBalancer) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

func (o NetworkLoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkLoadBalancer struct {
	value *NetworkLoadBalancer
	isSet bool
}

func (v NullableNetworkLoadBalancer) Get() *NetworkLoadBalancer {
	return v.value
}

func (v *NullableNetworkLoadBalancer) Set(val *NetworkLoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancer(val *NetworkLoadBalancer) *NullableNetworkLoadBalancer {
	return &NullableNetworkLoadBalancer{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
