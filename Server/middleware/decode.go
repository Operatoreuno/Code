package middleware

import (
	"encoding/json"
	"errors"
	apiErrors "oniplu/errors"
	"oniplu/pkg"

	"github.com/gin-gonic/gin"
)

// Dodifica e sanitizza il body JSON di una richiesta.
// Verifica che:
//   - Il Content-Type sia application/json
//   - Il body non sia vuoto
//   - Il JSON sia valido
//
// Dopo la decodifica, applica automaticamente la sanitizzazione tramite pkg.Sanitize.
//
// Parameters:
//   - ctx: contesto Gin della richiesta
//   - body: puntatore alla struct dove decodificare il JSON
//
// Returns:
//   - error: errore di validazione o decodifica, nil se successo
func decodeJSON[T any](ctx *gin.Context, body *T) error {
	contentType := ctx.Request.Header.Get("Content-Type")

	if contentType != "application/json" {
		return errors.New("content-Type deve essere application/json")
	}

	if ctx.Request.Body == nil || ctx.Request.ContentLength == 0 {
		return errors.New("body della richiesta mancante")
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(body); err != nil {
		return err
	}

	pkg.Sanitize(body)
	return nil
}

// Wrapper middleware che decodifica e sanitizza automaticamente il body JSON.
// Semplifica la gestione delle richieste POST/PUT/PATCH fornendo il body già decodificato e sanitizzato.
//
// Il middleware:
//   - Decodifica il JSON dal body della richiesta
//   - Sanitizza i dati tramite pkg.Sanitize
//   - Passa il body decodificato al controller
//   - Gestisce automaticamente gli errori di decodifica
//
// Esempio d'uso:
//
//	router.POST("/endpoint", middleware.WithJSONBody(func(ctx *gin.Context, body *MyStruct) error {
//	    // body è già decodificato e sanitizzato
//	    return nil
//	}))
//
// Parameters:
//   - controller: funzione handler che riceve il contesto Gin e il body decodificato
//
// Returns:
//   - gin.HandlerFunc: middleware da utilizzare con Gin router
func Decode[T any](controller func(*gin.Context, *T) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body T

		// Decode e Sanitizza JSON
		if err := decodeJSON(ctx, &body); err != nil {
			apiErr := apiErrors.BadRequestError(err.Error(), apiErrors.INVALID_REQUEST)
			ctx.Error(apiErr)
			ctx.Abort()
			return
		}

		// Chiama il controller con il body decodificato
		if err := controller(ctx, &body); err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}
	}
}
