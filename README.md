# youtube-cli
A Youtube CLI written in Go for using the Youtube API

# usage
To use the YouTube-API you need an api key. As long as you don't exceed your daily API request quota of 10,000 units per day, it's free. Required steps to register an api key can be found at [YouTube Data API](https://developers.google.com/youtube/v3/getting-started?hl=de)

# run
export API-Key=<your-api-key>

# listChannel
List channels for user YouTube
```bash
go run main.go yt listChannels YouTube
```