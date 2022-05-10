package main

import (
    "encoding/json"
    "log"
    "io/ioutil"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

type Auth struct {
    Token string `json:"token"`
    AppID string `json:"appId"`
    GuildID string `json:"guildId"`
}

var configPath = "config/"

func ReadAuth(auth *Auth) {
    authPath := fmt.Sprintf("%sauth.json", configPath)
    jsonFile, err := os.Open(authPath)
    
    defer jsonFile.Close()

    // TODO: Should throw an error instead
    if err != nil {
        fmt.Println(err)
        return
    }

    byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValue, auth)
}

func main() {
    var auth Auth

    ReadAuth(&auth)

    botToken := fmt.Sprintf("Bot %s", auth.Token)

    discord, err := discordgo.New(botToken)

    if err != nil {
        fmt.Println("error when creating Discord session,", err)
        return
    }

    discord.AddHandler(interactionHandler)
    // discord.Identify.Intents = discordgo.IntentsGuildMessages

    _, err = discord.ApplicationCommandCreate(auth.AppID, auth.GuildID, &discordgo.ApplicationCommand {
        Name: "test",
        Description: "Test me, daddy",
    })

    if err != nil {
        log.Fatalf("Cannot create command: %v", err)
    }



    err = discord.Open()

    if err != nil {
        fmt.Println("error when opening connection,", err)
    }
    
    defer discord.Close()

    fmt.Println("Bot is now running")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    switch i.Type {
    case discordgo.InteractionApplicationCommand:
        basicInteractionHandler(s, i)
    case discordgo.InteractionMessageComponent:
        componentId := i.MessageComponentData().CustomID
        componentInteractionHandler(s, i, componentId)
    }
}

func componentInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate, componentId string) {
    if componentId == "yes" {
        fmt.Println("Yes pressed")
    }
}


func basicInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData {
            Content: "Dupa",
            Flags: 1 << 6,
            Components: []discordgo.MessageComponent {
                discordgo.ActionsRow {
                    Components: yesOrNoButtons,
                },
                discordgo.ActionsRow {
                    Components: []discordgo.MessageComponent {
                        discordgo.SelectMenu {
                            CustomID: "select",
                            Placeholder: "Z czym potrzebujesz pomocy?",
                            Options: helpCategories,
                        },
                    },
                },
            },
        },
    })

    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseModal,
            Data: &discordgo.InteractionResponseData{
                CustomID: "modals_survey_" + i.Interaction.Member.User.ID,
                Title:    "Modals survey",
                Components: []discordgo.MessageComponent{
                    discordgo.ActionsRow{
                        Components: []discordgo.MessageComponent{
                            discordgo.TextInput{
                                CustomID:    "opinion",
                                Label:       "What is your opinion on them?",
                                Style:       discordgo.TextInputShort,
                                Placeholder: "Don't be shy, share your opinion with us",
                                Required:    true,
                                MaxLength:   300,
                                MinLength:   10,
                            },
                        },
                    },
                    discordgo.ActionsRow{
                        Components: []discordgo.MessageComponent{
                            discordgo.TextInput{
                                CustomID:  "suggestions",
                                Label:     "What would you suggest to improve them?",
                                Style:     discordgo.TextInputParagraph,
                                Required:  false,
                                MaxLength: 2000,
                            },
                        },
                    },
                },
            },
    })
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.GuildID != "973657194283274251" {
        return
    }
    // Ignore messages from self
    if m.Author.ID == s.State.User.ID {
        return
    }

    if m.Content == "ping" {
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
}


