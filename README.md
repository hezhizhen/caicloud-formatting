# caicloud-formatting

## imports

Adjust the imports order to satisfy our [convention](https://github.com/caicloud/engineering/blob/master/guidelines/golang.md#order).

### Getting Started

Suppose you have `go-common` and `config-admin` in your GOPATH and both of them are up to date, and you are currently in the root of the latter repository.

1. Simply run `go run ../go-common/format/main.go` and boom! Everything's done.
2. Add a Makefile target which actually executes the command above, and `make` it.

```makefile
GOCOMMON := $(ROOT)/vendor/github.com/caicloud/go-common
format:
    @go run $(GOCOMMON)/format/main.go
```

ATTENTION: to make such a Makefile target work, you need to add the following line to the top of `Gopkg.toml` file:

```toml
required = ["github.com/caicloud/go-common/format"]
```

And then execute `dep ensure -update github.com/caicloud/go-common`. Just a few minutes later, you'll see `format` package in your `vendor`.

### Usage

You can directly run the `main.go`, or you can specify the root, like `go run ../go-common/format/main.go .`. Running it in the root is highly recommended, because files in the vendor should be skipped, but it is identified by check if the path has a prefix of `vendor`.

It will be skipped if a file is:

* in the vendor directory
* a hidden one
* not a Go file (end with `.go`)

Besides, it will remove commented import lines to make imports clean.

### Git integration

Here's an idea about how to use `git pre-commit` to automatically format our code.

1. Make sure that `go-common/format` package is in your vendor directory
2. Add a `pre-commit` file to `.git/hooks` directory and make it executable
3. Add and commit as usual

#### Vendor `go-common/format`

1. Add a line to the top of `Gopkg.toml` file: `required = ["github.com/caicloud/go-common/format"]` and save it
2. Execute `dep ensure -update github.com/caicloud/go-common`

#### Add `pre-commit`

Save the following code to a file named `pre-commit` in the `.git/hooks` directory:

```sh
go run vendor/github.com/caicloud/go-common/format/main.go .
if [[ -n $(git diff) ]]; then
    echo "Imports have been re-ordered. Please add and commit again."
    exit 1
fi
exit 0
```

Don't forget to make it executable: `chmod +x .git/hooks/pre-commit`, Or you'll see such messages:

```text
hint: The 'pre-commit' hook was ignored because it's not set as executable.
hint: You can disable this warning with `git config advice.ignoredHook false`
```
