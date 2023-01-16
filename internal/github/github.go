package github

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) ExchangeCodeToAccessKey(ctx context.Context, clientID, clientSecret, code string) (string, error) {
	type exchangeCodeRequest struct {
		ClientID     string `json:"client_id,omitempty"`
		ClientSecret string `json:"client_secret,omitempty"`
		Code         string `json:"code,omitempty"`
	}

	req := exchangeCodeRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://github.com/login/oauth/access_token",
		bytes.NewReader(reqJSON),
	)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	var respJSON struct {
		AccessToken string `json:"access_token,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respJSON); err != nil {
		return "", err
	}

	return respJSON.AccessToken, nil
}

func (c *Client) IsMember(ctx context.Context, accessKey, org, user string) (bool, error) {
	githhubClient := getGithubClientWithAccessKey(ctx, accessKey)

}

func getGithubClientWithAccessKey(ctx context.Context, accessKey string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessKey})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
