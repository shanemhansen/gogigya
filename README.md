gogigya
=======

Bindings to the Gigya REST API http://developers.gigya.com/037_API_reference

Included:

   * Signing api keys, secrets, and all the machinery
   * Generic request/response binding
   * Support for User/Account schema
   * Demo application for reporting on users
   * Demo application making generic api calls


Installation
--------

    # install the library:
    go get github.com/shanemhansen/gogiya
    
    // use in your .go code:
    import (
         "github.com/shanemhansen/gogigya"
    )


Examples
--------

     #Add GOPATH/bin to your PATH
     go install github.com/shanemhansen/gogigya/...
     APIKEY=XXX SECRET=YYY gigdump -apiKey $APIKEY -secret $SECRET -report "UID:{{.UID}" -count 1
     APIKEY=XXX SECRET=YYY gig -apiKey $APIKEY -secret $SECRET socialize.getUserInfo UID 1
