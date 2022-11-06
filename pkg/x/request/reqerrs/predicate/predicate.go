package predicate

import (
	"strconv"
	"strings"
)

func strPtr(s string) *string {
	return &s
}

type buildE struct {
	entity string
}

func E(entity string) buildE {
	return buildE{entity: entity}
}

// F is field
func (p buildE) F(field string) buildField {
	return buildField{buildE: p, fields: map[string]fieldval{}, currfield: field}
}

type fieldval struct {
	op    *string
	value *string
}

type buildField struct {
	buildE
	fields    map[string]fieldval
	currfield string
}

func (b buildField) General() Predicate {
	b.fields[b.currfield] = fieldval{}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) Eq(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) IEq(value int) Predicate {
	return b.Eq(strconv.Itoa(value))
}

func (b buildField) Ne(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is not"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) INe(value int) Predicate {
	return b.Ne(strconv.Itoa(value))
}

func (b buildField) Gt(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is greater than"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) IGt(value int) Predicate {
	return b.Gt(strconv.Itoa(value))
}

func (b buildField) Lt(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is less than"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) ILt(value int) Predicate {
	return b.Lt(strconv.Itoa(value))
}

func (b buildField) Ge(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is greater than or equal to"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) IGe(value int) Predicate {
	return b.Ge(strconv.Itoa(value))
}

func (b buildField) Le(value string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("is less than or equal to"), value: strPtr(value)}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) ILe(value int) Predicate {
	return b.Le(strconv.Itoa(value))
}

func (b buildField) In(value []string) Predicate {
	b.fields[b.currfield] = fieldval{op: strPtr("in"), value: strPtr(strings.Join(value, ", "))}
	return Predicate{buildE: b.buildE, fields: b.fields}
}

func (b buildField) IIn(value []int) Predicate {
	s := make([]string, len(value))
	for i, v := range value {
		s[i] = strconv.Itoa(v)
	}
	return b.In(s)
}

type Predicate struct {
	buildE
	fields map[string]fieldval
}

func (p Predicate) F(field string) buildField {
	fields := p.fields
	fields[field] = fieldval{}
	return buildField{buildE: p.buildE, fields: fields}
}

func (p Predicate) Describe() string {
	var s strings.Builder
	s.WriteString(p.entity)
	if len(p.fields) != 0 {
		s.WriteString(" where ")
	}
	var i int
	for field, val := range p.fields {
		if i > 0 {
			s.WriteString(" and ")
		}
		s.WriteString(field)
		if val.op != nil {
			s.WriteString(" ")
			s.WriteString(*val.op)
			s.WriteString(" ")
			s.WriteString(*val.value)
		}
		i++
	}
	return s.String()
}
