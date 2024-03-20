package main

import(
	"net/http"
	"log"

) 

fund main(){
	http.HandleFunc("/", healthHandler)

	log.Println("server starting at :444")

	err := http.ListenAndServer(":444", nil)
	if err != nil{
		log.Fatal("server error", err)
	}
	log.Println("server is stopped")

}

func healthHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server is running..."))
}