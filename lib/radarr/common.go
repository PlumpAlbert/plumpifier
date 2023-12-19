package radarr

import "time"

type WebhookPayload struct {
	EventType      string `json:"eventType"`
	InstanceName   string `json:"instanceName"`
	ApplicationURL string `json:"applicationUrl"`
}

type CustomFormatInfo struct {
	CustomFormats     []CustomFormat `json:"customFormats"`
	CustomFormatScore int            `json:"customFormatScore"`
}

type CustomFormat struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Release struct {
	Quality           string   `json:"quality"`
	QualityVersion    int      `json:"qualityVersion"`
	ReleaseGroup      string   `json:"releaseGroup"`
	ReleaseTitle      string   `json:"releaseTitle"`
	Indexer           string   `json:"indexer"`
	Size              int64    `json:"size"`
	CustomFormatScore int      `json:"customFormatScore"`
	CustomFormats     []string `json:"customFormats"`
	IndexerFlags      []string `json:"indexerFlags"`
}

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	FilePath    string    `json:"filePath"`
	ReleaseDate time.Time `json:"releaseDate"`
	FolderPath  string    `json:"folderPath"`
	TMDB        int       `json:"tmdbId"`
	IMDB        string    `json:"imdbId"`
	Overview    string    `json:"overview"`
}

type RemoteMovie struct {
	TMDB  int    `json:"tmdbId"`
	IMDB  string `json:"imdbId"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type MovieFile struct {
	Id             int           `json:"id"`
	RelativePath   string        `json:"relativePath"`
	Path           string        `json:"path"`
	Quality        string        `json:"quality"`
	QualityVersion int           `json:"qualityVersion"`
	ReleaseGroup   string        `json:"releaseGroup"`
	SceneName      string        `json:"sceneName"`
	IndexerFlags   string        `json:"indexerFlags"`
	Size           int64         `json:"size"`
	DateAdded      time.Time     `json:"dateAdded"`
	MediaInfo      FileMediaInfo `json:"mediaInfo"`
}

type FileMediaInfo struct {
	AudioChannels         float32  `json:"audioChannels"`
	AudioCodec            string   `json:"audioCodec"`
	AudioLanguages        []string `json:"audioLanguages"`
	Height                int      `json:"height"`
	Width                 int      `json:"width"`
	Subtitles             []string `json:"subtitles"`
	VideoCodec            string   `json:"videoCodec"`
	VideoDynamicRange     string   `json:"videoDynamicRange"`
	VideoDynamicRangeType string   `json:"videoDynamicRangeType"`
}
