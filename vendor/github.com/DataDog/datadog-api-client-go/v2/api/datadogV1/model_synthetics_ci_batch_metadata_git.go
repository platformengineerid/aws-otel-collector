// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsCIBatchMetadataGit Git information.
type SyntheticsCIBatchMetadataGit struct {
	// Branch name.
	Branch *string `json:"branch,omitempty"`
	// The commit SHA.
	CommitSha *string `json:"commitSha,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsCIBatchMetadataGit instantiates a new SyntheticsCIBatchMetadataGit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsCIBatchMetadataGit() *SyntheticsCIBatchMetadataGit {
	this := SyntheticsCIBatchMetadataGit{}
	return &this
}

// NewSyntheticsCIBatchMetadataGitWithDefaults instantiates a new SyntheticsCIBatchMetadataGit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsCIBatchMetadataGitWithDefaults() *SyntheticsCIBatchMetadataGit {
	this := SyntheticsCIBatchMetadataGit{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *SyntheticsCIBatchMetadataGit) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCIBatchMetadataGit) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *SyntheticsCIBatchMetadataGit) HasBranch() bool {
	return o != nil && o.Branch != nil
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *SyntheticsCIBatchMetadataGit) SetBranch(v string) {
	o.Branch = &v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *SyntheticsCIBatchMetadataGit) GetCommitSha() string {
	if o == nil || o.CommitSha == nil {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCIBatchMetadataGit) GetCommitShaOk() (*string, bool) {
	if o == nil || o.CommitSha == nil {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *SyntheticsCIBatchMetadataGit) HasCommitSha() bool {
	return o != nil && o.CommitSha != nil
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *SyntheticsCIBatchMetadataGit) SetCommitSha(v string) {
	o.CommitSha = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsCIBatchMetadataGit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.CommitSha != nil {
		toSerialize["commitSha"] = o.CommitSha
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsCIBatchMetadataGit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branch    *string `json:"branch,omitempty"`
		CommitSha *string `json:"commitSha,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch", "commitSha"})
	} else {
		return err
	}
	o.Branch = all.Branch
	o.CommitSha = all.CommitSha

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
