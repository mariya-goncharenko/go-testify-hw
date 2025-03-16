package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
Тест 1:

Когда параметр count больше, чем есть кафе, должны вернуть все доступные кафе.
*/
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4 // В базе всего 4 кафе
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status 200 OK")

	// Получаем тело ответа и проверяем количество кафе
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// Проверяем, что количество кафе в ответе равно ожидаемому
	assert.Len(t, list, totalCount, "Expected cafe count: %d, got %d", totalCount, len(list))

	// Ожидаем, что ответ будет содержать все кафе
	expectedResponse := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedResponse, body, "Response should contain all cafes")
}

/*
Тест 2:

Успешный запрос с параметрами count и city.

Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
*/
func TestMainHandlerSuccessfulRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err, "Failed to create HTTP request") // Ошибка при создании запроса

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status 200 OK")
	// Проверяем, что ответ не пустой
	assert.NotEmpty(t, responseRecorder.Body.String(), "Response should not be empty")
}

/*
Тест 3:

Город, который передаётся в параметре city, не поддерживается.

Сервис возвращает код ответа 400 и ошибку "wrong city value" в теле ответа.
*/
func TestMainHandlerUnsupportedCity(t *testing.T) {
	// Создаем запрос с неподдерживаемым городом "unknown"
	req, err := http.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
	require.NoError(t, err, "Failed to create HTTP request") // Ошибка при создании запроса

	// Регистрируем рекордер для захвата ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 400 Bad Request
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status 400 Bad Request")

	// Проверяем, что в теле ответа правильное сообщение об ошибке
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Incorrect error message")
}
