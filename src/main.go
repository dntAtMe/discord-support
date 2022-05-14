package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
    var auth Auth

    readAuth(&auth, "auth.json")

    botToken := fmt.Sprintf("Bot %s", auth.Token)

    discord, err := discordgo.New(botToken)

    if err != nil {
        fmt.Println("error when creating Discord session,", err)
        return
    }

    discord.AddHandler(interactionHandler)
    discord.AddHandler(messageHandler)

    /*
    _, err = discord.ApplicationCommandCreate(auth.AppID, auth.GuildID, &discordgo.ApplicationCommand {
        Name: "test",
        Description: "Test me, daddy",
    })
    */

    if err != nil {
        log.Fatalf("Cannot create command: %v", err)
    }

    err = discord.Open()

    if err != nil {
        fmt.Println("error when opening connection,", err)
    }
    defer discord.Close()

    fmt.Println("Bot is now running")
    
    setupKillSignals()
}

func setupKillSignals() {
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
}


// Handles any type of interaction
func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    switch i.Type {
    // case discordgo.InteractionApplicationCommand:
    //    basicInteractionHandler(s, i)
    case discordgo.InteractionMessageComponent:
        componentInteractionHandler(s, i)
    }
}

// Generates permissions for roles assigned to this category and for the user who created this topic
//
// Interaction context is needed to get the user that started interaction.
// Flags can be specified via https://discord.com/developers/docs/topics/permissions
func generatePermissionsForCategory(categoryName string, i *discordgo.InteractionCreate, flags int64) []*discordgo.PermissionOverwrite {
    permissions := []*discordgo.PermissionOverwrite{}

    // Add default roles
    for _, role := range defaultCategoryRoles {
        permissions = append(permissions, &discordgo.PermissionOverwrite {
            ID: role.ID,
            Type: discordgo.PermissionOverwriteTypeRole,
            Allow: flags,
        })
    }

    // Add roles assigned to this topic
    allowedRoles := categoryRoles[categoryName]

    for _, role := range allowedRoles {
        permissions = append(permissions, &discordgo.PermissionOverwrite {
            ID: role.ID,
            Type: discordgo.PermissionOverwriteTypeRole,
            Allow: flags,
        })
    }

    // Remove @everyone
    permissions = append(permissions, &discordgo.PermissionOverwrite {
        ID: "331504113852612609",
        Type: discordgo.PermissionOverwriteTypeRole,
        Deny: flags,
    })

    // Add interacting user to this topic
    permissions = append(permissions, &discordgo.PermissionOverwrite {
        ID: i.Member.User.ID,
        Type: discordgo.PermissionOverwriteTypeMember,
        Allow: flags,
    })

    // Add self to this topic
    permissions = append(permissions, &discordgo.PermissionOverwrite {
        ID: i.Interaction.AppID,
        Type: discordgo.PermissionOverwriteTypeMember,
        Allow: flags,
    })

    return permissions
}

// Handles interaction with a button; note that buttons might have same ID as menu options for support topics, 
// but since we handle buttons separately, it's not really an issue 
//
func buttonInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    componentId := i.MessageComponentData().CustomID

    // Go through all existing categories and check if button pressed is confirmation of a category creation
    for _, categoryOption := range helpCategories {
        if componentId == categoryOption.Value {
            channel, err := s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData {
                Name: fmt.Sprintf("%s-%d", componentId, rand.Int31n(99999)),
                Type: 0,
                ParentID: supportCategories[ENV],
                PermissionOverwrites: generatePermissionsForCategory(componentId, i, 1 << 10),
            })

            if err != nil {
                fmt.Printf("Error when creating a channel")
            }

            s.ChannelMessageSendComplex(channel.ID, closeTopicMessage(componentId))

            s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
                Type: discordgo.InteractionResponseUpdateMessage,
                Data: &discordgo.InteractionResponseData {
                    Content: "Temat założony na kanale " + channel.Mention(),
                    Flags: 1 << 6,
                    Components: []discordgo.MessageComponent {},
                },
            })

            err = s.InteractionResponseDelete(i.Interaction)

            if err != nil {
                fmt.Println(err)
            }
        }    
    }

    if componentId == "close-topic" {
        s.ChannelDelete(i.ChannelID)
    }
}

// Handles interaction with any component
//
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
        categoryName := i.MessageComponentData().Values[0]

        err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseUpdateMessage,
            Data: &discordgo.InteractionResponseData {
                Content: categoryDescriptions[categoryName],
                Flags: 1 << 6,
                Components: yesOrNoButtons(categoryName, "no"),
            },
        })

        if err != nil {
            fmt.Println(err)
        }
    }
}


// Intercepts user messages and handles them
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages from self
    if m.Author.ID == s.State.User.ID {
        return
    }

    // Only consider messages from support channel 
    if m.ChannelID != supportChannels[ENV] {
        return
    }

    if m.Content == "ping" {
        _, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend {
            Content: "description",
            Components: helpButton,
        })

        if err != nil {
            fmt.Println(err)
        }
    }
}


