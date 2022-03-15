package controllertest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"merchant-api/controllers"

	"github.com/gin-gonic/gin"
)

func TestGetTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/transactions", controllers.GetTransaction)

	req, err := http.NewRequest(http.MethodGet, "/transactions", nil)

	if err != nil {
		t.Fatalf("Could not request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
