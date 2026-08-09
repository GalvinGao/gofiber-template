package dto

import (
	"encoding/json"
	"testing"
)

func TestUpdatePostDTOPreservesNullSemantics(t *testing.T) {
	t.Parallel()

	var update UpdatePostDTO
	if err := json.Unmarshal([]byte(`{"title":"Hello","description":null}`), &update); err != nil {
		t.Fatalf("unmarshal update: %v", err)
	}

	if !update.Title.Valid || update.Title.String != "Hello" {
		t.Fatalf("title = %#v, want a valid Hello value", update.Title)
	}
	if update.Description.Valid {
		t.Fatalf("description = %#v, want null", update.Description)
	}

	encoded, err := json.Marshal(update)
	if err != nil {
		t.Fatalf("marshal update: %v", err)
	}
	if got, want := string(encoded), `{"title":"Hello","description":null}`; got != want {
		t.Fatalf("JSON = %s, want %s", got, want)
	}
}
