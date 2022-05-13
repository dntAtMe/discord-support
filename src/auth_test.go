package main

import (
    "testing"
)

func TestReadAuth(t *testing.T) {

    // If any function panics, catch it
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("Reading authentication paniced")
        }
    }()

    var auth Auth
    ReadAuth(&auth, "example_auth.json")

    if auth.Token != "token" {
        t.Log("Auth token should be \"token\" but got", auth.Token)
        t.Fail()
    }
    if auth.GuildID != "guildId" {
        t.Log("Auth token should be \"guildId\" but got", auth.GuildID)
        t.Fail()
    }
    if auth.AppID != "appId" {
        t.Log("Auth token should be \"appId\" but got", auth.AppID)
        t.Fail()
    }

}
