package control

import (
        "fmt"
	"model"
	"share"
  )
// ---------------------------------------------------


func CuotList(inicio time.Time) {
        lisCuot, err         := model.CuotLim(inicio)
        if err != nil {
            log.Println(err)
         }
	 for  i := 0; i < len(lisCuot); i++{
		 fec    :=  lisCuot[i].Fecha.Format("2006-01-02")
		 monto  :=  share.Format64(lisCuot[i].Amount)
		 apt    :=  lisCuot[i].Apto
		 fmt.Printf("%s  %s   %s\n", apt, fec, monto )
	 }
 }

