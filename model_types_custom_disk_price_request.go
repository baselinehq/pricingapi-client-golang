/*
Pricing API

Provides API level access to the Costgraph Services.

API version: 1.0
Contact: contact@baselinehq.cloud
*/

package pricing_api_client

import (
	"encoding/json"
)

// checks if the TypesCustomDiskPriceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesCustomDiskPriceRequest{}

// TypesCustomDiskPriceRequest struct for TypesCustomDiskPriceRequest
type TypesCustomDiskPriceRequest struct {
	Entries []SchemaDiskPricingsRow `json:"entries"`
}

// NewTypesCustomDiskPriceRequest instantiates a new TypesCustomDiskPriceRequest object
func NewTypesCustomDiskPriceRequest() *TypesCustomDiskPriceRequest {
	this := TypesCustomDiskPriceRequest{}
	return &this
}

// GetEntries returns the Entries field value.
func (o *TypesCustomDiskPriceRequest) GetEntries() []SchemaDiskPricingsRow {
	if o == nil {
		var ret []SchemaDiskPricingsRow
		return ret
	}
	return o.Entries
}

// SetEntries sets the Entries field.
func (o *TypesCustomDiskPriceRequest) SetEntries(v []SchemaDiskPricingsRow) {
	o.Entries = v
}

func (o TypesCustomDiskPriceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesCustomDiskPriceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entries"] = o.Entries
	return toSerialize, nil
}

type NullableTypesCustomDiskPriceRequest struct {
	value *TypesCustomDiskPriceRequest
	isSet bool
}

func (v NullableTypesCustomDiskPriceRequest) Get() *TypesCustomDiskPriceRequest {
	return v.value
}

func (v *NullableTypesCustomDiskPriceRequest) Set(val *TypesCustomDiskPriceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesCustomDiskPriceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesCustomDiskPriceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesCustomDiskPriceRequest(val *TypesCustomDiskPriceRequest) *NullableTypesCustomDiskPriceRequest {
	return &NullableTypesCustomDiskPriceRequest{value: val, isSet: true}
}

func (v NullableTypesCustomDiskPriceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesCustomDiskPriceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
