package contest

import (
	"errors"
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
	"github.com/cloudfstrife/go-codeforces/problemset"
)

func Standings(cli *client.Client, standingsPar StandingsPar) (*StandingsResult, error) {
	var result = StandingsResult{}
	par, err := standingsPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/contest.standings", par, &result)
	return &result, err
}

type StandingsPar struct {
	ContestID        int64                      `json:"contestId"`        // (Required) 	Id of the contest. It is not the round number. It can be seen in contest URL. For example: /contest/566/status
	AsManager        bool                       `json:"asManager"`        // Boolean. If set to true, the response will contain information available to contest managers. Otherwise, the response will contain only the information available to the participants. You must be a contest manager to use it.
	From             int64                      `json:"from"`             // 1-based index of the standings row to start the ranklist.
	Count            int64                      `json:"count"`            // Number of standing rows to return.
	Handles          string                     `json:"handles"`          // Semicolon-separated list of handles. No more than 10000 handles is accepted.
	Room             string                     `json:"room"`             // If specified, than only participants from this room will be shown in the result. If not — all the participants will be shown.
	ShowUnofficial   bool                       `json:"showUnofficial"`   // If true than all participants (virtual, out of competition) are shown. Otherwise, only official contestants are shown.
	ParticipantTypes problemset.ParticipantType `json:"participantTypes"` // Comma-separated list of participant types without spaces. Possible values: CONTESTANT, PRACTICE, VIRTUAL, MANAGER, OUT_OF_COMPETITION. Only participants with the specified types will be displayed.
}

func (sp StandingsPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if sp.ContestID == 0 {
		return result, errors.New("contest id is zero")
	}
	result["contestId"] = []string{strconv.FormatInt(sp.ContestID, 10)}
	if sp.AsManager {
		result["asManager"] = []string{strconv.FormatBool(sp.AsManager)}
	}

	if sp.From != 0 {
		result["from"] = []string{strconv.FormatInt(sp.From, 10)}
	}
	if sp.Count != 0 {
		result["count"] = []string{strconv.FormatInt(sp.Count, 10)}
	}
	if sp.Handles != "" {
		result["gandles"] = []string{sp.Handles}
	}
	if sp.Room != "" {
		result["room"] = []string{sp.Room}
	}
	if sp.ShowUnofficial {
		result["showUnofficial"] = []string{strconv.FormatBool(sp.ShowUnofficial)}
	}
	if sp.ParticipantTypes != "" {
		result["participantTypes"] = []string{string(sp.ParticipantTypes)}
	}

	return result, nil
}

type StandingsResult struct {
	Contest  *Contest             `json:"contest"`
	Problems []problemset.Problem `json:"problems"`
	Rows     []RanklistRow        `json:"rows"`
}

type RanklistRow struct {
	Party                     problemset.Party `json:"party"`                     // . Party that took a corresponding place in the contest.
	Rank                      int64            `json:"rank"`                      // . Party place in the contest.
	Points                    float64          `json:"points"`                    // . point number. Total amount of points, scored by the party.
	Penalty                   int64            `json:"penalty"`                   // . Total penalty (in ICPC meaning) of the party.
	SuccessfulHackCount       int64            `json:"successfulHackCount"`       // .
	UnsuccessfulHackCount     int64            `json:"unsuccessfulHackCount"`     // .
	ProblemResults            []ProblemResult  `json:"problemResults"`            // . list of ProblemResult objects. Party results for each problem. Order of the problems is the same as in "problems" field of the returned object.
	LastSubmissionTimeSeconds int64            `json:"lastSubmissionTimeSeconds"` // . For IOI contests only. Time in seconds from the start of the contest to the last submission that added some points to the total score of the party. Can be absent.
}

type ProblemResult struct {
	Points                    float64           `json:"points"`                    // . Floating point number.
	Penalty                   int64             `json:"penalty"`                   // . Integer. Penalty (in ICPC meaning) of the party for this problem. Can be absent.
	RejectedAttemptCount      int64             `json:"rejectedAttemptCount"`      // . Integer. Number of incorrect submissions.
	ProblemResultType         ProblemResultType `json:"type"`                      // . Enum: PRELIMINARY, FINAL. If type is PRELIMINARY then points can decrease (if, for example, solution will fail during system test). Otherwise, party can only increase points for this problem by submitting better solutions.
	BestSubmissionTimeSeconds int64             `json:"bestSubmissionTimeSeconds"` // . Integer. Number of seconds after the start of the contest before the submission, that brought maximal amount of points for this problem. Can be absent.
}

type ProblemResultType string

const (
	ProblemResultType_PRELIMINARY ProblemResultType = "PRELIMINARY"
	ProblemResultType_FINAL       ProblemResultType = "FINAL"
)
