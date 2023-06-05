package main

import (
	"github.com/RemyRanger/go-etl-csv/internal/common/config"
	"github.com/RemyRanger/go-etl-csv/internal/common/database/mysql"
	"github.com/RemyRanger/go-etl-csv/internal/common/logs"
	"github.com/RemyRanger/go-etl-csv/internal/router"
)

func main() {
	// init config
	config.Conf = config.New("ticketstore.yaml")

	// init logger
	logs.New()

	// init db
	mysql := mysql.New()

	// init router
	router.NewTicketstoreRouter(mysql).Serve()
}
