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

func SuppliersaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		var requestBody models.Supplier
		
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
		url := fmt.Sprintf("%s/accounting/suppliers%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.CreateSupplierResponse
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

func CreateSuppliersaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_accounting_suppliers",
		mcp.WithDescription("Create Supplier"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("last_name", mcp.Description("Input parameter: The last name of the person.")),
		mcp.WithString("suffix", mcp.Description("")),
		mcp.WithArray("bank_accounts", mcp.Description("")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithArray("addresses", mcp.Description("")),
		mcp.WithObject("account", mcp.Description("")),
		mcp.WithString("downstream_id", mcp.Description("Input parameter: The third-party API ID of original entity")),
		mcp.WithBoolean("individual", mcp.Description("Input parameter: Is this an individual or business supplier")),
		mcp.WithString("company_name", mcp.Description("Input parameter: The name of the company.")),
		mcp.WithArray("websites", mcp.Description("")),
		mcp.WithArray("emails", mcp.Description("")),
		mcp.WithArray("phone_numbers", mcp.Description("")),
		mcp.WithString("notes", mcp.Description("Input parameter: Some notes about this supplier")),
		mcp.WithString("middle_name", mcp.Description("Input parameter: Middle name of the person.")),
		mcp.WithString("display_name", mcp.Description("Input parameter: Display name")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: A unique identifier for an object.")),
		mcp.WithString("payment_method", mcp.Description("Input parameter: Payment method used for the transaction, such as cash, credit card, bank transfer, or check")),
		mcp.WithString("first_name", mcp.Description("Input parameter: The first name of the person.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithObject("tax_rate", mcp.Description("")),
		mcp.WithString("channel", mcp.Description("Input parameter: The channel through which the transaction is processed.")),
		mcp.WithString("title", mcp.Description("Input parameter: The job title of the person.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithString("display_id", mcp.Description("Input parameter: Display ID")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithString("status", mcp.Description("Input parameter: Supplier status")),
		mcp.WithString("tax_number", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SuppliersaddHandler(cfg),
	}
}
