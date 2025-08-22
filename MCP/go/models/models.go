package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteCustomerResponse represents the DeleteCustomerResponse schema from the OpenAPI specification
type DeleteCustomerResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// LinkedSupplier represents the LinkedSupplier schema from the OpenAPI specification
type LinkedSupplier struct {
	Address Address `json:"address,omitempty"`
	Company_name string `json:"company_name,omitempty"` // The company name of the supplier.
	Display_id string `json:"display_id,omitempty"` // The display ID of the supplier.
	Display_name string `json:"display_name,omitempty"` // The display name of the supplier.
	Id string `json:"id"` // The ID of the supplier this entity is linked to.
}

// PassThroughQuery represents the PassThroughQuery schema from the OpenAPI specification
type PassThroughQuery struct {
	Example_downstream_property string `json:"example_downstream_property,omitempty"` // All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
}

// UpdatePurchaseOrderResponse represents the UpdatePurchaseOrderResponse schema from the OpenAPI specification
type UpdatePurchaseOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// UpdateJournalEntryResponse represents the UpdateJournalEntryResponse schema from the OpenAPI specification
type UpdateJournalEntryResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// InvoiceItemAssetAccount represents the InvoiceItemAssetAccount schema from the OpenAPI specification
type InvoiceItemAssetAccount struct {
	Nominal_code string `json:"nominal_code,omitempty"` // The nominal code of the account.
	Code string `json:"code,omitempty"` // The code assigned to the account.
	Id string `json:"id,omitempty"` // The unique identifier for the account.
	Name string `json:"name,omitempty"` // The name of the account.
}

// CreateTaxRateResponse represents the CreateTaxRateResponse schema from the OpenAPI specification
type CreateTaxRateResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// DeleteJournalEntryResponse represents the DeleteJournalEntryResponse schema from the OpenAPI specification
type DeleteJournalEntryResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UpdateCreditNoteResponse represents the UpdateCreditNoteResponse schema from the OpenAPI specification
type UpdateCreditNoteResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// DeleteInvoiceResponse represents the DeleteInvoiceResponse schema from the OpenAPI specification
type DeleteInvoiceResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data InvoiceResponse `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// SocialLink represents the SocialLink schema from the OpenAPI specification
type SocialLink struct {
	Id string `json:"id,omitempty"` // Unique identifier of the social link
	TypeField string `json:"type,omitempty"` // Type of the social link, e.g. twitter
	Url string `json:"url"` // URL of the social link, e.g. https://www.twitter.com/apideck
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// InvoiceItemsFilter represents the InvoiceItemsFilter schema from the OpenAPI specification
type InvoiceItemsFilter struct {
	Name string `json:"name,omitempty"` // Name of Invoice Items to search for
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Name string `json:"name,omitempty"` // The name of the address.
	Notes string `json:"notes,omitempty"` // Additional notes
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
	Website string `json:"website,omitempty"` // Website of the address
	Email string `json:"email,omitempty"` // Email address of the address
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Street_number string `json:"street_number,omitempty"` // Street number
	State string `json:"state,omitempty"` // Name of state
	TypeField string `json:"type,omitempty"` // The type of address.
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	City string `json:"city,omitempty"` // Name of city.
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// BankAccount represents the BankAccount schema from the OpenAPI specification
type BankAccount struct {
	Account_number string `json:"account_number,omitempty"` // A bank account number is a number that is tied to your bank account. If you have several bank accounts, such as personal, joint, business (and so on), each account will have a different account number.
	Iban string `json:"iban,omitempty"` // The International Bank Account Number (IBAN).
	Bic string `json:"bic,omitempty"` // The Bank Identifier Code (BIC).
	Branch_identifier string `json:"branch_identifier,omitempty"` // A branch identifier is a unique identifier for a branch of a bank or financial institution.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Routing_number string `json:"routing_number,omitempty"` // A routing number is a nine-digit code used to identify a financial institution in the United States.
	Account_name string `json:"account_name,omitempty"` // The name which you used in opening your bank account.
	Account_type string `json:"account_type,omitempty"` // The type of bank account.
	Bank_code string `json:"bank_code,omitempty"` // A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.
	Bank_name string `json:"bank_name,omitempty"` // The name of the bank
	Bsb_number string `json:"bsb_number,omitempty"` // A BSB is a 6 digit numeric code used for identifying the branch of an Australian or New Zealand bank or financial institution.
}

// LinkedTrackingCategory represents the LinkedTrackingCategory schema from the OpenAPI specification
type LinkedTrackingCategory struct {
	Id string `json:"id,omitempty"` // The unique identifier for the tracking category.
	Name string `json:"name,omitempty"` // The name of the tracking category.
}

// CustomersFilter represents the CustomersFilter schema from the OpenAPI specification
type CustomersFilter struct {
	Company_name string `json:"company_name,omitempty"` // Company Name of customer to search for
	Display_name string `json:"display_name,omitempty"` // Display Name of customer to search for
	Email string `json:"email,omitempty"` // Email of customer to search for
	First_name string `json:"first_name,omitempty"` // First name of customer to search for
	Last_name string `json:"last_name,omitempty"` // Last name of customer to search for
	Status string `json:"status,omitempty"` // Status of customer to filter on
}

// LinkedLedgerAccount represents the LinkedLedgerAccount schema from the OpenAPI specification
type LinkedLedgerAccount struct {
	Code string `json:"code,omitempty"` // The code assigned to the account.
	Id string `json:"id,omitempty"` // The unique identifier for the account.
	Name string `json:"name,omitempty"` // The name of the account.
	Nominal_code string `json:"nominal_code,omitempty"` // The nominal code of the account.
}

// TaxRate represents the TaxRate schema from the OpenAPI specification
type TaxRate struct {
	Original_tax_rate_id string `json:"original_tax_rate_id,omitempty"` // ID of the original tax rate from which the new tax rate is derived. Helps to understand the relationship between corresponding tax rate entities.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Name string `json:"name,omitempty"` // Name assigned to identify this tax rate.
	Report_tax_type string `json:"report_tax_type,omitempty"` // Report Tax type to aggregate tax collected or paid for reporting purposes
	Id string `json:"id,omitempty"` // ID assigned to identify this tax rate.
	Tax_payable_account_id string `json:"tax_payable_account_id,omitempty"` // Unique identifier for the account for tax collected.
	Code string `json:"code,omitempty"` // Tax code assigned to identify this tax rate.
	Status string `json:"status,omitempty"` // Tax rate status
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	TypeField string `json:"type,omitempty"` // Tax type used to indicate the source of tax collected or paid
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Components []interface{} `json:"components,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Effective_tax_rate float64 `json:"effective_tax_rate,omitempty"` // Effective tax rate
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Description string `json:"description,omitempty"` // Description of tax rate
	Tax_remitted_account_id string `json:"tax_remitted_account_id,omitempty"` // Unique identifier for the account for tax remitted.
	Total_tax_rate float64 `json:"total_tax_rate,omitempty"` // Not compounded sum of the components of a tax rate
}

