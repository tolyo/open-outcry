# Instrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Ticker-like name of the instrument. For monetary instruments, a currency pair is used. | [optional] 
**QuoteCurrency** | Pointer to **string** | ISO 4217 Currency symbol | [optional] 
**Enabled** | Pointer to **bool** | Availability for trading | [optional] 

## Methods

### NewInstrument

`func NewInstrument() *Instrument`

NewInstrument instantiates a new Instrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentWithDefaults

`func NewInstrumentWithDefaults() *Instrument`

NewInstrumentWithDefaults instantiates a new Instrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instrument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instrument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Instrument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instrument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instrument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instrument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *Instrument) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *Instrument) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *Instrument) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *Instrument) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetEnabled

`func (o *Instrument) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Instrument) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Instrument) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Instrument) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


