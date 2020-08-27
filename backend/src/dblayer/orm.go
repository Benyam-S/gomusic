package dblayer

import (
	"errors"

	"github.com/Benyam-S/gomusic/backend/src/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// DBORM is a type that defines the database handler of the system
type DBORM struct {
	conn *gorm.DB
}

// NewORM is a function that returns a new orm type
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		conn: db,
	}, err
}

// GetProduct is a method that returns a product that matchs the given id
func (db *DBORM) GetProduct(id uint) (product models.Product, err error) {
	return product, db.conn.First(&product, id).Error
}

// GetAllProducts is a method that returns all the products in the repository
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {

	products = make([]models.Product, 0)
	err = db.conn.Find(&products).Error
	return products, err
}

// GetPromos is a method that returns all the promo products in the repository
func (db *DBORM) GetPromos() (products []models.Product, err error) {

	products = make([]models.Product, 0)
	err = db.conn.Where("promotion IS NOT NULL").Find(&products).Error
	return products, err
}

// GetCustomerByName is a method that returns a customer using firstname and lastname
func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {

	// Where clause can also take structs rather than string
	err = db.conn.Where(&models.Customer{FirstName: firstname, LastName: lastname}).First(&customer).Error

	return customer, err
}

// GetCustomerByID is a method that returns a customer that matchs the given id
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.conn.First(&customer, id).Error
}

// AddUser is a method that adds a new user to the system
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Password)

	customer.LoggedIn = true
	err := db.conn.Create(&customer).Error
	customer.Password = ""

	return customer, err
}

// SignInUser is a method that sign in users to the system
func (db *DBORM) SignInUser(email, password string) (customer models.Customer, err error) {

	result := db.conn.Model(models.Customer{}).Where("email = ?", email)

	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if checkPassword(customer.Password, password) {
		return customer, errors.New("invalid password")
	}

	customer.LoggedIn = true
	err = result.Update(&customer).Error
	if err != nil {
		return customer, err
	}

	customer.Password = ""

	return customer, nil

}

// SignOutUserByID is a method that sign's out user from the system
func (db *DBORM) SignOutUserByID(id int) error {

	customer := new(models.Customer)
	customer.ID = uint(id)

	return db.conn.Model(customer).Update("logged_in", false).Error
}

// GetCustomerOrdersByID is a method that get's all the customers orders
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {

	return orders, db.conn.Model(&models.Order{}).Where("customer_id", id).Find(&orders).Error
}

// AddOrder is a method that add the order to the orders table
func (db *DBORM) AddOrder(order models.Order) error {
	return db.conn.Create(&order).Error
}

// GetCreditCardCID is a method that gets the id representing the credit card from the database
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	cusomterWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return cusomterWithCCID.CCID, db.conn.Model(&models.Customer{}).First(&cusomterWithCCID, id).Error
}

// SaveCreditCardForCustomer is a method that saves the credit card information for the customer
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.conn.Model(&models.Customer{}).Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}

func hashPassword(s *string) error {

	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(*s), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashedByte)

	return nil

}

func checkPassword(existingHash, incomingPass string) bool {
	//this method will return an error if the hash does not match the provided password string
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}
