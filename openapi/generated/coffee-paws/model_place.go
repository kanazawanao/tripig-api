/*
coffee-paws api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Place type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Place{}

// Place struct for Place
type Place struct {
	BusinessStatus *string `json:"businessStatus,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Name *string `json:"name,omitempty"`
	Lat *int32 `json:"lat,omitempty"`
	Lng *int32 `json:"lng,omitempty"`
	OpenNow *bool `json:"openNow,omitempty"`
	PhotoUrls []string `json:"photoUrls,omitempty"`
	PlaceId *string `json:"placeId,omitempty"`
	PriceLevel *int32 `json:"priceLevel,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	Types []string `json:"types,omitempty"`
	UserRatingsTotal *int32 `json:"userRatingsTotal,omitempty"`
	Vicinity *string `json:"vicinity,omitempty"`
}

// NewPlace instantiates a new Place object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlace() *Place {
	this := Place{}
	return &this
}

// NewPlaceWithDefaults instantiates a new Place object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceWithDefaults() *Place {
	this := Place{}
	return &this
}

// GetBusinessStatus returns the BusinessStatus field value if set, zero value otherwise.
func (o *Place) GetBusinessStatus() string {
	if o == nil || IsNil(o.BusinessStatus) {
		var ret string
		return ret
	}
	return *o.BusinessStatus
}

// GetBusinessStatusOk returns a tuple with the BusinessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetBusinessStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessStatus) {
		return nil, false
	}
	return o.BusinessStatus, true
}

// HasBusinessStatus returns a boolean if a field has been set.
func (o *Place) HasBusinessStatus() bool {
	if o != nil && !IsNil(o.BusinessStatus) {
		return true
	}

	return false
}

// SetBusinessStatus gets a reference to the given string and assigns it to the BusinessStatus field.
func (o *Place) SetBusinessStatus(v string) {
	o.BusinessStatus = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Place) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Place) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Place) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Place) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Place) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Place) SetName(v string) {
	o.Name = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *Place) GetLat() int32 {
	if o == nil || IsNil(o.Lat) {
		var ret int32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetLatOk() (*int32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *Place) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given int32 and assigns it to the Lat field.
func (o *Place) SetLat(v int32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *Place) GetLng() int32 {
	if o == nil || IsNil(o.Lng) {
		var ret int32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetLngOk() (*int32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *Place) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given int32 and assigns it to the Lng field.
func (o *Place) SetLng(v int32) {
	o.Lng = &v
}

// GetOpenNow returns the OpenNow field value if set, zero value otherwise.
func (o *Place) GetOpenNow() bool {
	if o == nil || IsNil(o.OpenNow) {
		var ret bool
		return ret
	}
	return *o.OpenNow
}

// GetOpenNowOk returns a tuple with the OpenNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetOpenNowOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenNow) {
		return nil, false
	}
	return o.OpenNow, true
}

// HasOpenNow returns a boolean if a field has been set.
func (o *Place) HasOpenNow() bool {
	if o != nil && !IsNil(o.OpenNow) {
		return true
	}

	return false
}

// SetOpenNow gets a reference to the given bool and assigns it to the OpenNow field.
func (o *Place) SetOpenNow(v bool) {
	o.OpenNow = &v
}

// GetPhotoUrls returns the PhotoUrls field value if set, zero value otherwise.
func (o *Place) GetPhotoUrls() []string {
	if o == nil || IsNil(o.PhotoUrls) {
		var ret []string
		return ret
	}
	return o.PhotoUrls
}

// GetPhotoUrlsOk returns a tuple with the PhotoUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetPhotoUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.PhotoUrls) {
		return nil, false
	}
	return o.PhotoUrls, true
}

// HasPhotoUrls returns a boolean if a field has been set.
func (o *Place) HasPhotoUrls() bool {
	if o != nil && !IsNil(o.PhotoUrls) {
		return true
	}

	return false
}

// SetPhotoUrls gets a reference to the given []string and assigns it to the PhotoUrls field.
func (o *Place) SetPhotoUrls(v []string) {
	o.PhotoUrls = v
}

// GetPlaceId returns the PlaceId field value if set, zero value otherwise.
func (o *Place) GetPlaceId() string {
	if o == nil || IsNil(o.PlaceId) {
		var ret string
		return ret
	}
	return *o.PlaceId
}

// GetPlaceIdOk returns a tuple with the PlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetPlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceId) {
		return nil, false
	}
	return o.PlaceId, true
}

// HasPlaceId returns a boolean if a field has been set.
func (o *Place) HasPlaceId() bool {
	if o != nil && !IsNil(o.PlaceId) {
		return true
	}

	return false
}

// SetPlaceId gets a reference to the given string and assigns it to the PlaceId field.
func (o *Place) SetPlaceId(v string) {
	o.PlaceId = &v
}

// GetPriceLevel returns the PriceLevel field value if set, zero value otherwise.
func (o *Place) GetPriceLevel() int32 {
	if o == nil || IsNil(o.PriceLevel) {
		var ret int32
		return ret
	}
	return *o.PriceLevel
}

// GetPriceLevelOk returns a tuple with the PriceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetPriceLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriceLevel) {
		return nil, false
	}
	return o.PriceLevel, true
}

// HasPriceLevel returns a boolean if a field has been set.
func (o *Place) HasPriceLevel() bool {
	if o != nil && !IsNil(o.PriceLevel) {
		return true
	}

	return false
}

// SetPriceLevel gets a reference to the given int32 and assigns it to the PriceLevel field.
func (o *Place) SetPriceLevel(v int32) {
	o.PriceLevel = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *Place) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *Place) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *Place) SetRating(v int32) {
	o.Rating = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *Place) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *Place) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *Place) SetTypes(v []string) {
	o.Types = v
}

// GetUserRatingsTotal returns the UserRatingsTotal field value if set, zero value otherwise.
func (o *Place) GetUserRatingsTotal() int32 {
	if o == nil || IsNil(o.UserRatingsTotal) {
		var ret int32
		return ret
	}
	return *o.UserRatingsTotal
}

// GetUserRatingsTotalOk returns a tuple with the UserRatingsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetUserRatingsTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.UserRatingsTotal) {
		return nil, false
	}
	return o.UserRatingsTotal, true
}

// HasUserRatingsTotal returns a boolean if a field has been set.
func (o *Place) HasUserRatingsTotal() bool {
	if o != nil && !IsNil(o.UserRatingsTotal) {
		return true
	}

	return false
}

// SetUserRatingsTotal gets a reference to the given int32 and assigns it to the UserRatingsTotal field.
func (o *Place) SetUserRatingsTotal(v int32) {
	o.UserRatingsTotal = &v
}

// GetVicinity returns the Vicinity field value if set, zero value otherwise.
func (o *Place) GetVicinity() string {
	if o == nil || IsNil(o.Vicinity) {
		var ret string
		return ret
	}
	return *o.Vicinity
}

// GetVicinityOk returns a tuple with the Vicinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Place) GetVicinityOk() (*string, bool) {
	if o == nil || IsNil(o.Vicinity) {
		return nil, false
	}
	return o.Vicinity, true
}

// HasVicinity returns a boolean if a field has been set.
func (o *Place) HasVicinity() bool {
	if o != nil && !IsNil(o.Vicinity) {
		return true
	}

	return false
}

// SetVicinity gets a reference to the given string and assigns it to the Vicinity field.
func (o *Place) SetVicinity(v string) {
	o.Vicinity = &v
}

func (o Place) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Place) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessStatus) {
		toSerialize["businessStatus"] = o.BusinessStatus
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !IsNil(o.OpenNow) {
		toSerialize["openNow"] = o.OpenNow
	}
	if !IsNil(o.PhotoUrls) {
		toSerialize["photoUrls"] = o.PhotoUrls
	}
	if !IsNil(o.PlaceId) {
		toSerialize["placeId"] = o.PlaceId
	}
	if !IsNil(o.PriceLevel) {
		toSerialize["priceLevel"] = o.PriceLevel
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.UserRatingsTotal) {
		toSerialize["userRatingsTotal"] = o.UserRatingsTotal
	}
	if !IsNil(o.Vicinity) {
		toSerialize["vicinity"] = o.Vicinity
	}
	return toSerialize, nil
}

type NullablePlace struct {
	value *Place
	isSet bool
}

func (v NullablePlace) Get() *Place {
	return v.value
}

func (v *NullablePlace) Set(val *Place) {
	v.value = val
	v.isSet = true
}

func (v NullablePlace) IsSet() bool {
	return v.isSet
}

func (v *NullablePlace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlace(val *Place) *NullablePlace {
	return &NullablePlace{value: val, isSet: true}
}

func (v NullablePlace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


