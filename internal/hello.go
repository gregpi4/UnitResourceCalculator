package main

import (
	"net/http"

	resinput "github.com/gregpi4/UnitResourceCalculator/internal/resinput"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	http.HandleFunc("/extensions/{extensionId}/products", resinput.GetAllProducts)

	//http.HandleFunc("/", resinput.GetProductInputs)

	http.ListenAndServe(":8080", nil)

}
