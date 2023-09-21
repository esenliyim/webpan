package github

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type githubHookHeaders struct {
	Agent                      string `hook:"User-Agent"`
	Host                       string `hook:"Host"`
	ContentType                string `hook:"Content-Type"`
	GithubDelivery             string `hook:"X-Github-Delivery"`
	Event                      string `hook:"X-Github-Event"`
	HookId                     int    `hook:"X-Gihub-Hook-Id"`
	HookInstallationTargetId   int    `hook:"X-Gihub-Hook-Installation-Target-Id"`
	HookInstallationTargetType string `hook:"X-Gihub-Hook-Installation-Target-Type"`
	Signature                  string `hook:"X-Hub-Signature"`
	Signature256               string `hook:"X-Hub-Signature-256"`
}

type GithubHandler struct {
	name string
}

func New(name string) GithubHandler {
	return GithubHandler{
		name: name,
	}
}

func (gh GithubHandler) GetName() string {
	return gh.name
}

func (gh GithubHandler) Handle(c *fiber.Ctx) error {

	headers := parseHeaders(c.GetReqHeaders())

	if !validateSignature(headers.Signature256, c.Body()) {
		return errors.New("invalid signature")
	}

	switch headers.Event {
	case "push":
		return handlePush(c)
	default:
		return errors.New("invalid hook event")
	}
}

func parseHeaders(headers map[string]string) (parsed githubHookHeaders) {
	headersType := reflect.TypeOf(parsed)
	for i := 0; i < headersType.NumField(); i++ {
		field := headersType.Field(i)
		tag := field.Tag.Get("hook")
		fieldName := field.Name
		setHeaderField(&parsed, fieldName, headers[tag])
	}
	return
}

func setHeaderField(obj interface{}, key string, value string) {
	structVal := reflect.ValueOf(obj).Elem()
	structFieldVal := structVal.FieldByName(key)

	if !structFieldVal.IsValid() {
		return
	}

	if !structFieldVal.CanSet() {
		return
	}

	structFieldType := structFieldVal.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return
	}

	structFieldVal.Set(val)
}

func validateSignature(signature string, body []byte) bool {
	secret := os.Getenv("GITHUB_SECRET")
	hash := hmac.New(sha256.New, []byte(secret))
	normalized, err := normalizeBody(body)
	if err != nil {
		return false
	}
	hash.Write(normalized)

	return hex.EncodeToString(hash.Sum(nil)) == string(stripHashFromHeaderValue(signature))
}

func stripHashFromHeaderValue(headerVal string) []byte {
	return []byte(strings.TrimPrefix(headerVal, "sha256="))
}

func normalizeBody(body []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, body); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
