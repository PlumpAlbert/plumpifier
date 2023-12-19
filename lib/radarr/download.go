package radarr

type MovieDownloaded struct {
	WebhookPayload
	Movie              Movie            `json:"movie"`
	RemoteMovie        RemoteMovie      `json:"remoteMovie"`
	MovieFile          MovieFile        `json:"movieFile"`
	Release            Release          `json:"release"`
	IsUpgrade          bool             `json:"isUpgrade"`
	DownloadClient     string           `json:"downloadClient"`
	DownloadClientType string           `json:"downloadClientType"`
	DownloadId         string           `json:"downloadId"`
	CustomFormatInfo   CustomFormatInfo `json:"customFormatInfo"`
	DeletedFiles       []MovieFile      `json:"deletedFiles"`
}
