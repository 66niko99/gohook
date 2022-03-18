package gohook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// NewWebhook creates a webhook with the optional "webhookUrl" parameter
func NewWebhook(webhookUrl ...string) *Webhook {
	if len(webhookUrl) > 0 {
		return &Webhook{WebhookUrl: webhookUrl[0]}
	}
	return &Webhook{}
}

// AddEmbed adds an embed to the webhook
func (wh *Webhook) AddEmbed(embed *Embed) *Webhook {
	wh.Embeds = append(wh.Embeds, embed)
	return wh
}

// NewEmbed creates an embed and adds it to the webhook
func (wh *Webhook) NewEmbed() *Embed {
	embed := &Embed{}
	wh.Embeds = append(wh.Embeds, embed)
	return embed
}

// Json returns the webhooks as json as bytes
func (wh *Webhook) Json() []byte {
	asJson, err := json.Marshal(wh)
	if err != nil {
		log.Println("webhook json marshal error: ", err)
		return nil
	}
	return asJson
}

// Send sends the webhook to the specified url
func (wh *Webhook) Send() (*http.Response, error) {
	return http.Post(wh.WebhookUrl, "application/json", bytes.NewBuffer(wh.Json()))
}

type Webhook struct {
	WebhookUrl string   `json:"-"`
	Username   string   `json:"username,omitempty"`
	AvatarUrl  string   `json:"avatar_url,omitempty"`
	Content    string   `json:"content,omitempty"`
	Embeds     []*Embed `json:"embeds,omitempty"`
}
