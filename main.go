package main

import (
	"log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/twitter-srv/handler"
	proto "github.com/micro/twitter-srv/proto/api"
	"github.com/micro/twitter-srv/twitter"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.twitter"),
		micro.Flags(
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
		),
		micro.Action(func(c *cli.Context) {
			twitter.Token = c.String("access_token")
			twitter.TokenSecret = c.String("access_token_secret")
			twitter.ConsumerKey = c.String("consumer_key")
			twitter.ConsumerSecret = c.String("consumer_secret")
		}),
	)

	service.Init()
	twitter.Init()

	proto.RegisterApiHandler(service.Server(), &handler.Api{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
