package main

import (
	"database/sql"
	"net/http"

	"github.com/juhonamnam/wedding-invitation-server/env"
	"github.com/juhonamnam/wedding-invitation-server/httphandler"
	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "file:./sql.db?_journal_mode=WAL&_sync=OFF")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	if env.UseGuestbook {
		mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	}
	if env.UseAttendance {
		mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))
	}

	corHandler := cors.New(cors.Options{
		AllowedOrigins: []string{env.AllowOrigin},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut},
		//AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)

}
