package client

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/cloudfstrife/log"
	"go.uber.org/zap"
)

type Client struct {
	conf *Config
	hc   *http.Client
}

type Response struct {
	Status  string          `json:"status"`
	Comment string          `json:"comment,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func NewClient(conf *Config) *Client {
	return &Client{
		conf: conf,
		hc:   &http.Client{},
	}
}

func (c *Client) Sig(method string, params map[string][]string) (string, error) {
	if c.conf == nil || c.conf.ApiKey == "" || c.conf.ApiSecret == "" {
		return "", errors.New("apiKey or apiSecret is blank")
	}
	q := url.Values{}
	for k, v := range params {
		q.Set(k, strings.Join(v, ";"))
	}
	r := fmt.Sprintf("%06d", rand.Intn(1000000))
	sum := sha512.Sum512([]byte(fmt.Sprintf("%s/%s?%s#%s", r, method, q.Encode(), c.conf.ApiSecret)))
	asig := hex.EncodeToString(sum[:])
	return r + asig, nil
}

func (c *Client) Call(apiURL string, params map[string][]string, result interface{}) error {
	var err error

	var u *url.URL
	if u, err = url.Parse(CF_URL); err != nil {
		log.Error("parse url failed", zap.Error(err))
		return err
	}
	u.Path = path.Join(u.Path, apiURL)

	if c.conf != nil && c.conf.Lang != "" {
		params["lang"] = []string{c.conf.Lang}
	}
	if c.conf != nil && c.conf.ApiKey != "" && c.conf.ApiSecret != "" {
		params["time"] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
		params["apiKey"] = []string{c.conf.ApiKey}
		apiSig, err := c.Sig(apiURL, params)
		if err != nil {
			log.Error("sig codeforces failed", zap.Error(err))
			return err
		}
		params["apiSig"] = []string{apiSig}
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, strings.Join(v, ";"))
	}
	u.RawQuery = q.Encode()

	log.Debug("call codeforces api", zap.String("url", u.String()))

	var resp *http.Response
	if resp, err = c.hc.Get(u.String()); err != nil {
		log.Error("call codeforces failed", zap.Error(err))
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.Error("call codeforces response statue is not OK", zap.String("url", u.String()), zap.String("status", resp.Status))
		return fmt.Errorf("call codeforces api : %s response unsuccess", apiURL)
	}

	defer resp.Body.Close()
	var rspBody []byte
	if rspBody, err = io.ReadAll(resp.Body); err != nil {
		log.Error("read codeforces response failed", zap.Error(err))
		return err
	}

	log.Debug("Call codeforces api response", zap.String("url", u.String()), zap.String("result", string(rspBody)))

	var res Response
	if err = json.Unmarshal(rspBody, &res); err != nil {
		log.Error("unmarshal codeforces response failed", zap.Error(err))
		return err
	}

	if res.Status == "FAILED" {
		return errors.New(res.Comment)
	}

	return json.Unmarshal(res.Result, result)
}