// LinkedInvoiceItem represents the LinkedInvoiceItem schema from the OpenAPI specification
type LinkedInvoiceItem struct {
	Name string `json:"name,omitempty"` // User defined item name
	Code string `json:"code,omitempty"` // User defined item code
	Id string `json:"id,omitempty"` // ID of the linked item. A reference to the [invoice item](https://developers.apideck.com/apis/accounting/reference#tag/Invoice-Items) that was used to create this line item
}

// ProfitAndLossFilter represents the ProfitAndLossFilter schema from the OpenAPI specification
type ProfitAndLossFilter struct {
	Customer_id string `json:"customer_id,omitempty"` // Filter by customer id
	End_date string `json:"end_date,omitempty"` // Filter by end date. If end date is given, start date is required.
	Start_date string `json:"start_date,omitempty"` // Filter by start date. If start date is given, end date is required.
}

// GetJournalEntryResponse represents the GetJournalEntryResponse schema from the OpenAPI specification
type GetJournalEntryResponse struct {
	Data JournalEntry `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// PaymentsFilter represents the PaymentsFilter schema from the OpenAPI specification
type PaymentsFilter struct {
	Invoice_number string `json:"invoice_number,omitempty"` // Invoice number for payments to filter on
}

// UpdateCustomerResponse represents the UpdateCustomerResponse schema from the OpenAPI specification
type UpdateCustomerResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// UpdateLedgerAccountResponse represents the UpdateLedgerAccountResponse schema from the OpenAPI specification
type UpdateLedgerAccountResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetCreditNotesResponse represents the GetCreditNotesResponse schema from the OpenAPI specification
type GetCreditNotesResponse struct {
	Data []CreditNote `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CompanyInfo represents the CompanyInfo schema from the OpenAPI specification
type CompanyInfo struct {
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Sales_tax_enabled bool `json:"sales_tax_enabled,omitempty"` // Whether sales tax is enabled for the company
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Company_start_date string `json:"company_start_date,omitempty"` // Date when company file was created
	Emails []Email `json:"emails,omitempty"`
	Status string `json:"status,omitempty"` // Based on the status some functionality is enabled or disabled.
	Sales_tax_number string `json:"sales_tax_number,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Fiscal_year_start_month string `json:"fiscal_year_start_month,omitempty"` // The start month of fiscal year.
	Legal_name string `json:"legal_name,omitempty"` // The legal name of the company
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Automated_sales_tax bool `json:"automated_sales_tax,omitempty"` // Whether sales tax is calculated automatically for the company
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Default_sales_tax TaxRate `json:"default_sales_tax,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Company_name string `json:"company_name,omitempty"` // The name of the company.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Addresses []Address `json:"addresses,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// SuppliersFilter represents the SuppliersFilter schema from the OpenAPI specification
type SuppliersFilter struct {
	First_name string `json:"first_name,omitempty"` // First name of supplier to search for
	Last_name string `json:"last_name,omitempty"` // Last name of supplier to search for
	Company_name string `json:"company_name,omitempty"` // Company Name of supplier to search for
	Display_name string `json:"display_name,omitempty"` // Display Name of supplier to search for
	Email string `json:"email,omitempty"` // Email of supplier to search for
}

// GetLedgerAccountResponse represents the GetLedgerAccountResponse schema from the OpenAPI specification
type GetLedgerAccountResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data LedgerAccount `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetBillResponse represents the GetBillResponse schema from the OpenAPI specification
type GetBillResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Bill `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// InvoicesSort represents the InvoicesSort schema from the OpenAPI specification
type InvoicesSort struct {
	By string `json:"by,omitempty"` // The field on which to sort the Invoices
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
}

// DeletePaymentResponse represents the DeletePaymentResponse schema from the OpenAPI specification
type DeletePaymentResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Parent LinkedParentCustomer `json:"parent,omitempty"` // The parent customer this entity is linked to.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	Tax_number string `json:"tax_number,omitempty"`
	Individual bool `json:"individual,omitempty"` // Is this an individual or business customer
	Display_id string `json:"display_id,omitempty"` // Display ID
	Project bool `json:"project,omitempty"` // If true, indicates this is a Project.
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id"` // A unique identifier for an object.
	Emails []Email `json:"emails,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Bank_accounts []BankAccount `json:"bank_accounts,omitempty"`
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Title string `json:"title,omitempty"` // The job title of the person.
	Addresses []Address `json:"addresses,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Websites []Website `json:"websites,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Display_name string `json:"display_name,omitempty"` // Display name
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Middle_name string `json:"middle_name,omitempty"` // Middle name of the person.
	Account LinkedLedgerAccount `json:"account,omitempty"`
	Status string `json:"status,omitempty"` // Customer status
	Notes string `json:"notes,omitempty"` // Some notes about this customer
	Company_name string `json:"company_name,omitempty"` // The name of the company.
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Id string `json:"id,omitempty"` // Unique identifier for the email address
	TypeField string `json:"type,omitempty"` // Email type
	Email string `json:"email"` // Email address
}

// CreateLedgerAccountResponse represents the CreateLedgerAccountResponse schema from the OpenAPI specification
type CreateLedgerAccountResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// LinkedTaxRate represents the LinkedTaxRate schema from the OpenAPI specification
type LinkedTaxRate struct {
	Code string `json:"code,omitempty"` // Tax rate code
	Id string `json:"id,omitempty"` // The ID of the object.
	Name string `json:"name,omitempty"` // Name of the tax rate
	Rate float64 `json:"rate,omitempty"` // Rate of the tax rate
}

// GetCustomersResponse represents the GetCustomersResponse schema from the OpenAPI specification
type GetCustomersResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Customer `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// UpdateInvoiceItemsResponse represents the UpdateInvoiceItemsResponse schema from the OpenAPI specification
type UpdateInvoiceItemsResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// InvoiceItem represents the InvoiceItem schema from the OpenAPI specification
type InvoiceItem struct {
	Taxable bool `json:"taxable,omitempty"` // If true, transactions for this item are taxable
	Code string `json:"code,omitempty"` // User defined item code
	Description string `json:"description,omitempty"` // A short description of the item
	Income_account LinkedLedgerAccount `json:"income_account,omitempty"`
	Sales_details map[string]interface{} `json:"sales_details,omitempty"`
	Tracked bool `json:"tracked,omitempty"` // Item is inventoried
	Name string `json:"name,omitempty"` // Item name
	Id string `json:"id,omitempty"` // The ID of the item.
	Purchase_details map[string]interface{} `json:"purchase_details,omitempty"`
	TypeField string `json:"type,omitempty"` // Item type
	Purchased bool `json:"purchased,omitempty"` // Item is available for purchase transactions
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Inventory_date string `json:"inventory_date,omitempty"` // The date of opening balance if inventory item is tracked - YYYY-MM-DD.
	Quantity float64 `json:"quantity,omitempty"`
	Sold bool `json:"sold,omitempty"` // Item will be available on sales transactions
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Expense_account LinkedLedgerAccount `json:"expense_account,omitempty"`
	Unit_price float64 `json:"unit_price,omitempty"`
	Active bool `json:"active,omitempty"`
	Asset_account LinkedLedgerAccount `json:"asset_account,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Tracking_category LinkedTrackingCategory `json:"tracking_category,omitempty"`
}

