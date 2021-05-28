package godacl

import (
	"github.com/bwmarrin/snowflake"
)

// Spec from discord API v9:
// https://discord.com/developers/docs/resources/user#user-object

type user_flags int

const (
	None                      = 0
	DiscordEmployee           = 1 << 0
	PartneredServerOwner      = 1 << 1
	HypeSquadEvents           = 1 << 2
	BugHunterLevel1           = 1 << 3
	HouseBravery              = 1 << 6
	HouserBrilliance          = 1 << 7
	HouseBalance              = 1 << 8
	EarlySupporter            = 1 << 9
	TeamUser                  = 1 << 10
	BugHunterLevel2           = 1 << 14
	VerifiedBot               = 1 << 16
	EarlyVerifiedBotDeveloper = 1 << 17
	DiscordCertifiedModerator = 1 << 18
)

type premium_type int

const (
	// None = 0
	Nitro_Classic = 1
	Nitro         = 2
)

type UserData struct {
	_id            snowflake.Node // the user's id
	_username      string         // the user's username, not unique across the platform
	_discriminator string         // the user's 4-digit discord-tag
	_avatar        string         // the user's avatar hash
	_bot           bool           // whether the user belongs to an OAuth2 application
	_system        bool           // whether the user is an Official Discord System user (part of the urgent message system)
	_mfa_enabled   bool           // whether the user has two factor enabled on their account
	_locale        string         // the user's chosen language option
	_verified      bool           // whether the email on this account has been verified
	_email         string         // the user's email
	_flags         user_flags     // the flags on a user account
	_premium_type  premium_type   // the type of Nitro subscription on a user's account
	_public_flags  user_flags     // the public flags on a user's account
}

type image_data string

// GET/users/@me
// Returns the user object of the requester's account. For OAuth2, this requires the identify scope, which will return the object without an email, and optionally the email scope, which returns the object with an email.
func (a *DiscordAPI) GetCurrentUser() (UserData, error) {
	return UserData{}, ErrorNotImplemented
}

// GET/users/{user.id}
// Returns a user object for a given user ID.
func (a *DiscordAPI) GetUser(id snowflake.Node) (UserData, error) {
	return UserData{}, ErrorNotImplemented
}

// PATCH/users/@me
// Modify the requester's user account settings. Returns a user object on success.
func (a *DiscordAPI) ModifyCurrentUser(username string, avatar image_data) (UserData, error) {
	return UserData{}, ErrorNotImplemented
}

// GET/users/@me/guilds
// Returns a list of partial guild objects the current user is a member of. Requires the guilds OAuth2 scope.
func (a *DiscordAPI) GetCurrentUserGuilds(before snowflake.Node, after snowflake.Node, limit int) ([]GuildData, error) {
	return nil, ErrorNotImplemented
}

// DELETE/users/@me/guilds/{guild.id}
// Leave a guild. Returns a 204 empty response on success.
func (a *DiscordAPI) LeaveGuild(id snowflake.Node) error {
	return ErrorNotImplemented
}

// POST/users/@me/channels
// Create a new DM channel with a user. Returns a DM channel object.
// ! You should not use this endpoint to DM everyone in a server about something. DMs should generally be initiated by a user action. If you open a significant amount of DMs too quickly, your bot may be rate limited or blocked from opening new ones.
func (a *DiscordAPI) CreateDM(recipient_id snowflake.Node) (ChannelData, error) {
	return ChannelData{}, ErrorNotImplemented
}

// POST/users/@me/channels
// Create a new group DM channel with multiple users. Returns a DM channel object. This endpoint was intended to be used with the now-deprecated GameBridge SDK. DMs created with this endpoint will not be shown in the Discord client
func (a *DiscordAPI) CreateGroupDM(access_tokens []string, nicks map[snowflake.Node]string) (ChannelData, error) {
	return ChannelData{}, ErrorNotImplemented
}

// GET/users/@me/connections
// Returns a list of connection objects. Requires the connections OAuth2 scope.
func (a *DiscordAPI) GetUserConnections() ([]ConnectionData, error) {
	return nil, ErrorNotImplemented
}
