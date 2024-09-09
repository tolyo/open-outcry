# PaymentAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | ISO 4217 Currency symbol | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**AmountReserved** | Pointer to **float64** |  | [optional] 
**AmountAvailable** | Pointer to **float64** |  | [optional] 

## Methods

### NewPaymentAccount

`func NewPaymentAccount() *PaymentAccount`

NewPaymentAccount instantiates a new PaymentAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountWithDefaults

`func NewPaymentAccountWithDefaults() *PaymentAccount`

NewPaymentAccountWithDefaults instantiates a new PaymentAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentAccount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAccount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAccount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountReserved

`func (o *PaymentAccount) GetAmountReserved() float64`

GetAmountReserved returns the AmountReserved field if non-nil, zero value otherwise.

### GetAmountReservedOk

`func (o *PaymentAccount) GetAmountReservedOk() (*float64, bool)`

GetAmountReservedOk returns a tuple with the AmountReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountReserved

`func (o *PaymentAccount) SetAmountReserved(v float64)`

SetAmountReserved sets AmountReserved field to given value.

### HasAmountReserved

`func (o *PaymentAccount) HasAmountReserved() bool`

HasAmountReserved returns a boolean if a field has been set.

### GetAmountAvailable

`func (o *PaymentAccount) GetAmountAvailable() float64`

GetAmountAvailable returns the AmountAvailable field if non-nil, zero value otherwise.

### GetAmountAvailableOk

`func (o *PaymentAccount) GetAmountAvailableOk() (*float64, bool)`

GetAmountAvailableOk returns a tuple with the AmountAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAvailable

`func (o *PaymentAccount) SetAmountAvailable(v float64)`

SetAmountAvailable sets AmountAvailable field to given value.

### HasAmountAvailable

`func (o *PaymentAccount) HasAmountAvailable() bool`

HasAmountAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


