# Password Cracking

## Easy
```
Title: Getting Started
Description: This is an MD5 hash of a easy password to crack.
Hint: This hash is probably listed online and can be viewed from a Google search.
Hash: 5f4dcc3b5aa765d61d8327deb882cf99
Password: password
```
```
Title: rockyou
Description: This is a SHA256 of a password found in a common dictionary attack. If new to password cracking, lookup how to use the tool 'hashcat'.
Hint: This hash can probably be cracked from using a very common dictionary attack that will rockyou
Hash: f70a65c5a5329bbac866029a2a500a70c4e5d538b638aaa463fd8b90028547fd
Password: jacksbbq1
```
```
Title: Missing Digits
Description: This is a SHA256 of a password that we know starts with `thepigbbq` and then has 5 digits after, but we don't know what the digits are.
Hint: This hash can be cracked by using a mask attack to brute force all combinations of 5 digits.
Hash: 0ccddfc1db55e9b637a89ff6328c234e6de98add7ace9acd64ecd32ba9883927
Password: thepigbbq29174
```

## Medium
```
Title: Qwerty
Description: This is a SHA256 of a password that we know is 10 - 16 characters and only uses the lowercase letters q, w, e, r, t, & y.
Hint: This hash can be cracked from using the -i --increment-min=10 flags for hashcat.
Hash: b97bdf8221d5c3913f95a8624e679d09be49bed43a2fbb87145cc1211cf67c17
Password: twyqwrewyqwtewt
```
```
Title: Names
Description: This is a SHA256 of a password that we know is a person's name followed by 4 digits.
Hint: This hash can be cracked using a wordlist of lowercase names and a mask attack adding 4 digits to the end.
Hash: e80742e2a469e8c5a7108149c6c48f5b2d1c65d0181c57b70649a6c7a97042ab
Password: shayla1986
```
```
Title: Known Pattern
Description: This is a SHA256 of a password that we know follows the pattern lowercase adjective + lowercase noun + 3 digits
Hint: This hash can be cracked from using a wordlist of adjective + noun that phrases that follow a similar pattern to netgear/spectrum default wifi router passwords.
Hash: a6684c79789af49701747d9c20a95bdf460fdf9de58e89203ffd72d307083b45
Password: redlightning284
```

## Hard
```
Title: 1337
Description: This is a SHA256 of a password that is an all lowercase word with the o's swapped with 0s and the e's replaced with 3's. Ex: wonder --> w0nd3r
Hint: Because the word is unknown a large wordlist should be used with a rule that changes o's to 0's and e's to 3's.
Hash: 12c160ace0afb52e9597bbeb98126706227bebf2002d7b460c2f3d69b0c94086
Password: st0n3h3ng3
```
```
Title: Passphrases
Description: This is a SHA256 of a password that was made following the guidelines of using a passphrase. In this case, they used a series of lowercase words separated by spaces. The words are not random.
Hint: Perhaps a passphrase dictionary attack could be useful.
Hash: 5500914759d0c942feb88ee404d1e377e1ce65bcd2dd2c6d3b0e6d7cfe159995
Password: rapid eye movement during sleep
```
```
Title: Guidelines
Description: This is a SHA256 of a password that was made following the guidelines to be at least 10 characters with at least 1 uppercase, lowercase, symbol, & number.
Hint: It can be assumed that this password follows the guidelines in ways that people tend follow them. Uppercase letters are usually first, the symbol usually comes after the word followed by some digits.
Hash: f5f0f5a663df92d0518db4ee573675ecd4b22812285c11672d50a641bd93de80
Password: Pigh3ad!839
```