package contest

import (
	"errors"
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
	"github.com/cloudfstrife/go-codeforces/problemset"
)

func Hacks(cli *client.Client, hacksPar HacksPar) ([]HacksResult, error) {
	var result = make([]HacksResult, 0)
	par, err := hacksPar.ParseToCFPar()
	if err != nil {
		return nil, err
	}
	err = cli.Call("/contest.hacks", par, &result)
	return result, err
}

type HacksPar struct {
	ContestID int64 `json:"contestId"` // . Id of the contest. It is not the round number. It can be seen in contest URL. For example: /contest/566/status
	AsManager bool  `json:"asManager"` // . If set to true, the response will contain information available to contest managers. Otherwise, the response will contain only the information available to the participants. You must be a contest manager to use it.
}

func (sp HacksPar) ParseToCFPar() (map[string][]string, error) {
	var result = make(map[string][]string)
	if sp.ContestID == 0 {
		return result, errors.New("contest id is zero")
	}
	result["contestId"] = []string{strconv.FormatInt(sp.ContestID, 10)}
	if sp.AsManager {
		result["asManager"] = []string{strconv.FormatBool(sp.AsManager)}
	}
	return result, nil
}

type HacksResult struct {
	ID                  int64               `json:"id"`                  // .
	CreationTimeSeconds int64               `json:"creationTimeSeconds"` // . Hack creation time in unix format.
	Hacker              problemset.Party    `json:"hacker"`              //  object.
	Defender            problemset.Party    `json:"defender"`            //  object.
	Verdict             problemset.Verdict  `json:"verdict"`             // : HACK_SUCCESSFUL, HACK_UNSUCCESSFUL, INVALID_INPUT, GENERATOR_INCOMPILABLE, GENERATOR_CRASHED, IGNORED, TESTING, OTHER. Can be absent.
	Problem             *problemset.Problem `json:"problem"`             //  object. Hacked problem.
	Test                string              `json:"test"`                // . Can be absent.
	JudgeProtocol       JudgeProtocol       `json:"judgeProtocol"`       //  with three fields: "manual", "protocol" and "verdict". Field manual can have values "true" and "false". If manual is "true" then test for the hack was entered manually. Fields "protocol" and "verdict" contain human-readable description of judge protocol and hack verdict. Localized. Can be absent.

}

type JudgeProtocol struct {
	Manual   string `json:"manual"`
	Protocol string `json:"protocol"`
	Verdict  string `json:"verdict"`
}
