package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/pusher/pusher-http-go"
	"net/http"
	. "safer.com/m/models"
	"safer.com/m/services"
)

func SendMessage(pusherClient pusher.Client) http.HandlerFunc{
	return func(w http.ResponseWriter,r* http.Request){
		claims:=r.Context().Value("user").(jwt.MapClaims)

		if services.CheckMessagePermissions(claims["id"].(string),UserType(claims["role"].(string)),chi.URLParam(r,"caseId")){
			var message Message
			json.NewDecoder(r.Body).Decode(&message)
			message.SenderId=claims["id"].(string)
			message.CaseId=chi.URLParam(r,"caseId")
			if message,err:=services.CreateMessage(message);err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
			}else{
				channels,err:=services.GetRecipientsChannelId(message.CaseId)
				if err==nil{
					fmt.Println(channels)
					for _,v:=range channels{
						pusherClient.Trigger(v, "new-message", message)
					}
				}
				json.NewEncoder(w).Encode(message)
			}

		}else{
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(HttpError{Message: "Unauthorized"})
		}
	}
}

func GetMessageToken(w http.ResponseWriter, r* http.Request){
	token:=services.GetChannelToken(r.Context().Value("user").(jwt.MapClaims)["id"].(string))
	json.NewEncoder(w).Encode(map[string]string{"token":token})
}

func GetMessages(w http.ResponseWriter, r* http.Request){
	messages:=services.GetMessagesByCaseId(chi.URLParam(r,"caseId"))
	json.NewEncoder(w).Encode(messages)
}