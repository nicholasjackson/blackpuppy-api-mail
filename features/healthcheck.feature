@helloworld
Feature: HealthCheck
	In order to ensure quality
	As a user
	I want to be able to test functionality of my API

Scenario: Check API health
	Given I send a GET request to "/mail/healthcheck"
	Then the response status should be "200"
	And the JSON response should have "$..StatusMessage" with the text "Hello World"
