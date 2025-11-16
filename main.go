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

	filesChanged := getChanges()
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
	// runCmd("git", "commit", "-m", commitMsg)
}

// ---------------------------Ai msg gen helpers---------------------------
func generateMsg(diff string) (string, error) {
    ctx := context.Background()

    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey: apiKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        return "", fmt.Errorf("Error creating client:", err)
    }

    prompt := fmt.Sprintf("Generate a 10 word or less commit message for this diff: %s", diff)
    resp, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(prompt),
        nil,
    )
    if err != nil {
        return "", fmt.Errorf("Error generating message:", err)
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

func getChanges() ([]string) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting changed files:", err)
		return nil
	}

	if strings.TrimSpace(string(out)) == "" {
		return []string{}
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n")
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
