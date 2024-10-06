# FxInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Ticker-like name of the instrument. For monetary instruments, a currency pair is used. | [optional] 
**QuoteCurrency** | Pointer to **string** | ISO 4217 Currency symbol | [optional] 
**BaseCurrency** | Pointer to **string** | ISO 4217 Currency symbol | [optional] 
**Enabled** | Pointer to **bool** | Availability for trading | [optional] 

## Methods

### NewFxInstrument

`func NewFxInstrument() *FxInstrument`

NewFxInstrument instantiates a new FxInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxInstrumentWithDefaults

`func NewFxInstrumentWithDefaults() *FxInstrument`

NewFxInstrumentWithDefaults instantiates a new FxInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FxInstrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FxInstrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FxInstrument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FxInstrument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FxInstrument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FxInstrument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FxInstrument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FxInstrument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *FxInstrument) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *FxInstrument) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *FxInstrument) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *FxInstrument) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *FxInstrument) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *FxInstrument) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *FxInstrument) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *FxInstrument) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetEnabled

`func (o *FxInstrument) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FxInstrument) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FxInstrument) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FxInstrument) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


