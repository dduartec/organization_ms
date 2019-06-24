package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"app/database"
	//"../database"
)

type move_struct struct {
	Owner   string `json:"owner"`
	Origin  string `json:"origin"`
	Destiny string `json:"destiny"`
}

type create_struct struct {
	Owner string `json:"owner"`
	Path  string `json:"path"`
}

type delete_struct struct {
	Owner string `json:"owner"`
	Path  string `json:"path"`
}
type error_struct struct {
	Advise string `json:"advise"`
}

//Move test
func Move(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Move")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	var t move_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	temp, err := json.Marshal(t)
	if err != nil {
		log.Panic(err)
	}
	aux, err := url.Parse("http://file-controller-ms:2870/move?source=" + t.Origin + "&target=" + t.Destiny)
	if err != nil {
		log.Panic(err)
	}
	res, err := http.Post(aux.String(), "application/json", bytes.NewBuffer(temp))
	if err != nil {
		log.Panic(err)
	}
	json.NewDecoder(res.Body).Decode(&t)
	fmt.Println(t)
	err = os.Rename(t.Origin, t.Destiny)
	if err != nil {
		var jsonerr error_struct
		jsonerr.Advise = "No existe el archivo/carpeta en la ruta de origen"
		log.Println("No existe el archivo/carpeta en la ruta de origen")
		json.NewEncoder(w).Encode(jsonerr)
		return
	}
	db := database.Connect()
	jsonres := database.Move(db, t.Owner, t.Origin, t.Destiny)
	json.NewEncoder(w).Encode(jsonres)
	fmt.Println(database.Disconnect(db))
}

func LogMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	db := database.Connect()
	aux := database.LogMove(db)
	json.NewEncoder(w).Encode(aux)
	fmt.Println(database.Disconnect(db))
}

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	decoder := json.NewDecoder(r.Body)
	var t create_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	aux, err := exists(t.Path)
	if aux == false {
		temp, err := json.Marshal(t)
		if err != nil {
			log.Panic(err)
		}
		aux, err := url.Parse("http://file-controller-ms:2870/createFolder?path=" + t.Path)
		if err != nil {
			log.Panic(err)
		}
		res, err := http.Post(aux.String(), "application/json", bytes.NewBuffer(temp))
		if err != nil {
			log.Panic(err)
		}
		json.NewDecoder(res.Body).Decode(&t)
		fmt.Println(t.Path)
		err = os.Mkdir(t.Path, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
		db := database.Connect()
		jsonres := database.Create(db, t.Owner, t.Path)
		json.NewEncoder(w).Encode(jsonres)
		fmt.Println(database.Disconnect(db))
	} else {
		var jsonerr error_struct
		jsonerr.Advise = "La carpeta ya existe"
		log.Println("La carpeta ya existe")
		json.NewEncoder(w).Encode(jsonerr)
		return

	}

}

func LogCreateFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	db := database.Connect()
	aux := database.LogCreate(db)
	json.NewEncoder(w).Encode(aux)
	fmt.Println(database.Disconnect(db))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	var t delete_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	aux, err := exists(t.Path)
	if aux == true {
		temp, err := json.Marshal(t)
		if err != nil {
			log.Panic(err)
		}
		aux, err := url.Parse("http://file-controller-ms:2870/delete?path=" + t.Path)
		if err != nil {
			log.Panic(err)
		}
		res, err := http.Post(aux.String(), "application/json", bytes.NewBuffer(temp))
		if err != nil {
			log.Panic(err)
		}
		json.NewDecoder(res.Body).Decode(&t)
		fmt.Println(t)
		err = os.RemoveAll(t.Path)
		if err != nil {
			panic(err)
		}
		db := database.Connect()
		jsonres := database.Delete(db, t.Owner, t.Path)
		json.NewEncoder(w).Encode(jsonres)
		fmt.Println(database.Disconnect(db))
	} else {
		var jsonerr error_struct
		jsonerr.Advise = "La carpeta/archivo no existe"
		log.Println("La carpeta/archivo no existe")
		json.NewEncoder(w).Encode(jsonerr)
		return
	}

}

func LogDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	db := database.Connect()
	aux := database.LogDelete(db)
	json.NewEncoder(w).Encode(aux)
	fmt.Println(database.Disconnect(db))
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
