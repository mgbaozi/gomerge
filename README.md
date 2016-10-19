Gomerge
============
Merge a struct with a map[string]interface{} in golang. (according to json-tag)

#### Why
Think that when writing a web server, it have a model and data of a user (tom)
``` go
    type People struct {
        Name  string `json:"name"`
        Sex   string `json:"sex"`
        Age   int    `json:"age"`
        Times int   `json:"times"`
    }

    var tom := People{
        Name: "tom",
        Sex: "male",
        Age: 18,
        Times: 1,
    }
```
When a user only want to update 'age' field, he give server a JSON like
``` json
{
    "name": "tom",
    "age": 19
}
```
If you use json.Unmarshal and merge to struct, you will get
``` go
People{
    Name: "tom",
    Sex: "male",
    Age: 18,
    Times: 0,
}
```
So you need to unmarshal a struct to a map[string]interface, and merge with origin struct.

``` go
import (
    "json"
    "github.com/mgbaozi/gomerge"
)

// body as string
func main() {
    var request_data map[string]interface{}
    if err := json.Unmarshal(body, &request_data); err != nil {
        // something wrong
    }
    if err := gomerge.Merge(&tom, request_data); err != nil {
        // something wrong
    }
    result, _ := json.Marshal(tom)
    fmt.Println(result)
}
```
