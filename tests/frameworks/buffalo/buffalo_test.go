package buffalo

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/xenon007/go-admin/tests/common"
)

func TestBuffalo(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
