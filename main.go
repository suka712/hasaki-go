package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/genai"
)

func main() {
	err := runCmd("git", "add", ".")
    check(err, "Error running 'git add'")

	filesChanged, err := getChanges()
    check(err, "Error getting changed files")
	if len(filesChanged) == 0 {
		fmt.Println("No file change. No commit made.")
		return
	}

	diff, err := getDiff()
    check(err, "Error getting diff")
	if len(diff) == 0 {
		fmt.Println("Empty diff. No commit made.")
		return
	}

    msg, err := generateMsg(diff)
    check(err, "Error generating message")
    if len(msg) == 0 {
		fmt.Println("Empty message. No commit made.")
        return
    }

    fmt.Println(msg)

	err = runCmd("git", "commit", "-m", msg)
    check(err, "Error running 'git commit'")
}

// ---------------------------Go boilerplate---------------------------
func check(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %v", msg, err)
    }
}

// ---------------------------Ai msg gen---------------------------
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

// ---------------------------Shell & Git---------------------------
func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error running command '%s': %w", name, err)
    }
    return nil
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

func getDiff() (string, error) {
	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running git diff: %w", err)
	}

	return string(out), nil
}
