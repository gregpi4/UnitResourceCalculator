package resinput

// import (
// 	"encoding/json"
// 	"net/http/httptest"
// 	"reflect"
// 	"testing"
// )

// func TestCalculateProduct(t *testing.T) {
// 	recorder := httptest.NewRecorder()
// 	wantBody := map[string]map[string]int{"inputs": {"atom": 100}}

// 	GetProductInputs(recorder, nil)
// 	if recorder.Code != 200 {
// 		t.Errorf("Expected status code 200, got %d", recorder.Code)
// 	}

// 	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json; charset=utf-8" {
// 		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
// 	}

// 	var gotBody map[string]map[string]int
// 	err := json.Unmarshal(recorder.Body.Bytes(), &gotBody)
// 	if err != nil {
// 		t.Errorf("failed to unmarshal response body: %v", err)
// 	}

// 	if !reflect.DeepEqual(gotBody, wantBody) {
// 		t.Errorf("expected response body to be %v, got %v", wantBody, recorder.Body)
// 	}
// }
