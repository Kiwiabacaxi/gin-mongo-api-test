package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// Omitempty = não retornar o campo se estiver vazio
	// validate = validação de dados
	// required = campo obrigatório

	Id       primitive.ObjectID `json:"id,omitempty"`                           // ID do usuário
	Name     string             `json:"name,omitempty" validate:"required"`     // Nome do usuário
	Location string             `json:"location,omitempty" validate:"required"` // Localização do usuário
	Title    string             `json:"title,omitempty" validate:"required"`    // Título do usuário
}
