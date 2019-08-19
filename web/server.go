package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/janithl/kottu2020/domain/blog"
	"github.com/janithl/kottu2020/domain/post"
)

// Server struct implements the REST API Server
type Server struct {
	port        int
	blogService blog.Service
	postService post.Service
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
			output, err = s.postService.FindPost(id)
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
		lang, ok := vars["language"]
		if !ok {
			lang = "en"
		}
		output := s.postService.FindLatestPosts(lang, page)
		s.outputJSON(w, output)
	}
}

// listBlogs returns listings of blogs
func (s *Server) listBlogs(criteria string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var output []*blog.Blog
		if criteria == "all" {
			page, err := strconv.Atoi(vars["page"])
			if err != nil {
				page = 1
			}
			output = s.blogService.FindAllBlogs(page)
		} else {
			output = s.blogService.FindPopularBlogs()
		}
		s.outputJSON(w, output)
	}
}

// getBlogCount returns the count of blogs
func (s *Server) getBlogCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count := s.blogService.BlogCount()
		s.outputJSON(w, map[string]int{"count": count})
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
	apirouter.HandleFunc("/latest/{language:[a-z]{2}}", s.latestPosts())
	apirouter.HandleFunc("/latest/{language:[a-z]{2}}/page/{page:[0-9]+}", s.latestPosts())

	apirouter.HandleFunc("/blogs/all", s.listBlogs("all"))
	apirouter.HandleFunc("/blogs/all/page/{page:[0-9]+}", s.listBlogs("all"))
	apirouter.HandleFunc("/blogs/popular", s.listBlogs("popular"))
	apirouter.HandleFunc("/blogs/count", s.getBlogCount())

	// static file serving
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(s.staticPath))))

	// serve on given port
	fmt.Printf("Serving HTTP on localhost: %d\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), r))
}

// NewServer returns a new instance of the server
func NewServer(port int, blogService blog.Service, postService post.Service) *Server {
	return &Server{
		port:        port,
		blogService: blogService,
		postService: postService,
		staticPath:  "static",
	}
}
