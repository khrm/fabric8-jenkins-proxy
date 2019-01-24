package configuration

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New("fileNot.yaml")
	assert.Error(t, err,
		"Error expected when file not found")
}

func TestConfig_GetPostgresHost(t *testing.T) {
	want := "localhost:9001"
	os.Setenv(pgHost, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresHost(), "Postgres Host don't match")
}

func TestConfig_GetPostgresPort(t *testing.T) {
	port := "999"
	want := 999
	os.Setenv(pgPort, port)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresPort(), "Postgres Port don't match")
}

func TestConfig_GetPostgresUser(t *testing.T) {
	want := "f8proxy"
	os.Setenv(pgUser, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresUser(), "Postgres User don't match")
}

func TestConfig_GetPostgresDatabase(t *testing.T) {
	want := "f8proxyDB"
	os.Setenv(pgDB, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresDatabase(), "Postgres DB don't match")
}

func TestConfig_GetPostgresPassword(t *testing.T) {
	want := "f8proxyPass"
	os.Setenv(pgPass, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresPassword(), "Postgres Password don't match")
}

func TestConfig_GetPostgresSSLMode(t *testing.T) {
	want := "enable"
	os.Setenv(pgSSLMode, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresSSLMode(), "Postgres SSL Mode don't match")
}

func TestConfig_GetPostgresConnectionTimeout(t *testing.T) {
	want := defaultPostgresConnectionTimeout
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresConnectionTimeout(), "Postgres DB don't match")
}

func TestConfig_GetPostgresConnectionMaxIdle(t *testing.T) {
	want := defaultPostgresConnectionMaxIdle
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresConnectionMaxIdle(),
		"Postgres Conn Max Idle don't match")
}

func TestConfig_GetPostgresConnectionMaxOpen(t *testing.T) {
	want := defaultPostgresConnectionMaxOpen
	c, _ := New("")
	assert.Equal(t, want, c.GetPostgresConnectionMaxOpen(),
		"Postgres Conn Max Open don't match")
}

func TestConfig_GetIdlerURL(t *testing.T) {
	want := "idler.openshift.io"
	os.Setenv(idlerAPIURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetIdlerURL(), "Idler URL don't match")
}

func TestConfig_GetAuthURL(t *testing.T) {
	want := "auth.openshift.io"
	os.Setenv(authURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetAuthURL(), "Auth URL don't match")
}

func TestConfig_GetTenantURL(t *testing.T) {
	want := "tenant.openshift.io"
	os.Setenv(f8TenantAPIURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetTenantURL(), "F8 Tenant URL don't match")
}

func TestConfig_GetWitURL(t *testing.T) {
	want := "wit.openshift.io"
	os.Setenv(witAPIURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetWitURL(), "WIT URL don't match")
}

func TestConfig_GetAuthToken(t *testing.T) {
	want := "secret"
	os.Setenv(authToken, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetAuthToken(), "Auth Token don't match")
}

func TestConfig_GetRedirectURL(t *testing.T) {
	want := "redirect.openshift.io"
	os.Setenv(redirectURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetRedirectURL(), "Redirect URL don't match")
}

func TestConfig_GetIndexPath(t *testing.T) {
	want := "indexPath"
	os.Setenv(indexPath, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetIndexPath(), "Index Path don't match")
}

func TestConfig_GetMaxRequestRetry(t *testing.T) {
	want := defaultMaxRequestRetry
	c, _ := New("")
	assert.Equal(t, want, c.GetMaxRequestRetry(), "Max Request Retry don't match")
}

func TestConfig_GetDebugMode(t *testing.T) {
	want := defaultDebugMode
	c, _ := New("")
	assert.Equal(t, want, c.GetDebugMode(), "Debug Mode don't match")
}

func TestConfig_GetHTTPSEnabled(t *testing.T) {
	want := defaultHTTPSEnabled
	c, _ := New("")
	assert.Equal(t, want, c.GetHTTPSEnabled(), "HTTPs Enable don't match")
}

func TestConfig_GetGatewayTimeout(t *testing.T) {
	want := defaultGatewayTimeout
	c, _ := New("")
	assert.Equal(t, want, c.GetGatewayTimeout(), "Gateway Timeout don't match")
}

func TestConfig_GetAllowedOrigins(t *testing.T) {
	origins := "openshift.io,local"
	want := []string{"openshift.io", "local"}
	os.Setenv(allowedOrigins, origins)
	c, _ := New("")
	assert.Equal(t, want, c.GetAllowedOrigins(), "Idler URL don't match")
}

func TestConfig_String(t *testing.T) {
	os.Clearenv()
	c, _ := New("")
	assert.True(t, strings.Contains(c.String(), "jc_index_path:"+defaultIndexPath),
		"Default IndexPath Config String doesn't match")
	assert.True(t, strings.Contains(c.String(),
		strings.ToLower(maxRequestRetry)+":"+
			strconv.Itoa(defaultMaxRequestRetry)),
		"Max Request retry Config String doesn't match")
	assert.True(t, strings.Contains(c.String(),
		strings.ToLower(enableHTTPS)+":"+
			"false"),
		"Enable HTTPs String doesn't match")
}
