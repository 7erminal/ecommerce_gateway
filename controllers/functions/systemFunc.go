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
