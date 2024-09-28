# golang-quotes
Just for fun

## Setup

sudo apt update && sudo apt upgrade

sudo apt install golang-go

go run .

## commands

### Gets:

curl http://localhost:8080/quotes/

curl http://localhost:8080/quotes/:author

### Posts:

curl http://localhost:8080/quotes \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "14","title": "any","author": "any"}'
