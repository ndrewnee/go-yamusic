package yamusic

import (
	"context"
	"net/http"
)

type (
	// GenresService is a service to deal with genres.
	GenresService struct {
		client *Client
	}

	// Genres describes genres method response.
	Genres struct {
		InvocationInfo struct {
			Hostname           string `json:"hostname"`
			ReqID              string `json:"req-id"`
			ExecDurationMillis string `json:"exec-duration-millis"`
		} `json:"invocationInfo"`
		Result []struct {
			ID          string `json:"id"`
			Weight      int    `json:"weight"`
			ComposerTop bool   `json:"composerTop"`
			TracksCount int    `json:"tracksCount"`
			Title       string `json:"title"`
			FullTitle   string `json:"fullTitle,omitempty"`
			Titles      struct {
				Kk struct {
					Title string `json:"title"`
				} `json:"kk"`
				En struct {
					Title string `json:"title"`
				} `json:"en"`
				Be struct {
					Title     string `json:"title"`
					FullTitle string `json:"fullTitle"`
				} `json:"be"`
				Uk struct {
					Title string `json:"title"`
				} `json:"uk"`
				Ru struct {
					Title     string `json:"title"`
					FullTitle string `json:"fullTitle"`
				} `json:"ru"`
			} `json:"titles"`
			Images struct {
			} `json:"images"`
			ShowInMenu bool   `json:"showInMenu"`
			URLPart    string `json:"urlPart,omitempty"`
			Color      string `json:"color,omitempty"`
			RadioIcon  struct {
				BackgroundColor string `json:"backgroundColor"`
				ImageURL        string `json:"imageUrl"`
			} `json:"radioIcon,omitempty"`
			SubGenres []struct {
				ID          string `json:"id"`
				Weight      int    `json:"weight"`
				ComposerTop bool   `json:"composerTop"`
				URLPart     string `json:"urlPart"`
				TracksCount int    `json:"tracksCount"`
				Title       string `json:"title"`
				FullTitle   string `json:"fullTitle,omitempty"`
				Titles      struct {
					En struct {
						Title     string `json:"title"`
						FullTitle string `json:"fullTitle"`
					} `json:"en"`
					Be struct {
						Title     string `json:"title"`
						FullTitle string `json:"fullTitle"`
					} `json:"be"`
					Uk struct {
						Title     string `json:"title"`
						FullTitle string `json:"fullTitle"`
					} `json:"uk"`
					Ru struct {
						Title     string `json:"title"`
						FullTitle string `json:"fullTitle"`
					} `json:"ru"`
				} `json:"titles"`
				Color  string `json:"color"`
				Images struct {
					Two08X208   string `json:"208x208"`
					Three00X300 string `json:"300x300"`
				} `json:"images"`
				ShowInMenu bool `json:"showInMenu"`
				RadioIcon  struct {
					BackgroundColor string `json:"backgroundColor"`
					ImageURL        string `json:"imageUrl"`
				} `json:"radioIcon"`
			} `json:"subGenres,omitempty"`
		} `json:"result"`
	}
)

// List returns list of existed genres.
func (s *GenresService) List(ctx context.Context) (*Genres, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "genres", nil)
	if err != nil {
		return nil, nil, err
	}

	var genres = new(Genres)
	resp, err := s.client.Do(ctx, req, genres)
	if err != nil {
		return nil, resp, err
	}

	return genres, resp, err
}
