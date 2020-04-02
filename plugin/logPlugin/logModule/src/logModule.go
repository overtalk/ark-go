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
	t := reflect.TypeOf((*AFCLogModule)(nil))
	if !t.Implements(reflect.TypeOf((*logModule.AFILogModule)(nil)).Elem()) {
		log.Fatal("AFILogModule is not implemented by AFCLogModule")
	}

	logModule.ModuleType = t.Elem()
	logModule.ModuleName = filepath.Join(logModule.ModuleType.PkgPath(), logModule.ModuleType.Name())
	logModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCLogModule{}).Update).Pointer()).Name()
}

type AFCLogModule struct {
	ark.AFCModule
	// other data
	logger *logrus.Logger
}

func (logModule *AFCLogModule) Init() error {
	logModule.logger = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: true,
		Level:        logrus.WarnLevel,
	}

	return nil
}

// ------------------- logrus options -------------------
func (logModule *AFCLogModule) SetFormatter(formatter logrus.Formatter) {
	logModule.logger.SetFormatter(formatter)
}

func (logModule *AFCLogModule) SetOutput(out io.Writer) {
	logModule.logger.SetOutput(out)
}

func (logModule *AFCLogModule) SetReportCaller(include bool) {
	logModule.logger.SetReportCaller(include)
}

func (logModule *AFCLogModule) SetLevel(level logrus.Level) {
	logModule.logger.SetLevel(level)
}

func (logModule *AFCLogModule) AddHook(hook logrus.Hook) {
	logModule.logger.AddHook(hook)
}

func (logModule *AFCLogModule) GetLogger() *logrus.Logger {
	return logModule.logger
}
