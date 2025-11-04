# QuickVariables (qv)

A tiny CLI to store, retrieve, update, list, and delete named variables from a local JSON file. Great for keeping quick config values handy during development.

## Quick Start

- Build the binary
  - Requires Go: `go build -o qv`
  - Optionally move it into your `PATH`: `mv qv /usr/local/bin/`
- Configure storage path
  - Create a `.env` file alongside where you run `qv` and set:
    - `QV_PATH=~/.config/quickvariable.json`
- Initialize the store
  - `qv init`

You’re ready to go.

## Usage

- Set a variable
  - `qv set --name api_key --var 12345`
- Get a variable
  - `qv get --name api_key`
- Update a variable
  - `qv update --name api_key --var 67890`
- Delete a variable
  - `qv delete --name api_key`
- List variable names
  - `qv list`

Example: export a value into your shell
- `export API_KEY=$(qv get --name api_key)`

## Commands & Flags

- `init`
  - Initializes the JSON store file.
- `set --name <key> --var <value>`
  - Creates a new variable with the given key/value.
- `get --name <key>`
  - Prints the value for the given key.
- `update --name <key> --var <value>`
  - Replaces the value for an existing key.
- `delete --name <key>`
  - Removes a key/value pair.
- `list`
  - Prints all stored keys.

## Configuration

- The tool reads `.env` using `QV_PATH` to locate your JSON store.
  - Example `.env`:
    - `QV_PATH="/Users/<you>/.config/quickvariable.json"`
- Keep `.env` in the directory from which you invoke `qv` (or use a global `.env` loader setup).

## Notes

- Data is stored as plain JSON on disk (no encryption). Do not store sensitive secrets unless your machine and file permissions are secured.
- If `qv init` doesn’t create the file at your `QV_PATH`, create it manually: `mkdir -p ~/.config && touch ~/.config/quickvariable.json` and re-run your command.
- Running `qv` with no arguments prints the available commands.

## Uninstall

- Delete the binary (`which qv`) and remove the JSON store at `QV_PATH` if you no longer need it.

