package problemset

import (
	"errors"
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
)

func RecentStatus(cli *client.Client, recentStatusPar RecentStatusPar) ([]Submission, error) {
	var result = make([]Submission, 0)
	par, err := recentStatusPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/problemset.recentStatus", par, &result)
	return result, err
}

type RecentStatusPar struct {
	Count          int64  `json:"count"`          // Number of submissions to return. Can be up to 1000.
	ProblemsetName string `json:"problemsetName"` // Custom problemset's short name, like 'acmsguru'
}

func (sp RecentStatusPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if sp.Count < 0 || sp.Count > 1000 {
		return nil, errors.New("count is invalid")
	}
	result["count"] = []string{strconv.FormatInt(sp.Count, 10)}
	if sp.ProblemsetName != "" {
		result["problemsetName"] = []string{sp.ProblemsetName}
	}
	return result, nil
}

type Submission struct {
	ID                  int64   `json:"id"`                  //.
	ContestID           int64   `json:"contestId"`           //. Can be absent.
	CreationTimeSeconds int64   `json:"creationTimeSeconds"` //. Time, when submission was created, in unix-format.
	RelativeTimeSeconds int64   `json:"relativeTimeSeconds"` //. Number of seconds, passed after the start of the contest (or a virtual start for virtual parties), before the submission.
	Problem             Problem `json:"problem"`             // object.
	Author              Party   `json:"author"`              // object.
	ProgrammingLanguage string  `json:"programmingLanguage"` //.
	Verdict             Verdict `json:"verdict"`             //: FAILED, OK, PARTIAL, COMPILATION_ERROR, RUNTIME_ERROR, WRONG_ANSWER, PRESENTATION_ERROR, TIME_LIMIT_EXCEEDED, MEMORY_LIMIT_EXCEEDED, IDLENESS_LIMIT_EXCEEDED, SECURITY_VIOLATED, CRASHED, INPUT_PREPARATION_CRASHED, CHALLENGED, SKIPPED, TESTING, REJECTED. Can be absent.
	Testset             Testset `json:"testset"`             //: SAMPLES, PRETESTS, TESTS, CHALLENGES, TESTS1, ..., TESTS10. Testset used for judging the submission.
	PassedTestCount     int64   `json:"passedTestCount"`     //. Number of passed tests.
	TimeConsumedMillis  int64   `json:"timeConsumedMillis"`  //. Maximum time in milliseconds, consumed by solution for one test.
	MemoryConsumedBytes int64   `json:"memoryConsumedBytes"` //. Maximum memory in bytes, consumed by solution for one test.
	Points              float64 `json:"points"`              // point number. Can be absent. Number of scored points for IOI-like contests.

}

type Party struct {
	ContestID        int64           `json:"contestId"`        // . Can be absent. Id of the contest, in which party is participating.
	Members          []Member        `json:"members"`          // . list of Member objects. Members of the party.
	ParticipantType  ParticipantType `json:"participantType"`  // : CONTESTANT, PRACTICE, VIRTUAL, MANAGER, OUT_OF_COMPETITION.
	TeamID           int64           `json:"teamId"`           // . Can be absent. If party is a team, then it is a unique team id. Otherwise, this field is absent.
	TeamName         string          `json:"teamName"`         // . Localized. Can be absent. If party is a team or ghost, then it is a localized name of the team. Otherwise, it is absent.
	Ghost            bool            `json:"ghost"`            // . If true then this party is a ghost. It participated in the contest, but not on Codeforces. For example, Andrew Stankevich Contests in Gym has ghosts of the participants from Petrozavodsk Training Camp.
	Room             int64           `json:"room"`             // . Can be absent. Room of the party. If absent, then the party has no room.
	StartTimeSeconds int64           `json:"startTimeSeconds"` // . Can be absent. Time, when this party started a contest.
}

type ParticipantType string

const (
	PARTICIPANT_TYPE_CONTESTANT         ParticipantType = "CONTESTANT"
	PARTICIPANT_TYPE_PRACTICE           ParticipantType = "PRACTICE"
	PARTICIPANT_TYPE_VIRTUAL            ParticipantType = "VIRTUAL"
	PARTICIPANT_TYPE_MANAGER            ParticipantType = "MANAGER"
	PARTICIPANT_TYPE_OUT_OF_COMPETITION ParticipantType = "OUT_OF_COMPETITION"
)

type Member struct {
	Handle string `json:"handle"` // . Codeforces user handle.
	Name   string `json:"name"`   // . Can be absent. User's name if available.

}

type Verdict string

const (
	VERDICT_FAILED                    Verdict = "FAILED"
	VERDICT_OK                        Verdict = "OK"
	VERDICT_PARTIAL                   Verdict = "PARTIAL"
	VERDICT_COMPILATION_ERROR         Verdict = "COMPILATION_ERROR"
	VERDICT_RUNTIME_ERROR             Verdict = "RUNTIME_ERROR"
	VERDICT_WRONG_ANSWER              Verdict = "WRONG_ANSWER"
	VERDICT_PRESENTATION_ERROR        Verdict = "PRESENTATION_ERROR"
	VERDICT_TIME_LIMIT_EXCEEDED       Verdict = "TIME_LIMIT_EXCEEDED"
	VERDICT_MEMORY_LIMIT_EXCEEDED     Verdict = "MEMORY_LIMIT_EXCEEDED"
	VERDICT_IDLENESS_LIMIT_EXCEEDED   Verdict = "IDLENESS_LIMIT_EXCEEDED"
	VERDICT_SECURITY_VIOLATED         Verdict = "SECURITY_VIOLATED"
	VERDICT_CRASHED                   Verdict = "CRASHED"
	VERDICT_INPUT_PREPARATION_CRASHED Verdict = "INPUT_PREPARATION_CRASHED"
	VERDICT_CHALLENGED                Verdict = "CHALLENGED"
	VERDICT_SKIPPED                   Verdict = "SKIPPED"
	VERDICT_TESTING                   Verdict = "TESTING"
	VERDICT_REJECTED                  Verdict = "REJECTED"
)

type Testset string

const (
	TESTSET_SAMPLES    Testset = "SAMPLES"
	TESTSET_PRETESTS   Testset = "PRETESTS"
	TESTSET_TESTS      Testset = "TESTS"
	TESTSET_CHALLENGES Testset = "CHALLENGES"
	TESTSET_TESTS1     Testset = "TESTS1"
	TESTSET_TESTS2     Testset = "TESTS2"
	TESTSET_TESTS3     Testset = "TESTS3"
	TESTSET_TESTS4     Testset = "TESTS4"
	TESTSET_TESTS5     Testset = "TESTS5"
	TESTSET_TESTS6     Testset = "TESTS6"
	TESTSET_TESTS7     Testset = "TESTS7"
	TESTSET_TESTS8     Testset = "TESTS8"
	TESTSET_TESTS9     Testset = "TESTS9"
	TESTSET_TESTS10    Testset = "TESTS10"
)
