package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/dntAtMe/discord-support_bot/src/locale"
)

var ENV = "prod"

var supportCategories = map[string]string {
    "dev": "974322254227841044",
    "prod": "975035840994627616",
}

var supportChannels = map[string]string {
    "dev": "973678484620722256",
    "prod": "975035944627503144",
}

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

func closeTopicMessage(category string) *discordgo.MessageSend {
    return &discordgo.MessageSend {
        Content: categoryCreationInfo[category],
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

/* You can define categories for support options like that

var helpCategories = []discordgo.SelectMenuOption {
    {
        Label: "Label",
        Value: "inner-value", // Not really important, it's an ID of sorts so needs to be unique
        Emoji: discordgo.ComponentEmoji {
            Name: "💼",
        },
        Default: false,
        Description: "short description",
    },
    {
        ...
    }
}
*/

/*
var categoryDescriptions = map[string]string {
    "inner-value": `
    Here you can insert longer description when confirmation message pops up
    `,
}
*/

/*
var categoryCreationInfo = map[string]string {
    "inner-value": `
    Here you can insert another longer description on a freshly created channel for this topic
    `,
}
*/


// Roles you want assigned to every support topic that will be created 
// var defaultCategoryRoles = []Role { roles["CommunityManager"], roles["ProjectManager"], }


/* Here you can assign specific roles for each topic 
var categoryRoles = map[string][]Role {
    "inner-value": {},
}
*/
