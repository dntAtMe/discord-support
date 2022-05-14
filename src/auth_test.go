package main

import (
    "testing"
)

func TestReadAuth(t *testing.T) {

    targetAuth := Auth {
        Token: "token",
        GuildID: "guildId",
        AppID: "appId",
    }

    // If any function panics, catch it
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("Reading authentication paniced")
        }
    }()

    var auth Auth
    readAuth(&auth, "example_auth.json")

    if auth.Token != targetAuth.Token {
        t.Logf("Auth token should be %s but got %s", targetAuth.Token, auth.Token)
        t.Fail()
    }
    if auth.GuildID != targetAuth.GuildID {
        t.Logf("Auth token should be %s but got %s", targetAuth.GuildID, auth.GuildID)
        t.Fail()
    }
    if auth.AppID != targetAuth.AppID {
        t.Logf("Auth token should be %s but got %s", targetAuth.AppID, auth.AppID)
        t.Fail()
    }

}
