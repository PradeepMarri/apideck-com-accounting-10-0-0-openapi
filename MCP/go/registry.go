package main

import (
	"github.com/accounting-api/mcp-server/config"
	"github.com/accounting-api/mcp-server/models"
	tools_tax_rates "github.com/accounting-api/mcp-server/tools/tax_rates"
	tools_purchase_orders "github.com/accounting-api/mcp-server/tools/purchase_orders"
	tools_profit_and_loss "github.com/accounting-api/mcp-server/tools/profit_and_loss"
	tools_suppliers "github.com/accounting-api/mcp-server/tools/suppliers"
	tools_invoice_items "github.com/accounting-api/mcp-server/tools/invoice_items"
	tools_customers "github.com/accounting-api/mcp-server/tools/customers"
	tools_credit_notes "github.com/accounting-api/mcp-server/tools/credit_notes"
	tools_payments "github.com/accounting-api/mcp-server/tools/payments"
	tools_invoices "github.com/accounting-api/mcp-server/tools/invoices"
	tools_ledger_accounts "github.com/accounting-api/mcp-server/tools/ledger_accounts"
	tools_balance_sheet "github.com/accounting-api/mcp-server/tools/balance_sheet"
	tools_bills "github.com/accounting-api/mcp-server/tools/bills"
	tools_journal_entries "github.com/accounting-api/mcp-server/tools/journal_entries"
	tools_company_info "github.com/accounting-api/mcp-server/tools/company_info"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tax_rates.CreateTaxratesdeleteTool(cfg),
		tools_tax_rates.CreateTaxratesoneTool(cfg),
		tools_tax_rates.CreateTaxratesupdateTool(cfg),
		tools_purchase_orders.CreatePurchaseordersallTool(cfg),
		tools_purchase_orders.CreatePurchaseordersaddTool(cfg),
		tools_profit_and_loss.CreateProfitandlossoneTool(cfg),
		tools_suppliers.CreateSuppliersallTool(cfg),
		tools_suppliers.CreateSuppliersaddTool(cfg),
		tools_suppliers.CreateSuppliersdeleteTool(cfg),
		tools_suppliers.CreateSuppliersoneTool(cfg),
		tools_suppliers.CreateSuppliersupdateTool(cfg),
		tools_invoice_items.CreateInvoiceitemsdeleteTool(cfg),
		tools_invoice_items.CreateInvoiceitemsoneTool(cfg),
		tools_invoice_items.CreateInvoiceitemsupdateTool(cfg),
		tools_tax_rates.CreateTaxratesallTool(cfg),
		tools_tax_rates.CreateTaxratesaddTool(cfg),
		tools_customers.CreateCustomersallTool(cfg),
		tools_customers.CreateCustomersaddTool(cfg),
		tools_credit_notes.CreateCreditnotesaddTool(cfg),
		tools_credit_notes.CreateCreditnotesallTool(cfg),
		tools_payments.CreatePaymentsupdateTool(cfg),
		tools_payments.CreatePaymentsdeleteTool(cfg),
		tools_payments.CreatePaymentsoneTool(cfg),
		tools_credit_notes.CreateCreditnotesdeleteTool(cfg),
		tools_credit_notes.CreateCreditnotesoneTool(cfg),
		tools_credit_notes.CreateCreditnotesupdateTool(cfg),
		tools_invoices.CreateInvoicesallTool(cfg),
		tools_invoices.CreateInvoicesaddTool(cfg),
		tools_ledger_accounts.CreateLedgeraccountsdeleteTool(cfg),
		tools_ledger_accounts.CreateLedgeraccountsoneTool(cfg),
		tools_ledger_accounts.CreateLedgeraccountsupdateTool(cfg),
		tools_invoice_items.CreateInvoiceitemsallTool(cfg),
		tools_invoice_items.CreateInvoiceitemsaddTool(cfg),
		tools_balance_sheet.CreateBalancesheetoneTool(cfg),
		tools_bills.CreateBillsallTool(cfg),
		tools_bills.CreateBillsaddTool(cfg),
		tools_journal_entries.CreateJournalentriesdeleteTool(cfg),
		tools_journal_entries.CreateJournalentriesoneTool(cfg),
		tools_journal_entries.CreateJournalentriesupdateTool(cfg),
		tools_invoices.CreateInvoicesupdateTool(cfg),
		tools_invoices.CreateInvoicesdeleteTool(cfg),
		tools_invoices.CreateInvoicesoneTool(cfg),
		tools_journal_entries.CreateJournalentriesallTool(cfg),
		tools_journal_entries.CreateJournalentriesaddTool(cfg),
		tools_ledger_accounts.CreateLedgeraccountsallTool(cfg),
		tools_ledger_accounts.CreateLedgeraccountsaddTool(cfg),
		tools_customers.CreateCustomersdeleteTool(cfg),
		tools_customers.CreateCustomersoneTool(cfg),
		tools_customers.CreateCustomersupdateTool(cfg),
		tools_payments.CreatePaymentsallTool(cfg),
		tools_payments.CreatePaymentsaddTool(cfg),
		tools_purchase_orders.CreatePurchaseordersdeleteTool(cfg),
		tools_purchase_orders.CreatePurchaseordersoneTool(cfg),
		tools_purchase_orders.CreatePurchaseordersupdateTool(cfg),
		tools_bills.CreateBillsdeleteTool(cfg),
		tools_bills.CreateBillsoneTool(cfg),
		tools_bills.CreateBillsupdateTool(cfg),
		tools_company_info.CreateCompanyinfooneTool(cfg),
	}
}
