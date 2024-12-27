package contest

import (
	"errors"
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
	"github.com/cloudfstrife/go-codeforces/problemset"
)

func Status(cli *client.Client, statusPar StatusPar) ([]StatusResult, error) {
	var result = make([]StatusResult, 0)
	par, err := statusPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/contest.list", par, &result)
	return result, err
}

type StatusPar struct {
	ContestID int64  `json:"contestId"` // . Id of the contest. It is not the round number. It can be seen in contest URL. For example: /contest/566/status
	AsManager bool   `json:"asManager"` // . If set to true, the response will contain information available to contest managers. Otherwise, the response will contain only the information available to the participants. You must be a contest manager to use it.
	Handle    string `json:"handle"`    // . Codeforces user handle.
	From      int64  `json:"from"`      // . 1-based index of the first submission to return.
	Count     int64  `json:"count"`     // . Number of returned submissions.
}

func (sp StatusPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if sp.ContestID == 0 {
		return result, errors.New("contest id is zero")
	}
	result["contestId"] = []string{strconv.FormatInt(sp.ContestID, 10)}
	if sp.AsManager {
		result["asManager"] = []string{strconv.FormatBool(sp.AsManager)}
	}
	if sp.Handle != "" {
		result["handle"] = []string{sp.Handle}
	}
	if sp.From != 0 {
		result["from"] = []string{strconv.FormatInt(sp.From, 10)}
	}
	if sp.Count != 0 {
		result["count"] = []string{strconv.FormatInt(sp.Count, 10)}
	}
	return result, nil
}

type StatusResult struct {
	ID                  int64               `json:"id"`                  // . ID
	ContestID           int64               `json:"contestId"`           // . Can be absent.
	CreationTimeSeconds int64               `json:"creationTimeSeconds"` // . Time, when submission was created, in unix-format.
	RelativeTimeSeconds int64               `json:"relativeTimeSeconds"` // . Number of seconds, passed after the start of the contest (or a virtual start for virtual parties), before the submission.
	Problem             *problemset.Problem `json:"problem"`             // . object.
	Author              problemset.Party    `json:"author"`              // . object.
	ProgrammingLanguage string              `json:"programmingLanguage"` // . Programming Language
	Verdict             problemset.Verdict  `json:"verdict"`             // : FAILED, OK, PARTIAL, COMPILATION_ERROR, RUNTIME_ERROR, WRONG_ANSWER, PRESENTATION_ERROR, TIME_LIMIT_EXCEEDED, MEMORY_LIMIT_EXCEEDED, IDLENESS_LIMIT_EXCEEDED, SECURITY_VIOLATED, CRASHED, INPUT_PREPARATION_CRASHED, CHALLENGED, SKIPPED, TESTING, REJECTED. Can be absent.
	Testset             problemset.Testset  `json:"testset"`             // : SAMPLES, PRETESTS, TESTS, CHALLENGES, TESTS1, ..., TESTS10. Testset used for judging the submission.
	PassedTestCount     int64               `json:"passedTestCount"`     // . Number of passed tests.
	TimeConsumedMillis  int64               `json:"timeConsumedMillis"`  // . Maximum time in milliseconds, consumed by solution for one test.
	MemoryConsumedBytes int64               `json:"memoryConsumedBytes"` // . Maximum memory in bytes, consumed by solution for one test.
	Points              float64             `json:"points"`              // . point number. Can be absent. Number of scored points for IOI-like contests.

}
