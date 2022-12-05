package yamusic

import (
	"context"
	"net/http"
	"time"
)

type (
	// FeedService is a service to deal with accounts.
	FeedService struct {
		client *Client
	}

	// FeedResp describes get feed method response
	FeedResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			CanGetMoreEvents   bool          `json:"canGetMoreEvents"`
			Pumpkin            bool          `json:"pumpkin"`
			Today              string        `json:"today"`
			GeneratedPlaylists []interface{} `json:"generatedPlaylists"`
			Headlines          []struct {
				Type    string `json:"type"`
				ID      string `json:"id"`
				Message string `json:"message"`
			} `json:"headlines"`
			Days []struct {
				Day    string `json:"day"`
				Events []struct {
					ID          string `json:"id"`
					Type        string `json:"type"`
					TypeForFrom string `json:"typeForFrom,omitempty"`
					Title       []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"title,omitempty"`
					Promo struct {
						PromoID       string        `json:"promoId"`
						Category      string        `json:"category"`
						TitleURL      string        `json:"titleUrl"`
						SubtitleURL   string        `json:"subtitleUrl"`
						Title         string        `json:"title"`
						Subtitle      string        `json:"subtitle"`
						Heading       string        `json:"heading"`
						Description   string        `json:"description"`
						Background    string        `json:"background"`
						ImagePosition string        `json:"imagePosition"`
						PromotionType string        `json:"promotionType"`
						Tags          []interface{} `json:"tags"`
						StartDate     time.Time     `json:"startDate"`
						Pager         struct {
							Total   int `json:"total"`
							Page    int `json:"page"`
							PerPage int `json:"perPage"`
						} `json:"pager"`
						Playlists []struct {
							Playlist struct {
								UID                  int       `json:"uid"`
								Kind                 int       `json:"kind"`
								Revision             int       `json:"revision"`
								TrackCount           int       `json:"trackCount"`
								DurationMs           int       `json:"durationMs"`
								Collective           bool      `json:"collective"`
								IsBanner             bool      `json:"isBanner"`
								IsPremiere           bool      `json:"isPremiere"`
								Available            bool      `json:"available"`
								Title                string    `json:"title"`
								Description          string    `json:"description"`
								DescriptionFormatted string    `json:"descriptionFormatted"`
								Visibility           string    `json:"visibility"`
								BackgroundColor      string    `json:"backgroundColor"`
								TextColor            string    `json:"textColor"`
								Image                string    `json:"image"`
								OgImage              string    `json:"ogImage"`
								Created              time.Time `json:"created"`
								Modified             time.Time `json:"modified"`
								Owner                struct {
									UID      int    `json:"uid"`
									Login    string `json:"login"`
									Name     string `json:"name"`
									Verified bool   `json:"verified"`
								} `json:"owner"`
								Cover struct {
									Type    string `json:"type"`
									Dir     string `json:"dir"`
									Version string `json:"version"`
									URI     string `json:"uri"`
									Custom  bool   `json:"custom"`
								} `json:"cover"`
								Tags []struct {
									ID    string `json:"id"`
									Value string `json:"value"`
								} `json:"tags"`
							} `json:"playlist"`
							SomeArtists []struct {
								Various          bool     `json:"various"`
								Composer         bool     `json:"composer"`
								Available        bool     `json:"available"`
								TicketsAvailable bool     `json:"ticketsAvailable"`
								ID               string   `json:"id"`
								Name             string   `json:"name"`
								OgImage          string   `json:"ogImage"`
								Genres           []string `json:"genres"`
								Cover            struct {
									Type   string `json:"type"`
									Prefix string `json:"prefix"`
									URI    string `json:"uri"`
								} `json:"cover,omitempty"`
								Counts struct {
									Tracks       int `json:"tracks"`
									DirectAlbums int `json:"directAlbums"`
									AlsoAlbums   int `json:"alsoAlbums"`
									AlsoTracks   int `json:"alsoTracks"`
								} `json:"counts"`
								Ratings struct {
									Day   int `json:"day"`
									Week  int `json:"week"`
									Month int `json:"month"`
								} `json:"ratings,omitempty"`
								Links []struct {
									Title         string `json:"title"`
									Href          string `json:"href"`
									Type          string `json:"type"`
									SocialNetwork string `json:"socialNetwork,omitempty"`
								} `json:"links"`
							} `json:"someArtists"`
							ArtistsCount int `json:"artistsCount"`
						} `json:"playlists"`
					} `json:"promo,omitempty"`
					Message          string `json:"message,omitempty"`
					Genre            string `json:"genre,omitempty"`
					RadioIsAvailable bool   `json:"radioIsAvailable,omitempty"`
					Tracks           []struct {
						DurationMs               int    `json:"durationMs"`
						FileSize                 int    `json:"fileSize"`
						Available                bool   `json:"available"`
						AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`
						LyricsAvailable          bool   `json:"lyricsAvailable"`
						ID                       string `json:"id"`
						RealID                   string `json:"realId"`
						Title                    string `json:"title"`
						StorageDir               string `json:"storageDir"`
						CoverURI                 string `json:"coverUri"`
						OgImage                  string `json:"ogImage"`
						Major                    struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"major"`
						Normalization struct {
							Gain float64 `json:"gain"`
							Peak int     `json:"peak"`
						} `json:"normalization"`
						Artists []struct {
							ID       int           `json:"id"`
							Name     string        `json:"name"`
							Various  bool          `json:"various"`
							Composer bool          `json:"composer"`
							Genres   []interface{} `json:"genres"`
							Cover    struct {
								Type   string `json:"type"`
								Prefix string `json:"prefix"`
								URI    string `json:"uri"`
							} `json:"cover"`
						} `json:"artists"`
						Albums []struct {
							ID                       int           `json:"id"`
							Year                     int           `json:"year"`
							TrackCount               int           `json:"trackCount"`
							Recent                   bool          `json:"recent"`
							VeryImportant            bool          `json:"veryImportant"`
							Available                bool          `json:"available"`
							AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
							AvailableForMobile       bool          `json:"availableForMobile"`
							AvailablePartially       bool          `json:"availablePartially"`
							Title                    string        `json:"title"`
							CoverURI                 string        `json:"coverUri"`
							OgImage                  string        `json:"ogImage"`
							Genre                    string        `json:"genre"`
							Buy                      []interface{} `json:"buy"`
							Bests                    []interface{} `json:"bests"`
							Artists                  []struct {
								ID       int    `json:"id"`
								Name     string `json:"name"`
								Various  bool   `json:"various"`
								Composer bool   `json:"composer"`
								Cover    struct {
									Type   string `json:"type"`
									Prefix string `json:"prefix"`
									URI    string `json:"uri"`
								} `json:"cover"`
								Genres []interface{} `json:"genres"`
							} `json:"artists"`
							Labels []struct {
								ID          int    `json:"id"`
								Name        string `json:"name"`
								Description string `json:"description"`
								Image       string `json:"image"`
							} `json:"labels"`
							TrackPosition struct {
								Volume int `json:"volume"`
								Index  int `json:"index"`
							} `json:"trackPosition"`
						} `json:"albums"`
					} `json:"tracks,omitempty"`
				} `json:"events"`
				TracksToPlay []struct {
					DurationMs               int    `json:"durationMs"`
					FileSize                 int    `json:"fileSize"`
					Available                bool   `json:"available"`
					AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`
					Best                     bool   `json:"best,omitempty"`
					LyricsAvailable          bool   `json:"lyricsAvailable"`
					ID                       string `json:"id"`
					RealID                   string `json:"realId"`
					Title                    string `json:"title"`
					StorageDir               string `json:"storageDir"`
					CoverURI                 string `json:"coverUri"`
					OgImage                  string `json:"ogImage"`
					Version                  string `json:"version,omitempty"`
					ContentWarning           string `json:"contentWarning,omitempty"`
					Major                    struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"major"`
					Normalization struct {
						Gain float64 `json:"gain"`
						Peak int     `json:"peak"`
					} `json:"normalization"`
					Artists []struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Various  bool   `json:"various"`
						Composer bool   `json:"composer"`
						Cover    struct {
							Type   string `json:"type"`
							Prefix string `json:"prefix"`
							URI    string `json:"uri"`
						} `json:"cover"`
						Genres []interface{} `json:"genres"`
					} `json:"artists"`
					Albums []struct {
						ID                       int           `json:"id"`
						Year                     int           `json:"year"`
						TrackCount               int           `json:"trackCount"`
						Recent                   bool          `json:"recent"`
						VeryImportant            bool          `json:"veryImportant"`
						Available                bool          `json:"available"`
						AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
						AvailableForMobile       bool          `json:"availableForMobile"`
						AvailablePartially       bool          `json:"availablePartially"`
						Title                    string        `json:"title"`
						Type                     string        `json:"type"`
						CoverURI                 string        `json:"coverUri"`
						OgImage                  string        `json:"ogImage"`
						Genre                    string        `json:"genre"`
						Bests                    []interface{} `json:"bests"`
						Buy                      []interface{} `json:"buy"`
						ReleaseDate              time.Time     `json:"releaseDate"`
						Artists                  []struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Various  bool   `json:"various"`
							Composer bool   `json:"composer"`
							Cover    struct {
								Type   string `json:"type"`
								Prefix string `json:"prefix"`
								URI    string `json:"uri"`
							} `json:"cover"`
							Genres []interface{} `json:"genres"`
						} `json:"artists"`
						Labels []struct {
							ID          int    `json:"id"`
							Name        string `json:"name"`
							Description string `json:"description"`
							Image       string `json:"image"`
						} `json:"labels"`
						TrackPosition struct {
							Volume int `json:"volume"`
							Index  int `json:"index"`
						} `json:"trackPosition"`
					} `json:"albums"`
				} `json:"tracksToPlay"`
				TracksToPlayWithAds []struct {
					Type  string `json:"type"`
					Track struct {
						ID                       string `json:"id"`
						RealID                   string `json:"realId"`
						Title                    string `json:"title"`
						DurationMs               int    `json:"durationMs"`
						FileSize                 int    `json:"fileSize"`
						Available                bool   `json:"available"`
						AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`
						LyricsAvailable          bool   `json:"lyricsAvailable"`
						StorageDir               string `json:"storageDir"`
						CoverURI                 string `json:"coverUri"`
						OgImage                  string `json:"ogImage"`
						Major                    struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"major"`
						Normalization struct {
							Gain float64 `json:"gain"`
							Peak int     `json:"peak"`
						} `json:"normalization"`
						Artists []struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Various  bool   `json:"various"`
							Composer bool   `json:"composer"`
							Cover    struct {
								Type   string `json:"type"`
								Prefix string `json:"prefix"`
								URI    string `json:"uri"`
							} `json:"cover"`
							Genres []interface{} `json:"genres"`
						} `json:"artists"`
						Albums []struct {
							ID                       int           `json:"id"`
							Year                     int           `json:"year"`
							TrackCount               int           `json:"trackCount"`
							Recent                   bool          `json:"recent"`
							VeryImportant            bool          `json:"veryImportant"`
							Available                bool          `json:"available"`
							AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
							AvailableForMobile       bool          `json:"availableForMobile"`
							AvailablePartially       bool          `json:"availablePartially"`
							Title                    string        `json:"title"`
							Type                     string        `json:"type"`
							CoverURI                 string        `json:"coverUri"`
							OgImage                  string        `json:"ogImage"`
							Genre                    string        `json:"genre"`
							Bests                    []interface{} `json:"bests"`
							Buy                      []interface{} `json:"buy"`
							ReleaseDate              time.Time     `json:"releaseDate"`
							Artists                  []struct {
								ID       int    `json:"id"`
								Name     string `json:"name"`
								Various  bool   `json:"various"`
								Composer bool   `json:"composer"`
								Cover    struct {
									Type   string `json:"type"`
									Prefix string `json:"prefix"`
									URI    string `json:"uri"`
								} `json:"cover"`
								Genres []interface{} `json:"genres"`
							} `json:"artists"`
							Labels []struct {
								ID          int    `json:"id"`
								Name        string `json:"name"`
								Description string `json:"description"`
								Image       string `json:"image"`
							} `json:"labels"`
							TrackPosition struct {
								Volume int `json:"volume"`
								Index  int `json:"index"`
							} `json:"trackPosition"`
						} `json:"albums"`
					} `json:"track"`
				} `json:"tracksToPlayWithAds"`
			} `json:"days"`
		} `json:"result"`
	}
)

// Get returns feed of current user or base feed if there is no access token
func (s *FeedService) Get(
	ctx context.Context,
) (*FeedResp, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "feed", nil)
	if err != nil {
		return nil, nil, err
	}

	feed := new(FeedResp)
	resp, err := s.client.Do(ctx, req, feed)
	return feed, resp, err
}
