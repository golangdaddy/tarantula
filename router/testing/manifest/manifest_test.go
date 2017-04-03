package manifest

import (
    "fmt"
    "testing"
)

func TestOpenManifest(t *testing.T) {

    manifest, err := (&App{}).OpenManifest("./manifest.example.json")
    if err != nil {
        t.Error(err)
        return
    }

    if len(manifest.Endpoints) != 1 {
        t.Error("INVALID RESULT")
        return
    }

    if len(manifest.Variables) != 3 {
        t.Error("INVALID RESULT")
        return
    }

    fmt.Println(manifest)
}
