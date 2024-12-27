package user

import (
	"errors"
	"strconv"

	"github.com/cloudfstrife/go-codeforces/blog"
	"github.com/cloudfstrife/go-codeforces/client"
	"github.com/cloudfstrife/go-codeforces/problemset"
)

func BlogEntries(cli *client.Client, handle string) ([]blog.BlogEntry, error) {
	var result = make([]blog.BlogEntry, 0)
	var err = cli.Call("/user.blogEntries", map[string][]string{
		"handle": {handle},
	}, &result)
	return result, err
}

func Friends(cli *client.Client, onlyOnline bool) ([]string, error) {
	var result = make([]string, 0)
	var err = cli.Call("/user.friends", map[string][]string{
		"onlyOnline": {strconv.FormatBool(onlyOnline)},
	}, result)
	return result, err
}

type User struct {
	Handle                  string `json:"handle"`                  // . Codeforces user handle.
	Email                   string `json:"email"`                   // . Shown only if user allowed to share his contact info.
	VkId                    string `json:"vkId"`                    // . User id for VK social network. Shown only if user allowed to share his contact info.
	OpenId                  string `json:"openId"`                  // . Shown only if user allowed to share his contact info.
	FirstName               string `json:"firstName"`               // . Localized. Can be absent.
	LastName                string `json:"lastName"`                // . Localized. Can be absent.
	Country                 string `json:"country"`                 // . Localized. Can be absent.
	City                    string `json:"city"`                    // . Localized. Can be absent.
	Organization            string `json:"organization"`            // . Localized. Can be absent.
	Contribution            int64  `json:"contribution"`            // . User contribution.
	Rank                    string `json:"rank"`                    // . Localized.
	Rating                  int64  `json:"rating"`                  // .
	MaxRank                 string `json:"maxRank"`                 // . Localized.
	MaxRating               int64  `json:"maxRating"`               // .
	LastOnlineTimeSeconds   int64  `json:"lastOnlineTimeSeconds"`   // . Time, when user was last seen online, in unix format.
	RegistrationTimeSeconds int64  `json:"registrationTimeSeconds"` // . Time, when user was registered, in unix format.
	FriendOfCount           int64  `json:"friendOfCount"`           // . Amount of users who have this user in friends.
	Avatar                  string `json:"avatar"`                  // . User's avatar URL.
	TitlePhoto              string `json:"titlePhoto"`              // . User's title photo URL.
}

func Info(cli *client.Client, handles []string, checkHistoricHandles bool) ([]User, error) {
	var result = make([]User, 0)
	var err = cli.Call("/user.info", map[string][]string{
		"handles":              handles,
		"checkHistoricHandles": {strconv.FormatBool(checkHistoricHandles)},
	}, &result)
	return result, err
}

type RatedListPar struct {
	ActiveOnly     bool  `json:"activeOnly"`     //. If true then only users, who participated in rated contest during the last month are returned. Otherwise, all users with at least one rated contest are returned.
	IncludeRetired bool  `json:"includeRetired"` //. If true, the method returns all rated users, otherwise the method returns only users, that were online at last month.
	ContestID      int64 `json:"contestId"`      //. Id of the contest. It is not the round number. It can be seen in contest URL. For example: /contest/566/status
}

func (sp RatedListPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	result["activeOnly"] = []string{strconv.FormatBool(sp.ActiveOnly)}
	result["includeRetired"] = []string{strconv.FormatBool(sp.IncludeRetired)}
	if sp.ContestID != 0 {
		result["contestId"] = []string{strconv.FormatInt(sp.ContestID, 10)}
	}
	return result, nil
}

func RatedList(cli *client.Client, rlPar RatedListPar) ([]User, error) {
	var result = make([]User, 0)
	par, err := rlPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/user.ratedList", par, &result)
	return result, err
}

type RatingChange struct {
	ContestId               int64  `json:"contestId"`               //.
	ContestName             string `json:"contestName"`             //. Localized.
	Handle                  string `json:"handle"`                  //. Codeforces user handle.
	Rank                    int64  `json:"rank"`                    //. Place of the user in the contest. This field contains user rank on the moment of rating update. If afterwards rank changes (e.g. someone get disqualified), this field will not be update and will contain old rank.
	RatingUpdateTimeSeconds int64  `json:"ratingUpdateTimeSeconds"` //. Time, when rating for the contest was update, in unix-format.
	OldRating               int64  `json:"oldRating"`               //. User rating before the contest.
	NewRating               int64  `json:"newRating"`               //. User rating after the contest.
}

func Rating(cli *client.Client, handle string) ([]RatingChange, error) {
	var result = make([]RatingChange, 0)
	var err = cli.Call("/user.rating", map[string][]string{
		"handle": {handle},
	}, &result)
	return result, err
}

type StatusPar struct {
	Handle string `json:"handle"` // (Required) 	Codeforces user handle.
	From   int64  `json:"from"`   // 	1-based index of the first submission to return.
	Count  int64  `json:"count"`  // 	Number of returned submissions.
}

func (sp StatusPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if sp.Handle == "" {
		return nil, errors.New("handle is required")
	}
	if sp.From == 0 {
		sp.From = 1
	}
	if sp.Count == 0 {
		sp.Count = 10
	}

	result["handle"] = []string{sp.Handle}
	result["from"] = []string{strconv.FormatInt(sp.From, 10)}
	result["count"] = []string{strconv.FormatInt(sp.Count, 10)}
	return result, nil
}

func Status(cli *client.Client, spar StatusPar) ([]problemset.Submission, error) {
	var result = make([]problemset.Submission, 0)
	par, err := spar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/user.status", par, &result)
	return result, err
}
