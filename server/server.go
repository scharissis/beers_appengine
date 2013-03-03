package hello

import (
    "appengine"
    "appengine/datastore"
    "appengine/user"
    "encoding/json"
    "strconv"
    "net/http"
    "time"
    "log"
    "os"
    //"fmt"
)


var (
	index_file = "index.html"
	
	bootstrap_path = "lib/bootstrap"
	asset_path = "assets"
)

type User struct {
	Name string
	LogoutUrl string
}

type Beer struct {
	Name string
	Brewery string
	Country string
	Style string
	Score float64
	Comment string
	Timestamp time.Time
}


func init() {
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.HandleFunc("/store", store)
    http.HandleFunc("/test", index)
    
    http.HandleFunc("/api/user", api_get_user)
    http.HandleFunc("/api/beers", api_beers)
    http.HandleFunc("/api/beers/add", api_beers_add)
    
    // Static Resources
	http.Handle("/static/bootstrap/js/", http.StripPrefix("/static/bootstrap/", http.FileServer(http.Dir(bootstrap_path))))
	http.Handle("/static/bootstrap/css/", http.StripPrefix("/static/bootstrap/", http.FileServer(http.Dir(bootstrap_path))))
	http.Handle("/static/bootstrap/img/", http.StripPrefix("/static/bootstrap/", http.FileServer(http.Dir(bootstrap_path))))	
	http.Handle("/assets/js/", http.StripPrefix("/assets/", http.FileServer(http.Dir(asset_path))))
	http.Handle("/assets/partials/", http.StripPrefix("/assets/partials/", http.FileServer(http.Dir(asset_path+"/partials"))))
}

func login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }    
}

func logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
    	url, _ := user.LogoutURL(c, "/")
    	http.Redirect(w, r, url, http.StatusFound)
    }
   
}
/*
func root(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
    greetings := make([]Greeting, 0, 10)
    if _, err := q.GetAll(c, &greetings); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := guestbookTemplate.Execute(w, greetings); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func sign(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    log.Println(r.FormValue("content"));
    g := Greeting{
        Content: r.FormValue("content"),
        Date:    time.Now(),
    }
    if u := user.Current(c); u != nil {
        g.Author = u.String()
    }
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Greeting", nil), &g)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/", http.StatusFound)
}
*/
func index(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(index_file); err == nil {
		http.ServeFile(w, r, index_file)
	} else {
		log.Println("Error: ", index_file, " not found!")
	}
}

func store(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    log.Println(r.FormValue("content"));
    b := Beer{
        Name: r.FormValue("content"),
    }
    if u := user.Current(c); u != nil {
        //g.Author = u.String()
    }
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Beer", nil), &b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/beers", http.StatusFound)
}

func api_get_user(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	if u := user.Current(c); u != nil {
		url, _ := user.LogoutURL(c, "/")
	
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w,u.Email);
		user := User{
			Name: u.Email,
			LogoutUrl: url,
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}


func api_beers(w http.ResponseWriter, r *http.Request) {	
	c := appengine.NewContext(r)
	limit := 50
    q := datastore.NewQuery("Beer").Limit(limit)
    beers := make([]Beer, 0, limit)
    if _, err := q.GetAll(c, &beers); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(beers)
    
}

func api_beers_add(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    
    score, err := strconv.ParseFloat(r.FormValue("beerScore"), 64);
    if err != nil {
    	log.Println("Error: ", err);
    	score = 0.0;
    }
    
    b := Beer{
		Name: r.FormValue("beerName"),
		Brewery: r.FormValue("beerBrewery"),
		Country: r.FormValue("beerCountry"),
		Style: r.FormValue("beerStyle"),
		Score: score,
		Comment: r.FormValue("beerComment"),
		Timestamp:    time.Now(),
    }
    _, err = datastore.Put(c, datastore.NewIncompleteKey(c, "Beer", nil), &b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/beers", http.StatusOK)
}
