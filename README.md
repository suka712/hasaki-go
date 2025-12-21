# Hasaki Go
An app to automate git commits.

## To use the app
**Step 1:** Run
```bash
git clone git@github.com:suka712/hasaki-go.git

cd hasaki-go
```
**Step 2:** Create `secret.go` similar to `secret.go.example` and fill in your own Gemini API key.

**Step 3:** Then run
```bash
go build -o hx && sudo mv hx /usr/local/bin/
```
And the app should be available via `hx` in the terminal.

## Example uses
For automatic message generation, run `hx`.
```bash
hasaki-go ❯ hx
Generating commit message...
[master 107e0e8] Clarify boilerplate comment.
 1 file changed, 1 insertion(+), 1 deletion(-)
╭──────────────────────────────────────────────────────────╮
│ Message: Clarify boilerplate comment.                    │
│ Changed: main.go                                         │
╰──────────────────────────────────────────────────────────╯
khiem@mcx ~/Repos/hasaki-go ❯ hx -m "Add ability to add manual message."
```
To manually insert your command, run `hx -m "My manual commit message."`
```bash
hasaki-go ❯ hx -m "Remove empty line."
[master 4d9cf18] Remove empty line.
 1 file changed, 1 insertion(+), 3 deletions(-)
╭──────────────────────────────────────────────────────────╮
│ Message: Remove empty line.                              │
│ Changed: main.go                                         │
╰──────────────────────────────────────────────────────────╯
```