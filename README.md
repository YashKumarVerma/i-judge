# I Judge

and i do it good. This project is **not an online judge**. It only executes programs.

## Is it of my use
It's a simple prototype to run a program / binary / object file using system calls, with provisions to pipe input into stdin, and capture stdout. If that's what you are looking for, then yes, it's of your use.

## Build files
Try not using them as they can be outdated, I'd recommend building them yourselves by
```
make pack
```

## Running locally
Simply run the program by 
```
make run
```
You might need to change the absolute path given in `internal/main.go` to make it work.