# TradingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Instruments** | [**[]TradingAccountInstrument**](TradingAccountInstrument.md) |  | 

## Methods

### NewTradingAccount

`func NewTradingAccount(id string, instruments []TradingAccountInstrument, ) *TradingAccount`

NewTradingAccount instantiates a new TradingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingAccountWithDefaults

`func NewTradingAccountWithDefaults() *TradingAccount`

NewTradingAccountWithDefaults instantiates a new TradingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TradingAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradingAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradingAccount) SetId(v string)`

SetId sets Id field to given value.


### GetInstruments

`func (o *TradingAccount) GetInstruments() []TradingAccountInstrument`

GetInstruments returns the Instruments field if non-nil, zero value otherwise.

### GetInstrumentsOk

`func (o *TradingAccount) GetInstrumentsOk() (*[]TradingAccountInstrument, bool)`

GetInstrumentsOk returns a tuple with the Instruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruments

`func (o *TradingAccount) SetInstruments(v []TradingAccountInstrument)`

SetInstruments sets Instruments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


