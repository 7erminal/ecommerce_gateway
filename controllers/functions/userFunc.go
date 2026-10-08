package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func GetUserDetails(c *beego.Controller, userid string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user details ", userid)

	request := api.NewRequest(
		host,
		"/v1/users/"+userid,
		api.GET)
	// request.Params["username"] = username
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

	// var prettyJSON bytes.Buffer
	// if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
	// 	logs.Info("Raw response received is ", string(read))
	// } else {
	// 	logs.Info("Raw response received is \n", prettyJSON.String())
	// }
	logs.Info("Raw response received is ", string(read))
	// data := map[string]interface{}{}
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func InviteUser(c *beego.Controller, email string, role int64, link string, inviteby string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", email)

	request := api.NewRequest(
		host,
		"/v1/users/invite-user/",
		api.POST)
	request.InterfaceParams["Email"] = email
	request.InterfaceParams["Link"] = link
	request.InterfaceParams["Role"] = strconv.FormatInt(role, 10)
	request.InterfaceParams["InviteBy"] = inviteby
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetUsers(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetUsersWithRole(c *beego.Controller, role_id string, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users/role/"+role_id,
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetUsersWithBranch(c *beego.Controller, branch_id string, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users/branch/"+branch_id,
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetRoles(c *beego.Controller) (resp responses.RolesAllResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/roles",
		api.GET)
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
	var data responses.RolesAllResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetInvites(c *beego.Controller) (resp responses.UserInvitesResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/users/get-user-invites",
		api.GET)
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
	var data responses.UserInvitesResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetRole(c *beego.Controller, role string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/roles/"+role,
		api.GET)
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetRoleWithRoleName(c *beego.Controller, role string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting role with ", role)

	request := api.NewRequest(
		host,
		"/v1/roles/role/"+role,
		api.GET)
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func AddRole(c *beego.Controller, req requests.AddRoleRequestDTO, createdBy string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Adding role with name ", req.Name)

	request := api.NewRequest(
		host,
		"/v1/roles",
		api.POST)
	request.InterfaceParams["Role"] = req.Name
	request.InterfaceParams["Description"] = req.Description
	request.InterfaceParams["AddedBy"] = createdBy
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func DeleteRole(c *beego.Controller, role string, deletedBy string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Deleting role with name ", role)

	request := api.NewRequest(
		host,
		"/v1/roles/"+role,
		api.DELETE)
	request.Params["DeletedBy"] = deletedBy
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetPermissions(c *beego.Controller) (resp responses.PermissionsResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	request := api.NewRequest(
		host,
		"/v1//permissions",
		api.GET)
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
	var data responses.PermissionsResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetActions(c *beego.Controller) (resp responses.ActionsResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting actions")

	request := api.NewRequest(
		host,
		"/v1/actions",
		api.GET)
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
	var data responses.ActionsResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func AddRolePermission(c *beego.Controller, role string, req requests.UpdateRolePermissionRequestDTO, addedBy string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Updating role permissions for role ", role)

	request := api.NewRequest(
		host,
		"/v1/role-permissions",
		api.POST)
	request.InterfaceParams["Role"] = req.Role
	request.InterfaceParams["PermissionCode"] = req.PermissionCode
	request.InterfaceParams["Action"] = req.Action
	request.InterfaceParams["AddedBy"] = addedBy

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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func RemoveRolePermission(c *beego.Controller, role string, permission string, action string, removedBy string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Removing role permissions for role ", role)

	request := api.NewRequest(
		host,
		"/v1/role-permissions/"+role,
		api.DELETE)
	request.Params["Role"] = role
	request.Params["PermissionCode"] = permission
	request.Params["Action"] = action
	request.Params["RemovedBy"] = removedBy

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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetInvite(c *beego.Controller, token string) (resp responses.UserInviteResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user token ", token)

	request := api.NewRequest(
		host,
		"/v1/users/user-invite/"+token,
		api.GET)
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
	var data responses.UserInviteResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateUser(c *beego.Controller, id string, req requests.UpdateUserRequestDTO, updatedBy string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.FirstName)

	request := api.NewRequest(
		host,
		"/v1/users/"+id,
		api.PUT)
	request.InterfaceParams["FullName"] = req.FirstName + " | " + req.LastName
	request.InterfaceParams["Address"] = req.Address
	request.InterfaceParams["Dob"] = req.Dob
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Username"] = req.Username
	request.InterfaceParams["MaritalStatus"] = ""
	request.InterfaceParams["BranchId"] = req.BranchId
	request.InterfaceParams["ImagePath"] = ""
	request.InterfaceParams["RoleId"] = req.RoleId
	request.InterfaceParams["UpdatedBy"] = updatedBy
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
	// var dataOri responses.UserOriResponseDTO
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func UpdateUserRole(c *beego.Controller, id string, req requests.UpdateUserRoleRequestDTO, updatedBy string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.RoleId)

	request := api.NewRequest(
		host,
		"/v1/users/role/"+id,
		api.PUT)
	request.InterfaceParams["RoleId"] = req.RoleId
	request.InterfaceParams["UpdatedBy"] = updatedBy
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
	// var dataOri responses.UserOriResponseDTO
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func UpdateUserBranch(c *beego.Controller, id string, req requests.UpdateUserBranchRequestDTO, updatedBy string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending branch ID ", req.BranchId)

	request := api.NewRequest(
		host,
		"/v1/users/branch/"+id,
		api.PUT)
	request.InterfaceParams["BranchId"] = req.BranchId
	request.InterfaceParams["UpdatedBy"] = updatedBy
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
	// var dataOri responses.UserOriResponseDTO
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func UpdateUserImage(c *beego.Controller, userImage string, userId string, updatedBy string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending file ", userImage)

	request := api.NewRequest(
		host,
		"/v1/users/update-user-image/",
		api.POST)

	request.FileField["UserImage"] = userImage
	request.Params["UserId"] = userId
	request.Params["UpdatedBy"] = updatedBy
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateUserInvite(c *beego.Controller, id string, status string) (resp responses.UserInviteResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Updating invite ", status)

	request := api.NewRequest(
		host,
		"/v1/users/update-user-invite/"+id,
		api.PUT)
	request.InterfaceParams["Status"] = status

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
	// var dataOri responses.UserOriResponseDTO
	var data responses.UserInviteResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data)

	return data
}

func LogOut(c *beego.Controller, userToken string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("token", userToken)

	request := api.NewRequest(
		host,
		"/v1/auth/log-out/",
		api.POST)

	request.InterfaceParams["Token"] = userToken
	// request.HeaderField["content-type"] = "multipart/form-data"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
