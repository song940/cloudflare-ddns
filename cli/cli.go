package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var (
	api      *cloudflare.API
	zoneName string
	hostName string
)

func Run() {

	key := os.Getenv("CLOUDFLARE_KEY")
	email := os.Getenv("CLOUDFLARE_EMAIL")
	token := os.Getenv("CLOUDFLARE_TOKEN")

	var err error

	if token != "" {
		api, err = cloudflare.NewWithAPIToken(token)
	} else {
		api, err = cloudflare.New(key, email)
	}

	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&zoneName, "zone", "", "zone name")
	flag.StringVar(&hostName, "hostname", "", "hostname")
	flag.Parse()

	ctx := context.Background()

	user, err := api.UserDetails(ctx)
	fmt.Println("User:", user.LastName, user.FirstName)
	fmt.Println("User:", user.Email)
	fmt.Println("User:", user.Username)

	zoneID, err := api.ZoneIDByName(zoneName)
	fmt.Println("Zone:", zoneID, zoneName)
	dnsRecords, err := api.DNSRecords(ctx, zoneID, cloudflare.DNSRecord{Name: hostName})
	for _, record := range dnsRecords {
		fmt.Println("Record:", record.ID, record.Name, record.Type, record.Content)
	}
	// err := api.UpdateDNSRecord(ctx, record.ZoneID, record.ID, *record);
}
