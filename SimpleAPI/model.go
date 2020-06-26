package main

// Response Container for the data
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *[]Pegawai
}

// Pegawai The data model for Pegawai table
type Pegawai struct {
	NpkID  string `json:"npkid"`
	Nama   string `json:"nama"`
	Active bool   `json:"active"`
}
