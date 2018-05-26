package main

//Warning usado para enviar avisos na API
type Warning struct {
	Fail    bool   `json:"fail"`
	Message string `json:"message"`
}
