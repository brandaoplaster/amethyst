package router

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func New() *Router {
	return &Router{mux: mux.NewRouter()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) Resources(resource string, controller interface{}) {
	r.mux.HandleFunc("/"+resource, callMethod(controller, "Index")).Methods("GET")
	r.mux.HandleFunc("/"+resource, callMethod(controller, "Create")).Methods("POST")
	r.mux.HandleFunc("/"+resource+"/{id}", callMethod(controller, "Show")).Methods("GET")
	r.mux.HandleFunc("/"+resource+"/{id}", callMethod(controller, "Update")).Methods("PUT")
	r.mux.HandleFunc("/"+resource+"/{id}", callMethod(controller, "Delete")).Methods("DELETE")
}

func (r *Router) CustomAction(resource string, action string, controller interface{}, method string) {
	r.mux.HandleFunc("/"+resource+"/{id}/"+action, callMethod(controller, action)).Methods(method)
}

func callMethod(controller interface{}, methodName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controllerMethod := reflect.ValueOf(controller).MethodByName(methodName)
		if !controllerMethod.IsValid() {
			http.Error(w, "Método não encontrado", http.StatusNotFound)
			return
		}
		controllerMethod.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
	}
}
