package models

type URL struct {
	Key      string `json:"key"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type UrlArgs struct {
	URL string `json:"url" binding:"required"`
}
