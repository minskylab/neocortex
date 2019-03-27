package neocortex

import "errors"

var ErrSessionNotExist = errors.New("session not exist")
var ErrSessionExpired = errors.New("session expired")
var ErrInvalidResponseFromCognitiveService = errors.New("invalid response from cognitive service")
var ErrInvalidInputType = errors.New("invalid or unimplemented input type")
