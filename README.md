# go-internet

My sandbox for re-implementing things from the internet. Both out of curiosity, and out of practice.

- WIP - HTTP 1.1 https://www.rfc-editor.org/rfc/rfc2616

## Commands

Basic API server outputs requests it receives and returns a 200 with "Hello World!"

```sh
go run http_api/main.go

curl localhost:42069
# -> {{GET / HTTP/1.1} [{Host localhost:42069} {User-Agent curl/7.81.0} {Accept */*}] }

curl -X POST localhost:42069 -d '{"ahoy":"matey"}'
# -> {{POST / HTTP/1.1} [{Host localhost:42069} {User-Agent curl/7.81.0} {Accept */*} {Content-Length 16} {Content-Type application/x-www-form-urlencoded}] {"ahoy":"matey"}}
```

## Run tsts

```bash
# from src
go test -v $(find . -name '*_test.go' -exec dirname {} \; | sort -u)
```
