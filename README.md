# Go & Gin Tutotial
1. Install GIN
```bash
go get -u github.com/gin-gonic/gin
```
2. Import Gin library
```go
import "github.com/gin-gonic/gin"
```
## Serialization
To serialize in go we use the following:
```go
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
```
`json:"id" is for serialization` so to not use the capitalized prop when passing data through a request

## Controller
To create an endpoint we gotta use the ```gin.Context``` which is the most important part of Gin for it carries the request details, validates and serializes JSON.
Calling ```Context.IndentedJSON``` to serialize the struct into JSON and add it to the response:
```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
```
So we define the variable that is going to use the gin.context

## Main function
First we init the router using ```gin.Default()```
```go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
```

## Blanc identifyer ```_```
So when looping through an array we will have something like this:
```go
animals := []string{"dog","fish","cow"}
for i, animal := range animals{
    log.Println(i, animal)
}
```
and the output would be something like this:
```bash
0 dog
1 fish
2 cow
```
What if I dont care about the index?
Then we use the 'blanc identifyer' to tell go compiler that we dont care about the fist value being returned and we dont have an error
```go
animals := []string{"dog","fish","cow"}
for _, animal := range animals{
    log.Println(animal)
}
```
and the output would be something like this:
```bash
    dog
    fish
    cow
```
## Running the app
in the command line type the following:
```bash
go run .
```
