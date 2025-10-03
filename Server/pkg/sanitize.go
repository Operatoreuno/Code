package pkg

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// Policy molto restrittiva che rimuove tutto l'HTML, lasciando solo testo.
var xssPolicy = bluemonday.StrictPolicy()

// Identifica emoji Unicode.
var emojiRegex = regexp.MustCompile(`[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{1F1E0}-\x{1F1FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]|[\x{1F900}-\x{1F9FF}]|[\x{1FA00}-\x{1FA6F}]|[\x{1FA70}-\x{1FAFF}]|[\x{FE00}-\x{FE0F}]|[\x{1F004}]|[\x{1F0CF}]|[\x{1F18E}]|[\x{3030}]|[\x{2B50}]|[\x{2B55}]|[\x{203C}]|[\x{2049}]|[\x{25AA}]|[\x{25AB}]|[\x{25B6}]|[\x{25C0}]|[\x{25FB}]|[\x{25FC}]|[\x{25FD}]|[\x{25FE}]`)

// Sanitizza tutti i campi stringa di un oggetto struct.
// Applica le seguenti trasformazioni:
//   - Rimuove HTML/XSS utilizzando bluemonday StrictPolicy
//   - Rimuove spazi all'inizio e alla fine
//   - Normalizza le email in minuscolo
//   - Rimuove emoji dai campi name e surname
//   - Ignora i campi "password"

// Parameters:
//   - object: puntatore a una struct da sanitizzare
func Sanitize(object any) {
	v := reflect.ValueOf(object).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag != "" {
			if parts := strings.Split(jsonTag, ","); len(parts) > 0 {
				fieldName = parts[0]
			}
		}

		if strings.Contains(strings.ToLower(fieldName), "password") {
			continue
		}

		if field.Kind() == reflect.String {
			original := field.String()
			cleaned := xssPolicy.Sanitize(original)
			cleaned = strings.TrimSpace(cleaned)

			// Rimuovi emoji da name e surname
			if strings.Contains(strings.ToLower(fieldName), "name") ||
				strings.Contains(strings.ToLower(fieldName), "surname") {
				cleaned = emojiRegex.ReplaceAllString(cleaned, "")
				cleaned = strings.TrimSpace(cleaned)
			}

			// Se è un campo email, converti in lowercase
			if strings.Contains(strings.ToLower(fieldName), "email") {
				cleaned = strings.ToLower(cleaned)
			}

			field.SetString(cleaned)
		}
	}
}
