package configuration

import (
	"fmt"
	"strings"
	"time"

	"github.com/fabric8-services/fabric8-jenkins-proxy/internal/util"
	"github.com/pkg/errors"
	errs "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = log.WithFields(log.Fields{"component": "configuration"})
)

const (
	pgHost        = "JC_POSTGRES_HOST"
	pgPort        = "JC_POSTGRES_PORT"
	pgDB          = "JC_POSTGRES_DATABASE"
	pgUser        = "JC_POSTGRES_USER"
	pgPass        = "JC_POSTGRES_PASSWORD"
	pgSSLMode     = "JC_POSTGRES_SSL_MODE"
	pgConnTimeout = "JC_POSTGRES_CONNECTION_TIMEOUT"
	pgConnMaxIdle = "JC_POSTGRES_CONNECTION_MAX_IDLE"
	pgConnMaxOpen = "JC_POSTGRES_CONNECTION_MAX_OPEN"

	idlerAPIURL    = "JC_IDLER_API_URL"
	authURL        = "JC_AUTH_URL"
	authToken      = "JC_AUTH_TOKEN"
	f8TenantAPIURL = "JC_F8TENANT_API_URL"
	witAPIURL      = "JC_WIT_API_URL"

	redirectURL     = "JC_REDIRECT_URL"
	indexPath       = "JC_INDEX_PATH"
	maxRequestRetry = "JC_MAX_REQUEST_RETRY"
	debugMode       = "JC_DEBUG_MODE"
	enableHTTPS     = "JC_ENABLE_HTTPS"
	gatewayTimeout  = "JC_GATEWAY_TIMEOUT"
	allowedOrigins  = "JC_ALLOWED_ORIGINS"

	defaultPostgresSSLMode           = "disable"
	defaultPostgresConnectionTimeout = 5
	defaultPostgresConnectionMaxIdle = -1
	defaultPostgresConnectionMaxOpen = -1
	defaultIndexPath                 = "/opt/fabric8-jenkins-proxy/index.html"
	defaultMaxRequestRetry           = 10
	defaultDebugMode                 = false
	defaultHTTPSEnabled              = false
	defaultGatewayTimeout            = 5 * time.Second
	defaultAllowedOrigins            = "https://*openshift.io,https://localhost:*,http://localhost:*"
)

// New creates a configuration reader object using a configurable configuration
// file path.
func New(configFilePath string) (Configuration, error) {
	c := Config{
		v: viper.New(),
	}
	c.v.AutomaticEnv()
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.v.SetTypeByDefaultValue(true)
	c.setConfigDefaults()

	if configFilePath != "" {
		c.v.SetConfigType("yaml")
		c.v.SetConfigFile(configFilePath)
		err := c.v.ReadInConfig() // Find and read the config file
		if err != nil {           // Handle errors reading the config file
			return nil, errs.Errorf("Fatal error config file: %s \n", err)
		}
	}

	colErr := c.verify()
	if !colErr.Empty() {
		for _, err := range colErr.Errors {
			logger.Error(err)
		}
		return nil, errors.Wrap(colErr.ToError(),
			"some config variables are missing or invalid")
	}
	return &c, nil
}

// Config encapsulates the Viper configuration registry which stores the
// configuration data in-memory.
type Config struct {
	v *viper.Viper
}

func (c *Config) setConfigDefaults() {

	c.v.SetDefault(pgSSLMode, defaultPostgresSSLMode)
	c.v.SetDefault(pgConnTimeout, defaultPostgresConnectionTimeout)
	c.v.SetDefault(pgConnMaxIdle, defaultPostgresConnectionMaxIdle)
	c.v.SetDefault(pgConnMaxOpen, defaultPostgresConnectionMaxOpen)

	c.v.SetDefault(indexPath, defaultIndexPath)
	c.v.SetDefault(maxRequestRetry, defaultMaxRequestRetry)
	c.v.SetDefault(debugMode, defaultDebugMode)
	c.v.SetDefault(enableHTTPS, defaultHTTPSEnabled)
	c.v.SetDefault(gatewayTimeout, defaultGatewayTimeout)
	c.v.SetDefault(allowedOrigins, defaultAllowedOrigins)

}

// GetPostgresHost returns the postgres host as set via default, config file, or environment variable.
func (c *Config) GetPostgresHost() string {
	return c.v.GetString(pgHost)
}

// GetPostgresPort returns the postgres port as set via default, config file, or environment variable.
func (c *Config) GetPostgresPort() int {
	return c.v.GetInt(pgPort)
}

// GetPostgresUser returns the postgres user as set via default, config file, or environment variable.
func (c *Config) GetPostgresUser() string {
	return c.v.GetString(pgUser)
}

