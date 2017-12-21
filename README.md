# Pathgather Coding Challenge
A very simple in-memory key-value store.


## Description
Write a command-line REPL that interfaces with a simple in-memory kv-store.
The system should allow for nested transactions, which can be committed or
aborted.


### Design
The basic design is composed of a `map[string]string` in Golang. However to 
faciliate the history of transactions, the naive solution is to store a copy of
the data store in a stack. This implementation takes it a step further and adds
a stack for each key in the store. This greatly reduces the memory consumption
as it doesn't need to store a complete copy of the stack for every transaction,
only history for the fields which have been modified during transactions.


## Usage
Supported functions are:

| Function                      | Description |
| --------                      | ----------- |
|  *READ* <key>                 | Reads and prints value associated with key, else error                |
|  *WRITE* <key> <value>        | Stores value in key                                                   |
|  *START*                      | Start a transaction                                                   |
|  *COMMIT*                     | Commit a transaction (return error if no transaction)                 |
|  *ABORT*                      | Abort a transaction and discard all actions in current transaction                |
|  *QUIT*                       | Exit REPL cleanly                                                     |
