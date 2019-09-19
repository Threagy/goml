package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/goml/gomlet/db"
	"github.com/google/wire"
	"gopkg.in/yaml.v2"
)

type config struct {
	DB struct {
		URI    string `yaml:"uri"`
		DBName string `yaml:"dbname"`
	} `yaml:"db"`
}

func (c *config) getConfig() *config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {

	srv, cleanup, err := setup()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	addr := ":7777"
	log.Fatal(srv.ListenAndServe(addr))
}

func initialize() error {
	// load config
	var c config
	c.getConfig()

	db.Connect(c.DB.URI, c.DB.DBName)

	return nil
}

var applicationSet = wire.NewSet(
	newApplication,
	newRouter,
	wire.Bind(new(http.Handler), new(*http.ServeMux)),
)

type application struct {
}

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte("hello")); err != nil {
		log.Println("writing response:", err)
	}
}

func newApplication() *application {
	initialize()

	return &application{}
}

func newRouter(app *application) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", app.index)
	return r
}
