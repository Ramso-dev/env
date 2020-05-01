package env

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/Ramso-dev/srv"

	"net/http"
	"os"
)

//InitEnvVars sets the environment
func InitEnvVars(configuration interface{}) {

	switch os.Getenv("APP_ENV") {
	case "PROD":
		log.Println("Backend using PROD cloud environment")

		file, err := os.Open("env/prod.json")
		if err != nil {
			log.Println("Loading config from environment variables.")

			t := reflect.TypeOf(configuration)
			for i := 0; i < t.NumField(); i++ {
				key := t.Field(i).Name
				log.Println(key)
			}

		} else {
			log.Println("Loading config from json files.")
			decoder := json.NewDecoder(file)
			_ = decoder.Decode(&configuration)

			aMap, _ := InterfaceMap(configuration)
			for key, val := range aMap {
				os.Setenv(key, val)
				log.Println(key)
			}

		}
	case "CLOUDTEST":
		log.Println("Backend using TEST CLOUD environment")
		file, err := os.Open("env/cloudtest.json")
		if err != nil {
			log.Println("Loading config from environment variables.")

			t := reflect.TypeOf(configuration)
			for i := 0; i < t.NumField(); i++ {
				key := t.Field(i).Name
				log.Println(key, ":", os.Getenv(key))
			}

		} else {
			log.Println("Loading config from json files.")
			decoder := json.NewDecoder(file)
			_ = decoder.Decode(&configuration)

			aMap, _ := InterfaceMap(configuration)
			for key, val := range aMap {
				os.Setenv(key, val)
				log.Println(key, ":", val)
			}

		}
		os.Setenv("INSECURE", "true")
		os.Setenv("LOGS", "DEBUG")
	default:
		log.Println("Backend using LOCAL environment")

		file, err := os.Open("env/local.json")
		if err != nil {
			log.Println("Loading config from environment variables.")

			t := reflect.TypeOf(configuration)
			for i := 0; i < t.NumField(); i++ {
				key := t.Field(i).Name
				log.Println(key, ":", os.Getenv(key))
			}

		} else {
			log.Println("Loading config from json files.")
			decoder := json.NewDecoder(file)
			_ = decoder.Decode(&configuration)

			aMap, _ := InterfaceMap(configuration)
			for key, val := range aMap {
				os.Setenv(key, val)
				log.Println(key, ":", val)
			}

		}

		os.Setenv("INSECURE", "true")
		os.Setenv("LOGS", "DEBUG")
	}

	if os.Getenv("INSECURE") == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		srv.CustomClient = &http.Client{Timeout: 120 * time.Second, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

		log.Println("INSECURE is set to true. Warning: Insecure certificates are automatically goning to be trusted")
	}

}

func InterfaceMap(i interface{}) (map[string]string, error) {
	// Get type
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Map:

		//var newmap map[string]string = {}
		newmap := map[string]string{}
		// Get the value of the provided map
		v := reflect.ValueOf(i)

		// The "only" way of making a reflect.Type with interface{}
		//it := reflect.TypeOf((*interface{})(nil)).Elem()

		// Create the map of the specific type. Key type is t.Key(), and element type is it
		//m := reflect.MakeMap(reflect.MapOf(t.Key(), it))

		// Copy values to new map
		for _, mk := range v.MapKeys() {
			//log.Println(mk, v.MapIndex(mk))
			newmap[mk.String()] = v.MapIndex(mk).Interface().(string)
			//m.SetMapIndex(mk, v.MapIndex(mk))
		}

		return newmap, nil

	}

	return nil, errors.New("Unsupported type")
}