// GetProfitAndLossResponse represents the GetProfitAndLossResponse schema from the OpenAPI specification
type GetProfitAndLossResponse struct {
	Data ProfitAndLoss `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// DeleteLedgerAccountResponse represents the DeleteLedgerAccountResponse schema from the OpenAPI specification
type DeleteLedgerAccountResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// CreditNote represents the CreditNote schema from the OpenAPI specification
type CreditNote struct {
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Balance float64 `json:"balance,omitempty"` // The balance reflecting any payments made against the transaction.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Tax_code string `json:"tax_code,omitempty"` // Applicable tax id/code override if tax is not supplied on a line item basis.
	TypeField string `json:"type,omitempty"` // Type of payment
	Status string `json:"status,omitempty"` // Status of credit notes
	Terms string `json:"terms,omitempty"` // Optional terms to be associated with the credit note.
	Account LinkedLedgerAccount `json:"account,omitempty"`
	Allocations []interface{} `json:"allocations,omitempty"`
	Date_issued string `json:"date_issued,omitempty"` // Date credit note issued - YYYY:MM::DDThh:mm:ss.sTZD
	Date_paid string `json:"date_paid,omitempty"` // Date credit note paid - YYYY:MM::DDThh:mm:ss.sTZD
	Tax_inclusive bool `json:"tax_inclusive,omitempty"` // Amounts are including tax
	Total_amount float64 `json:"total_amount"` // Amount of transaction
	Note string `json:"note,omitempty"` // Optional note to be associated with the credit note.
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Sub_total float64 `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Reference string `json:"reference,omitempty"` // Optional reference message ie: Debit remittance detail.
	Total_tax float64 `json:"total_tax,omitempty"` // Total tax amount applied to this invoice.
	Id string `json:"id"` // Unique identifier representing the entity
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Customer LinkedCustomer `json:"customer,omitempty"` // The customer this entity is linked to.
	Line_items []InvoiceLineItem `json:"line_items,omitempty"`
	Number string `json:"number,omitempty"` // Credit note number.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Remaining_credit float64 `json:"remaining_credit,omitempty"` // Indicates the total credit amount still available to apply towards the payment.
}

// JournalEntryLineItem represents the JournalEntryLineItem schema from the OpenAPI specification
type JournalEntryLineItem struct {
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Total_amount float64 `json:"total_amount,omitempty"` // Debit entries are considered positive, and credit entries are considered negative.
	Tracking_category LinkedTrackingCategory `json:"tracking_category,omitempty"`
	Description string `json:"description,omitempty"` // User defined description
	Supplier LinkedSupplier `json:"supplier,omitempty"` // The supplier this entity is linked to.
	TypeField string `json:"type"` // Debit entries are considered positive, and credit entries are considered negative.
	Customer LinkedCustomer `json:"customer,omitempty"` // The customer this entity is linked to.
	Department_id string `json:"department_id,omitempty"` // A unique identifier for an object.
	Ledger_account LinkedLedgerAccount `json:"ledger_account"`
	Location_id string `json:"location_id,omitempty"` // A unique identifier for an object.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Tax_amount float64 `json:"tax_amount,omitempty"` // Tax amount
	Sub_total float64 `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
}

// DeleteBillResponse represents the DeleteBillResponse schema from the OpenAPI specification
type DeleteBillResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	Url string `json:"url"` // The website URL
	Id string `json:"id,omitempty"` // Unique identifier for the website
	TypeField string `json:"type,omitempty"` // The type of website
}

// UpdateBillResponse represents the UpdateBillResponse schema from the OpenAPI specification
type UpdateBillResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// InvoiceLineItem represents the InvoiceLineItem schema from the OpenAPI specification
type InvoiceLineItem struct {
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Line_number int `json:"line_number,omitempty"` // Line number in the invoice
	Department_id string `json:"department_id,omitempty"` // Department id
	Discount_amount float64 `json:"discount_amount,omitempty"` // Discount amount applied to the line item when supported downstream.
	Unit_of_measure string `json:"unit_of_measure,omitempty"` // Description of the unit type the item is sold as, ie: kg, hour.
	Code string `json:"code,omitempty"` // User defined item code
	Unit_price float64 `json:"unit_price,omitempty"`
	TypeField string `json:"type,omitempty"` // Item type
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Quantity float64 `json:"quantity,omitempty"`
	Row_id string `json:"row_id,omitempty"` // Row ID
	Description string `json:"description,omitempty"` // User defined description
	Location_id string `json:"location_id,omitempty"` // Location id
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Discount_percentage float64 `json:"discount_percentage,omitempty"` // Discount percentage applied to the line item when supported downstream.
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Ledger_account LinkedLedgerAccount `json:"ledger_account,omitempty"`
	Total_amount float64 `json:"total_amount,omitempty"` // Total amount of the line item
	Item LinkedInvoiceItem `json:"item,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Tax_amount float64 `json:"tax_amount,omitempty"` // Tax amount
}

