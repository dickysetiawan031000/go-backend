package usecase

import (
	"github.com/dickysetiawan031000/go-backend/dto/item"
	"github.com/dickysetiawan031000/go-backend/mapper"
	"github.com/dickysetiawan031000/go-backend/model"
	"github.com/dickysetiawan031000/go-backend/repository"
)

type ItemUseCase interface {
	Create(input item.CreateItemRequest) (model.Item, error)
	GetAll() ([]model.Item, error)
	GetByID(id uint) (model.Item, error)
	Update(id uint, input item.UpdateItemRequest) (model.Item, error)
	Delete(id uint) error
}

type itemUseCase struct {
	repo repository.ItemRepository
}

func NewItemUseCase(repo repository.ItemRepository) ItemUseCase {
	return &itemUseCase{repo: repo}
}

func (uc *itemUseCase) Create(input item.CreateItemRequest) (model.Item, error) {
	itemModel := mapper.ToItemModel(input)
	return uc.repo.Create(itemModel)
}

func (uc *itemUseCase) GetAll() ([]model.Item, error) {
	return uc.repo.FindAll()
}

func (uc *itemUseCase) GetByID(id uint) (model.Item, error) {
	return uc.repo.FindByID(id)
}

func (uc *itemUseCase) Update(id uint, input item.UpdateItemRequest) (model.Item, error) {
	itemModel := mapper.ToItemModel(input)
	return uc.repo.Update(id, itemModel)
}

func (uc *itemUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
