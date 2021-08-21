package main

import (
        "fmt"
        "time"
        "encoding/json"
        "io/ioutil"
         "errors"
	 "github.com/gatopardo/test/ch14/rptcondo/control"
	 "github.com/gatopardo/test/ch14/rptcondo/model"
	 "github.com/gatopardo/test/ch14/rptcondo/share"
  )
// ---------------------------------------------------

      const(
	      layout      = "2006-01-02"
              timeLayout = "15:04:05"
            )

func main() {
    var inf Info
    data, err := ioutil.ReadFile("./config.json")
    if err != nil {
        fmt.Print(err)
    }
     err = json.Unmarshal(data, &inf)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(inf)
    model.Connect(inf )
    defer model.Db.Close()
    stInic := "2021-07-01"
    dInic,_ := time.Parse(layout,stInic)
    control.CuotList(dInic)

}

