package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proyecto/graph"
	"regexp"
	"time"
	"unicode"

	"strings"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/pocketbase/pocketbase-go"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

const dbConection = "host=localhost port=5432 dbname=CRUD user=postgres password=postgresql"

type Data struct {
    ID   int    `gorm:"primaryKey"`
    Text string `gorm:"column:text"`
}


func Parrafo2() {
	p := "En las vacaciones de este año primero fui diez días a un hotel con mi familia, en el hotel de al lado estaba un amigo y todos los días salíamos juntos a la" + "piscina, a la playa, etc. En el hotel había un bufete donde íbamos a desayunar, almorzar y cenar, por la noche paseábamos por el paseo marítimo mientras nos tomábamos" + "un helado, algunos días íbamos a otras playas que estaban más lejos del hotel, algunas eran salvajes, tenían el agua cristalina y había muchos peces y plantas, otros" + " días los pasábamos en la piscina que era muy grande."

	var mapa = make(map[string]int)
	start := time.Now()

	palabras := strings.FieldsFunc(p, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, palabra := range palabras {
		d := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return unicode.ToLower(r)
			}
			return -1
		}, palabra)

		if d != "" {
			mapa[d] = mapa[d] + 1
		}
	}

	for clave, valor := range mapa {
		log.Println(clave, ":", valor)
	}

	timeLapse := time.Since(start)
	log.Println(timeLapse)

}

func Parrafo() {
	p := "En las vacaciones de este año primero fui diez días a un hotel con mi familia, en el hotel de al lado estaba un amigo y todos los días salíamos juntos a la" + "piscina, a la playa, etc. En el hotel había un bufete donde íbamos a desayunar, almorzar y cenar, por la noche paseábamos por el paseo marítimo mientras nos tomábamos" + "un helado, algunos días íbamos a otras playas que estaban más lejos del hotel, algunas eran salvajes, tenían el agua cristalina y había muchos peces y plantas, otros" + " días los pasábamos en la piscina que era muy grande."
	var mapa = make(map[string]int)
	start := time.Now()
	palabras := strings.Fields(p)
	slice := []string{}

	for _, palabra := range palabras {
		d := strings.ToLower(palabra)

		slice = append(slice, d)

	}
	r, _ := regexp.Compile("[^\\w\\s]+")

	for i, palabra := range slice {
		slice[i] = r.ReplaceAllString(palabra, "")
	}

	for _, palabra := range slice {
		if value, ok := mapa[palabra]; ok {
			mapa[palabra] = value + 1
		} else {
			mapa[palabra] = 1
		}

	}

	for clave, valor := range mapa {
		log.Println(clave, ":", valor)
	}

	timeLapse := time.Since(start)

	log.Println(timeLapse)

}
func main() {
    Parrafo()
	Parrafo2()
   
  
}






func Read() ([]string, error) {
    file, err := os.Open("C:/Users/Luis Acuna/Desktop/mis textos/prueba.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    fileInfo, err := file.Stat()
    if err != nil {
        return nil, err
    }

    fileSize := fileInfo.Size()
    buffer := make([]byte, fileSize)

    _, err = file.Read(buffer)
    if err != nil {
        return nil, err
    }

    content := string(buffer)
    lines := strings.Split(content, "\n")

    var result []string
    result = append(result, lines...)

	for _, line := range result {
        fmt.Println(line)
    }

    
    return result, nil
}



func Server() {
    dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            panic("error al conectar a la base de datos: " + err.Error())
        }else{
            log.Println("conexion con postgres existosa con gorm")
        }
        var data []Data

    // Seleccionar todos los registros
    db.Find(&data)

    for _, d := range data {
        fmt.Println(d.ID, d.Text)
    }
        

        port := os.Getenv("PORT")
        if port == "" {
            port = defaultPort
        }

        srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
