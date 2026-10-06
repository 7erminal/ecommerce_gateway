package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/responses"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
)

// StatsController operations for Stats
type StatsController struct {
	BaseController
}

// URLMapping ...
func (c *StatsController) URLMapping() {
	c.Mapping("GetGeneralStats", c.GetGeneralStats)
}

// GetGeneralStats ...
// @Title Get Stats Per Branch
// @Description Get Stats of Branch
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.ItemsResponseDTO
// @Failure 403 body is empty
// @router /get-stats [get]
func (c *StatsController) GetGeneralStats() {
	if !c.RequirePermission("REPORT", "READ") {
		c.Data["json"] = map[string]string{"error": "forbidden"}
		c.ServeJSON()
		return
	}
	v := c.Ctx.Input.GetData("user")
	userData, err := v.(*responses.UsersOri)
	if err {
		logs.Error("Unable to get user data")
	}

	var isSuccess bool = false

	logs.Info("Success response received")
	branchidStr := strconv.FormatInt(userData.UserDetails.Branch.BranchId, 10)

	getItemStatsResp := functions.GetItemStats(&c.Controller, branchidStr)

	if getItemStatsResp.StatusCode == 200 {
		logs.Info("Item stats returned: ")
		isSuccess = true

		itemStats := responses.StatsDTO{}
		if getItemStatsResp.Stats != nil {

			itemStats = *getItemStatsResp.Stats
		}

		isSuccess = true

		var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: &itemStats, StatusDesc: getItemStatsResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
