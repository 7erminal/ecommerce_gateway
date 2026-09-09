package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func AddBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.POST)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64, branchId string) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchId,
		api.PUT)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetBranch(c *beego.Controller, branchid string) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchid,
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func DeleteBranch(c *beego.Controller, branchid string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchid,
		api.DELETE)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetSystemDetails(c *beego.Controller, branchid string) (resp responses.SystemDetailsResponseDTO, err error) {
	message := "An error occurred"
	proceed := true
	response := responses.SystemDetailsResponseDTO{}

	getBranchResp := GetBranch(c, branchid)

	// if getBranchResp.StatusCode == 200 {
	// userId := verifyToken.User.UserId
	if getBranchResp.StatusCode != 200 {
		err = fmt.Errorf("Branch provided does not exist")
		message = "Branch provided does not exist"
		proceed = false
	}

	if proceed {
		getCountryResp := GetCountry(c, strconv.FormatInt(getBranchResp.Result.Country.CountryId, 10))

		if getCountryResp.StatusCode != 200 {
			err = fmt.Errorf("Country provided does not exist")
			message = "Country provided does not exist"
			proceed = false
		}

		if proceed {
			// getCurrencyResp := GetCurrency(c, strconv.FormatInt(getCountryResp.Result.DefaultCurrency, 10))

			// if getCurrencyResp.StatusCode != 200 {
			// 	err = fmt.Errorf("Currency provided does not exist")
			// 	message = "Currency provided does not exist"
			// 	proceed = false
			// } else {}

			if proceed {
				currencyResp_ := responses.CurrencyResp{
					CurrencyId: strconv.FormatInt(getBranchResp.Result.Country.Currency.CurrencyId, 10),
					Currency:   getBranchResp.Result.Country.Currency.Currency,
					Symbol:     getBranchResp.Result.Country.Currency.Symbol,
				}

				countryResp_ := responses.CountryResp{
					Country:     getCountryResp.Result.Country,
					CountryCode: getCountryResp.Result.CountryCode,
					Currency:    &currencyResp_,
				}
				branchResp_ := responses.BranchResp{
					BranchId:      getBranchResp.Result.BranchId,
					Branch:        getBranchResp.Result.BranchName,
					Description:   getBranchResp.Result.Description,
					Location:      getBranchResp.Result.Location,
					PhoneNumber:   getBranchResp.Result.PhoneNumber,
					BranchManager: nil,
					DateCreated:   getBranchResp.Result.DateCreated,
					Country:       &countryResp_,
				}

				message = "System details retrieved successfully"

				response = responses.SystemDetailsResponseDTO{
					Result: &responses.SystemDetailsData{
						Branch: &branchResp_,
					},
					Success:    proceed,
					StatusDesc: message,
				}
				err = nil
				resp = response
			}
		}
	}

	return resp, err
}

