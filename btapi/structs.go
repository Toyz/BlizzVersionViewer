package btapi

// All Games
type Game struct {
	Name         string    `json:"name"`
	ImageCode    string    `json:"image_code"`
	URL          string    `json:"url"`
	BlogCode     string    `json:"blog_code,omitempty"`
	HasNewForums bool      `json:"has_new_forums"`
	Channels     []Channel `json:"channels"`
}

type Channel struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	GameType  string `json:"game_type"`
	NotesCode string `json:"notes_code,omitempty"`
	Code      string `json:"code"`
}

//Region Data
type RegionInfo struct {
	Buildconfig   string `json:"buildconfig"`
	Buildid       string `json:"buildid"`
	Cdnconfig     string `json:"cdnconfig"`
	Keyring       string `json:"keyring"`
	Region        string `json:"region"`
	Regionname    string `json:"regionname"`
	Versionsname  string `json:"versionsname"`
	Productconfig string `json:"productconfig"`
	Updated       string `json:"updated"`
}

//Patch Notes
type PatchNotes struct {
	PatchNotes []PatchNote `json:"patchNotes"`
	Pagination Pagination  `json:"pagination"`
}

type PatchNote struct {
	Program      string `json:"program"`
	Locale       string `json:"locale"`
	Type         string `json:"type"`
	PatchVersion string `json:"patchVersion"`
	Status       string `json:"status"`
	Detail       string `json:"detail"`
	BuildNumber  int    `json:"buildNumber"`
	Publish      int64  `json:"publish"`
	Created      int64  `json:"created"`
	Updated      int64  `json:"updated"`
	Develop      bool   `json:"develop"`
	Slug         string `json:"slug"`
	Version      string `json:"version"`
}

type Pagination struct {
	TotalEntries int `json:"totalEntries"`
	TotalPages   int `json:"totalPages"`
	PageSize     int `json:"pageSize"`
	Page         int `json:"page"`
}