// GetPostgresDatabase returns the postgres database as set via default, config file, or environment variable.
func (c *Config) GetPostgresDatabase() string {
	return c.v.GetString(pgDB)
}

// GetPostgresPassword returns the postgres password as set via default, config file, or environment variable.
func (c *Config) GetPostgresPassword() string {
	return c.v.GetString(pgPass)
}

// GetPostgresSSLMode returns the postgres sslmode as set via default, config file, or environment variable.
func (c *Config) GetPostgresSSLMode() string {
	return c.v.GetString(pgSSLMode)
}

// GetPostgresConnectionTimeout returns the postgres connection timeout as set via default, config file, or environment variable.
func (c *Config) GetPostgresConnectionTimeout() int {
	return c.v.GetInt(pgConnTimeout)
}

// GetPostgresConnectionMaxIdle returns the number of connections that should be keept alive in the database connection pool at
// any given time. -1 represents no restrictions/default behavior.
func (c *Config) GetPostgresConnectionMaxIdle() int {
	return c.v.GetInt(pgConnMaxIdle)
}

// GetPostgresConnectionMaxOpen returns the max number of open connections that should be open in the database connection pool.
// -1 represents no restrictions/default behavior.
func (c *Config) GetPostgresConnectionMaxOpen() int {
	return c.v.GetInt(pgConnMaxOpen)
}

// GetIdlerURL returns the Idler API URL as set via default, config file, or environment variable.
func (c *Config) GetIdlerURL() string {
	return c.v.GetString(idlerAPIURL)
}

// GetAuthURL returns the Auth API URL as set via default, config file, or environment variable.
func (c *Config) GetAuthURL() string {
	return c.v.GetString(authURL)
}

// GetTenantURL returns the F8 Tenant API URL as set via default, config file, or environment variable.
func (c *Config) GetTenantURL() string {
	return c.v.GetString(f8TenantAPIURL)
}

// GetWitURL returns the WIT API URL as set via default, config file, or environment variable.
func (c *Config) GetWitURL() string {
	return c.v.GetString(witAPIURL)
}

// GetAuthToken returns the Auth token as set via default, config file, or environment variable.
func (c *Config) GetAuthToken() string {
	return c.v.GetString(authToken)
}

// GetRedirectURL returns the redirect url to be passed to Auth as set via default, config file, or environment variable.
func (c *Config) GetRedirectURL() string {
	return c.v.GetString(redirectURL)
}

// GetIndexPath returns the path to loading page template as set via default, config file, or environment variable.
func (c *Config) GetIndexPath() string {
	return c.v.GetString(indexPath)
}

// GetMaxRequestRetry returns the number of retries for webhook request forwarding as set via default, config file,
// or environment variable.
func (c *Config) GetMaxRequestRetry() int {
	return c.v.GetInt(maxRequestRetry)
}

// GetDebugMode returns if debug mode should be enabled as set via default, config file, or environment variable.
func (c *Config) GetDebugMode() bool {
	return c.v.GetBool(debugMode)
}

// GetHTTPSEnabled returns if https should be enabled as set via default, config file, or environment variable.
func (c *Config) GetHTTPSEnabled() bool {
	return c.v.GetBool(enableHTTPS)
}

// GetGatewayTimeout returns the duration within which the reverse-proxy expects
// a response.
func (c *Config) GetGatewayTimeout() time.Duration {
	return c.v.GetDuration(gatewayTimeout)
}

// GetAllowedOrigins returns string containing allowed origins separated with ", "
func (c *Config) GetAllowedOrigins() []string {
	return strings.Split(c.v.GetString(allowedOrigins), ",")
}

// String returns string representation of configuration
func (c *Config) String() string {
	all := c.v.AllSettings()
	for k := range all {
		// don't echo tokens or secret
		if strings.Contains(k, "TOKEN") ||
			strings.Contains(k, "token") {
			all[k] = "***"
		}

		if strings.Contains(k, "PASSWORD") ||
			strings.Contains(k, "password") {
			all[k] = "***"
		}
	}
	return fmt.Sprintf("%v", all)
}

// verify checks whether all needed config options are set.
func (c *Config) verify() util.MultiError {
	config := c.v.AllSettings()
	var errors util.MultiError
	for k, v := range config {
		switch strings.ToUpper(k) {
		case idlerAPIURL:
			continue
		case f8TenantAPIURL:
			continue
		case witAPIURL:
			continue
		case redirectURL:
			continue
		case authURL:
			errors.Collect(util.IsURL(v, k))
		case pgHost:
			continue
		case pgDB:
			continue
		case pgUser:
			continue
		case pgPass:
			continue
		case authToken:
			continue
		case allowedOrigins:
			continue
		case indexPath:
			errors.Collect(util.IsNotEmpty(v, k))
		}
	}
	return errors
}
