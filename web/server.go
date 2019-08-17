package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/janithl/kottu2020/domain/blog"
)

// Server struct implements the REST API Server
type Server struct {
	port        int
	blogService blog.Service
	staticPath  string
}

// outputJSON is a helper that converts the output object to JSON
func (s *Server) outputJSON(w http.ResponseWriter, output interface{}) {
	outputJSON, err := json.Marshal(output)
	if s.outputErrorJSON(w, err, http.StatusInternalServerError) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}

// outputErrorJSON is a helper that checks if here's an error, and if so outputs JSON
func (s *Server) outputErrorJSON(w http.ResponseWriter, err error, errorCode int) bool {
	if err != nil {
		json, _ := json.Marshal(ServerError{
			ErrorCode:    errorCode,
			ErrorMessage: err.Error(),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorCode)
		w.Write(json)
		return true
	}
	return false
}

// listDetails, given the ID and object type, lists the object's information
func (s *Server) listDetails(objtype string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if s.outputErrorJSON(w, err, http.StatusBadRequest) {
			return
		}

		var output interface{}
		if objtype == "blog" {
			output, err = s.blogService.FindBlog(id)
			if s.outputErrorJSON(w, err, http.StatusNotFound) {
				return
			}

		} else if objtype == "post" {
			output, err = s.blogService.FindPost(id)
			if s.outputErrorJSON(w, err, http.StatusNotFound) {
				return
			}
		}
		s.outputJSON(w, output)
	}
}

// latestPosts returns the newest posts
func (s *Server) latestPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			page = 1
		}
		output := s.blogService.FindLatestPosts("en", 20, page)
		s.outputJSON(w, output)
	}
}

// Serve serves HTTP
func (s *Server) Serve() {
	r := mux.NewRouter()

	// /api route
	apirouter := r.PathPrefix("/api").Subrouter()
	apirouter.HandleFunc("/blog/{id:[0-9]+}", s.listDetails("blog"))
	apirouter.HandleFunc("/post/{id:[0-9]+}", s.listDetails("post"))
	apirouter.HandleFunc("/latest", s.latestPosts())
	apirouter.HandleFunc("/latest/page/{page:[0-9]+}", s.latestPosts())

	// static file serving
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(s.staticPath))))

	// serve on given port
	fmt.Printf("Serving HTTP on localhost: %d\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), r))
}

// NewServer returns a new instance of the server
func NewServer(port int, blogService blog.Service) *Server {
	return &Server{
		port:        port,
		blogService: blogService,
		staticPath:  "static",
	}
}
