package yamusic

import (
	"context"
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
		Error struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

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
