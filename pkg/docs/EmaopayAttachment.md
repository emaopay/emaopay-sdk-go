# EmaopayAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | 可用的 | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PostId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewEmaopayAttachment

`func NewEmaopayAttachment() *EmaopayAttachment`

NewEmaopayAttachment instantiates a new EmaopayAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmaopayAttachmentWithDefaults

`func NewEmaopayAttachmentWithDefaults() *EmaopayAttachment`

NewEmaopayAttachmentWithDefaults instantiates a new EmaopayAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *EmaopayAttachment) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EmaopayAttachment) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EmaopayAttachment) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *EmaopayAttachment) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmaopayAttachment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmaopayAttachment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmaopayAttachment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmaopayAttachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EmaopayAttachment) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EmaopayAttachment) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EmaopayAttachment) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EmaopayAttachment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *EmaopayAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmaopayAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmaopayAttachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmaopayAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostId

`func (o *EmaopayAttachment) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *EmaopayAttachment) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *EmaopayAttachment) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *EmaopayAttachment) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetProvider

`func (o *EmaopayAttachment) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmaopayAttachment) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmaopayAttachment) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmaopayAttachment) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecret

`func (o *EmaopayAttachment) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EmaopayAttachment) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EmaopayAttachment) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *EmaopayAttachment) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EmaopayAttachment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmaopayAttachment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmaopayAttachment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmaopayAttachment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *EmaopayAttachment) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EmaopayAttachment) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EmaopayAttachment) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EmaopayAttachment) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


