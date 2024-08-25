package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jetnoli/go-admin/db/models/user"
	"github.com/jetnoli/go-admin/services"
	// "github.com/jetnoli/go-admin/view/components/module"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body := &user.Properties{}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Error Parsing Json Body: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(body)

	if err != nil {
		http.Error(w, "Error Creating User: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&user)

	if err != nil {
		http.Error(w, "Error Returing User Details: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	users, err := services.GetAllUsers()

	if err != nil {
		http.Error(w, "Error Fetching Users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resType, ok := r.Header["Response-Type"]

	if !ok {
		resType = []string{"htmx"}
	}

	if resType[0] == "json" {
		w.Header().Add("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(&users)

		if err != nil {
			http.Error(w, "Error Returning Users: "+err.Error(), http.StatusInternalServerError)
		}

		return
	}

	// // Assume resType == "htmx"
	component := UserItemList(users)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error Rendering Users List: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid Id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(id)

	if err != nil {
		http.Error(w, "User Not Found: "+err.Error(), http.StatusNotFound)
		return
	}

	resType, ok := r.Header["Response-Type"]

	if !ok {
		resType = []string{"htmx"}
	}

	if resType[0] == "json" {
		w.Header().Add("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(&user)

		if err != nil {
			http.Error(w, "Error Returning User: "+err.Error(), http.StatusInternalServerError)
		}

		return
	}

	query := r.URL.Query()
	asList := query.Get("as_list")

	// Assume resType == "htmx"
	component := UserDetail(*user)

	if asList == "true" {
		component = UserListItem(*user)
	}

	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error Rendering User: "+err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid Id: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	body := &user.Properties{}

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Error Parsing Json Body: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.UpdateUserById(id, body)

	if err != nil {
		http.Error(w, "User Not Found: "+err.Error(), http.StatusNotFound)
		return
	}

	resType, ok := r.Header["Response-Type"]

	if !ok {
		resType = []string{"htmx"}
	}

	if resType[0] == "json" {
		w.Header().Add("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(&user)

		if err != nil {
			http.Error(w, "Error Returning User: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("HX-Trigger", "postFormCompleted")

	component := UserDetail(*user)
	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error Rendering User: "+err.Error(), http.StatusInternalServerError)
	}

}

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	err := services.DeleteAllUsers()

	if err != nil {
		http.Error(w, "Error Deleting All Users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(`{
		"success": true
	}`))
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid Id: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteUserById(id)

	if err != nil {
		http.Error(w, "Error Deleting All Users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(`{
		"success": true
	}`))
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	user, err := services.GetUserByUsername(username)

	if err != nil {
		http.Error(w, "User Not Found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&user)

	if err != nil {
		http.Error(w, "Error Returning User: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
