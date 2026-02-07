# dsame

`dsame` is a simple command-line tool inspired by [`anew`](https://github.com/tomnomnom/anew) by tomnomnom.  
While `anew` appends new lines that do **not exist** in a file, `dsame` does the **opposite**: it filters out any lines from stdin that already exist in a comparison file.  

This makes `dsame` useful for tasks like filtering domain or subdomain from lists, where you want to **remove duplicates** based on another reference file.

---

## Example case: Filtering HTTP subdomains already available on HTTPS

Imagine you have two lists of alive subdomains after probing with `httpx`/`httprobe`:

```text
[~] $ > cat http-domains.txt
target.com
sub1.target.com
sub2.target.com

[~] $ > cat https-domains.txt
target.com
sub2.target.com
```

If you want to get only HTTP subdomains that are not available on HTTPS, you can use dsame:

### Basic Usage :

```text
[~] $ > cat http-domains.txt | dsame -f https-domains.txt
sub1.target.com
```
Or with flag and arguments
- -o : Output string, save the output to file
```text
[~] $ > cat http-domains.txt | dsame -f https-domains.txt -o pure-http-domains-only.txt
```
- --no-trim : Do not trim leading/trailing whitespace before comparison
```text
[~] $ > cat http-domains.txt | dsame -f https-domains.txt --no-trim -o pure-http-domains-only.txt
```
- Other just check with
```text
[~] $ > dsame --help
```
## Installation
Make sure you have Go installed, then run:

```bash
go install github.com/alb-soul/dsame@latest
```

Make sure $GOPATH/bin or $HOME/go/bin is in your PATH.
