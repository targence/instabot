package instabot

import "time"

// PostBackEvent defines flatten data for postback event.
type PostBackEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	Data      *Postback
}

// GetPostBackEvent returns postback event.
// Call this only when the event type is WebhookEventTypePostBack.
func (m *Messaging) GetPostBackEvent() *PostBackEvent {
	return &PostBackEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
		Data:      m.PostBack,
	}
}

// QuickReplyEvent defines flatten quick reply event data.
type QuickReplyEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	MID       string
	Text      string
	Data      *WebhookQuickReply
}

// GetQuickReplyEvent returns quick reply event.
// Call this only when the event type is WebhookEventTypeQuickReply.
func (m *Messaging) GetQuickReplyEvent() *QuickReplyEvent {
	quickReply := &QuickReplyEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		quickReply.MID = m.Message.MID
		quickReply.Text = m.Message.Text
		quickReply.Data = m.Message.QuickReply
	}

	return quickReply
}

// StoryMentionEvent defines flatten story mention event.
type StoryMentionEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	MID       string
	Story     *ReplyToStory
}

// GetStoryMentionEvent returns story mention event.
// Call this only when the event type is WebhookEventTypeStoryMention.
func (m *Messaging) GetStoryMentionEvent() *StoryMentionEvent {
	storyMentionEvent := &StoryMentionEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
		MID:       m.Message.MID,
	}

	if m.Message != nil {
		storyMentionEvent.MID = m.Message.MID

		if len(m.Message.Attachments) > 0 {
			storyMentionEvent.Story = &ReplyToStory{
				URL: m.Message.Attachments[0].Payload.URL,
			}
		}
	}

	return storyMentionEvent
}

// StoryReplyEvent defines flatten story reply event.
type StoryReplyEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	MID       string
	Text      string
	Story     *ReplyToStory
}

// GetStoryReplyEvent returns story reply event.
// Call this only when the event type is WebhookEventTypeStoryReply.
func (m *Messaging) GetStoryReplyEvent() *StoryReplyEvent {
	storyReplyEvent := &StoryReplyEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		storyReplyEvent.MID = m.Message.MID
		storyReplyEvent.Text = m.Message.Text

		if m.Message.ReplyTo != nil {
			storyReplyEvent.Story = m.Message.ReplyTo.Story
		}
	}

	return storyReplyEvent
}

type ReferralMessageEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	Referral  *ReferralLink
}

func (m *Messaging) GetReferralEvent() *ReferralMessageEvent {
	referralMessageEvent := &ReferralMessageEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
		Referral:  m.Referral,
	}

	return referralMessageEvent
}

// TextMessageEvent defines flatten text message event.
type TextMessageEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	MID       string
	Text      string
	Referral  *ReferralLink
}

// GetTextMessageEvent returns text message event.
// Call this only when the event type is WebhookEventTypeTextMessage.
func (m *Messaging) GetTextMessageEvent() *TextMessageEvent {
	textMessageEvent := &TextMessageEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
		MID:       m.Message.MID,
		Text:      m.Message.Text,
		Referral:  m.Message.Referral,
	}

	if m.Message != nil {
		textMessageEvent.MID = m.Message.MID
		textMessageEvent.Text = m.Message.Text
		textMessageEvent.Referral = m.Message.Referral
	}

	return textMessageEvent
}

// MediaMessageEvent defines flatten media message event.
type MediaMessageEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	MID       string
	Type      WebhookEventType
	Media     *AttachmentPayload
}

// GetMediaMessageEvent returns media (image, audio, video, file) message event.
// Call this only when the event type is one of
// WebhookEventTypeImageMessage, WebhookEventTypeAudioeMessage,
// WebhookEventTypeVideoMessage or WebhookEventTypeFileMessage.
func (m *Messaging) GetMediaMessageEvent() *MediaMessageEvent {
	mediaMessageEvent := &MediaMessageEvent{
		Type:      m.Type,
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		mediaMessageEvent.MID = m.Message.MID

		if len(m.Message.Attachments) > 0 {
			mediaMessageEvent.Media = &m.Message.Attachments[0].Payload
		}
	}

	return mediaMessageEvent
}

// MessageReplyEvent defines flatten message reply event.
type MessageReplyEvent struct {
	Sender     *Sender
	Recipient  *Recipient
	Timestamp  time.Time
	MID        string
	Text       string
	ReplyToMID string
}

// GetMessageReplyEvent returns message reply event.
// Call this only when the event type is WebhookEventTypeMessageReply.
func (m *Messaging) GetMessageReplyEvent() *MessageReplyEvent {
	messageReplyEvent := &MessageReplyEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		messageReplyEvent.MID = m.Message.MID
		messageReplyEvent.Text = m.Message.Text

		if m.Message.ReplyTo != nil {
			messageReplyEvent.ReplyToMID = m.Message.ReplyTo.MID
		}
	}

	return messageReplyEvent
}

// MessageReactionEvent defines flatten message react event.
type MessageReactionEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	Reaction  *Reaction
}

// GetMessageReactionEvent returns message reaction event.
// Call this only when the event type is WebhookEventTypeReaction.
func (m *Messaging) GetMessageReactionEvent() *MessageReactionEvent {
	return &MessageReactionEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
		Reaction:  m.Reaction,
	}
}

// MessageSeenEvent defines message seen event.
type MessageSeenEvent struct {
	Sender    *Sender
	Recipient *Recipient
	Timestamp time.Time
	SeenMID   string
}

// GetMessageSeenEvent returns message seen event.
// Call this only when the event type is WebhookEventTypeMessageSeen.
func (m *Messaging) GetMessageSeenEvent() *MessageSeenEvent {
	messageSeenEvent := &MessageSeenEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Read != nil {
		messageSeenEvent.SeenMID = m.Read.MID
	}

	return messageSeenEvent
}

// MessageShareEvent defines message share events.
type MessageShareEvent struct {
	Sender           *Sender
	Recipient        *Recipient
	Timestamp        time.Time
	SharedPayloadURL string
}

// GetMessageShareEvent returns message share event.
// Call this only when the event type is WebhookEventTypeShare.
func (m *Messaging) GetMessageShareEvent() *MessageShareEvent {
	messageShareEvent := &MessageShareEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		if len(m.Message.Attachments) > 0 {
			messageShareEvent.SharedPayloadURL = m.Message.Attachments[0].Payload.URL
		}
	}

	return messageShareEvent
}

// MessageDeleteEvent defines message delete events.
type MessageDeleteEvent struct {
	Sender     *Sender
	Recipient  *Recipient
	Timestamp  time.Time
	DeletedMID string
}

// GetMessageDeleteEvent returns message delete event.
// Call this only when the event type is WebhookEventTypeDeleted.
func (m *Messaging) GetMessageDeleteEvent() *MessageDeleteEvent {
	messageDeletevent := &MessageDeleteEvent{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Timestamp: time.Unix(m.Timestamp, 0).UTC(),
	}

	if m.Message != nil {
		messageDeletevent.DeletedMID = m.Message.MID
	}

	return messageDeletevent
}
