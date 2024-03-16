package go_cache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       int
	Username string
}
type Server struct {
	db    map[int]*User
	cache map[int]*User
	dbHit int
}

func NewServer() *Server {
	db := make(map[int]*User, 0)

	for i := 1; i <= 100; i++ {
		db[i] = &User{
			ID:       i,
			Username: fmt.Sprintf("user_%d", i),
		}
	}
	return &Server{
		db:    db,
		cache: make(map[int]*User, 0),
	}
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	// Lookup in Cache first
	user, ok := s.cache[id]
	if ok {
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	// If not found, request the database
	if user, ok = s.db[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
	}
	s.dbHit++

	s.cache[id] = user
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
