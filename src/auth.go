package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// TODO: Write unit tests for reading auth data

type Auth struct {
    Token string `json:"token"`
    AppID string `json:"appId"`
    GuildID string `json:"guildId"`
}

var configPath = "config/"

func ReadAuth(auth *Auth) {
    authPath := fmt.Sprintf("%sauth.json", configPath)
    jsonFile, err := os.Open(authPath)
    
    if err != nil {
        panic(err)
    }

    // If success, close the file on function return
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValue, auth)
}
