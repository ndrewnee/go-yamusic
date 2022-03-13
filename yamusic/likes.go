package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type (
	LikesService struct {
		client *Client
	}
	LikeResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Revision int `json:"revision"`
		} `json:"result"`
	}
)

func (l *LikesService) likeDislikeAction(
	ctx context.Context,
	mode string,
	toType string,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	if userID == 0 {
		userID = l.client.userID
	}
	action := "add-multiple"
	if remove {
		action = "remove"
	}
	uri := fmt.Sprintf(
		"%s/users/%d/%slikes/%ss/%s",
		l.client.baseURL, userID, mode, toType, action,
	)

	var strids string
	for i, id := range ids {
		strids += strconv.Itoa(id)
		if i < len(ids)-1 {
			strids += ","
		}
	}
	form := url.Values{}
	form.Set(toType+"-ids", strids)

	req, err := l.client.NewRequest(http.MethodPost, uri, form)

	lresp := new(LikeResp)
	resp, err := l.client.Do(ctx, req, &lresp)
	return lresp, resp, err
}

func (l *LikesService) Like(
	ctx context.Context,
	toType string,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.likeDislikeAction(ctx, "", toType, ids, remove, userID)
}

func (l *LikesService) Dislike(
	ctx context.Context,
	toType string,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.likeDislikeAction(ctx, "dis", toType, ids, remove, userID)
}

func (l *LikesService) LikeTrack(
	ctx context.Context,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.Like(ctx, "track", ids, remove, userID)
}

func (l *LikesService) DislikeTrack(
	ctx context.Context,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.Dislike(ctx, "track", ids, remove, userID)
}

func (l *LikesService) LikeAlbum(
	ctx context.Context,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.Like(ctx, "album", ids, remove, userID)
}

func (l *LikesService) LikePlaylist(
	ctx context.Context,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.Like(ctx, "playlist", ids, remove, userID)
}

func (l *LikesService) LikeUser(
	ctx context.Context,
	ids []int,
	remove bool,
	userID int,
) (*LikeResp, *http.Response, error) {
	return l.Like(ctx, "playlist", ids, remove, userID)
}

func (l *LikesService) getLikesDislikesAction(
	ctx context.Context,
	mode string,
	fromType string,
	to interface{},
	userID int,
) (*http.Response, error) {
	uri := fmt.Sprintf(
		"%s/users/%d/%slikes/%ss",
		l.client.baseURL, userID, mode, fromType,
	)
	req, err := l.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req, to)
	return resp, err
}

func (l *LikesService) GetLikes(
	ctx context.Context,
	fromType string,
	to interface{},
	userID int,
) (*http.Response, error) {
	return l.getLikesDislikesAction(ctx, "", fromType, to, userID)
}

func (l *LikesService) GetDisikes(
	ctx context.Context,
	fromType string,
	to interface{},
	userID int,
) (*http.Response, error) {
	return l.getLikesDislikesAction(ctx, "dis", fromType, to, userID)
}
