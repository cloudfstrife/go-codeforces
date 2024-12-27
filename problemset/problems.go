package problemset

import "github.com/cloudfstrife/go-codeforces/client"

func Problems(cli *client.Client, problemsPar ProblemsPar) (*ProblemsResult, error) {
	var result = ProblemsResult{}
	par, err := problemsPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/problemset.problems", par, &result)
	return &result, err
}

type ProblemsPar struct {
	Tags           []string `json:"tags"`           // 	Semicilon-separated list of tags.
	ProblemsetName string   `json:"problemsetName"` // 	Custom problemset's short name, like 'acmsguru'
}

func (sp ProblemsPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if len(sp.Tags) > 0 {
		result["tags"] = sp.Tags
	}
	if sp.ProblemsetName != "" {
		result["problemsetName"] = []string{sp.ProblemsetName}
	}
	return result, nil
}

type ProblemsResult struct {
	Problems          []Problem          `json:"problems"`
	ProblemStatistics []ProblemStatistic `json:"problemStatistics"`
}

type Problem struct {
	ContestID      int64       `json:"contestId"`      // . Can be absent. Id of the contest, containing the problem.
	ProblemsetName string      `json:"problemsetName"` // . Can be absent. Short name of the problemset the problem belongs to.
	Index          string      `json:"index"`          // . Usually, a letter or letter with digit(s) indicating the problem index in a contest.
	Name           string      `json:"name"`           // . Localized.
	Type           ProblemType `json:"type"`           // . problem type  PROGRAMMING, QUESTION.
	Points         float64     `json:"points"`         // . point number. Can be absent. Maximum amount of points for the problem.
	Rating         int         `json:"rating"`         // . Can be absent. Problem rating (difficulty).
	Tags           []string    `json:"tags"`           // . Problem tags.
}

type ProblemType string

const (
	ProblemType_PROGRAMMING ProblemType = "PROGRAMMING"
	ProblemType_QUESTION    ProblemType = "QUESTION"
)

type ProblemStatistic struct {
	ContestID   int64  `json:"contestId"`   // . Can be absent. Id of the contest, containing the problem.
	Index       string `json:"index"`       // . Usually, a letter or letter with digit(s) indicating the problem index in a contest.
	SolvedCount int64  `json:"solvedCount"` // . Number of users, who solved the problem.
}
