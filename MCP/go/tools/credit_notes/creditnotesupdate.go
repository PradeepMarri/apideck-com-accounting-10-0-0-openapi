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

func CreditnotesupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.CreditNote
		
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
		url := fmt.Sprintf("%s/accounting/credit-notes/%s%s", cfg.BaseURL, id, queryString)
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
		var result models.UpdateCreditNoteResponse
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

func CreateCreditnotesupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_accounting_credit-notes_id",
		mcp.WithDescription("Update Credit Note"),
		mcp.WithString("id", mcp.Required(), mcp.Description("ID of the record you are acting upon.")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("reference", mcp.Description("Input parameter: Optional reference message ie: Debit remittance detail.")),
		mcp.WithString("total_tax", mcp.Description("Input parameter: Total tax amount applied to this invoice.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: Unique identifier representing the entity")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithObject("customer", mcp.Description("Input parameter: The customer this entity is linked to.")),
		mcp.WithArray("line_items", mcp.Description("")),
		mcp.WithString("number", mcp.Description("Input parameter: Credit note number.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("remaining_credit", mcp.Description("Input parameter: Indicates the total credit amount still available to apply towards the payment.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("balance", mcp.Description("Input parameter: The balance reflecting any payments made against the transaction.")),
		mcp.WithString("currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithString("tax_code", mcp.Description("Input parameter: Applicable tax id/code override if tax is not supplied on a line item basis.")),
		mcp.WithString("type", mcp.Description("Input parameter: Type of payment")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of credit notes")),
		mcp.WithString("terms", mcp.Description("Input parameter: Optional terms to be associated with the credit note.")),
		mcp.WithObject("account", mcp.Description("")),
		mcp.WithArray("allocations", mcp.Description("")),
		mcp.WithString("date_issued", mcp.Description("Input parameter: Date credit note issued - YYYY:MM::DDThh:mm:ss.sTZD")),
		mcp.WithString("date_paid", mcp.Description("Input parameter: Date credit note paid - YYYY:MM::DDThh:mm:ss.sTZD")),
		mcp.WithBoolean("tax_inclusive", mcp.Description("Input parameter: Amounts are including tax")),
		mcp.WithString("total_amount", mcp.Required(), mcp.Description("Input parameter: Amount of transaction")),
		mcp.WithString("note", mcp.Description("Input parameter: Optional note to be associated with the credit note.")),
		mcp.WithString("currency_rate", mcp.Description("Input parameter: Currency Exchange Rate at the time entity was recorded/generated.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("sub_total", mcp.Description("Input parameter: Sub-total amount, normally before tax.")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreditnotesupdateHandler(cfg),
	}
}
