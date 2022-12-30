package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	IPAddress := r.Header.Get("X-Real-Ip")

	fmt.Fprintf(w, IPAddress)


	// fmt.Fprint(w, `
	// 	<h1>Frisco ISD HAC API</h1>
	// 	<p>Documentation:</p>
	// 	<a href="https://friscoisdhacapidocs.vercel.app/" target="_blank">https://friscoisdhacapidocs.vercel.app/</a>
	// `)
}
