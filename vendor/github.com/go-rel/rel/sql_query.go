package rel

import "strings"

// SQLQuery allows querying using native query supported by database.
type SQLQuery struct {
	Statement string
	Values    []any
}

// Build Raw Query.
func (sq SQLQuery) Build(query *Query) {
	query.SQLQuery = sq
}

func (sq SQLQuery) String() string {
	var builder strings.Builder
	builder.WriteString("rel.SQL(\"")
	builder.WriteString(sq.Statement)
	builder.WriteString("\"")

	if len(sq.Values) != 0 {
		builder.WriteString(", ")
		builder.WriteString(fmtAnys(sq.Values))
	}

	builder.WriteString(")")
	return builder.String()
}

// SQL Query.
func SQL(statement string, values ...any) SQLQuery {
	return SQLQuery{
		Statement: statement,
		Values:    values,
	}
}
