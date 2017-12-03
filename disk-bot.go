package main

import "fmt"

type DiskBot struct {
	TwitterClient *TwitterClient
	Conf          *Configuration
}

func (d *DiskBot) Start() {
	fmt.Println("Starting diskbot")
	// TODO: Search hashtags for #globe
	// TODO: Pick one at random to reply to
	// TODO: Create tweet body with image
	// TODO: Send tweet
	resp := d.TwitterClient.GetTweetsForHashtag("globe")
	fmt.Println(resp)
}

func NewDiskBot(conf *Configuration) *DiskBot {
	d := new(DiskBot)
	d.Conf = conf
	d.TwitterClient = NewTwitterClient(d.Conf.TwitterConsumerKey,
		d.Conf.TwitterConsumerSecret)
	return d
}
