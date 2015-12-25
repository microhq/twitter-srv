package handler

import (
	"fmt"
	"net/url"

	"github.com/micro/go-micro/errors"
	api "github.com/micro/twitter-srv/proto/api"
	"github.com/micro/twitter-srv/twitter"
	"golang.org/x/net/context"
)

type Api struct{}

func (a *Api) Tweet(ctx context.Context, req *api.TweetRequest, rsp *api.TweetResponse) error {
	if len(req.Status) == 0 {
		return errors.BadRequest("go.micro.srv.twitter.Api.Tweet", "Status cannot be blank")
	}
	if len(req.Status) > 140 {
		return errors.BadRequest("go.micro.srv.twitter.Api.Tweet", "Status cannot be longer than 140 chars")
	}

	u := url.Values{}

	if req.InReplyToStatusId > 0 {
		u.Set("in_reply_to_status_id", fmt.Sprintf("%d", req.InReplyToStatusId))
	}

	if req.PossiblySensitive {
		u.Set("possibly_sensitive", "true")
	}

	if req.LatLng != nil {
		u.Set("lat", fmt.Sprintf("%.6f", req.LatLng.Lat))
		u.Set("long", fmt.Sprintf("%.6f", req.LatLng.Lng))
	}

	if len(req.PlaceId) > 0 {
		u.Set("place_id", req.PlaceId)
	}

	if req.DisplayCoordinates {
		u.Set("display_coordinates", "true")
	}

	if req.TrimUser {
		u.Set("trim_user", "true")
	}

	for _, id := range req.MediaIds {
		u.Set("media_ids", fmt.Sprintf("%d", id))
	}

	if err := twitter.Tweet(req.Status, u, rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.twitter.Api.Tweet", err.Error())
	}

	return nil
}
