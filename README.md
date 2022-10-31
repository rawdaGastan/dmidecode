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
