# Pland
`pland` is a daemon that watches the .plan file and syncs it with the server on edit. It is purposefully kept as a separate file to keep system overhead as small as possible, and only imports from the shared cli libraries as necessary.

# Features
- `install` which allows it to be run on startup
- `uninstall` which removes it from startup