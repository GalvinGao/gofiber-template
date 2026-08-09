package route

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestCreateGroupsRegistersExpectedPrefixes(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	groups := CreateGroups(app)
	groups.APIV1.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("api")
	})
	groups.Internal.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("internal")
	})

	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{name: "API v1", path: "/api/v1/health", wantStatus: http.StatusOK, wantBody: "api"},
		{name: "internal", path: "/internal/health", wantStatus: http.StatusOK, wantBody: "internal"},
		{name: "unprefixed", path: "/health", wantStatus: http.StatusNotFound, wantBody: "Not Found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := app.Test(httptest.NewRequest(http.MethodGet, tt.path, nil))
			if err != nil {
				t.Fatalf("test request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("status code = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			if got := string(body); got != tt.wantBody {
				t.Fatalf("body = %q, want %q", got, tt.wantBody)
			}
		})
	}
}
