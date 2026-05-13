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

// checks if the TypesCustomDiskPricingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesCustomDiskPricingResponse{}

// TypesCustomDiskPricingResponse struct for TypesCustomDiskPricingResponse
type TypesCustomDiskPricingResponse struct {
	Status  *string                 `json:"status,omitempty"`
	Entries []SchemaDiskPricingsRow `json:"entries,omitempty"`
}

// NewTypesCustomDiskPricingResponse instantiates a new TypesCustomDiskPricingResponse object
func NewTypesCustomDiskPricingResponse() *TypesCustomDiskPricingResponse {
	this := TypesCustomDiskPricingResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TypesCustomDiskPricingResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise.
func (o *TypesCustomDiskPricingResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TypesCustomDiskPricingResponse) HasStatus() bool {
	return o != nil && !IsNil(o.Status)
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TypesCustomDiskPricingResponse) SetStatus(v string) {
	o.Status = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *TypesCustomDiskPricingResponse) GetEntries() []SchemaDiskPricingsRow {
	if o == nil {
		var ret []SchemaDiskPricingsRow
		return ret
	}
	return o.Entries
}

// SetEntries sets the Entries field.
func (o *TypesCustomDiskPricingResponse) SetEntries(v []SchemaDiskPricingsRow) {
	o.Entries = v
}

func (o TypesCustomDiskPricingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesCustomDiskPricingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableTypesCustomDiskPricingResponse struct {
	value *TypesCustomDiskPricingResponse
	isSet bool
}

func (v NullableTypesCustomDiskPricingResponse) Get() *TypesCustomDiskPricingResponse {
	return v.value
}

func (v *NullableTypesCustomDiskPricingResponse) Set(val *TypesCustomDiskPricingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesCustomDiskPricingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesCustomDiskPricingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesCustomDiskPricingResponse(val *TypesCustomDiskPricingResponse) *NullableTypesCustomDiskPricingResponse {
	return &NullableTypesCustomDiskPricingResponse{value: val, isSet: true}
}

func (v NullableTypesCustomDiskPricingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesCustomDiskPricingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
