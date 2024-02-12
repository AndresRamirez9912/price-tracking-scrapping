package constants

// Hosts
const (
	EXITO_HOST  = "www.exito.com"
	JUMBO_HOST  = "www.tiendasjumbo.co"
	AMAZON_HOST = "www.amazon.com"
)

// Exito Const
const (
	EXITO_STORE                     = "Exito"
	EXITO_WAIT_UNTIL                = "div[id='inpage_container']"
	EXITO_LOWER_PRICE_SELECTOR      = "span[data-fs-price='true'][data-testid='store-price'][data-variant='selling'].price_fs-price__7Y_0s"
	EXITO_HIGHER_PRICE_SELECTOR     = "p.priceSection_container-promotion_price-dashed__Pzc_z"
	EXITO_OTHER_PRICE_SELECTOR      = "p[data-fs-container-price-otros='true']"
	EXITO_PRODUCT_NAME_SELECTOR     = "h1.product-title_product-title__heading__eJJqz"
	EXITO_PRODUCT_BRAND_SELECTOR    = "span.product-title_product-title__specification__B5pYV"
	EXITO_PRODUCT_IMAGE_SELECTOR    = "img[width='750px']"
	EXITO_PRODUCT_DISCOUNT_SELECTOR = "span[data-percentage='true']"
)

// Jumbo Const
const (
	JUMBO_STORE                     = "Jumbo"
	JUMBO_WAIT_UNTIL                = "span.t-heading-4.v-mid"
	JUMBO_PRODUCT_NAME_SELECTOR     = "span.vtex-store-components-3-x-productBrand"
	JUMBO_PRICES_SELECTOR           = "div.tiendasjumboqaio-jumbo-minicart-2-x-price"
	JUMBO_PRODUCT_BRAND_SELECTOR    = "span.vtex-store-components-3-x-productBrandName"
	JUMBO_PRODUCT_IMAGE_SELECTOR    = "img.vtex-store-components-3-x-productImageTag.vtex-store-components-3-x-productImageTag--image-product-pdp.vtex-store-components-3-x-productImageTag--main.vtex-store-components-3-x-productImageTag--image-product-pdp--main"
	JUMBO_PRODUCT_DISCOUNT_SELECTOR = "span.tiendasjumboqaio-jumbo-minicart-2-x-containerPercentageFlag"
	JUMBO_PRODUCT_ID_SELECTOR       = "span.vtex-product-identifier-0-x-product-identifier__value"
)
