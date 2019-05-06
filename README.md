# go-smite
Starter code for interacting with the SMITE API. SMITE is a MOBA video game published by Hi-Rez studios. This code can be used for accessing the API if you have valid credentials.

# To get started
- First apply to get developer credentials from Hi-Rez. This will include a Developer ID and a Authorization Key. You will need these to access the API. 
- Once you have the credentials make a json and make three string fields in the JSON: "dev_id", "auth_key", and "session_id". Put in your credentials into the "dev_id" and "auth_key" fields and you can leave "session_id" as the empty string if you want.
- Next set up a couple paths to files that you will want to use for storing Credentials, Item, and God data and put these paths into their respective variables in /src/smite/general/settings.go.
- Now that you have all that set up feel free to change the Language code if you want or even the response format if you want to get the xml and parse it not using the methods that I have provided.

# Using the API
After settings is all set up using the API is pretty simple. The code below shows an example call to get all items from the API into your json file.

```go

package main

import (
	S "smite/connection"
)

func main() {
	S.GetCredentialsAndUpdateSession().GetItems()
}
```
At the moment all of the API calls have not been set up and error handling consists of panicking. This will be updated shortly!
