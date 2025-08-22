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

func InvoiceitemsaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.InvoiceItem
		
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
		url := fmt.Sprintf("%s/accounting/invoice-items%s", cfg.BaseURL, queryString)
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
		var result models.CreateInvoiceItemResponse
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

func CreateInvoiceitemsaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_accounting_invoice-items",
		mcp.WithDescription("Create Invoice Item"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithObject("asset_account", mcp.Description("")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithObject("tracking_category", mcp.Description("")),
		mcp.WithBoolean("taxable", mcp.Description("Input parameter: If true, transactions for this item are taxable")),
		mcp.WithString("code", mcp.Description("Input parameter: User defined item code")),
		mcp.WithString("description", mcp.Description("Input parameter: A short description of the item")),
		mcp.WithObject("income_account", mcp.Description("")),
		mcp.WithObject("sales_details", mcp.Description("")),
		mcp.WithBoolean("tracked", mcp.Description("Input parameter: Item is inventoried")),
		mcp.WithString("name", mcp.Description("Input parameter: Item name")),
		mcp.WithString("id", mcp.Description("Input parameter: The ID of the item.")),
		mcp.WithObject("purchase_details", mcp.Description("")),
		mcp.WithString("type", mcp.Description("Input parameter: Item type")),
		mcp.WithBoolean("purchased", mcp.Description("Input parameter: Item is available for purchase transactions")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("inventory_date", mcp.Description("Input parameter: The date of opening balance if inventory item is tracked - YYYY-MM-DD.")),
		mcp.WithString("quantity", mcp.Description("")),
		mcp.WithBoolean("sold", mcp.Description("Input parameter: Item will be available on sales transactions")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithObject("expense_account", mcp.Description("")),
		mcp.WithString("unit_price", mcp.Description("")),
		mcp.WithBoolean("active", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    InvoiceitemsaddHandler(cfg),
	}
}
