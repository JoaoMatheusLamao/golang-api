package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

func BurcarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuários"))
}

func BurcarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário"))
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza um Usuário"))
}

func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apagando um Usuário"))
}
