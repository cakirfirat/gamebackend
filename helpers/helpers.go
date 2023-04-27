package helpers

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
func Md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SendSms(phoneno, message string) {
	_ = godotenv.Load("../.env")

	url := "https://api.netgsm.com.tr/sms/send/get"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("usercode", os.Getenv("USERNAME"))
	_ = writer.WriteField("password", os.Getenv("PASSWORD"))
	_ = writer.WriteField("gsmno", phoneno)
	_ = writer.WriteField("message", message)
	_ = writer.WriteField("msgheader", os.Getenv("USERNAME"))
	err := writer.Close()
	CheckError(err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	CheckError(err)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	CheckError(err)
	defer res.Body.Close()

	// read response body
	scanner := bufio.NewScanner(res.Body)
	var response []byte
	for scanner.Scan() {
		response = append(response, scanner.Bytes()...)
	}

	// print response body
	fmt.Println(string(response))

}

var SECRET = []byte("super-secret-auth")

func CreateJwt() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 730).Unix()

	tokenStr, err := token.SignedString(SECRET)

	CheckError(err)

	return tokenStr, nil

}

func ValidateJwt(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Not authorized"))
				}
				return SECRET, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not authorized"))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
		}
	})
}

func CreateOtp() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(999999))
}

func Localizate(lang, text string) string {

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	switch lang {
	case "tr-tr":
		bundle.LoadMessageFile("../helpers/lang/tr-TR.json")
	case "en-en":
		bundle.LoadMessageFile("../helpers/lang/en-EN.json")
	default:
		bundle.LoadMessageFile("../helpers/lang/en-EN.json")
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	return localizer.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: text}})
}

func GenerateUUID() (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}
