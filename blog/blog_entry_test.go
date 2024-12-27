package blog

import (
	"testing"

	"github.com/cloudfstrife/go-codeforces/client"
)

func TestView(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := View(cli, 136855)
	if err != nil {
		t.Error("get blog entry view failed", err)
		return
	}
	t.Logf("get blog entry view success %#v", result)
}

func TestComments(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	result, err := Comments(cli, 136855)
	if err != nil {
		t.Error("get blog entry view failed", err)
		return
	}
	t.Logf("get blog entry view success %#v", len(result))
}
