package task

import "time"

const (
	TWOCAP_KEY    = ""
	RECAP_SITEKEY = "6LfFeqMfAAAAAPTHzfiw74_7tDqsXa2Qq4UgZnQJ"

	ITEM_OOS         = "ItemOutOfStock"
	INVALID_EXCHANGE = "InvalidExchange"
)

type AuthLogin struct {
	AccessToken   string `json:"access_token"`
	APIServerURL  string `json:"api_server_url"`
	AppID         string `json:"app_id"`
	UserID        string `json:"user_id"`
	UniqueAppID   string `json:"unique_app_id"`
	SignedAt      int64  `json:"signed_at"`
	KeyVersion    string `json:"key_version"`
	UniqueUserID  string `json:"unique_user_id"`
	EcomSignature string `json:"ecom_signature"`
	AppVersion    string `json:"app_version"`
}

type JwtResp struct {
	Jwt      string `json:"jwt"`
	UserInfo struct {
		Firstname    string `json:"firstname"`
		Lastname     string `json:"lastname"`
		EmailAddress string `json:"email_address"`
		PhoneNumber  string `json:"phone_number"`
		Role         string `json:"role"`
	} `json:"user_info"`
	JwtExp    float64 `json:"jwt_exp"`
	StoreInfo struct {
		StoreID       string `json:"store_id"`
		StoreDispName string `json:"store_disp_name"`
		EcomStoreType string `json:"ecom_store_type"`
		StoreSegment  string `json:"store_segment"`
		ImageLogoURL  string `json:"image_logo_url"`
		LegacyPlanID  string `json:"legacy_plan_id"`
	} `json:"store_info"`
	EppVerified bool `json:"epp_verified"`
	AuthInfo    struct {
		AuthType  string `json:"auth_type"`
		AuthValue string `json:"auth_value"`
	} `json:"auth_info"`
}

type SamsungUser struct {
	UserInfo struct {
		Firstname    string `json:"firstname"`
		Lastname     string `json:"lastname"`
		EmailAddress string `json:"email_address"`
		PhoneNumber  string `json:"phone_number"`
		UserGroups   []any  `json:"user_groups"`
		CountryCode  string `json:"countryCode"`
	} `json:"user_info"`
	UserIdentity struct {
		IDProvider string `json:"id_provider"`
		HqGUID     string `json:"hq_guid"`
		UserID     string `json:"user_id"`
		SmbStatus  string `json:"smb_status"`
	} `json:"user_identity"`
	LoginType string `json:"login_type"`
	StoreID   string `json:"store_id"`
	StoreInfo struct {
		StoreID       string `json:"store_id"`
		StoreDispName string `json:"store_disp_name"`
		EcomStoreType string `json:"ecom_store_type"`
		StoreSegment  string `json:"store_segment"`
		ImageLogoURL  string `json:"image_logo_url"`
		LegacyPlanID  string `json:"legacy_plan_id"`
	} `json:"store_info"`
	EppVerified bool `json:"epp_verified"`
	AuthInfo    struct {
		AuthType  string `json:"auth_type"`
		AuthValue string `json:"auth_value"`
	} `json:"auth_info"`
}

type CsrfResp struct {
	Csrf struct {
		Token         string `json:"token"`
		ParameterName string `json:"parameterName"`
		HeaderName    string `json:"headerName"`
	} `json:"_csrf"`
	RemIDChkYN      bool   `json:"remIdChkYN"`
	BackgroundImage string `json:"backgroundImage"`
	WipCancelURI    string `json:"wipCancelURI"`
}

type AppCreateCart struct {
	CartID      string   `json:"cart_id,omitempty"`
	TriggerTags []string `json:"trigger_tags,omitempty"`
	Timezone    string   `json:"timezone"`
	StoreID     string   `json:"store_id"`
	Experiments []Exp    `json:"experiments"`
}

type CaptchaPayload struct {
	Token    string `json:"token"`
	Action   string `json:"action"`
	Username string `json:"username"`
}

type EmailPayload struct {
	LoginID    string `json:"loginId"`
	RememberID bool   `json:"rememberId"`
	StaySignIn bool   `json:"staySignIn"`
}

type EmailResp struct {
	RtnCd   string `json:"rtnCd"`
	NextURL string `json:"nextURL"`
	RtnMsg  string `json:"rtnMsg"`
}

type LoginPayload struct {
	Email            string `json:"iptLgnID"`
	Password         string `json:"iptLgnPD"`
	LoginKey         string `json:"svcIptLgnKY"`
	LoginIV          string `json:"svcIptLgnIV"`
	RememberPassword bool   `json:"remIdChkYN"`
	Captcha          any    `json:"captchaAnswer"`
	Assertion        any    `json:"assertion"`
}

type CartPayload struct {
	CartId      string     `json:"cart_id,omitempty"`
	LineItems   []CartItem `json:"line_items"`
	PostalCode  string     `json:"postal_code"`
	StoreId     string     `json:"store_id,omitempty"`
	Experiments []Exp      `json:"experiments"`
}

