package logs

import 	(
		"fmt"
		"sync"
		"time"
		"errors"
		"strings"
		"runtime"
		"reflect"
		"encoding/json"
		//
		"golang.org/x/net/context"
		"cloud.google.com/go/logging"
		"google.golang.org/appengine/log"
	)

func fn() string {

	names := []string{}

	pc, _, _, ok := runtime.Caller(5)

	if ok { names = append(names, runtime.FuncForPC(pc).Name()) }

	pc, _, _, ok = runtime.Caller(4)

	if ok { names = append(names, runtime.FuncForPC(pc).Name()) }

	pc, _, _, ok = runtime.Caller(3)

	if ok { names = append(names, runtime.FuncForPC(pc).Name()) }

	return strings.Join(names, "/")
}

// creates a new logging client
func NewClient(googleProjectName string, ctx context.Context) *LogClient {

	client, err := logging.NewClient(ctx, googleProjectName); if err != nil { panic(err) }

	return &LogClient{
		ctx,
		client,
		map[string]*Logger{},
		sync.RWMutex{},
	}
}

type LogClient struct {
	ctx context.Context
	client *logging.Client
	loggers map[string]*Logger
	sync.RWMutex
}

func (lc *LogClient) Close() {
	lc.client.Close()
}

// creates a new logger based on the input name
func (lc *LogClient) NewLogger(logFuncNames ...string) *Logger {

	var logFuncName string

	if len(logFuncNames) == 0 {

		pc, _, _, _ := runtime.Caller(1)
		logFuncName = runtime.FuncForPC(pc).Name()

		// remove illegal chars from log name
		logFuncName = strings.Replace(logFuncName, "*", "", -1)
		logFuncName = strings.Replace(logFuncName, "(", "", -1)
		logFuncName = strings.Replace(logFuncName, ")", "", -1)

	} else {

		logFuncName = logFuncNames[0]

	}

	lc.Lock()
	defer lc.Unlock()

	if lc.loggers[logFuncName] == nil {
		lc.loggers[logFuncName] = &Logger{
			lc.ctx,
			lc.client,
			lc.client.Logger(logFuncName),
		}
	}

	return lc.loggers[logFuncName]
}

type Logger struct {
	ctx context.Context
	client *logging.Client
	logger *logging.Logger
}

func (lg *Logger) Flush() {
	lg.logger.Flush()
}

// creates and executes a logging entry
func (lg *Logger) Log(msg interface{}, severity logging.Severity) {

	fn := fn()

	// show only last n chars of fn
	n := 64
	l := len(fn)
	if l > n { fn = fn[l-n:] }

	payload := fmt.Sprintf(
		"%v:%s: %v",
		time.Now().UTC().Unix(),
		fn,
		msg,
	)

	switch severity {

		case logging.Error:

			log.Errorf(lg.ctx, payload)

		case logging.Debug:

			log.Debugf(lg.ctx, payload)

	}

}

// Debug log
func (lg *Logger) Debug(msg interface{}) { lg.Log(msg, logging.Debug) }

// Debug log with formatting
func (lg *Logger) Debugf(s string, args ...interface{}) {

	msg := fmt.Sprintf(s, args...)

	lg.Log(msg, logging.Debug)

}

// Error log
func (lg *Logger) NewError(msg string) error {

	lg.Log(msg, logging.Error)

	return errors.New(msg)
}

// Error log with formatting
func (lg *Logger) NewErrorf(s string, args ...interface{}) error {

	msg := fmt.Sprintf(s, args...)

	lg.Log(msg, logging.Error)

	return errors.New(msg)
}

// error log
func (lg *Logger) Error(err error) bool {

	if err != nil { lg.Log(err, logging.Error) }

	return err != nil
}

// critical log
func (lg *Logger) Panic(msg interface{}) {

	if msg == nil { return }

	lg.Log(msg, logging.Critical)

	lg.logger.Flush()

	panic(msg)
}

// critical log
func (lg *Logger) Fatal(msg interface{}) {

	if msg == nil { return }

	lg.Log(msg, logging.Critical)

	lg.logger.Flush()

	panic(msg)
}

// type assertion fail log
func (lg *Logger) Reflect(e interface{}) {

	msg := "REFLECT VALUE IS NIL"

	if e != nil { msg = "REFLECT VALUE IS "+reflect.TypeOf(e).String() }

	lg.Log(msg, logging.Error)
}

// debug json output log
func (lg *Logger) DebugJSON(i interface{}) {

	b, err := json.Marshal(i); if err != nil { lg.Error(err); return }

	lg.Log(string(b), logging.Debug)
}

// error json output log
func (lg *Logger) ErrorJSON(i interface{}) {

	b, err := json.Marshal(i); if err != nil { lg.Error(err); return }

	lg.Log(string(b), logging.Error)
}
