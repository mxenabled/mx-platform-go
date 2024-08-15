# MonthlyCashFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Unique identifier for the monthly cash flow profile. Defined by MX. | [optional] 
**UserGuid** | Pointer to **string** | Unique identifier for the user the monthly cash flow profile is attached to. Defined by MX. | [optional] 
**BudgetedIncome** | Pointer to **float32** | The amount of the budgeted income for the user. | [optional] 
**BudgetedExpenses** | Pointer to **float32** | The amount of the budgeted expenses for the user. | [optional] 
**GoalsContribution** | Pointer to **float32** | The monthly dollar amount allocated for goals. | [optional] 
**EstimatedGoalsContribution** | Pointer to **int32** | The estimated monthly dollar amount allocated for goals calculated from income and budgets. | [optional] 
**UsesEstimatedGoalsContribution** | Pointer to **bool** |  | [optional] 

## Methods

### NewMonthlyCashFlowResponse

`func NewMonthlyCashFlowResponse() *MonthlyCashFlowResponse`

NewMonthlyCashFlowResponse instantiates a new MonthlyCashFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyCashFlowResponseWithDefaults

`func NewMonthlyCashFlowResponseWithDefaults() *MonthlyCashFlowResponse`

NewMonthlyCashFlowResponseWithDefaults instantiates a new MonthlyCashFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *MonthlyCashFlowResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *MonthlyCashFlowResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *MonthlyCashFlowResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *MonthlyCashFlowResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetUserGuid

`func (o *MonthlyCashFlowResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *MonthlyCashFlowResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *MonthlyCashFlowResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *MonthlyCashFlowResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### GetBudgetedIncome

`func (o *MonthlyCashFlowResponse) GetBudgetedIncome() float32`

GetBudgetedIncome returns the BudgetedIncome field if non-nil, zero value otherwise.

### GetBudgetedIncomeOk

`func (o *MonthlyCashFlowResponse) GetBudgetedIncomeOk() (*float32, bool)`

GetBudgetedIncomeOk returns a tuple with the BudgetedIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetedIncome

`func (o *MonthlyCashFlowResponse) SetBudgetedIncome(v float32)`

SetBudgetedIncome sets BudgetedIncome field to given value.

### HasBudgetedIncome

`func (o *MonthlyCashFlowResponse) HasBudgetedIncome() bool`

HasBudgetedIncome returns a boolean if a field has been set.

### GetBudgetedExpenses

`func (o *MonthlyCashFlowResponse) GetBudgetedExpenses() float32`

GetBudgetedExpenses returns the BudgetedExpenses field if non-nil, zero value otherwise.

### GetBudgetedExpensesOk

`func (o *MonthlyCashFlowResponse) GetBudgetedExpensesOk() (*float32, bool)`

GetBudgetedExpensesOk returns a tuple with the BudgetedExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetedExpenses

`func (o *MonthlyCashFlowResponse) SetBudgetedExpenses(v float32)`

SetBudgetedExpenses sets BudgetedExpenses field to given value.

### HasBudgetedExpenses

`func (o *MonthlyCashFlowResponse) HasBudgetedExpenses() bool`

HasBudgetedExpenses returns a boolean if a field has been set.

### GetGoalsContribution

`func (o *MonthlyCashFlowResponse) GetGoalsContribution() float32`

GetGoalsContribution returns the GoalsContribution field if non-nil, zero value otherwise.

### GetGoalsContributionOk

`func (o *MonthlyCashFlowResponse) GetGoalsContributionOk() (*float32, bool)`

GetGoalsContributionOk returns a tuple with the GoalsContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalsContribution

`func (o *MonthlyCashFlowResponse) SetGoalsContribution(v float32)`

SetGoalsContribution sets GoalsContribution field to given value.

### HasGoalsContribution

`func (o *MonthlyCashFlowResponse) HasGoalsContribution() bool`

HasGoalsContribution returns a boolean if a field has been set.

### GetEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) GetEstimatedGoalsContribution() int32`

GetEstimatedGoalsContribution returns the EstimatedGoalsContribution field if non-nil, zero value otherwise.

### GetEstimatedGoalsContributionOk

`func (o *MonthlyCashFlowResponse) GetEstimatedGoalsContributionOk() (*int32, bool)`

GetEstimatedGoalsContributionOk returns a tuple with the EstimatedGoalsContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) SetEstimatedGoalsContribution(v int32)`

SetEstimatedGoalsContribution sets EstimatedGoalsContribution field to given value.

### HasEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) HasEstimatedGoalsContribution() bool`

HasEstimatedGoalsContribution returns a boolean if a field has been set.

### GetUsesEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) GetUsesEstimatedGoalsContribution() bool`

GetUsesEstimatedGoalsContribution returns the UsesEstimatedGoalsContribution field if non-nil, zero value otherwise.

### GetUsesEstimatedGoalsContributionOk

`func (o *MonthlyCashFlowResponse) GetUsesEstimatedGoalsContributionOk() (*bool, bool)`

GetUsesEstimatedGoalsContributionOk returns a tuple with the UsesEstimatedGoalsContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) SetUsesEstimatedGoalsContribution(v bool)`

SetUsesEstimatedGoalsContribution sets UsesEstimatedGoalsContribution field to given value.

### HasUsesEstimatedGoalsContribution

`func (o *MonthlyCashFlowResponse) HasUsesEstimatedGoalsContribution() bool`

HasUsesEstimatedGoalsContribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


