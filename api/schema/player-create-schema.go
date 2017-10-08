package schema

const PlayerCreateSchema = `
{
	"$schema": "http://json-schema.org/draft-04/schema#", 
	"additionalProperties": false, 
	"definitions": {}, 
	"id": "http://example.com/example.json", 
	"properties": {
	  "mail": {
		"description": "Player's mail", 
		"examples": [
		  "test@test.fr"
		], 
		"id": "/properties/mail", 
		"title": "The mail schema.", 
		"type": "string"
	  }, 
	  "name": {
		"description": "Player's name", 
		"examples": [
		  "TEST"
		], 
		"id": "/properties/name", 
		"title": "The name schema.", 
		"type": "string"
	  }, 
	  "nationality": {
		"description": "Player's nationality", 
		"examples": [
		  "FRANCE"
		], 
		"id": "/properties/nationality", 
		"title": "The nationality schema.", 
		"type": "string"
	  }, 
	  "password": {
		"description": "Player's password", 
		"examples": [
		  "TEST"
		], 
		"id": "/properties/password", 
		"title": "The password schema.", 
		"type": "string"
	  }
	}, 
	"required": [
	  "mail", 
	  "password", 
	  "name", 
	  "nationality"
	], 
	"type": "object"
  }
`
