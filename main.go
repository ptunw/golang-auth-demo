package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/context"
    "strings"
)

func main() {
    router := mux.NewRouter()
    amw := authMiddleWare{}
    router.Use(amw.ValidateJWT)
    router.HandleFunc("/login", LoginUser).Methods("POST")
    router.HandleFunc("/me", FetchLoggedInUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

type User struct {
    Email string `json:"email"`
    Name string `json:"name"`
}

type Credential struct {
    Email string `json:"email"`
    Password string `json: "password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    jsonDecoder := json.NewDecoder(r.Body)
    var cred Credential
    err := jsonDecoder.Decode(&cred)
    if err != nil {
        panic(err)
    }

    //These values should come from database
    const validEmail = "ak.sskg@gmail.com"
    const validPassword = "password" //password should have been hashed
    if cred.Email == validEmail && cred.Password == validPassword {
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ "email": cred.Email })
        tokenString, error := token.SignedString([]byte("mysecretkey"))
        if error != nil {
            panic(error)
        }
        w.Write([]byte(tokenString))
    }
}

func FetchLoggedInUser(w http.ResponseWriter, r *http.Request) {
    
    loggedInUserEmail := context.Get(r, "email").(string)
    u := &User{Email: loggedInUserEmail, Name: "Unknown"}
    umarshalled, _ := json.Marshal(u)
    w.Write(umarshalled)
}


type authMiddleWare struct {

}

func (amw *authMiddleWare) ValidateJWT(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" &&  r.URL.Path == "/login" {
            LoginUser(w, r)
        } else {
            tokenToken := strings.Split(token, " ")
            if len(tokenToken) == 2 {
                parsed, error := jwt.Parse(tokenToken[1], func(token *jwt.Token) (interface{}, error) {
                    return []byte("mysecretkey"), nil
                })

                if error != nil {
                    log.Println(error)
                    w.WriteHeader(http.StatusUnauthorized)

                } else {
                    c := parsed.Claims.(jwt.MapClaims)
                    context.Set(r, "email", c["email"])
                    next.ServeHTTP(w, r)
                }
            }
        }
    })
}