package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nicholasjackson/blackpuppy-api-mail/business"
	"github.com/nicholasjackson/blackpuppy-api-mail/email"
)

func ContactUsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	response := business.ContactUsResponse{}
	response.StatusMessage = "OK"

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
