package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Harbingerrrr/bot/commands"
	"github.com/Harbingerrrr/bot/db"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session, error := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if error != nil {
		log.Fatal("Error creating Discord session: ", error)
		return
	}
	defer session.Close()

	// Postgres
	// db.Connect(os.Getenv("DB_CONNECTION_STRING"))
	// defer db.Close()

	// OpenSearch
	db.ConnectOpenSearch(os.Getenv("OPENSEARCH_DOMAIN"), os.Getenv("OPENSEARCH_BASIC_AUTH"))

	// commands.InitOpenSearchClient(&db.OpenSearchClient{Client: db.GetOpenSearchClient()})

	session.AddHandler(handleMessage)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	error = session.Open()
	if error != nil {
		log.Fatal("Error opening connection to Discord: ", error)
		return
	}
	defer session.Close()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	shutdown()
}

func handleMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	log.Print("Author: ", msg.Author, " (", msg.Author.ID, ")\nMessage received: ", msg.Content)
	commands.HandleCommand(session, msg)

}

func shutdown() {
	db.Close()

	shudownChannel := make(chan os.Signal, 1)
	signal.Notify(shudownChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-shudownChannel

	fmt.Println("Shutting down...")
}
