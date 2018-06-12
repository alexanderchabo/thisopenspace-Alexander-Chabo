# thisopenspace-Alexander-Chabo : Challenge
## Overview
I chose to retrieve the information from the JSON endpoint to a back-end server written in go. One of the reasons I chose to develop the server in go is due to the fact that it is a language I'm currently learning. In addition to this, I thought it would be an exciting challenge to create a back-end that handles all logic, including pagination.

The front-end is written only in HTML & CSS and handles all information regarding spaces and pagination through parameters that are passed through from the backend.

It took me roughly 8h (including research) to get everything up and running. I hope you enjoy my solution to this challenge!

### Directory layout

    .
    ├── ...
    ├── src                       # Source files (Application)
    │   ├── main.go               # Main file
    │   ├── handler.go            # Handles the http request
    │   ├── model.go              # Define structs
    │   └── templates             # HTML template folder
    │       ├── index.html        # View for spaces and paginator
    │       └── styles            # Stylesheets
    │           └── css.index     # CSS for spaces and paginator
    └── ...

## Install instructions

The following instructions are written in consideration to a Mac environment.
```
1. Follow the official instructions to install go on the golang site: https://golang.org/dl/
```
```
2. Verify that go is installed with command 'go version'
```
```
3. Get repo with command 'go get github.com/alexanderchabo/thisopenspace-Alexander-Chabo'
```
```
4. cd to folder 'thisopenspace-Alexander-Chabo'. The default go root path $GOPATH$ is "/Users/{Your username}/go" on Mac. If needed, use command 'go env' to find your corresponding $GOPATH$.
```
```
5. Run command 'go run src/*.go' to start the server
```
Browse to http://localhost:8080/ to see solution.

## Challenge
This test has no time limit, but for reference, candidates have traditionally done this in a couple of hours. There will be no penalty for extra time taken.

We have a JSON endpoint that displays information on spaces here https://thisopenspace.com/lhl-test?page=(x) where the (x) is a page number.

We want to have a UI where we can easily see all the listings, through pagination controls or infinite scrolling. It should not show all the spaces all at once.

The UI can be done in any form, even command-line if you are more comfortable with back-end.

You'll be evaluated on:
- Code/component organization
- Code maintainability
- If things actually work
- Library choices - do you choose libraries that you don't need?

Please submit:
- The code
- Any instructions to run it if needed
- A quick comment on how much time was taken, split between research (if needed) and implementation.