// GetSupplierResponse represents the GetSupplierResponse schema from the OpenAPI specification
type GetSupplierResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Supplier `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// Invoice represents the Invoice schema from the OpenAPI specification
type Invoice struct {
	Due_date string `json:"due_date,omitempty"` // The invoice due date is the date on which a payment or invoice is scheduled to be received by the seller - YYYY-MM-DD.
	Reference string `json:"reference,omitempty"` // Optional invoice reference.
	Tax_code string `json:"tax_code,omitempty"` // Applicable tax id/code override if tax is not supplied on a line item basis.
	Invoice_date string `json:"invoice_date,omitempty"` // Date invoice was issued - YYYY-MM-DD.
	Invoice_sent bool `json:"invoice_sent,omitempty"` // Invoice sent to contact/customer.
	Number string `json:"number,omitempty"` // Invoice number.
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Status string `json:"status,omitempty"` // Invoice status
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Line_items []InvoiceLineItem `json:"line_items,omitempty"`
	Bank_account BankAccount `json:"bank_account,omitempty"`
	Tax_inclusive bool `json:"tax_inclusive,omitempty"` // Amounts are including tax
	Sub_total float64 `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Shipping_address Address `json:"shipping_address,omitempty"`
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Customer_memo string `json:"customer_memo,omitempty"` // Customer memo
	Total_tax float64 `json:"total_tax,omitempty"` // Total tax amount applied to this invoice.
	Deposit float64 `json:"deposit,omitempty"` // Amount of deposit made to this invoice.
	TypeField string `json:"type,omitempty"` // Invoice type
	Discount_amount float64 `json:"discount_amount,omitempty"` // Discount amount applied to this invoice.
	Discount_percentage float64 `json:"discount_percentage,omitempty"` // Discount percentage applied to this invoice.
	Source_document_url string `json:"source_document_url,omitempty"` // URL link to a source document - shown as 'Go to [appName]' in the downstream app. Currently only supported for Xero.
	Total float64 `json:"total,omitempty"` // Total amount of invoice, including tax.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	Po_number string `json:"po_number,omitempty"` // A PO Number uniquely identifies a purchase order and is generally defined by the buyer. The buyer will match the PO number in the invoice to the Purchase Order.
	Customer LinkedCustomer `json:"customer,omitempty"` // The customer this entity is linked to.
	Tracking_category LinkedTrackingCategory `json:"tracking_category,omitempty"`
	Balance float64 `json:"balance,omitempty"` // Balance of invoice due.
	Ledger_account LinkedLedgerAccount `json:"ledger_account,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Billing_address Address `json:"billing_address,omitempty"`
	Template_id string `json:"template_id,omitempty"` // Optional invoice template
	Accounting_by_row bool `json:"accounting_by_row,omitempty"` // Indicates if accounting by row is used (true) or not (false). Accounting by row means that a separate ledger transaction is created for each row.
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Terms string `json:"terms,omitempty"` // Terms of payment.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
}

// LinkedParentCustomer represents the LinkedParentCustomer schema from the OpenAPI specification
type LinkedParentCustomer struct {
	Id string `json:"id"` // The parent ID of the customer this entity is linked to.
	Name string `json:"name,omitempty"` // The name of the parent customer.
}

// DeleteCreditNoteResponse represents the DeleteCreditNoteResponse schema from the OpenAPI specification
type DeleteCreditNoteResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetInvoicesResponse represents the GetInvoicesResponse schema from the OpenAPI specification
type GetInvoicesResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Invoice `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// ProfitAndLossSection represents the ProfitAndLossSection schema from the OpenAPI specification
type ProfitAndLossSection struct {
	TypeField string `json:"type"`
	Id string `json:"id,omitempty"`
	Records []interface{} `json:"records,omitempty"`
	Title string `json:"title,omitempty"`
	Total float64 `json:"total,omitempty"`
}

// InvoiceResponse represents the InvoiceResponse schema from the OpenAPI specification
type InvoiceResponse struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
}

// UpdateTaxRateResponse represents the UpdateTaxRateResponse schema from the OpenAPI specification
type UpdateTaxRateResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetCustomerResponse represents the GetCustomerResponse schema from the OpenAPI specification
type GetCustomerResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Customer `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetPurchaseOrderResponse represents the GetPurchaseOrderResponse schema from the OpenAPI specification
type GetPurchaseOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data PurchaseOrder `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// CreateJournalEntryResponse represents the CreateJournalEntryResponse schema from the OpenAPI specification
type CreateJournalEntryResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// PurchaseOrder represents the PurchaseOrder schema from the OpenAPI specification
type PurchaseOrder struct {
	Accounting_by_row bool `json:"accounting_by_row,omitempty"` // Indicates if accounting by row is used (true) or not (false). Accounting by row means that a separate ledger transaction is created for each row.
	Due_date string `json:"due_date,omitempty"` // The due date is the date on which a payment is scheduled to be received - YYYY-MM-DD.
	Expected_arrival_date string `json:"expected_arrival_date,omitempty"` // The date on which the order is expected to arrive - YYYY-MM-DD.
	Ledger_account LinkedLedgerAccount `json:"ledger_account,omitempty"`
	Line_items []InvoiceLineItem `json:"line_items,omitempty"`
	Memo string `json:"memo,omitempty"` // Message for the supplier. This text appears on the Purchase Order.
	Sub_total float64 `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
	Delivery_date string `json:"delivery_date,omitempty"` // The date on which the purchase order is to be delivered - YYYY-MM-DD.
	Total float64 `json:"total,omitempty"` // Total amount of invoice, including tax.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Total_tax float64 `json:"total_tax,omitempty"` // Total tax amount applied to this invoice.
	Shipping_address Address `json:"shipping_address,omitempty"`
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Template_id string `json:"template_id,omitempty"` // Optional purchase order template
	Tax_inclusive bool `json:"tax_inclusive,omitempty"` // Amounts are including tax
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
	Tax_code string `json:"tax_code,omitempty"` // Applicable tax id/code override if tax is not supplied on a line item basis.
	Status string `json:"status,omitempty"`
	Supplier LinkedSupplier `json:"supplier,omitempty"` // The supplier this entity is linked to.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Issued_date string `json:"issued_date,omitempty"` // Date purchase order was issued - YYYY-MM-DD.
	Bank_account BankAccount `json:"bank_account,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Reference string `json:"reference,omitempty"` // Optional purchase order reference.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Discount_percentage float64 `json:"discount_percentage,omitempty"` // Discount percentage applied to this transaction.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Po_number string `json:"po_number,omitempty"` // A PO Number uniquely identifies a purchase order and is generally defined by the buyer.
}

