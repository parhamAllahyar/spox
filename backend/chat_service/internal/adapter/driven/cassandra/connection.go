package cassandra

import (
	"chat/configs"

	"github.com/gocql/gocql"
)

func CassandraConnection(config configs.CassandaraDBConfig) (*gocql.Session, error) {

	//TODO: Congif
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "restfulapi"
	return cluster.CreateSession()

}
