package router

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	// We could use http.Handler as a type here; using the specific type has
	// the advantage that static analysis tools can link directly from
	// h.UserHandler.ServeHTTP to the correct definition. The disadvantage is
	// that we have slightly stronger coupling. Do the tradeoff yourself.
	UserHandler *UserHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "user" {
		h.UserHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

type UserHandler struct {
}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}
	switch req.Method {
	case "GET":
		h.handleGet(id)
	case "PUT":
		h.handlePut(id)
	default:
		http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}

// func main() {
//     a := &App{
//         UserHandler: new(UserHandler),
//     }
//     http.ListenAndServe(":8000", a)
// }

type UserHandler struct {
	ProfileHandler *ProfileHandler
}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}

	if req.URL.Path != "/" {
		head, tail := ShiftPath(req.URL.Path)
		switch head {
		case "profile":
			// We can't just make ProfileHandler an http.Handler; it needs the
			// user id. Let's instead…
			h.ProfileHandler.Handler(id).ServeHTTP(res, req)
		case "account":
			h.AccountHandler.Handler(accountId).serveHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
		return
	}
	// As before
	// ...
}

type ProfileHandler struct {
}

func (h *ProfileHandler) Handler(id int) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Do whatever
	})
}
