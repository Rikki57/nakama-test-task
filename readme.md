# Nakama Custom RPC Function

## Overview

This repository contains a custom RPC function for Nakama that reads a file from disk, calculates its hash, and saves the data to the database. The function accepts a payload with optional parameters (type, version, hash) and responds with the file content and hash.

## Running the Solution

1. Ensure you have Docker and Docker Compose installed.
2. Clone this repository.
3. Navigate to the project directory.
4. Run the following command to start the services:

```sh
docker-compose up
```
## Running the test
The test client is located in
https://github.com/Rikki57/nakama-test-task-test
See readme.md of this repo for details

## Explanation of the solution
The standard approach offered by the Nakama team is utilized: a Go module containing the implementation for an RPC function, along with a special function for registering this RPC function within the Nakama framework. The function is described in the count-hash.go file. Every file read operation creates a record in the corresponding SQL table. However, the current implementation is basic and requires further clarification of requirements. Presently, it serves as a proof of concept for the ability to interact with the database.

## Thoughts, and ideas about the task
If this were a real task, I would like to clarify several details (unfortunately, I couldn't do so because I started the task over the weekend). What is the purpose of this process? Are these files with some content obtained from outside the app? If so, we need some security checks. If it is internal storage, some caching could be useful to avoid reading the physical file every time. We also need some system that will bring tht files to the reading folder. Moreover, what is the typical file size? Maybe it is better to read the file immediately after upload and store its content in a MongoDB-like storage? What is the purpose of the database? Do we need to use it as a logbook for every request, or do we need to record the latest reading fact, or something else? Maybe we need to try searching for the content inside the DB, and if it's not present, read the file? Additionally, the hash check is odd: we return the hash every time, even if the requestor provides the wrong hash. This means the requestor can always make two requests: get the hash from the first response and use it in the second request, ensuring they always have the correct hash on the second request. If the hash is required for security reasons, this approach is unsafe, and we need to hide the hash if it does not match the requested one.

## How I would improve it if I had more time.

1. When I implemented the task, I had no experience in Go and Nakama. The first step would be improving this basic solution by adhering to code conventions and best practices regarding project structure, code style etc.
2. The solution needs some parametrized configuration (DB passwords and other parameters definitely must not be stored inside the code)
3. This solution lacks security checks. Anyone can send a request, so it must be filtered. Some security issues are already described in the previous section (insecure file reading, hash-based control not working with the current requirements).
4. The test does not check database changes. Additionally, a better technology might be chosen (I used Java because I could implement it quickly).
5. The logic needs significant changes after requirements clarification (the current solution has a lot of logical gaps).
6. The solution currently does nos suppose any devops automation.