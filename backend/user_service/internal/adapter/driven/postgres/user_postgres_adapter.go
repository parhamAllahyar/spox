package postgres

import (
	"time"
	"user/internal/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName  string    `gorm:"default:NULL"`
	LastName   string    `gorm:"default:NULL"`
	Email      string    `gorm:"default:NULL"`
	Phone      string    `gorm:"default:NULL"`
	Password   string
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Deleted_at gorm.DeletedAt
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userRepo {
	return userRepo{
		db: db,
	}
}

func (r userRepo) Create(user entity.User) error {

	newUser := User{
		ID:        uuid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
		Email:     user.Email,
	}

	// Create the user in the database
	query := r.db.Create(&newUser)
	if query.Error != nil {
		return query.Error
	}
	return nil

}

func (r userRepo) Update(data entity.User) error {

	var user User
	query := r.db.First(&user, "id = ?", data.ID)
	if query.Error != nil {
		return query.Error
	}
	user = User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	query = r.db.Model(&user).Where("id = ?", data.ID).Updates(&user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r userRepo) Delete(id uuid.UUID) error {

	var user User
	query := r.db.First(&user, id)

	if query.Error != nil {
		return query.Error
	}

	query = r.db.Delete(&user)

	if query.Error != nil {
		return query.Error
	}

	return nil

}

func (r userRepo) GetByID(id uuid.UUID) (entity.User, error) {

	var user User
	query := r.db.Where("id = ?", id).First(&user)

	if query.RowsAffected == 0 {
		return entity.User{}, nil
	}

	if query.Error != nil {
		return entity.User{}, query.Error
	}

	return entity.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil

}

func (r userRepo) GetByEmail(email string) (entity.User, error) {

	var user User

	query := r.db.Where("email = ?", email).First(&user)

	if query.RowsAffected == 0 {
		return entity.User{}, nil
	}

	if query.Error != nil {
		return entity.User{}, query.Error
	}

	return entity.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}, nil

}

func (r userRepo) Index(page int, pageSize int) ([]entity.User, error) {

	var usersList []entity.User

	var users []User

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, u := range users {
		usersList = append(usersList, entity.User{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		})
	}

	return usersList, nil
}

func (r userRepo) UpdatePhone(id uuid.UUID, phone string) error {
	var user User
	query := r.db.First(&user, "id = ?", id)
	// TODO:
	if query.Error != nil {
		return query.Error
	}
	r.db.Model(&user).Update("phone", phone)
	return nil
}

func (r userRepo) UpdateEmail(id uuid.UUID, email string) error {
	var user User
	query := r.db.First(&user, "id = ?", id)
	// TODO:
	if query.Error != nil {
		return query.Error
	}
	r.db.Model(&user).Update("email", email)
	return nil
}

func (r userRepo) UpdatePassword(id uuid.UUID, newPassword string) error {

	var user User
	query := r.db.First(&user, "id = ?", id)
	// TODO:
	if query.Error != nil {
		return query.Error
	}
	r.db.Model(&user).Update("password", newPassword)
	return nil

}

func (r userRepo) Count() (int64, error) {
	var count int64
	query := r.db.Table("users").Model(&User{}).Count(&count)
	return count, query.Error
}

func (r userRepo) Search(string) ([]entity.User, error) {
	return nil, nil
}

func (r userRepo) SearchByEmail(string) (entity.User, error) {

	return entity.User{}, nil

}
