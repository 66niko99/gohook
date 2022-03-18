# gohook
Go package to help you send Discord webhooks

## Usage: 
```golang
// create the webhook
webhook := gohook.NewWebhook("") // your webhook url here, or alternatively set it with webhook.WebhookUrl
// customize the webhook
webhook.Username = "Gohook"
webhook.AvatarUrl = "https://avatars.githubusercontent.com/u/69372253?v=4"
webhook.Content = "This is a test"

// create the embed
embed := webhook.NewEmbed()
// customize the embed
embed.Title = "Test Embed"
embed.Footer.Text = "Test footer"
embed.Footer.IconUrl = "https://avatars.githubusercontent.com/u/69372253?v=4"
embed.Image.Url = "https://avatars.githubusercontent.com/u/69372253?v=4"
embed.Url = "https://github.com/66niko99/gohook"
embed.Color = 15417396
embed.Thumbnail.Url = "https://avatars.githubusercontent.com/u/69372253?v=4"
// set author
embed.Author.Name = "66niko999"
embed.Author.IconUrl = "https://avatars.githubusercontent.com/u/69372253?v=4"
embed.Author.Url = "https://github.com/66niko99"

// add fields
embed.AddField("Test Field", "Test Value", false).
  AddField("Test Field", "Test Value", true).
  AddField("Test Field", "Test Value", true)

// send webhook
_, err := webhook.Send()
if err != nil {
  log.Fatalln(err)
}
```
## Result:
![image](https://user-images.githubusercontent.com/69372253/159091371-110a29d1-7e62-4b49-9455-a469b98bd6cf.png)
