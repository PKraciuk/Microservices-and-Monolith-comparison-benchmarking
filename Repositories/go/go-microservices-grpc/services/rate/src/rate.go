package src

import (
	"encoding/json"
	"fmt"
	rate "github.com/harlow/go-micro-services/rate/proto"
	"log"
	"net"

	"github.com/harlow/go-micro-services/rate/data"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// New returns a new server
func New() *Rate {
	return &Rate{
		rateTable: loadRateTable("data/inventory.json"),
	}
}

// Rate implements the rate service
type Rate struct {
	rateTable map[stay]*rate.RatePlan
}

// Run starts the server
func (s *Rate) Run(port int) error {
	srv := grpc.NewServer()
	rate.RegisterRateServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

// GetRates gets rates for hotels for specific date range.
func (s *Rate) GetRates(ctx context.Context, req *rate.Request) (*rate.Result, error) {
	res := new(rate.Result)

	for _, hotelID := range req.HotelIds {
		stay := stay{
			HotelID: hotelID,
			InDate:  req.InDate,
			OutDate: req.OutDate,
		}
		if s.rateTable[stay] != nil {
			res.RatePlans = append(res.RatePlans, s.rateTable[stay])
		}
	}

	return res, nil
}

// loadRates loads rate codes from JSON file.
func loadRateTable(path string) map[stay]*rate.RatePlan {
	file := data.MustAsset(path)

	var rates []*rate.RatePlan
	if err := json.Unmarshal(file, &rates); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	rateTable := make(map[stay]*rate.RatePlan)
	for _, ratePlan := range rates {
		stay := stay{
			HotelID: ratePlan.HotelId,
			InDate:  ratePlan.InDate,
			OutDate: ratePlan.OutDate,
		}
		rateTable[stay] = ratePlan
	}

	return rateTable
}

type stay struct {
	HotelID string
	InDate  string
	OutDate string
}
