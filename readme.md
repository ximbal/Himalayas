# Himalayas: An email alias manager for Cpanel (Create Aliases from Anywhere!)

## Why?

Well in essence because I've been tired of suffering with spam mail for 25 years, the logic that I use is:

You're asked for an email address, let's say by any ecommerce merchant/airline/bank/you name it!.

The merchant might:

* Suffer from a leak of your email address (typical) and then falls into the wrong hands.

* Have unethical practices and their unsubscribe button on their emails either won't work or they simply won't stop.

Me on the other hand would like to:

* Be able to stop the spamming on its tracks, no more regex editing.
* Be able to tell who leaked my info (and potentially decide to do something about it)
* Integrate with **CPANEL** (because it is easy to setup your own email)
* Have the ability to cross compile and install in my Mac, in my Linux Box and on my mobile phone (I have it working on a Mac and on an Android Phone)
* and most important of all, be able to create a custom email address (an alias) that I can create on the fly and not give away my real email address.
* Bonus tip, what I usually do is create an email address with the name of the business@myseconddomain.com and let the forwarder send it to myrealemailaddess@mymaindomain.com



### To Compile/Run

```
go build .
```

### to Execute the binary
```
./emailalias
```
### To configure it to work with your domain:

rename `config/config.json.example` to `config/config.json` then edit the values accordingly:

{
    "cpanelAPIKey": "USERNAME:SOMEKEY",
    "cpanelHost": "cpanel.yourDOMAIN.com:2083",
    "defaultForwardersDomain": "domain1.com",
    "yourRealEmail":"niceperson@domain1.com",
    "domains": ["domain1.com", "domain2.com"]
}

Where the USERNAME:SOMEKEY are your cpanel username and the api key. You need to get those out of your cpanel admin console.

### to package for Mac OS
```
fyne package -os darwin -icon iconAlias.png

Then

Drag and drop to your Applications folder
```

I have cross compiled and created a Fyne Android App for my mobile phone.

Please refer to Fyne's Guide on packaging the App for Mobile.

https://developer.fyne.io/started/mobile

