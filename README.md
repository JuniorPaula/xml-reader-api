# XML Reader API

## Introdução

**XML Reader API** é uma API simples de consulta de um arquivo xml armazena na propria aplicação. O modo de operação consiste em carregar o arquivo em memória quando a aplicação é iniciada, disponibilizando assim os dados para posterior consulta via enpoint.
Alem da leitura do arquivo a aplicação ainda conta com sistema de autenticação que logará o cliente disponibilizando o token para relalizar as consultas.
  
### Dependências

[Golang] https://go.dev/doc/install
[SQLite3] https://www.sqlite.org/

## Referências

[SQLite3] https://github.com/mattn/go-sqlite3/blob/master/example/simple/simple.go

[Bcript] https://pkg.go.dev/golang.org/x/crypto/bcrypt

[Chi] https://github.com/go-chi/chi

[Viper] https://github.com/spf13/viper

[Excelize] https://github.com/qax-os/excelize

### Modo de usar
```bash
$ git clone url
```

```bash
$ go mod tidy
```
#### .env
Criar um arquivo .env e copiar o .env.example para configurar as variáveis de ambiente.

### Execultar
```bash
$ go run ./cmd/api/main.go 
```
### Build
```bash
$ go build -o app ./cmd/api/main.go && ./app
```

### Testes
```bash
$ go test -v ./...
```

#### Cobertura de testes
```bash
$ go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```

### Endpoints
#### Cadastro
[POST] `/signup`
- body
```json
{
	"name": string,
	"email": string,
	"password": string
}
```

**Resposta**

**201**
```json
{
	"message": string,
	"error": boolean,
	"data": {
		"id": number,
		"name": string,
		"email": string
	}
}
```

**400**
```json
{
	"message": string
	"error": boolean
}
```

#### Login
[POST] `/login`
- body
```json
{
	"email": string,
	"password": string
}
```

**Resposta**

**200**
```json
{
	"message": string,
	"error": boolean,
	"data": {
		"id": number,
		"name": string,
		"email": string,
		"access_token": string
	}
}
```

**401**
```json
{
	"message": string
	"error": boolean
}
```

#### Login
[GET] `/suppliers`
- Header
`Authorization` 
`Bearer {token}`

- Query
`limit` default = 200
`offset` default = 0

**Resposta**

**200**
```json
{
	"message": string,
	"error": boolean,
	"data": {
		"limit": number,
		"offset": number,
		"suppliers": [
			{
				"ParterID": string,
				"PartnerName": string,
				"CustomerID": string,
				"CustomerName": string,
				"CustomerDomainName": string,
				"CustomerCountry": string,
				"MpnID": number,
				"Tier2MpnID": number,
				"InvoiceNumber": string,
				"ProductID": string,
				"SKUID": string,
				"AvailabilityID": string,
				"SKUName": string,
				"ProductName": string,
				"PublisherName": string,
				"PublisherID": string,
				"SubscriptionDescription": string,
				"SubscriptionID": string,
				"ChargeStartDate": string,
				"ChargeEndDate": string,
				"UsageDate": string,
				"MeterType": string,
				"MeterCategory": string,
				"MeterID": string,
				"MeterSubCategory": string,
				"MeterName": string,
				"MeterRegion": string,
				"Unit": string,
				"ResourceLocation": string,
				"CostomerService": string,
				"ResourceGroup": string,
				"ResourceURI": string,
				"ChargeType": string,
				"UnitPrice": number,
				"Quantity": number,
				"UnitType": string,
				"BillingPreTaxTotal": number,
				"BillingCurrency": string,
				"PricingPreTaxTotal": number,
				"PricingCurrency": string,
				"ServiceInfo1": string,
				"ServiceInfo2": string,
				"Tags": string,
				"AdditionalInfo": string,
				"EffectiveUnitPrice": number,
				"PCToBCExchangeRate": number,
				"PCToBCExchangeRateDate": string,
				"EntitlementId": string,
				"EntitlementDescription": string,
				"PartnerEarnedCreditPercentage": number,
				"CreditPercentage": number,
				"CreditType": string,
				"BenefitOrderId": string,
				"BenefitId": string,
				"BenefitType": string
			}
		],
		"total": number
	}
}
```

**401**
```json
{
	"message": string
	"error": boolean
}
```