type CartItem struct {
	SkuID      string `json:"sku_id"`
	Quantity   string `json:"quantity"`
	ExchangeID string `json:"exchange_id,omitempty"`
	LineItems  []Item `json:"line_items,omitempty"`
}

type Item struct {
	SkuID     string `json:"sku_id"`
	Quantity  string `json:"quantity"`
	FlashSale bool   `json:"flash_sale,omitempty"`
}

type Exp struct {
	SignifydPreAuthActive bool `json:"signifyd_pre_auth_active,omitempty"`
	CartSpa               bool `json:"cart_spa,omitempty"`
}

type Tracking struct {
	AdobeVisitorID string `json:"adobe_visitor_id"`
}

type OrderPayload struct {
	PaymentMethod           string     `json:"payment_method"`
	EncryptedPaymentContext EncPayment `json:"encrypted_payment_context"`
}

type BillingPayload struct {
	PaymentMethod           string      `json:"payment_method"`
	EncryptedPaymentContext EncPayment  `json:"encrypted_payment_context"`
	BillingInfo             BillingInfo `json:"billing_info"`
	Token                   string      `json:"token"`
	CaptchaProvider         string      `json:"captcha_provider"`
}

type EncPayment struct {
	EncryptedPayload  string `json:"encrypted_payload"`
	EncryptedPassword string `json:"encrypted_password"`
}

