# Test Plan - Postman API Testing

Questo documento descrive tutti i test da eseguire su Postman per validare le route e la gestione degli errori.

---

## Setup Postman

**Base URL**: `http://localhost:3000/api`

### Environment Variables
```
base_url = http://localhost:3000/api
```

---

## 1. POST /auth/signup - Registrazione Utente

### 1.1 ✅ Caso Success - Registrazione Valida

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe",
  "phone": "+391234567890"
}
```

**Expected Response:** `200 OK`
```json
{
  "message": "Utente registrato con successo",
  "data": {
    "id": "uuid",
    "email": "test@example.com",
    "name": "John",
    "surname": "Doe",
    "phone": "+39 123 456 7890",
    "role": "STANDARD",
    "is_active": false,
    "created_at": "2025-10-02T10:30:00Z"
  }
}
```

**Validazioni:**
- ✓ Status code = 200
- ✓ Response contiene `id` (UUID)
- ✓ Email è lowercase
- ✓ Phone è formattato correttamente
- ✓ Role default = "STANDARD"
- ✓ is_active = false

---

### 1.2 ✅ Caso Success - Senza Phone (Opzionale)

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "user2@example.com",
  "password": "Password123",
  "name": "Jane",
  "surname": "Smith"
}
```

**Expected Response:** `200 OK`
```json
{
  "message": "Utente registrato con successo",
  "data": {
    "id": "uuid",
    "email": "user2@example.com",
    "name": "Jane",
    "surname": "Smith",
    "phone": null,
    "role": "STANDARD",
    "is_active": false,
    "created_at": "2025-10-02T10:30:00Z"
  }
}
```

**Validazioni:**
- ✓ Status code = 200
- ✓ `phone` = null

---

### 1.3 ❌ Errore - Content-Type Mancante

**Request:**
```http
POST {{base_url}}/auth/signup

{
  "email": "test@example.com",
  "password": "Password123"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Formato JSON non valido",
  "data": {
    "details": "Content-Type deve essere application/json"
  }
}
```

**Validazioni:**
- ✓ Status code = 400
- ✓ errorCode = 1002 (INVALID_REQUEST)

---

### 1.4 ❌ Errore - Content-Type Sbagliato

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: text/plain

{
  "email": "test@example.com"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Formato JSON non valido",
  "data": {
    "details": "Content-Type deve essere application/json"
  }
}
```

---

### 1.5 ❌ Errore - Body Vuoto

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Formato JSON non valido",
  "data": {
    "details": "EOF"
  }
}

```

---

### 1.6 ❌ Errore - JSON Malformato

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123"
  "name": "John"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Formato JSON non valido",
  "data": {
    "details": "invalid character '\"' after object key:value pair"
  }
}
```

---

### 1.7 ❌ Errore - Email Mancante

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "email",
      "message": "is required"
    }
  ]
}
```

---

### 1.8 ❌ Errore - Email Formato Invalido

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "notanemail",
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "email",
      "message": "invalid email format"
    }
  ]
}
```

---

### 1.9 ❌ Errore - Email con Spazi (Deve Essere Trimmata)

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "   test@example.com   ",
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `200 OK` (email viene trimmata automaticamente)
```json
{
  "data": {
    "email": "test@example.com"
  }
}
```

**Validazioni:**
- ✓ Email salvata senza spazi

---

### 1.10 ❌ Errore - Email Uppercase (Deve Essere Lowercase)

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "USER@EXAMPLE.COM",
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `200 OK` (email convertita in lowercase)
```json
{
  "data": {
    "email": "user@example.com"
  }
}
```

---

### 1.11 ❌ Errore - Password Mancante

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "is required"
    }
  ]
}
```

---

### 1.12 ❌ Errore - Password Troppo Corta

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Pass1",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "must be at least 8 characters long"
    }
  ]
}
```

---

### 1.13 ❌ Errore - Password Senza Uppercase

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "must contain at least one uppercase letter"
    }
  ]
}
```

---

### 1.14 ❌ Errore - Password Senza Lowercase

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "PASSWORD123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "must contain at least one lowercase letter"
    }
  ]
}
```

---

### 1.15 ❌ Errore - Password Senza Numeri

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "must contain at least one digit"
    }
  ]
}
```

---

### 1.16 ❌ Errore - Password con Spazi

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Pass word 123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "password",
      "message": "cannot contain spaces"
    }
  ]
}
```

---

### 1.17 ❌ Errore - Name Mancante

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "name",
      "message": "is required"
    }
  ]
}
```

---

### 1.18 ❌ Errore - Name Troppo Lungo (>50 caratteri)

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "JohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohn",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "name",
      "message": "must be at most 50 characters long"
    }
  ]
}
```

---

### 1.19 ❌ Errore - Surname Mancante

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "surname",
      "message": "is required"
    }
  ]
}
```

---

### 1.20 ❌ Errore - Phone Formato Invalido

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe",
  "phone": "abc123"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "phone",
      "message": "invalid phone number format"
    }
  ]
}
```

---

### 1.21 ❌ Errore - Phone Troppo Corto

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe",
  "phone": "123"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "phone",
      "message": "must be between 8 and 15 characters"
    }
  ]
}
```

---

