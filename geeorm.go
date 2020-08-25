package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (engine *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	engine = &Engine{db: db}
	log.Info("connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("failed to close database")
		return
	}
	log.Info("close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db)
}
