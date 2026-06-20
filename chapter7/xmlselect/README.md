## Search & print content of `h2` which is 2 level deep inside `div`s
With XML we can fetch `token-after-token` and need not to fetch the whole tree

Structure
```sh
<div>
  <div>
    <h2>xxx</h2>
  </div>
</div>
```
Then we get 
```sh
xxx
```
Run 
```sh
go mod init xmlselect
go build .
cat sample.xml | ./xmlselect div div h2
```

NOTE: Structure of the xml library
```sh
type Name struct { Local string }
// e.g. id="myNavbar" => Name="id", Value="myNavbar"
type Attr struct {
	Name Name
	Value string
}

type Token interface{}
type StartElement struct { // <div id="myNavbar" style="xxx">
	Name Name // "div"
	Attr []Attr
}
type EndElement struct { // </div>
	Name Name // "div"
}
type CharData []byte // what's inside <div>...</div>
type Comment []byte

type Decoder struct {...}
// method of Decoder to fetch the next Token
func (*Decoder) Token() (Token, error)

// create new Decoder from whatever has method Read()
func NewDecoder(io.Reader) *Decoder
```