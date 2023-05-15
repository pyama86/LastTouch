package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type Response struct {
	UpdateTime string `json:"update_time"`
}

func main() {
	http.HandleFunc("/getUpdateTime", getUpdateTimeHandler)
	log.Fatal(http.ListenAndServe(getEnv("LASTTOUCH_LISTEN", "127.0.0.1:8080"), nil))
}

func checkAuth(r *http.Request) bool {
	clientID, clientSecret, ok := r.BasicAuth()
	if ok == false {
		return false
	}
	return clientID == os.Getenv("LASTTOUCH_USER") && clientSecret == os.Getenv("LASTTOUCH_PASSWORD")
}

func getUpdateTimeHandler(w http.ResponseWriter, r *http.Request) {
	if !checkAuth(r) {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}

	db := r.URL.Query().Get("db")
	table := r.URL.Query().Get("table")
	if db == "" || table == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: Both 'db' and 'table' parameters must be provided\n"))
		return
	}

	updateTime, err := getUpdateTime(db, table)
	if err != nil {
		logrus.WithError(err).Error("Failed to get update time")
		http.Error(w, "Failed to get update time", http.StatusInternalServerError)
		return
	}

	resp := Response{UpdateTime: updateTime}
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal response")
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getUpdateTime(database, tableName string) (string, error) {

	var mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB string

	mysqlUser = getEnv("MYSQL_USER", "root")
	mysqlPassword = getEnv("MYSQL_PASSWORD", "")
	mysqlHost = getEnv("MYSQL_HOST", "localhost")
	mysqlPort = getEnv("MYSQL_PORT", "3306")
	mysqlDB = getEnv("MYSQL_DB", "lasttouch")

	dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return "", err
	}
	defer db.Close()

	const query = `SELECT UPDATE_TIME FROM information_schema.tables WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`
	row := db.QueryRow(query, database, tableName)

	var updateTime string
	err = row.Scan(&updateTime)
	if err != nil {
		return "", err
	}

	return updateTime, nil
}