### 1.22 ❌ Errore - Phone Troppo Lungo

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe",
  "phone": "1234567890123456"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "phone",
      "message": "must be between 8 and 15 characters"
    }
  ]
}
```

---

### 1.23 ❌ Errore - Email Già Esistente

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "existing@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `409 Conflict`
```json
{
  "errorCode": 1001,
  "message": "Email già in uso",
  "data": {
    "field": "email"
  }
}
```

**Validazioni:**
- ✓ Status code = 409
- ✓ errorCode = 1001 (ALREADY_EXISTS)

---

### 1.24 ❌ Errore - Multipli Errori di Validazione

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "notanemail",
  "password": "pass",
  "name": "",
  "surname": ""
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "email",
      "message": "invalid email format"
    },
    {
      "field": "password",
      "message": "must be at least 8 characters long"
    },
    {
      "field": "name",
      "message": "is required"
    },
    {
      "field": "surname",
      "message": "is required"
    }
  ]
}
```

**Validazioni:**
- ✓ Tutti gli errori sono elencati
- ✓ Array con 4 errori

---

### 1.25 ❌ Errore - XSS Attack in Name

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "<script>alert('xss')</script>",
  "surname": "Doe"
}
```

**Expected Response:** `200 OK` (HTML rimosso da sanitizzazione)
```json
{
  "data": {
    "name": ""
  }
}
```

**Validazioni:**
- ✓ Script tag rimosso
- ✓ Name è stringa vuota dopo sanitizzazione
- ✓ Dovrebbe poi fallire validazione "is required"

---

### 1.26 ❌ Errore - Campo Extra Non Richiesto

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John",
  "surname": "Doe",
  "extraField": "should be ignored"
}
```

**Expected Response:** `200 OK` (campo extra ignorato)
```json
{
  "data": {
    "email": "test@example.com",
    "name": "John",
    "surname": "Doe"
  }
}
```

**Validazioni:**
- ✓ Campo extra ignorato
- ✓ Registrazione completata con successo

---

## 2. GET /auth/health - Health Check

### 2.1 ✅ Caso Success

**Request:**
```http
GET {{base_url}}/auth/health
```

**Expected Response:** `200 OK`
```json
{
  "message": "Auth service is running",
  "data": {
    "service": "auth",
    "status": "healthy"
  }
}
```

**Validazioni:**
- ✓ Status code = 200
- ✓ Nessun body richiesto

---

### 2.2 ❌ Errore - Metodo Sbagliato

**Request:**
```http
POST {{base_url}}/auth/health
```

**Expected Response:** `404 Not Found` o `405 Method Not Allowed`


---

## 3. GET /auth/test-panic - Test Panic Handler

### 3.1 ✅ Panic Gestito

**Request:**
```http
GET {{base_url}}/auth/test-panic
```

**Expected Response:** `500 Internal Server Error`
```json
{
  "errorCode": 1004,
  "message": "Si è verificato un errore critico"
}
```

**Validazioni:**
- ✓ Status code = 500
- ✓ errorCode = 1004 (INTERNAL_SERVER_ERR)
- ✓ Server non crasha

---

## 4. Test Generici - Middleware

### 4.1 ❌ Route Non Esistente

**Request:**
```http
GET {{base_url}}/nonexistent
```

**Expected Response:** `404 Not Found`

---

### 4.2 ✅ CORS Headers

**Request:**
```http
OPTIONS {{base_url}}/auth/signup
Origin: http://localhost:3000
```

**Expected Response:** `200 OK`

**Headers da verificare:**
- `Access-Control-Allow-Origin: *` (o dominio specifico)
- `Access-Control-Allow-Methods: GET, POST, PUT, PATCH, DELETE, OPTIONS`
- `Access-Control-Allow-Headers: Content-Type, Authorization`

---

### 4.3 ✅ Security Headers

**Request:**
```http
GET {{base_url}}/auth/health
```

**Headers da verificare:**
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security: max-age=31536000; includeSubDomains`

---

## 5. Edge Cases

### 5.1 ❌ Request Troppo Grande

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "A very long string..." (>1MB)
}
```

**Expected Response:** `413 Payload Too Large` o `400 Bad Request`

---

### 5.2 ❌ Unicode e Emoji in Campi

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123",
  "name": "John 😀",
  "surname": "Doe"
}
```

**Expected Response:** `200 OK` (emoji mantenuto)
```json
{
  "data": {
    "name": "John 😀"
  }
}
```
---

### 5.3 ❌ Null Values

**Request:**
```http
POST {{base_url}}/auth/signup
Content-Type: application/json

{
  "email": null,
  "password": "Password123",
  "name": "John",
  "surname": "Doe"
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "errorCode": 1002,
  "message": "Dati di input non validi",
  "data": [
    {
      "field": "email",
      "message": "is required"
    }
  ]
}
```

---

## 6. Performance Tests

### 6.1 Concurrent Requests

**Test:**
- Invia 10 richieste simultanee di signup con email diverse
- Verifica che tutte completino con successo

**Expected:**
- ✓ Tutte le richieste ritornano 200
- ✓ Tempo di risposta < 200ms per richiesta

---

### 6.2 Database Connection

**Test:**
- Disconnetti database
- Invia richiesta di signup

**Expected Response:** `500 Internal Server Error`
```json
{
  "errorCode": 1004,
  "message": "Si è verificato un errore imprevisto"
}
```

---