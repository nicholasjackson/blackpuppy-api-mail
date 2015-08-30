require 'cucumber/rest_api'
require 'cucumber/pickle_mongodb'
require 'cucumber/mailcatcher'

$SERVER_PATH = ENV['WEB_SERVER_URI']
$EMAIL_SERVER_PATH = ENV['EMAIL_SERVER_URI']
Cucumber::Mailcatcher::HttpClient.server_url = $EMAIL_SERVER_PATH
