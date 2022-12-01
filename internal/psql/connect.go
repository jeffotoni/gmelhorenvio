package psql

// Back-End in Go server
// @jeffotoni

import (
	"context"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/jeffotoni/gmelhorenvio/internal/fmts"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ConnGroup struct {
	Log    *DbConnection
	Monge  *DbConnection
	Design *DbConnection
}

type config struct {
	Name, Host, User, Pass, Port, Ssl, Source string
}

var ConfigLog = config{
	DB_NAME_LOG,
	DB_HOST_LOG,
	DB_USER_LOG,
	DB_PASSWORD_LOG,
	DB_PORT_LOG,
	DB_SSL,
	source_log,
}

type DbConnection struct {
	Once       *sync.Once
	Ctx        *context.Context
	ReopenConn chan bool
	IsWaiting  bool
	mu         sync.Mutex
	Conn       *pgxpool.Pool
	Config     config
}

func NewConn(ctx *context.Context, conf config) *DbConnection {
	conn := &DbConnection{
		Once:       &sync.Once{},
		Ctx:        ctx,
		ReopenConn: make(chan bool),
		Config:     conf,
	}

	err := conn.Connect()
	if err != nil {
		log.Println("error on connecting to the Db:", err.Error())
		if DB_WAIT_START {
			conn.WaitForConnection()
		}
	}

	return conn
}

func (c *DbConnection) Connect() error {
	var err error
	c.Once.Do(func() {
		var connStr string
		if len(dbCertServerCaPath) > 0 &&
			len(dbCertClientPath) > 0 &&
			len(dbCertKeyPath) > 0 {
			connStr = fmts.ConcatStr(
				"postgres://",
				url.UserPassword(c.Config.User, c.Config.Pass).String(),
				"@", c.Config.Host,
				":", c.Config.Port,
				"/", c.Config.Name,
				"?sslmode=", c.Config.Ssl,
				"&sslrootcert=", dbCertServerCaPath,
				"&sslcert=", dbCertClientPath,
				"&sslkey=", dbCertKeyPath)
		} else {
			connStr = fmts.ConcatStr(
				"postgres://",
				url.UserPassword(c.Config.User, c.Config.Pass).String(),
				"@", c.Config.Host,
				":", c.Config.Port,
				"/", c.Config.Name)
			// "?sslmode=", c.Config.Ssl)
		}
		db, e := pgxpool.Connect(*c.Ctx, connStr)
		if e != nil {
			err = e
			return
		}
		c.Conn = db
	})
	return err
}

func (c *DbConnection) WaitForConnection() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.IsWaiting = true
	log.Println("WAITING FOR CONNECTION")
	log.Println("Start")
	for {
		log.Println("... trying to restore Db connection ...")
		if c.Conn == nil || c.Conn.Ping(*c.Ctx) != nil {
			time.Sleep(time.Second * time.Duration(DB_RECONNECT))
			c.Once = &sync.Once{}
			c.Connect()
			continue
		}
		break
	}
	log.Println("End")
	log.Println("CONNECTION RESTORED")
	c.IsWaiting = false
}

func (c *DbConnection) LoopCheckConnection() {
	go func() {
		for { // for range chan?
			select {
			case reopen := <-c.ReopenConn:
				if reopen {
					go c.WaitForConnection()
				}
			}
		}
	}()
}

func (c *DbConnection) reConnect() {
	if c.IsWaiting {
		return
	}
	go func() {
		c.ReopenConn <- true
	}()
}

func (c *DbConnection) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if c.Conn == nil || c.Conn.Ping(*c.Ctx) != nil {
		c.reConnect()
		return nil, ErrDatabaseDown
	}
	return c.Conn.Exec(ctx, sql, args...)
}

func (c *DbConnection) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if c.Conn == nil || c.Conn.Ping(*c.Ctx) != nil {
		c.reConnect()
		return nil, ErrDatabaseDown
	}
	return c.Conn.Query(ctx, sql, args...)
}

func (c *DbConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) (pgx.Row, error) {
	if c.Conn == nil || c.Conn.Ping(*c.Ctx) != nil {
		c.reConnect()
		return nil, ErrDatabaseDown
	}
	return c.Conn.QueryRow(ctx, sql, args...), nil
}

func (c *DbConnection) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	if c.Conn == nil || c.Conn.Ping(*c.Ctx) != nil {
		c.reConnect()
		return nil, ErrDatabaseDown
	}
	return c.Conn.QueryFunc(ctx, sql, args, scans, f)
}
