# EmaopayUploadCredits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | 存储桶名 | [optional] 
**Credentials** | Pointer to [**EmaopayCredentials**](EmaopayCredentials.md) | 临时证书 | [optional] 
**Expiration** | Pointer to **string** | 证书有效的时间，以 iso8601 格式的 UTC 时间表示 注意：此字段可能返回 null，表示取不到有效值。 | [optional] 
**ExpiredTime** | Pointer to **int32** | 临时证书有效的时间，返回 Unix 时间戳，精确到秒 | [optional] 
**Region** | Pointer to **string** | 区域 | [optional] 

## Methods

### NewEmaopayUploadCredits

`func NewEmaopayUploadCredits() *EmaopayUploadCredits`

NewEmaopayUploadCredits instantiates a new EmaopayUploadCredits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmaopayUploadCreditsWithDefaults

`func NewEmaopayUploadCreditsWithDefaults() *EmaopayUploadCredits`

NewEmaopayUploadCreditsWithDefaults instantiates a new EmaopayUploadCredits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *EmaopayUploadCredits) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *EmaopayUploadCredits) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *EmaopayUploadCredits) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *EmaopayUploadCredits) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCredentials

`func (o *EmaopayUploadCredits) GetCredentials() EmaopayCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EmaopayUploadCredits) GetCredentialsOk() (*EmaopayCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EmaopayUploadCredits) SetCredentials(v EmaopayCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *EmaopayUploadCredits) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetExpiration

`func (o *EmaopayUploadCredits) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EmaopayUploadCredits) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EmaopayUploadCredits) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *EmaopayUploadCredits) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpiredTime

`func (o *EmaopayUploadCredits) GetExpiredTime() int32`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *EmaopayUploadCredits) GetExpiredTimeOk() (*int32, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *EmaopayUploadCredits) SetExpiredTime(v int32)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *EmaopayUploadCredits) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetRegion

`func (o *EmaopayUploadCredits) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EmaopayUploadCredits) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EmaopayUploadCredits) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EmaopayUploadCredits) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


