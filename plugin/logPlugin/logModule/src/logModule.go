package src

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
)

func init() {
	t := reflect.TypeOf((*CLogModule)(nil))
	if !t.Implements(reflect.TypeOf((*logModule.ILogModule)(nil)).Elem()) {
		log.Fatal("ILogModule is not implemented by CLogModule")
	}

	logModule.ModuleType = t.Elem()
	logModule.ModuleName = filepath.Join(logModule.ModuleType.PkgPath(), logModule.ModuleType.Name())
	logModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CLogModule{}).Update).Pointer()).Name()
}

type CLogModule struct {
	ark.Module
	// other data
	logger *logrus.Logger
}

func (logModule *CLogModule) Init() error {
	logModule.logger = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: true,
		Level:        logrus.WarnLevel,
	}

	return nil
}

// ------------------- logrus options -------------------
func (logModule *CLogModule) SetFormatter(formatter logrus.Formatter) {
	logModule.logger.SetFormatter(formatter)
}

func (logModule *CLogModule) SetOutput(out io.Writer) {
	logModule.logger.SetOutput(out)
}

func (logModule *CLogModule) SetReportCaller(include bool) {
	logModule.logger.SetReportCaller(include)
}

func (logModule *CLogModule) SetLevel(level logrus.Level) {
	logModule.logger.SetLevel(level)
}

func (logModule *CLogModule) AddHook(hook logrus.Hook) {
	logModule.logger.AddHook(hook)
}

func (logModule *CLogModule) GetLogger() *logrus.Logger {
	return logModule.logger
}
