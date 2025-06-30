package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
)

func (cfg *config)handleGetMenu(w http.ResponseWriter, r *http.Request){
	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		//fmt.Println("No Bearer token")
		return
	}

	_, err = auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		//respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		//fmt.Println(err)
		return
	}

	classificationMap := make(map[string][]database.Drug)
	drugData, err := cfg.db.ListDrugsByClassification(r.Context())
	if err != nil{
		//errs.General = "Server Error"
		//templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	for _, data := range drugData{
		_, ok := classificationMap[data.Name]
		if !ok{
			classificationMap[data.Name] = []database.Drug{}
		}
		classificationMap[data.Name] = append(classificationMap[data.Name], data.Drug)
	}

	templates.Game(classificationMap).Render(r.Context(), w)
}