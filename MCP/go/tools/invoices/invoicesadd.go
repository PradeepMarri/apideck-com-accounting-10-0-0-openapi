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

func InvoicesaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.Invoice
		
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
		url := fmt.Sprintf("%s/accounting/invoices%s", cfg.BaseURL, queryString)
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
		var result models.CreateInvoiceResponse
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

func CreateInvoicesaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_accounting_invoices",
		mcp.WithDescription("Create Invoice"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithString("due_date", mcp.Description("Input parameter: The invoice due date is the date on which a payment or invoice is scheduled to be received by the seller - YYYY-MM-DD.")),
		mcp.WithString("reference", mcp.Description("Input parameter: Optional invoice reference.")),
		mcp.WithString("tax_code", mcp.Description("Input parameter: Applicable tax id/code override if tax is not supplied on a line item basis.")),
		mcp.WithString("invoice_date", mcp.Description("Input parameter: Date invoice was issued - YYYY-MM-DD.")),
		mcp.WithBoolean("invoice_sent", mcp.Description("Input parameter: Invoice sent to contact/customer.")),
		mcp.WithString("number", mcp.Description("Input parameter: Invoice number.")),
		mcp.WithString("currency_rate", mcp.Description("Input parameter: Currency Exchange Rate at the time entity was recorded/generated.")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("status", mcp.Description("Input parameter: Invoice status")),
		mcp.WithString("downstream_id", mcp.Description("Input parameter: The third-party API ID of original entity")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithArray("line_items", mcp.Description("")),
		mcp.WithObject("bank_account", mcp.Description("")),
		mcp.WithBoolean("tax_inclusive", mcp.Description("Input parameter: Amounts are including tax")),
		mcp.WithString("sub_total", mcp.Description("Input parameter: Sub-total amount, normally before tax.")),
		mcp.WithString("currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithString("id", mcp.Description("Input parameter: A unique identifier for an object.")),
		mcp.WithObject("shipping_address", mcp.Description("")),
		mcp.WithString("language", mcp.Description("Input parameter: language code according to ISO 639-1. For the United States - EN")),
		mcp.WithString("customer_memo", mcp.Description("Input parameter: Customer memo")),
		mcp.WithString("total_tax", mcp.Description("Input parameter: Total tax amount applied to this invoice.")),
		mcp.WithString("deposit", mcp.Description("Input parameter: Amount of deposit made to this invoice.")),
		mcp.WithString("type", mcp.Description("Input parameter: Invoice type")),
		mcp.WithString("discount_amount", mcp.Description("Input parameter: Discount amount applied to this invoice.")),
		mcp.WithString("discount_percentage", mcp.Description("Input parameter: Discount percentage applied to this invoice.")),
		mcp.WithString("source_document_url", mcp.Description("Input parameter: URL link to a source document - shown as 'Go to [appName]' in the downstream app. Currently only supported for Xero.")),
		mcp.WithString("total", mcp.Description("Input parameter: Total amount of invoice, including tax.")),
		mcp.WithString("payment_method", mcp.Description("Input parameter: Payment method used for the transaction, such as cash, credit card, bank transfer, or check")),
		mcp.WithString("po_number", mcp.Description("Input parameter: A PO Number uniquely identifies a purchase order and is generally defined by the buyer. The buyer will match the PO number in the invoice to the Purchase Order.")),
		mcp.WithObject("customer", mcp.Description("Input parameter: The customer this entity is linked to.")),
		mcp.WithObject("tracking_category", mcp.Description("")),
		mcp.WithString("balance", mcp.Description("Input parameter: Balance of invoice due.")),
		mcp.WithObject("ledger_account", mcp.Description("")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithObject("billing_address", mcp.Description("")),
		mcp.WithString("template_id", mcp.Description("Input parameter: Optional invoice template")),
		mcp.WithBoolean("accounting_by_row", mcp.Description("Input parameter: Indicates if accounting by row is used (true) or not (false). Accounting by row means that a separate ledger transaction is created for each row.")),
		mcp.WithString("channel", mcp.Description("Input parameter: The channel through which the transaction is processed.")),
		mcp.WithString("terms", mcp.Description("Input parameter: Terms of payment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    InvoicesaddHandler(cfg),
	}
}
