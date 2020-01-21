# caicloud-formatting

## imports

Adjust the imports order to satisfy our [convention](https://github.com/caicloud/engineering/blob/master/guidelines/golang.md#order).

### Getting Started

Suppose you have `caicloud-formatting` and `config-admin` in your GOPATH and both of them are up to date, and you are currently in the root of the latter repository.

1. Simply run `go run $GOPATH/src/github.com/hezhizhen/caicloud-formatting/main.go` and boom! Everything's done.
2. Add a Makefile target which actually executes the command above, and `make` it.

```makefile
FORMATTING := $(GOPATH)/src/github.com/hezhizhen/caicloud-formatting
format:
    @go run $(FORMATTING)/main.go
```

### Usage

You can directly run the `main.go`, or you can specify the root, like `go run $GOPATH/src/github.com/hezhizhen/caicloud-formatting/main.go .`. Running it in the root of a repository is highly recommended, because files in the vendor should be skipped, but they are identified by checking if paths have prefixes of `vendor`.

A file will be skipped if it is:

* in the vendor directory
* a hidden one
* not a Go file (end with `.go`)

Besides, commented import lines will be removed to make imports clean.

### Git integration

Here's an idea about how to use `git pre-commit` to automatically format our code.

1. Download binary from `hezhizhen/caicloud-formatting` and save it to somewhere in the repository
2. Add a `pre-commit` file to `.git/hooks` directory and make it executable
3. Add and commit as usual

#### Download binary

1. Visit [here](https://github.com/hezhizhen/caicloud-formatting/releases) and download the latest one that is compatible with your OS
2. Save to somewhere in the repository (e.g.: `tools/caicloud-formatting`)
3. Give it executable permission if it isn't exectuable `chmod +x tools/caicloud-formatting`

#### Add `pre-commit`

Save the following code to a file named `pre-commit` in the `.git/hooks` directory:

```sh
./tools/caicloud-formatting .
if [[ -n $(git diff) ]]; then
    echo "Imports have been changed. Please commit again."
    exit 1
fi
exit 0
```

Don't forget to make it executable: `chmod +x .git/hooks/pre-commit`, Or you'll see such messages:

```text
hint: The 'pre-commit' hook was ignored because it's not set as executable.
hint: You can disable this warning with `git config advice.ignoredHook false`
```
