package gohook

// NewEmbed creates a new embed
func NewEmbed() *Embed {
	return &Embed{}
}

// AddField adds a field to the embed
func (embed *Embed) AddField(Name string, Value string, Inline bool) *Embed {
	embed.Fields = append(embed.Fields, Field{
		Name:   Name,
		Value:  Value,
		Inline: Inline,
	})
	return embed
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Url         string  `json:"url,omitempty"`
	Color       int     `json:"color,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	Fields      []Field `json:"fields,omitempty"`

	Author    Author    `json:"author,omitempty"`
	Footer    Footer    `json:"footer,omitempty"`
	Image     Image     `json:"image,omitempty"`
	Thumbnail Thumbnail `json:"thumbnail,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Author struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Image struct {
	Url string `json:"url,omitempty"`
}

type Thumbnail struct {
	Url string `json:"url,omitempty"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}
