package main

import (
	"crypto/tls"
	"encoding/json"
	dataBase "forum/back/create_db"
	server "forum/back/server"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Config contient les informations de configuration
type Config struct {
	CertFilePath string `json:"CertFilePath"`
	KeyFilePath  string `json:"KeyFilePath"`
}

func main() {

	dataBase.InitiateDatabase()

	// load css
	css := http.FileServer(http.Dir("./front/static/css"))
	http.Handle("/css/", http.StripPrefix("/css", css))
	// load images
	img := http.FileServer(http.Dir("./front/static/img"))
	http.Handle("/img/", http.StripPrefix("/img", img))
	// load scripts
	scpt := http.FileServer(http.Dir("./front/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts", scpt))

	// Configuration du serveur TLS
	tlsConfig, err := setupTLSConfig("config.json")
	if err != nil {
		log.Fatalf("Error setting up TLS configuration: %v", err)
	}

	server.Handlers()

	// Démarrage du serveur TLS
	startTLSServer(":8080", tlsConfig)
}

// setupTLSConfig configure la configuration TLS en chargeant les informations depuis un fichier de configuration
func setupTLSConfig(configFile string) (*tls.Config, error) {
	config, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}
	serverTLSCert, err := tls.LoadX509KeyPair(config.CertFilePath, config.KeyFilePath)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}, nil
}

// startTLSServer démarre le serveur TLS
func startTLSServer(addr string, tlsConfig *tls.Config) {
	server := http.Server{
		Addr:         addr,
		TLSConfig:    tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}

// loadConfig charge la configuration depuis un fichier JSON
func loadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
