package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    cert, err := tls.LoadX509KeyPair("configs/spire-quickstart-svid.pem", "configs/spire-quickstart-key.pem")
    if err != nil {
        log.Fatal(err)
    }

    caCert, err := ioutil.ReadFile("configs/ca-cert.pem")
    if err != nil {
        log.Fatal(err)
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientCAs:    caCertPool,
        ClientAuth:   tls.RequireAndVerifyClientCert,
    }

    server := &http.Server{
        Addr:      ":443",
        TLSConfig: tlsConfig,
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from spire-quickstart!")
    })

    log.Fatal(server.ListenAndServeTLS("", ""))
}
