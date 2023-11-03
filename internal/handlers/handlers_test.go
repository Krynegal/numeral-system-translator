package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Krynegal/numeral-system-translator.git/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		converter := &ConverterMock{
			ConvertFunc: func(num string, base int, toBase int) (string, error) {

				return "", nil
			},
		}

		r := NewRouter(converter)
		r.POST("/translate", TranslateHandler(converter))

		number := new(string)
		*number = "A"
		base := new(int)
		*base = 16
		toBase := new(int)
		*toBase = 10

		company := models.Request{
			Number: number,
			Base:   base,
			ToBase: toBase,
		}

		jsonValue, _ := json.Marshal(company)
		req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("converter", func(t *testing.T) {
			converter := &ConverterMock{
				ConvertFunc: func(num string, base int, toBase int) (string, error) {

					return "", errors.New("error")
				},
			}

			r := NewRouter(converter)
			r.POST("/translate", TranslateHandler(converter))

			number := new(string)
			*number = "A"
			base := new(int)
			*base = 16
			toBase := new(int)
			*toBase = 10

			company := models.Request{
				Number: number,
				Base:   base,
				ToBase: toBase,
			}

			jsonValue, _ := json.Marshal(company)
			req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("base", func(t *testing.T) {
			converter := &ConverterMock{
				ConvertFunc: func(num string, base int, toBase int) (string, error) {

					return "", nil
				},
			}

			r := NewRouter(converter)
			r.POST("/translate", TranslateHandler(converter))

			number := new(string)
			*number = "1245"
			base := new(int)
			*base = 16

			company := models.Request{
				Number: number,
				Base:   base,
			}

			jsonValue, _ := json.Marshal(company)
			req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("number", func(t *testing.T) {
			converter := &ConverterMock{
				ConvertFunc: func(num string, base int, toBase int) (string, error) {

					return "", nil
				},
			}

			r := NewRouter(converter)
			r.POST("/translate", TranslateHandler(converter))

			number := new(string)
			*number = "-1245"
			base := new(int)
			*base = 16

			company := models.Request{
				Number: number,
				Base:   base,
			}

			jsonValue, _ := json.Marshal(company)
			req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	})
}
