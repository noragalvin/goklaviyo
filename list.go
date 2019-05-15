package goklaviyo

import (
	"errors"
	"fmt"
)

const (
	lists_path         = "/list"
	multiple_list_path = "/lists"
	single_list_path   = lists_path + "/%s"

	interest_categories_path      = "/lists/%s/interest-categories"
	single_interest_category_path = interest_categories_path + "/%s"
)

type ListQueryParams struct {
	ExtendedQueryParams

	BeforeDateCreated      string
	SinceDateCreated       string
	BeforeCampaignLastSent string
	SinceCampaignLastSent  string
	Email                  string
}

func (q ListQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["before_date_created"] = q.BeforeDateCreated
	m["since_date_created"] = q.SinceDateCreated
	m["before_campaign_last_sent"] = q.BeforeCampaignLastSent
	m["since_campaign_last_sent"] = q.SinceCampaignLastSent
	m["email"] = q.Email
	return m
}

type ListOfLists struct {
	Lists []ListResponse `json:"lists"`
}

type ListCreationRequest struct {
	Name                string           `json:"name"`
	Contact             Contact          `json:"contact"`
	PermissionReminder  string           `json:"permission_reminder"`
	UseArchiveBar       bool             `json:"use_archive_bar"`
	CampaignDefaults    CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe   string           `json:"notify_on_subscribe"`
	NotifyOnUnsubscribe string           `json:"notify_on_unsubscribe"`
	EmailTypeOption     bool             `json:"email_type_option"`
	Visibility          string           `json:"visibility"`
}

type ListResponse struct {
	ListID      string `json:"list_id"`
	ListName    string `json:"list_name"`
	ListType    string `json:"list_type"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	PersonCount int    `json:"person_count"`

	api *API
}

func (list ListResponse) CanMakeRequest() error {
	if list.ListID == "" {
		return errors.New("No ID provided on list")
	}

	return nil
}

type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

func (api API) GetLists(params *ListQueryParams) (*[]ListResponse, error) {
	response := new([]ListResponse)

	err := api.Request("GET", multiple_list_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, res := range *response {
		res.api = &api
	}

	return response, nil
}

// NewListResponse returns a *ListResponse that is minimally viable for making
// API requests. This is useful for such API requests that depend on a
// ListResponse for its ID (e.g. CreateMember) without having to make a second
// network request to get the list itself.
func (api API) NewListResponse(id string) *ListResponse {
	return &ListResponse{
		ListID: id,
		api:    &api,
	}
}

func (api API) GetList(id string, params *BasicQueryParams) (*ListResponse, error) {
	endpoint := fmt.Sprintf(single_list_path, id)
	response := new(ListResponse)
	response.api = &api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api API) CreateList(body *ListCreationRequest) (*ListResponse, error) {
	response := new(ListResponse)
	response.api = &api
	return response, api.Request("POST", lists_path, nil, body, response)
}

func (api API) UpdateList(id string, body *ListCreationRequest) (*ListResponse, error) {
	endpoint := fmt.Sprintf(single_list_path, id)

	response := new(ListResponse)
	response.api = &api

	return response, api.Request("PATCH", endpoint, nil, body, response)
}

func (api API) DeleteList(id string) (bool, error) {
	endpoint := fmt.Sprintf(single_list_path, id)
	return api.RequestOk("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Interest Categories
// ------------------------------------------------------------------------------------------------

type ListOfInterestCategories struct {
	ListID     string             `json:"listID"`
	Categories []InterestCategory `json:"categories"`
}

type InterestCategoryRequest struct {
	Title        string `json:"title"`
	DisplayOrder int    `json:"display_order"`
	Type         string `json:"type"`
}

type InterestCategory struct {
	InterestCategoryRequest

	ListID string `json:"listID"`
	ID     string `json:"id"`

	withLinks
	api *API
}

func (interestCatgory InterestCategory) CanMakeRequest() error {
	if interestCatgory.ID == "" {
		return errors.New("No ID provided on interest category")
	}

	return nil
}

type InterestCategoriesQueryParams struct {
	ExtendedQueryParams

	ListID string `json:"listID"`
	Type   string `json:"type"`
}

func (q InterestCategoriesQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["type"] = q.Type
	return m
}

func (list ListResponse) GetInterestCategories(params *InterestCategoriesQueryParams) (*ListOfInterestCategories, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(interest_categories_path, list.ListID)
	response := new(ListOfInterestCategories)

	err := list.api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	for i, _ := range response.Categories {
		response.Categories[i].api = list.api
	}

	return response, nil
}

func (list ListResponse) GetInterestCategory(id string, params *BasicQueryParams) (*InterestCategory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ListID, id)
	response := new(InterestCategory)
	response.api = list.api

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) CreateInterestCategory(body *InterestCategoryRequest) (*InterestCategory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(interest_categories_path, list.ListID)
	response := new(InterestCategory)
	response.api = list.api

	return response, list.api.Request("POST", endpoint, nil, body, response)
}

func (list ListResponse) UpdateInterestCategory(id string, body *InterestCategoryRequest) (*InterestCategory, error) {
	if list.ListID == "" {
		return nil, errors.New("No ID provided on list")
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ListID, id)
	response := new(InterestCategory)
	response.api = list.api

	return response, list.api.Request("PATCH", endpoint, nil, body, response)
}

func (list ListResponse) DeleteInterestCategory(id string) (bool, error) {
	if list.ListID == "" {
		return false, errors.New("No ID provided on list")
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ListID, id)
	return list.api.RequestOk("DELETE", endpoint)
}
