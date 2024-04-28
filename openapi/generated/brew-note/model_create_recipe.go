/*
brew-note api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateRecipe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecipe{}

// CreateRecipe struct for CreateRecipe
type CreateRecipe struct {
	ExtractionEquipment string `json:"extractionEquipment"`
	CoffeeType string `json:"coffeeType"`
	Steps []RecipeStep `json:"steps"`
}

// NewCreateRecipe instantiates a new CreateRecipe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecipe(extractionEquipment string, coffeeType string, steps []RecipeStep) *CreateRecipe {
	this := CreateRecipe{}
	this.ExtractionEquipment = extractionEquipment
	this.CoffeeType = coffeeType
	this.Steps = steps
	return &this
}

// NewCreateRecipeWithDefaults instantiates a new CreateRecipe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecipeWithDefaults() *CreateRecipe {
	this := CreateRecipe{}
	return &this
}

// GetExtractionEquipment returns the ExtractionEquipment field value
func (o *CreateRecipe) GetExtractionEquipment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtractionEquipment
}

// GetExtractionEquipmentOk returns a tuple with the ExtractionEquipment field value
// and a boolean to check if the value has been set.
func (o *CreateRecipe) GetExtractionEquipmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtractionEquipment, true
}

// SetExtractionEquipment sets field value
func (o *CreateRecipe) SetExtractionEquipment(v string) {
	o.ExtractionEquipment = v
}

// GetCoffeeType returns the CoffeeType field value
func (o *CreateRecipe) GetCoffeeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CoffeeType
}

// GetCoffeeTypeOk returns a tuple with the CoffeeType field value
// and a boolean to check if the value has been set.
func (o *CreateRecipe) GetCoffeeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoffeeType, true
}

// SetCoffeeType sets field value
func (o *CreateRecipe) SetCoffeeType(v string) {
	o.CoffeeType = v
}

// GetSteps returns the Steps field value
func (o *CreateRecipe) GetSteps() []RecipeStep {
	if o == nil {
		var ret []RecipeStep
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *CreateRecipe) GetStepsOk() ([]RecipeStep, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *CreateRecipe) SetSteps(v []RecipeStep) {
	o.Steps = v
}

func (o CreateRecipe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecipe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extractionEquipment"] = o.ExtractionEquipment
	toSerialize["coffeeType"] = o.CoffeeType
	toSerialize["steps"] = o.Steps
	return toSerialize, nil
}

type NullableCreateRecipe struct {
	value *CreateRecipe
	isSet bool
}

func (v NullableCreateRecipe) Get() *CreateRecipe {
	return v.value
}

func (v *NullableCreateRecipe) Set(val *CreateRecipe) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecipe) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecipe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecipe(val *CreateRecipe) *NullableCreateRecipe {
	return &NullableCreateRecipe{value: val, isSet: true}
}

func (v NullableCreateRecipe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecipe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


