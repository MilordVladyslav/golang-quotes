# golang-quotes
Just for fun

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
