package db

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Tratar erros para permitir incluir msgs de onde ocorreram
var ErrInvalidType = errors.New("invalid type")
var ErrInvalidStructType = errors.New("invalid type, expected struct")
var ErrNotAddressable = errors.New("cannot assign to the item passed, item must be a pointer in order to assign")
var ErrFieldNotFound = errors.New("cannot assign to the item passed, field not found")

var ErrInvalidHex = primitive.ErrInvalidHex
