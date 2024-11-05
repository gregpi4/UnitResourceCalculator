package resinput

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func getWildcard(pattern string) (bool, string) {
	if strings.HasPrefix(pattern, "{") && strings.HasSuffix(pattern, "}") {
		var identifier = strings.TrimPrefix(pattern, "{")
		identifier = strings.TrimSuffix(identifier, "}")
		return true, identifier
	}
	return false, ""
}

func parseURL(r *http.Request) (map[string]string, error) {
	pattern := strings.Split(r.Pattern, "/")
	var value = strings.Split(r.URL.Path, "/")
	var parsedWildcard = map[string]string{}

	for i := 0; i < len(pattern); i++ {
		isWildcard, wildcard := getWildcard(pattern[i])
		if isWildcard {
			parsedWildcard[wildcard] = value[i]
		}
	}
	return parsedWildcard, nil
}

func logCall(r *http.Request) {
	log.Printf("URL called: %s", r.URL)
}

func formatResponse(w http.ResponseWriter, data map[string][]string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	logCall(r)
	var parsed_url, _ = parseURL(r)
	log.Printf("Parsed URL: %s", parsed_url)

	inputs_needed := GetProductsName(parsed_url["extensionId"])
	formatResponse(w, map[string][]string{"inputs": inputs_needed}, http.StatusOK)
}

func DesignController(w http.ResponseWriter, r *http.Request) {
	logCall(r)
	var parsed_url, _ = parseURL(r)
	log.Printf("Parsed URL: %s", parsed_url)
}

func GetProductInputs(w http.ResponseWriter, r *http.Request) {
	logCall(r)
	var parsed_url, _ = parseURL(r)
	log.Printf("Parsed URL: %s", parsed_url)

	//inputs_needed, _ := Calculate("A")
	//formatResponse(w, map[string]map[string]int{"inputs": inputs_needed}, http.StatusOK)
}
