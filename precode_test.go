package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"

	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Если в параметре count указано больше, чем заданно всего, должны вернуться все доступные кафе.
*/
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4 // В базе есть только 4 кафе
	req, err := http.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount+10)+"&city=moscow", nil)
	assert.NoError(t, err)

	// Создаем новый объект для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Обработчик
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 200
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Ожидаемый результат
	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент" // Переменная для ожидаемого результата

	// Проверяем, что в теле ответа то же, что в ожидаемом значении
	assert.Equal(t, expected, responseRecorder.Body.String())
}
