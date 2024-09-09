# CreateTradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instrument** | **string** | Ticker-like name of the instrument. For monetary instruments, a currency pair is used. | 
**Side** | [**TradeOrderSide**](TradeOrderSide.md) |  | 
**Type** | [**TradeOrderType**](TradeOrderType.md) |  | 
**TimeInForce** | [**TradeOrderTimeInForce**](TradeOrderTimeInForce.md) |  | 
**Amount** | **float64** |  | 
**Price** | Pointer to **float64** |  | [optional] 

## Methods

### NewCreateTradeRequest

`func NewCreateTradeRequest(instrument string, side TradeOrderSide, type_ TradeOrderType, timeInForce TradeOrderTimeInForce, amount float64, ) *CreateTradeRequest`

NewCreateTradeRequest instantiates a new CreateTradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeRequestWithDefaults

`func NewCreateTradeRequestWithDefaults() *CreateTradeRequest`

NewCreateTradeRequestWithDefaults instantiates a new CreateTradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrument

`func (o *CreateTradeRequest) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *CreateTradeRequest) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *CreateTradeRequest) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.


### GetSide

`func (o *CreateTradeRequest) GetSide() TradeOrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateTradeRequest) GetSideOk() (*TradeOrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateTradeRequest) SetSide(v TradeOrderSide)`

SetSide sets Side field to given value.


### GetType

`func (o *CreateTradeRequest) GetType() TradeOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTradeRequest) GetTypeOk() (*TradeOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTradeRequest) SetType(v TradeOrderType)`

SetType sets Type field to given value.


### GetTimeInForce

`func (o *CreateTradeRequest) GetTimeInForce() TradeOrderTimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CreateTradeRequest) GetTimeInForceOk() (*TradeOrderTimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CreateTradeRequest) SetTimeInForce(v TradeOrderTimeInForce)`

SetTimeInForce sets TimeInForce field to given value.


### GetAmount

`func (o *CreateTradeRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTradeRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTradeRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPrice

`func (o *CreateTradeRequest) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateTradeRequest) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateTradeRequest) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateTradeRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


