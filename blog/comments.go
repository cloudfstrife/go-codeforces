package blog

import (
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
)

func Comments(cli *client.Client, id int64) ([]Comment, error) {
	var result = make([]Comment, 0)
	var err = cli.Call("/blogEntry.comments", map[string][]string{
		"blogEntryId": {strconv.FormatInt(id, 10)},
	}, &result)

	return result, err
}

type Comment struct {
	ID                  int64  `json:"id"`                  // . ID
	CreationTimeSeconds int64  `json:"creationTimeSeconds"` // . Time, when comment was created, in unix format.
	CommentatorHandle   string `json:"commentatorHandle"`   // . CommentatorHandle
	Locale              string `json:"locale"`              // . Locale
	Text                string `json:"text"`                // . Text
	ParentCommentID     int64  `json:"parentCommentId"`     // . Can be absent.
	Rating              int64  `json:"rating"`              // . Rating
}
