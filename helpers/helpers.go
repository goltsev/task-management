package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/goltsev/task-management/entities"
)

func GenerateString(len int) (string, error) {
	buf := &strings.Builder{}
	buf.Grow(len)
	for i := 0; i < len; i++ {
		r := rand.Int63()%('z'-'A') + 'A'
		buf.WriteRune(rune(r))
	}
	return buf.String(), nil
}

func PrintProjectList(list []*entities.Project) *bytes.Buffer {
	buf := &bytes.Buffer{}
	for i, item := range list {
		fmt.Fprintf(buf, "%d : %s\n", i, item.Name)
	}
	return buf
}

func PrintProject(p *entities.Project) *bytes.Buffer {
	buf := &bytes.Buffer{}
	e := json.NewEncoder(buf)
	e.Encode(p)
	for _, item := range p.Columns {
		fmt.Fprintf(buf, "\t%v\n", *item)
	}
	return buf
}

func PrintColumn(c *entities.Column) *bytes.Buffer {
	buf := &bytes.Buffer{}
	e := json.NewEncoder(buf)
	e.Encode(c)
	for _, item := range c.Tasks {
		fmt.Fprintf(buf, "\t%v\n", *item)
	}
	return buf
}

func PrintColumnList(list []*entities.Column) *bytes.Buffer {
	buf := &bytes.Buffer{}
	for _, item := range list {
		buf.Write(PrintColumn(item).Bytes())
		buf.WriteByte('\n')
	}
	return buf
}

func PrintTasks(t []*entities.Task) *bytes.Buffer {
	buf := &bytes.Buffer{}
	for _, item := range t {
		fmt.Fprintf(buf, "%v\n", *item)
	}
	return buf
}

func PrintComments(cmts []*entities.Comment) *bytes.Buffer {
	buf := &bytes.Buffer{}
	for _, item := range cmts {
		fmt.Fprintf(buf, "%v\n", *item)
	}
	return buf
}
