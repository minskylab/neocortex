package neocortex

import "errors"

var ErrSessionNotExist = errors.New("session not exist")

// var ErrSessionExpired = errors.Default("session expired")
var ErrInvalidResponseFromCognitiveService = errors.New("invalid response from cognitive service")
var ErrInvalidInputType = errors.New("invalid or unimplemented input type")
var ErrContextNotExist = errors.New("context is not valid or not exist")
var ErrChannelIsNotRegistered = errors.New("channel not exist on this engine instance")
