package yinyuetai

type yinyuetaiMvData struct {
	Error     bool      `json:"error"`
	Message   string    `json:"message"`
	VideoInfo videoInfo `json:"videoInfo"`
}

type videoInfo struct {
	CoreVideoInfo coreVideoInfo `json:"coreVideoInfo"`
}

type coreVideoInfo struct {
	ArtistNames    string          `json:"artistNames"`
	Duration       int             `json:"duration"`
	Error          bool            `json:"error"`
	ErrorMsg       string          `json:"errorMsg"`
	VideoId        int             `json:"videoId"`
	VideoName      string          `json:"videoName"`
	VideoUrlModels []videoUrlModel `json:"videoUrlModels"`
}

type videoUrlModel struct {
	Bitrate          int    `json:"bitrate"`
	BitrateType      int    `json:"bitrateType"`
	FileSize         int64  `json:"fileSize"`
	MD5              string `json:"md5"`
	SHA1             string `json:"sha1"`
	QualityLevel     string `json:"qualityLevel"`
	QualityLevelName string `json:"qualityLevelName"`
	VideoUrl         string `json:"videoUrl"`
}
