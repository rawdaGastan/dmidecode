# dmidecode

- A decoder for dmidecode output

## How to use

- Create a new dmidecode struct

```go
import "github.com/rawdaGastan/dmidecode/pkg"

dmiDecode := pkg.NewDMIDecoder()
```

- Then decode it

```go
dmiDecode.Decode()
```

## Functions

- `parser.Decode()` &rarr; to get your decoded map
- `parser.GetSections()` &rarr; to get your sections' names
- `parser.GetSection( sectionKey )` &rarr; to get the content of the specified section key
- `parser.GetOptions( sectionKey )` &rarr; to get the options of the specified section key
- `parser.Get( sectionKey, optionKey )` &rarr; to get the string value of an option key inside a section
- `parser.GetList( sectionKey, optionKey )` &rarr; to get the list value of an option key inside a section

## Testing

Use this command to run the tests

```bash
go test -v ./...
```
