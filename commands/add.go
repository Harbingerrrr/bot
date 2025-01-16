package commands

import (
	"fmt"
	"strconv"

	"github.com/Harbingerrrr/bot/util"
	"github.com/bwmarrin/discordgo"
)

// Add adds xp to the specified users
// Example: !add 1 username1 username2 username3
func Add(session *discordgo.Session, msg *discordgo.MessageCreate, args *[]string) {
	// Error handling
	if len(*args) < 2 {
		session.ChannelMessageSend(msg.ChannelID, "Usage: !add 1 OutStrike [...]")
		return
	}

	xp, err := strconv.Atoi((*args)[0])
	if err != nil {
		session.ChannelMessageSend(msg.ChannelID, "Second argument must be an integer\nUsage: !add 1 OutStrike [...]")
		return
	}

	// Get group ID
	groupID := util.GetGroupID(&msg.GuildID)

	usernames := (*args)[1:]

	// Get user IDs
	userIDs := util.GetUserIDsFromUsernames(&usernames)
	fmt.Println(userIDs)

	// Add xp to users (not implemented)
	addXP(&userIDs, &xp, &groupID)

}

func addXP(userIDs *[]int, xp *int, groupID *int) {
	for _, userID := range *userIDs {
		fmt.Println(userID, xp)
		// update data in postgres database

		// check if user can promote
		util.QuickCheck(&userID, groupID)
	}
}
