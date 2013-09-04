package gogigya

import "flag"


var ApiKey *string
var SecretKey *string

func init() {
    ApiKey = flag.String("apiKey","", "API Key")
    SecretKey = flag.String("secret","", "Secret Key")
}
