package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(), // Remove from production
)

func Home(w http.ResponseWriter, r *http.Request) {
	
}
