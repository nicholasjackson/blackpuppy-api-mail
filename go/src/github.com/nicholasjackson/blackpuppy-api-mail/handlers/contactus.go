package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/nicholasjackson/blackpuppy-api-mail/business"
	"github.com/nicholasjackson/blackpuppy-api-mail/email"
)

func ContactUsHandler(rw http.ResponseWriter, r *http.Request) {
	response := business.ContactUsResponse{}
	response.StatusMessage = "Invalid request object"

	encoder := json.NewEncoder(rw)

	data, _ := ioutil.ReadAll(r.Body)
	var contactUsRequest business.ContactUsRequest

	err := json.Unmarshal(data, &contactUsRequest)
	if err != nil || !business.ValidateRequest(&contactUsRequest) {
		http.Error(rw, "Invalid request object", http.StatusBadRequest)
		return
	}

	mail := email.SmtpMail{}

	err = business.SendEmail(contactUsRequest.Name, contactUsRequest.Email, contactUsRequest.Body, &mail)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder.Encode(&response)
}
