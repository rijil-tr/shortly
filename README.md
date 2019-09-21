# Shortly
**Shortly** is a URL shortner service

## Code example

```
# Execute the program
go run main.go
```

Using CURL

Generate a short URL
`curl 'http://localhost:5000/' -H 'Content-Type: application/x-www-form-urlencoded' --data 'link=https://google.com'`

Redirect
`curl http://localhost:5000/l/3E84B09B18848F91`

Visitor Count
`curl http://localhost:5000/s/3E84B09B18848F91`