/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// ConnectorAllOf struct for ConnectorAllOf
type ConnectorAllOf struct {
	Metadata           *ConnectorAllOfMetadata           `json:"metadata,omitempty"`
	DeploymentLocation *ConnectorAllOfDeploymentLocation `json:"deployment_location,omitempty"`
	ConnectorTypeId    *string                           `json:"connector_type_id,omitempty"`
	ConnectorSpec      *map[string]interface{}           `json:"connector_spec,omitempty"`
	Status             *string                           `json:"status,omitempty"`
}

// NewConnectorAllOf instantiates a new ConnectorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorAllOf() *ConnectorAllOf {
	this := ConnectorAllOf{}
	return &this
}

// NewConnectorAllOfWithDefaults instantiates a new ConnectorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorAllOfWithDefaults() *ConnectorAllOf {
	this := ConnectorAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConnectorAllOf) GetMetadata() ConnectorAllOfMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConnectorAllOfMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAllOf) GetMetadataOk() (*ConnectorAllOfMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConnectorAllOf) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConnectorAllOfMetadata and assigns it to the Metadata field.
func (o *ConnectorAllOf) SetMetadata(v ConnectorAllOfMetadata) {
	o.Metadata = &v
}

// GetDeploymentLocation returns the DeploymentLocation field value if set, zero value otherwise.
func (o *ConnectorAllOf) GetDeploymentLocation() ConnectorAllOfDeploymentLocation {
	if o == nil || o.DeploymentLocation == nil {
		var ret ConnectorAllOfDeploymentLocation
		return ret
	}
	return *o.DeploymentLocation
}

// GetDeploymentLocationOk returns a tuple with the DeploymentLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAllOf) GetDeploymentLocationOk() (*ConnectorAllOfDeploymentLocation, bool) {
	if o == nil || o.DeploymentLocation == nil {
		return nil, false
	}
	return o.DeploymentLocation, true
}

// HasDeploymentLocation returns a boolean if a field has been set.
func (o *ConnectorAllOf) HasDeploymentLocation() bool {
	if o != nil && o.DeploymentLocation != nil {
		return true
	}

	return false
}

// SetDeploymentLocation gets a reference to the given ConnectorAllOfDeploymentLocation and assigns it to the DeploymentLocation field.
func (o *ConnectorAllOf) SetDeploymentLocation(v ConnectorAllOfDeploymentLocation) {
	o.DeploymentLocation = &v
}

// GetConnectorTypeId returns the ConnectorTypeId field value if set, zero value otherwise.
func (o *ConnectorAllOf) GetConnectorTypeId() string {
	if o == nil || o.ConnectorTypeId == nil {
		var ret string
		return ret
	}
	return *o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAllOf) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil || o.ConnectorTypeId == nil {
		return nil, false
	}
	return o.ConnectorTypeId, true
}

// HasConnectorTypeId returns a boolean if a field has been set.
func (o *ConnectorAllOf) HasConnectorTypeId() bool {
	if o != nil && o.ConnectorTypeId != nil {
		return true
	}

	return false
}

// SetConnectorTypeId gets a reference to the given string and assigns it to the ConnectorTypeId field.
func (o *ConnectorAllOf) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = &v
}

// GetConnectorSpec returns the ConnectorSpec field value if set, zero value otherwise.
func (o *ConnectorAllOf) GetConnectorSpec() map[string]interface{} {
	if o == nil || o.ConnectorSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ConnectorSpec
}

// GetConnectorSpecOk returns a tuple with the ConnectorSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAllOf) GetConnectorSpecOk() (*map[string]interface{}, bool) {
	if o == nil || o.ConnectorSpec == nil {
		return nil, false
	}
	return o.ConnectorSpec, true
}

// HasConnectorSpec returns a boolean if a field has been set.
func (o *ConnectorAllOf) HasConnectorSpec() bool {
	if o != nil && o.ConnectorSpec != nil {
		return true
	}

	return false
}

// SetConnectorSpec gets a reference to the given map[string]interface{} and assigns it to the ConnectorSpec field.
func (o *ConnectorAllOf) SetConnectorSpec(v map[string]interface{}) {
	o.ConnectorSpec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectorAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectorAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConnectorAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o ConnectorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.DeploymentLocation != nil {
		toSerialize["deployment_location"] = o.DeploymentLocation
	}
	if o.ConnectorTypeId != nil {
		toSerialize["connector_type_id"] = o.ConnectorTypeId
	}
	if o.ConnectorSpec != nil {
		toSerialize["connector_spec"] = o.ConnectorSpec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorAllOf struct {
	value *ConnectorAllOf
	isSet bool
}

func (v NullableConnectorAllOf) Get() *ConnectorAllOf {
	return v.value
}

func (v *NullableConnectorAllOf) Set(val *ConnectorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorAllOf(val *ConnectorAllOf) *NullableConnectorAllOf {
	return &NullableConnectorAllOf{value: val, isSet: true}
}

func (v NullableConnectorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
