package business

type ContactUsRequest struct {
	Name  string
	Email string
	Body  string
}

type ContactUsResponse struct {
	StatusMessage string
}

func ValidateRequest(request *ContactUsRequest) bool {

	if request.Name != "" && request.Email != "" && request.Body != "" {
		return true
	}

	return false
}
