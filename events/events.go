package events

import "github.com/stewart/slack/types"

// Structs that map to the responses laid out by Slack's RTM docs
// https://api.slack.com/rtm

type Event struct {
	Type string `json:"type"`
}

// Hello Event
// https://api.slack.com/events/hello
type HelloEvent struct{ Event }

// Message Event (subtypes currently not handled)
// https://api.slack.com/events/message
type MessageEvent struct {
	Event

	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`

	Subtype string `json:"subtype,omitempty"`
	Hidden  bool   `json:"hidden,omitempty"`

	DeletedTS string `json:"deleted_ts,omitempty"`
	EventTS   string `json:"event_ts,omitempty"`

	// an Edited message is a simple message with additional properties
	Edited struct {
		User string `json:"user"`
		TS   string `json:"ts"`
	} `json:"edited,omitempty"`
}

// User Typing Event
// https://api.slack.com/events/user_typing
type UserTypingEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// Channel Marked Event
// https://api.slack.com/events/channel_marked
type ChannelMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// Channel Created Event
// https://api.slack.com/events/channel_created
type ChannelCreatedEvent struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created string `json:"created"`
		Creator string `json:"creator"`
	} `json:"channel"`
}

// Channel Joined Event
// https://api.slack.com/events/channel_joined
type ChannelJoinedEvent struct {
	Event
	Channel types.Channel `json:"channel"`
}

// Channel Left Event
// https://api.slack.com/events/channel_left
type ChannelLeftEvent struct {
	Event
	Channel string `json:"channel"`
}

// Channel Deleted Event
// https://api.slack.com/events/channel_deleted
type ChannelDeletedEvent struct {
	Event
	Channel string `json:"channel"`
}

// Channel Rename Event
// https://api.slack.com/events/channel_rename
type ChannelRenameEvent struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
	}
}

// Channel Archive Event
// https://api.slack.com/events/channel_archive
type ChannelArchiveEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// Channel Unarchive Event
// https://api.slack.com/events/channel_unarchive
type ChannelUnarchiveEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// Channel History Changed Event
// https://api.slack.com/events/channel_history_changed
type ChannelHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// IM Created Event
// https://api.slack.com/events/im_created
type IMCreatedEvent struct {
	Event
	User    string        `json:"user"`
	Channel types.Channel `json:"channel"`
}

// IM Open Event
// https://api.slack.com/events/im_open
type IMOpenEvent struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// IM Close Event
// https://api.slack.com/events/im_close
type IMCloseEvent struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// IM Marked Event
// https://api.slack.com/events/im_marked
type IMMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// IM History Changed Event
// https://api.slack.com/events/im_history_changed
type IMHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// Group Joined Event
// https://api.slack.com/events/group_joined
type GroupJoinedEvent struct {
	Event
	Channel types.Channel `json:"channel"`
}

// Group Left Event
// https://api.slack.com/events/group_left
type GroupLeftEvent struct {
	Event
	Channel types.Channel `json:"channel"`
}

// Group Open Event
// https://api.slack.com/events/group_open
type GroupOpenEvent struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// Group Close Event
// https://api.slack.com/events/group_close
type GroupCloseEvent struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// Group Archive Event
// https://api.slack.com/events/group_archive
type GroupArchiveEvent struct {
	Event
	Channel string `json:"channel"`
}

// Group Unarchive Event
// https://api.slack.com/events/group_unarchive
type GroupUnarchiveEvent struct {
	Event
	Channel string `json:"channel"`
}

// Group Rename Event
// https://api.slack.com/events/group_rename
type GroupRenameEvent struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
	}
}

// Group Marked Event
// https://api.slack.com/events/group_marked
type GroupMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// Group History Changed Event
// https://api.slack.com/events/group_history_changed
type GroupHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// File Created Event
// https://api.slack.com/events/file_created
type FileCreatedEvent struct {
	Event
	File types.File `json:"file"`
}

// File Shared Event
// https://api.slack.com/events/file_shared
type FileSharedEvent struct {
	Event
	File types.File `json:"file"`
}

// File Unshared Event
// https://api.slack.com/events/file_unshared
type FileUnsharedEvent struct {
	Event
	File types.File `json:"file"`
}

// File Public Event
// https://api.slack.com/events/file_public
type FilePublicEvent struct {
	Event
	File types.File `json:"file"`
}

// File Private Event
// https://api.slack.com/events/file_private
type FilePrivateEvent struct {
	Event
	File string `json:"file"`
}

// File Change Event
// https://api.slack.com/events/file_change
type FileChangeEvent struct {
	Event
	File types.File `json:"file"`
}

// File Deleted Event
// https://api.slack.com/events/file_deleted
type FileDeletedEvent struct {
	Event
	FileID  string `json:"file_id"`
	EventTS string `json:"event_ts"`
}

// File Comment Added Event
// https://api.slack.com/events/file_comment_added
type FileCommentAddedEvent struct {
	Event
	File types.File `json:"file"`
}

// File Comment Edited Event
// https://api.slack.com/events/file_comment_edited
type FileCommentEditedEvent struct {
	Event
	File types.File `json:"file"`
}

// File Comment Deleted Event
// https://api.slack.com/events/file_comment_deleted
type FileCommentDeletedEvent struct {
	Event
	File    types.File `json:"file"`
	Comment string     `json:"comment"`
}

// Pin Added Event
// https://api.slack.com/events/pin_added
type PinAddedEvent struct{ Event }

// Pin Removed Event
// https://api.slack.com/events/pin_removed
type PinRemovedEvent struct{ Event }

// Presence Change Event
// https://api.slack.com/events/presence_change
type PresenceChangeEvent struct {
	Event
	User     string `json:"user"`
	Presence string `json:"presence"`
}

// Manual Presence Change Event
// https://api.slack.com/events/manual_presence_change
type ManualPresenceChangeEvent struct {
	Event
	Presence string `json:"presence"`
}

// Preference Change Event
// https://api.slack.com/events/pref_change
type PrefChangeEvent struct {
	Event
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// User Change Event
// https://api.slack.com/events/user_change
type UserChangeEvent struct {
	Event
	User types.User `json:"user"`
}

// Team Join Event
// https://api.slack.com/events/team_join
type TeamJoinEvent struct {
	Event
	User types.User `json:"user"`
}

// Star Added Event
// https://api.slack.com/events/star_added
type StarAddedEvent struct {
	Event
	User    string                 `json:"user"`
	Item    map[string]interface{} `json:"item"` // TODO: implement this
	EventTS string                 `json:"event_ts"`
}

// Star Removed Event
// https://api.slack.com/events/star_removed
type StarRemovedEvent struct {
	Event
	User    string                 `json:"user"`
	Item    map[string]interface{} `json:"item"` // TODO: implement this
	EventTS string                 `json:"event_ts"`
}

// Emoji Changed Event
// https://api.slack.com/events/emoji_changed
type EmojiChangedEvent struct {
	Event
	EventTS string `json:"event_ts"`
}

// Commands Changed Event
// https://api.slack.com/events/commands_changed
type CommandsChangedEvent struct {
	Event
	EventTS string `json:"event_ts"`
}

// Team Plan Change Event
// https://api.slack.com/events/team_plan_change
type TeamPlanChangeEvent struct {
	Event
	Plan string `json:"plan"`
}

// Team Pref Change Event
// https://api.slack.com/events/team_pref_change
type TeamPrefChangeEvent struct {
	Event
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// Team Rename Event
// https://api.slack.com/events/team_rename
type TeamRenameEvent struct {
	Event
	Name string `json:"name"`
}

// Team Domain Change Event
// https://api.slack.com/events/team_domain_change
type TeamDomainChangeEvent struct {
	Event
	URL    string `json:"url"`
	Domain string `json:"domain"`
}

// Email Domain Changed Event
// https://api.slack.com/events/email_domain_changed
type EmailDomainChangedEvent struct {
	Event
	EmailDomain string `json:"email_domain"`
	EventTS     string `json:"event_ts"`
}

// Bot Added Event
// https://api.slack.com/events/bot_added
type BotAddedEvent struct {
	Event
	Bot struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Icons struct {
			Image48 string `json:"image_48"`
		} `json:"icons"`
	} `json:"bot"`
}

// Bot Changed Event
// https://api.slack.com/events/bot_changed
type BotChangedEvent struct {
	Event
	Bot struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Icons struct {
			Image48 string `json:"image_48"`
		} `json:"icons"`
	} `json:"bot"`
}

// Accounts Changed Event
// https://api.slack.com/events/accounts_changed
type AccountsChangedEvent struct{ Event }

// Team Migration Started Event
// https://api.slack.com/events/team_migration_started
type TeamMigrationStartedEvent struct{ Event }
