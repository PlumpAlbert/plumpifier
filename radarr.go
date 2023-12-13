package main

type MovieDownloaded struct {
	Movie Movie     `json:"movie"`
	File  MovieFile `json:"movieFile"`
}

type Movie struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
	IMDB  string `json:"imdbId"`
}

type MovieFile struct {
	Id      int    `json:"id"`
	Quality string `json:"quality"`
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
