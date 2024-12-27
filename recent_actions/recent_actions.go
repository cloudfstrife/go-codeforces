package recentactions

import (
	"strconv"

	"github.com/cloudfstrife/go-codeforces/blog"
	"github.com/cloudfstrife/go-codeforces/client"
)

func RecentActions(cli *client.Client, maxCount int64) ([]RecentAction, error) {
	var result = make([]RecentAction, 0)
	var err = cli.Call("/recentActions", map[string][]string{
		"maxCount": {strconv.FormatInt(maxCount, 10)},
	}, &result)

	return result, err
}

type RecentAction struct {
	TimeSeconds int64          `json:"timeSeconds"` // . Action time, in unix format.
	BlogEntry   blog.BlogEntry `json:"blogEntry"`   //  object in short form. Can be absent.
	Comment     blog.Comment   `json:"comment"`     //  object. Can be absent.
}
