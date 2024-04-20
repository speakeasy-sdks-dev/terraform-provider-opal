// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/opalsecurity/terraform-provider-opal/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GroupUserResource{}
var _ resource.ResourceWithImportState = &GroupUserResource{}

func NewGroupUserResource() resource.Resource {
	return &GroupUserResource{}
}

// GroupUserResource defines the resource implementation.
type GroupUserResource struct {
	client *sdk.OpalAPI
}

// GroupUserResourceModel describes the resource data model.
type GroupUserResourceModel struct {
	AccessLevel         *tfTypes.ResourceAccessLevel `tfsdk:"access_level"`
	AccessLevelRemoteID types.String                 `tfsdk:"access_level_remote_id"`
	DurationMinutes     types.Int64                  `tfsdk:"duration_minutes"`
	Email               types.String                 `tfsdk:"email"`
	ExpirationDate      types.String                 `tfsdk:"expiration_date"`
	FullName            types.String                 `tfsdk:"full_name"`
	GroupID             types.String                 `tfsdk:"group_id"`
	UserID              types.String                 `tfsdk:"user_id"`
}

func (r *GroupUserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_user"
}

func (r *GroupUserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GroupUser Resource",
		Attributes: map[string]schema.Attribute{
			"access_level": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_level_name": schema.StringAttribute{
						Computed:    true,
						Description: `The human-readable name of the access level.`,
					},
					"access_level_remote_id": schema.StringAttribute{
						Computed:    true,
						Description: `The machine-readable identifier of the access level.`,
					},
				},
				MarkdownDescription: `# Access Level Object` + "\n" +
					`### Description` + "\n" +
					`The ` + "`" + `GroupAccessLevel` + "`" + ` object is used to represent the level of access that a user has to a group or a group has to a group. The "default" access` + "\n" +
					`level is a ` + "`" + `GroupAccessLevel` + "`" + ` object whose fields are all empty strings.` + "\n" +
					`` + "\n" +
					`### Usage Example` + "\n" +
					`View the ` + "`" + `GroupAccessLevel` + "`" + ` of a group/user or group/group pair to see the level of access granted to the group.`,
			},
			"access_level_remote_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. Requires replacement if changed. `,
			},
			"duration_minutes": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(0),
				Description: `Must be set to 0. Any nonzerovalue in terraform does not make sense. Requires replacement if changed. ; must be one of ["0"]; Default: 0`,
				Validators: []validator.Int64{
					int64validator.OneOf(
						[]int64{
							0,
						}...,
					),
				},
			},
			"email": schema.StringAttribute{
				Computed:    true,
				Description: `The user's email.`,
			},
			"expiration_date": schema.StringAttribute{
				Computed:    true,
				Description: `The day and time the user's access will expire.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"full_name": schema.StringAttribute{
				Computed:    true,
				Description: `The user's full name.`,
			},
			"group_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `The ID of the group. Requires replacement if changed. `,
			},
			"user_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `The ID of the user to add. Requires replacement if changed. `,
			},
		},
	}
}

func (r *GroupUserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.OpalAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.OpalAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GroupUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GroupUserResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := *data.ToOperationsCreateGroupUserRequestBody()
	groupID := data.GroupID.ValueString()
	userID := data.UserID.ValueString()
	request := operations.CreateGroupUserRequest{
		RequestBody: requestBody,
		GroupID:     groupID,
		UserID:      userID,
	}
	res, err := r.client.Groups.CreateUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GroupUser == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGroupUser(res.GroupUser)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GroupUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GroupUserResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GroupUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GroupUserResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GroupUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GroupUserResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	groupID := data.GroupID.ValueString()
	userID := data.UserID.ValueString()
	request := operations.DeleteGroupUserRequest{
		GroupID: groupID,
		UserID:  userID,
	}
	res, err := r.client.Groups.DeleteUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GroupUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource group_user.")
}
