package api

import "reflection_prototype/internal/storage"

type Handler struct {
	S storage.Storer
}
