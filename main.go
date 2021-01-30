package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello!")

		// !!!!test code
		// Gets and prints SQL Server version
		// Use background context
		//db, err := sql.Open("sqlserver", ConnString())
		//checkErr(err)
		//defer db.Close()
		//
		//result:=0.0
		//err = db.QueryRow("select sum(loss/10000) from CS_AnnualLosses").Scan(&result)
		//if err != nil {
		//
		//	fmt.Fprintf(w, "server failure: %s", err.Error())
		//} else {
		//	fmt.Fprintf(w, "%f", result)
		//}

		//		fmt.Printf("%f\n", result)

	} else {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}

