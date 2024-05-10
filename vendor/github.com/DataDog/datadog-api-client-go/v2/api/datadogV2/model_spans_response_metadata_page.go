// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansResponseMetadataPage Paging attributes.
type SpansResponseMetadataPage struct {
	// The cursor to use to get the next results, if any. To make the next request, use the same
	// parameters with the addition of the `page[cursor]`.
	After *string `json:"after,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansResponseMetadataPage instantiates a new SpansResponseMetadataPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansResponseMetadataPage() *SpansResponseMetadataPage {
	this := SpansResponseMetadataPage{}
	return &this
}

// NewSpansResponseMetadataPageWithDefaults instantiates a new SpansResponseMetadataPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansResponseMetadataPageWithDefaults() *SpansResponseMetadataPage {
	this := SpansResponseMetadataPage{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *SpansResponseMetadataPage) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansResponseMetadataPage) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *SpansResponseMetadataPage) HasAfter() bool {
	return o != nil && o.After != nil
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *SpansResponseMetadataPage) SetAfter(v string) {
	o.After = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansResponseMetadataPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.After != nil {
		toSerialize["after"] = o.After
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansResponseMetadataPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		After *string `json:"after,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"after"})
	} else {
		return err
	}
	o.After = all.After

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
