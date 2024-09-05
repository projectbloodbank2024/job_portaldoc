/*
 * JOB_PORTAL
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"ApiV1ApplicationDelete66d6ac98c62ea8144431d04cDelete",
		strings.ToUpper("Delete"),
		"/api/v1/application/delete/66d6ac98c62ea8144431d04c",
		ApiV1ApplicationDelete66d6ac98c62ea8144431d04cDelete,
	},

	Route{
		"ApiV1ApplicationEmployerGetallGet",
		strings.ToUpper("Get"),
		"/api/v1/application/employer/getall",
		ApiV1ApplicationEmployerGetallGet,
	},

	Route{
		"ApiV1ApplicationJobseekerGetallGet",
		strings.ToUpper("Get"),
		"/api/v1/application/jobseeker/getall",
		ApiV1ApplicationJobseekerGetallGet,
	},

	Route{
		"ApiV1ApplicationPost66d5501fe452d7bffe026053Post",
		strings.ToUpper("Post"),
		"/api/v1/application/post/66d5501fe452d7bffe026053",
		ApiV1ApplicationPost66d5501fe452d7bffe026053Post,
	},

	Route{
		"ApiV1JobDelete66d54d94e452d7bffe026047Delete",
		strings.ToUpper("Delete"),
		"/api/v1/job/delete/66d54d94e452d7bffe026047",
		ApiV1JobDelete66d54d94e452d7bffe026047Delete,
	},

	Route{
		"ApiV1JobGet66d30863fbf23475bbcd92a1Get",
		strings.ToUpper("Get"),
		"/api/v1/job/get/66d30863fbf23475bbcd92a1",
		ApiV1JobGet66d30863fbf23475bbcd92a1Get,
	},

	Route{
		"ApiV1JobGetallGet",
		strings.ToUpper("Get"),
		"/api/v1/job/getall",
		ApiV1JobGetallGet,
	},

	Route{
		"ApiV1JobGetmyjobsGet",
		strings.ToUpper("Get"),
		"/api/v1/job/getmyjobs",
		ApiV1JobGetmyjobsGet,
	},

	Route{
		"ApiV1JobPostPost",
		strings.ToUpper("Post"),
		"/api/v1/job/post",
		ApiV1JobPostPost,
	},

	Route{
		"ApiV1UserGetuserGet",
		strings.ToUpper("Get"),
		"/api/v1/user/getuser",
		ApiV1UserGetuserGet,
	},

	Route{
		"ApiV1UserLoginPost",
		strings.ToUpper("Post"),
		"/api/v1/user/login",
		ApiV1UserLoginPost,
	},

	Route{
		"ApiV1UserLogoutGet",
		strings.ToUpper("Get"),
		"/api/v1/user/logout",
		ApiV1UserLogoutGet,
	},

	Route{
		"ApiV1UserRegisterPost",
		strings.ToUpper("Post"),
		"/api/v1/user/register",
		ApiV1UserRegisterPost,
	},

	Route{
		"ApiV1UserUpdatePasswordPut",
		strings.ToUpper("Put"),
		"/api/v1/user/update/password",
		ApiV1UserUpdatePasswordPut,
	},

	Route{
		"ApiV1UserUpdateProfilePut",
		strings.ToUpper("Put"),
		"/api/v1/user/update/profile",
		ApiV1UserUpdateProfilePut,
	},
}
