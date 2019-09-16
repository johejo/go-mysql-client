package gomysqlclient

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/c-bata/go-prompt"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/olekukonko/tablewriter"
	"github.com/xwb1989/sqlparser"
)

type Cli struct {
	db     *sql.DB
	config *Config
}

func NewCli(c *Config) (*Cli, error) {
	db, err := sql.Open("mysql", c.String())
	if err != nil {
		return nil, err
	}
	return &Cli{
		db:     db,
		config: c,
	}, nil
}

func (c *Cli) execute(q string) (*Results, error) {
	if q == "" {
		return nil, nil
	}
	rows, err := c.db.Query(q)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results [][]string
	for rows.Next() {
		rs := make([]sql.NullString, len(columns))
		rsp := make([]interface{}, len(columns))
		for i := range rs {
			rsp[i] = &rs[i]
		}
		if err = rows.Scan(rsp...); err != nil {
			break
		}
		_rs := make([]string, len(columns))
		for i := range rs {
			_rs[i] = rs[i].String
		}
		results = append(results, _rs)
	}
	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &Results{
		Columns: columns,
		Rows:    results,
	}, nil
}

func (c *Cli) Run(execute ...string) error {
	if len(execute) == 0 || execute[0] == "" {
		return c.RunPrompt()
	}
	for _, e := range execute {
		results, err := c.execute(e)
		if err != nil {
			return err
		}
		fmt.Println(results.String())
	}
	return nil
}

func (c *Cli) executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit", "quit":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		q := in
		if _, err := sqlparser.Parse(q); err != nil {
			fmt.Println(err)
			return
		}
		results, err := c.execute(q)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(results.String())
		return
	}
}

func (c *Cli) RunPrompt() error {
	tables, err := c.getTables()
	if err != nil {
		return err
	}
	suggests := make([]prompt.Suggest, 0, len(tables)+len(mysqlWords)+len(mysqlFuncs))
	for _, w := range mysqlWords {
		suggests = append(suggests, prompt.Suggest{Text: w})
	}
	for _, f := range mysqlFuncs {
		suggests = append(suggests, prompt.Suggest{Text: f})
	}
	for _, table := range tables {
		suggests = append(suggests, prompt.Suggest{Text: table})
	}
	completer := func(document prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggests, document.GetWordBeforeCursor(), true)
	}
	p := prompt.New(c.executor, completer, prompt.OptionPrefix("> "), prompt.OptionHistory(make([]string, 0)))
	p.Run()
	return nil
}

func (c *Cli) getTables() ([]string, error) {
	rows, err := c.db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	var tables []string
	for rows.Next() {
		var table string
		if err = rows.Scan(&table); err != nil {
			break
		}
		tables = append(tables, table)
	}
	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func QueriesFromReader(r io.Reader) []string {
	s := bufio.NewScanner(r)
	commentRegexp := regexp.MustCompile(`--.*$`)
	var lines []string
	for s.Scan() {
		line := s.Text()
		if line == "\n" {
			break
		}
		q := commentRegexp.ReplaceAllString(line, "")
		if q != "" {
			lines = append(lines, q)
		}
	}
	return strings.Split(strings.Join(lines, ""), ";")
}

type Results struct {
	Columns []string
	Rows    [][]string
}

func (r *Results) String() string {
	buf := bytes.NewBufferString("")
	table := tablewriter.NewWriter(buf)
	table.SetHeader(r.Columns)
	table.AppendBulk(r.Rows)
	table.Render()
	return buf.String()
}
