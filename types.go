package slack

// Representations of Slack's API Types
// https://api.slack.com/types
//
// Some of these aren't used right now, but might be in the future

// A Channel representation
type Type struct {
	ID string `json:"id"`
}

type Topic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

// Channel
// https://api.slack.com/types/channel
type Channel struct {
	Type

	Name    string `json:"name"`
	Created int    `json:"created"`
	Creator string `json:"creator"`

	Members []string `json:"members"`

	Topic   Topic `json:"topic"`
	Purpose Topic `json:"purpose"`

	IsChannel  bool `json:"is_channel"`
	IsGeneral  bool `json:"is_general"`
	IsArchived bool `json:"is_archived"`
	IsMember   bool `json:"is_member"`
}

// File
// https://api.slack.com/types/channel
type File struct {
	Type
	Created   int `json:"created"`
	Timestamp int `json:"timestamp"`

	Name       string `json:"name"`
	Title      string `json:"title"`
	User       string `json:"user"`
	Mimetype   string `json:"mimetype"`
	Filetype   string `json:"filetype"`
	PrettyType string `json:"pretty_type"`

	Mode         string `json:"mode"`
	Editable     bool   `json:"editable"`
	IsExternal   bool   `json:"is_external"`
	ExternalType string `json:"external_type"`

	Size int `json:"size"`

	URL                string `json:"url"`
	URLDownload        string `json:"url_download"`
	URLPrivate         string `json:"url_private"`
	URLPrivateDownload string `json:"url_private_download"`

	Thumb64     string `json:"thumb_64"`
	Thumb80     string `json:"thumb_80"`
	Thumb360    string `json:"thumb_360"`
	Thumb360Gif string `json:"thumb_360_gif"`
	Thumb360W   int    `json:"thumb_360_w"`
	Thumb360H   int    `json:"thumb_360_h"`

	Permalink        string `json:"permalink"`
	EditLink         string `json:"edit_link"`
	Preview          string `json:"preview"`
	PreviewHighlight string `json:"preview_highlight"`
	Lines            int    `json:"lines"`
	LinesMore        int    `json:"lines_more"`

	IsPublic        bool     `json:"is_public"`
	PublicURLShared bool     `json:"public_url_shared"`
	Channels        []string `json:"channels"`
	Groups          []string `json:"groups"`
	NumStars        int      `json:"num_stars"`
	IsStarred       bool     `json:"is_starred"`
}

// Group
// https://api.slack.com/types/group
type Group struct {
	Type

	Name    string `json:"name"`
	Created int    `json:"created"`
	Creator string `json:"creator"`

	Members []string `json:"members"`

	Topic   Topic `json:"topic"`
	Purpose Topic `json:"purpose"`

	IsGroup    bool `json:"is_group"`
	IsArchived bool `json:"is_archived"`
}

// IM
// https://api.slack.com/types/im
type IM struct {
	Type

	User    string `json:"user"`
	Created int    `json:"created"`

	IsIM          bool `json:"is_im"`
	IsUserDeleted bool `json:"is_user_deleted"`
}

// User
// https://api.slack.com/types/user
type User struct {
	Type

	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
	Color   string `json:"color"`

	Profile struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RealName  string `json:"real_name"`
		Email     string `json:"email"`
		Skype     string `json:"skype"`
		Phone     string `json:"phone"`
		Image24   string `json:"image_24"`
		Image32   string `json:"image_32"`
		Image48   string `json:"image_48"`
		Image72   string `json:"image_72"`
		Image192  string `json:"image_192"`
	} `json:"profile"`

	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`

	Has2FA   bool `json:"has_2fa"`
	HasFiles bool `json:"has_files"`
}
