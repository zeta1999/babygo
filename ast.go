package main

var astCon string = "Con"
var astTyp string = "Typ"
var astVar string = "Var"
var astFun string = "Fun"
var astPkg string = "Pkg"

type signature struct {
	params  *astFieldList
	results *astFieldList
}

type astObject struct {
	Kind     string
	Name     string
	Decl     interface{} // *astValueSpec|*astFuncDecl|*astTypeSpec|*astField|*astAssignStmt
	Variable *Variable
}

type astExpr interface{}

type astField struct {
	Name   *astIdent
	Type   astExpr
	Offset int
}

type astFieldList struct {
	List []*astField
}

type astIdent struct {
	Name string
	Obj  *astObject
}

type astEllipsis struct {
	Elt astExpr
}

type astBasicLit struct {
	Kind  string // token.INT, token.CHAR, or token.STRING
	Value string
}

type astCompositeLit struct {
	Type astExpr
	Elts []astExpr
}

type astKeyValueExpr struct {
	Key   astExpr
	Value astExpr
}

type astParenExpr struct {
	X astExpr
}

type astSelectorExpr struct {
	X   astExpr
	Sel *astIdent
}

type astIndexExpr struct {
	X     astExpr
	Index astExpr
}

type astSliceExpr struct {
	X      astExpr
	Low    astExpr
	High   astExpr
	Max    astExpr
	Slice3 bool
}

type astCallExpr struct {
	Fun      astExpr   // function expression
	Args     []astExpr // function arguments; or nil
	Ellipsis bool
}

type astStarExpr struct {
	X astExpr
}

type astUnaryExpr struct {
	X  astExpr
	Op string
}

type astBinaryExpr struct {
	X  astExpr
	Y  astExpr
	Op string
}

type astTypeAssertExpr struct {
	X    astExpr
	Type astExpr // asserted type; nil means type switch X.(type)
}

// Type nodes
type astArrayType struct {
	Len astExpr
	Elt astExpr
}

type astStructType struct {
	Fields *astFieldList
}

type astInterfaceType struct {
	methods []string
}

type astFuncType struct {
	Params  *astFieldList
	Results *astFieldList
}

type astStmt interface{}

type astDeclStmt struct {
	Decl astDecl
}

type astExprStmt struct {
	X astExpr
}

type astIncDecStmt struct {
	X   astExpr
	Tok string
}

type astAssignStmt struct {
	Lhs     []astExpr
	Tok     string
	Rhs     []astExpr
	isRange bool
}

type astReturnStmt struct {
	Results []astExpr
	node    *nodeReturnStmt
}

type astBranchStmt struct {
	Tok        string
	Label      string
	currentFor astStmt
}

type astBlockStmt struct {
	List []astStmt
}

type astIfStmt struct {
	Init astStmt
	Cond astExpr
	Body *astBlockStmt
	Else astStmt
}

type astCaseClause struct {
	List []astExpr
	Body []astStmt
}

type astSwitchStmt struct {
	Tag  astExpr
	Body *astBlockStmt
	// lableExit string
}

type astTypeSwitchStmt struct {
	Assign astStmt
	Body   *astBlockStmt
	node   *nodeTypeSwitchStmt
}

type nodeReturnStmt struct {
	fnc *Func
}

type nodeTypeSwitchStmt struct {
	subject         astExpr
	subjectVariable *Variable
	assignIdent     *astIdent
	cases           []*TypeSwitchCaseClose
}

type TypeSwitchCaseClose struct {
	variable     *Variable
	variableType *Type
	orig         *astCaseClause
}

type astForStmt struct {
	Init      astStmt
	Cond      astExpr
	Post      astStmt
	Body      *astBlockStmt
	Outer     astStmt // outer loop
	labelPost string
	labelExit string
}

type astRangeStmt struct {
	Key       astExpr
	Value     astExpr
	X         astExpr
	Body      *astBlockStmt
	Outer     astStmt // outer loop
	labelPost string
	labelExit string
	lenvar    *Variable
	indexvar  *Variable
	Tok       string
}

type astImportSpec struct {
	Path string
}

type astValueSpec struct {
	Name  *astIdent
	Type  astExpr
	Value astExpr
}

type astTypeSpec struct {
	Name *astIdent
	Type astExpr
}

// Pseudo interface for *ast.Decl
// *astGenDecl | *astFuncDecl
type astDecl interface {
}

type astSpec interface{}

type astGenDecl struct {
	Spec astSpec // *astValueSpec | *TypeSpec
}

type astFuncDecl struct {
	Recv *astFieldList
	Name *astIdent
	Type *astFuncType
	Body *astBlockStmt
}

type astFile struct {
	Name       string
	Imports    []*astImportSpec
	Decls      []astDecl
	Unresolved []*astIdent
	scope      *astScope
}

type astScope struct {
	Outer   *astScope
	Objects []*objectEntry
}

type objectEntry struct {
	name string
	obj  *astObject
}

func astNewScope(outer *astScope) *astScope {
	return &astScope{
		Outer: outer,
	}
}

func (s *astScope) Insert(obj *astObject) {
	if s == nil {
		panic2(__func__, "s sholud not be nil\n")
	}

	s.Objects = append(s.Objects, &objectEntry{
		name: obj.Name,
		obj:  obj,
	})
}

func (s *astScope) Lookup(name string) *astObject {
	for _, oe := range s.Objects {
		if oe.name == name {
			return oe.obj
		}
	}

	return nil
}
