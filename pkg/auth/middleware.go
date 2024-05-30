package auth

import (
	"context"
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/qiniu/qmgo"

	"supplychain/graph/model"
	"supplychain/pkg/jwt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCtxKey = &contextKey{"AIzaSyDMi3RRz0nP0Seh141cUAL9LpaxI1R_3Uo"}

type contextKey struct {
	name string
}

type Response struct {
	ResponseCode int32  `json:"responseCode"`
	ResponseMsg  string `json:"responseMsg"`
}

func Middleware(db *qmgo.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return

				// resp := `{responseCode: 401, responseMsg: "Token is Empty"}`
				// var response Response
				// if err := json.Unmarshal([]byte(resp), &response); err != nil {
				// 	fmt.Println("unmarshal error", err)
				// }
				// http.Error(w, response.ResponseMsg, http.StatusUnauthorized)
				// return
			}

			//validate jwt token
			sessionID, err := jwt.ParseToken(tokenStr)
			if err != nil {
				var resp string
				if err.Error() == "Token is expired" {
					resp = `{"responseCode": 403, "responseMsg": "Token is expired"}`
				} else if err.Error() == "signature is invalid" {
					resp = `{"responseCode": 403, "responseMsg": "Invalid Token"}`
				}
				// var response Response
				// if err := json.Unmarshal([]byte(resp), &response); err != nil {
				// 	return
				// }
				// fmt.Println("response return", sessionID)
				http.Error(w, resp, http.StatusUnauthorized)
				return
			}

			// create user and check if user exists in db

			var m model.LoginData
			sID, _ := primitive.ObjectIDFromHex(sessionID)
			if err := db.Collection("users").Find(context.Background(), bson.M{"_id": sID}).One(&m); err != nil {
				next.ServeHTTP(w, r)
				return
			} else {
				fmt.Println("User verified in session", m.ID)
				ctx := context.WithValue(r.Context(), userCtxKey, &m.ID)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			}
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *string {
	raw, _ := ctx.Value(userCtxKey).(*string)
	return raw
}
