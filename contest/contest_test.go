package contest

import (
	"testing"

	"github.com/cloudfstrife/go-codeforces/client"
)

func TestList(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result []Contest

	result, err = List(cli, true)
	if err != nil {
		t.Error("get contest list failed", err)
	}
	t.Log("get contest list success", len(result))

	result, err = List(cli, false)
	if err != nil {
		t.Error("get contest list failed", err)
		return
	}
	t.Log("get contest list success", result[0])
}

func TestStatus(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result []StatusResult

	var par StatusPar = StatusPar{
		ContestID: 2043,
	}

	result, err = Status(cli, par)
	if err != nil {
		t.Error("get contest status failed", err)
		return
	}
	t.Logf("get contest status success %#v", result[0])
}

func TestHacks(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result []HacksResult

	var par HacksPar = HacksPar{
		ContestID: 566,
	}

	result, err = Hacks(cli, par)
	if err != nil {
		t.Error("get contest hacks failed", err)
		return
	}
	t.Logf("get contest hacks success %#v", result[0])
}

func TestStandings(t *testing.T) {
	cli := client.NewClient(&client.Config{
		Lang: "en",
	})

	var err error
	var result *StandingsResult

	var par = StandingsPar{
		ContestID:      566,
		AsManager:      false,
		From:           1,
		Count:          5,
		ShowUnofficial: true,
	}

	result, err = Standings(cli, par)
	if err != nil {
		t.Error("get contest standings failed", err)
		return
	}
	t.Logf("get contest standings success %#v", result)
}
