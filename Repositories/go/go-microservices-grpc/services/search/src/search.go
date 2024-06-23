package src

import (
	"fmt"
	search "github.com/harlow/go-micro-services/services/search/proto"
	geo "github.com/harlow/go-micro-services/services/search/proto/geo"
	rate "github.com/harlow/go-micro-services/services/search/proto/rate"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func New(geoConnection, rateConnection *grpc.ClientConn) *Search {
	return &Search{
		geoClient:  geo.NewGeoClient(geoConnection),
		rateClient: rate.NewRateClient(rateConnection),
	}
}

type Search struct {
	geoClient  geo.GeoClient
	rateClient rate.RateClient
}

func (s *Search) Run(port int) error {
	srv := grpc.NewServer()
	search.RegisterSearchServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return srv.Serve(lis)
}

// Nearby returns ids of nearby hotels ordered by ranking algo
func (s *Search) Nearby(ctx context.Context, req *search.NearbyRequest) (*search.SearchResult, error) {
	// find nearby hotels
	nearby, err := s.geoClient.Nearby(ctx, &geo.Request{
		Lat: req.Lat,
		Lon: req.Lon,
	})
	if err != nil {
		log.Fatalf("nearby error: %v", err)
	}

	// find rates for hotels
	rates, err := s.rateClient.GetRates(ctx, &rate.Request{
		HotelIds: nearby.HotelIds,
		InDate:   req.InDate,
		OutDate:  req.OutDate,
	})
	if err != nil {
		log.Fatalf("rates error: %v", err)
	}

	// build the response
	res := new(search.SearchResult)
	for _, ratePlan := range rates.RatePlans {
		res.HotelIds = append(res.HotelIds, ratePlan.HotelId)
	}

	return res, nil
}
