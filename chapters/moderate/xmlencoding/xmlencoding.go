package xmlencoding

/*
XML Encoding: Since XML is very verbose, you are given a way of encoding it where each tag gets
mapped to a pre-defined integer value. The language/grammar is as follows:
Element --> Tag Attribute s END Children END
Attribut e --> Tag Value
END --> 0
Tag --> some predefined mapping to in t
Value --> strin g value
For example, the following XML might be converted into the compressed string below (assuming a
mapping of family -> 1, person ->2 , firstName -> 3, lastName ->4 , state -> 5).
<family lastName="McDowell" state="CA">
•cperson firstName="Gayle">Some Message</person>
</family>
Becomes:
1 4 McDowell 5 CA 0 2 3 Gayle 0 Some Message 0 0
Write code to print the encoded version of an XML element (passed in Element and Attribute
objects).
*/

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Tag   string
	Value string
}

type Element struct {
	Tag        string
	Attributes []Attribute
	Children   []*Element
	Value      string
}

func Encode(element *Element, mapping map[string]int) string {
	var b strings.Builder
	encodeElement(&b, element, mapping)
	return strings.TrimSpace(b.String())
}

func encodeElement(b *strings.Builder, e *Element, m map[string]int) {
	fmt.Fprintf(b, "%d ", m[e.Tag])
	for _, a := range e.Attributes {
		fmt.Fprintf(b, "%d %s ", m[a.Tag], a.Value)
	}
	fmt.Fprint(b, "0 ")
	if len(e.Children) > 0 {
		for _, c := range e.Children {
			encodeElement(b, c, m)
		}
	} else if e.Value != "" {
		fmt.Fprintf(b, "%s ", e.Value)
	}
	fmt.Fprint(b, "0 ")
}
