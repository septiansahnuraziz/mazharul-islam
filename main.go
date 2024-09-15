package main

import (
	"github.com/mazharul-islam/internal/cmd"
)

// @termsOfService				http://your-term-of-service-url.com
// @contact.name				Eraspace
// @contact.url				eraspace.com
// @license.name				Apache 2.0
// @license.ur					http://www.apache.org/licenses/LICENSE-2.0.html
// @query.collection.			format multi
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService				http://swagger.io/terms/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cmd.Execute()
}
