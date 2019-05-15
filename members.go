package goklaviyo

import (
	"fmt"
)

const (
	members_path       = "/list/%s/members"
	single_member_path = members_path + "/%s"

	member_activity_path = single_member_path + "/activity"
)

type ListOfMembers struct {
	Contacts []Contact `json:"contacts"`
	Paging   Paging    `json:"paging"`
}

type MemberRequest struct {
	Email string `json:"email,omitempty"`
	AdditionalMemBerRequest
}

type AdditionalMemBerRequest struct {
	CreatedAt   string `json:"$createdAt,omitempty"`
	FirstName   string `json:"$firstName,omitempty"`
	LastName    string `json:"$lastName,omitempty"`
	Country     string `json:"$country,omitempty"`
	CountryCode string `json:"$countryCode,omitempty"`
	State       string `json:"$state,omitempty"`
	City        string `json:"$city,omitempty"`
	Address     string `json:"$address,omitempty"`
	PostalCode  string `json:"$postalCode,omitempty"`
	Gender      string `json:"$gender,omitempty"`
	Phone       string `json:"$phone,omitempty"`
	Birthdate   string `json:"$birthdate,omitempty"`
}

type MemberSubscribeRequest struct {
	Profiles []MemberRequest `json:"profiles"`
}

type Contact struct {
	MemberRequest
	ID    string `json:"id"`
	Email string `json:"email"`

	api *API
}

type ContactCheck struct {
	Emails []string `json:"emails"`
}

type Statuses struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

func (api API) GetMembers(params *InterestCategoriesQueryParams) (*ListOfMembers, error) {
	endpoint := members_path
	response := new(ListOfMembers)

	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, m := range response.Contacts {
		m.api = &api
	}
	return response, nil
}

func (api API) GetMember(listID string, id string, params *BasicQueryParams) (*Contact, error) {

	endpoint := fmt.Sprintf(single_member_path, listID, id)
	response := new(Contact)
	response.api = &api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api API) CheckMember(listID string, body *ContactCheck) (*[]Contact, error) {
	endpoint := fmt.Sprintf(members_path, listID)
	response := new([]Contact)
	// response.api = &api
	for _, res := range *response {
		res.api = &api
	}

	return response, api.Request("GET", endpoint, nil, body, response)
}

func (api API) CreateMember(listID string, body *MemberSubscribeRequest) (*[]Contact, error) {
	endpoint := fmt.Sprintf(members_path, listID)
	response := new([]Contact)
	// response.api = &api
	for _, res := range *response {
		res.api = &api
	}

	return response, api.Request("POST", endpoint, nil, body, response)
}

// func (api API) UpdateMember(id string, body *MemberRequest) (*Member, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return nil, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	response := new(Member)
// 	response.api = list.api

// 	return response, list.api.Request("PATCH", endpoint, nil, body, response)
// }

// func (api API) AddOrUpdateMember(id string, body *MemberRequest) (*Member, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return nil, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	response := new(Member)
// 	response.api = list.api

// 	return response, list.api.Request("PUT", endpoint, nil, body, response)
// }

// func (api API) DeleteMember(id string) (bool, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return false, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	return list.api.RequestOk("DELETE", endpoint)
// }
