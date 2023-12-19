package lib

import "time"

type MovieDownloaded struct {
	Movie              Movie       `json:"movie"`
	File               MovieFile   `json:"movieFile"`
	RemoteMovie        RemoteMovie `json:"remoteMovie"`
	IsUpgrade          bool        `json:"isUpgrade"`
	DownloadClient     string      `json:"downloadClient"`
	DownloadClientType string      `json:"downloadClientType"`
	DownloadId         string      `json:"downloadId"`
	EventType          string      `json:"eventType"`
}

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"releaseDate"`
	FolderPath  string    `json:"folderPath"`
	TMDB        int       `json:"tmdbId"`
	IMDB        string    `json:"imdbId"`
}

type RemoteMovie struct {
	TMDB  int    `json:"tmdbId"`
	IMDB  string `json:"imdbId"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type MovieFile struct {
	Id             int    `json:"id"`
	RelativePath   string `json:"relativePath"`
	Path           string `json:"path"`
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	ReleaseGroup   string `json:"releaseGroup"`
	SceneName      string `json:"sceneName"`
	IndexerFlags   string `json:"indexerFlags"`
	Size           int    `json:"size"`
}

// {
//   "movie": {
//     "id": 686,
//     "title": "<movieTitle>",
//     "year": <int>,
//     "releaseDate": "yyyy-MM-dd",
//     "folderPath": "<pathToFolder>",
//     "tmdbId": <int>,
//     "imdbId": "tt<id>"
//   },
//   "remoteMovie": {
//     "tmdbId": <int>,
//     "imdbId": "tt<id>",
//     "title": "<movieTitle>",
//     "year": <int>
//   },
//   "movieFile": {
//     "id": 36745,
//     "relativePath": "<filenameAfterRename",
//     "path": "<pathToDownloadedFile>",
//     "quality": "WEBDL-1080p",
//     "qualityVersion": 1,
//     "releaseGroup": "<rlsGroup>",
//     "sceneName": "The.Movie.Title.2012.1080p.BluRay.x265-<rlsGroup>",
//     "indexerFlags": "0",
//     "size": <intBytes>
//   },
//   "isUpgrade": false,
//   "downloadClient": "<downloadClient>",
//   "downloadClientType": "<type>",
//   "downloadId": "<downloadClient>_The.Movie.Title.2012.1080p.BluRay.x265-<rlsGroup>_<longInt>",
//   "eventType": "Download"
// }
