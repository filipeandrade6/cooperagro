package handler

// func TestCreateUser(t *testing.T) {

// 	bpDB := baseproduct.newInmem()
// 	bpservice := baseproduct.NewService(bpDB)

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":alsdkfjalçsfjd}`))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	h := &handler{mockDB}

// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
