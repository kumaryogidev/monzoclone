package utils

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func ConnectCassandra() {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "monzobank"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 10 * time.Second

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	log.Println("Connected to Cassandra successfully!")
}
