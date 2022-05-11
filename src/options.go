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

var yesOrNoButtons = []discordgo.MessageComponent {
    discordgo.ActionsRow {
        Components: []discordgo.MessageComponent {
            discordgo.Button {
                Label: usedLocale.BUTTON_YES,
                Style: discordgo.SuccessButton,
                Disabled: false,
                CustomID: "yes",
            },
            discordgo.Button {
                Label: usedLocale.BUTTON_NO,
                Style: discordgo.DangerButton,
                Disabled: false,
                CustomID: "no",
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
    {
        Label: "Propozycja biznesu",
        Value: "gang",
        Emoji: discordgo.ComponentEmoji {
            Name: "💼",
        },
        Default: false,
        Description: "Jeśli masz pomysł na biznes który chciałbyś prowadzić, tutaj możesz go opisać.",
    },
}
