package main

type Role struct {
	ID string
}

var roles = map[string]Role{
	"CommunityManager": {ID: "434331919895756800"},
	"ProjectManager":   {ID: "434330942342168576"},
	"ServerAdmin":      {ID: "434334962070716418"},
	"Businessman":      {ID: "856510661076451369"},
	"CrimeManager":     {ID: "1035574197767909458"},
}
