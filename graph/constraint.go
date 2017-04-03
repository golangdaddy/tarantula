package graph

type Constraint interface {
	SQL() string
	Value() string
}

//

type AUTO_INCREMENT struct {
	Val string
}

func (c AUTO_INCREMENT) SQL() string   { return "AUTO_INCREMENT" }
func (c AUTO_INCREMENT) Value() string { return c.Val }

//

type PRIMARY_KEY struct {
	Val string
}

func (c PRIMARY_KEY) SQL() string   { return "PRIMARY KEY" }
func (c PRIMARY_KEY) Value() string { return c.Val }

type FOREIGN_KEY struct {
	Val string
}

func (c FOREIGN_KEY) SQL() string   { return "FOREIGN KEY" }
func (c FOREIGN_KEY) Value() string { return c.Val }

type NULL struct{}

func (c NULL) SQL() string   { return "" }
func (c NULL) Value() string { panic("CONSTRAINT MISUSE: " + c.SQL()); return "" }

type NOT_NULL struct{}

func (c NOT_NULL) SQL() string   { return "NOT NULL" }
func (c NOT_NULL) Value() string { panic("CONSTRAINT MISUSE: " + c.SQL()); return "" }

type UNIQUE struct{}

func (c UNIQUE) SQL() string   { return "UNIQUE" }
func (c UNIQUE) Value() string { panic("CONSTRAINT MISUSE: " + c.SQL()); return "" }

type CHECK struct{}

func (c CHECK) SQL() string   { return "CHECK" }
func (c CHECK) Value() string { panic("CONSTRAINT MISUSE: " + c.SQL()); return "" }

type DEFAULT struct{}

func (c DEFAULT) SQL() string   { return "DEFAULT" }
func (c DEFAULT) Value() string { panic("CONSTRAINT MISUSE: " + c.SQL()); return "" }

type Constraints struct {
	PRIMARY_KEY PRIMARY_KEY
	FOREIGN_KEY FOREIGN_KEY
	NOT_NULL    NOT_NULL
	DEFAULT     DEFAULT
	UNIQUE      UNIQUE
	CHECK       CHECK
}

func (c *Constraints) List() []Constraint {

	return []Constraint{
		c.PRIMARY_KEY,
		c.FOREIGN_KEY,
		c.NOT_NULL,
		c.DEFAULT,
		c.UNIQUE,
		c.CHECK,
	}
}