type BillingInfo struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Line1      string `json:"line_1"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	City       string `json:"city"`
}

// Define the structs based on the JSON structure
type UserAlerts struct {
	CartAlerts      []interface{} `json:"cart_alerts"`
	LineItemsAlerts []interface{} `json:"line_items_alerts"`
}

type DeviceAttributes struct {
	ClientType string `json:"client_type"`
	OS         string `json:"os"`
	Platform   string `json:"platform"`
	Browser    string `json:"browser"`
	DeviceType string `json:"device_type"`
}

type UserInfo struct {
	Locale                     string           `json:"locale"`
	IdentityID                 *string          `json:"identity_id"`
	VisitorID                  string           `json:"visitor_id"`
	UserID                     *string          `json:"user_id"`
	DeviceAttributes           DeviceAttributes `json:"device_attributes"`
	StoreType                  string           `json:"store_type"`
	VATReversalThresholdAmount float64          `json:"vat_reversal_threshold_amount"`
}

type Discount struct {
	Amount float64       `json:"amount"`
	Split  []interface{} `json:"split"`
}

type QuantityGroup struct {
	AssociatedQuantities []interface{} `json:"associated_quantities"`
	Tax                  float64       `json:"tax"`
	Shipping             float64       `json:"shipping"`
	Discount             Discount      `json:"discount"`
	Subtotal             float64       `json:"subtotal"`
}

type SalePriceInfo struct {
	Public struct {
		Unfinanced float64 `json:"unfinanced"`
	} `json:"public"`
	Store struct {
		Unfinanced float64 `json:"unfinanced"`
	} `json:"store"`
}

type LineItemCost struct {
	UnitListPrice  float64         `json:"unit_list_price"`
	SalePrice      float64         `json:"sale_price"`
	Total          float64         `json:"total"`
	QuantityGroups []QuantityGroup `json:"quantity_groups"`
	RegularPrice   float64         `json:"regular_price"`
	SalePriceInfo  SalePriceInfo   `json:"sale_price_info"`
	Subtotal       float64         `json:"subtotal"`
	DisplayPrice   float64         `json:"display_price"`
	UnitPrice      float64         `json:"unit_price"`
}

type Attributes struct {
	B2BEcomFlag                bool          `json:"b2bEcomFlag"`
	EcomFlag                   bool          `json:"ecom_flag"`
	DisplayName                string        `json:"display_name"`
	ShortDescription           string        `json:"short_description"`
	LongDescription            string        `json:"long_description"`
	ThumbnailURL               string        `json:"thumbnail_url"`
	ImageURL                   string        `json:"image_url"`
	IsAccessory                bool          `json:"is_accessory"`
	ProductDetailsPageURL      *string       `json:"product_details_page_url"`
	ProductDivision            *string       `json:"product_division"`
	ProductFamily              *string       `json:"product_family"`
	ProductType                string        `json:"product_type"`
	ShipAlone                  bool          `json:"ship_alone"`
	ShipAlongWith              []interface{} `json:"ship_along_with"`
	ShipAlongWithQuantity      int           `json:"ship_along_with_quantity"`
	ProductTypeCustom          *string       `json:"product_type_custom"`
	DiagonalScreenSize         *float64      `json:"diagonal_screen_size"`
	ShippingWeight             float64       `json:"shipping_weight"`
	TermsAndConditionsID       *string       `json:"terms_and_conditions_id"`
	ShortRedemptionInstruction *string       `json:"short_redemption_instruction"`
	Channel                    []interface{} `json:"channel"`
	ModelName                  string        `json:"model_name"`
	Options                    interface{}   `json:"options"`
	Requires                   []interface{} `json:"requires"`
	Supports                   []interface{} `json:"supports"`
	CarrierDeliveryModeID      *string       `json:"carrier_delivery_mode_id"`
	IsUpgradeEligible          bool          `json:"is_upgrade_eligible"`
	DeliveryModeType           *string       `json:"delivery_mode_type"`
	DeliveryServiceType        *string       `json:"delivery_service_type"`
	DeliveryIndicator          *string       `json:"delivery_indicator"`
	DeliveryInstruction        *string       `json:"delivery_instruction"`
	ExternalMerchantID         *string       `json:"external_merchant_id"`
	IsFlagshipSKU              bool          `json:"is_flagship_sku"`
	DelayedShippingMessage     *string       `json:"delayed_shipping_message"`
	SamsungCareProductID       *string       `json:"samsung_care_product_id"`
	PartnerAttributesConfig    *string       `json:"partner_attributes_config"`
	Carrier                    *string       `json:"carrier"`
	IsBuyFromStoreEligible     bool          `json:"is_buy_from_store_eligible"`
	B2CEcomEnabled             bool          `json:"b2c_ecom_enabled"`
	B2BEcomEnabled             bool          `json:"b2b_ecom_enabled"`
	IsHADelivery               bool          `json:"is_ha_delivery"`
	IsInsuranceProduct         bool          `json:"is_insurance_product"`
	IsSimCard                  bool          `json:"is_sim_card"`
	IsTradeInEligible          bool          `json:"is_trade_in_eligible"`
	IsSubscription             bool          `json:"is_subscription"`
	IsHazmat                   bool          `json:"is_hazmat"`
	IsCustomizable             bool          `json:"is_customizable"`
	IsBuyNowEligible           bool          `json:"is_buynow_eligible"`
	ApplicablePaymentOptions   []interface{} `json:"applicable_payment_options"`
	IsEnergyStarCertified      *bool         `json:"is_energy_star_certified"`
	ModelDivisionCode          *string       `json:"model_division_code"`
	CisSKU                     *string       `json:"cis_sku"`
	IsCancellable              bool          `json:"is_cancellable"`
	DaasServiceSKUs            []interface{} `json:"daas_service_skus"`
	IsGRVSKU                   bool          `json:"is_grv_sku"`
	SalesOrg                   *string       `json:"sales_org"`
	RenewalSKUID               *string       `json:"renewal_sku_id"`
	UseCase                    *string       `json:"use_case"`
	TrialInfo                  struct {
		IsApplicable bool `json:"is_applicable"`
	} `json:"trial_info"`
	PlanTierValue              float64       `json:"plan_tier_value"`
	PartnerContractMarginValue float64       `json:"partner_contract_margin_value"`
	PartnerContractMarginUnit  string        `json:"partner_contract_margin_unit"`
	AllowBundleCancel          bool          `json:"allow_bundle_cancel"`
	PartnerAdminPortalURL      *string       `json:"partner_admin_portal_url"`
	PartnerSupportNumber       *string       `json:"partner_support_number"`
	AssociatedSKUs             []interface{} `json:"associated_skus"`
	Tags                       []string      `json:"tags"`
	Bulletin                   []interface{} `json:"bulletin"`
	IsESimEnabled              bool          `json:"is_esim_enabled"`
	SimTypeForActivation       string        `json:"sim_type_for_activation"`
	BlackListedStates          []interface{} `json:"black_listed_states"`
	LocalizedFields            []interface{} `json:"localized_fields"`
	VATReversal                bool          `json:"vat_reversal"`
}

type LineItem struct {
	Quantity           int                `json:"quantity"`
	ExchangeAttributes ExchangeAttributes `json:"exchange_attributes"`
	SKUID              string             `json:"sku_id"`
	ExchangeID         *string            `json:"exchange_id"`
	LineItemID         string             `json:"line_item_id"`
	LineItemCost       LineItemCost       `json:"line_item_cost"`
	OriginalSKUID      string             `json:"original_sku_id"`
	UIOrder            int                `json:"ui_order"`
	PricingOrder       int                `json:"pricing_order"`
	Attributes         Attributes         `json:"attributes"`
	SkipInventory      bool               `json:"skip_inventory"`
	FlashSale          bool               `json:"flash_sale"`
	InventoryStatus    struct {
		Status string `json:"status"`
	} `json:"inventory_status"`
}

type ExchangeAttributes struct {
	MultiQuantityDevices []struct {
		ID         string   `json:"id"`
		SKU        string   `json:"sku"`
		Type       string   `json:"type"`
		OfferID    string   `json:"offer_id"`
		Quantity   int      `json:"quantity"`
		DeviceID   string   `json:"device_id"`
		OfferIDs   []string `json:"offer_ids"`
		DeviceInfo struct {
			Name  string `json:"name"`
			Brand string `json:"brand"`
			State struct {
				DataWiped     bool `json:"data_wiped"`
				DeviceWorking bool `json:"device_working"`
				ScreenWorking bool `json:"screen_working"`
			} `json:"state"`
			Category  string `json:"category"`
			ImageURL  string `json:"image_url"`
			Valuation struct {
				IsExternalDeviceValuation bool `json:"is_external_device_valuation"`
			} `json:"valuation"`
			ExternalAttributes struct {
				Grade string `json:"grade"`
			} `json:"external_attributes"`
		} `json:"device_info"`
		InitiatedDate     string `json:"initiated_date"`
		EstimatedDiscount struct {
			Total struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"total"`
			ExchangeValue struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"exchange_value"`
			ExchangeDiscount struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"exchange_discount"`
		} `json:"estimated_discount"`
		ExternalAttributes struct {
			Reference string `json:"reference"`
		} `json:"external_attributes"`
		DeviceGrade string `json:"device_grade"`
	} `json:"multi_quantity_devices"`
	DeviceInfo struct {
		EstimatedDiscount struct {
			Total struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"total"`
			ExchangeValue struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"exchange_value"`
			ExchangeDiscount struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"exchange_discount"`
		} `json:"estimated_discount"`
	} `json:"device_info"`
	Type string `json:"type"`
}

type Cost struct {
	Discount struct {
		Amount float64       `json:"amount"`
		Split  []interface{} `json:"split"`
	} `json:"discount"`
	Shipping           float64 `json:"shipping"`
	Subtotal           float64 `json:"subtotal"`
	Tax                float64 `json:"tax"`
	Total              float64 `json:"total"`
	DisplayTax         float64 `json:"display_tax"`
	RemainingDueAmount float64 `json:"remaining_due_amount"`
}

type ChannelInfo struct {
	ChannelCode string `json:"channel_code"`
}

type AffiliateTrackingInfo struct {
	AdobeVisitorID string `json:"adobe_visitor_id"`
}

type MetaInformation struct {
	Experiments []struct {
		RemoveTaxBloat        bool `json:"remove_tax_bloat"`
		SignifydPreAuthActive bool `json:"signifyd_pre_auth_active"`
		CartSpa               bool `json:"cart_spa"`
	} `json:"experiments"`
	Tags []string `json:"tags"`
}

type RetryFlags struct {
	PaymentOptions bool `json:"payment_options"`
	DeliveryModes  bool `json:"delivery_modes"`
}

type OfferGroup struct {
	GroupID    string `json:"group_id"`
	Discounted []struct {
		LineItemID           string `json:"line_item_id"`
		AssociatedQuantities []int  `json:"associated_quantities"`
	} `json:"discounted"`
	OfferID   string `json:"offer_id"`
	Triggered []struct {
		LineItemID           string `json:"line_item_id"`
		AssociatedQuantities []int  `json:"associated_quantities"`
	} `json:"triggered"`
	DeviceID             string `json:"device_id"`
	DeviceGrade          string `json:"device_grade"`
	AssociatedQuantities []int  `json:"associated_quantities"`
	ChildExchangeID      string `json:"child_exchange_id"`
}

type Rewards struct {
	Accrued []struct {
		RewardPoints     int     `json:"reward_points"`
		RewardPercentage float64 `json:"reward_percentage"`
		Type             string  `json:"type"`
		Split            []struct {
			OfferID      string `json:"offer_id"`
			OfferName    string `json:"offer_name"`
			RewardPoints int    `json:"reward_points"`
		} `json:"split"`
	} `json:"accrued"`
	DollarConversionMultiplier float64 `json:"dollar_conversion_multiplier"`
}

type PaymentAttributes struct {
	CostSplitType    string `json:"cost_split_type"`
	CostSplitVersion string `json:"cost_split_version"`
}

type LineItemFinancePlan struct {
	LineItemID  string `json:"line_item_id"`
	SKUID       string `json:"sku_id"`
	FinanceInfo struct {
		FinancePlans []struct {
			PlanName               string  `json:"plan_name"`
			PlanID                 string  `json:"plan_id"`
			UpgradePlanID          int     `json:"upgrade_plan_id"`
			PlanType               string  `json:"plan_type"`
			MonthlyRateCoefficient float64 `json:"monthly_rate_coefficient"`
			FinanceProviderType    string  `json:"finance_provider_type"`
			ExternalPlanID         int     `json:"external_plan_id"`
			Tenure                 struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"tenure"`
			InterestRate        string  `json:"interest_rate"`
			InstallmentAmount   string  `json:"installment_amount"`
			TotalAmount         string  `json:"total_amount"`
			MinimumAmount       float64 `json:"minimum_amount"`
			DownPayment         string  `json:"down_payment"`
			FinancingAttributes struct {
				IsUpgradeEligible bool `json:"is_upgrade_eligible"`
			} `json:"financing_attributes"`
			MonthlyPayment      float64 `json:"monthly_payment"`
			FirstMonthlyPayment float64 `json:"first_monthly_payment"`
		} `json:"financePlans"`
		FinancingOptions []struct {
			SKUID                  string        `json:"sku_id"`
			PlanID                 string        `json:"plan_id"`
			FinancingEligibleSKUs  []interface{} `json:"financing_eligible_skus"`
			RelatedProductFamilies []interface{} `json:"related_product_families"`
		} `json:"financingOptions"`
	} `json:"finance_info"`
}

