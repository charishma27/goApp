package store

import (
    "errors"
    "goApp/models"
    "sync"
)

type Store struct {
    users map[int]models.User
    mu    sync.Mutex
    nextID int
}

var UserStore = Store{
    users:  make(map[int]models.User),
    nextID: 1000,
}

func (s *Store) Create(user models.User) models.User {
    s.mu.Lock()
    defer s.mu.Unlock()
    user.ID = s.nextID
    s.nextID++
    s.users[user.ID] = user
    return user
}

func (s *Store) GetAll() []models.User {
    s.mu.Lock()
    defer s.mu.Unlock()
    var users []models.User
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
}

func (s *Store) GetByID(id int) (models.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    user, exists := s.users[id]
    if !exists {
        return models.User{}, errors.New("user not found")
    }
    return user, nil
}

func (s *Store) Update(id int, newUser models.User) (models.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    _, exists := s.users[id]
    if !exists {
        return models.User{}, errors.New("user not found")
    }
    newUser.ID = id
    s.users[id] = newUser
    return newUser, nil
}

func (s *Store) Delete(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    _, exists := s.users[id]
    if !exists {
        return errors.New("user not found")
    }
    delete(s.users, id)
    return nil
}
