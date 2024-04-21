package db

import "fmt"

// TODO: Tratar erros para permitir incluir msgs de onde ocorreram
var ErrInvalidType = fmt.Errorf("invalid type")
var ErrInvalidStructType = fmt.Errorf("invalid type, expected struct")
var ErrNotAddressable = fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
var ErrFieldNotFound = fmt.Errorf("cannot assign to the item passed, field not found")
