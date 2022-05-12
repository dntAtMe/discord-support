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
    discord.AddHandler(messageHandler)
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
        componentInteractionHandler(s, i)
    }
}

func buttonInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    componentId := i.MessageComponentData().CustomID

    // Confirmation on business topic creation
    // TODO: Automatically assign name, parent ID and role overwrites based on chosen option
    if componentId == "business" {
        s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData {
            Name: "test-topic",
            Type: 0,
            ParentID: "974322254227841044",
            PermissionOverwrites: []*discordgo.PermissionOverwrite {
                {
                    ID: i.Member.User.ID,
                    Type: discordgo.PermissionOverwriteTypeMember,
                    Allow: 1 << 10,
                },
            },
        })
    }
}

func componentInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    componentId := i.MessageComponentData().CustomID

    // Triggered by a button
    if i.Message != nil {
        buttonInteractionHandler(s, i)
    }

    switch componentId {
    case "help":
        err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData {
                Content: "Wybierz kategorię",
                Flags: 1 << 6,
                Components: helpMenu,
            },
        })

        if err != nil {
            fmt.Println(err)
        }
    case "select_category":
        // value: i.MessageComponentData().Values[0]

        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseUpdateMessage,
            Data: &discordgo.InteractionResponseData {
                Content: "Aby ubiegać się o biznes, to x y z. Czy chcesz założyć nowy temat?",
                Flags: 1 << 6,
                Components: yesOrNoButtons("business", "no"),
            },
        })
    }
}


func basicInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData {
            Content: "Dupa",
            Flags: 1 << 6,
            Components: helpMenu,
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
    // Only consider messages from dev guild
    if m.GuildID != "973657194283274251" {
        return
    }
    // Ignore messages from self
    if m.Author.ID == s.State.User.ID {
        return
    }

    if m.Content == "ping" {
        _, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend {
            Content: "Opis że przycisk wcisnac jesli jest potrzeba kontaktu z administracja",
            Components: helpButton,
        })

        if err != nil {
            fmt.Println(err)
        }
    }
}


