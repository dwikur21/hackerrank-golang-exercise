package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func checkerror(w http.ResponseWriter, err error) {
	var response Response

	log.Println(err)

	response.Status = 2
	response.Message = err.Error()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

func returnAllPegawai(w http.ResponseWriter, r *http.Request) {
	var varPegawai Pegawai
	var arrPegawai []Pegawai
	var response Response

	db := connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT [PEG_NPKID],[PEG_NAME],[ACTIVE] FROM [OCBC_EMERGING_DATA].[dbo].[RFPEGAWAI]")
	if err != nil {
		checkerror(w, err)
	}

	for rows.Next() {
		if err := rows.Scan(&varPegawai.NpkID, &varPegawai.Nama, &varPegawai.Active); err != nil {
			checkerror(w, err)
		}

		arrPegawai = append(arrPegawai, varPegawai)
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = &arrPegawai

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updatePegawai(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var response Response

	params := mux.Vars(r)

	var pegawai Pegawai

	_ = json.NewDecoder(r.Body).Decode(&pegawai)

	sqlresult, err := db.Exec(
		`UPDATE RFPEGAWAI 
		SET PEG_NAME = @p1
		,ACTIVE = @p2
		WHERE PEG_NPKID = @p3`, pegawai.Nama, pegawai.Active, params["id"])

	if err != nil {
		checkerror(w, err)
	}

	result, err := sqlresult.RowsAffected()
	if err != nil {
		checkerror(w, err)
	}

	response.Status = 1
	response.Message = fmt.Sprintf("%v row(s) affected", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createPegawai(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var response Response

	var pegawai Pegawai

	_ = json.NewDecoder(r.Body).Decode(&pegawai)

	sqlresult, err := db.Exec(
		`INSERT INTO RFPEGAWAI
		(PEG_NPKID, PEG_NAME, ACTIVE)
		VALUES
		(@p1, @p2, @p3)`, pegawai.NpkID, pegawai.Nama, pegawai.Active)

	if err != nil {
		checkerror(w, err)
	}

	result, err := sqlresult.RowsAffected()
	if err != nil {
		checkerror(w, err)
	}

	response.Status = 1
	response.Message = fmt.Sprintf("%v row(s) affected", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deletePegawai(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var response Response

	params := mux.Vars(r)

	sqlresult, err := db.Exec(
		`DELETE FROM RFPEGAWAI
		WHERE PEG_NPKID = @p1`, params["id"])

	if err != nil {
		checkerror(w, err)
	}

	result, err := sqlresult.RowsAffected()
	if err != nil {
		checkerror(w, err)
	}

	response.Status = 1
	response.Message = fmt.Sprintf("%v row(s) affected", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
