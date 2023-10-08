package sql_test

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/database/sql"
	"github.com/shogo82148/std/net/http"
)

var pool, db *sql.DB

var ctx = context.Background()

type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request)

func Ping(ctx context.Context)

func Query(ctx context.Context, id int64)
