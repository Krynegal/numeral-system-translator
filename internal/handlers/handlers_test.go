package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Krynegal/numeral-system-translator.git/internal/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, method, path string, body []byte) (*http.Response, models.Response) {
	ts := httptest.NewServer(NewHandler().Router)
	defer ts.Close()
	req, err := http.NewRequest(method, ts.URL+path, bytes.NewBuffer(body))
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	defer resp.Body.Close()

	var response models.Response
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)
	return resp, response
}

func TestHandlerAdd(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  string
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "5",
			operand2: "7",
			wantRes:  "12",
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "-5",
			operand2: "7",
			wantRes:  "2",
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "",
			wantRes:  "",
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0",
			operand2: "3",
			wantRes:  "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//reqBody := `{"":"", "":""}`
			reqBody := smpl_calc_api.Data{
				Operand1: smpl_calc_api.Operand{
					Value: tt.operand1,
					Base:  10,
				},
				Operand2: smpl_calc_api.Operand{
					Value: tt.operand2,
					Base:  10,
				},
				ToBase: 10,
			}
			r, _ := json.Marshal(reqBody)
			resp, body := testRequest(t, http.MethodPost, "/api/add", r)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Result)
		})
	}
}

func TestHandlerSub(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  string
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "-11",
			operand2: "33",
			wantRes:  "-44",
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "0",
			operand2: "7",
			wantRes:  "-7",
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "",
			wantRes:  "",
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0",
			operand2: "-3",
			wantRes:  "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := smpl_calc_api.Data{
				Operand1: smpl_calc_api.Operand{
					Value: tt.operand1,
					Base:  10,
				},
				Operand2: smpl_calc_api.Operand{
					Value: tt.operand2,
					Base:  10,
				},
				ToBase: 10,
			}
			r, _ := json.Marshal(reqBody)
			resp, body := testRequest(t, http.MethodPost, "/api/sub", r)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Result)
		})
	}
}

func TestHandlerMul(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  string
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "5",
			operand2: "7",
			wantRes:  "35",
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "-9",
			operand2: "0",
			wantRes:  "0",
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-1",
			operand2: "",
			wantRes:  "",
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0",
			operand2: "3",
			wantRes:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := smpl_calc_api.Data{
				Operand1: smpl_calc_api.Operand{
					Value: tt.operand1,
					Base:  10,
				},
				Operand2: smpl_calc_api.Operand{
					Value: tt.operand2,
					Base:  10,
				},
				ToBase: 10,
			}
			r, _ := json.Marshal(reqBody)
			resp, body := testRequest(t, http.MethodPost, "/api/mul", r)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Result)
		})
	}
}

func TestHandlerDiv(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  string
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "44",
			operand2: "2",
			wantRes:  "22",
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "3",
			operand2: "8",
			wantRes:  "0",
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "0",
			wantRes:  "",
		},
		{
			name:     "test #4",
			success:  false,
			operand1: "-0",
			operand2: "",
			wantRes:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := smpl_calc_api.Data{
				Operand1: smpl_calc_api.Operand{
					Value: tt.operand1,
					Base:  10,
				},
				Operand2: smpl_calc_api.Operand{
					Value: tt.operand2,
					Base:  10,
				},
				ToBase: 10,
			}
			r, _ := json.Marshal(reqBody)
			resp, body := testRequest(t, http.MethodPost, "/api/div", r)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Result)
		})
	}
}
