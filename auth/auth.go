package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"

	"github.com/scottmckendry/mnemstart/config"
)

type AuthService struct{}

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		github.New(
			config.Envs.GithubClientID,
			config.Envs.GithubClientSecret,
			buildCallbackURL("github"), "user:email",
		),
		discord.New(
			config.Envs.DiscordClientID,
			config.Envs.DiscordClientSecret,
			buildCallbackURL("discord"), "identify", "email",
		),
	)

	return &AuthService{}
}

func (a *AuthService) GetSessionUser(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, err
	}

	user := session.Values["user"]
	if user == nil {
		return goth.User{}, fmt.Errorf("user is unauthenticated! %v", user)
	}

	return user.(goth.User), nil
}

func (a *AuthService) StoreUserSession(
	w http.ResponseWriter,
	r *http.Request,
	user goth.User,
) error {
	session, _ := gothic.Store.Get(r, SessionName)
	session.Values["user"] = user

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func (a *AuthService) RemoveUserSession(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		log.Println("User session not found")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	session.Values["user"] = goth.User{}
	session.Options.MaxAge = -1

	session.Save(r, w)
}

func RequireAuth(auth *AuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := auth.GetSessionUser(r)
			if err != nil {
				log.Println("User is not authenticated!")
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
				return
			}

			log.Printf("Authenticated user: %v", session.Email)
			next.ServeHTTP(w, r)
		})
	}
}

func buildCallbackURL(provider string) string {
	hostUrl := fmt.Sprintf("%s", config.Envs.PublicHost)
	if config.Envs.SendPortInCallback {
		hostUrl = fmt.Sprintf("%s:%s", hostUrl, config.Envs.Port)
	}
	return fmt.Sprintf(
		"%s/auth/%s/callback",
		hostUrl,
		provider,
	)
}
