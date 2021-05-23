package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/meetsoni15/CoWin-Vaccine-Notifier/sounds"
)

var (
	// For how many days you want to search
	num_of_days = 0
	// Age
	age = 25
	// Pincode
	pincode = []string{"462003"}
	// For notification
	PlaySound = false
	//Running script
	Ticker = 5 * time.Minute
)

//Centers -> Centers array
type Centers struct {
	Center []Center `json:"centers"`
}

type Center struct {
	CenterID int       `json:"center_id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	State    string    `json:"state_name"`
	District string    `json:"district_name"`
	Block    string    `json:"block_name"`
	Pincode  int       `json:"pincode"`
	Lat      float64   `json:"lat"`
	Long     float64   `json:"long"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	FeeType  string    `json:"fee_type"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	SessionID              string   `json:"session_id"`
	Date                   string   `json:"date"`
	AvailableCapacity      int      `json:"available_capacity"`
	MinAgeLimit            int      `json:"min_age_limit"`
	Vaccine                string   `json:"vaccine"`
	Slots                  []string `json:"slots"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
}

func main() {
	//For initial time run once
	CheckSlotsAvailability()
	runtimeTicker := time.NewTicker(Ticker)
	for {
		select {
		case <-runtimeTicker.C:
			print("\n uptime")
			CheckSlotsAvailability()
		}
	}

}

func CheckSlotsAvailability() {
	counter := 0
	dates := DateRange()
	for _, zipcode := range pincode {
		for _, date := range dates {
			CenterArr := SendReq(zipcode, date)
			if len(CenterArr.Center) > 0 {
				for _, center := range CenterArr.Center {
					if len(center.Sessions) > 0 {
						for _, session := range center.Sessions {
							if session.MinAgeLimit <= age && session.AvailableCapacity > 0 {
								print("Pincode: " + zipcode)
								print("Available on: {}" + date)
								print("\n", center.Name)
								print("\n", center.Block)
								print("\n Price: ", center.FeeType)
								print("\n Availablity : ", session.AvailableCapacity)

								if session.Vaccine != "" {
									print("\n Vaccine type : ", session.Vaccine)
								}
								counter++
							}
						}
					}
				}
			} else {
				print("\n Pincode: " + zipcode)
				print("\n Not Available on: {}" + date)
				print("\n No Response!")
				print("\n")
			}
		}
	}

	if counter == 0 {
		print("\n No Vaccination slot available!")
	} else {
		if PlaySound {
			sounds.PlaySound()
		}
		print("\n Search Completed!")
	}
}

func DateRange() []string {
	date_arr := []string{}
	start := time.Now()
	end := start.AddDate(0, 0, num_of_days)
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		date_arr = append(date_arr, d.Format("02-01-2006"))
	}

	return date_arr
}

func SendReq(pincode, date string) Centers {
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=" + pincode + "&date=" + date + ""
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return Centers{}
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.76 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Centers{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Centers{}
	}

	var c = Centers{}
	json.Unmarshal(body, &c)
	return c
}
