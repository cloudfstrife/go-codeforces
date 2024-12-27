package contest

import (
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
)

func RatingChanges(cli *client.Client, id int64) ([]RatingChange, error) {
	var result = make([]RatingChange, 0)
	var err = cli.Call("/contest.ratingChanges", map[string][]string{
		"contestId": {strconv.FormatInt(id, 10)},
	}, &result)

	return result, err
}

type RatingChange struct {
	ContestID               int64  `json:"contestId"`               // .
	ContestName             string `json:"contestName"`             // . Localized.
	Handle                  string `json:"handle"`                  // . Codeforces user handle.
	Rank                    int64  `json:"rank"`                    // . Place of the user in the contest. This field contains user rank on the moment of rating update. If afterwards rank changes (e.g. someone get disqualified), this field will not be update and will contain old rank.
	RatingUpdateTimeSeconds int64  `json:"ratingUpdateTimeSeconds"` // . Time, when rating for the contest was update, in unix-format.
	OldRating               int64  `json:"oldRating"`               // . User rating before the contest.
	NewRating               int64  `json:"newRating"`               // . User rating after the contest.
}
