package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// SecurityMiddleware configura gli headers di sicurezza HTTP.
// Applica le seguenti protezioni:
//   - FrameDeny: impedisce il rendering della pagina in iframe (protezione clickjacking)
//   - ContentTypeNosniff: impedisce al browser di sniffare il content-type
//   - BrowserXssFilter: abilita il filtro XSS del browser
//   - ReferrerPolicy: imposta la policy del referrer a "same-origin"
//   - STSIncludeSubdomains e STSPreload: configurazione HSTS per produzione
func SecurityMiddleware() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		ReferrerPolicy:       "same-origin",
		STSSeconds:           0, // Strict-Transport-Security (TODO: Aggiornare in produzione a 31536000)
		STSIncludeSubdomains: true,
		STSPreload:           true,
		// ContentSecurityPolicy: "default-src 'self'", (TODO: Aggiornare in produzione)
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
