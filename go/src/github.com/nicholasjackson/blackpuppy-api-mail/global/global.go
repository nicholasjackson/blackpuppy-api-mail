package global

import (
	"encoding/json"
	"os"
	"log"
)

type smtpServerSettings struct {
	Server   string
	User     string
	Password string
	UseAuth  bool
}

type ConfigStruct struct {
	ContactUsEmail     string
	ContactUsSubject   string
	ContactUsTemplate  string
	SmtpServerSettings smtpServerSettings
	RootFolder         string
}

var Config ConfigStruct

func LoadConfig(config string, rootfolder string) {
	file, err := os.Open(config)
	if err != nil {
		panic("Unable to open config")
	}

	decoder := json.NewDecoder(file)
	Config = ConfigStruct{}
	err = decoder.Decode(&Config)
	Config.RootFolder = rootfolder

	log.Println("Email:", Config.ContactUsEmail)
}
