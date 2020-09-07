package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Benyam-S/gomusic/backend/src/dblayer"
	"github.com/Benyam-S/gomusic/backend/src/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"

	"github.com/gin-gonic/gin"
)

// HandlerInterface is an interface that defines all the interface required by our fronted
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

// Handler is a struct that defines a handler type
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler is a function that returns a new handler type
func NewHandler(dbname, conn string) (*Handler, error) {
	db, err := dblayer.NewORM(dbname, conn)
	if err != nil {
		return nil, err
	}

	return &Handler{db: db}, nil
}

// GetProducts is a method that handles the request for getting all the products
func (h *Handler) GetProducts(c *gin.Context) {

	if h.db == nil {
		return
	}

	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetPromos is a method that handles the request for getting all the promos
func (h *Handler) GetPromos(c *gin.Context) {

	if h.db == nil {
		return
	}

	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promos)
}

// SignIn is a method that handles a singin request
func (h *Handler) SignIn(c *gin.Context) {

	if h.db == nil {
		return
	}

	customer := models.Customer{}
	c.ShouldBindJSON(&customer)

	customer, err := h.db.SignInUser(customer.Email, customer.Password)
	if err != nil {

		if err.Error() == "invalid password" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// AddUser is a method that handles add user request
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// SignOut is a method that handles sign out request
func (h *Handler) SignOut(c *gin.Context) {

	if h.db == nil {
		return
	}

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// GetOrders is a method that handles get orders for a specific user
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Charge is a method that handles credit card charging process
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}

	request := struct {
		models.Order
		Remember    bool   `json:"rememberCard"`
		UseExisting bool   `json:"useExisting"`
		Token       string `json:"token"`
	}{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}

	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	chargeP := &stripe.ChargeParams{
		//the price we obtained from the incoming request:
		Amount: stripe.Int64(int64(request.Price)),
		//the currency:
		Currency: stripe.String("usd"),
		//the description:
		Description: stripe.String("GoMusic charge..."),
	}

	stripeCustomerID := ""

	if request.UseExisting {
		//use existing
		log.Println("Getting credit card id...")
		//This is a new method which retrieve the stripe customer id from the database
		stripeCustomerID, err = h.db.GetCreditCardCID(request.CustomerID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cp := &stripe.CustomerParams{}
		cp.SetSource(request.Token)
		customer, err := customer.New(cp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stripeCustomerID = customer.ID
	}

	if request.Remember {
		//save the stripe customer id, and link it to the actual customer id in our database
		err = h.db.SaveCreditCardForCustomer(request.CustomerID, stripeCustomerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//Assign the stipe customer id to the *stripe.ChargeParams object:
	chargeP.Customer = stripe.String(stripeCustomerID)

	//Charge the credit card
	_, err = charge.New(chargeP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.AddOrder(request.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
