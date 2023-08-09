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
