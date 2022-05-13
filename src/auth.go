package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Auth struct {
    Token string `json:"token"`
    AppID string `json:"appId"`
    GuildID string `json:"guildId"`
}

var configPath = "../config/"

func ReadAuth(auth *Auth, fileName string) {

    authPath := fmt.Sprintf("%s%s", configPath, fileName)
    jsonFile, err := os.Open(authPath)
    
    if err != nil {
        panic(err)
    }

    // If success, close the file on function return
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValue, auth)
}
