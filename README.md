# Twitter Server 

Twitter server is lets you tweet, it's as simple as that.

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Get your API keys from the twitter developer console.

4. Download and start the service

	```shell
	go get github.com/micro/twitter-srv
	twitter-srv --access_token=XXX --access_token-secret=XXX --consumer_key=XXX --consumer_secret=XXX
	```

	OR as a docker container

	```shell
	docker run microhq/twitter-srv --access_token=XXX --access_token-secret=XXX --consumer_key=XXX --consumer_secret=XXX --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API

```shell
micro query go.micro.srv.twitter Api.Tweet '{"status": "Tweeting via a microservice"}'
{
	"created_at": "Fri Dec 25 01:20:51 +0000 2015",
	"entities": {
		"url": {}
	},
	"extended_entities": {
		"url": {}
	},
	"id": 6.801964754737848e+17,
	"id_str": "680196475473784833",
	"lang": "en",
	"place": {
		"bounding_box": {},
		"geometry": {}
	},
	"source": "\u003ca href=\"http://asl.am\" rel=\"nofollow\"\u003easl.am\u003c/a\u003e",
	"text": "Tweeting via a microservice",
	"user": {
		"created_at": "Wed Sep 23 18:10:11 +0000 2015",
		"default_profile": true,
		"description": "Tech and world news",
		"entities": {
			"url": {}
		},
		"followers_count": 26,
		"id": 3.751898902e+09,
		"id_str": "3751898902",
		"lang": "en",
		"listed_count": 6,
		"location": "Everywhere",
		"name": "asl.am",
		"profile_background_color": "C0DEED",
		"profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
		"profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
		"profile_image_url": "http://pbs.twimg.com/profile_images/646768994062610436/VpCJnZVO_normal.jpg",
		"profile_image_url_https": "https://pbs.twimg.com/profile_images/646768994062610436/VpCJnZVO_normal.jpg",
		"profile_link_color": "0084B4",
		"profile_sidebar_border_color": "C0DEED",
		"profile_sidebar_fill_color": "DDEEF6",
		"profile_text_color": "333333",
		"profile_use_background_image": true,
		"screen_name": "_asl_am",
		"statuses_count": 11855,
		"url": "http://t.co/OMhazIHzZk"
	}
}
```
