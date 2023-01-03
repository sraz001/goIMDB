package goIMDB

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"io"
	"net/http"
	"net/url"
	"time"
)

/*
Authentication and encryption of Http calls to the API.

The api is hosted on Amazon, so the requests must be signed after getting a valid login credentials.
The login credentials are kept until they expire to avoid repeated calls to the api just to login.

Using the AWS official sdk for the crypto part.
todo: keep credentials in a temp file (useful for CLI usage)
*/

// holds various constants used in the program
const (
	BaseUri       = "https://api.imdbws.com"
	SearchBaseUri = `https://v2.sg.media-imdb.com`
	UserAgent     = `IMDb/8.3.1 (iPhone9,4; iOS 11.2.1)`
	AppKey        = `76a6cc20-6073-4290-8a2c-951b4580ae4a`
	SoonExpire    = time.Minute * 5
	ServiceName   = "imdbapi"
	ServiceRegion = "us-east-1"
)

// json equivalent struct to get the temporary session/access information
type authJson struct {
	Resource struct {
		Type                string    `json:"@type"`
		AccessKeyId         string    `json:"accessKeyId"`
		ExpirationTimeStamp time.Time `json:"expirationTimeStamp"`
		SecretAccessKey     string    `json:"secretAccessKey"`
		SessionToken        string    `json:"sessionToken"`
	} `json:"resource"`
}

// authImdb keep current credentials and expiration time
type authImdb struct {
	cred   *credentials.Credentials
	expiry time.Time
}

// IMDBPie - contain information about the parameters of the package.  No exposed fields
type IMDBPie struct {
	auth      authImdb
	c         *http.Client
	excludeEp bool
	region    string
}

// InitIMDBPie - create the api structure that allows the searchs.
func InitIMDBPie() *IMDBPie {
	return &IMDBPie{region: "US"}
}

// ExcludeEpisodes  - set a flag to exclude the episodes from the result sets (use if you don't need such details)
func (m *IMDBPie) ExcludeEpisodes(s bool) {
	m.excludeEp = s
}

// SetRegion - set the region, default is 'en_US'
func (m *IMDBPie) SetRegion(r string) {
	m.region = r
}

// urlGetCompose create urls from the get enpoints with imdb code
func urlGetCompose(enpoint string, im string) string {
	u, e := url.JoinPath(BaseUri, fmt.Sprintf(enpoint, im))
	if e != nil {
		panic(e)
	}
	return u
}

// getWebCredentials get temporary credentials from aws and return them as AWS sdk credentials
func (m *IMDBPie) getWebCredentials() error {
	url := fmt.Sprintf("%s/authentication/credentials/temporary/ios82", BaseUri)
	jsonStr := []byte(fmt.Sprintf(`{"appKey": "%s"}`, AppKey))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", UserAgent)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var aJson = authJson{}
	err = json.Unmarshal(body, &aJson)
	if err != nil {
		return err
	}
	//fmt.Println(aJson.Resource.ExpirationTimeStamp, time.Now(), "\n", aJson.Resource.SessionToken, aJson.Resource.SecretAccessKey, aJson.Resource.ExpirationTimeStamp)
	m.auth.cred = credentials.NewStaticCredentials(aJson.Resource.AccessKeyId, aJson.Resource.SecretAccessKey, aJson.Resource.SessionToken)
	m.auth.expiry = aJson.Resource.ExpirationTimeStamp
	return nil
}

// getCredentials -  Get either the existing or new credentials from aws.
func (m *IMDBPie) getCredentials() error {
	if m.auth.cred == nil || time.Now().UTC().Sub(m.auth.expiry) < SoonExpire {
		return m.getWebCredentials()
	}
	return nil
}

// makeImdbRequest - Make a request to the main API - with possible parameters, body & method
// return http Response. All requests are signed with the AWS crypto library
func (m *IMDBPie) makeImdbRequest(method string, baseurl string, p *map[string]string, body *[]byte) (*http.Response, error) {
	e := m.getCredentials()
	if e != nil {
		return nil, e
	}
	c := &http.Client{}
	var req *http.Request
	var r *bytes.Reader
	if body == nil || len(*body) == 0 {
		req, e = http.NewRequest(method, baseurl, nil)
	} else {
		r = bytes.NewReader(*body)
		req, e = http.NewRequest(method, baseurl, r)
	}
	req.Header.Set("User-Agent", UserAgent)
	if e != nil {
		return nil, e
	}
	q := req.URL.Query()
	if p != nil {
		for k, v := range *p {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
	signer := v4.Signer{Credentials: m.auth.cred}
	_, e = signer.Sign(req, nil, ServiceName, ServiceRegion, time.Now().UTC())
	if e != nil {
		return nil, e
	}
	resp, e := c.Do(req)
	if e != nil {
		return resp, e
	}
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New(fmt.Sprintf("bad status from http Get: %d", resp.StatusCode))
	}
	return resp, e
}
