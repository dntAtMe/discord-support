package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/dntAtMe/discord-support_bot/src/locale"
)

var usedLocale locale.Locale = locale.Locale_pl 

var helpButton = []discordgo.MessageComponent {
    discordgo.ActionsRow {
        Components: []discordgo.MessageComponent {
            discordgo.Button {
                Label: usedLocale.BUTTON_HELP,
                Style: discordgo.PrimaryButton,
                Disabled: false, 
                CustomID: "help",
            },
        },
    },
}

func yesOrNoButtons(yesId string, noId string) []discordgo.MessageComponent {
    return []discordgo.MessageComponent {
        discordgo.ActionsRow {
            Components: []discordgo.MessageComponent {
                discordgo.Button {
                    Label: usedLocale.BUTTON_YES,
                    Style: discordgo.SuccessButton,
                    Disabled: false,
                    CustomID: yesId,
                },
                discordgo.Button {
                    Label: usedLocale.BUTTON_NO,
                    Style: discordgo.DangerButton,
                    Disabled: false,
                    CustomID: noId,
                },
            },
        },
    }
}

var closeTopicMessage  = &discordgo.MessageSend {
    Content: "",
    Components: []discordgo.MessageComponent {
        discordgo.ActionsRow {
            Components: []discordgo.MessageComponent {
                discordgo.Button {
                    Label: usedLocale.BUTTON_CLOSE_TOPIC,
                    Style: discordgo.DangerButton,
                    Disabled: false,
                    CustomID: "close-topic",
                },
            },
        },
    },
}

var helpMenu = []discordgo.MessageComponent {
    discordgo.ActionsRow {
        Components: []discordgo.MessageComponent {
            discordgo.SelectMenu {
                CustomID: "select_category",
                Placeholder: usedLocale.MENU_HELP_PLACEHOLDER,
                Options: helpCategories,
            },
        },
    },
}

var helpCategories = []discordgo.SelectMenuOption {
    {
        Label: "Propozycja biznesu",
        Value: "business",
        Emoji: discordgo.ComponentEmoji {
            Name: "💼",
        },
        Default: false,
        Description: "Jeśli masz pomysł na biznes który chciałbyś prowadzić, tutaj możesz go opisać.",
    },
}

var categoryRoles = map[string][]Role {
    "business": { roles["CommunityManager"] },
}
