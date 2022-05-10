package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/dntAtMe/discord-support_bot/src/locale"
)

var usedLocale locale.Locale = locale.Locale_pl 

var yesOrNoButtons = []discordgo.MessageComponent {
    discordgo.Button {
        Label: usedLocale.Yes,
        Style: discordgo.SuccessButton,
        Disabled: false,
        CustomID: "yes",
    },
    discordgo.Button {
        Label: usedLocale.No,
        Style: discordgo.DangerButton,
        Disabled: false,
        CustomID: "no",
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
