package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling incoming orders..✅")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling users..✅")
	})

	PORT := 3000

	// Load the TLS cert and key
	cert:= "cert.pem"
	key:= "key.pem"

	// Configure TLS
	tlsConfig:= &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server:= http.Server{
		Addr: fmt.Sprintf(":%d",PORT),
		Handler: nil,
		TLSConfig: tlsConfig,
	}


	// Enable http2
	//! 🟢 HTTP2 Server with TLS
	http2.ConfigureServer(&server,&http2.Server{})

	fmt.Println("🟢 Server is running on PORT:",PORT)

	err:=server.ListenAndServeTLS(cert,key)
	if err!=nil{
		log.Fatal("⚠️ Could not start the server:",err)
	}



	//! 🟢 HTTP 1.1 Server without TLS
	// err:=http.ListenAndServe(fmt.Sprintf(":%d",PORT),nil)
	// if err!=nil{
	// 	log.Fatal("⚠️ Could not start the server:",err)
	// }

}