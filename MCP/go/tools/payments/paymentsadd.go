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

func PaymentsaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.Payment
		
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
		url := fmt.Sprintf("%s/accounting/payments%s", cfg.BaseURL, queryString)
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
		var result models.CreatePaymentResponse
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

func CreatePaymentsaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_accounting_payments",
		mcp.WithDescription("Create Payment"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithObject("account", mcp.Description("")),
		mcp.WithBoolean("reconciled", mcp.Description("Input parameter: Payment has been reconciled")),
		mcp.WithString("type", mcp.Description("Input parameter: Type of payment")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithString("reference", mcp.Description("Input parameter: Optional payment reference message ie: Debit remittance detail.")),
		mcp.WithString("display_id", mcp.Description("Input parameter: Payment id to be displayed.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: Unique identifier representing the entity")),
		mcp.WithString("payment_method_id", mcp.Description("Input parameter: Unique identifier for the payment method.")),
		mcp.WithString("accounts_receivable_account_type", mcp.Description("Input parameter: Type of accounts receivable account.")),
		mcp.WithString("downstream_id", mcp.Description("Input parameter: The third-party API ID of original entity")),
		mcp.WithString("currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithString("payment_method", mcp.Description("Input parameter: Payment method name")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of payment")),
		mcp.WithObject("supplier", mcp.Description("Input parameter: The supplier this entity is linked to.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("currency_rate", mcp.Description("Input parameter: Currency Exchange Rate at the time entity was recorded/generated.")),
		mcp.WithObject("customer", mcp.Description("Input parameter: The customer this entity is linked to.")),
		mcp.WithString("total_amount", mcp.Required(), mcp.Description("Input parameter: Amount of payment")),
		mcp.WithString("note", mcp.Description("Input parameter: Optional note to be associated with the payment.")),
		mcp.WithString("accounts_receivable_account_id", mcp.Description("Input parameter: Unique identifier for the account to allocate payment to.")),
		mcp.WithString("transaction_date", mcp.Required(), mcp.Description("Input parameter: Date transaction was entered - YYYY:MM::DDThh:mm:ss.sTZD")),
		mcp.WithArray("allocations", mcp.Description("")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PaymentsaddHandler(cfg),
	}
}
