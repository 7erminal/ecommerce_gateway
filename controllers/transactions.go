package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// TransactionsController operations for Transactions
type TransactionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *TransactionsController) URLMapping() {
	c.Mapping("PlaceRentalRequest", c.PlaceRentalRequest)
	c.Mapping("PlaceSalesRequest", c.PlaceSalesRequest)
	c.Mapping("PlaceOrderRequest", c.PlaceOrderRequest)
	c.Mapping("GetAllTransactions", c.GetAllTransactions)
	c.Mapping("GetAllOrders", c.GetAllOrders)
	c.Mapping("GetOneOrder", c.GetOneOrder)
}

// PlaceRentalRequest ...
// @Title Place Rental Request
// @Description Place an order to rent a product
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.RentalRequestDTO	true		"body for Transactions content"
// @Success 201 {object} models.Transactions
// @Failure 403 body is empty
// @router /place-rental-request [post]
func (c *TransactionsController) PlaceRentalRequest() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if err != false {
		logs.Error("An error occurred ", err)
	}
	var v requests.RentalRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	requestType := "RENTAL"

	isSuccess := false

	itemImage := functions.UploadPaymentProof(&c.Controller, v.PaymentProofImageUrl)

	if itemImage.StatusCode == 200 {
		logs.Info("Payment proof uploaded successfully")
	}

	products := []requests.Item{}

	for _, i := range v.Products {
		tProduct := requests.Item{ItemId: i.ProductId, Quantity: i.Quantity}
		products = append(products, tProduct)
	}

	userIdStr := strconv.FormatInt(userData.UserId, 10)
	q := requests.PostTransactionRequest{
		Items:           products,
		RequestType:     requestType,
		PaymentMethodId: v.PaymentMethodId,
		Comment:         "",
		OrderLocation:   v.OrderLocation,
		OrderBy:         userIdStr,
		OrderStartDate:  v.OrderStartDate,
		OrderEndDate:    v.OrderEndDate,
		CustomerId:      v.CustomerId,
		CurrencyId:      "1",
	}

	fmt.Printf("Item request of v: %+v\n", q)

	postRequestResponse := functions.PlaceOrder(&c.Controller, q)

	message := "Unable to complete order"

	if postRequestResponse.StatusCode == 200 {
		isSuccess = true

		message = "Order completed successfully"
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: postRequestResponse.Transaction, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// PlaceSalesRequest ...
// @Title Place Sales Request
// @Description Place an order to buy a product
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.SalesRequestDTO	true		"body for Transactions content"
// @Success 201 {object} models.Transactions
// @Failure 403 body is empty
// @router /place-sales-request [post]
func (c *TransactionsController) PlaceSalesRequest() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if err != false {
		logs.Error("An error occurred ", err)
	}
	var v requests.SalesRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	requestType := "SALES"

	isSuccess := false

	products := []requests.Item{}

	for _, i := range v.Products {
		tProduct := requests.Item{ItemId: i.ProductId, Quantity: i.Quantity}
		products = append(products, tProduct)
	}

	userIdStr := strconv.FormatInt(userData.UserId, 10)
	q := requests.PostTransactionRequest{
		Items:           products,
		RequestType:     requestType,
		PaymentMethodId: v.PaymentMethodId,
		Comment:         "",
		OrderLocation:   "",
		OrderBy:         userIdStr,
		OrderStartDate:  "",
		OrderEndDate:    "",
		CustomerId:      v.CustomerId,
		CurrencyId:      "1",
	}

	postRequestResponse := functions.PlaceOrder(&c.Controller, q)

	message := "Unable to complete order"

	if postRequestResponse.StatusCode == 200 {
		isSuccess = true

		message = "Order completed successfully"
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: postRequestResponse.Transaction, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// PlaceOrderRequest ...
// @Title Place Order Request
// @Description Place an order request
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.OrderRequest2DTO	true		"body for Transactions content"
// @Success 201 {object} responses.TransactionResponseDTO
// @Failure 403 body is empty
// @router /place-order-request [post]
func (c *TransactionsController) PlaceOrderRequest() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if err != false {
		logs.Error("An error occurred ", err)
	}
	var v requests.OrderRequest2DTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	rjson, marshalErr := json.Marshal(v)
	if marshalErr != nil {
		logs.Error("Error marshalling user to JSON: ", marshalErr.Error())
	} else {
		logs.Info("Request received is \n", string(rjson))
	}

	requestType := "PURCHASE"

	isSuccess := false

	products := []requests.Item{}

	for _, i := range v.Products {
		tProduct := requests.Item{ItemId: i.ProductId, Quantity: i.Quantity}
		products = append(products, tProduct)
	}

	// Currency      string
	// Items         []Cart
	// RequestType   string
	// Comment       string
	// CreatedBy     string
	// OrderDate     string
	// OrderEndDate  string
	// OrderLocation string
	// Customer      string
	// Branch        string

	userIdStr := strconv.FormatInt(userData.UserId, 10)

	q := requests.PostTransactionRequest{
		Items:           products,
		RequestType:     requestType,
		PaymentMethodId: v.PaymentMethodId,
		Comment:         "",
		OrderLocation:   "",
		OrderBy:         userIdStr,
		OrderStartDate:  "",
		OrderEndDate:    "",
		CustomerId:      v.CustomerId,
		CurrencyId:      v.Currency,
	}

	postRequestResponse := functions.PlaceOrder(&c.Controller, q)

	message := "Unable to complete order"

	if postRequestResponse.StatusCode == 200 {
		isSuccess = true

		message = "Order completed successfully"
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: postRequestResponse.Transaction, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		resp := responses.TransactionResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllTransactions ...
// @Title Get All Transactions
// @Description Get all transactions
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.OrdersGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-all-transactions [get]
func (c *TransactionsController) GetAllTransactions() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if userData == nil {
		logs.Error("An error occurred ", err)
		return
	}

	userIdStr := strconv.FormatInt(userData.UserId, 10)

	logs.Info("About to get transactions. Ordering by desc")
	serviceResp := functions.GetTransactions(&c.Controller, requests.GetTransactionsRequestDTO{
		Limit:  "20",
		Offset: "0",
		Query:  "userId=" + userIdStr,
		Order:  "desc",
	})

	if serviceResp.StatusCode == 200 {
		orderResp := responses.OrdersGatewayResponseDTO{
			Success:    serviceResp.StatusCode == 200,
			Result:     serviceResp.Orders,
			StatusDesc: serviceResp.StatusDesc,
		}
		c.Data["json"] = orderResp
	} else {
		logs.Error("An error occurred while fetching transactions: ", serviceResp.StatusDesc)
		orderResp := responses.OrdersGatewayResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: "Failed to fetch transactions",
		}
		c.Data["json"] = orderResp
	}

	c.ServeJSON()
}

// GetAllOrders ...
// @Title Get All Orders
// @Description Get all orders
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.OrdersGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-orders [get]
func (c *TransactionsController) GetAllOrders() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	var order string

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if userData == nil {
		logs.Error("An error occurred ", err)
		return
	}

	if v := c.GetString("order"); v != "" {
		logs.Info("Order is ", v)
		order = v
	}

	userIdStr := strconv.FormatInt(userData.UserId, 10)

	logs.Info("Order from frontend is by ", order)
	serviceResp := functions.GetOrders(&c.Controller, requests.GetOrdersRequestDTO{
		Limit:  "10",
		Offset: "0",
		Query:  "userId=" + userIdStr,
		Order:  order,
	})

	if serviceResp.StatusCode == 200 {
		orderResp := responses.OrdersGatewayResponseDTO{
			Success:    serviceResp.StatusCode == 200,
			Result:     serviceResp.Orders,
			StatusDesc: serviceResp.StatusDesc,
		}
		c.Data["json"] = orderResp
	} else {
		logs.Error("An error occurred while fetching orders: ", serviceResp.StatusDesc)
		orderResp := responses.OrdersGatewayResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: "Failed to fetch orders",
		}
		c.Data["json"] = orderResp
	}

	c.ServeJSON()
}

