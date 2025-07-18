package main

import(
	"net/http"
	"rxcheck/internal/auth"
	//"rxcheck/internal/database"
	"time"
	//"fmt"
	//"github.com/google/uuid"
	"encoding/json"
	"rxcheck/templates"
)

func(cfg *config) handleLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	password := r.FormValue("password")

	// params := templates.LoginParams{
	// 	Title: "Login",
	// 	SwapMessage: "Don't Have an account? Create Account",
	// 	SubmitURL: "/login_user",
	// 	SwapURL: "/swap_create",
	// }

	errs := templates.LoginError{
	}

	if userName == ""{
		errs.Name = "Enter a Username"
		templates.Login(templates.LOGIN_PARAMS, errs).Render(r.Context(), w)
		return
	}

	if password == ""{
		errs.Password = "Enter a password"
		templates.Login(templates.LOGIN_PARAMS, errs).Render(r.Context(), w)
		return
	}

	user, err := cfg.db.GetUserByUserName(r.Context(), userName)
	if err != nil {
		errs.General = "Incorrect email or password"
		templates.Login(templates.LOGIN_PARAMS, errs).Render(r.Context(), w)
		//respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(password, user.HashedPassword)
	if err != nil{
		errs.General = "Incorrect email or password"
		templates.Login(templates.LOGIN_PARAMS, errs).Render(r.Context(), w)
		//respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	//token, err := auth.MakeJWT(uuid.New(), cfg.jwtSecret, time.Hour)
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, 24 * time.Hour)
	if err != nil{
		//respondWithError(w, http.StatusInternalServerError, "Could not create token", err)
		return
	}

	// refreshTokenString, err := auth.MakeRefreshToken()
	// if err != nil{
	// 	//respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
	// 	return
	// }

	// refreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{refreshTokenString, user.ID, time.Now().UTC().Add(60 * 24 * time.Hour)})
	// if err != nil{
	// 	//respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
	// 	return
	// }

	type TokenData struct {
		Login struct {
			Token string `json:"token"`
		} `json:"login"`
	}

	data := TokenData{Login: struct{Token string `json:"token"`}{Token: token}}

	json, err := json.Marshal(data)
	if err != nil {
		//fmt.Printf("Error marshalling JSON: %s", err)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//w.Header().Add("HX-Trigger", fmt.Sprintf("{ \"login\": { \"token\": \"%s\"}}", token))
	//fmt.Println(string(json))
	w.Header().Add("HX-Trigger", string(json))

	//classificationMap := make(map[string][]database.Drug)
	drugData, err := cfg.db.ListDrugsByClassification(r.Context())
	if err != nil{
		errs.General = "Server Error"
		//fmt.Println(string(drugData[0].Drugs))
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	// for _, data := range drugData{
	// 	_, ok := classificationMap[data.Name]
	// 	if !ok{
	// 		classificationMap[data.Name] = []database.Drug{}
	// 	}
	// 	classificationMap[data.Name] = append(classificationMap[data.Name], data.Drug)
	// }

	// classifications, err := cfg.db.ListClassifications(r.Context())
	// if err != nil{
	// 	errs.General = "Server Error"
	// 	templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
	// 	return
	// }

	templates.Game(drugData).Render(r.Context(), w)
}