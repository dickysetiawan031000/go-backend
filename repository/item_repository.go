package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/dickysetiawan031000/go-backend/model"
)

type ItemRepository interface {
	Create(item model.Item) (model.Item, error)
	FindAll() ([]model.Item, error)
	FindByID(id uint) (model.Item, error)
	Update(id uint, updated model.Item) (model.Item, error)
	Delete(id uint) error
}

type itemRepo struct {
	data   map[uint]model.Item
	mu     sync.RWMutex
	autoID uint
}

func NewItemRepository() ItemRepository {
	return &itemRepo{
		data: make(map[uint]model.Item),
	}
}

func (r *itemRepo) Create(item model.Item) (model.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.autoID++
	item.ID = r.autoID
	item.CreatedAt = now()
	item.UpdatedAt = now()

	r.data[item.ID] = item
	return item, nil
}

func (r *itemRepo) FindAll() ([]model.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var items []model.Item
	for _, v := range r.data {
		items = append(items, v)
	}
	return items, nil
}

func (r *itemRepo) FindByID(id uint) (model.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.data[id]
	if !exists {
		return model.Item{}, errors.New("item not found")
	}
	return item, nil
}

func (r *itemRepo) Update(id uint, updated model.Item) (model.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.data[id]
	if !exists {
		return model.Item{}, errors.New("item not found")
	}

	updated.ID = id
	updated.CreatedAt = item.CreatedAt
	updated.UpdatedAt = now()
	r.data[id] = updated

	return updated, nil
}

func (r *itemRepo) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("item not found")
	}
	delete(r.data, id)
	return nil
}

// helper
func now() int64 {
	return int64(uint64(time.Now().Unix()))
}
