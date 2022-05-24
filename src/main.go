package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
	case discordgo.InteractionModalSubmit:
		modalInteractionHandler(s, i)
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
		permissions = append(permissions, &discordgo.PermissionOverwrite{
			ID:    role.ID,
			Type:  discordgo.PermissionOverwriteTypeRole,
			Allow: flags,
		})
	}

	// Add roles assigned to this topic
	allowedRoles := categoryRoles[categoryName]

	for _, role := range allowedRoles {
		permissions = append(permissions, &discordgo.PermissionOverwrite{
			ID:    role.ID,
			Type:  discordgo.PermissionOverwriteTypeRole,
			Allow: flags,
		})
	}

	// Remove @everyone
	permissions = append(permissions, &discordgo.PermissionOverwrite{
		ID:   "331504113852612609",
		Type: discordgo.PermissionOverwriteTypeRole,
		Deny: flags,
	})

	// Add interacting user to this topic
	permissions = append(permissions, &discordgo.PermissionOverwrite{
		ID:    i.Member.User.ID,
		Type:  discordgo.PermissionOverwriteTypeMember,
		Allow: flags,
	})

	// Add self to this topic
	permissions = append(permissions, &discordgo.PermissionOverwrite{
		ID:    i.Interaction.AppID,
		Type:  discordgo.PermissionOverwriteTypeMember,
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
			uuid := uuid.New()
			channel, err := s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
				Name:                 fmt.Sprintf("%s_%s", componentId, uuid.String()),
				Type:                 0,
				ParentID:             supportCategories[ENV],
				PermissionOverwrites: generatePermissionsForCategory(componentId, i, 1<<10),
			})

			if err != nil {
				fmt.Printf("Error when creating a channel")
			}

			s.ChannelMessageSendComplex(channel.ID, closeTopicMessage(i.Member.User, componentId))

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    "Temat założony na kanale " + channel.Mention(),
					Flags:      1 << 6,
					Components: []discordgo.MessageComponent{},
				},
			})

			err = s.InteractionResponseDelete(i.Interaction)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if componentId == "close-topic" {
		roleFound := false
		for _, role := range i.Interaction.Member.Roles {
			if role == roles["CommunityManager"].ID || role == roles["ProjectManager"].ID {
				roleFound = true

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseModal,
					Data: &discordgo.InteractionResponseData{
						CustomID: "close-topic-m_" + i.Message.Mentions[0].ID + "_" + i.ChannelID,
						Title:    "Zamknięcie tematu",
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.TextInput{
										CustomID:    "response",
										Label:       "Odpowiedź:",
										Style:       discordgo.TextInputParagraph,
										Placeholder: "Puste jeśli bez odpowiedzi",
										Required:    false,
										MaxLength:   300,
										MinLength:   0,
									},
								},
							},
						},
					},
				})

				if err != nil {
					fmt.Println(err)
				}

			}
		}

		// Not a Community Manager or Project Manager
		if !roleFound {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    "Nie masz uprawnień do zamykania tematu",
					Flags:      1 << 6,
					Components: []discordgo.MessageComponent{},
				},
			})
		}
	}

	if componentId == "leave" {
		s.ChannelPermissionSet(i.ChannelID, i.Member.User.ID, discordgo.PermissionOverwriteTypeMember, 0, 1<<10)
		s.ChannelMessageSend(i.ChannelID, "<@"+i.Member.User.ID+"> opuścił kanał")
	}

	if componentId == "no" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    ">>> Wybierz kategorię",
				Flags:      1 << 6,
				Components: helpMenu,
			},
		})
	}
}

func reverse(arr []*discordgo.Message) []*discordgo.Message {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func modalInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	componentId := data.CustomID

	if strings.HasPrefix(componentId, "close-topic-m") {
		requesteeID := strings.Split(componentId, "_")[1]
		requestee, err := s.User(requesteeID)

		channelID := strings.Split(componentId, "_")[2]

		fmt.Printf("Requestee: %s\n", requesteeID)
		fmt.Printf("ChannelID: %s\n", channelID)

		responseMessage := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

		fmt.Printf("Response: %s\n", responseMessage)

		if strings.Trim(responseMessage, " \n") != "" {
			// Estabilish private chat with requestee
			privateChannel, _ := s.UserChannelCreate(requesteeID)

			// Send requestee a message if there is something to send
			s.ChannelMessageSend(privateChannel.ID, "Zamknięto temat: "+responseMessage)
		}
		// Copy all messages to database
		messages, err := s.ChannelMessages(channelID, 100, "", "", "")

		if err != nil {
			fmt.Println(err)
		}

		supportChannel, err := s.Channel(channelID)

		// Create channel in archive
		archiveCategory, _ := s.Channel("978772539667009687")
		archiveChannel, _ := s.GuildChannelCreateComplex(archiveCategory.GuildID, discordgo.GuildChannelCreateData{
			Name:     strings.Split(supportChannel.Name, "_")[0] + "_" + requestee.Username,
			Type:     0,
			ParentID: archiveCategory.ID,
		})

		for _, message := range reverse(messages) {
			/*
				      db.LogMessage(&db.Message{
								AuthorId: message.Author.ID,
								Content:  message.Content,
								Date:     message.Timestamp.String(),
							}, supportChannel.Name)
			*/
			var attachmentsContent string = ""
			for _, attach := range message.Attachments {
				attachmentsContent += "\n" + attach.ProxyURL
			}

			s.ChannelMessageSendComplex(archiveChannel.ID, &discordgo.MessageSend{
				Content:    fmt.Sprintf("%s (%s): %s %s", message.Author.Mention(), message.Author.Username, message.Content, attachmentsContent),
				Embeds:     message.Embeds,
				Components: message.Components,
			})
		}

		s.ChannelDelete(channelID)
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
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    ">>> Wybierz kategorię",
				Flags:      1 << 6,
				Components: helpMenu,
			},
		})

		if err != nil {
			fmt.Println(err)
		}
	case "select_category":
		categoryName := i.MessageComponentData().Values[0]

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    categoryDescriptions[categoryName],
				Flags:      1 << 6,
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
		_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: `
>>> Jeśli masz jakiś temat do omówienia z administracją i nie wiesz do kogo się zgłosić, Support przeprowadzi Cię przez następne kroki, aby ułatwić nam komunikację i przyśpieszyć cały proces. 
Przed założeniem nowego tematu, upewnij się że Twój problem nie został rozwiązany w <#753276779715756113>.
            `,
			Components: helpButton,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

}
