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

	// GenresListResp describes genres method response.
	GenresListResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         []struct {
			Weight      int    `json:"weight"`
			TracksCount int    `json:"tracksCount"`
			ComposerTop bool   `json:"composerTop"`
			ShowInMenu  bool   `json:"showInMenu"`
			ID          string `json:"id"`
			Title       string `json:"title"`
			FullTitle   string `json:"fullTitle,omitempty"`
			URLPart     string `json:"urlPart,omitempty"`
			Color       string `json:"color,omitempty"`
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
			RadioIcon struct {
				BackgroundColor string `json:"backgroundColor"`
				ImageURL        string `json:"imageUrl"`
			} `json:"radioIcon,omitempty"`
			SubGenres []struct {
				Weight      int    `json:"weight"`
				TracksCount int    `json:"tracksCount"`
				ComposerTop bool   `json:"composerTop"`
				ShowInMenu  bool   `json:"showInMenu"`
				URLPart     string `json:"urlPart"`
				ID          string `json:"id"`
				Title       string `json:"title"`
				FullTitle   string `json:"fullTitle,omitempty"`
				Color       string `json:"color"`
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
				Images struct {
					Two08X208   string `json:"208x208"`
					Three00X300 string `json:"300x300"`
				} `json:"images"`
				RadioIcon struct {
					BackgroundColor string `json:"backgroundColor"`
					ImageURL        string `json:"imageUrl"`
				} `json:"radioIcon"`
			} `json:"subGenres,omitempty"`
		} `json:"result"`
	}
)

// List returns list of existed genres.
func (s *GenresService) List(
	ctx context.Context,
) (*GenresListResp, *http.Response, error) {

	req, err := s.client.NewRequest(http.MethodGet, "genres", nil)
	if err != nil {
		return nil, nil, err
	}

	genres := new(GenresListResp)
	resp, err := s.client.Do(ctx, req, genres)
	return genres, resp, err
}
