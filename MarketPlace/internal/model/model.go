// Package model models that use in this project
package model

// Vendor : struct for vendor
type Vendor struct {
	ID          string `bson,json:"id"`
	Name        string `bson,json:"name"`
	Password    string `bson,json:"password"`
	Balance     uint   `bson,json:"accessToken"`
	AccessToken string `bson,json:"accessToken"` // ???
}

//CозданиеНовогоПродавца()
//ДобавлениеТовара()
//ПополнениеБалансаЗаПроданныйТовар()

// Consumer : struct for consumer
type Consumer struct {
	ID          string `bson,json:"id"`
	Name        string `bson,json:"name"`
	Password    string `bson,json:"password"`
	AccessToken string `bson,json:"accessToken"` // ???
	Balance     uint   `bson,json:"accessToken"`
}

//CозданиеНовогоПокупателя()
//ДобавлениеТовараВКорзину()
//ПокупкаТовара()----//СписаниеСредствСБаланса()

// Config struct create config
type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
