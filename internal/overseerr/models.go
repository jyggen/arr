package overseerr

import "time"

const (
	MediaRequestStatusPending = iota + 1
	MediaRequestStatusApproved
	MediaRequestStatusDeclined
	MediaRequestStatusFailed
)

const (
	MediaStatusUnknown = iota + 1
	MediaStatusPending
	MediaStatusProcessing
	MediaStatusPartiallyAvailable
	MediaStatusAvailable
)

const (
	MediaTypeMovie = "movie"
	MediaTypeTv    = "tv"
)

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Pages    int `json:"pages"`
	Results  int `json:"results"`
}

type Media struct {
	Id                    int       `json:"id"`
	MediaType             string    `json:"mediaType"`
	TmdbId                int       `json:"tmdbId"`
	TvdbId                int       `json:"tvdbId"`
	ImdbId                string    `json:"imdbId"`
	Status                int       `json:"status"`
	Status4K              int       `json:"status4k"`
	ExternalServiceId     int       `json:"externalServiceId"`
	ExternalServiceId4K   int       `json:"externalServiceId4k"`
	ExternalServiceSlug   string    `json:"externalServiceSlug"`
	ExternalServiceSlug4K string    `json:"externalServiceSlug4k"`
	RatingKey             string    `json:"ratingKey"`
	RatingKey4K           string    `json:"RatingKey4k"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	LastSeasonChange      time.Time `json:"lastSeasonChange"`
	MediaAddedAt          time.Time `json:"mediaAddedAt"`
}

type Season struct {
	Id           int       `json:"id"`
	SeasonNumber int       `json:"seasonNumber"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type RequestedBy struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PlexToken    string    `json:"plexToken"`
	PlexUsername string    `json:"plexUsername"`
	UserType     int       `json:"userType"`
	Permissions  int       `json:"permissions"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	RequestCount int       `json:"requestCount"`
}

type ModifiedBy struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PlexToken    string    `json:"plexToken"`
	PlexUsername string    `json:"plexUsername"`
	UserType     int       `json:"userType"`
	Permissions  int       `json:"permissions"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	RequestCount int       `json:"requestCount"`
}

type Result struct {
	Id                int         `json:"id"`
	Status            int         `json:"status"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
	Type              string      `json:"type"`
	Is4K              bool        `json:"is4k"`
	ServerId          int         `json:"serverId"`
	ProfileId         int         `json:"profileId"`
	RootFolder        string      `json:"rootFolder"`
	LanguageProfileId int         `json:"languageProfileId"`
	Tags              []int       `json:"tags"`
	IsAutoRequest     bool        `json:"isAutoRequest"`
	Media             Media       `json:"media"`
	RequestedBy       RequestedBy `json:"requestedBy"`
	ModifiedBy        ModifiedBy  `json:"modifiedBy"`
	SeasonCount       int         `json:"seasonCount"`
}

type RequestsResponse struct {
	PageInfo PageInfo `json:"pageInfo"`
	Results  []Result `json:"results"`
}
