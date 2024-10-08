package user

import (
    "gorm.io/gorm"
)

type Repository interface {
    CreateUser(user *User) error
    GetUserByID(id uint) (*User, error)
    GetUserByEmail(email string) (*User, error)
    GetAllUsers() ([]User, error)
    UpdateUser(user *User) error
    DeleteUser(id uint) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db}
}

func (r *repository) CreateUser(user *User) error {
    return r.db.Create(user).Error
}

func (r *repository) GetUserByID(id uint) (*User, error) {
    var user User
    if err := r.db.Preload("PriceOffers").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
    var user User
    if err := r.db.Preload("PriceOffers").Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *repository) GetAllUsers() ([]User, error) {
    var users []User
    if err := r.db.Preload("PriceOffers").Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *repository) UpdateUser(user *User) error {
    return r.db.Save(user).Error
}

func (r *repository) DeleteUser(id uint) error {
    return r.db.Delete(&User{}, id).Error
}
