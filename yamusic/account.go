package yamusic

import (
	"context"
	"net/http"
	"time"
)

type (
	// AccountService is a service to deal with accounts.
	AccountService struct {
		client *Client
	}
	// AccountStatus describes account get status method response
	AccountStatus struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result struct {
			Account struct {
				UID              int       `json:"uid"`
				Region           int       `json:"region"`
				Login            string    `json:"login"`
				FullName         string    `json:"fullName"`
				SecondName       string    `json:"secondName"`
				FirstName        string    `json:"firstName"`
				DisplayName      string    `json:"displayName"`
				Birthday         string    `json:"birthday"`
				ServiceAvailable bool      `json:"serviceAvailable"`
				HostedUser       bool      `json:"hostedUser"`
				RegisteredAt     time.Time `json:"registeredAt"`
				Now              time.Time `json:"now"`
				PassportPhones   []struct {
					Phone string `json:"phone"`
				} `json:"passport-phones"`
			} `json:"account"`
			Permissions struct {
				Until   time.Time `json:"until"`
				Values  []string  `json:"values"`
				Default []string  `json:"default"`
			} `json:"permissions"`
			Subscription struct {
				CanStartTrial bool `json:"canStartTrial"`
				Mcdonalds     bool `json:"mcdonalds"`
			} `json:"subscription"`
		} `json:"result"`
	}
)

// GetStatus returns account's status
func (s *AccountService) GetStatus(
	ctx context.Context,
) (*AccountStatus, *http.Response, error) {

	req, err := s.client.NewRequest(http.MethodGet, "account/status", nil)
	if err != nil {
		return nil, nil, err
	}

	accountStatus := new(AccountStatus)
	resp, err := s.client.Do(ctx, req, accountStatus)
	return accountStatus, resp, err
}
