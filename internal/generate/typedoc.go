package main

// Kind ...
type Kind uint

const (
	ExternalModule       Kind = 1 << iota
	Module                    //2
	Enumeration               //4
	_                         //8
	EnumerationMember         //16
	Variable                  //32
	Function                  //64
	Class                     //128
	Interface                 //256
	Constructor               //512
	Property                  //1024
	Method                    //2048
	CallSignature             //4096
	IndexSignature            //8192
	ConstructorSignature      //16384
	Parameter                 //32768
	TypeLiteral               //65536
	TypeParameter             //131072
	_                         // 262144
	_                         // 524288
	_                         // 1048576
	_                         // 2097152
	TypeAlias                 // 4194304
)

// Flags ...
type Flags struct {
	IsExported bool `json:"isExported,omitempty"`
	IsExternal bool `json:"isExternal,omitempty"`
	IsOptional bool `json:"isOptional,omitempty"`
}

// Group ...
type Group struct {
	Title    string `json:"title"`
	Kind     int    `json:"kind"`
	Children []int  `json:"children,omitempty"`
}

// Decolaration ...
type Decolaration struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Kind           Kind         `json:"kind"`
	KindString     string       `json:"kindString"`
	Flags          *Flags       `json:"flags,omitempty"`
	IndexSignature []*Signature `json:"indexSignature,omitempty"`
	Signatures     []*Signature `json:"signatures,omitempty"`
}

// Type ...
type Type struct {
	Type         string        `json:"type"` // intrinsic/array/reference/union/reflection
	Name         string        `json:"name"`
	Types        []*Type       `json:"types,omitempty"`       // for union
	ID           int           `json:"id,omitempty"`          // for reference
	ElementType  *Type         `json:"elementType,omitempty"` // for array
	Decolaration *Decolaration `json:"declaration,omitempty"` // for reference inline
	Value        string        `json:"value,omitempty"`       // for stringLiteral
}

// Comment ...
type Comment struct {
	Text string `json:"text"`
}

// Param ...
type Param struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Kind       Kind     `json:"kind"`
	KindString string   `json:"kindString"`
	Flags      *Flags   `json:"flags,omitempty"`
	Type       *Type    `json:"type,omitempty"`
	Comment    *Comment `json:"comment,omitempty"`
}

// Source ...
type Source struct {
	FileName  string `json:"fileName"`
	Line      int    `json:"line"`
	Character int    `json:"character"`
}

// Signature ...
type Signature struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Kind       Kind     `json:"kind"`
	KindString string   `json:"kindString"`
	Flags      *Flags   `json:"flags,omitempty"`
	Params     []*Param `json:"parameters,omitempty"`
	Type       *Type    `json:"type,omitempty"` //return type
}

// Define ...
type Define struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	OriginalName  string       `json:"originalName"`
	Kind          Kind         `json:"kind"`
	KindString    string       `json:"kindString"`
	Flags         *Flags       `json:"flags,omitempty"`
	Children      []*Define    `json:"children,omitempty"`
	Groups        []*Group     `json:"groups,omitempty"`
	Signatures    []*Signature `json:"signatures,omitempty"`
	Type          *Type        `json:"type,omitempty"`
	InheritedFrom *Type        `json:"inheritedFrom,omitempty"`
	Sources       []*Source    `json:"sources,omitempty"`
	ExtendedTypes []*Type      `json:"extendedTypes,omitempty"`
}

/* type memo

以下のパターンがある。
- Class Def / 1-file
- Module Def / 1-file　<- サブパッケージとする

以下のunion定義型は「*CSSRule」とする。
"type": {
	"type": "union",
	"types": [
		{
			"type": "reference",
			"name": "CSSRule",
			"id": 7792
		},
		{
			"type": "intrinsic",
			"name": "null"
		}
	]
}

colors, vectorsのTypeLiteralに関しては既存の定義型に置き換える

*/