// BalanceSheet represents the BalanceSheet schema from the OpenAPI specification
type BalanceSheet struct {
	Report_name string `json:"report_name"` // The name of the report
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Equity map[string]interface{} `json:"equity"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Assets map[string]interface{} `json:"assets"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Start_date string `json:"start_date"` // The start date of the report
	End_date string `json:"end_date,omitempty"` // The start date of the report
	Liabilities map[string]interface{} `json:"liabilities"`
}

// Company represents the Company schema from the OpenAPI specification
type Company struct {
	Custom_fields []CustomField `json:"custom_fields,omitempty"`
	Image string `json:"image,omitempty"` // The Image URL of the company
	Interaction_count int `json:"interaction_count,omitempty"` // Number of interactions
	Annual_revenue string `json:"annual_revenue,omitempty"` // The annual revenue of the company
	Payee_number string `json:"payee_number,omitempty"` // A payee number is a unique number that identifies a payee for tax purposes.
	Created_by string `json:"created_by,omitempty"` // Created by user ID
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Tags []string `json:"tags,omitempty"`
	Owner_id string `json:"owner_id,omitempty"` // Owner ID
	Salutation string `json:"salutation,omitempty"` // A formal salutation for the person. For example, 'Mr', 'Mrs'
	Fax string `json:"fax,omitempty"` // The fax number of the company
	Number_of_employees string `json:"number_of_employees,omitempty"` // Number of employees
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Sales_tax_number string `json:"sales_tax_number,omitempty"` // A sales tax number is a unique number that identifies a company for tax purposes.
	Abn_branch string `json:"abn_branch,omitempty"` // An ABN Branch (also known as a GST Branch) is used if part of your business needs to account for GST separately from its parent entity.
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Abn_or_tfn string `json:"abn_or_tfn,omitempty"` // An ABN is necessary for operating a business, while a TFN (Tax File Number) is required for any person working in Australia.
	Id string `json:"id,omitempty"` // Unique identifier for the company
	Websites []Website `json:"websites,omitempty"`
	Status string `json:"status,omitempty"` // The status of the company
	Updated_at string `json:"updated_at,omitempty"` // Last updated date
	Addresses []Address `json:"addresses,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Creation date
	Vat_number string `json:"vat_number,omitempty"` // The VAT number of the company
	Acn string `json:"acn,omitempty"` // The Australian Company Number (ACN) is a nine digit number with the last digit being a check digit calculated using a modified modulus 10 calculation. ASIC has adopted a convention of always printing and displaying the ACN in the format XXX XXX XXX; three blocks of three characters, each block separated by a blank.
	Birthday string `json:"birthday,omitempty"` // The date of birth of the person.
	Description string `json:"description,omitempty"` // A description of the company
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Row_type map[string]interface{} `json:"row_type,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // Updated by user ID
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Deleted bool `json:"deleted,omitempty"` // Whether the company is deleted or not
	Name string `json:"name"` // Name of the company
	Last_activity_at string `json:"last_activity_at,omitempty"` // Last activity date
	Emails []Email `json:"emails,omitempty"`
	Bank_accounts []BankAccount `json:"bank_accounts,omitempty"`
	Parent_id string `json:"parent_id,omitempty"` // Parent ID
	Social_links []SocialLink `json:"social_links,omitempty"`
	Ownership string `json:"ownership,omitempty"` // The ownership indicates the type of ownership of the company.
	Read_only bool `json:"read_only,omitempty"` // Whether the company is read-only or not
	Industry string `json:"industry,omitempty"` // The industry represents the type of business the company is in.
}

