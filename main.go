package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/micro/twitter-srv/handler"
	proto "github.com/micro/twitter-srv/proto/api"
	"github.com/micro/twitter-srv/twitter"
)

func main() {
	cmd.Flags = append(cmd.Flags,
		cli.StringFlag{
			Name:   "access_token",
			EnvVar: "ACCESS_TOKEN",
			Usage:  "Twitter access token",
		},
		cli.StringFlag{
			Name:   "access_token_secret",
			EnvVar: "ACCESS_TOKEN_SECRET",
			Usage:  "Twitter access token secret",
		},
		cli.StringFlag{
			Name:   "consumer_key",
			EnvVar: "CONSUMER_KEY",
			Usage:  "Twitter consumer key",
		},
		cli.StringFlag{
			Name:   "consumer_secret",
			EnvVar: "ACCESS_TOKEN",
			Usage:  "Twitter consumer secret",
		},
	)

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		twitter.Token = c.String("access_token")
		twitter.TokenSecret = c.String("access_token_secret")
		twitter.ConsumerKey = c.String("consumer_key")
		twitter.ConsumerSecret = c.String("consumer_secret")
	})

	cmd.Init()

	twitter.Init()

	server.Init(
		server.Name("go.micro.srv.twitter"),
	)

	proto.RegisterApiHandler(server.DefaultServer, &handler.Api{})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
