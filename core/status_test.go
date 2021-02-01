package core

import (
	"github.com/golang/mock/gomock"
	"github.com/nuts-foundation/nuts-node/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestNewStatusEngine_Routes(t *testing.T) {
	t.Run("Registers a single route for listing all engines", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		echo := mock.NewMockEchoRouter(ctrl)

		echo.EXPECT().GET("/status/diagnostics", gomock.Any())
		echo.EXPECT().GET("/status", gomock.Any())

		NewStatusEngine(NewSystem()).Routes(echo)
	})
}

func TestNewStatusEngine_Cmd(t *testing.T) {
	system := NewSystem()
	t.Run("Cmd returns a cobra command", func(t *testing.T) {
		e := NewStatusEngine(system).Cmd
		assert.Equal(t, "diagnostics", e.Name())
	})

	t.Run("Executed Cmd writes diagnostics to prompt", func(t *testing.T) {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		NewStatusEngine(system).Cmd.Execute()

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		assert.Equal(t, "", string(out))
	})
}

func TestNewStatusEngine_Diagnostics(t *testing.T) {
	system := NewSystem()
	system.RegisterEngine(NewStatusEngine(system))
	system.RegisterEngine(NewLoggerEngine())
	system.RegisterEngine(NewMetricsEngine())

	t.Run("diagnostics() returns engine list", func(t *testing.T) {
		ds := NewStatusEngine(system).Diagnostics()
		assert.Len(t, ds, 1)
		assert.Equal(t, "Registered engines", ds[0].Name())
		assert.Equal(t, "Status,Logging,Metrics", ds[0].String())
	})

	t.Run("diagnosticsOverview() renders text output of diagnostics", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		echo := mock.NewMockContext(ctrl)

		echo.EXPECT().String(http.StatusOK, "Status\n\tRegistered engines: Status,Logging,Metrics\nLogging\n\tverbosity: ")

		(&status{system}).diagnosticsOverview(echo)
	})
}

func TestStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	echo := mock.NewMockContext(ctrl)

	echo.EXPECT().String(http.StatusOK, "OK")

	statusOK(echo)
}