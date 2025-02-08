package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"time"
	//"github.com/google/uuid"
)

func(cfg *config) handleLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	password := r.FormValue("password")

	user, err := cfg.db.GetUserByUserName(r.Context(), userName)
	if err != nil {
		//respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(password, user.HashedPassword)
	if err != nil{
		//respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil{
		//respondWithError(w, http.StatusInternalServerError, "Could not create token", err)
		return
	}

	refreshTokenString, err := auth.MakeRefreshToken()
	if err != nil{
		//respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
		return
	}

	refreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{refreshTokenString, user.ID, time.Now().UTC().Add(60 * 24 * time.Hour)})
	if err != nil{
		//respondWithError(w, http.StatusInternalServerError, "Could not create refresh token", err)
		return
	}

}