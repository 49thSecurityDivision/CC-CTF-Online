# The Pig BBQ

The Pig BBQ is a series of simple web exploits.

### Flags:
1. The first flag can be found commented out in the `index.html` page.  
`flag{LWhRU-uUq62Seohdsj8hX0K5a0q5KywQidq8zYtPKMWs4q2Royr51meLCWLu7Zan}`
1. The second flag can be found by running the js function `getFlag` from the console on the order page. These kinds of flags can be found by searching the front end pages for `flag`.  
`flag{rVS11g-XBgWah7HTZWzfmUvoMcnmIvDs}`
1. The `robots.txt` file features 2 interesting lines. The first is a link to a flag.  
`flag{hL7K5FsztoA8dx2gDWtaUPvDSX6P7uV7}`  
The second line reveals that there is an unlisted login page.
1. If the user tries to login with the username `admin` the login will be successful (we ran out of time to make this more interesting)  
After a successful login it is revealed that the menu now shows a list of hack4hire services. One of those services is a flag.  
`flag{AHHmt-uZyqy4Y3FNHXi8YBuRmPL9P1Sh}`
1. The api has 2 hidden flags. The first is just from making a get request the `/flag` endpoint.  
`flag{uHSgglifeO9tlIJkGGxcMC1tKXCRwLN5}`
1. The second is by customizing a post request to the `/order` endpoint with `flag` or `null` in the `order` field.  
`flag{EI46mGgilJPWOJhwfmEqsrdbsyibzDz1}`