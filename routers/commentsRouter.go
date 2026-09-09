package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "ChangePassword",
            Router: `/change-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "RefreshAccessToken",
            Router: `/refresh-access-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "RefreshCustomerAccessToken",
            Router: `/refresh-customer-access-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "SignIn",
            Router: `/sign-in`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "VerifyToken",
            Router: `/verify-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "UpdateCustomerImage",
            Router: `/upload-customer-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: `/add-category`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddFeature",
            Router: `/add-feature`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddPurpose",
            Router: `/add-purpose`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddRentalsItem",
            Router: `/add-rental-product`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddSalesItem",
            Router: `/add-sales-product`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "DeleteCategory",
            Router: `/delete-category/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "DeleteFeature",
            Router: `/delete-feature/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "DeleteItem",
            Router: `/delete-item/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "DeletePurpose",
            Router: `/delete-purpose/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetCategories",
            Router: `/get-categories`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetFeatures",
            Router: `/get-features`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: `/get-item/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItems",
            Router: `/get-items`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetPurposes",
            Router: `/get-purposes`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "UpdateItem",
            Router: `/update-product/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "UpdateItemImage",
            Router: `/upload-product-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "GetPaymentMethods",
            Router: `/payment-methods`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "UploadPaymentProof",
            Router: `/upload-payment-proof`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:RegistrationController"],
        beego.ControllerComments{
            Method: "RegisterAlt",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:StatsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:StatsController"],
        beego.ControllerComments{
            Method: "GetGeneralStats",
            Router: `/get-stats`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "AddApplication",
            Router: `/add-application`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "AddBranch",
            Router: `/add-branch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "AddTheme",
            Router: `/add-theme`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete-branch/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetApplication",
            Router: `/get-application/:code`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetApplications",
            Router: `/get-applications`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetOneBranch",
            Router: `/get-branch/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetAllBranches",
            Router: `/get-branches`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetAllCountries",
            Router: `/get-countries`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetIdTypes",
            Router: `/get-id-types`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetRoles",
            Router: `/get-roles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetSystemDetails",
            Router: `/get-system-details/:branchid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "UpdateApplication",
            Router: `/update-application/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "UpdateBranch",
            Router: `/update-branch/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "UpdateTheme",
            Router: `/update-theme/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "UploadSystemImage",
            Router: `/upload-system-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "GetAllTransactions",
            Router: `/get-all-transactions`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "GetOneOrder",
            Router: `/get-order/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "GetAllOrders",
            Router: `/get-orders`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "PlaceOrderRequest",
            Router: `/place-order-request`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "PlaceRentalRequest",
            Router: `/place-rental-request`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "PlaceSalesRequest",
            Router: `/place-sales-request`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetBranchManagers",
            Router: `/get-branch-managers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserInvites",
            Router: `/get-invites`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetRoles",
            Router: `/get-roles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserInvite",
            Router: `/get-user-invite/:token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/get-user-session`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserWithId",
            Router: `/get-user-with-id/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: `/get-users`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUsersUnderBranch",
            Router: `/get-users-under-branch/:branch_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "InviteUserReg",
            Router: `/invite-user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "LogOut",
            Router: `/log-out`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateInviteToken",
            Router: `/revoke-invite/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserBranch",
            Router: `/update-user-branch/:userid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserImage",
            Router: `/update-user-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserRole",
            Router: `/update-user-role/:userid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update-user/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "VerifyInvite",
            Router: `/verify-invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
