package yamusic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type (
	// TracksService is a service to deal with tracks
	TracksService struct {
		client *Client
	}
	TrackResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         []struct {
			ID     string `json:"id"`
			RealID string `json:"realId"`
			Title  string `json:"title"`
			Major  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"major"`
			Available                      bool   `json:"available"`
			AvailableForPremiumUsers       bool   `json:"availableForPremiumUsers"`
			AvailableFullWithoutPermission bool   `json:"availableFullWithoutPermission"`
			StorageDir                     string `json:"storageDir"`
			DurationMs                     int    `json:"durationMs"`
			FileSize                       int    `json:"fileSize"`
			R128                           struct {
				I  float64 `json:"i"`
				Tp float64 `json:"tp"`
			} `json:"r128"`
			PreviewDurationMs int `json:"previewDurationMs"`
			Artists           []struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Various  bool   `json:"various"`
				Composer bool   `json:"composer"`
				Cover    struct {
					Type   string `json:"type"`
					URI    string `json:"uri"`
					Prefix string `json:"prefix"`
				} `json:"cover"`
				Genres []interface{} `json:"genres"`
			} `json:"artists"`
			Albums []struct {
				ID            int           `json:"id"`
				Title         string        `json:"title"`
				MetaType      string        `json:"metaType"`
				Year          int           `json:"year"`
				ReleaseDate   time.Time     `json:"releaseDate"`
				CoverURI      string        `json:"coverUri"`
				OgImage       string        `json:"ogImage"`
				Genre         string        `json:"genre"`
				Buy           []interface{} `json:"buy"`
				TrackCount    int           `json:"trackCount"`
				LikesCount    int           `json:"likesCount"`
				Recent        bool          `json:"recent"`
				VeryImportant bool          `json:"veryImportant"`
				Artists       []struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Various  bool   `json:"various"`
					Composer bool   `json:"composer"`
					Cover    struct {
						Type   string `json:"type"`
						URI    string `json:"uri"`
						Prefix string `json:"prefix"`
					} `json:"cover"`
					Genres []interface{} `json:"genres"`
				} `json:"artists"`
				Labels []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"labels"`
				Available                bool  `json:"available"`
				AvailableForPremiumUsers bool  `json:"availableForPremiumUsers"`
				AvailableForMobile       bool  `json:"availableForMobile"`
				AvailablePartially       bool  `json:"availablePartially"`
				Bests                    []int `json:"bests"`
				TrackPosition            struct {
					Volume int `json:"volume"`
					Index  int `json:"index"`
				} `json:"trackPosition"`
			} `json:"albums"`
			CoverURI         string `json:"coverUri"`
			OgImage          string `json:"ogImage"`
			LyricsAvailable  bool   `json:"lyricsAvailable"`
			Type             string `json:"type"`
			RememberPosition bool   `json:"rememberPosition"`
			TrackSharingFlag string `json:"trackSharingFlag"`
			LyricsInfo       struct {
				HasAvailableSyncLyrics bool `json:"hasAvailableSyncLyrics"`
				HasAvailableTextLyrics bool `json:"hasAvailableTextLyrics"`
			} `json:"lyricsInfo"`
			HasSyncLyrics bool `json:"hasSyncLyrics"`
			HasLyrics     bool `json:"hasLyrics"`
		} `json:"result"`
	}
	// Response of track/%d/download_info
	DownloadInfoResp struct {
		InvocationInfo InvocationInfo
		Error          Error `json:"error"`
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
	var sign = md5.Sum([]byte("XGRlBW9FXlekgbPrRHuSiA" + dlInfo.Path[1:] + dlInfo.S))
	uri := fmt.Sprintf(
		"https://%s/get-mp3/%s/%s%s",
		dlInfo.Host,
		hex.EncodeToString(sign[:15]), // from MarshalX/yandex_music_api
		dlInfo.TS, dlInfo.Path,
	)
	return uri, nil
}
