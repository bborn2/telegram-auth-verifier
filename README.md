# what is this
Fork form [here](https://github.com/electrofocus/telegram-auth-verifier) and fix many bugs

# telegram-auth-verifier
Golang package for [Telegram Website Login](https://core.telegram.org/widgets/login#checking-authorization) credentials verification. Check documentation [here](https://pkg.go.dev/github.com/electrofocus/telegram-auth-verifier).

## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```
go get github.com/bborn2/telegram-auth-verifier
```

## Example

Let's verify credentials:

```go
import (
	"encoding/json"
	"fmt"
	"net/url"

	tgverifier "github.com/electrofocus/telegram-auth-verifier"
)

func main() {
	token := []byte("Your Telegram Bot Token")
	
	creds := tgverifier.Credentials{
		ID:        111111111,
		FirstName: "Kun",
		LastName:  "Song",
		Username:  "recoba",
		PhotoURL:  "http://tg.me/xxx",
		AuthDate:  1617443847,
		Hash:      ”ae1b08443b7bb50295be3961084c106072798cb65e91995a1b49927cd4cc5b0c“,
	}

	if err := creds.Verify(token); err != nil {
		fmt.Println("Credentials are not from Telegram")
		return
	}

	fmt.Println("Credentials are from Telegram")
}
```