func GetBranches(c *beego.Controller) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranchBranchManger(c *beego.Controller, userid string, branchid string) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/branch-manager/"+branchid,
		api.PUT)
	request.InterfaceParams["BranchManager"] = userid
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountries(c *beego.Controller) (resp responses.CountriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/",
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CountriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountry(c *beego.Controller, id string) (resp responses.CountryOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/"+id,
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CountryOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountryByCode(c *beego.Controller, code string) (resp responses.CountryOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/code/"+code,
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CountryOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCurrency(c *beego.Controller, id string) (resp responses.CurrencyOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/currencies/"+id,
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CurrencyOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCurrencyByCode(c *beego.Controller, code string) (resp responses.CountriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/currencies/code/"+code,
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CountriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func AddApplication(c *beego.Controller, req requests.ApplicationRequest, addedBy int64) (resp responses.ApplicationResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Adding application with code: ", req.ApplicationCode)

	request := api.NewRequest(
		host,
		"/v1/applications/",
		api.POST)
	request.InterfaceParams["ApplicationCode"] = req.ApplicationCode
	request.InterfaceParams["ApplicationName"] = req.ApplicationName
	request.InterfaceParams["ApplicationLogo"] = req.ApplicationLogo
	request.InterfaceParams["ThemeColors"] = req.ThemeColors
	request.InterfaceParams["DefaultFontsize"] = req.DefaultFontsize
	request.InterfaceParams["ApplicationImage"] = req.ApplicationImage
	request.InterfaceParams["ThemeCode"] = req.ThemeCode

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ApplicationResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 201 || backendResp.StatusCode == 200 {
		resp = responses.ApplicationResponseDTO{
			Success: true,
			Result: &responses.ApplicationResp{
				ApplicationId:    backendResp.Result.ApplicationId,
				ApplicationCode:  backendResp.Result.ApplicationCode,
				ApplicationName:  backendResp.Result.ApplicationName,
				ApplicationLogo:  backendResp.Result.ApplicationLogo,
				ThemeColors:      backendResp.Result.ThemeColors,
				DefaultFontsize:  backendResp.Result.DefaultFontsize,
				ApplicationImage: backendResp.Result.ApplicationImage,
				DateCreated:      backendResp.Result.DateCreated,
				DateModified:     backendResp.Result.DateModified,
				Active:           backendResp.Result.Active,
				Theme: &responses.ThemeResp{
					ThemeId:   backendResp.Result.Theme.ThemeId,
					ThemeCode: backendResp.Result.Theme.ThemeCode,
					ThemeName: backendResp.Result.Theme.ThemeName,
				},
			},
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ApplicationResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: backendResp.StatusMessage,
		}
	}

	logs.Info("Resp is ", resp)
	return resp
}

func UpdateApplication(c *beego.Controller, req requests.UpdateApplicationRequest, applicationId string) (resp responses.ApplicationResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Updating application with id: ", applicationId)

	request := api.NewRequest(
		host,
		"/v1/applications/"+applicationId,
		api.PUT)
	request.InterfaceParams["ApplicationCode"] = req.ApplicationCode
	request.InterfaceParams["ApplicationName"] = req.ApplicationName
	request.InterfaceParams["ApplicationLogo"] = req.ApplicationLogo
	request.InterfaceParams["ThemeColors"] = req.ThemeColors
	request.InterfaceParams["DefaultFontsize"] = req.DefaultFontsize
	request.InterfaceParams["ApplicationImage"] = req.ApplicationImage
	request.InterfaceParams["ThemeCode"] = req.ThemeCode
	request.InterfaceParams["UpdatedBy"] = req.UpdatedBy

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ApplicationResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 200 {
		resp = responses.ApplicationResponseDTO{
			Success: true,
			Result: &responses.ApplicationResp{
				ApplicationId:    backendResp.Result.ApplicationId,
				ApplicationCode:  backendResp.Result.ApplicationCode,
				ApplicationName:  backendResp.Result.ApplicationName,
				ApplicationLogo:  backendResp.Result.ApplicationLogo,
				ThemeColors:      backendResp.Result.ThemeColors,
				DefaultFontsize:  backendResp.Result.DefaultFontsize,
				ApplicationImage: backendResp.Result.ApplicationImage,
				DateCreated:      backendResp.Result.DateCreated,
				DateModified:     backendResp.Result.DateModified,
				Active:           backendResp.Result.Active,
				Theme: &responses.ThemeResp{
					ThemeId:   backendResp.Result.Theme.ThemeId,
					ThemeCode: backendResp.Result.Theme.ThemeCode,
					ThemeName: backendResp.Result.Theme.ThemeName,
				},
			},
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ApplicationResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: backendResp.StatusMessage,
		}
	}

	logs.Info("Resp is ", resp)
	return resp
}

func AddTheme(c *beego.Controller, req requests.ThemeRequest) (resp responses.ThemeResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Adding theme with code: ", req.ThemeCode)

	request := api.NewRequest(
		host,
		"/v1/themes/",
		api.POST)
	request.InterfaceParams["ThemeCode"] = req.ThemeCode
	request.InterfaceParams["ThemeName"] = req.ThemeName

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ThemeResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 201 || backendResp.StatusCode == 200 {
		resp = responses.ThemeResponseDTO{
			Success: true,
			Result: &responses.ThemeResp{
				ThemeId:   backendResp.Result.ThemeId,
				ThemeCode: backendResp.Result.ThemeCode,
				ThemeName: backendResp.Result.ThemeName,
			},
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ThemeResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: backendResp.StatusMessage,
		}
	}

	logs.Info("Resp is ", resp)
	return resp
}

func UpdateTheme(c *beego.Controller, req requests.ThemeRequest, themeId string) (resp responses.ThemeResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Updating theme with id: ", themeId)

	request := api.NewRequest(
		host,
		"/v1/themes/"+themeId,
		api.PUT)
	request.InterfaceParams["ThemeCode"] = req.ThemeCode
	request.InterfaceParams["ThemeName"] = req.ThemeName

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ThemeResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 200 {
		resp = responses.ThemeResponseDTO{
			Success: true,
			Result: &responses.ThemeResp{
				ThemeId:   backendResp.Result.ThemeId,
				ThemeCode: backendResp.Result.ThemeCode,
				ThemeName: backendResp.Result.ThemeName,
			},
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ThemeResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: backendResp.StatusMessage,
		}
	}

	logs.Info("Resp is ", resp)
	return resp
}
func GetAllApplications(c *beego.Controller) (resp responses.ApplicationsResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting all applications")

	request := api.NewRequest(
		host,
		"/v1/applications/",
		api.GET)

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ApplicationsResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 200 {
		resp = responses.ApplicationsResponseDTO{
			Success:    true,
			Result:     backendResp.Result,
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ApplicationsResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: "An error occurred while fetching the applications",
		}
	}
	logs.Info("Resp is ", resp)
	return resp
}
func GetApplicationByCode(c *beego.Controller, code string) (resp responses.ApplicationResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting application by code: ", code)

	request := api.NewRequest(
		host,
		"/v1/applications/"+code,
		api.GET)

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	var backendResp responses.ApplicationResponse
	json.Unmarshal(read, &backendResp)

	// Transform backend response to gateway response
	if backendResp.StatusCode == 200 {
		resp = responses.ApplicationResponseDTO{
			Success: true,
			Result: &responses.ApplicationResp{
				ApplicationId:    backendResp.Result.ApplicationId,
				ApplicationCode:  backendResp.Result.ApplicationCode,
				ApplicationName:  backendResp.Result.ApplicationName,
				ApplicationLogo:  backendResp.Result.ApplicationLogo,
				ThemeColors:      backendResp.Result.ThemeColors,
				DefaultFontsize:  backendResp.Result.DefaultFontsize,
				ApplicationImage: backendResp.Result.ApplicationImage,
				DateCreated:      backendResp.Result.DateCreated,
				DateModified:     backendResp.Result.DateModified,
				Active:           backendResp.Result.Active,
				Theme: &responses.ThemeResp{
					ThemeId:   backendResp.Result.Theme.ThemeId,
					ThemeCode: backendResp.Result.Theme.ThemeCode,
					ThemeName: backendResp.Result.Theme.ThemeName,
				},
			},
			StatusDesc: backendResp.StatusMessage,
		}
	} else {
		resp = responses.ApplicationResponseDTO{
			Success:    false,
			Result:     nil,
			StatusDesc: backendResp.StatusMessage,
		}
	}

	logs.Info("Resp is ", resp)
	return resp
}

func UploadSystemImage(c *beego.Controller, systemImage string, system string) (resp responses.SystemImageOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending file ", systemImage)

	request := api.NewRequest(
		host,
		"/v1/application/upload-pictures",
		api.POST)

	request.FileField["Image"] = systemImage
	request.InterfaceParams["System"] = system
	// request.HeaderField["content-type"] = "multipart/form-data"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}

	// client.Request.HeaderField["content-type"] = "multipart/form-data"
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.SystemImageOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
