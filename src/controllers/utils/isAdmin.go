package utils

import (
	. "Bot1/src/config"

	"github.com/disgoorg/disgo/discord"
)

func IsAdmin(member discord.Member) bool {
	permissions := Client.Caches().MemberPermissions(member)
	return permissions.Has(discord.PermissionAdministrator)
}
