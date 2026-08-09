package http_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"

	serverhttp "github.com/GalvinGao/gofiber-template/internal/server/http"
)

func TestCreateServesSuccessfulRequests(t *testing.T) {
	t.Parallel()

	app := serverhttp.Create()
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	resp := testRequest(t, app, http.MethodGet, "/health", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}

func TestCreateRecoversPanics(t *testing.T) {
	t.Parallel()

	app := serverhttp.Create()
	app.Get("/panic", func(c fiber.Ctx) error {
		panic("boom")
	})

	resp := testRequest(t, app, http.MethodGet, "/panic", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
}

func TestCreateHandlesCORSPreflight(t *testing.T) {
	t.Parallel()

	app := serverhttp.Create()
	req := httptest.NewRequest(http.MethodOptions, "/resource", nil)
	req.Header.Set("Origin", "https://example.com")
	req.Header.Set("Access-Control-Request-Method", http.MethodGet)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("test request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
	if got := resp.Header.Get("Access-Control-Allow-Origin"); got != "*" {
		t.Fatalf("Access-Control-Allow-Origin = %q, want %q", got, "*")
	}
}

func TestCreateRateLimitsRequests(t *testing.T) {
	t.Parallel()

	app := serverhttp.Create()
	app.Get("/limited", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	for requestNumber := 1; requestNumber <= 21; requestNumber++ {
		resp := testRequest(t, app, http.MethodGet, "/limited", nil)
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()

		want := http.StatusOK
		if requestNumber == 21 {
			want = http.StatusTooManyRequests
		}
		if resp.StatusCode != want {
			t.Fatalf("request %d status code = %d, want %d", requestNumber, resp.StatusCode, want)
		}
	}
}

func testRequest(t *testing.T, app *fiber.App, method, path string, body io.Reader) *http.Response {
	t.Helper()

	resp, err := app.Test(httptest.NewRequest(method, path, body))
	if err != nil {
		t.Fatalf("test request: %v", err)
	}
	return resp
}
