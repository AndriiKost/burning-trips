package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andriikost/burning-gateway/api/models"
	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/gorilla/mux"
)

type FeaturedStops struct {
	NearestStops   []models.Stop  `json:"nearestStops"`
	TopRankedStops []models.Stop  `json:"topRankedStops"`
	Stories        []models.Story `json:"stories"`
	Routes         []models.Route `json:"routes"`
	Stop           models.Stop    `json:"stop"`
}

func (server *Server) GetFeaturedStops(w http.ResponseWriter, r *http.Request) {
	var err error

	vars := mux.Vars(r)
	stopId, err := strconv.ParseUint(vars["id"], 10, 64)

	vm := FeaturedStops{}

	stopRepo := models.Stop{}
	routeRepo := models.Route{}
	storyRepo := models.Story{}

	curStop, err := stopRepo.Find(server.DB, stopId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		vm.Stop = *curStop
	}

	nearestStops, err := stopRepo.GetNearestStops(server.DB, stopId)
	if err != nil {
		fmt.Println("stopRepo GetNearestStops ", err)
	} else {
		vm.NearestStops = nearestStops
	}

	topRankedStops, err := stopRepo.GetTopRankedStops(server.DB, stopId)
	if err != nil {
		fmt.Println("stopRepo GetTopRankedStops ", err)
	} else {
		vm.TopRankedStops = topRankedStops
	}

	routes, err := routeRepo.FindAllByStop(server.DB, stopId)
	if err != nil {
		fmt.Println("routeRepo FindAllByStop ", err)
	} else {
		vm.Routes = routes
	}

	stories, err := storyRepo.FindAllByStop(server.DB, stopId)
	if err != nil {
		fmt.Println("storyRepo FindAllByStop ", err)
	} else {
		vm.Stories = stories
	}

	responses.JSON(w, http.StatusOK, vm)
}
