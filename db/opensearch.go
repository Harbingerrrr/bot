package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/opensearch-project/opensearch-go"
)

type OpenSearchClient struct {
	Client *opensearch.Client
}

var openSearchClient *OpenSearchClient

func ConnectOpenSearch(domain string, auth string) error {
	client, error := opensearch.NewClient(opensearch.Config{
		Addresses: []string{domain},
		Username:  strings.Split(auth, ":")[0],
		Password:  strings.Split(auth, ":")[1],
	})

	if error != nil {
		log.Fatalf("Error creating OpenSearch client: %q", error)
	}

	openSearchClient = &OpenSearchClient{Client: client}
	fmt.Println("Connected to OpenSearch")
	return nil
}

func GetOpenSearchClient() *opensearch.Client {
	if openSearchClient == nil {
		log.Fatalf("OpenSearch client not connected")
	}

	return openSearchClient.Client
}

/*
	OpenSearch Indexes:
	* /opensearch/guilds_configurations
	- guilds_configurations
	-- each indexed document is a guild object
	---- group_id, group_name, roles, role_per_xp, blacklisted_users, subscription_status, subscription_end_date, subscription_start_date, subscription_plan

	* /opensearch/group_xp
	- group_xp
	-- each indexed document is an array of user objects
	---- user_id, username, xp_total, audit_log

	* /opensearch/admins
	- admins
	-- each indexed document is a discord id that is an admin of Harbinger
	---- discord_id
*/

func (os *OpenSearchClient) ConfigureGuild(guildID *int) {

}

// QueryGuild queries the OpenSearch index for the specified guild
func (os *OpenSearchClient) QueryGuild(guildID *int) {

}

// QueryXP queries the OpenSearch index for the specified user in the specified group
func (os *OpenSearchClient) QueryXP(userID *int, groupID *int) int {
	if openSearchClient == nil {
		log.Fatalf("OpenSearch client not connected")
	}

	// search /opensearch/group_xp/<group_id> for user_id
	resp, error := openSearchClient.Client.Search()
	if error != nil {
		log.Fatalf("Error querying OpenSearch: %q", error)
	}
	defer resp.Body.Close()

	log.Printf("Querying XP for user %d in group %d: %d", *userID, *groupID, 123)

	return 0
}
