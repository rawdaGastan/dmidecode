# dmidecode

- A decoder for dmidecode output

## How to use

- Create a new dmidecode struct

```go
import "github.com/rawdaGastan/dmidecode/pkg"

dmi := pkg.NewDMIDecoder()
```

- Get the dmidecode output

```go
output, _ := dmi.GetDMIDecodeOutput()
```

- Then decode it

```go
dmi.Decode(output)
```

## Functions

- `dmiDecode.Decode()` &rarr; to get your decoded map
- `dmiDecode.GetSections()` &rarr; to get your sections' names
- `dmiDecode.GetSection( sectionKey )` &rarr; to get the content of the specified section key
- `dmiDecode.GetOptions( sectionKey )` &rarr; to get the options of the specified section key
- `dmiDecode.Get( sectionKey, optionKey )` &rarr; to get the string value of an option key inside a section
- `dmiDecode.GetList( sectionKey, optionKey )` &rarr; to get the list value of an option key inside a section

## Testing

Use this command to run the tests

```bash
go test -v ./...
```

```bash
task test
```

```bash
make test
```

## Coverage

- create a `coverage` folder
- Use this command to see the coverage

```bash
mkdir coverage
go test -v ./... -coverprofile=coverage/coverage.out
go tool cover -html=coverage/coverage.out -o coverage/coverage.html
```

```bash
task coverage
```

```bash
make coverage
```

- Open coverage/coverage.html to trace the coverage

## Benchmarks

Use this command to run the tests

```bash
go test -v ./... -bench=. -count 1 -benchtime=10s -benchmem -run=^#
```

```bash
make benchmarks
```

```bash
task benchmarks
```
