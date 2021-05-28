package godacl

import (
	"errors"

	"github.com/bwmarrin/snowflake"
)

type IDiscordAPI interface {

	// User
	GetCurrentUser() (UserData, error)
	GetUser(id snowflake.Node) (UserData, error)
	ModifyCurrentUser(username string, avatar image_data) (UserData, error)
	GetCurrentUserGuilds(before snowflake.Node, after snowflake.Node, limit int) ([]GuildData, error)
	LeaveGuild(id snowflake.Node) error
	CreateDM(recipient_id snowflake.Node) (ChannelData, error)
	CreateGroupDM(access_tokens []string, nicks map[snowflake.Node]string) (ChannelData, error)
	GetUserConnections() ([]ConnectionData, error)
}

type DiscordAPI struct{}

var ErrorNotImplemented = errors.New("This function is not implemented.")
