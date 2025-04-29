package model

import (
    "os/exec"
    "strings"
)

func PredictText(text string) (string, error) {
    out, err := exec.Command("python3", "model_runner/run_model.py", text).Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)), nil
}

