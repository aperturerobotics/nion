package testutil

import (
	"bytes"
	"testing"

	"github.com/rgraphql/nion/qtree"
	"github.com/rgraphql/nion/schema"
	proto "github.com/rgraphql/rgraphql"
)

// MockSchemaSrc is a mock schema source code fragment.
var MockSchemaSrc = `
type Person {
	name: String
	height: Int
}

type RootQuery {
	allPeople: [Person]
	names: [String]!
}

schema {
	query: RootQuery
}
`

// BuildMockTree builds a mock schema and query tree.
func BuildMockTree(t *testing.T, schemaFrags ...string) (
	*schema.Schema,
	*qtree.QueryTreeNode,
	<-chan *proto.RGQLQueryError,
) {
	var schemaBuf bytes.Buffer
	if len(schemaFrags) == 0 {
		_, _ = schemaBuf.WriteString(MockSchemaSrc)
	} else {
		for _, frag := range schemaFrags {
			_, _ = schemaBuf.WriteString(frag)
			_, _ = schemaBuf.WriteString("\n")
		}
	}

	sch, err := schema.Parse(schemaBuf.String())
	if err != nil {
		t.Fatal(err.Error())
	}
	errCh := make(chan *proto.RGQLQueryError, 10)
	qt, err := sch.BuildQueryTree(errCh)
	if err != nil {
		t.Fatal(err.Error())
	}
	return sch, qt, errCh
}
