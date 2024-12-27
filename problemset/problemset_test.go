package problemset

import (
	"testing"

	"github.com/cloudfstrife/go-codeforces/client"
)

func TestProblems(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result *ProblemsResult

	var par = ProblemsPar{
		Tags:           nil,
		ProblemsetName: "",
	}

	result, err = Problems(cli, par)
	if err != nil {
		t.Error("get problems failed", err)
		return
	}
	t.Logf("get problems success %#v", result)
}

func TestRecentStatus(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result = []Submission{}

	var par = RecentStatusPar{
		Count:          10,
		ProblemsetName: "",
	}

	result, err = RecentStatus(cli, par)
	if err != nil {
		t.Error("get recent status failed", err)
		return
	}
	t.Logf("get recent status success %#v", result)
}
