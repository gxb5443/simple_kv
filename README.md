# Pathgather Coding Challenge
A very simple in-memory key-value store.

## Description
Write a command-line REPL that interfaces with a simple in-memory kv-store.
The system should allow for nested transactions, which can be committed or
aborted.

## Usage
Supported functions are:
| Function                      | Description |
| --------                      | ----------- |
|  *READ* <key>                 | Reads and prints value associated with key, else error                |
|  *WRITE* <key> <value>        | Stores value in key                                                   |
|  *DELETE* <key>               | Removes key from store                                                |
|  *START*                      | Start a transaction                                                   |
|  *COMMIT*                     | Commit a transaction (return error if no transaction)                 |
|  *ABORT*                      | Abort a transaction and discard all actions in current transaction                |
|  *QUIT*                       | Exit REPL cleanly                                                     |
