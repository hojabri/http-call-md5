# HTTP-CALL-MD5
A tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

### Instructions

1. Make sure you have Go installed ([download](https://golang.org/dl/)).
2. To build an executable commandline file, type in the terminal: `make build`. This will build and create `myhttp` executable file.
5. To run unit tests, type in the terminal: ``make test``

### How to run
First, make sure you have built the program using ``make build``.
Then you can run it typing the command name ``./myhttp`` (Mac/Linux) or ``myhttp.exe`` (Windows)

Syntax:
```
Usage of ./myhttp:
./myhttp [-parallel <size> ] <URL1> <URL2> ...
  -parallel int
        number of parallel requests (default 10)

```

By default, the number of parallel http calls is 10, but you can specify an integer after ``-parallel`` argument, after the command name,
For example:

```
./myhttp -parallel 20 google.com www.google.com https://adjust.com

```
Output:
```
http://www.google.com 3be1646b2df8bf467531e8b687c4535b
http://google.com 23c85714a8ff78569c2a0b0a55796865
https://adjust.com 3744f1e6f3e1517324b5462529ac51ae

```
### Logs
Error logs for http calls will be saved into the ``logs.txt`` file, where the command exists.