// GetPaymentsResponse represents the GetPaymentsResponse schema from the OpenAPI specification
type GetPaymentsResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Payment `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
}

// CreateBillResponse represents the CreateBillResponse schema from the OpenAPI specification
type CreateBillResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// CreateSupplierResponse represents the CreateSupplierResponse schema from the OpenAPI specification
type CreateSupplierResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CreatePaymentResponse represents the CreatePaymentResponse schema from the OpenAPI specification
type CreatePaymentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// LedgerAccount represents the LedgerAccount schema from the OpenAPI specification
type LedgerAccount struct {
	Level float64 `json:"level,omitempty"`
	Bank_account BankAccount `json:"bank_account,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Sub_accounts []interface{} `json:"sub_accounts,omitempty"` // The sub accounts of the account.
	Categories []map[string]interface{} `json:"categories,omitempty"` // The categories of the account.
	Display_id string `json:"display_id,omitempty"` // The human readable display ID used when displaying the account
	Active bool `json:"active,omitempty"` // Whether the account is active or not.
	Nominal_code string `json:"nominal_code,omitempty"` // The nominal code of the ledger account.
	Last_reconciliation_date string `json:"last_reconciliation_date,omitempty"` // Reconciliation Date means the last calendar day of each Reconciliation Period.
	Header bool `json:"header,omitempty"` // Whether the account is a header or not.
	Current_balance float64 `json:"current_balance,omitempty"` // The current balance of the account.
	Code string `json:"code,omitempty"` // The code assigned to the account.
	Fully_qualified_name string `json:"fully_qualified_name,omitempty"` // The fully qualified name of the account.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Sub_account bool `json:"sub_account,omitempty"` // Whether the account is a sub account or not.
	Status string `json:"status,omitempty"` // The status of the account.
	TypeField string `json:"type,omitempty"` // The type of account.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Tax_type string `json:"tax_type,omitempty"` // The tax type of the account.
	Description string `json:"description,omitempty"` // The description of the account.
	Classification string `json:"classification,omitempty"` // The classification of account.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Opening_balance float64 `json:"opening_balance,omitempty"` // The opening balance of the account.
	Parent_account map[string]interface{} `json:"parent_account,omitempty"`
	Sub_type string `json:"sub_type,omitempty"` // The sub type of account.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Name string `json:"name,omitempty"` // The name of the account.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Lead_id string `json:"lead_id,omitempty"` // The lead the contact is associated with.
	Department string `json:"department,omitempty"` // The department of the contact.
	Last_name string `json:"last_name,omitempty"` // The last name of the contact.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	First_email_at string `json:"first_email_at,omitempty"` // The first email date of the contact.
	Prefix string `json:"prefix,omitempty"` // The prefix of the contact.
	First_name string `json:"first_name,omitempty"` // The first name of the contact.
	Current_balance float64 `json:"current_balance,omitempty"` // The current balance of the contact.
	TypeField string `json:"type,omitempty"` // The type of the contact.
	Emails []Email `json:"emails,omitempty"`
	Status string `json:"status,omitempty"` // The status of the contact.
	Addresses []Address `json:"addresses,omitempty"`
	Fax string `json:"fax,omitempty"` // The fax number of the contact.
	Id string `json:"id,omitempty"` // Unique identifier for the contact.
	Photo_url string `json:"photo_url,omitempty"` // The URL of the photo of a person.
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Description string `json:"description,omitempty"` // The description of the contact.
	Last_activity_at string `json:"last_activity_at,omitempty"` // The last activity date of the contact.
	Company_name string `json:"company_name,omitempty"` // The name of the company the contact is associated with.
	Suffix string `json:"suffix,omitempty"` // The suffix of the contact.
	Created_at string `json:"created_at,omitempty"` // The creation date of the contact.
	Name string `json:"name"` // Full name of the contact.
	Owner_id string `json:"owner_id,omitempty"` // The owner of the contact.
	Gender string `json:"gender,omitempty"` // The gender of the contact.
	Title string `json:"title,omitempty"` // The job title of the contact.
	Image string `json:"image,omitempty"`
	Middle_name string `json:"middle_name,omitempty"` // The middle name of the contact.
	Custom_fields []CustomField `json:"custom_fields,omitempty"`
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Company_id string `json:"company_id,omitempty"` // The company the contact is associated with.
	Websites []Website `json:"websites,omitempty"`
	First_call_at string `json:"first_call_at,omitempty"` // The first call date of the contact.
	Social_links []SocialLink `json:"social_links,omitempty"`
	Email_domain string `json:"email_domain,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The last update date of the contact.
	Birthday string `json:"birthday,omitempty"` // The birthday of the contact.
	Active bool `json:"active,omitempty"` // The active status of the contact.
	Lead_source string `json:"lead_source,omitempty"` // The lead source of the contact.
}

// DeleteSupplierResponse represents the DeleteSupplierResponse schema from the OpenAPI specification
type DeleteSupplierResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// JournalEntry represents the JournalEntry schema from the OpenAPI specification
type JournalEntry struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Tax_type string `json:"tax_type,omitempty"` // The specific category of tax associated with a transaction like sales or purchase
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Memo string `json:"memo,omitempty"` // Reference for the journal entry.
	Posted_at string `json:"posted_at,omitempty"` // This is the date on which the journal entry was added. This can be different from the creation date and can also be backdated.
	Number string `json:"number,omitempty"` // Journal entry number.
	Tax_code string `json:"tax_code,omitempty"` // Applicable tax id/code override if tax is not supplied on a line item basis.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Journal_symbol string `json:"journal_symbol,omitempty"` // Journal symbol of the entry. For example IND for indirect costs
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Title string `json:"title,omitempty"` // Journal entry title
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Line_items []JournalEntryLineItem `json:"line_items,omitempty"` // Requires a minimum of 2 line items that sum to 0
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
}

// BillsSort represents the BillsSort schema from the OpenAPI specification
type BillsSort struct {
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
	By string `json:"by,omitempty"` // The field on which to sort the Bills
}

// TaxRatesFilter represents the TaxRatesFilter schema from the OpenAPI specification
type TaxRatesFilter struct {
	Assets bool `json:"assets,omitempty"` // Boolean to describe if tax rate can be used for asset accounts
	Equity bool `json:"equity,omitempty"` // Boolean to describe if tax rate can be used for equity accounts
	Expenses bool `json:"expenses,omitempty"` // Boolean to describe if tax rate can be used for expense accounts
	Liabilities bool `json:"liabilities,omitempty"` // Boolean to describe if tax rate can be used for liability accounts
	Revenue bool `json:"revenue,omitempty"` // Boolean to describe if tax rate can be used for revenue accounts
}

// UpdatePaymentResponse represents the UpdatePaymentResponse schema from the OpenAPI specification
type UpdatePaymentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// BillLineItem represents the BillLineItem schema from the OpenAPI specification
type BillLineItem struct {
	Department_id string `json:"department_id,omitempty"` // Department id
	Tax_amount float64 `json:"tax_amount,omitempty"` // Tax amount
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Total_amount float64 `json:"total_amount,omitempty"` // Total amount of the line item
	Unit_price float64 `json:"unit_price,omitempty"`
	Code string `json:"code,omitempty"` // User defined item code
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Description string `json:"description,omitempty"` // User defined description
	Quantity float64 `json:"quantity,omitempty"`
	Line_number int `json:"line_number,omitempty"` // Line number in the invoice
	TypeField string `json:"type,omitempty"` // Bill Line Item type
	Discount_amount float64 `json:"discount_amount,omitempty"` // Discount amount applied to the line item when supported downstream.
	Location_id string `json:"location_id,omitempty"` // Location id
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Discount_percentage float64 `json:"discount_percentage,omitempty"` // Discount percentage applied to the line item when supported downstream.
	Ledger_account LinkedLedgerAccount `json:"ledger_account,omitempty"`
	Row_id string `json:"row_id,omitempty"` // Row ID
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Unit_of_measure string `json:"unit_of_measure,omitempty"` // Description of the unit type the item is sold as, ie: kg, hour.
	Item LinkedInvoiceItem `json:"item,omitempty"`
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// UpdateInvoiceResponse represents the UpdateInvoiceResponse schema from the OpenAPI specification
type UpdateInvoiceResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data InvoiceResponse `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// CreatePurchaseOrderResponse represents the CreatePurchaseOrderResponse schema from the OpenAPI specification
type CreatePurchaseOrderResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetPaymentResponse represents the GetPaymentResponse schema from the OpenAPI specification
type GetPaymentResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Payment `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// ProfitAndLoss represents the ProfitAndLoss schema from the OpenAPI specification
type ProfitAndLoss struct {
	Currency string `json:"currency"`
	End_date string `json:"end_date,omitempty"` // The start date of the report
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Income map[string]interface{} `json:"income"`
	Report_name string `json:"report_name"` // The name of the report
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Gross_profit map[string]interface{} `json:"gross_profit,omitempty"`
	Net_income map[string]interface{} `json:"net_income,omitempty"`
	Customer_id string `json:"customer_id,omitempty"` // Customer id
	Expenses map[string]interface{} `json:"expenses"`
	Net_operating_income map[string]interface{} `json:"net_operating_income,omitempty"`
	Start_date string `json:"start_date,omitempty"` // The start date of the report
}

// Payment represents the Payment schema from the OpenAPI specification
type Payment struct {
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
	Customer LinkedCustomer `json:"customer,omitempty"` // The customer this entity is linked to.
	Total_amount float64 `json:"total_amount"` // Amount of payment
	Note string `json:"note,omitempty"` // Optional note to be associated with the payment.
	Accounts_receivable_account_id string `json:"accounts_receivable_account_id,omitempty"` // Unique identifier for the account to allocate payment to.
	Transaction_date string `json:"transaction_date"` // Date transaction was entered - YYYY:MM::DDThh:mm:ss.sTZD
	Allocations []map[string]interface{} `json:"allocations,omitempty"`
	Payment_method_reference string `json:"payment_method_reference,omitempty"` // Optional reference message returned by payment method on processing
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Account LinkedLedgerAccount `json:"account,omitempty"`
	Reconciled bool `json:"reconciled,omitempty"` // Payment has been reconciled
	TypeField string `json:"type,omitempty"` // Type of payment
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Reference string `json:"reference,omitempty"` // Optional payment reference message ie: Debit remittance detail.
	Display_id string `json:"display_id,omitempty"` // Payment id to be displayed.
	Id string `json:"id"` // Unique identifier representing the entity
	Payment_method_id string `json:"payment_method_id,omitempty"` // Unique identifier for the payment method.
	Accounts_receivable_account_type string `json:"accounts_receivable_account_type,omitempty"` // Type of accounts receivable account.
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Payment_method string `json:"payment_method,omitempty"` // Payment method name
	Status string `json:"status,omitempty"` // Status of payment
	Supplier LinkedSupplier `json:"supplier,omitempty"` // The supplier this entity is linked to.
}

// AccountingWebhookEvent represents the AccountingWebhookEvent schema from the OpenAPI specification
type AccountingWebhookEvent struct {
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Event_type string `json:"event_type,omitempty"`
}

// DeleteInvoiceItemResponse represents the DeleteInvoiceItemResponse schema from the OpenAPI specification
type DeleteInvoiceItemResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetJournalEntriesResponse represents the GetJournalEntriesResponse schema from the OpenAPI specification
type GetJournalEntriesResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []JournalEntry `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetPurchaseOrdersResponse represents the GetPurchaseOrdersResponse schema from the OpenAPI specification
type GetPurchaseOrdersResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []PurchaseOrder `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// CreateInvoiceResponse represents the CreateInvoiceResponse schema from the OpenAPI specification
type CreateInvoiceResponse struct {
	Data InvoiceResponse `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// InvoiceItemExpenseAccount represents the InvoiceItemExpenseAccount schema from the OpenAPI specification
