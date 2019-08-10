package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/janithl/kottu2020/domain/blog"
)

// Server struct implements the REST API Server
type Server struct {
	port        string
	blogService blog.Service
}

// ServerError struct is used to convey API errors
type ServerError struct {
	ErrorCode    int
	ErrorMessage string
}

// defaultHandler serves out the index.html file
func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "index.html")
	http.ServeFile(w, r, fp)
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
func (s *Server) listDetails(objType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := "/" + objType + "/"
		id, err := strconv.Atoi(r.URL.Path[len(slug):])
		if s.outputErrorJSON(w, err, http.StatusBadRequest) {
			return
		}

		if objType == "blog" {
			blog, err := s.blogService.FindBlog(id)
			if s.outputErrorJSON(w, err, http.StatusNotFound) {
				return
			}

			s.outputJSON(w, blog)
		} else if objType == "post" {
			post, err := s.blogService.FindPost(id)
			if s.outputErrorJSON(w, err, http.StatusNotFound) {
				return
			}

			defer s.outputJSON(w, post)
		}
	}
}

// Serve serves HTTP
func (s *Server) Serve() {
	// define the routes
	http.HandleFunc("/blog/", s.listDetails("blog"))
	http.HandleFunc("/post/", s.listDetails("post"))
	http.HandleFunc("/", s.defaultHandler)

	// serve on given port
	fmt.Println("Serving HTTP on localhost:" + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

// NewServer returns a new instance of the server
func NewServer(port string, blogService blog.Service) *Server {
	return &Server{
		port:        port,
		blogService: blogService,
	}
}
