package endToEndTests

import (

)

func (s *MyFirstSuite) TestAddClient(t provider.T) { 

}


gin.SetMode(gin.TestMode)

	router := gin.Default()
	mockService := &mockUserService{users: make(map[string]pkg.User)}
	handler := handlers.NewHandler(mockService) // Предполагается, что у вас есть конструктор для Handler

	router.POST("/signup", handler.SignUp)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, input.Email, response["id"])
