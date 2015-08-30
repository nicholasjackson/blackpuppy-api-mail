package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Message_Validator_With_Valid_Request(t *testing.T) {
	request := ContactUsRequest{}
	request.Body = "Something"
	request.Email = "nic@thatlondon.com"
	request.Name = "Somethings"

	response := ValidateRequest(&request)

	assert.Equal(t, true, response)
}

func Test_Message_Validator_With_Empty_Request(t *testing.T) {
	request := ContactUsRequest{}

	response := ValidateRequest(&request)

	assert.Equal(t, false, response)
}

func Test_Message_Validator_With_MissingName(t *testing.T) {
	request := ContactUsRequest{}
	request.Body = "Something"
	request.Email = "nic@thatlondon.com"

	response := ValidateRequest(&request)

	assert.Equal(t, false, response)
}

func Test_Message_Validator_With_Missing_Email(t *testing.T) {
	request := ContactUsRequest{}
	request.Body = "Something"
	request.Name = "Somethings"

	response := ValidateRequest(&request)

	assert.Equal(t, false, response)
}

func Test_Message_Validator_With_Missing_Body(t *testing.T) {
	request := ContactUsRequest{}
	request.Email = "nic@thatlondon.com"
	request.Name = "Somethings"

	response := ValidateRequest(&request)

	assert.Equal(t, false, response)
}
