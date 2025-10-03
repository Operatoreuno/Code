package pkg

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Controlla se un valore vuoto è accettabile.
// Ritorna true se la validazione deve continuare, false se il campo è vuoto (valido o errore già gestito).
//
// Parameters:
//   - value: il valore da validare
//   - required: se true, il campo è obbligatorio
//
// Returns:
//   - bool: true se la validazione deve continuare, false altrimenti
//   - error: errore se il campo è obbligatorio ma vuoto, nil altrimenti
func checkRequired(value string, required bool) (bool, error) {
	if value == "" {
		if required {
			return false, errors.New("campo obbligatorio")
		}
		return false, nil // Opzionale e vuoto = valido, non continuare
	}
	return true, nil // Continua con le altre validazioni
}

// Controlla il formato dell'email utilizzando una regex.
//
// Parameters:
//   - value: l'email da validare
//   - required: se true, il campo è obbligatorio
//
// Returns:
//   - error: errore se l'email non è valida, nil altrimenti
func EmailValidator(value string, required bool) error {
	if shouldContinue, err := checkRequired(value, required); !shouldContinue {
		return err
	}

	// Validazione formato email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return errors.New("formato email non valido")
	}

	return nil
}

// Controlla che la password abbia almeno 8 caratteri, contenga almeno una maiuscola,
// una minuscola, un numero e non contenga spazi.
//
// Parameters:
//   - value: la password da validare
//   - required: se true, il campo è obbligatorio
//
// Returns:
//   - error: errore se la password non rispetta i requisiti, nil altrimenti
func PasswordValidator(value string, required bool) error {
	if shouldContinue, err := checkRequired(value, required); !shouldContinue {
		return err
	}

	if len(value) < 8 {
		return errors.New("deve contenere almeno 8 caratteri")
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, r := range value {
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		}
	}

	if !hasUpper {
		return errors.New("deve contenere almeno una lettera maiuscola")
	}
	if !hasLower {
		return errors.New("deve contenere almeno una lettera minuscola")
	}
	if !hasDigit {
		return errors.New("deve contenere almeno un numero")
	}

	if strings.Contains(value, " ") {
		return errors.New("non può contenere spazi")
	}

	return nil
}

// Verifica che la lunghezza della stringa sia compresa tra i valori min e max specificati.
//
// Parameters:
//   - value: la stringa da validare
//   - min: lunghezza minima (0 per nessun limite)
//   - max: lunghezza massima (0 per nessun limite)
//   - required: se true, il campo è obbligatorio
//
// Returns:
//   - error: errore se la stringa non rispetta i limiti di lunghezza, nil altrimenti
func StringValidator(value string, min, max int, required bool) error {
	if shouldContinue, err := checkRequired(value, required); !shouldContinue {
		return err
	}

	length := len(value)

	if min > 0 && length < min {
		return fmt.Errorf("deve contenere almeno %d caratteri", min)
	}
	if max > 0 && length > max {
		return fmt.Errorf("deve contenere al massimo %d caratteri", max)
	}

	return nil
}

// Controlla che il numero sia tra 8 e 16 caratteri e contenga solo un + iniziale opzionale, numeri e spazi.
//
// Parameters:
//   - value: il numero di telefono da validare
//   - required: se true, il campo è obbligatorio
//
// Returns:
//   - error: errore se il numero non è valido, nil altrimenti
func PhoneValidator(value string, required bool) error {
	if shouldContinue, err := checkRequired(value, required); !shouldContinue {
		return err
	}

	// Validazione lunghezza telefono
	if len(value) < 8 || len(value) > 16 {
		return errors.New("deve contenere tra 8 e 16 caratteri")
	}

	// Può contenere un solo + iniziale, poi solo numeri e spazi
	phoneRegex := regexp.MustCompile(`^\+?[\d\s]+$`)
	if !phoneRegex.MatchString(value) {
		return errors.New("formato numero di telefono non valido")
	}

	return nil
}
