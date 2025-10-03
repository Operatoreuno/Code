package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware configura gli headers CORS (Cross-Origin Resource Sharing).
//
// Headers configurati:
//   - Access-Control-Allow-Origin: origini permesse
//   - Access-Control-Allow-Credentials: abilita invio credenziali
//   - Access-Control-Allow-Methods: metodi HTTP permessi
//   - Access-Control-Allow-Headers: headers personalizzati permessi
//   - Access-Control-Expose-Headers: headers esposti al client (include X-Request-Id)
//   - Access-Control-Max-Age: durata cache preflight (3600 secondi)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // (TODO: Aggiornare in produzione)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-Request-Id")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-Request-Id")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		c.Next()
	}
}