type InvoiceItemExpenseAccount struct {
	Code string `json:"code,omitempty"` // The code assigned to the account.
	Id string `json:"id,omitempty"` // The unique identifier for the account.
	Name string `json:"name,omitempty"` // The name of the account.
	Nominal_code string `json:"nominal_code,omitempty"` // The nominal code of the account.
}

// GetInvoiceResponse represents the GetInvoiceResponse schema from the OpenAPI specification
type GetInvoiceResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Invoice `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// LinkedCustomer represents the LinkedCustomer schema from the OpenAPI specification
type LinkedCustomer struct {
	Company_name string `json:"company_name,omitempty"` // The company name of the customer.
	Display_id string `json:"display_id,omitempty"` // The display ID of the customer.
	Display_name string `json:"display_name,omitempty"` // The display name of the customer.
	Id string `json:"id"` // The ID of the customer this entity is linked to.
	Name string `json:"name,omitempty"` // The name of the customer. Deprecated, use display_name instead.
}

// CreateCreditNoteResponse represents the CreateCreditNoteResponse schema from the OpenAPI specification
type CreateCreditNoteResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetCompanyInfoResponse represents the GetCompanyInfoResponse schema from the OpenAPI specification
type GetCompanyInfoResponse struct {
	Data CompanyInfo `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// InvoiceItemIncomeAccount represents the InvoiceItemIncomeAccount schema from the OpenAPI specification
type InvoiceItemIncomeAccount struct {
	Code string `json:"code,omitempty"` // The code assigned to the account.
	Id string `json:"id,omitempty"` // The unique identifier for the account.
	Name string `json:"name,omitempty"` // The name of the account.
	Nominal_code string `json:"nominal_code,omitempty"` // The nominal code of the account.
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	TypeField string `json:"type,omitempty"` // The type of phone number
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
	Number string `json:"number"` // The phone number
}

// Supplier represents the Supplier schema from the OpenAPI specification
type Supplier struct {
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Suffix string `json:"suffix,omitempty"`
	Bank_accounts []BankAccount `json:"bank_accounts,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Addresses []Address `json:"addresses,omitempty"`
	Account LinkedLedgerAccount `json:"account,omitempty"`
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Individual bool `json:"individual,omitempty"` // Is this an individual or business supplier
	Company_name string `json:"company_name,omitempty"` // The name of the company.
	Websites []Website `json:"websites,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Notes string `json:"notes,omitempty"` // Some notes about this supplier
	Middle_name string `json:"middle_name,omitempty"` // Middle name of the person.
	Display_name string `json:"display_name,omitempty"` // Display name
	Id string `json:"id"` // A unique identifier for an object.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Title string `json:"title,omitempty"` // The job title of the person.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Display_id string `json:"display_id,omitempty"` // Display ID
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Status string `json:"status,omitempty"` // Supplier status
	Tax_number string `json:"tax_number,omitempty"`
}

