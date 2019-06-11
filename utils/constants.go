package utils

// Constants
const (
	// Error codes
	CODEALLAOK              string = "000" // Success
	CODENOTFOUND            string = "101" // resource not found
	CODEUNKNOWNINVOKE       string = "102" // Unknown invoke
	CODEUNPROCESSABLEENTITY string = "103" // Invalid input
	CODEGENEXCEPTION        string = "201" // Unknown exception
	CODEAlRDEXIST           string = "202" // Not unique
	CODENOTALLWD            string = "104" // Operation not allowed

	//Asset Object/Doctypes
	BANKOO string = "BANKOO" //bitmask is 00
	BNKBRA string = "BNKBRA" //bitmask is 01
	BNKCFG string = "BNKCFG" //bitmask is 02
	BUYERO string = "BUYERO" //bitmask is 03
	LOANOO string = "LOANOO" //bitmask is 04
	LONBUY string = "LONBUY" //bitmask is 05
	LONDOC string = "LONDOC" //bitmask is 06
	LONMRK string = "LONMRK" //bitmask is 07
	LONRAT string = "LONRAT" //bitmask is 08
	MESSAG string = "MESSAG" //bitmask is 09
	PERMSN string = "PERMSN" //bitmask is 10
	PRPRTY string = "PRPRTY" //bitmask is 11
	ROLEOO string = "ROLEOO" //bitmask is 12
	SELLER string = "SELLER" //bitmask is 13
	TRNSAC string = "TRNSAC" //bitmask is 14
	USEROO string = "USEROO" //bitmask is 15
	USECAT string = "USECAT" //bitmask is 16

	// Range index name - to perform range queries
	INDXNM string = "bitmask~txnID~id"

	FIXEDPT int32 = 4 // All currency values rounded off to 4 decimals i.e. 0.0000

	//function names for read and write

	Init string = "Init"

	BankR string = "BankR"
	BankW string = "BankW"

	BankBranchR string = "BankBranchR"
	BankBranchW string = "BankBranchW"

	BranchConfigR string = "BranchConfigR"
	BranchConfigW string = "BranchConfigW"

	BuyerR string = "BuyerR"
	BuyerW string = "BuyerW"

	LoanR string = "LoanR"
	LoanW string = "LoanW"

	LoanBuyerR string = "LoanBuyerR"
	LoanBuyerW string = "LoanBuyerW"

	LoanDocR string = "LoanDocR"
	LoanDocW string = "LoanDocW"

	LoanMarketShareR string = "LoanMarketShareR"
	LoanMarketShareW string = "LoanMarketShareW"

	LoanRatingR string = "LoanRatingR"
	LoanRatingW string = "LoanRatingW"

	MessageR string = "MessageR"
	MessageW string = "MessageW"

	PermissionsR string = "PermissionsR"
	PermissionsW string = "PermissionsW"

	PropertyR string = "PropertyR"
	PropertyW string = "PropertyW"

	RoleR string = "RoleR"
	RoleW string = "RoleW"

	SellerR string = "SellerR"
	SellerW string = "SellerW"

	TransactionR string = "TransactionR"
	TransactionW string = "TransactionW"

	UserR string = "UserR"
	UserW string = "UserW"

	UserCategoryR string = "UserCategoryR"
	UserCategoryW string = "UserCategoryW"
)
