package route

import (
    "github.com/gorilla/mux"
    hubspot "github.com/heaptracetechnology/hubspot/hubspot"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateOrUpdateContact",
        "POST",
        "/createOrUpdate",
        hubspot.CreateOrUpdateContact,
    },
    Route{
        "GetContactByVID",
        "POST",
        "/getContact",
        hubspot.GetContactByVID,
    },
    Route{
        "DeleteContactByVID",
        "POST",
        "/deleteContact",
        hubspot.DeleteContactByVID,
    },
    Route{
        "CreateDeal",
        "POST",
        "/createDeal",
        hubspot.CreateDeal,
    },
    Route{
        "CreateTicket",
        "POST",
        "/createTicket",
        hubspot.CreateTicket,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
