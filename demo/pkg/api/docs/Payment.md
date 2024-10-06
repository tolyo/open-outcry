# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**PaymentType**](PaymentType.md) |  | 
**Amount** | **float64** |  | 
**Currency** | **string** | ISO 4217 Currency symbol | 
**SenderAccountId** | **string** |  | 
**BeneficiaryAccountId** | **string** |  | 
**Details** | **string** |  | 
**ExternalReferenceNumber** | **string** |  | 
**Status** | **string** |  | 
**DebitBalanceAmount** | Pointer to **float64** |  | [optional] 
**CreditBalanceAmount** | Pointer to **float64** |  | [optional] 

## Methods

### NewPayment

`func NewPayment(id string, type_ PaymentType, amount float64, currency string, senderAccountId string, beneficiaryAccountId string, details string, externalReferenceNumber string, status string, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Payment) GetType() PaymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payment) GetTypeOk() (*PaymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payment) SetType(v PaymentType)`

SetType sets Type field to given value.


### GetAmount

`func (o *Payment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Payment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSenderAccountId

`func (o *Payment) GetSenderAccountId() string`

GetSenderAccountId returns the SenderAccountId field if non-nil, zero value otherwise.

### GetSenderAccountIdOk

`func (o *Payment) GetSenderAccountIdOk() (*string, bool)`

GetSenderAccountIdOk returns a tuple with the SenderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountId

`func (o *Payment) SetSenderAccountId(v string)`

SetSenderAccountId sets SenderAccountId field to given value.


### GetBeneficiaryAccountId

`func (o *Payment) GetBeneficiaryAccountId() string`

GetBeneficiaryAccountId returns the BeneficiaryAccountId field if non-nil, zero value otherwise.

### GetBeneficiaryAccountIdOk

`func (o *Payment) GetBeneficiaryAccountIdOk() (*string, bool)`

GetBeneficiaryAccountIdOk returns a tuple with the BeneficiaryAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountId

`func (o *Payment) SetBeneficiaryAccountId(v string)`

SetBeneficiaryAccountId sets BeneficiaryAccountId field to given value.


### GetDetails

`func (o *Payment) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Payment) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Payment) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetExternalReferenceNumber

`func (o *Payment) GetExternalReferenceNumber() string`

GetExternalReferenceNumber returns the ExternalReferenceNumber field if non-nil, zero value otherwise.

### GetExternalReferenceNumberOk

`func (o *Payment) GetExternalReferenceNumberOk() (*string, bool)`

GetExternalReferenceNumberOk returns a tuple with the ExternalReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceNumber

`func (o *Payment) SetExternalReferenceNumber(v string)`

SetExternalReferenceNumber sets ExternalReferenceNumber field to given value.


### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDebitBalanceAmount

`func (o *Payment) GetDebitBalanceAmount() float64`

GetDebitBalanceAmount returns the DebitBalanceAmount field if non-nil, zero value otherwise.

### GetDebitBalanceAmountOk

`func (o *Payment) GetDebitBalanceAmountOk() (*float64, bool)`

GetDebitBalanceAmountOk returns a tuple with the DebitBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalanceAmount

`func (o *Payment) SetDebitBalanceAmount(v float64)`

SetDebitBalanceAmount sets DebitBalanceAmount field to given value.

### HasDebitBalanceAmount

`func (o *Payment) HasDebitBalanceAmount() bool`

HasDebitBalanceAmount returns a boolean if a field has been set.

### GetCreditBalanceAmount

`func (o *Payment) GetCreditBalanceAmount() float64`

GetCreditBalanceAmount returns the CreditBalanceAmount field if non-nil, zero value otherwise.

### GetCreditBalanceAmountOk

`func (o *Payment) GetCreditBalanceAmountOk() (*float64, bool)`

GetCreditBalanceAmountOk returns a tuple with the CreditBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceAmount

`func (o *Payment) SetCreditBalanceAmount(v float64)`

SetCreditBalanceAmount sets CreditBalanceAmount field to given value.

### HasCreditBalanceAmount

`func (o *Payment) HasCreditBalanceAmount() bool`

HasCreditBalanceAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


