package goklaviyo

import (
	"log"
	"testing"
)

// func TestGetList(t *testing.T) {
//     apiKey := "pk_fd63d1b7a65f150c2d70994539db5077bb"
// 	listID := "QGnDEg"
// 	client := New(apiKey)

// 	list, err := client.GetList(listID, nil)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}	

// 	log.Println(list)
// }

func TestGetLists(t *testing.T) {
    apiKey := "pk_fd63d1b7a65f150c2d70994539db5077bb"
	client := New(apiKey)

	list, err := client.GetLists(nil)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(list)
}

// func TestGetListMembers(t *testing.T) {
// 	apiKey := "pk_fd63d1b7a65f150c2d70994539db5077bb"
// 	listID := "QGnDEg"
	
// 	client := New(apiKey)

// 	params := InterestCategoriesQueryParams{}
// 	params.ListID = listID
// 	members, err := client.GetMembers(&params)
// 	if err != nil {
// 		log.Println("ERROR: ", err.Error())
// 	} else {
// 		log.Println(members)
// 	}
// }



func TestCheckMember(t *testing.T) {
	apiKey := "pk_fd63d1b7a65f150c2d70994539db5077bb"
	listID := "QGnDEg"
	
	client := New(apiKey)

	lists := []string{}
	lists = append(lists, "clonenora04@gmail.com")

	contactCheck := ContactCheck{}
	contactCheck.Emails = lists

	members, err := client.CheckMember(listID, &contactCheck)
	if err != nil {
		log.Println("ERROR: ", err.Error())
	} else {
		log.Println(len(*members))
	}
}

func TestCreateMember(t *testing.T) {
	apiKey := "pk_fd63d1b7a65f150c2d70994539db5077bb"
	listID := "QGnDEg"
	
	client := New(apiKey)

	member := MemberRequest{}
	member.Email = "clonenora05@gmail.com"
	member.FirstName = "No"
	member.LastName = "Ga"
	member.Phone = "+84367998641"

	lists := []MemberRequest{}
	lists = append(lists, member)

	memberSubscriber := MemberSubscribeRequest{}
	memberSubscriber.Profiles = lists

	members, err := client.CreateMember(listID, &memberSubscriber)
	if err != nil {
		log.Println("ERROR: ", err.Error())
	} else {
		log.Println(members)
	}
}