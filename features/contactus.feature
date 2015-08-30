@helloworld
Feature: HealthCheck
	In order to ensure quality
	As a user
	I want to be able to test functionality of my API

Scenario: Send a contact us email with a valid message
  Given I send a POST request to "/mail/contactus" with the following:
    | Name   | Nic Jackson |
    | Email | nic@thatlondon.com  |
    | Body | This is the message body  |
  Then the response status should be "200"
  And I should receive 1 email with a html body containing "Nic Jackson"
  And I should receive 1 email with a html body containing "nic@thatlondon.com"
  And I should receive 1 email with a html body containing "This is the message body"

Scenario: Send a contact us email with an invalid message
  Given I send a POST request to "/mail/contactus" with the following:
    | name   | Nic Jackson |
  Then the response status should be "400"

Scenario: Send a contact us email with an no message
	Given I send a POST request to "/mail/contactus"
	Then the response status should be "400"