// GetLedgerAccountsResponse represents the GetLedgerAccountsResponse schema from the OpenAPI specification
type GetLedgerAccountsResponse struct {
	Data []LedgerAccount `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetInvoiceItemResponse represents the GetInvoiceItemResponse schema from the OpenAPI specification
type GetInvoiceItemResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data InvoiceItem `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
}

// Bill represents the Bill schema from the OpenAPI specification
type Bill struct {
	Sub_total float64 `json:"sub_total,omitempty"` // Sub-total amount, normally before tax.
	Total_tax float64 `json:"total_tax,omitempty"` // Total tax amount applied to this bill.
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Reference string `json:"reference,omitempty"` // Optional bill reference.
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Tax_code string `json:"tax_code,omitempty"` // Applicable tax id/code override if tax is not supplied on a line item basis.
	Deposit float64 `json:"deposit,omitempty"` // Amount of deposit made to this bill.
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	Discount_percentage float64 `json:"discount_percentage,omitempty"` // Discount percentage applied to this transaction.
	Paid_date string `json:"paid_date,omitempty"` // The paid date is the date on which a payment was sent to the supplier - YYYY-MM-DD.
	Line_items []BillLineItem `json:"line_items,omitempty"`
	Supplier LinkedSupplier `json:"supplier,omitempty"` // The supplier this entity is linked to.
	Status string `json:"status,omitempty"` // Invoice status
	Bill_number string `json:"bill_number,omitempty"` // Reference to supplier bill number
	Notes string `json:"notes,omitempty"`
	Tax_inclusive bool `json:"tax_inclusive,omitempty"` // Amounts are including tax
	Terms string `json:"terms,omitempty"` // Terms of payment.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Accounting_by_row bool `json:"accounting_by_row,omitempty"` // Indicates if accounting by row is used (true) or not (false). Accounting by row means that a separate ledger transaction is created for each row.
	Due_date string `json:"due_date,omitempty"` // The due date is the date on which a payment is scheduled to be received - YYYY-MM-DD.
	Po_number string `json:"po_number,omitempty"` // A PO Number uniquely identifies a purchase order and is generally defined by the buyer. The buyer will match the PO number in the invoice to the Purchase Order.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Bank_account BankAccount `json:"bank_account,omitempty"`
	Balance float64 `json:"balance,omitempty"` // Balance of bill due.
	Bill_date string `json:"bill_date,omitempty"` // Date bill was issued - YYYY-MM-DD.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Total float64 `json:"total,omitempty"` // Total amount of bill, including tax.
	Ledger_account LinkedLedgerAccount `json:"ledger_account,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Currency_rate float64 `json:"currency_rate,omitempty"` // Currency Exchange Rate at the time entity was recorded/generated.
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// AccountingCustomer represents the AccountingCustomer schema from the OpenAPI specification
type AccountingCustomer struct {
	Downstream_id string `json:"downstream_id,omitempty"` // The third-party API ID of original entity
	Bank_accounts []BankAccount `json:"bank_accounts,omitempty"`
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Title string `json:"title,omitempty"` // The job title of the person.
	Addresses []Address `json:"addresses,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Websites []Website `json:"websites,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Display_name string `json:"display_name,omitempty"` // Display name
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Tax_rate LinkedTaxRate `json:"tax_rate,omitempty"`
	Middle_name string `json:"middle_name,omitempty"` // Middle name of the person.
	Account LinkedLedgerAccount `json:"account,omitempty"`
	Status string `json:"status,omitempty"` // Customer status
	Notes string `json:"notes,omitempty"` // Some notes about this customer
	Company_name string `json:"company_name,omitempty"` // The name of the company.
	Parent LinkedParentCustomer `json:"parent,omitempty"` // The parent customer this entity is linked to.
	Payment_method string `json:"payment_method,omitempty"` // Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	Tax_number string `json:"tax_number,omitempty"`
	Individual bool `json:"individual,omitempty"` // Is this an individual or business customer
	Display_id string `json:"display_id,omitempty"` // Display ID
	Project bool `json:"project,omitempty"` // If true, indicates this is a Project.
	Channel string `json:"channel,omitempty"` // The channel through which the transaction is processed.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id"` // A unique identifier for an object.
	Emails []Email `json:"emails,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// GetTaxRatesResponse represents the GetTaxRatesResponse schema from the OpenAPI specification
type GetTaxRatesResponse struct {
	Data []TaxRate `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetBillsResponse represents the GetBillsResponse schema from the OpenAPI specification
type GetBillsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Bill `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// BalanceSheetFilter represents the BalanceSheetFilter schema from the OpenAPI specification
type BalanceSheetFilter struct {
	End_date string `json:"end_date,omitempty"` // Filter by end date. If end date is given, start date is required.
	Start_date string `json:"start_date,omitempty"` // Filter by start date. If start date is given, end date is required.
}

// DeletePurchaseOrderResponse represents the DeletePurchaseOrderResponse schema from the OpenAPI specification
type DeletePurchaseOrderResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetCreditNoteResponse represents the GetCreditNoteResponse schema from the OpenAPI specification
type GetCreditNoteResponse struct {
	Data CreditNote `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetSuppliersResponse represents the GetSuppliersResponse schema from the OpenAPI specification
type GetSuppliersResponse struct {
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Supplier `json:"data"`
}

// GetTaxRateResponse represents the GetTaxRateResponse schema from the OpenAPI specification
type GetTaxRateResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data TaxRate `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetInvoiceItemsResponse represents the GetInvoiceItemsResponse schema from the OpenAPI specification
type GetInvoiceItemsResponse struct {
	Data []InvoiceItem `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// DeleteTaxRateResponse represents the DeleteTaxRateResponse schema from the OpenAPI specification
type DeleteTaxRateResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// CreateCustomerResponse represents the CreateCustomerResponse schema from the OpenAPI specification
type CreateCustomerResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// ProfitAndLossRecord represents the ProfitAndLossRecord schema from the OpenAPI specification
type ProfitAndLossRecord struct {
	TypeField string `json:"type"`
	Value float64 `json:"value,omitempty"`
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// UpdateSupplierResponse represents the UpdateSupplierResponse schema from the OpenAPI specification
type UpdateSupplierResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// CreateInvoiceItemResponse represents the CreateInvoiceItemResponse schema from the OpenAPI specification
type CreateInvoiceItemResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetBalanceSheetResponse represents the GetBalanceSheetResponse schema from the OpenAPI specification
type GetBalanceSheetResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data BalanceSheet `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}
