# TradeOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Instrument** | Pointer to **string** | Ticker-like name of the instrument. For monetary instruments, a currency pair is used. | [optional] 
**Side** | Pointer to [**TradeOrderSide**](TradeOrderSide.md) |  | [optional] 
**Type** | Pointer to [**TradeOrderType**](TradeOrderType.md) |  | [optional] 
**TimeInForce** | Pointer to [**TradeOrderTimeInForce**](TradeOrderTimeInForce.md) |  | [optional] 
**Status** | Pointer to [**TradeOrderStatus**](TradeOrderStatus.md) |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**OpenAmount** | Pointer to **float64** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTradeOrder

`func NewTradeOrder() *TradeOrder`

NewTradeOrder instantiates a new TradeOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeOrderWithDefaults

`func NewTradeOrderWithDefaults() *TradeOrder`

NewTradeOrderWithDefaults instantiates a new TradeOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TradeOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradeOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradeOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TradeOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstrument

`func (o *TradeOrder) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *TradeOrder) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *TradeOrder) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *TradeOrder) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetSide

`func (o *TradeOrder) GetSide() TradeOrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *TradeOrder) GetSideOk() (*TradeOrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *TradeOrder) SetSide(v TradeOrderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *TradeOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *TradeOrder) GetType() TradeOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TradeOrder) GetTypeOk() (*TradeOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TradeOrder) SetType(v TradeOrderType)`

SetType sets Type field to given value.

### HasType

`func (o *TradeOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *TradeOrder) GetTimeInForce() TradeOrderTimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *TradeOrder) GetTimeInForceOk() (*TradeOrderTimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *TradeOrder) SetTimeInForce(v TradeOrderTimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *TradeOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetStatus

`func (o *TradeOrder) GetStatus() TradeOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradeOrder) GetStatusOk() (*TradeOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradeOrder) SetStatus(v TradeOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TradeOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *TradeOrder) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TradeOrder) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TradeOrder) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TradeOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAmount

`func (o *TradeOrder) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TradeOrder) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TradeOrder) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TradeOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOpenAmount

`func (o *TradeOrder) GetOpenAmount() float64`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *TradeOrder) GetOpenAmountOk() (*float64, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *TradeOrder) SetOpenAmount(v float64)`

SetOpenAmount sets OpenAmount field to given value.

### HasOpenAmount

`func (o *TradeOrder) HasOpenAmount() bool`

HasOpenAmount returns a boolean if a field has been set.

### GetCreated

`func (o *TradeOrder) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TradeOrder) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TradeOrder) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TradeOrder) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


