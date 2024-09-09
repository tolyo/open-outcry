# TradingAccountInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Ticker-like name of the instrument. For monetary instruments, a currency pair is used. | 
**Amount** | **float64** |  | 
**AmountReserved** | **float64** |  | 
**AmountAvailable** | **float64** |  | 
**Value** | **float64** |  | 
**Currency** | **string** | ISO 4217 Currency symbol | 

## Methods

### NewTradingAccountInstrument

`func NewTradingAccountInstrument(name string, amount float64, amountReserved float64, amountAvailable float64, value float64, currency string, ) *TradingAccountInstrument`

NewTradingAccountInstrument instantiates a new TradingAccountInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingAccountInstrumentWithDefaults

`func NewTradingAccountInstrumentWithDefaults() *TradingAccountInstrument`

NewTradingAccountInstrumentWithDefaults instantiates a new TradingAccountInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TradingAccountInstrument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TradingAccountInstrument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TradingAccountInstrument) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *TradingAccountInstrument) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TradingAccountInstrument) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TradingAccountInstrument) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetAmountReserved

`func (o *TradingAccountInstrument) GetAmountReserved() float64`

GetAmountReserved returns the AmountReserved field if non-nil, zero value otherwise.

### GetAmountReservedOk

`func (o *TradingAccountInstrument) GetAmountReservedOk() (*float64, bool)`

GetAmountReservedOk returns a tuple with the AmountReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountReserved

`func (o *TradingAccountInstrument) SetAmountReserved(v float64)`

SetAmountReserved sets AmountReserved field to given value.


### GetAmountAvailable

`func (o *TradingAccountInstrument) GetAmountAvailable() float64`

GetAmountAvailable returns the AmountAvailable field if non-nil, zero value otherwise.

### GetAmountAvailableOk

`func (o *TradingAccountInstrument) GetAmountAvailableOk() (*float64, bool)`

GetAmountAvailableOk returns a tuple with the AmountAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAvailable

`func (o *TradingAccountInstrument) SetAmountAvailable(v float64)`

SetAmountAvailable sets AmountAvailable field to given value.


### GetValue

`func (o *TradingAccountInstrument) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TradingAccountInstrument) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TradingAccountInstrument) SetValue(v float64)`

SetValue sets Value field to given value.


### GetCurrency

`func (o *TradingAccountInstrument) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TradingAccountInstrument) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TradingAccountInstrument) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


