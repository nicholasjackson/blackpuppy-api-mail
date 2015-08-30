package business

import (
	"net/smtp"
	"testing"
	//"github.com/nicholasjackson/blackpuppy-api-mail/global"
	"github.com/nicholasjackson/blackpuppy-api-mail/email"
	"github.com/stretchr/testify/mock"
)

type MockSmtp struct {
	mock.Mock
	email.BaseMail
}

func (m *MockSmtp) SendMail(addr string, a smtp.Auth, from string, to []string, subject string, msg []byte) error {
	args := m.Mock.Called()
	return args.Error(0)
}

func (m *MockSmtp) BuildMessage(subject string, template string, data interface{}) []byte {
	args := m.Mock.Called()
	return args.Get(0).([]byte)
}

func Test_Send_Email(t *testing.T) {
	var mockSmtp *MockSmtp = new(MockSmtp)
	message := []byte("This is the message")

	mockSmtp.Mock.On("BuildMessage").Return(message)
	mockSmtp.Mock.On("SendMail").Return(nil)

	_ = SendEmail("test@test.com", "subject", "body", mockSmtp)

	mockSmtp.Mock.AssertCalled(t, "BuildMessage", mock.Anything, mock.Anything, mock.Anything)

	mockSmtp.Mock.AssertNumberOfCalls(t, "BuildMessage", 1)
	mockSmtp.Mock.AssertNumberOfCalls(t, "SendMail", 1)
}
