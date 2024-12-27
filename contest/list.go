package contest

import (
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
)

func List(cli *client.Client, gym bool) ([]Contest, error) {
	var result = make([]Contest, 0)
	var err = cli.Call("/contest.list", map[string][]string{
		"gym": {strconv.FormatBool(gym)},
	}, &result)

	return result, err
}

type Contest struct {
	ID                    int64        `json:"id"`                    // .ID
	Name                  string       `json:"name"`                  // .Localized.
	Type                  ContestType  `json:"type"`                  // .Scoring system used for the contest.
	Phase                 ContestPhase `json:"phase"`                 // .BEFORE, CODING, PENDING_SYSTEM_TEST, SYSTEM_TEST, FINISHED.
	Frozen                bool         `json:"frozen"`                // .If true, then the ranklist for the contest is frozen and shows only submissions, created before freeze.
	DurationSeconds       int64        `json:"durationSeconds"`       // .Duration of the contest in seconds.
	FreezeDurationSeconds int64        `json:"freezeDurationSeconds"` // .Can be absent. The ranklist freeze duration of the contest in seconds if any.
	StartTimeSeconds      int64        `json:"startTimeSeconds"`      // .Can be absent. Contest start time in unix format.
	RelativeTimeSeconds   int64        `json:"relativeTimeSeconds"`   // .Can be absent. Number of seconds, passed after the start of the contest. Can be negative.
	PreparedBy            string       `json:"preparedBy"`            // .Can be absent. Handle of the user, how created the contest.
	WebsiteUrl            string       `json:"websiteUrl"`            // .Can be absent. URL for contest-related website.
	Description           string       `json:"description"`           // .Localized. Can be absent.
	Difficulty            int64        `json:"difficulty"`            // .Can be absent. From 1 to 5. Larger number means more difficult problems.
	Kind                  string       `json:"kind"`                  // .Kind Localized. Can be absent. Human-readable type of the contest from the following categories: Official ICPC Contest, Official School Contest, Opencup Contest, School/University/City/Region Championship, Training Camp Contest, Official International Personal Contest, Training Contest.
	IcpcRegion            string       `json:"icpcRegion"`            // .IcpcRegion Localized. Can be absent. Name of the Region for official ICPC contests.
	Country               string       `json:"country"`               // .Country Localized. Can be absent.
	City                  string       `json:"city"`                  // .City Localized. Can be absent.
	Season                string       `json:"season"`                // .Season Can be absent.

}

type ContestType string

const (
	CONTEST_TYPE_CF   = "CF"
	CONTEST_TYPE_IOI  = "IOI"
	CONTEST_TYPE_ICPC = "ICPC"
)

type ContestPhase string

const (
	CONTEST_PHASE_BEFORE              = "BEFORE"
	CONTEST_PHASE_CODING              = "CODING"
	CONTEST_PHASE_PENDING_SYSTEM_TEST = "PENDING_SYSTEM_TEST"
	CONTEST_PHASE_SYSTEM_TEST         = "SYSTEM_TEST"
	CONTEST_PHASE_FINISHED            = "FINISHED"
)
