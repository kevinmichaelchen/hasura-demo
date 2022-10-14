// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __getAnimalInput is used internally by genqlient
type __getAnimalInput struct {
	Species string `json:"species"`
}

// GetSpecies returns __getAnimalInput.Species, and is useful for accessing the field via an interface.
func (v *__getAnimalInput) GetSpecies() string { return v.Species }

// getAnimalAnimal includes the requested fields of the GraphQL type animal.
// The GraphQL type's documentation follows.
//
// animal
type getAnimalAnimal struct {
	Id      string `json:"id"`
	Species string `json:"species"`
}

// GetId returns getAnimalAnimal.Id, and is useful for accessing the field via an interface.
func (v *getAnimalAnimal) GetId() string { return v.Id }

// GetSpecies returns getAnimalAnimal.Species, and is useful for accessing the field via an interface.
func (v *getAnimalAnimal) GetSpecies() string { return v.Species }

// getAnimalResponse is returned by getAnimal on success.
type getAnimalResponse struct {
	// fetch data from the table: "animal"
	Animal []getAnimalAnimal `json:"animal"`
}

// GetAnimal returns getAnimalResponse.Animal, and is useful for accessing the field via an interface.
func (v *getAnimalResponse) GetAnimal() []getAnimalAnimal { return v.Animal }

func getAnimal(
	ctx context.Context,
	client graphql.Client,
	species string,
) (*getAnimalResponse, error) {
	req := &graphql.Request{
		OpName: "getAnimal",
		Query: `
query getAnimal ($species: String!) {
	animal(where: {species:{_eq:$species}}) {
		id
		species
	}
}
`,
		Variables: &__getAnimalInput{
			Species: species,
		},
	}
	var err error

	var data getAnimalResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
