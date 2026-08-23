package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// ItemsController operations for Items
type ItemsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ItemsController) URLMapping() {
	c.Mapping("AddSalesItem", c.AddSalesItem)
	c.Mapping("UpdateItemImage", c.UpdateItemImage)
	c.Mapping("AddCategory", c.AddCategory)
	c.Mapping("GetCategories", c.GetCategories)
	c.Mapping("GetItems", c.GetItems)
	c.Mapping("GetProduct", c.GetProduct)
	c.Mapping("UpdateItem", c.UpdateItem)
	c.Mapping("AddRentalsItem", c.AddRentalsItem)
	c.Mapping("AddPurpose", c.AddPurpose)
	c.Mapping("AddFeature", c.AddFeature)
}

// AddSalesItem ...
// @Title Add Item
// @Description Add an item for sale
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.AddSalesItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /add-sales-product [post]
func (c *ItemsController) AddSalesItem() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)
	logs.Info("Error is ", err)
	var v requests.AddSalesItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var isSuccess bool = false

	logs.Info("Received \nProduct name: ", v.ProductName, "Branch ID:: ", userData.UserDetails.Branch.BranchId, "Cost price:: ", v.CostPrice, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Selling price:: ", v.SellingPrice)
	proceed := true
	errorMessage := "An error occurred"
	logs.Info("Token verified!")
	branchId := strconv.FormatInt(userData.UserDetails.Branch.BranchId, 10)

	getBranchResp, erri := functions.GetSystemDetails(&c.Controller, branchId)
	if erri != nil {
		errorMessage = "Failed to fetch branch details"
		proceed = false
	}

	if proceed == true {
		if getBranchResp.Success != true {
			errorMessage = getBranchResp.StatusDesc
			proceed = false
		}
	}

	getProductTypes := functions.GetCategory(&c.Controller, strconv.FormatInt(v.CategoryId, 10))
	if getProductTypes.StatusCode != 200 {
		errorMessage = "Product type provided does not exist"
		proceed = false
	}

	if proceed {
		req := requests.AddItemRequestDTO{
			ProductName:     v.ProductName,
			Description:     v.Description,
			Weight:          v.Weight,
			Quantity:        v.Quantity,
			ReorderLevel:    0,
			CostPrice:       v.CostPrice,
			SellingPrice:    v.SellingPrice,
			BranchId:        userData.UserDetails.Branch.BranchId,
			ImagePath:       v.ImagePath,
			AvailableSizes:  v.AvailableSizes,
			AvailableColors: v.AvailableColors,
			Purposes:        v.Purposes,
			Features:        v.Features,
			Country:         "GHA",
		}
		addItemResp := functions.AddItem(
			&c.Controller,
			req,
			getProductTypes.Category.CategoryId,
			getBranchResp.Result.Branch.Country.CountryCode,
			userData.UserDetails.Branch.BranchId,
			int(userData.UserId))

		itemResp := responses.Item{}
		if addItemResp.StatusCode == 200 {
			if addItemResp.Item != nil {
				logs.Info("About to update item image")
				itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
				if itemImageUpdateResp.StatusCode == 200 {
					logs.Info("Successfully updated item image")
				} else {
					logs.Error("Failed update")
				}

				var features []responses.Feature

				if v.Features != nil {
					if len(*v.Features) > 0 {
						for _, featureId := range *v.Features {

							var ifResp requests.AddProductFeatureRequestDTO = requests.AddProductFeatureRequestDTO{
								ProductId: addItemResp.Item.ItemId,
								FeatureId: featureId,
							}

							addItemFeatureResp := functions.AddItemFeatures(&c.Controller, ifResp)

							if addItemFeatureResp.StatusCode == 200 {
								logs.Info("Successfully added item feature")
								features = append(features, *addItemFeatureResp.Result.Feature)
							} else {
								logs.Error("Failed to add item feature")
							}
						}
					}
				}

				var purposes []responses.Purpose
				if v.Purposes != nil {
					if len(*v.Purposes) > 0 {
						for _, purposeId := range *v.Purposes {

							var ipResp requests.AddProductPurposeRequestDTO = requests.AddProductPurposeRequestDTO{
								ProductId: addItemResp.Item.ItemId,
								PurposeId: purposeId,
							}

							addItemPurposeResp := functions.AddItemPurposes(&c.Controller, ipResp)

							if addItemPurposeResp.StatusCode == 200 {
								logs.Info("Successfully added item purpose")
								purposes = append(purposes, *addItemPurposeResp.Result.Purpose)
							} else {
								logs.Error("Failed to add item purpose")
							}
						}
					}
				}

				availableSizes := strings.Split(addItemResp.Item.AvailableSizes, ",")
				availableColors := strings.Split(addItemResp.Item.AvailableColors, ",")

				itemResp = responses.Item{
					ProductId:        addItemResp.Item.ItemId,
					ProductName:      addItemResp.Item.ItemName,
					Description:      addItemResp.Item.Description,
					ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
					ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
					ImagePath:        itemImageUpdateResp.Item.ImagePath,
					Quantity:         addItemResp.Item.Quantity,
					Branch:           addItemResp.Item.Branch,
					Category:         addItemResp.Item.Category,
					AvailableSizes:   &availableSizes,
					AvailableColors:  &availableColors,
					Purposes:         &purposes,
					Features:         &features,
				}

				isSuccess = true
			} else {
				isSuccess = false
			}

		}

		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

		c.Data["json"] = resp
		// } else {
		// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

		// 	c.Data["json"] = resp
		// }
	} else {
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// AddRentalsItem ...
// @Title Add Rentals Item
// @Description Add an item for rent
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.AddRentalItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /add-rental-product [post]
func (c *ItemsController) AddRentalsItem() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)
	logs.Info("Error is ", err)
	// userIdStr := strconv.FormatInt(userData.UserId, 10)
	var v requests.AddRentalItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var isSuccess bool = false

	logs.Info("Received \nProduct name: ", v.ProductName, "Branch ID:: ", userData.UserDetails.Branch.BranchId, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Rental price:: ", v.RentalPrice)
	proceed := true
	errorMessage := "An error occurred"
	logs.Info("Token verified!")
	logs.Info("Branch is !", userData.UserDetails.Branch)
	branchId := strconv.FormatInt(userData.UserDetails.Branch.BranchId, 10)
	getBranchResp, erri := functions.GetSystemDetails(&c.Controller, branchId)
	if erri != nil {
		errorMessage = "Failed to fetch branch details"
		proceed = false
	}

	if proceed == true {
		if getBranchResp.Success != true {
			errorMessage = getBranchResp.StatusDesc
			proceed = false
		}
	}

	sales_product_type_name, _ := beego.AppConfig.String("rentalsProductType")

	getProductTypes := functions.GetCategoryByName(&c.Controller, sales_product_type_name)

	if getProductTypes.StatusCode != 200 {
		errorMessage = "Product type provided does not exist"
		proceed = false
	}

	if proceed {
		req := requests.AddItemRequestDTO{ProductName: v.ProductName, Quantity: v.Quantity, ReorderLevel: v.ReorderLevel, CostPrice: 0, SellingPrice: v.RentalPrice, BranchId: userData.UserDetails.Branch.BranchId, ImagePath: v.ImagePath}
		addItemResp := functions.AddItem(&c.Controller, req, getProductTypes.Category.CategoryId, getBranchResp.Result.Branch.Country.CountryCode, userData.UserDetails.Branch.BranchId, int(userData.UserId))

		itemResp := responses.Item{}
		if addItemResp.StatusCode == 200 {
			if addItemResp.Item != nil {
				logs.Info("About to update item image")
				itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
				if itemImageUpdateResp.StatusCode == 200 {
					logs.Info("Successfully updated item image")
				} else {
					logs.Error("Failed update")
				}

				itemResp = responses.Item{
					ProductId:        addItemResp.Item.ItemId,
					ProductName:      addItemResp.Item.ItemName,
					Description:      addItemResp.Item.Description,
					ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
					ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
					ImagePath:        itemImageUpdateResp.Item.ImagePath,
					Quantity:         addItemResp.Item.Quantity,
					Branch:           addItemResp.Item.Branch,
				}

				isSuccess = true
			} else {
				isSuccess = false
			}

		}

		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

		c.Data["json"] = resp
		// } else {
		// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

		// 	c.Data["json"] = resp
		// }
	} else {
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateItem ...
// @Title Update Item
// @Description Add an item
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /update-product/:id [put]
func (c *ItemsController) UpdateItem() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)
	logs.Info("Error is ", err)
	var v requests.UpdateItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	idStr := c.Ctx.Input.Param(":id")

	logs.Info("Received \nProduct name: ", v.ProductName, "Branch ID:: ", v.BranchId, "Cost price:: ", v.CostPrice, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Selling price:: ", v.SellingPrice)

	var isSuccess bool = false

	proceed := true
	errorMessage := "An error occurred"
	logs.Info("Token verified!")
	branchId := strconv.FormatInt(v.BranchId, 10)
	getBranchResp, erri := functions.GetSystemDetails(&c.Controller, branchId)
	if erri != nil {
		errorMessage = "Failed to fetch branch details"
		proceed = false
	}

	if proceed == true {
		if getBranchResp.Success != true {
			errorMessage = getBranchResp.StatusDesc
			proceed = false
		}
	}

	categoryId := strconv.FormatInt(v.CategoryId, 10)

	getProductTypes := functions.GetCategory(&c.Controller, categoryId)

	if getProductTypes.StatusCode != 200 {
		errorMessage = "Product type provided does not exist"
		proceed = false
	}

	getItemResp := functions.GetItem(&c.Controller, idStr)

	if getItemResp.StatusCode != 200 {
		errorMessage = "Item provided does not exist"
		proceed = false
	}

	if proceed {
		addItemResp := functions.UpdateItem(&c.Controller, v, getBranchResp.Result.Branch.Country.CountryCode, v.BranchId, int(userData.UserId), idStr)

		itemResp := responses.Item{}
		if addItemResp.StatusCode == 200 {
			if addItemResp.Item != nil {
				logs.Info("About to update item image")
				itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
				if itemImageUpdateResp.StatusCode == 200 {
					logs.Info("Successfully updated item image")
				} else {
					logs.Error("Failed update")
				}

				itemFeaturesResp := functions.GetItemFeatures(&c.Controller, requests.AddProductFeatureRequestDTO{ProductId: addItemResp.Item.ItemId})
				if itemFeaturesResp.StatusCode == 200 {
					logs.Info("Successfully fetched item features")
					for _, feature := range *itemFeaturesResp.Result {
						logs.Info("Feature is ", feature.Feature)
						exists := false
						for _, featureId := range *v.Features {
							if feature.Feature.FeatureId == featureId {
								exists = true
							}
						}

						if !exists {
							logs.Info("Feature does not exist, deleting ", feature.Feature.FeatureId)
							deleteResp := functions.DeleteItemFeature(&c.Controller, strconv.FormatInt(feature.ItemFeatureId, 10))
							if deleteResp.StatusCode == 200 {
								logs.Info("Successfully deleted item feature")
							} else {
								logs.Error("Failed to delete item feature")
							}
						}
					}

					exists := false
					for _, featureId := range *v.Features {
						for _, feature := range *itemFeaturesResp.Result {
							if featureId == feature.Feature.FeatureId {
								exists = true
							}
						}

						if !exists {
							logs.Info("Feature does not exist, adding ", featureId)
							var ifResp requests.AddProductFeatureRequestDTO = requests.AddProductFeatureRequestDTO{
								ProductId: addItemResp.Item.ItemId,
								FeatureId: featureId,
							}

							addItemFeatureResp := functions.AddItemFeatures(&c.Controller, ifResp)

							if addItemFeatureResp.StatusCode == 200 {
								logs.Info("Successfully added item feature")
							} else {
								logs.Error("Failed to add item feature")
							}
						}
					}
				} else {
					logs.Error("Failed to fetch item features")
				}

				itemPurposesResp := functions.GetItemPurposes(&c.Controller, requests.AddProductPurposeRequestDTO{ProductId: addItemResp.Item.ItemId})
				if itemPurposesResp.StatusCode == 200 {
					logs.Info("Successfully fetched item purposes")
					for _, purpose := range *itemPurposesResp.Result {
						logs.Info("Purpose is ", purpose.Purpose)
						exists := false
						for _, purposeId := range *v.Purposes {
							if purpose.Purpose.PurposeId == purposeId {
								exists = true
							}
						}

						if !exists {
							logs.Info("Purpose does not exist, deleting ", purpose.Purpose.PurposeId)
							deleteResp := functions.DeleteItemPurpose(&c.Controller, strconv.FormatInt(purpose.ItemPurposeId, 10))
							if deleteResp.StatusCode == 200 {
								logs.Info("Successfully deleted item purpose")
							} else {
								logs.Error("Failed to delete item purpose")
							}
						}
					}

					exists := false
					for _, purposeId := range *v.Purposes {
						for _, purpose := range *itemPurposesResp.Result {
							if purposeId == purpose.Purpose.PurposeId {
								exists = true
							}
						}

						if !exists {
							logs.Info("Purpose does not exist, adding ", purposeId)
							var ifResp requests.AddProductPurposeRequestDTO = requests.AddProductPurposeRequestDTO{
								ProductId: addItemResp.Item.ItemId,
								PurposeId: purposeId,
							}

							addItemPurposeResp := functions.AddItemPurposes(&c.Controller, ifResp)

							if addItemPurposeResp.StatusCode == 200 {
								logs.Info("Successfully added item purpose")
							} else {
								logs.Error("Failed to add item purpose")
							}
						}
					}
				} else {
					logs.Error("Failed to fetch item purposes")
				}

				availableSizes := strings.Split(addItemResp.Item.AvailableSizes, ",")
				availableColors := strings.Split(addItemResp.Item.AvailableColors, ",")

				itemResp = responses.Item{
					ProductId:        addItemResp.Item.ItemId,
					ProductName:      addItemResp.Item.ItemName,
					Description:      addItemResp.Item.Description,
					Weight:           addItemResp.Item.Weight,
					ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
					ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
					ImagePath:        itemImageUpdateResp.Item.ImagePath,
					Quantity:         addItemResp.Item.Quantity,
					Branch:           addItemResp.Item.Branch,
					Category:         addItemResp.Item.Category,
					AvailableSizes:   &availableSizes,
					AvailableColors:  &availableColors,
					// Features:         &features,
					// Purposes:         &purposes,
				}

				isSuccess = true
			} else {
				isSuccess = false
			}

		}

		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

		c.Data["json"] = resp
		// } else {
		// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

		// 	c.Data["json"] = resp
		// }
	} else {
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateImage ...
// @Title UpdateImage
// @Description Update User's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /upload-product-image [post]
func (c *ItemsController) UpdateItemImage() {

	var isSuccess bool = false

	image, header, err := c.GetFile("Image")

	if err != nil {
		var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "Image", image, *header)

		if respCode == 200 {
			itemImage := functions.UploadItemImage(&c.Controller, filePath)

			if itemImage.StatusCode == 200 {
				logs.Info("Item image returned: ", itemImage.Value)

				isSuccess = true

				var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: itemImage.Value, StatusDesc: itemImage.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to upload file. Tmp"}
			c.Data["json"] = resp
		}

	}

	c.ServeJSON()
}

// AddCategory ...
// @Title Add Category
// @Description Add Category
// @Param	Authorization		header 	string true		"header for User"
// @Param	CategoryImage		formData 	file	true		"Category Image"
// @Param	CategoryName		formData 	string	true		"Category name"
// @Success 200 {object} responses.CategoryResponseDTO
// @Failure 403 body is empty
// @router /add-category [post]
func (c *ItemsController) AddCategory() {
	var isSuccess bool = false

	image, header, err := c.GetFile("CategoryImage")

	if err != nil {
		var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "CategoryImage", image, *header)

		if respCode == 200 {

			categoryName := c.Ctx.Input.Query("CategoryName")
			categoryDescription := c.Ctx.Input.Query("CategoryDescription")

			categoryResp := functions.AddCategory(&c.Controller, filePath, categoryName, categoryDescription)

			if categoryResp.StatusCode == 200 {

				isSuccess = true

				var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: categoryResp.Category, StatusDesc: categoryResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. File upload failed"}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// AddFeature ...
// @Title Add Feature
// @Description Add Feature
// @Param	Authorization		header 	string true		"header for User"
// @Param	FeatureImage		formData 	file	true		"Feature Image"
// @Param	FeatureName		formData 	string	true		"Feature name"
// @Success 200 {object} responses.CategoryResponseDTO
// @Failure 403 body is empty
// @router /add-feature [post]
func (c *ItemsController) AddFeature() {
	var isSuccess bool = false

	image, header, err := c.GetFile("FeatureImage")

	if err != nil {
		var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "FeatureImage", image, *header)

		if respCode == 200 {

			featureName := c.Ctx.Input.Query("FeatureName")
			featureDescription := c.Ctx.Input.Query("FeatureDescription")

			featureResp := functions.AddFeature(&c.Controller, filePath, featureName, featureDescription)

			if featureResp.StatusCode == 200 {

				isSuccess = true

				var resp responses.FeaturesResponseDTO = responses.FeaturesResponseDTO{Success: isSuccess, Result: featureResp.Features, StatusDesc: featureResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.FeaturesResponseDTO = responses.FeaturesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.FeaturesResponseDTO = responses.FeaturesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. File upload failed"}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// AddPurpose ...
// @Title Add Purpose
// @Description Add Purpose
// @Param	Authorization		header 	string true		"header for User"
// @Param	PurposeImage		formData 	file	true		"Purpose Image"
// @Param	PurposeName		formData 	string	true		"Purpose name"
// @Success 200 {object} responses.CategoryResponseDTO
// @Failure 403 body is empty
// @router /add-purpose [post]
func (c *ItemsController) AddPurpose() {
	var isSuccess bool = false

	image, header, err := c.GetFile("PurposeImage")

	if err != nil {
		var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "PurposeImage", image, *header)

		if respCode == 200 {

			purposeName := c.Ctx.Input.Query("PurposeName")
			purposeDescription := c.Ctx.Input.Query("PurposeDescription")

			purposeResp := functions.AddPurpose(&c.Controller, filePath, purposeName, purposeDescription)

			if purposeResp.StatusCode == 200 {

				isSuccess = true

				var resp responses.PurposesResponseDTO = responses.PurposesResponseDTO{Success: isSuccess, Result: purposeResp.Purposes, StatusDesc: purposeResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.PurposesResponseDTO = responses.PurposesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.PurposesResponseDTO = responses.PurposesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. File upload failed"}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// GetCategories ...
// @Title Get Categories
// @Description Get Categories
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.CategoriesResponseDTO
// @Failure 403 body is empty
// @router /get-categories [get]
func (c *ItemsController) GetCategories() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = true

	categoryResponse := functions.GetCategories(&c.Controller)

	if categoryResponse.StatusCode == 200 {
		logs.Info("Categories returned: ", categoryResponse.Categories)

		isSuccess = true

		var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: categoryResponse.Categories, StatusDesc: "Categories fetched successfully"}
		c.Data["json"] = resp
	} else {
		var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetFeatures ...
// @Title Get Features
// @Description Get Features
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.FeaturesResponseDTO
// @Failure 403 body is empty
// @router /get-features [get]
func (c *ItemsController) GetFeatures() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = true

	featureResponse := functions.GetFeatures(&c.Controller)

	if featureResponse.StatusCode == 200 {
		logs.Info("Features returned: ", featureResponse.Features)

		isSuccess = true

		var resp responses.FeaturesResponseDTO = responses.FeaturesResponseDTO{Success: isSuccess, Result: featureResponse.Features, StatusDesc: "Features fetched successfully"}
		c.Data["json"] = resp
	} else {
		var resp responses.FeaturesResponseDTO = responses.FeaturesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetPurposes ...
// @Title Get Purposes
// @Description Get Purposes
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.PurposesResponseDTO
// @Failure 403 body is empty
// @router /get-purposes [get]
func (c *ItemsController) GetPurposes() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = true

	purposeResponse := functions.GetPurposes(&c.Controller)

	if purposeResponse.StatusCode == 200 {
		logs.Info("Purposes returned: ", purposeResponse.Purposes)

		isSuccess = true

		var resp responses.PurposesResponseDTO = responses.PurposesResponseDTO{Success: isSuccess, Result: purposeResponse.Purposes, StatusDesc: "Purposes fetched successfully"}
		c.Data["json"] = resp
	} else {
		var resp responses.PurposesResponseDTO = responses.PurposesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetItems ...
// @Title Get Items
// @Description Get Items
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.ItemsResponseDTO
// @Failure 403 body is empty
// @router /get-items [get]
func (c *ItemsController) GetItems() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			var fields string
			var sortby string
			var order string
			var query string
			var limit string
			var offset string

			// limit: 10 (default is 10)
			if v := c.GetString("limit"); v != "" {
				limit = v
			}
			// offset: 0 (default is 0)
			if v := c.GetString("offset"); v != "" {
				offset = v
			}
			// sortby: col1,col2
			if v := c.GetString("sortby"); v != "" {
				sortby = v
			}
			// order: desc,asc
			if v := c.GetString("order"); v != "" {
				order = v
			}
			// query: k:v,k:v
			if v := c.GetString("query"); v != "" {
				query = v
			}
			logs.Info("Success response received")
			isSuccess = false

			logs.Info("User data received ", verifyToken.User.UserDetails.Branch)

			if verifyToken.User.UserDetails.Branch != nil {
				branchId := strconv.FormatInt(verifyToken.User.UserDetails.Branch.BranchId, 10)

				// Depending on the role, fetch items
				var getItemsResp responses.ItemsOriResponseDTO
				if verifyToken.User.Role.Role == "SUPER_ADMIN" {
					getItemsResp = functions.GetItems(&c.Controller, query, fields, sortby, order, offset, limit)
				} else {
					getItemsResp = functions.GetItemsByBranch(&c.Controller, branchId, query, fields, sortby, order, offset, limit)
				}

				if getItemsResp.StatusCode == 200 {
					logs.Info("Items returned: ", getItemsResp.Items)

					items := []responses.Item{}
					if getItemsResp.Items != nil && len(*getItemsResp.Items) > 0 {
						for _, item := range *getItemsResp.Items {
							availableSizes := strings.Split(item.AvailableSizes, ",")
							availableColors := strings.Split(item.AvailableColors, ",")
							itemT := responses.Item{
								ProductId:        item.ItemId,
								ProductName:      item.ItemName,
								Description:      item.Description,
								ProductPrice:     float64(item.ItemPrice.ItemPrice),
								ProductCostPrice: float64(item.ItemPrice.AltItemPrice),
								ImagePath:        item.ImagePath,
								Quantity:         item.Quantity,
								Branch:           item.Branch,
								Category:         item.Category,
								AvailableSizes:   &availableSizes,
								AvailableColors:  &availableColors,
								Features:         item.ItemFeatures,
								Purposes:         item.ItemPurposes,
								Status:           "ACTIVE",
							}

							items = append(items, itemT)
						}

					} else {
						items = []responses.Item{}
					}

					isSuccess = true

					data := responses.ItemsData{}
					data.Data = &items
					data.Count = len(items)

					var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: &data, StatusDesc: getItemsResp.StatusDesc}
					c.Data["json"] = resp
				} else {
					var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
					c.Data["json"] = resp
				}
			} else {
				logs.Error("User is not linked to a branch")
				var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. User is not linked to a branch"}
				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetItem ...
// @Title Get Item
// @Description Get Item
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /get-item/:id [get]
func (c *ItemsController) GetProduct() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = false

	idStr := c.Ctx.Input.Param(":id")

	itemResp := functions.GetItem(&c.Controller, idStr)

	item := responses.Item{}
	if itemResp.StatusCode == 200 {
		// logs.Info("Categories returned: ", categoryResponse.Categories)
		availableSizes := strings.Split(itemResp.Item.AvailableSizes, ",")
		availableColors := strings.Split(itemResp.Item.AvailableColors, ",")
		item = responses.Item{
			ProductId:        itemResp.Item.ItemId,
			ProductName:      itemResp.Item.ItemName,
			Description:      itemResp.Item.Description,
			ProductPrice:     float64(itemResp.Item.ItemPrice.ItemPrice),
			ProductCostPrice: float64(itemResp.Item.ItemPrice.AltItemPrice),
			ImagePath:        itemResp.Item.ImagePath,
			Quantity:         itemResp.Item.Quantity,
			Branch:           itemResp.Item.Branch,
			Status:           "ACTIVE",
			Category:         itemResp.Item.Category,
			Features:         itemResp.Item.ItemFeatures,
			Purposes:         itemResp.Item.ItemPurposes,
			AvailableSizes:   &availableSizes,
			AvailableColors:  &availableColors,
		}

		isSuccess = true

		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &item, StatusDesc: itemResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// DeleteCategory ...
// @Title Delete Category
// @Description Delete Category
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /delete-category/:id [delete]
func (c *ItemsController) DeleteCategory() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = false

	idStr := c.Ctx.Input.Param(":id")

	deleteCategoryResp := functions.DeleteCategory(&c.Controller, idStr)

	if deleteCategoryResp.StatusCode == 200 {
		isSuccess = true

		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &deleteCategoryResp.StatusDesc, StatusDesc: deleteCategoryResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// DeleteFeature ...
// @Title Delete Feature
// @Description Delete Feature
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /delete-feature/:id [delete]
func (c *ItemsController) DeleteFeature() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = false

	idStr := c.Ctx.Input.Param(":id")

	deleteFeatureResp := functions.DeleteFeature(&c.Controller, idStr)

	if deleteFeatureResp.StatusCode == 200 {
		isSuccess = true

		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &deleteFeatureResp.StatusDesc, StatusDesc: deleteFeatureResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// DeletePurpose ...
// @Title Delete Purpose
// @Description Delete Purpose
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /delete-purpose/:id [delete]
func (c *ItemsController) DeletePurpose() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = false

	idStr := c.Ctx.Input.Param(":id")

	deletePurposeResp := functions.DeletePurpose(&c.Controller, idStr)

	if deletePurposeResp.StatusCode == 200 {
		isSuccess = true

		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &deletePurposeResp.StatusDesc, StatusDesc: deletePurposeResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// DeleteItem ...
// @Title Delete Item
// @Description Delete Item
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /delete-item/:id [delete]
func (c *ItemsController) DeleteItem() {
	var isSuccess bool = false

	logs.Info("Success response received")
	isSuccess = false

	idStr := c.Ctx.Input.Param(":id")

	deleteItemResp := functions.DeleteItem(&c.Controller, idStr)

	if deleteItemResp.StatusCode == 200 {
		isSuccess = true

		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &deleteItemResp.StatusDesc, StatusDesc: deleteItemResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
