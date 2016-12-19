package main

import (
	"flag"
	"log"
	"net/http"

  "github.com/julienschmidt/httprouter"
)

//IndexHandler returns an httprouter.Handle which serves entrypoint
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        http.ServeFile(w, r, entrypoint)
    }

    return httprouter.Handle(fn)
}

func main() {
  entry := "./landing.html"
	var addr = flag.String("addr", ":8080", "The addr of the app.")
	flag.Parse() //parse flags

  r := httprouter.New()

  r.ServeFiles("/static/*filepath", http.Dir("static"))
  r.GET("/", IndexHandler(entry))

	log.Fatal(http.ListenAndServe(*addr, r))
}
