# reef-pi/client

An http client library for [reef-pi](http://reef-pi.com)

[![GoDoc](https://godoc.org/github.com/reef-pi/client?status.svg)](https://godoc.org/github.com/reef-pi/client)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/reef-pi/client/blob/master/LICENSE.txt)

## Example

```go
  c, err := New("http://reef-pi.local")
  if err != nil {
    panic(err)
  }
  if err := c.SignIn("reef-pi","reef-pi"); err != nil {
    panic(err)
  }
  defer c.SignOut()
  outlets, err := c.Outlets()
  if err != nil {
    panic(err)
  }
  for _, outlet := range outlets {
     fmt.Println(outlet.Name)
  }
```
