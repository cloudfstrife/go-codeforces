package user

import (
	"testing"

	"github.com/cloudfstrife/go-codeforces/client"
)

func TestBlogEntries(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := BlogEntries(cli, "Fefer_Ivan")
	if err != nil {
		t.Error("get user blog entries failed", err)
		return
	}
	t.Logf("get user blog entries success %#v", result)
}

// func TestFriends(t *testing.T) {
// 	cli := client.NewClient(&client.Config{
// 		Lang: "en",
// 	})

// 	result, err := Friends(cli, false)
// 	if err != nil {
// 		t.Error("get user friends failed", err)
// 		return
// 	}
// 	t.Logf("get user friends success %#v", result)
// }

func TestInfo(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := Info(cli, []string{"", "Fefer_Ivan"}, false)
	if err != nil {
		t.Error("get user info failed", err)
		return
	}
	t.Logf("get user info success %#v", result)
}

func TestRatedList(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := RatedList(cli, RatedListPar{
		ActiveOnly:     true,
		IncludeRetired: false,
		ContestID:      2044,
	})
	if err != nil {
		t.Error("get user rated list failed", err)
		return
	}
	t.Logf("get user rated list success %#v", result)
}

func TestRating(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := Rating(cli, "Fefer_Ivan")
	if err != nil {
		t.Error("get user rating failed", err)
		return
	}
	t.Logf("get user rating success %#v", result)
}

func TestStatus(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := Status(cli, StatusPar{
		Handle: "Fefer_Ivan",
		From:   1,
		Count:  10,
	})
	if err != nil {
		t.Error("get user status failed", err)
		return
	}
	t.Logf("get user status success %#v", result)
}
