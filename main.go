package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/genai"
)

func main() {
	runCmd("git", "add", ".")

	filesChanged, err := getChanges()
    if err != nil {
        fmt.Printf("Error getting file change: %s", err)
        return
    }
	if len(filesChanged) == 0 {
		fmt.Println("No file change. No commit made.")
		return
	}

	diff := getDiff()
	if len(diff) == 0 {
		fmt.Println("Empty commit message. No commit made.")
		return
	}

    msg, err := generateMsg(diff)
    if len(msg) == 0 {
		fmt.Println("Error generating message:", err)
        return
    }

    fmt.Println(msg)

	// commitMsg := "Auto generated: Made changes to the code."
	runCmd("git", "commit", "-m", msg)
}

// ---------------------------Ai msg gen helpers---------------------------
func generateMsg(diff string) (string, error) {
    ctx := context.Background()

    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey: apiKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        return "", fmt.Errorf("error creating client: %w", err)
    }

    prompt := fmt.Sprintf("Generate a 10 word or less commit message for this diff: %s", diff)
    resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(prompt),
        nil,
    )
    if err != nil {
        return "", fmt.Errorf("error generating message: %w", err)
    }

    text := resp.Text()
    result := strings.TrimSpace(text)
    result = strings.ReplaceAll(result, "\"", "")
    result = strings.ReplaceAll(result, "\\", "")

    return result, nil
}

// ---------------------------Shell & Git helpers---------------------------
func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func getChanges() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting changed files: %w", err)
	}

	if strings.TrimSpace(string(out)) == "" {
		return []string{}, nil
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func getDiff() string {
	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting diff:", err)
		return ""
	}

	return string(out)
}
