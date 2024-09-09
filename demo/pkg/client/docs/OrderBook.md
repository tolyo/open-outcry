# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sell** | Pointer to [**[]PriceVolume**](PriceVolume.md) |  | [optional] 
**Buy** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewOrderBook

`func NewOrderBook() *OrderBook`

NewOrderBook instantiates a new OrderBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookWithDefaults

`func NewOrderBookWithDefaults() *OrderBook`

NewOrderBookWithDefaults instantiates a new OrderBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSell

`func (o *OrderBook) GetSell() []PriceVolume`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *OrderBook) GetSellOk() (*[]PriceVolume, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *OrderBook) SetSell(v []PriceVolume)`

SetSell sets Sell field to given value.

### HasSell

`func (o *OrderBook) HasSell() bool`

HasSell returns a boolean if a field has been set.

### GetBuy

`func (o *OrderBook) GetBuy() interface{}`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *OrderBook) GetBuyOk() (*interface{}, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *OrderBook) SetBuy(v interface{})`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *OrderBook) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### SetBuyNil

`func (o *OrderBook) SetBuyNil(b bool)`

 SetBuyNil sets the value for Buy to be an explicit nil

### UnsetBuy
`func (o *OrderBook) UnsetBuy()`

UnsetBuy ensures that no value is present for Buy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


