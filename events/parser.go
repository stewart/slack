package events

import "encoding/json"

func Parse(msg []byte) (interface{}, error) {
	var event = Event{}

	if err := json.Unmarshal(msg, &event); err != nil {
		return nil, err
	}

	switch event.Type {
	case "accounts_changed":
		var event = AccountsChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "bot_added":
		var event = BotAdded{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "bot_changed":
		var event = BotChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_archive":
		var event = ChannelArchive{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_created":
		var event = ChannelCreated{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_deleted":
		var event = ChannelDeleted{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_history_changed":
		var event = ChannelHistoryChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_joined":
		var event = ChannelJoined{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_left":
		var event = ChannelLeft{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_marked":
		var event = ChannelMarked{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_rename":
		var event = ChannelRename{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_unarchive":
		var event = ChannelUnarchive{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "commands_changed":
		var event = CommandsChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "email_domain_changed":
		var event = EmailDomainChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "emoji_changed":
		var event = EmojiChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_change":
		var event = FileChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_added":
		var event = FileCommentAdded{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_deleted":
		var event = FileCommentDeleted{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_edited":
		var event = FileCommentEdited{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_created":
		var event = FileCreated{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_deleted":
		var event = FileDeleted{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_private":
		var event = FilePrivate{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_public":
		var event = FilePublic{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_shared":
		var event = FileShared{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_unshared":
		var event = FileUnshared{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_archive":
		var event = GroupArchive{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_close":
		var event = GroupClose{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_history_changed":
		var event = GroupHistoryChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_joined":
		var event = GroupJoined{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_left":
		var event = GroupLeft{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_marked":
		var event = GroupMarked{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_open":
		var event = GroupOpen{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_rename":
		var event = GroupRename{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_unarchive":
		var event = GroupUnarchive{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "hello":
		var event = Hello{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_close":
		var event = IMClose{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_created":
		var event = IMCreated{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_history_changed":
		var event = IMHistoryChanged{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_marked":
		var event = IMMarked{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_open":
		var event = IMOpen{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "manual_presence_change":
		var event = ManualPresenceChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "message":
		var event = Message{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pin_added":
		var event = PinAdded{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pin_removed":
		var event = PinRemoved{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pref_change":
		var event = PrefChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "presence_change":
		var event = PresenceChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "star_added":
		var event = StarAdded{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "star_removed":
		var event = StarRemoved{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_domain_change":
		var event = TeamDomainChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_join":
		var event = TeamJoin{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_migration_started":
		var event = TeamMigrationStarted{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_plan_change":
		var event = TeamPlanChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_pref_change":
		var event = TeamPrefChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_rename":
		var event = TeamRename{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "user_change":
		var event = UserChange{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "user_typing":
		var event = UserTyping{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		var event = Event{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	}
}
