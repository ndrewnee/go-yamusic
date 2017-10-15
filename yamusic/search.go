package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type (
	searchType string
)

const (
	searchTypeArtist searchType = "artist"
	searchTypeAlbum             = "album"
	searchTypeTrack             = "track"
	searchTypeAll               = "all"
)

type (
	// SearchService is a service to deal with genres.
	SearchService struct {
		client *Client
	}
	// SearchOptions defines search options like pagination and correction
	SearchOptions struct {
		Page      int
		NoCorrect bool
	}
	// Search describes search method response.
	Search struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result struct {
			MisspellCorrected bool   `json:"misspellCorrected"`
			Nocorrect         bool   `json:"nocorrect"`
			SearchRequestID   string `json:"searchRequestId"`
			Text              string `json:"text"`
			MisspellResult    string `json:"misspellResult"`
			MisspellOriginal  string `json:"misspellOriginal"`
			Best              struct {
				Type   string `json:"type"`
				Result struct {
					ID    int `json:"id"`
					Cover struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Composer bool   `json:"composer"`
					Various  bool   `json:"various"`
					Name     string `json:"name"`
					Counts   struct {
						Tracks       int `json:"tracks"`
						DirectAlbums int `json:"directAlbums"`
						AlsoAlbums   int `json:"alsoAlbums"`
						AlsoTracks   int `json:"alsoTracks"`
					} `json:"counts"`
					Genres           []string `json:"genres"`
					TicketsAvailable bool     `json:"ticketsAvailable"`
					PopularTracks    []struct {
						ID             int  `json:"id"`
						Available      bool `json:"available"`
						AvailableAsRbt bool `json:"availableAsRbt"`
						Albums         []struct {
							ID            int           `json:"id"`
							StorageDir    string        `json:"storageDir"`
							Title         string        `json:"title"`
							Artists       []interface{} `json:"artists"`
							CoverURI      string        `json:"coverUri"`
							TrackCount    int           `json:"trackCount"`
							Genre         string        `json:"genre"`
							Available     bool          `json:"available"`
							TrackPosition struct {
								Volume int `json:"volume"`
								Index  int `json:"index"`
							} `json:"trackPosition"`
						} `json:"albums"`
						StorageDir string        `json:"storageDir"`
						DurationMs int           `json:"durationMs"`
						Explicit   bool          `json:"explicit"`
						Title      string        `json:"title"`
						Artists    []interface{} `json:"artists"`
						Regions    []string      `json:"regions"`
					} `json:"popularTracks"`
					Regions []string `json:"regions"`
				} `json:"result"`
			} `json:"best"`
			Tracks struct {
				Total   int `json:"total"`
				PerPage int `json:"perPage"`
				Results []struct {
					ID             int  `json:"id"`
					Available      bool `json:"available"`
					AvailableAsRbt bool `json:"availableAsRbt"`
					Albums         []struct {
						ID                  int           `json:"id"`
						StorageDir          string        `json:"storageDir"`
						OriginalReleaseYear int           `json:"originalReleaseYear"`
						Year                int           `json:"year"`
						Title               string        `json:"title"`
						Artists             []interface{} `json:"artists"`
						CoverURI            string        `json:"coverUri"`
						TrackCount          int           `json:"trackCount"`
						Genre               string        `json:"genre"`
						Available           bool          `json:"available"`
						TrackPosition       struct {
							Volume int `json:"volume"`
							Index  int `json:"index"`
						} `json:"trackPosition"`
					} `json:"albums"`
					StorageDir string `json:"storageDir"`
					DurationMs int    `json:"durationMs"`
					Explicit   bool   `json:"explicit"`
					Title      string `json:"title"`
					Artists    []struct {
						ID    int `json:"id"`
						Cover struct {
							Type   string `json:"type"`
							Prefix string `json:"prefix"`
							URI    string `json:"uri"`
						} `json:"cover"`
						Composer   bool          `json:"composer"`
						Various    bool          `json:"various"`
						Name       string        `json:"name"`
						Decomposed []interface{} `json:"decomposed"`
					} `json:"artists"`
					Regions []string `json:"regions"`
					Version string   `json:"version,omitempty"`
				} `json:"results"`
			} `json:"tracks"`
			Playlists struct {
				Total   int `json:"total"`
				PerPage int `json:"perPage"`
				Results []struct {
					UID        int    `json:"uid"`
					Kind       int    `json:"kind"`
					TrackCount int    `json:"trackCount"`
					Title      string `json:"title"`
					Owner      struct {
						UID      int    `json:"uid"`
						Login    string `json:"login"`
						Name     string `json:"name"`
						Verified bool   `json:"verified"`
					} `json:"owner"`
					Cover struct {
						Type     string   `json:"type"`
						ItemsURI []string `json:"itemsUri"`
						Custom   bool     `json:"custom"`
					} `json:"cover"`
					Tags    []interface{} `json:"tags"`
					Regions []string      `json:"regions"`
				} `json:"results"`
			} `json:"playlists"`
			Artists struct {
				Total   int `json:"total"`
				PerPage int `json:"perPage"`
				Results []struct {
					ID    int `json:"id"`
					Cover struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Composer bool   `json:"composer"`
					Various  bool   `json:"various"`
					Name     string `json:"name"`
					Counts   struct {
						Tracks       int `json:"tracks"`
						DirectAlbums int `json:"directAlbums"`
						AlsoAlbums   int `json:"alsoAlbums"`
						AlsoTracks   int `json:"alsoTracks"`
					} `json:"counts"`
					Genres           []string `json:"genres"`
					TicketsAvailable bool     `json:"ticketsAvailable"`
					PopularTracks    []struct {
						ID             int  `json:"id"`
						Available      bool `json:"available"`
						AvailableAsRbt bool `json:"availableAsRbt"`
						Albums         []struct {
							ID            int           `json:"id"`
							StorageDir    string        `json:"storageDir"`
							Title         string        `json:"title"`
							Artists       []interface{} `json:"artists"`
							CoverURI      string        `json:"coverUri"`
							TrackCount    int           `json:"trackCount"`
							Genre         string        `json:"genre"`
							Available     bool          `json:"available"`
							TrackPosition struct {
								Volume int `json:"volume"`
								Index  int `json:"index"`
							} `json:"trackPosition"`
						} `json:"albums"`
						StorageDir string        `json:"storageDir"`
						DurationMs int           `json:"durationMs"`
						Explicit   bool          `json:"explicit"`
						Title      string        `json:"title"`
						Artists    []interface{} `json:"artists"`
						Regions    []string      `json:"regions"`
					} `json:"popularTracks"`
					Regions []string `json:"regions"`
				} `json:"results"`
			} `json:"artists"`
			Videos struct {
				Total   int `json:"total"`
				PerPage int `json:"perPage"`
				Results []struct {
					YoutubeURL              string   `json:"youtubeUrl"`
					ThumbnailURL            string   `json:"thumbnailUrl"`
					Title                   string   `json:"title"`
					Duration                int      `json:"duration"`
					Text                    string   `json:"text"`
					HTMLAutoPlayVideoPlayer string   `json:"htmlAutoPlayVideoPlayer"`
					Regions                 []string `json:"regions"`
				} `json:"results"`
			} `json:"videos"`
			Albums struct {
				Total   int `json:"total"`
				PerPage int `json:"perPage"`
				Results []struct {
					ID                  int    `json:"id"`
					StorageDir          string `json:"storageDir"`
					OriginalReleaseYear int    `json:"originalReleaseYear"`
					Year                int    `json:"year"`
					Title               string `json:"title"`
					Artists             []struct {
						ID    int `json:"id"`
						Cover struct {
							Type   string `json:"type"`
							Prefix string `json:"prefix"`
							URI    string `json:"uri"`
						} `json:"cover"`
						Composer   bool          `json:"composer"`
						Various    bool          `json:"various"`
						Name       string        `json:"name"`
						Decomposed []interface{} `json:"decomposed"`
					} `json:"artists"`
					CoverURI   string   `json:"coverUri"`
					TrackCount int      `json:"trackCount"`
					Genre      string   `json:"genre"`
					Available  bool     `json:"available"`
					Regions    []string `json:"regions"`
				} `json:"results"`
			} `json:"albums"`
		} `json:"result"`
	}
)

