package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		// Cadastra um usuário
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		// retorna dados de todos os usuarios
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BurcarUsuarios,
		RequerAutenticacao: false,
	},
	{
		// retorna o usuario por id
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BurcarUsuario,
		RequerAutenticacao: false,
	},
	{
		// Atualiza dados de um usuario
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUsuario,
		RequerAutenticacao: false,
	},
	{
		// deleta um usuario
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletaUsuario,
		RequerAutenticacao: false,
	},
}
