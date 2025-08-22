package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/accounting-api/mcp-server/config"
	"github.com/accounting-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func TaxratesupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["raw"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("raw=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.TaxRate
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/accounting/tax-rates/%s%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-apideck-consumer-id"]; ok {
			req.Header.Set("x-apideck-consumer-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-app-id"]; ok {
			req.Header.Set("x-apideck-app-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-service-id"]; ok {
			req.Header.Set("x-apideck-service-id", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UpdateTaxRateResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateTaxratesupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_accounting_tax-rates_id",
		mcp.WithDescription("Update Tax Rate"),
		mcp.WithString("id", mcp.Required(), mcp.Description("ID of the record you are acting upon.")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("type", mcp.Description("Input parameter: Tax type used to indicate the source of tax collected or paid")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithArray("components", mcp.Description("")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("effective_tax_rate", mcp.Description("Input parameter: Effective tax rate")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithString("description", mcp.Description("Input parameter: Description of tax rate")),
		mcp.WithString("tax_remitted_account_id", mcp.Description("Input parameter: Unique identifier for the account for tax remitted.")),
		mcp.WithString("total_tax_rate", mcp.Description("Input parameter: Not compounded sum of the components of a tax rate")),
		mcp.WithString("original_tax_rate_id", mcp.Description("Input parameter: ID of the original tax rate from which the new tax rate is derived. Helps to understand the relationship between corresponding tax rate entities.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("name", mcp.Description("Input parameter: Name assigned to identify this tax rate.")),
		mcp.WithString("report_tax_type", mcp.Description("Input parameter: Report Tax type to aggregate tax collected or paid for reporting purposes")),
		mcp.WithString("id", mcp.Description("Input parameter: ID assigned to identify this tax rate.")),
		mcp.WithString("tax_payable_account_id", mcp.Description("Input parameter: Unique identifier for the account for tax collected.")),
		mcp.WithString("code", mcp.Description("Input parameter: Tax code assigned to identify this tax rate.")),
		mcp.WithString("status", mcp.Description("Input parameter: Tax rate status")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TaxratesupdateHandler(cfg),
	}
}
