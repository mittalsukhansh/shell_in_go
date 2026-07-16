# K-Shell: A Custom Unix-Style Shell in Go

Most of it I wrote and yes AI has been used but only for learning the syntax and some specific parts which I found hard, apart from that Stackoverflow (the goat) has been referred and basically skipped yt tuts but took help of https://blog.init-io.net/post/2018/07-01-go-unix-shell .
GO docs are freaking overwhelming to read but trying my best to understand it.

## Intro to this project(kind of I would say)
A lightweight, cross-platform command-line interpreter built entirely in Go. 
This project bridges the gap between basic Go syntax and core operating system concepts by interacting directly with the OS kernel to manage standard I/O streams and execute system processes.

## MY Learnings
* **Deeper GO syntax:** Gone deeper into my GO learning journey. 
* **Customizable:** Right now I have just added a few commands like "pwd", "ls" and "cd" using switch (ofc) , just basically experimenting rn.
* **Error Handling:** Learned how to handle errors gracefully. Still working on this part


## Features
* **Cross-Platform Execution:** Utilizes Go's `os/exec` package to natively execute both Windows and Linux binaries.
* **Dynamic Prompt:** Real-time working directory state management using `os.Getwd()`.
* **Custom Built-ins:** Intercepts specific commands (`cd`, `ls`, `pwd`, `exit`) to manage shell state without spawning external processes.
* **Graceful Error Handling:** Prevents shell termination on invalid commands or execution failures.

## Technical Implementation
Unlike standard tutorials, this shell implements its own argument parsing using `bufio` and `strings` to prevent "giant file name" execution errors, and cleanly routes `Stdout` and `Stderr` streams back to the parent terminal window.

## How to Run
```bash
go run shell.go
