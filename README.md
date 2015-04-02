# Go UTM Stripper
Strips UTM parameters from URLs

#### Usage
Input must have 1 arguments:

1.  -url

#### Example Input - Output
-
Input:
```
go run main.go "postgres://user:pass@host.com:5432/path?k=v&utf_blah=2&utm_holder=10&aaa=5#f"
```
Output (Success):
```
{"result":"postgres://user:pass@host.com:5432/path?aaa=5\u0026k=v\u0026utf_blah=2#f"}
```
-
Output (Failure):
```
{
  "error":"<error message>"
}
```

#### How to build docker image
Requirements:

1. Golang environment set up
2. Git
3. boot2docker running

```
go get https://github.com/cloudspace/Go-UTM-Stripper
cd <Go-UTM-Stripper directory>/Go-UTM-Stripper
docker run --rm -v $(pwd):/src centurylink/golang-builder
docker build -t <username>/go-utm-stripper:0.3 ./

```

In order for `docker run --rm -v $(pwd):/src centurylink/golang-builder` to work you need to have the github url on the top line of main.go. It should look like this:
```
package main // import "github.com/cloudspace/Go-UTM-Stripper"
```
You also must *push your code* to github *before building* the docker image.