type Response struct {
	UserAlerts         UserAlerts          `json:"user_alerts"`
	SchemaVersion      string              `json:"schema_version"`
	CartID             string              `json:"cart_id"`
	UserInfo           UserInfo            `json:"user_info"`
	Status             string              `json:"status"`
	Currency           string              `json:"currency"`
	Type               string              `json:"type"`
	LineItems          map[string]LineItem `json:"line_items"`
	ExchangeAttributes ExchangeAttributes  `json:"exchange_attributes"`
	OriginalSKUID      string              `json:"original_sku_id"`
	UIOrder            int                 `json:"ui_order"`
	PricingOrder       int                 `json:"pricing_order"`
	Attributes         Attributes          `json:"attributes"`
	SkipInventory      bool                `json:"skip_inventory"`
	FlashSale          bool                `json:"flash_sale"`
	InventoryStatus    struct {
		Status string `json:"status"`
	} `json:"inventory_status"`
	OffersApplied map[string]struct {
		StartDate              string        `json:"start_date"`
		EndDate                string        `json:"end_date"`
		Name                   string        `json:"name"`
		OfferID                string        `json:"offer_id"`
		Type                   string        `json:"type"`
		IsDelayedGratification bool          `json:"is_delayed_gratification"`
		SalesPitch             []interface{} `json:"sales_pitch"`
		FundingInfo            []struct {
			MdfID string `json:"mdf_id"`
		} `json:"funding_info"`
	} `json:"offers_applied"`
	CouponCodes               []interface{}         `json:"coupon_codes"`
	CouponCodesFailureCount   int                   `json:"coupon_codes_failure_count"`
	Cost                      Cost                  `json:"cost"`
	ChannelInfo               ChannelInfo           `json:"channel_info"`
	MultiQuantity             bool                  `json:"multi_quantity"`
	IsStrictZeroDiscount      bool                  `json:"is_strict_zero_discount"`
	PoID                      string                `json:"po_id"`
	PostalCode                string                `json:"postal_code"`
	AffiliateTrackingInfo     AffiliateTrackingInfo `json:"affiliate_tracking_info"`
	MetaInformation           MetaInformation       `json:"meta_information"`
	RetryFlags                RetryFlags            `json:"retry_flags"`
	Locale                    string                `json:"locale"`
	IsStorePickupEligible     bool                  `json:"is_store_pickup_eligible"`
	IsExchangeTransformed     bool                  `json:"is_exchange_transformed"`
	RewardsRedemptionEligible bool                  `json:"rewards_redemption_eligible"`
	RejectedCouponCodes       []interface{}         `json:"rejected_coupon_codes"`
	OfferGroups               []OfferGroup          `json:"offer_groups"`
	Rewards                   Rewards               `json:"rewards"`
	// PricingValidationChecksum string                `json:"pricing_validation_checksum"`
	IsLoginRequired        bool                  `json:"is_login_required"`
	IsEPPAuthRequired      bool                  `json:"is_epp_auth_required"`
	LineItemGroups         map[string][]string   `json:"line_item_groups"`
	PaymentAttributes      PaymentAttributes     `json:"payment_attributes"`
	LineItemFinancePlans   []LineItemFinancePlan `json:"line_item_finance_plans"`
	ApplicableFinancePlans []struct {
		PlanName               string  `json:"plan_name"`
		PlanID                 string  `json:"plan_id"`
		UpgradePlanID          int     `json:"upgrade_plan_id"`
		PlanType               string  `json:"plan_type"`
		MonthlyRateCoefficient float64 `json:"monthly_rate_coefficient"`
		FinanceProviderType    string  `json:"finance_provider_type"`
		ExternalPlanID         int     `json:"external_plan_id"`
		Tenure                 struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"tenure"`
		InterestRate        string  `json:"interest_rate"`
		InstallmentAmount   string  `json:"installment_amount"`
		TotalAmount         string  `json:"total_amount"`
		MinimumAmount       float64 `json:"minimum_amount"`
		DownPayment         string  `json:"down_payment"`
		FinancingAttributes struct {
			IsUpgradeEligible bool `json:"is_upgrade_eligible"`
		} `json:"financing_attributes"`
		MonthlyPayment      float64 `json:"monthly_payment"`
		FirstMonthlyPayment float64 `json:"first_monthly_payment"`
	} `json:"applicable_finance_plans"`
	IsSupCart     bool          `json:"is_sup_cart"`
	ReferralCodes []interface{} `json:"referral_codes"`
	ExpiryDate    string        `json:"expiry_date"`
}

type JWTPayload struct {
	AccessToken   string `json:"access_token"`
	APIServerURL  string `json:"api_server_url"`
	AppID         string `json:"app_id"`
	UserID        string `json:"user_id"`
	UniqueAppID   string `json:"unique_app_id"`
	SignedAt      int64  `json:"signed_at"`
	KeyVersion    string `json:"key_version"`
	UniqueUserID  string `json:"unique_user_id"`
	EcomSignature string `json:"ecom_signature"`
	AppVersion    string `json:"app_version"`
}

type ShippingPayload struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	City             string `json:"city"`
	PostalCode       string `json:"postal_code"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	AddressSave      bool   `json:"address_save"`
	IsEnteredAddress bool   `json:"is_entered_address"`
	Line1            string `json:"line_1"`
	Line2            string `json:"line_2"`
	State            string `json:"state"`
	Country          string `json:"country"`
}

type PaymentInfo struct {
	Number   string
	ExpMonth string
	ExpYear  string
	CVV      string
}

type DeliveryGroups struct {
	CartPayload struct {
		BlockSavedCards bool   `json:"block_saved_cards"`
		CartID          string `json:"cart_id"`
		ChannelInfo     struct {
			ChannelCode string `json:"channel_code"`
		} `json:"channel_info"`
		ChosenPayments []struct {
			PaymentMethod string `json:"payment_method"`
			PaymentOption string `json:"payment_option"`
		} `json:"chosen_payments"`
		CouponCodes             []any     `json:"coupon_codes"`
		CouponCodesFailureCount int       `json:"coupon_codes_failure_count"`
		Currency                string    `json:"currency"`
		ExpiryDate              time.Time `json:"expiry_date"`
		IsEppAuthRequired       bool      `json:"is_epp_auth_required"`
		IsExchangeTransformed   bool      `json:"is_exchange_transformed"`
		IsLoginRequired         bool      `json:"is_login_required"`
		IsStorePickupEligible   bool      `json:"is_store_pickup_eligible"`
		IsStrictZeroDiscount    bool      `json:"is_strict_zero_discount"`
		IsSupCart               bool      `json:"is_sup_cart"`
		Locale                  string    `json:"locale"`
		MultiQuantity           bool      `json:"multi_quantity"`
		OfferTriggerTags        []string  `json:"offer_trigger_tags"`
		PaymentAttributes       struct {
			BloatTaxPercentage int    `json:"bloat_tax_percentage"`
			CostSplitType      string `json:"cost_split_type"`
			CostSplitVersion   string `json:"cost_split_version"`
		} `json:"payment_attributes"`
		PoID                      string `json:"po_id"`
		PostalCode                string `json:"postal_code"`
		PricingValidationChecksum string `json:"pricing_validation_checksum"`
		ReferralCouponCodeDetails struct {
			CouponCode string `json:"coupon_code"`
			Index      int    `json:"index"`
		} `json:"referralCouponCodeDetails"`
		ReferralCodes       []string `json:"referral_codes"`
		RejectedCouponCodes []any    `json:"rejected_coupon_codes"`
		RetryFlags          struct {
			DeliveryModes  bool `json:"delivery_modes"`
			PaymentOptions bool `json:"payment_options"`
		} `json:"retry_flags"`
		RewardsRedemptionEligible bool   `json:"rewards_redemption_eligible"`
		SchemaVersion             string `json:"schema_version"`
		ShippingInfo              struct {
			AddressSave      bool   `json:"address_save"`
			City             string `json:"city"`
			Country          string `json:"country"`
			Email            string `json:"email"`
			FirstName        string `json:"first_name"`
			IsEnteredAddress bool   `json:"is_entered_address"`
			IsPrimaryAddress bool   `json:"is_primary_address"`
			LastName         string `json:"last_name"`
			Line1            string `json:"line_1"`
			Line2            string `json:"line_2"`
			Phone            string `json:"phone"`
			PostalCode       string `json:"postal_code"`
			State            string `json:"state"`
		} `json:"shipping_info"`
		Status            string `json:"status"`
		SystemCouponCodes []any  `json:"system_coupon_codes"`
		Type              string `json:"type"`
		UserAlerts        struct {
			CartAlerts      []any `json:"cart_alerts"`
			LineItemsAlerts []any `json:"line_items_alerts"`
		} `json:"user_alerts"`
		DeliveryGroups map[string]interface{} `json:"delivery_groups"`
		OfferGroups    []struct {
			GroupID    string `json:"group_id"`
			Discounted []struct {
				LineItemID           string `json:"line_item_id"`
				AssociatedQuantities []int  `json:"associated_quantities"`
			} `json:"discounted"`
			OfferID   string `json:"offer_id"`
			Triggered []struct {
				LineItemID           string `json:"line_item_id"`
				AssociatedQuantities []int  `json:"associated_quantities"`
			} `json:"triggered"`
			DeviceID             string `json:"device_id,omitempty"`
			DeviceGrade          string `json:"device_grade,omitempty"`
			AssociatedQuantities []int  `json:"associated_quantities,omitempty"`
			ChildExchangeID      string `json:"child_exchange_id,omitempty"`
		} `json:"offer_groups"`
		RejectedOfferTriggerTags []struct {
			TriggerTag      string `json:"trigger_tag"`
			RejectionReason []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"rejection_reason"`
		} `json:"rejected_offer_trigger_tags"`
		DeliverySignature struct {
			Status string `json:"status"`
		} `json:"delivery_signature"`
	} `json:"cart_payload"`
	DeliveryResponse []struct {
		DeliveryGroupID string `json:"delivery_group_id"`
		LineItems       []struct {
			LineItemID string `json:"line_item_id"`
			Quantity   int    `json:"quantity"`
		} `json:"line_items"`
		DeliveryModes []struct {
			AdditionalTaxes          []any `json:"additional_taxes"`
			ApplicablePaymentOptions []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"applicable_payment_options"`
			CarrierMappingAttributes struct {
				IsEsimEnabled        bool   `json:"is_esim_enabled"`
				SimTypeForActivation string `json:"sim_type_for_activation"`
			} `json:"carrier_mapping_attributes"`
			CommonCode       string `json:"common_code"`
			CurrentVersion   int    `json:"currentVersion"`
			CustomProperties struct {
				AllowBundleCancel                         bool   `json:"allow_bundle_cancel"`
				CarrierCode                               any    `json:"carrier_code"`
				CarrierDeliveryModeID                     any    `json:"carrier_delivery_mode_id"`
				ChargePerItem                             bool   `json:"charge_per_item"`
				CommonDivision                            any    `json:"common_division"`
				CutOffTime                                any    `json:"cut_off_time"`
				DefaultDeliveryMode                       bool   `json:"default_delivery_mode"`
				DeliveryEnabledByDefaultForAllPostalCodes bool   `json:"delivery_enabled_by_default_for_all_postal_codes"`
				DeliveryModeType                          string `json:"delivery_mode_type"`
				DeliveryServiceType                       any    `json:"delivery_service_type"`
				ExchangeCarrierCode                       any    `json:"exchange_carrier_code"`
				ExchangeServiceCode                       any    `json:"exchange_service_code"`
				ExternalMerchantID                        any    `json:"external_merchant_id"`
				GerpSku                                   any    `json:"gerp_sku"`
				GracePeriod                               int    `json:"grace_period"`
				IsEnergyStarCertified                     bool   `json:"is_energy_star_certified"`
				IsGiftcardSku                             bool   `json:"is_giftcard_sku"`
				IsTrailAvailable                          bool   `json:"is_trail_available"`
				ModelDivision                             any    `json:"model_division"`
				OverrideDeliveryGroup                     any    `json:"override_delivery_group"`
				PartnerContractMarginUnit                 string `json:"partner_contract_margin_unit"`
				PartnerContractMarginValue                int    `json:"partner_contract_margin_value"`
				PlanTierValue                             int    `json:"plan_tier_value"`
				ReturnCarrierCode                         any    `json:"return_carrier_code"`
				ReturnServiceCode                         any    `json:"return_service_code"`
				SaCode                                    any    `json:"sa_code"`
				ServiceCode                               any    `json:"service_code"`
				ShipCondition                             any    `json:"ship_condition"`
				ShipConditionExchange                     any    `json:"ship_condition_exchange"`
				ShipConditionReturns                      any    `json:"ship_condition_returns"`
				ShipType                                  any    `json:"ship_type"`
				ShipTypeExchange                          any    `json:"ship_type_exchange"`
				ShipTypeReturns                           any    `json:"ship_type_returns"`
				Spi                                       any    `json:"spi"`
				SpiExchange                               any    `json:"spi_exchange"`
				SpiReturns                                any    `json:"spi_returns"`
				SuspensionPeriod                          int    `json:"suspension_period"`
				SuspensionPeriodUnit                      string `json:"suspension_period_unit"`
				TaxApplied                                bool   `json:"tax_applied"`
				TrackingURL                               any    `json:"tracking_url"`
				TrailPeriodUnit                           string `json:"trail_period_unit"`
				TrailPeriodValue                          int    `json:"trail_period_value"`
				UIOrder                                   string `json:"ui_order"`
				UseSpiValuesForGerpSku                    any    `json:"use_spi_values_for_gerp_sku"`
				VatReversal                               bool   `json:"vat_reversal"`
			} `json:"custom_properties,omitempty"`
			DeliveryGroup        string `json:"delivery_group"`
			DeviceAsAServiceSkus []any  `json:"device_as_a_service_skus"`
			Flags                struct {
				EcomFlag                      bool `json:"ecom_flag"`
				IsAccessory                   bool `json:"is_accessory"`
				IsCancellable                 bool `json:"is_cancellable"`
				IsCustomizable                bool `json:"is_customizable"`
				IsFinancingEligible           bool `json:"is_financing_eligible"`
				IsInsuranceProduct            bool `json:"is_insurance_product"`
				IsReturnRequired              bool `json:"is_return_required"`
				IsSecure                      bool `json:"is_secure"`
				IsSubscription                bool `json:"is_subscription"`
				IsTradeInEligible             bool `json:"is_trade_in_eligible"`
				IsReturnShippingLabelEligible bool `json:"is_return_shipping_label_eligible"`
				IsUpgradeEligible             bool `json:"is_upgrade_eligible"`
			} `json:"flags"`
			FulfillerAttributes []struct {
				CarrierName                 string `json:"carrier_name"`
				CommonDivision              string `json:"common_division"`
				GerpSku                     string `json:"gerp_sku"`
				IsActive                    bool   `json:"is_active"`
				ModelDivision               string `json:"model_division"`
				ShipCondition               string `json:"ship_condition"`
				ShipConditionExchange       string `json:"ship_condition_exchange"`
				ShipConditionReturns        string `json:"ship_condition_returns"`
				ShipType                    string `json:"ship_type"`
				ShipTypeExchange            string `json:"ship_type_exchange"`
				ShipTypeReturns             string `json:"ship_type_returns"`
				ShippingMethodFulfillerCode string `json:"shipping_method_fulfiller_code"`
				Spi                         string `json:"spi"`
				SpiExchange                 string `json:"spi_exchange"`
				SpiReturns                  string `json:"spi_returns"`
			} `json:"fulfiller_attributes"`
			Images struct {
				LargeImage struct {
				} `json:"large_image"`
				SmallImage struct {
				} `json:"small_image"`
			} `json:"images"`
			IsB2B                      string `json:"is_b2b"`
			MaxPurchasableLimit        int    `json:"max_purchasable_limit"`
			ModelCode                  string `json:"model_code"`
			ModelName                  string `json:"model_name"`
			NumberOfReviews            any    `json:"number_of_reviews"`
			ProductDisplayName         string `json:"product_display_name"`
			ProductType                string `json:"product_type"`
			ReviewRating               any    `json:"review_rating"`
			SalesOrg                   any    `json:"sales_org"`
			ShipAlone                  bool   `json:"ship_alone"`
			ShortDescription           string `json:"short_description"`
			ShortRedemptionInstruction any    `json:"short_redemption_instruction"`
			Sku                        string `json:"sku"`
			SkuSource                  string `json:"sku_source"`
			Tags                       []any  `json:"tags"`
			TenantID                   string `json:"tenant_id"`
			TradeIn                    struct {
				TradeInMapping []any `json:"trade_in_mapping"`
			} `json:"trade_in"`
			UpcCode any `json:"upc_code"`
			Urls    struct {
			} `json:"urls"`
			UserProperties struct {
			} `json:"user_properties"`
			UIOrder                string `json:"ui_order"`
			DefaultDeliveryMode    bool   `json:"default_delivery_mode"`
			GroupName              string `json:"group_name"`
			GroupID                string `json:"group_id"`
			DeliveryServices       []any  `json:"delivery_services"`
			PostalCodeAvailability bool   `json:"postal_code_availability"`
			SalePrice              struct {
				Value    float64 `json:"value"`
				Currency string  `json:"currency"`
			} `json:"sale_price"`
			Price struct {
				SalePrice struct {
					Value    float64 `json:"value"`
					Currency string  `json:"currency"`
				} `json:"sale_price"`
				MsrpPrice struct {
					Value    float64 `json:"value"`
					Currency string  `json:"currency"`
				} `json:"msrp_price"`
			} `json:"price"`
			ApplicablePaymentProviders []any `json:"applicable_payment_providers,omitempty"`
			BestTradeInDeviceIds       []any `json:"best_trade_in_device_ids,omitempty"`
			CustomProperties0          struct {
				AllowBundleCancel                         bool   `json:"allow_bundle_cancel"`
				ChargePerItem                             bool   `json:"charge_per_item"`
				DefaultDeliveryMode                       bool   `json:"default_delivery_mode"`
				DeliveryEnabledByDefaultForAllPostalCodes bool   `json:"delivery_enabled_by_default_for_all_postal_codes"`
				DeliveryModeType                          string `json:"delivery_mode_type"`
				GracePeriod                               int    `json:"grace_period"`
				IsEnergyStarCertified                     bool   `json:"is_energy_star_certified"`
				IsGiftcardSku                             bool   `json:"is_giftcard_sku"`
				IsTrailAvailable                          bool   `json:"is_trail_available"`
				PartnerContractMarginUnit                 string `json:"partner_contract_margin_unit"`
				PartnerContractMarginValue                int    `json:"partner_contract_margin_value"`
				PlanTierValue                             int    `json:"plan_tier_value"`
				SuspensionPeriod                          int    `json:"suspension_period"`
				SuspensionPeriodUnit                      string `json:"suspension_period_unit"`
				TaxApplied                                bool   `json:"tax_applied"`
				TrailPeriodUnit                           string `json:"trail_period_unit"`
				TrailPeriodValue                          int    `json:"trail_period_value"`
				UIOrder                                   string `json:"ui_order"`
				VatReversal                               bool   `json:"vat_reversal"`
			} `json:"custom_properties,omitempty"`
			Frequency []any `json:"frequency,omitempty"`
			Includes  []any `json:"includes,omitempty"`
			Requires  []any `json:"requires,omitempty"`
			Supports  []any `json:"supports,omitempty"`
		} `json:"delivery_modes"`
	} `json:"delivery_response"`
}
