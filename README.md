# [Golang-CoWin-Vaccine-Notifier]()
Automated Golang Script to retrieve vaccine slots availability and get notified when a slot is available.

In this used all inbuild library, for sound notification install PortAudio plugin in your os.

You can change below mention variable as per your need.

We only want to check for slots availability for the same day for 
let's givenum_days as 0 here.
```
num_of_days = 0
```    
if we want to check slot availability for next 2 days, so let's givenum_days as 2 here.
```
num_of_days = 2
```    

The Slots are available in two categories, Age 45+ and Age 18+. Let's proceed with Age 18+ for now, you can set it according to your need.
```
age = 25
```    
 Can perform the search based on Pincode. Since we are looking for the slots available nearby, I am making use of Pincode. You can also pass multiple pin-codes, separated by commans in the list.

```
pincode = []string{"395006"}
    or 
pincode = []string{"395006,395009"}
```

Used for notification, If you want to turn it on so pass true flag
```
PlaySound = false
```

Run script after every 5 minutes and check slot availability
```
Ticker = 5 * time.Minute
```

## pre requisite
```
To run sounds package you must first have the PortAudio development headers and libraries installed. Some systems provide a package for this; e.g., on Ubuntu you would want to run apt-get install portaudio19-dev. On other systems you might have to install from source.
``` 

LICENSE:
==========================
Copyright (c) 2021 Meet Soni

This project is licensed under the MIT License

<p align="center">
  <b><i>Let's connect! Find me on the web.</i></b>

[<img height="30" src="https://img.shields.io/badge/linkedin-blue.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />][LinkedIn]
<br />
<hr />

[linkedin]: https://www.linkedin.com/in/meetsoni1511/

  
If you have any Queries or Suggestions, feel free to reach out to me.