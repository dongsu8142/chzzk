package chzzk

type ChannelsResponse struct {
	Code    int     `json:"code"`
	Content ChannelContent `json:"content"`
}

type ChannelContent struct {
	Data []ChannelData `json:"data"`
}

type ChannelData struct {
	Channel `json:"channel"`
}

type Channel struct {
	ChannelID          string       `json:"channelId"`
	ChannelName        string       `json:"channelName"`
	ChannelImageURL    string       `json:"channelImageUrl"`
	VerifiedMark       bool         `json:"verifiedMark"`
	ChannelDescription string       `json:"channelDescription"`
	FollowerCount      int          `json:"followerCount"`
	OpenLive           bool         `json:"openLive"`
	PersonalData       ChannelPersonalData `json:"personalData"`
}

type ChannelPersonalData struct {
	Following        Following `json:"following"`
	PrivateUserBlock bool      `json:"privateUserBlock"`
}

type Following struct {
	Following    bool   `json:"following"`
	Notification bool   `json:"notification"`
	FollowDate   string `json:"followDate"`
}

type LivesResponse struct {
	Code    int     `json:"code"`
	Content LivesContent `json:"content"`
}

type LivesContent struct {
	Data []LivesData `json:"data"`
}

type LivesData struct {
	Live    Live    `json:"live"`
	LivesChannel LivesChannel `json:"channel"`
}

type Live struct {
	LiveTitle                string   `json:"liveTitle"`
	LiveImageURL             string   `json:"liveImageUrl"`
	DefaultThumbnailImageURL string   `json:"defaultThumbnailImageUrl"`
	ConcurrentUserCount      int      `json:"concurrentUserCount"`
	AccumulateCount          int      `json:"accumulateCount"`
	OpenDate                 string   `json:"openDate"`
	LiveID                   int      `json:"liveId"`
	Adult                    bool     `json:"adult"`
	Tags                     []string `json:"tags"`
	ChatChannelID            string   `json:"chatChannelId"`
	CategoryType             string   `json:"categoryType"`
	LiveCategory             string   `json:"liveCategory"`
	LiveCategoryValue        string   `json:"liveCategoryValue"`
	ChannelID                string   `json:"channelId"`
	LivePlaybackJSON         string   `json:"livePlaybackJson"`
	BlindType                any      `json:"blindType"`
}

type PersonalData struct {
	PrivateUserBlock bool `json:"privateUserBlock"`
}

type LivesChannel struct {
	ChannelID       string       `json:"channelId"`
	ChannelName     string       `json:"channelName"`
	ChannelImageURL string       `json:"channelImageUrl"`
	VerifiedMark    bool         `json:"verifiedMark"`
	PersonalData    PersonalData `json:"personalData"`
}

type VideosResponse struct {
	Code    int     `json:"code"`
	Content VideosContent `json:"content"`
}

type Video struct {
	VideoNo            int    `json:"videoNo"`
	VideoID            string `json:"videoId"`
	VideoTitle         string `json:"videoTitle"`
	VideoType          string `json:"videoType"`
	PublishDate        string `json:"publishDate"`
	ThumbnailImageURL  string `json:"thumbnailImageUrl"`
	Duration           int    `json:"duration"`
	ReadCount          int    `json:"readCount"`
	ChannelID          string `json:"channelId"`
	PublishDateAt      int64  `json:"publishDateAt"`
	Adult              bool   `json:"adult"`
	CategoryType       string `json:"categoryType"`
	VideoCategory      string `json:"videoCategory"`
	VideoCategoryValue string `json:"videoCategoryValue"`
	BlindType          any    `json:"blindType"`
}

type VideosChannel struct {
	ChannelID       string       `json:"channelId"`
	ChannelName     string       `json:"channelName"`
	ChannelImageURL string       `json:"channelImageUrl"`
	VerifiedMark    bool         `json:"verifiedMark"`
	PersonalData    PersonalData `json:"personalData"`
}

type VideosData struct {
	Video   Video   `json:"video"`
	Channel VideosChannel `json:"channel"`
}

type VideosContent struct {
	Data []VideosData `json:"data"`
}

type ChzzkClient struct {
	NidAuth    string
	NidSession string
}