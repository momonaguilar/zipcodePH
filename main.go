package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ZipCode struct {
	Zipcode      string `json:"zipcode"`
	Area         string `json:"area"`
	ProvinceCity string `json:"provinceCity"`
}

func main() {
	http.HandleFunc("/zipcode", handleGetZip)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("ERROR: Error reading config file. Err:", err)
	}
	viper.SetDefault("APP.PORT", "8081")
	PORT, ok := viper.Get("APP.PORT").(string)

	if !ok {
		log.Fatalln("ERROR: Invalid type assertion")
	}

	fmt.Println("INFO: Starting web server at port", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalln("ERROR: Unable to listen and serve, Err:", err)
	}

}

func handleGetZip(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.URL.Path != "/zipcode" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method not supported.", http.StatusNotFound)
		return
	}

	key := r.FormValue("key")

	//open file
	f, err := os.Open("data/zipcode.csv")
	if err != nil {
		log.Fatal(err)
	}

	//close when program ends
	defer f.Close()

	//read values
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	resultZipcode := []ZipCode{}

	for _, line := range data {
		if line[0] == key || strings.EqualFold(line[1], key) {
			z := ZipCode{
				Zipcode:      line[0],
				Area:         line[1],
				ProvinceCity: line[2],
			}
			resultZipcode = append(resultZipcode, z)
		}
	}

	result, err := json.Marshal(resultZipcode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(rw, string(result))

	elapsed := time.Since(start)

	logStr := "Key:" + key + ", Result:" + string(result) + ". Completed after " + fmt.Sprint(elapsed.Milliseconds()) + "ms"
	fmt.Println("INFO:", logStr)

}