// Artists searches artists by query
func (s *SearchService) Artists(ctx context.Context, query string, opts *SearchOptions) (*Search, *http.Response, error) {
	result := new(Search)
	resp, err := s.search(ctx, searchTypeArtist, query, opts, result)
	return result, resp, err
}

// Tracks searches tracks by query
func (s *SearchService) Tracks(ctx context.Context, query string, opts *SearchOptions) (*Search, *http.Response, error) {
	result := new(Search)
	resp, err := s.search(ctx, searchTypeTrack, query, opts, result)
	return result, resp, err
}

// Albums searches albums by query
func (s *SearchService) Albums(ctx context.Context, query string, opts *SearchOptions) (*Search, *http.Response, error) {
	result := new(Search)
	resp, err := s.search(ctx, searchTypeAlbum, query, opts, result)
	return result, resp, err
}

// All searches all(artists, albums, tracks) by query
func (s *SearchService) All(ctx context.Context, query string, opts *SearchOptions) (*Search, *http.Response, error) {
	result := new(Search)
	resp, err := s.search(ctx, searchTypeAll, query, opts, result)
	return result, resp, err
}

func (s *SearchService) search(ctx context.Context, searchTyp searchType, query string, opts *SearchOptions, result interface{}) (*http.Response, error) {
	if opts == nil {
		opts = &SearchOptions{}
	}

	queryParams := url.Values{
		"type":      []string{string(searchTyp)},
		"text":      []string{query},
		"page":      []string{strconv.Itoa(opts.Page)},
		"nocorrect": []string{strconv.FormatBool(opts.NoCorrect)},
	}

	uri := fmt.Sprintf("search?%v", queryParams.Encode())

	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, result)
}
