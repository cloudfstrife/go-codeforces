package recentactions

import (
	"testing"

	"github.com/cloudfstrife/go-codeforces/client"
)

func TestRecentActions(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := RecentActions(cli, 30)
	if err != nil {
		t.Error("get recent actions failed", err)
		return
	}
	t.Logf("get recent actions success %#v", result)
}
