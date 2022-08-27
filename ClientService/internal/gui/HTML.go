package gui_html

import (
	gRPC_Client "ClientService/external/gRPC-Client"
	"ClientService/internal/logging"
	"ClientService/internal/user"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	URL_MAIN                      = "/"
	URL_GET_STORAGE               = "/GetStorage"
	URL_ADD_UPDATE_STORAGE_RECORD = "/AddUpdateStorageRecord"
	URL_DELETE_RECORD             = "/DeleteRecord"
	URL_AUTH                      = "/Auth"
)

type Handler struct {
	HtmlRoot string
	User     *user.User
	*logging.Logger
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL_MAIN, h.MainPage)
	router.HandlerFunc(http.MethodPost, URL_GET_STORAGE, h.GetStorage)
	router.HandlerFunc(http.MethodPost, URL_ADD_UPDATE_STORAGE_RECORD, h.AddUpdateStorageRecord)
	router.HandlerFunc(http.MethodPost, URL_DELETE_RECORD, h.DeleteRecord)
	router.HandlerFunc(http.MethodPost, URL_AUTH, h.Auth)
}

// Auth  TODO: change return header codes
// @Summary Update or create storage on Encryption Server, through gRPC
// @Tags gui
// @Success 200
// @Failure 401
// @Router /Auth [post]
func (h *Handler) Auth(w http.ResponseWriter, _ *http.Request) {
	//TODO: Username/Password from configuration
	h.Logger.Info().Msg("Started Auth on auth server...")
	h.User.Auth.Authentificate(h.User.GetUserName(), h.User.GetPassword())
	h.Logger.Info().Msgf("Token for name:%s password:%s \t Token:%X", h.User.GetUserName(), h.User.GetPassword(), h.User.GetToken())
	if _, err := w.Write([]byte("Одноразовый токен успешно получен")); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	h.Logger.Info().Msg("Auth completed.")

}

// MainPage TODO: Set gui template to config file?
// @Summary Main html page
// @Tags gui
// @Success 200
// @Failure 500
// @Router / [get]
func (h *Handler) MainPage(w http.ResponseWriter, _ *http.Request) {
	ts, err := template.ParseFiles(h.HtmlRoot)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetStorage TODO: change return header codes
// @Summary Getting storage from Encryption Server, through gRPC
// @Tags gui
// @Success 200
// @Failure 401
// @Router /GetStorage [post]
func (h *Handler) GetStorage(w http.ResponseWriter, _ *http.Request) {
	//TODO: изменить хардкод порта/адреса и изменить имя переменной
	gRPC_Client.StartConnectionAndGetStorage("127.0.0.1:9999", h.User.GetToken(), 10, h.Logger)
	if _, err := w.Write([]byte("Hello from GetStorage")); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// AddUpdateStorageRecord  TODO: change return header codes
// @Summary Update or create storage on Encryption Server, through gRPC
// @Tags gui
// @Success 200
// @Failure 401
// @Router /AddUpdateStorageRecord [post]
func (h *Handler) AddUpdateStorageRecord(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Hello from AddUpdateStorage")); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteRecord   TODO: change return header codes
// @Summary DeleteRecord by storage from Encryption Server, through gRPC
// @Tags gui
// @Success 200
// @Failure 401
// @Router /DeleteRecord [post]
func (h *Handler) DeleteRecord(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Hello from DeleteRecord")); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
