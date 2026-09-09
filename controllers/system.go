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

// SystemController operations for System
type SystemController struct {
	beego.Controller
}

// URLMapping ...
func (c *SystemController) URLMapping() {
	c.Mapping("AddBranch", c.AddBranch)
	c.Mapping("GetOneBranch", c.GetOneBranch)
	c.Mapping("GetAllBranches", c.GetAllBranches)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UpdateBranch", c.UpdateBranch)
	c.Mapping("GetRoles", c.GetRoles)
	c.Mapping("GetSystemDetails", c.GetSystemDetails)
	c.Mapping("AddApplication", c.AddApplication)
	c.Mapping("UpdateApplication", c.UpdateApplication)
	c.Mapping("GetApplication", c.GetApplication)
	c.Mapping("AddTheme", c.AddTheme)
	c.Mapping("UpdateTheme", c.UpdateTheme)
	c.Mapping("GetApplications", c.GetApplications)
}

// GetRoles ...
// @Title Get Roles
// @Description Get all roles
// @Success 200 {object} responses.RolesAllGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-roles [get]
func (c *SystemController) GetRoles() {

	var isSuccess bool = false

	rolesResp := functions.GetRoles(&c.Controller)

	// var message string

	if rolesResp.StatusCode == 200 {

		isSuccess = true
		// message = "Email sent"

		var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: rolesResp.Roles, StatusDesc: rolesResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetSystemDetails ...
// @Title Get System Details
// @Description Get system details
// @Param	branchid		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.SystemDetailsResponseDTO
// @Failure 403 body is empty
// @router /get-system-details/:branchid [get]
func (c *SystemController) GetSystemDetails() {

	var isSuccess bool = false
	message := "An Error occurred"
	resp := responses.SystemDetailsData{}

	branchId := c.Ctx.Input.Param(":branchid")
	systemDetails, err := functions.GetSystemDetails(&c.Controller, branchId)

	// var message string

	if err == nil && systemDetails.Success == true {

		isSuccess = true
		// message = "Email sent"

		resp = *systemDetails.Result
		message = systemDetails.StatusDesc

		var resp responses.SystemDetailsResponseDTO = responses.SystemDetailsResponseDTO{Success: isSuccess, Result: &resp, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		var resp responses.SystemDetailsResponseDTO = responses.SystemDetailsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create System
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.BranchRequestDTO	true		"body for System content"
// @Success 200 {object} responses.BranchResponseDTO
// @Failure 403 body is empty
// @router /add-branch [post]
func (c *SystemController) AddBranch() {
	var v requests.BranchRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			userDetailsResp := functions.GetUserDetails(&c.Controller, v.BranchManager)
			if userDetailsResp.StatusCode == 200 {
				addBranchResp := functions.AddBranch(&c.Controller, v, verifyToken.User.UserId)

				if addBranchResp.StatusCode == 200 {
					// Assign branch manager to added branch
					splitName := strings.Split(userDetailsResp.User.FullName, " | ")
					firstname := ""
					lastname := ""
					if len(splitName) > 1 {
						firstname = splitName[0]
						lastname = splitName[1]
					} else {
						firstname = splitName[0]
					}
					userDetails := requests.UpdateUserRequestDTO{BranchId: addBranchResp.Result.BranchId, FirstName: firstname, LastName: lastname, Username: userDetailsResp.User.Username, PhoneNumber: userDetailsResp.User.PhoneNumber, Gender: userDetailsResp.User.Gender, Dob: userDetailsResp.User.Dob.GoString(), Address: userDetailsResp.User.Address}
					userId := strconv.FormatInt(userDetailsResp.User.UserId, 10)
					updateUserResp := functions.UpdateUser(&c.Controller, userId, userDetails)
					branchIdStr := strconv.FormatInt(addBranchResp.Result.BranchId, 10)
					updateBranchResp := functions.UpdateBranchBranchManger(&c.Controller, userId, branchIdStr)
					message := "Branch Added Successfully"
					if updateUserResp.StatusCode != 200 {
						message = "Branch added but failed to assign manager"
					}
					if updateBranchResp.StatusCode != 200 {
						message = "Branch added but failed to assign branch manager"
					}
					// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: addBranchResp.Branch.Country.DefaultCurrency.Symbol, Currency: addBranchResp.Branch.Country.DefaultCurrency.Currency}
					// var country responses.CountryResp = responses.CountryResp{Country: addBranchResp.Branch.Country.Country, CountryCode: addBranchResp.Branch.Country.CountryCode, Currency: &curr}
					var data responses.BranchResp = responses.BranchResp{
						BranchId:    addBranchResp.Result.BranchId,
						Branch:      addBranchResp.Result.BranchName,
						Description: addBranchResp.Result.Description,
						// Country:     &country,
						Location:    addBranchResp.Result.Location,
						PhoneNumber: addBranchResp.Result.PhoneNumber,
					}

					isSuccess = true
					var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: &data, StatusDesc: message}
					c.Data["json"] = resp
				} else {
					var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: addBranchResp.StatusDesc}
					c.Data["json"] = resp
				}
			} else {
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Branch manager does not exist"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred." + verifyToken.StatusDesc}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOneBranch ...
// @Title GetOneBranch
// @Description get Branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.BranchOriResponseDTO
// @Failure 403 :id is empty
// @router /get-branch/:id [get]
func (c *SystemController) GetOneBranch() {
	idStr := c.Ctx.Input.Param(":id")

	var isSuccess bool = false

	getBranchResp, err := functions.GetSystemDetails(&c.Controller, idStr)

	if err != nil {
		logs.Error("An error occurred fetching branch details: ", err)
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	if getBranchResp.Success == true {
		// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: getBranchResp.Branch.Country.DefaultCurrency.Symbol, Currency: getBranchResp.Branch.Country.DefaultCurrency.Currency}
		// var country responses.CountryResp = responses.CountryResp{Country: getBranchResp.Branch.Country.Country, CountryCode: getBranchResp.Branch.Country.CountryCode, Currency: &curr}

		var curr responses.CurrencyResp = responses.CurrencyResp{
			CurrencyId: getBranchResp.Result.Branch.Country.Currency.CurrencyId,
			Symbol:     getBranchResp.Result.Branch.Country.Currency.Symbol,
			Currency:   getBranchResp.Result.Branch.Country.Currency.Currency}

		var country responses.CountryResp = responses.CountryResp{
			Country:     getBranchResp.Result.Branch.Country.Country,
			CountryCode: getBranchResp.Result.Branch.Country.CountryCode,
			Currency:    &curr,
		}
		var data responses.BranchResp = responses.BranchResp{
			BranchId:    getBranchResp.Result.Branch.BranchId,
			Branch:      getBranchResp.Result.Branch.Branch,
			Country:     &country,
			Location:    getBranchResp.Result.Branch.Location,
			PhoneNumber: getBranchResp.Result.Branch.PhoneNumber,
			DateCreated: getBranchResp.Result.Branch.DateCreated,
		}

		isSuccess = true
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: &data, StatusDesc: "Branch details fetched Successfully"}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred fetching branches from api")
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllCountries ...
// @Title GetAllCountries
// @Description get Countries
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.CountriesResponseDTO
// @Failure 403
// @router /get-countries [get]
func (c *SystemController) GetAllCountries() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			getCountriesResp := functions.GetCountries(&c.Controller)

			// json.Unmarshal(getBranchResp.Branches, &v)

			if getCountriesResp.StatusCode == 200 {
				logs.Info("Response is 200")
				var countries []responses.CountryResp
				if getCountriesResp.Countries != nil && len(*getCountriesResp.Countries) > 0 {
					for _, country := range *getCountriesResp.Countries {

						var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: country.Currency.Symbol, Currency: country.Currency.Currency}
						var countryResp responses.CountryResp = responses.CountryResp{Country: country.Country, CountryCode: country.CountryCode, Currency: &curr}

						countries = append(countries, countryResp)
					}
				} else {
					countries = []responses.CountryResp{}
				}

				isSuccess = true
				var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: &countries, StatusDesc: "Countries fetched Successfully"}
				c.Data["json"] = resp
			} else {
				logs.Error("An error occurred fetching countries from api")
				var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllBranches ...
// @Title GetAllBranches
// @Description get branches
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.BranchesResponseDTO
// @Failure 403
// @router /get-branches [get]
func (c *SystemController) GetAllBranches() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			getBranchResp := functions.GetBranches(&c.Controller)

			// json.Unmarshal(getBranchResp.Branches, &v)

			if getBranchResp.StatusCode == 200 {
				logs.Info("Response is 200")
				var branches []responses.BranchResp
				if getBranchResp.Branches != nil && len(*getBranchResp.Branches) > 0 {
					for _, branch := range *getBranchResp.Branches {

						var branchManager *responses.UserGateway

						logs.Info("Branch value is ", branch.BranchManager)
						if branch.BranchManager != nil {
							splitName := strings.Split(branch.BranchManager.FullName, " | ")
							firstname := ""
							lastname := ""
							if len(splitName) > 1 {
								firstname = splitName[0]
								lastname = splitName[1]
							} else {
								firstname = splitName[0]
							}

							// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: branch.Country.DefaultCurrency.Symbol, Currency: branch.Country.DefaultCurrency.Currency}
							// var country responses.CountryResp = responses.CountryResp{Country: branch.Country.Country, CountryCode: branch.Country.CountryCode, Currency: &curr}
							branchManager = &responses.UserGateway{UserId: branch.BranchManager.UserId, FirstName: firstname, LastName: lastname, Username: branch.BranchManager.Username, Email: branch.BranchManager.Email, PhoneNumber: branch.BranchManager.PhoneNumber, ImagePath: branch.BranchManager.ImagePath}
						} else {
							branchManager = nil
						}

						var data responses.BranchResp = responses.BranchResp{
							BranchId:    branch.BranchId,
							Branch:      branch.BranchName,
							Description: branch.Description,
							// Country:     &country,
							Location:      branch.Location,
							PhoneNumber:   branch.PhoneNumber,
							DateCreated:   branch.DateCreated,
							BranchManager: branchManager,
						}

						branches = append(branches, data)
					}
				} else {
					branches = []responses.BranchResp{}
				}

				branchesData := responses.BranchesData{}
				branchesData.Data = &branches
				branchesData.Count = len(branches)

				isSuccess = true
				var resp responses.BranchesResponseDTO = responses.BranchesResponseDTO{Success: isSuccess, Result: &branchesData, StatusDesc: "Branches fetched Successfully"}
				c.Data["json"] = resp
			} else {
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateBranch ...
// @Title Update Branch
// @Description update a Branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.BranchRequestDTO	true		"body for Branches content"
// @Success 200 {object} responses.BranchResponseDTO
// @Failure 403 :id is not int
// @router /update-branch/:id [put]
func (c *SystemController) UpdateBranch() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			idStr := c.Ctx.Input.Param(":id")
			var r requests.BranchRequestDTO
			json.Unmarshal(c.Ctx.Input.RequestBody, &r)
			message := "Branch updated successfully"
			userDetailsResp := functions.GetUserDetails(&c.Controller, r.BranchManager)
			branchResp := &responses.BranchResp{}

			if userDetailsResp.StatusCode == 200 {
				updateBranch := functions.UpdateBranch(&c.Controller, r, verifyToken.User.UserId, idStr)

				if updateBranch.StatusCode == 200 {
					splitName := strings.Split(userDetailsResp.User.FullName, " | ")
					firstname := ""
					lastname := ""
					if len(splitName) > 1 {
						firstname = splitName[0]
						lastname = splitName[1]
					} else {
						firstname = splitName[0]
					}
					role_name, _ := beego.AppConfig.String("branchManagerRoleName")
					logs.Info("About to get data for role ", role_name)
					role := functions.GetRoleWithRoleName(&c.Controller, role_name)
					logs.Info("Get role response is ", role.Role.RoleId)
					var roleId int64 = 0
					if role.StatusCode == 200 {
						roleId = role.Role.RoleId
					}
					logs.Info("Sending role ", roleId)
					userDetails := requests.UpdateUserRequestDTO{RoleId: roleId, BranchId: updateBranch.Result.BranchId, FirstName: firstname, LastName: lastname, Username: userDetailsResp.User.Username, PhoneNumber: userDetailsResp.User.PhoneNumber, Gender: userDetailsResp.User.Gender, Dob: userDetailsResp.User.Dob.GoString(), Address: userDetailsResp.User.Address}
					userId := strconv.FormatInt(userDetailsResp.User.UserId, 10)
					updateUserResp := functions.UpdateUser(&c.Controller, userId, userDetails)
					if updateUserResp.StatusCode == 200 {
						logs.Info("Update user response is ", updateUserResp.StatusDesc)
						branchIdStr := strconv.FormatInt(updateBranch.Result.BranchId, 10)
						updateBranchResp := functions.UpdateBranchBranchManger(&c.Controller, userId, branchIdStr)
						splitName := strings.Split(updateUserResp.User.FullName, " | ")
						firstname := ""
						lastname := ""
						if len(splitName) > 1 {
							firstname = splitName[0]
							lastname = splitName[1]
						} else {
							firstname = splitName[0]
						}

						branchManager := responses.UserGateway{
							UserId:      updateUserResp.User.UserId,
							FirstName:   firstname,
							LastName:    lastname,
							Username:    updateUserResp.User.Username,
							Email:       updateUserResp.User.Email,
							PhoneNumber: updateUserResp.User.PhoneNumber,
							ImagePath:   updateUserResp.User.ImagePath,
							Customer:    updateUserResp.User.UserDetails,
							// Gender:
							// Dob:
							// Address:
							// IdType:
							// IdNumber:
							// Active:
							IsVerified: updateUserResp.User.IsVerified,
							Role:       updateUserResp.User.Role,
						}

						branchResp = &responses.BranchResp{
							BranchId:    updateBranch.Result.BranchId,
							Branch:      updateBranch.Result.BranchName,
							Description: updateBranch.Result.Description,
							// Country:       &country,
							Location:      updateBranch.Result.Location,
							PhoneNumber:   updateBranch.Result.PhoneNumber,
							DateCreated:   updateBranch.Result.DateCreated,
							BranchManager: &branchManager,
						}

						if updateBranchResp.StatusCode != 200 {
							message = "Branch updated. Failed to update branch's branch manager"
						}
					} else {
						message = "Branch updated. Failed to update branch manager"
						logs.Error("Failed to update user", updateBranch.StatusDesc)
						resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Branch update failed. " + updateUserResp.StatusDesc}
						c.Data["json"] = resp
					}

					isSuccess = true
				} else {
					message = "Failed to update branch"
					branchResp = nil
				}
			} else {
				logs.Info("Failed to get branch manager")
				message = "Failed to get specified branch manager"
			}
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: branchResp, StatusDesc: message}
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = resp
		} else {
			logs.Error("Failed to verify token")
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to verify token"}
			c.Data["json"] = resp
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Delete Branch ...
// @Title Delete Branch
// @Description delete the Branches
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete-branch/:id [delete]
func (c *SystemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			message := "Deleted"
			deleteResp := functions.DeleteBranch(&c.Controller, idStr)
			if deleteResp.StatusCode == 200 {
				resp := responses.StringResponseDTO{Success: true, Result: &message, StatusDesc: "Branch deleted successfully"}
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = resp
			} else {
				message = "Deletion Failed:: "
				resp := responses.StringResponseDTO{Success: false, Result: &message, StatusDesc: deleteResp.StatusDesc}
				c.Ctx.Output.SetStatus(301)
				c.Data["json"] = resp
			}
		} else {
			c.Ctx.Output.SetStatus(200)
			logs.Info("Failed to verify token")
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to verify token"}
			c.Data["json"] = resp
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Get All ID Types ...
// @Title Get All ID Types
// @Description get all ID types
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.IDTypesGatewayResponseDTO
// @Failure 403
// @router /get-id-types [get]
func (c *SystemController) GetIdTypes() {
	// v := c.Ctx.Input.GetData("user")
	// userData, err := v.(*responses.UsersOri)

	// fmt.Printf("Type of v: %T\n", v)
	// fmt.Printf("Value of v: %+v\n", v)

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

	var isSuccess bool = false
	var message string = "An error occurred"

	idTypes := functions.GetIdTypes(&c.Controller, query, fields, sortby, order, offset, limit)

	// json.Unmarshal(getBranchResp.Branches, &v)

	if idTypes.StatusCode == 200 {
		logs.Info("Response is 200")

		isSuccess = true
		message = "ID Types fetched successfully"
		var resp responses.IDTypesGatewayResponseDTO = responses.IDTypesGatewayResponseDTO{Success: isSuccess, Result: idTypes.IdTypes, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred fetching id types from api ", idTypes.StatusDesc)
		message = "Id type fetch error"
		var resp responses.IDTypesGatewayResponseDTO = responses.IDTypesGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// AddApplication ...
// @Title Create Application
// @Description create an application
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.ApplicationRequest	true		"body for Application content"
// @Success 200 {object} responses.ApplicationResponseDTO
// @Failure 403 body is empty
// @router /add-application [post]
func (c *SystemController) AddApplication() {
	userData, _ := c.Ctx.Input.GetData("user").(*responses.UsersOri)
	var v requests.ApplicationRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var isSuccess bool = false

	logs.Info("Token verified")
	isSuccess = false
	addedBy := userData.UserId
	appResp := functions.AddApplication(&c.Controller, v, addedBy)
	if appResp.Success {
		isSuccess = true
		message := "Application added successfully"
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: appResp.Result, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred while adding the application"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetApplication ...
// @Title Get Application
// @Description get an application by code
// @Param	code		path 	string	true		"The application code"
// @Success 200 {object} responses.ApplicationResponseDTO
// @Failure 403 code is empty
// @router /get-application/:code [get]
func (c *SystemController) GetApplication() {
	code := c.Ctx.Input.Param(":code")

	logs.Info("Getting application with code: ", code)

	isSuccess := false

	appResp := functions.GetApplicationByCode(&c.Controller, code)
	if appResp.Success {
		isSuccess = true
		message := "Application fetched successfully"
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: appResp.Result, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred while fetching the application"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetApplications ...
// @Title Get Applications
// @Description get all applications
// @Success 200 {object} responses.ApplicationsResponseDTO
// @Failure 403 an error occurred
// @router /get-applications [get]
func (c *SystemController) GetApplications() {
	isSuccess := false

	appResp := functions.GetAllApplications(&c.Controller)
	if appResp.Success {
		isSuccess = true
		message := "Applications fetched successfully"
		var resp responses.ApplicationsResponseDTO = responses.ApplicationsResponseDTO{Success: isSuccess, Result: appResp.Result, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		var resp responses.ApplicationsResponseDTO = responses.ApplicationsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred while fetching the applications"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateApplication ...
// @Title Update Application
// @Description update an application
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ApplicationRequest	true		"body for Application content"
// @Success 200 {object} responses.ApplicationResponseDTO
// @Failure 403 :id is not int
// @router /update-application/:id [put]
func (c *SystemController) UpdateApplication() {
	userData, _ := c.Ctx.Input.GetData("user").(*responses.UsersOri)
	idStr := c.Ctx.Input.Param(":id")
	var v requests.ApplicationRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var isSuccess bool = false

	logs.Info("Token verified")
	isSuccess = true
	updateRequest := requests.UpdateApplicationRequest{
		ApplicationCode:  v.ApplicationCode,
		ApplicationName:  v.ApplicationName,
		ApplicationLogo:  v.ApplicationLogo,
		ThemeColors:      v.ThemeColors,
		DefaultFontsize:  v.DefaultFontsize,
		ApplicationImage: v.ApplicationImage,
		ThemeCode:        v.ThemeCode,
		UpdatedBy:        userData.UserId,
	}
	appResp := functions.UpdateApplication(&c.Controller, updateRequest, idStr)
	if appResp.Success {
		isSuccess = true
		message := "Application updated successfully"
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: appResp.Result, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		message := "An error occurred while updating the application"
		var resp responses.ApplicationResponseDTO = responses.ApplicationResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// AddTheme ...
// @Title Create Theme
// @Description create a theme
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.ThemeRequest	true		"body for Theme content"
// @Success 200 {object} responses.ThemeResponseDTO
// @Failure 403 body is empty
// @router /add-theme [post]
func (c *SystemController) AddTheme() {
	var v requests.ThemeRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			logs.Info("Token verified")
			isSuccess = true
			themeResp := functions.AddTheme(&c.Controller, v)
			c.Data["json"] = themeResp
		} else {
			var resp responses.ThemeResponseDTO = responses.ThemeResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.ThemeResponseDTO = responses.ThemeResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateTheme ...
// @Title Update Theme
// @Description update a theme
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ThemeRequest	true		"body for Theme content"
// @Success 200 {object} responses.ThemeResponseDTO
// @Failure 403 :id is not int
// @router /update-theme/:id [put]
func (c *SystemController) UpdateTheme() {
	idStr := c.Ctx.Input.Param(":id")
	var v requests.ThemeRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			logs.Info("Token verified")
			isSuccess = true
			themeResp := functions.UpdateTheme(&c.Controller, v, idStr)
			c.Data["json"] = themeResp
		} else {
			var resp responses.ThemeResponseDTO = responses.ThemeResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.ThemeResponseDTO = responses.ThemeResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateImage ...
// @Title UpdateImage
// @Description Update User's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	Image		formData 	file	true		"System Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /upload-system-image [post]
func (c *SystemController) UploadSystemImage() {

	var isSuccess bool = false

	image, header, err := c.GetFile("Image")

	if err != nil {
		var resp responses.SystemImageResponseDTO = responses.SystemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		system := c.Ctx.Input.Query("System")
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "Image", image, *header)

		if respCode == 200 {
			itemImage := functions.UploadSystemImage(&c.Controller, filePath, system)

			if itemImage.StatusCode == 200 {
				logs.Info("Item image returned: ", itemImage.Value)

				isSuccess = true

				var resp responses.SystemImageResponseDTO = responses.SystemImageResponseDTO{Success: isSuccess, Result: itemImage.Value, StatusDesc: itemImage.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.SystemImageResponseDTO = responses.SystemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.SystemImageResponseDTO = responses.SystemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to upload file. Tmp"}
			c.Data["json"] = resp
		}

	}

	c.ServeJSON()
}
