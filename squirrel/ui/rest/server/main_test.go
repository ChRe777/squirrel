package main

import (
	"testing"
)

/*
func TestGetAllProducts(t *testing.T) {
  // Setup
    repositoryMock := catalog.DefaultRepository{}
    repositoryMock.AddProduct(catalog.Product{Id: 1, Name: "Schuhe"})
    router := NewRouter(&repositoryMock)
 
    // When: GET /catalog/products is called
    req, _ := http.NewRequest("GET", "/repl", nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
 
    // Then: status is 200
    assert.Equal(t, http.StatusOK, rr.Code)
 
    // And: Body contains 1 product
    expected := 
      `[{"Id":1,"Name":"Schuhe","Description":"","Category":"","Price":0}]`
    assert.Equal(t, expected, rr.Body.String())
}
*/