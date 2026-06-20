## Query XML tree

Run
```sh
go mod init exercise7_17
go build
cat sample.xml | ./exercise7_17 div h2
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