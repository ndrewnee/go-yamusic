package yamusic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"net/http"
)

type (
	// TracksService is a service to deal with tracks
	TracksService struct {
		client *Client
	}
	TrackResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         []Track        `json:"result"`
	}
	// Response of track/%d/download_info
	DownloadInfoResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         []struct {
			Codec           string `json:"codec"`
			Gain            bool   `json:"gain"`
			Preview         bool   `json:"preview"`
			DownloadInfoURL string `json:"downloadInfoUrl"`
			Direct          bool   `json:"direct"`
			BitrateInKbps   int    `json:"bitrateInKbps"`
		} `json:"result"`
	}
	// DownloadInfo is a response of URL from DownloadInfoResp's `DownloadInfoURL` field
	DownloadInfo struct {
		XMLName xml.Name `xml:"download-info"`
		Text    string   `xml:",chardata"`
		Host    string   `xml:"host"`
		Path    string   `xml:"path"`
		TS      string   `xml:"ts"`
		Region  string   `xml:"region"`
		S       string   `xml:"s"`
	}
)

type TrackError string

func (te TrackError) Error() string { return string(te) }

var (
	ErrNilDownloadInfoResp = TrackError("got nil download info resp pointer")
	ErrNilDownloadInfo     = TrackError("got nil download info pointer")
	ErrNilPath             = TrackError("got nil path")
	ErrEmptyPath           = TrackError("got empty path")
	ErrZeroResultLen       = TrackError("len of download inf response's result field is zero")
)

// Get returns track by its ID
func (t *TracksService) Get(ctx context.Context, id int) (*TrackResp, *http.Response, error) {
	uri := fmt.Sprintf("tracks/%v", id)
	req, err := t.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}
	track := new(TrackResp)
	resp, err := t.client.Do(ctx, req, track)
	return track, resp, err
}

// GetDownloadInfoResp returns DownloadInfoResp byt track's ID
// Be careful: you can get DownloadInfo by DownloadInfoURL only
// for one minute since you called GetDownloadInfoResp
func (t *TracksService) GetDownloadInfoResp(ctx context.Context, id int) (*DownloadInfoResp, *http.Response, error) {
	uri := fmt.Sprintf("tracks/%v/download-info", id)
	req, err := t.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	dlInfoResp := new(DownloadInfoResp)
	resp, err := t.client.Do(ctx, req, dlInfoResp)
	return dlInfoResp, resp, err
}

// GetDownloadInfo returns DownloadInfo by id of track.
// Be careful: it uses the same context for GetDownloadInfoResp
// and its request
func (t *TracksService) GetDownloadInfo(ctx context.Context, id int) (*DownloadInfo, *http.Response, error) {
	dlInfoResp, dirResp, err := t.GetDownloadInfoResp(ctx, id)
	if err != nil {
		return nil, dirResp, err
	}

	if dlInfoResp == nil {
		return nil, nil, ErrNilDownloadInfoResp
	}

	if len(dlInfoResp.Result) == 0 {
		return nil, nil, ErrZeroResultLen
	}

	req, err := t.client.NewRequest(http.MethodGet, dlInfoResp.Result[0].DownloadInfoURL, nil)

	dlInfo := new(DownloadInfo)
	resp, err := t.client.Do(ctx, req, dlInfo)
	return dlInfo, resp, err
}

// GetDownloadURL computes path to track by ID
func (t *TracksService) GetDownloadURL(ctx context.Context, id int) (string, error) {
	dlInfo, _, err := t.GetDownloadInfo(ctx, id)
	if err != nil {
		return "", err
	}
	if dlInfo == nil {
		return "", ErrNilDownloadInfo
	} else if len(dlInfo.Path) == 0 {
		return "", ErrEmptyPath
	}
	// a bit of magic
	const signPrefix = "XGRlBW9FXlekgbPrRHuSiA"
	var sign = md5.Sum([]byte(signPrefix + dlInfo.Path[1:] + dlInfo.S))
	uri := fmt.Sprintf(
		"https://%s/get-mp3/%s/%s%s",
		dlInfo.Host,
		hex.EncodeToString(sign[:]),
		dlInfo.TS, dlInfo.Path,
	)
	return uri, nil
}
