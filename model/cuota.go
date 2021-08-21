package model

import (
	"log"
        "fmt"
        "time"
	"database/sql"
	"errors"
       _ "github.com/lib/pq"   // PostgreSQL driver
	"gopkg.in/mgo.v2"
  )
// ---------------------------------------------------

var (
         Db *sql.DB
	databases share.Info
  )


// Periods table contains the information for each period
type Periodo struct {
	Id            uint32        `db:"id" bson:"id,omitempty"`
	Inicio        time.Time     `db:"inicio" bson:"inicio"`
	Final         time.Time     `db:"final" bson:"final"`
	CreatedAt     time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at" bson:"updated_at"`
}

type CuotaN struct {
	Id               uint32     `db:"id" bson:"id,omitempty"`
	PeriodId         uint32     `db:"periodid" bson:"periodid,omitempty"`
        Period        time.Time     `db:"period" bson:"period"`
	ApartaId         uint32     `db:"apartaid" bson:"apartaid,omitempty"`
	Apto             string     `db:"acodigo" bson:"acodigo,omitempty"`
	TipoId           uint32     `db:"tipoid" bson:"tipoid,omitempty"`
	Tipo             string     `db:"tdescripcion" bson:"tdescripcion,omitempty"`
        Fecha         time.Time     `db:"fecha" bson:"fecha"`
        Amount           int64     `db:"amount" bson:"amount"`
	CreatedAt     time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at" bson:"updated_at"`
}

      const(
	      layout      = "2006-01-02"
              timeLayout = "15:04:05"
            )

var (
	// ErrCode is a config or an internal error
	ErrCode = errors.New("Sentencia Case en codigo no es correcta.")
	// ErrNoResult is a not results error
	ErrNoResult = errors.New("Result  no encontrado.")
	// ErrUnavailable is a database not available error
	ErrUnavailable = errors.New("Database no disponible.")
	// ErrUnauthorized is a permissions violation
	ErrUnauthorized = errors.New("Usuario sin permiso para realizar esta operacion.")
        // Postgresql wrapper
         Db *sql.DB
	databases Info

)


// standardizeErrors returns the same error regardless of the database used
func standardizeError(err error) error {
	if err == sql.ErrNoRows || err == mgo.ErrNotFound {
		return ErrNoResult
	}

	return err
}

// Connect to the database
func Connect(d Info) {
	var err error
	// Store the config
	databases = d
//     fmt.Println(PgDNS(d))
          if Db,err  = sql.Open("postgres", PgDNS(d)); err !=  nil{
			log.Println("SQL Driver Error", err)
                        log.Fatal("Connection to database error")
           }
		if err = Db.Ping(); err != nil {
			log.Println("Database Error", err)
		}
}


  func PgDNS(ci Info  ) string {
         return   fmt.Sprintf("user=%s dbname=%s port=%d host=localhost sslmode=%s",ci.Username, ci.Name, ci.Port, ci.Parameter)
     }

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
// Get all periods in the database and returns the list
  func Periods() (periods []Periodo, err error) {
        var period Periodo
        stq :=   "SELECT id,  inicio, final, created_at, updated_at FROM periods order by  inicio desc"
        rows, err := Db.Query(stq)
        if err != nil {

                return
        }
        defer rows.Close()
        for rows.Next() {
                if err = rows.Scan(&period.Id, &period.Inicio, &period.Final, &period.CreatedAt, &period.UpdatedAt); err != nil {
                return
                }
                periods = append(periods, period)
        }
        return
 }

// PeriodByCode tenemos el period dado inicio
func (period * Periodo)PeriodByCode() (err error) {
        stq  :=   "SELECT id, inicio, final, created_at, updated_at FROM periods WHERE inicio = $1"
	err = Db.QueryRow(stq, &period.Inicio).Scan(&period.Id, &period.Inicio, &period.Final, &period.CreatedAt, &period.UpdatedAt)

	return  standardizeError(err)
}

// Get cuotas from a period 
  func CuotLim(inicio time.Time ) (cuotas []CuotaN, err error) {
        stq :=  " SELECT c.id, c.period_id, p.inicio, c.aparta_id, a.codigo, c.tipo_id, t.descripcion, " +
	        " c.fecha, c.amount, c.created_at, c.updated_at FROM cuotas c " +
		" JOIN  periods p ON c.Period_id = p.id JOIN  apartas a ON c.aparta_id = a.id " +
		" JOIN  tipos t ON c.tipo_id = t.id  WHERE   p.inicio = $1 ORDER BY p.inicio, c.fecha" 

	rows, err := Db.Query(stq, inicio)
	if err != nil {
		log.Println(err)
            return
	}
        defer rows.Close()
        cuot := CuotaN{}
        for rows.Next() {
           if err = rows.Scan(&cuot.Id,&cuot.PeriodId,&cuot.Period, &cuot.ApartaId, &cuot.Apto, &cuot.TipoId, &cuot.Tipo, &cuot.Fecha, &cuot.Amount,  &cuot.CreatedAt, &cuot.UpdatedAt); err != nil {
		log.Println(err)
                 return
            }
           cuotas = append(cuotas, cuot)
         }
       return
 }
