/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.56.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferIntentCreate Represents a transfer intent within Transfer UI.
type TransferIntentCreate struct {
	// Plaid's unique identifier for the transfer intent object.
	Id string `json:"id"`
	// The datetime the transfer was created. This will be of the form `2006-01-02T15:04:05Z`.
	Created time.Time `json:"created"`
	// The status of the transfer intent.  - `PENDING` – The transfer intent is pending. - `SUCCEEDED` – The transfer intent was successfully created. - `FAILED` – The transfer intent was unable to be created.
	Status string `json:"status"`
	// The Plaid `account_id` for the account that will be debited or credited. Returned only if `account_id` was set on intent creation.
	AccountId NullableString `json:"account_id,omitempty"`
	// Plaid’s unique identifier for the origination account for the intent. If not provided, the default account will be used.
	OriginationAccountId string `json:"origination_account_id"`
	// The amount of the transfer (decimal string with two digits of precision e.g. “10.00”).
	Amount string `json:"amount"`
	Mode TransferIntentCreateMode `json:"mode"`
	AchClass ACHClass `json:"ach_class"`
	User TransferUserInResponse `json:"user"`
	// A description for the underlying transfer. Maximum of 8 characters.
	Description string `json:"description"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentCreate TransferIntentCreate

// NewTransferIntentCreate instantiates a new TransferIntentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentCreate(id string, created time.Time, status string, originationAccountId string, amount string, mode TransferIntentCreateMode, achClass ACHClass, user TransferUserInResponse, description string) *TransferIntentCreate {
	this := TransferIntentCreate{}
	this.Id = id
	this.Created = created
	this.Status = status
	this.OriginationAccountId = originationAccountId
	this.Amount = amount
	this.Mode = mode
	this.AchClass = achClass
	this.User = user
	this.Description = description
	return &this
}

// NewTransferIntentCreateWithDefaults instantiates a new TransferIntentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentCreateWithDefaults() *TransferIntentCreate {
	this := TransferIntentCreate{}
	return &this
}

// GetId returns the Id field value
func (o *TransferIntentCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferIntentCreate) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TransferIntentCreate) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferIntentCreate) SetCreated(v time.Time) {
	o.Created = v
}

// GetStatus returns the Status field value
func (o *TransferIntentCreate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferIntentCreate) SetStatus(v string) {
	o.Status = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreate) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreate) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreate) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransferIntentCreate) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransferIntentCreate) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransferIntentCreate) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *TransferIntentCreate) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *TransferIntentCreate) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetAmount returns the Amount field value
func (o *TransferIntentCreate) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferIntentCreate) SetAmount(v string) {
	o.Amount = v
}

// GetMode returns the Mode field value
func (o *TransferIntentCreate) GetMode() TransferIntentCreateMode {
	if o == nil {
		var ret TransferIntentCreateMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetModeOk() (*TransferIntentCreateMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TransferIntentCreate) SetMode(v TransferIntentCreateMode) {
	o.Mode = v
}

// GetAchClass returns the AchClass field value
func (o *TransferIntentCreate) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferIntentCreate) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetUser returns the User field value
func (o *TransferIntentCreate) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferIntentCreate) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetDescription returns the Description field value
func (o *TransferIntentCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferIntentCreate) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreate) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreate) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferIntentCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferIntentCreate) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o TransferIntentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentCreate) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentCreate := _TransferIntentCreate{}

	if err = json.Unmarshal(bytes, &varTransferIntentCreate); err == nil {
		*o = TransferIntentCreate(varTransferIntentCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "status")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "user")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentCreate struct {
	value *TransferIntentCreate
	isSet bool
}

func (v NullableTransferIntentCreate) Get() *TransferIntentCreate {
	return v.value
}

func (v *NullableTransferIntentCreate) Set(val *TransferIntentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreate(val *TransferIntentCreate) *NullableTransferIntentCreate {
	return &NullableTransferIntentCreate{value: val, isSet: true}
}

func (v NullableTransferIntentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


