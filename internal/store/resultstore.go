package store

import (
	"sync"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)


type Store struct{
	data map[string] *domain.Result
	mu sync.RWMutex
}

type Storer interface {
	Set(key string, value *domain.Result)
	Get(key string) (*domain.Result, bool)
	Delete(key string)
}

func NewStore()*Store{

	return &Store{
		data: make(map[string]*domain.Result),
	}

}

func(s *Store)Set(key string,value *domain.Result){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func(s *Store)Get(key string) (*domain.Result, bool){
	s.mu.RLock()
	defer s.mu.RUnlock()

	val,ok:= s.data[key]
	
	return val,ok
}


func (s *Store)Delete(key string){
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data,key)
}