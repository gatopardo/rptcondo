package share

import (
        "fmt"
  )
// ---------------------------------------------------

type Info struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

func commas(s string) string {
    lon := len(s)
    if lon <= 3 {
        return s
    } else {
        return commas(s[0:len(s)-3]) + "," + s[len(s)-3:]
    }
}

func Format64(dat int64)(str string) {
          str  = fmt.Sprintf("%d", dat)
          lon  := len(str)
	  if  dat < 0 {
	          str = str[1:lon]
	  }
          lon  = len(str)  - 2
	  if lon > 0 {
              sini := str[:lon]
              sfin := str[lon:]
              scom := commas(sini)
              str  =  scom + "." + sfin
          }else{
	      pre := "0."
              if lon < 0 {
                  pre = "0.0"
              }
               str = pre + str
	  }
	  if  dat < 0 {
	       str = "-"+str
	  }
          return
     }
