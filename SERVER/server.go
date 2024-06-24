package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type GLOBAL struct {
	A ARTIST
	D DATES
	L LOCATION
	R RELATION
	C int
}

type GLOBAL_MAIN struct {
	A []ARTIST
	D DATES
	L LOCATION
	R RELATION
}
type ARTIST struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"location"`
	ConcertDates string   `json:"concertDates"`
	Relation     string   `json:"relations"`
}

type LOCATION struct {
	Index []struct {
		Id       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

type DATES struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RELATION struct {
	Index []struct {
		Id       int                 `json:"id"`
		LocaDate map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var Global_main GLOBAL_MAIN
var Artists []ARTIST
var Location LOCATION

func main() {
	http.HandleFunc("/", Main_Page)
	http.HandleFunc("/artist", artist)
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8081", nil)
}

func Main_Page(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rawContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(rawContent, &Artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl := template.Must(template.ParseFiles("./static/src/index.html"))
	if r.Method == "POST" {
		val_input := r.FormValue("input")
		Filtre := SearchFilter(val_input)
		tpl.Execute(w, Filtre)
	} else {
		tpl.Execute(w, Artists)
	}
}

var Global GLOBAL

func artist(w http.ResponseWriter, r *http.Request) {
	IdClick := r.URL.Query().Get("id")

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + IdClick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rawContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(rawContent, &Global.A)

	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/dates/" + IdClick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rawContent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(rawContent, &Global.D)

	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/locations/" + IdClick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rawContent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(rawContent, &Global.L)

	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/relation/" + IdClick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rawContent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(rawContent, &Global.R)
	Global.C = len(Global.D.Dates)
	tpl := template.Must(template.ParseFiles("./static/src/info.html"))
	tpl.Execute(w, Global)
}

func SearchFilter(filter string) []ARTIST {
	endlist := []ARTIST{}
	for _, v := range Artists {
		for _, m := range v.Members {
			if strings.Contains(strings.ToUpper(m), strings.ToUpper(filter)) {
				endlist = append(endlist, v)
				break
			}
		}

		if strings.Contains(strings.ToUpper(v.Name), strings.ToUpper(filter)) {
			endlist = append(endlist, v)
		}

		if strings.Contains(strings.ToUpper(v.FirstAlbum), strings.ToUpper(filter)) {
			endlist = append(endlist, v)
		}

		if strings.Contains(strconv.Itoa(v.CreationDate), filter) {
			endlist = append(endlist, v)
		}
	}
	return endlist
}