// GetOneOrder ...
// @Title Get One Order
// @Description Get one order by id
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to get"
// @Success 200 {object} responses.OrderItemResponseDTO
// @Failure 403 id is empty
// @router /get-order/:id [get]
func (c *TransactionsController) GetOneOrder() {
	idStr := c.Ctx.Input.Param(":id")

	success := false
	statusDesc := "An error occurred while fetching order item"
	result := &responses.OrdersCustom{}

	orderResp := functions.GetOrder(&c.Controller, requests.GetOrderRequestDTO{
		OrderId: idStr,
	})

	if orderResp.StatusCode == 200 {
		// if err := functions.UpdateItem(&c.Controller, requests.UpdateItemRequest{ItemId: item.Item.ItemId, Status: status}); err != nil {
		// 	logs.Error("An Error occurred while updating item: ", err.Error())
		// }
		// orderIdStr := strconv.FormatInt(orderItem.OrderId, 10)

		o := responses.OrdersCustom{
			OrderId:      orderResp.Order.OrderId,
			OrderNumber:  orderResp.Order.OrderNumber,
			Quantity:     orderResp.Order.Quantity,
			Cost:         orderResp.Order.Cost,
			Currency:     orderResp.Order.Currency,
			OrderDate:    orderResp.Order.OrderDate,
			DateCreated:  orderResp.Order.DateCreated,
			DateModified: orderResp.Order.DateModified,
			Customer:     orderResp.Order.Customer,
			OrderDetails: orderResp.Order.OrderDetails,
		}

		result = &o
		success = true
		statusDesc = "Order fetched successfully"
		resp := responses.OrderGatewayResponseDTO{Success: success, StatusDesc: statusDesc, Result: result}
		c.Data["json"] = resp

	} else {
		logs.Error("An Error occurred while fetching order item: ", orderResp.StatusDesc)
		statusDesc = "An error occurred while fetching order item: " + orderResp.StatusDesc
		resp := responses.OrderGatewayResponseDTO{Success: success, StatusDesc: statusDesc, Result: nil}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
