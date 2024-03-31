// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/chapter"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/livecategory"
	"github.com/zibbp/ganymede/ent/playback"
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/schema"
	"github.com/zibbp/ganymede/ent/twitchcategory"
	"github.com/zibbp/ganymede/ent/user"
	"github.com/zibbp/ganymede/ent/vod"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	channelFields := schema.Channel{}.Fields()
	_ = channelFields
	// channelDescRetention is the schema descriptor for retention field.
	channelDescRetention := channelFields[5].Descriptor()
	// channel.DefaultRetention holds the default value on creation for the retention field.
	channel.DefaultRetention = channelDescRetention.Default.(bool)
	// channelDescUpdatedAt is the schema descriptor for updated_at field.
	channelDescUpdatedAt := channelFields[7].Descriptor()
	// channel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	channel.DefaultUpdatedAt = channelDescUpdatedAt.Default.(func() time.Time)
	// channel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	channel.UpdateDefaultUpdatedAt = channelDescUpdatedAt.UpdateDefault.(func() time.Time)
	// channelDescCreatedAt is the schema descriptor for created_at field.
	channelDescCreatedAt := channelFields[8].Descriptor()
	// channel.DefaultCreatedAt holds the default value on creation for the created_at field.
	channel.DefaultCreatedAt = channelDescCreatedAt.Default.(func() time.Time)
	// channelDescID is the schema descriptor for id field.
	channelDescID := channelFields[0].Descriptor()
	// channel.DefaultID holds the default value on creation for the id field.
	channel.DefaultID = channelDescID.Default.(func() uuid.UUID)
	chapterFields := schema.Chapter{}.Fields()
	_ = chapterFields
	// chapterDescID is the schema descriptor for id field.
	chapterDescID := chapterFields[0].Descriptor()
	// chapter.DefaultID holds the default value on creation for the id field.
	chapter.DefaultID = chapterDescID.Default.(func() uuid.UUID)
	liveFields := schema.Live{}.Fields()
	_ = liveFields
	// liveDescWatchLive is the schema descriptor for watch_live field.
	liveDescWatchLive := liveFields[1].Descriptor()
	// live.DefaultWatchLive holds the default value on creation for the watch_live field.
	live.DefaultWatchLive = liveDescWatchLive.Default.(bool)
	// liveDescWatchVod is the schema descriptor for watch_vod field.
	liveDescWatchVod := liveFields[2].Descriptor()
	// live.DefaultWatchVod holds the default value on creation for the watch_vod field.
	live.DefaultWatchVod = liveDescWatchVod.Default.(bool)
	// liveDescDownloadArchives is the schema descriptor for download_archives field.
	liveDescDownloadArchives := liveFields[3].Descriptor()
	// live.DefaultDownloadArchives holds the default value on creation for the download_archives field.
	live.DefaultDownloadArchives = liveDescDownloadArchives.Default.(bool)
	// liveDescDownloadHighlights is the schema descriptor for download_highlights field.
	liveDescDownloadHighlights := liveFields[4].Descriptor()
	// live.DefaultDownloadHighlights holds the default value on creation for the download_highlights field.
	live.DefaultDownloadHighlights = liveDescDownloadHighlights.Default.(bool)
	// liveDescDownloadUploads is the schema descriptor for download_uploads field.
	liveDescDownloadUploads := liveFields[5].Descriptor()
	// live.DefaultDownloadUploads holds the default value on creation for the download_uploads field.
	live.DefaultDownloadUploads = liveDescDownloadUploads.Default.(bool)
	// liveDescDownloadSubOnly is the schema descriptor for download_sub_only field.
	liveDescDownloadSubOnly := liveFields[6].Descriptor()
	// live.DefaultDownloadSubOnly holds the default value on creation for the download_sub_only field.
	live.DefaultDownloadSubOnly = liveDescDownloadSubOnly.Default.(bool)
	// liveDescIsLive is the schema descriptor for is_live field.
	liveDescIsLive := liveFields[7].Descriptor()
	// live.DefaultIsLive holds the default value on creation for the is_live field.
	live.DefaultIsLive = liveDescIsLive.Default.(bool)
	// liveDescArchiveChat is the schema descriptor for archive_chat field.
	liveDescArchiveChat := liveFields[8].Descriptor()
	// live.DefaultArchiveChat holds the default value on creation for the archive_chat field.
	live.DefaultArchiveChat = liveDescArchiveChat.Default.(bool)
	// liveDescResolution is the schema descriptor for resolution field.
	liveDescResolution := liveFields[9].Descriptor()
	// live.DefaultResolution holds the default value on creation for the resolution field.
	live.DefaultResolution = liveDescResolution.Default.(string)
	// liveDescLastLive is the schema descriptor for last_live field.
	liveDescLastLive := liveFields[10].Descriptor()
	// live.DefaultLastLive holds the default value on creation for the last_live field.
	live.DefaultLastLive = liveDescLastLive.Default.(func() time.Time)
	// liveDescRenderChat is the schema descriptor for render_chat field.
	liveDescRenderChat := liveFields[11].Descriptor()
	// live.DefaultRenderChat holds the default value on creation for the render_chat field.
	live.DefaultRenderChat = liveDescRenderChat.Default.(bool)
	// liveDescVideoAge is the schema descriptor for video_age field.
	liveDescVideoAge := liveFields[12].Descriptor()
	// live.DefaultVideoAge holds the default value on creation for the video_age field.
	live.DefaultVideoAge = liveDescVideoAge.Default.(int64)
	// liveDescUpdatedAt is the schema descriptor for updated_at field.
	liveDescUpdatedAt := liveFields[13].Descriptor()
	// live.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	live.DefaultUpdatedAt = liveDescUpdatedAt.Default.(func() time.Time)
	// live.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	live.UpdateDefaultUpdatedAt = liveDescUpdatedAt.UpdateDefault.(func() time.Time)
	// liveDescCreatedAt is the schema descriptor for created_at field.
	liveDescCreatedAt := liveFields[14].Descriptor()
	// live.DefaultCreatedAt holds the default value on creation for the created_at field.
	live.DefaultCreatedAt = liveDescCreatedAt.Default.(func() time.Time)
	// liveDescID is the schema descriptor for id field.
	liveDescID := liveFields[0].Descriptor()
	// live.DefaultID holds the default value on creation for the id field.
	live.DefaultID = liveDescID.Default.(func() uuid.UUID)
	livecategoryFields := schema.LiveCategory{}.Fields()
	_ = livecategoryFields
	// livecategoryDescID is the schema descriptor for id field.
	livecategoryDescID := livecategoryFields[0].Descriptor()
	// livecategory.DefaultID holds the default value on creation for the id field.
	livecategory.DefaultID = livecategoryDescID.Default.(func() uuid.UUID)
	playbackFields := schema.Playback{}.Fields()
	_ = playbackFields
	// playbackDescTime is the schema descriptor for time field.
	playbackDescTime := playbackFields[3].Descriptor()
	// playback.DefaultTime holds the default value on creation for the time field.
	playback.DefaultTime = playbackDescTime.Default.(int)
	// playbackDescUpdatedAt is the schema descriptor for updated_at field.
	playbackDescUpdatedAt := playbackFields[5].Descriptor()
	// playback.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	playback.DefaultUpdatedAt = playbackDescUpdatedAt.Default.(func() time.Time)
	// playback.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	playback.UpdateDefaultUpdatedAt = playbackDescUpdatedAt.UpdateDefault.(func() time.Time)
	// playbackDescCreatedAt is the schema descriptor for created_at field.
	playbackDescCreatedAt := playbackFields[6].Descriptor()
	// playback.DefaultCreatedAt holds the default value on creation for the created_at field.
	playback.DefaultCreatedAt = playbackDescCreatedAt.Default.(func() time.Time)
	// playbackDescID is the schema descriptor for id field.
	playbackDescID := playbackFields[0].Descriptor()
	// playback.DefaultID holds the default value on creation for the id field.
	playback.DefaultID = playbackDescID.Default.(func() uuid.UUID)
	playlistFields := schema.Playlist{}.Fields()
	_ = playlistFields
	// playlistDescUpdatedAt is the schema descriptor for updated_at field.
	playlistDescUpdatedAt := playlistFields[4].Descriptor()
	// playlist.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	playlist.DefaultUpdatedAt = playlistDescUpdatedAt.Default.(func() time.Time)
	// playlist.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	playlist.UpdateDefaultUpdatedAt = playlistDescUpdatedAt.UpdateDefault.(func() time.Time)
	// playlistDescCreatedAt is the schema descriptor for created_at field.
	playlistDescCreatedAt := playlistFields[5].Descriptor()
	// playlist.DefaultCreatedAt holds the default value on creation for the created_at field.
	playlist.DefaultCreatedAt = playlistDescCreatedAt.Default.(func() time.Time)
	// playlistDescID is the schema descriptor for id field.
	playlistDescID := playlistFields[0].Descriptor()
	// playlist.DefaultID holds the default value on creation for the id field.
	playlist.DefaultID = playlistDescID.Default.(func() uuid.UUID)
	queueFields := schema.Queue{}.Fields()
	_ = queueFields
	// queueDescLiveArchive is the schema descriptor for live_archive field.
	queueDescLiveArchive := queueFields[1].Descriptor()
	// queue.DefaultLiveArchive holds the default value on creation for the live_archive field.
	queue.DefaultLiveArchive = queueDescLiveArchive.Default.(bool)
	// queueDescOnHold is the schema descriptor for on_hold field.
	queueDescOnHold := queueFields[2].Descriptor()
	// queue.DefaultOnHold holds the default value on creation for the on_hold field.
	queue.DefaultOnHold = queueDescOnHold.Default.(bool)
	// queueDescVideoProcessing is the schema descriptor for video_processing field.
	queueDescVideoProcessing := queueFields[3].Descriptor()
	// queue.DefaultVideoProcessing holds the default value on creation for the video_processing field.
	queue.DefaultVideoProcessing = queueDescVideoProcessing.Default.(bool)
	// queueDescChatProcessing is the schema descriptor for chat_processing field.
	queueDescChatProcessing := queueFields[4].Descriptor()
	// queue.DefaultChatProcessing holds the default value on creation for the chat_processing field.
	queue.DefaultChatProcessing = queueDescChatProcessing.Default.(bool)
	// queueDescProcessing is the schema descriptor for processing field.
	queueDescProcessing := queueFields[5].Descriptor()
	// queue.DefaultProcessing holds the default value on creation for the processing field.
	queue.DefaultProcessing = queueDescProcessing.Default.(bool)
	// queueDescRenderChat is the schema descriptor for render_chat field.
	queueDescRenderChat := queueFields[17].Descriptor()
	// queue.DefaultRenderChat holds the default value on creation for the render_chat field.
	queue.DefaultRenderChat = queueDescRenderChat.Default.(bool)
	// queueDescUpdatedAt is the schema descriptor for updated_at field.
	queueDescUpdatedAt := queueFields[20].Descriptor()
	// queue.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	queue.DefaultUpdatedAt = queueDescUpdatedAt.Default.(func() time.Time)
	// queue.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	queue.UpdateDefaultUpdatedAt = queueDescUpdatedAt.UpdateDefault.(func() time.Time)
	// queueDescCreatedAt is the schema descriptor for created_at field.
	queueDescCreatedAt := queueFields[21].Descriptor()
	// queue.DefaultCreatedAt holds the default value on creation for the created_at field.
	queue.DefaultCreatedAt = queueDescCreatedAt.Default.(func() time.Time)
	// queueDescID is the schema descriptor for id field.
	queueDescID := queueFields[0].Descriptor()
	// queue.DefaultID holds the default value on creation for the id field.
	queue.DefaultID = queueDescID.Default.(func() uuid.UUID)
	twitchcategoryFields := schema.TwitchCategory{}.Fields()
	_ = twitchcategoryFields
	// twitchcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	twitchcategoryDescUpdatedAt := twitchcategoryFields[4].Descriptor()
	// twitchcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	twitchcategory.DefaultUpdatedAt = twitchcategoryDescUpdatedAt.Default.(func() time.Time)
	// twitchcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	twitchcategory.UpdateDefaultUpdatedAt = twitchcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// twitchcategoryDescCreatedAt is the schema descriptor for created_at field.
	twitchcategoryDescCreatedAt := twitchcategoryFields[5].Descriptor()
	// twitchcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	twitchcategory.DefaultCreatedAt = twitchcategoryDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescOauth is the schema descriptor for oauth field.
	userDescOauth := userFields[4].Descriptor()
	// user.DefaultOauth holds the default value on creation for the oauth field.
	user.DefaultOauth = userDescOauth.Default.(bool)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[8].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	vodFields := schema.Vod{}.Fields()
	_ = vodFields
	// vodDescDuration is the schema descriptor for duration field.
	vodDescDuration := vodFields[5].Descriptor()
	// vod.DefaultDuration holds the default value on creation for the duration field.
	vod.DefaultDuration = vodDescDuration.Default.(int)
	// vodDescViews is the schema descriptor for views field.
	vodDescViews := vodFields[6].Descriptor()
	// vod.DefaultViews holds the default value on creation for the views field.
	vod.DefaultViews = vodDescViews.Default.(int)
	// vodDescProcessing is the schema descriptor for processing field.
	vodDescProcessing := vodFields[8].Descriptor()
	// vod.DefaultProcessing holds the default value on creation for the processing field.
	vod.DefaultProcessing = vodDescProcessing.Default.(bool)
	// vodDescLocked is the schema descriptor for locked field.
	vodDescLocked := vodFields[18].Descriptor()
	// vod.DefaultLocked holds the default value on creation for the locked field.
	vod.DefaultLocked = vodDescLocked.Default.(bool)
	// vodDescLocalViews is the schema descriptor for local_views field.
	vodDescLocalViews := vodFields[19].Descriptor()
	// vod.DefaultLocalViews holds the default value on creation for the local_views field.
	vod.DefaultLocalViews = vodDescLocalViews.Default.(int)
	// vodDescStreamedAt is the schema descriptor for streamed_at field.
	vodDescStreamedAt := vodFields[20].Descriptor()
	// vod.DefaultStreamedAt holds the default value on creation for the streamed_at field.
	vod.DefaultStreamedAt = vodDescStreamedAt.Default.(func() time.Time)
	// vodDescUpdatedAt is the schema descriptor for updated_at field.
	vodDescUpdatedAt := vodFields[21].Descriptor()
	// vod.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vod.DefaultUpdatedAt = vodDescUpdatedAt.Default.(func() time.Time)
	// vod.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vod.UpdateDefaultUpdatedAt = vodDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vodDescCreatedAt is the schema descriptor for created_at field.
	vodDescCreatedAt := vodFields[22].Descriptor()
	// vod.DefaultCreatedAt holds the default value on creation for the created_at field.
	vod.DefaultCreatedAt = vodDescCreatedAt.Default.(func() time.Time)
	// vodDescID is the schema descriptor for id field.
	vodDescID := vodFields[0].Descriptor()
	// vod.DefaultID holds the default value on creation for the id field.
	vod.DefaultID = vodDescID.Default.(func() uuid.UUID)
}
