package main

import (
  "errors"
  "log"
  "net/http"
  "encoding/json"
  "fmt"
)

//the following go code shall represent the json data structure:
//{"data": {"children": [
//    {"data": {
//      "title":  "The Go homepage",
//      "url":    "http://golang.org",
//      ...
//     }
//    },
//  ]}}

type Item struct {
  Title string
  URL   string
}

type Response struct {
  Data struct {
    Children []struct { //this is a slice, its like an array
      Data Item
    }
  }
}


func Get(reddit string) ([]Item, error) {
  url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
  resp, err := http.Get(url)

  if err != nil {
    return nil, err
  }

  //close the tcp connection. 'defer' makes sre that this statement is 
  // executed after the function returns, no matter what.
  
  defer resp.Body.Close() 

  if resp.StatusCode != http.StatusOK {
    return nil, errors.New(resp.Status)
  }

  r := new(Response)
  err = json.NewDecoder(resp.Body).Decode(r)

  if err != nil {
    return nil, err
  }

  items := make([]Item, len(r.Data.Children)) //create empty slice
  for i, child := range r.Data.Children {
    items[i] = child.Data
  }

  return items, nil
}

func main() {
  items, err := Get("golang")
  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item.Title)
  }
}
