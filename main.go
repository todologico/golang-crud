package main

import (
	"fmt"
	"log"
	"net/http"
)

//var db *sql.DB

//--------------------------------------
// home
//--------------------------------------
func home(w http.ResponseWriter, r *http.Request) {  
    
    fmt.Fprintf(w, "home url")
    
}

//--------------------------------------
// services
//--------------------------------------
func services(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "services url")
}


//--------------------------------------
// products
//--------------------------------------
func products(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "products url!\n")        

    numeros := make([]int, 5)

    numeros = append(numeros, 6)
    
    fmt.Println(numeros)   

}

//--------------------------------------
// main method routing
//--------------------------------------
func main() {

    /*

        // Capture connection properties.
        cfg := mysql.Config{
            User:   "golang",
            Passwd: "00000000",
            Net:    "tcp",
            Addr:   "127.0.0.1:3310",
            DBName: "golang",
            AllowNativePasswords: true,
        }
        // Get a database handle.
        var err error
        db, err = sql.Open("mysql", cfg.FormatDSN())
        if err != nil {
            log.Fatal(err)
        }
    
        pingErr := db.Ping()
        if pingErr != nil {
            log.Fatal(pingErr)
        }

        fmt.Println("Connected to mysql!")

    fmt.Println("Connected ya!")

    */

    http.HandleFunc("/", home)

    http.HandleFunc("/services", services)

    http.HandleFunc("/products", products)

    log.Fatal(http.ListenAndServe(":8080", nil))
}