package events

import "encoding/json"

func Parse(msg []byte) (interface{}, error) {
	var event = Event{}

	if err := json.Unmarshal(msg, &event); err != nil {
		return nil, err
	}

	switch event.Type {
	case "accounts_changed":
		var event = AccountsChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "bot_added":
		var event = BotAddedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "bot_changed":
		var event = BotChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_archive":
		var event = ChannelArchiveEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_created":
		var event = ChannelCreatedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_deleted":
		var event = ChannelDeletedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_history_changed":
		var event = ChannelHistoryChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_joined":
		var event = ChannelJoinedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_left":
		var event = ChannelLeftEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_marked":
		var event = ChannelMarkedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_rename":
		var event = ChannelRenameEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "channel_unarchive":
		var event = ChannelUnarchiveEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "commands_changed":
		var event = CommandsChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "email_domain_changed":
		var event = EmailDomainChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "emoji_changed":
		var event = EmojiChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_change":
		var event = FileChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_added":
		var event = FileCommentAddedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_deleted":
		var event = FileCommentDeletedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_comment_edited":
		var event = FileCommentEditedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_created":
		var event = FileCreatedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_deleted":
		var event = FileDeletedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_private":
		var event = FilePrivateEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_public":
		var event = FilePublicEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_shared":
		var event = FileSharedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "file_unshared":
		var event = FileUnsharedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_archive":
		var event = GroupArchiveEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_close":
		var event = GroupCloseEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_history_changed":
		var event = GroupHistoryChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_joined":
		var event = GroupJoinedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_left":
		var event = GroupLeftEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_marked":
		var event = GroupMarkedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_open":
		var event = GroupOpenEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_rename":
		var event = GroupRenameEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "group_unarchive":
		var event = GroupUnarchiveEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "hello":
		var event = HelloEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_close":
		var event = IMCloseEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_created":
		var event = IMCreatedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_history_changed":
		var event = IMHistoryChangedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_marked":
		var event = IMMarkedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "im_open":
		var event = IMOpenEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "manual_presence_change":
		var event = ManualPresenceChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "message":
		var event = MessageEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pin_added":
		var event = PinAddedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pin_removed":
		var event = PinRemovedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "pref_change":
		var event = PrefChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "presence_change":
		var event = PresenceChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "star_added":
		var event = StarAddedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "star_removed":
		var event = StarRemovedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_domain_change":
		var event = TeamDomainChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_join":
		var event = TeamJoinEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_migration_started":
		var event = TeamMigrationStartedEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_plan_change":
		var event = TeamPlanChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_pref_change":
		var event = TeamPrefChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "team_rename":
		var event = TeamRenameEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "user_change":
		var event = UserChangeEvent{}
		if err := json.Unmarshal(msg, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "user_typing":
		var event = UserTypingEvent{}
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
