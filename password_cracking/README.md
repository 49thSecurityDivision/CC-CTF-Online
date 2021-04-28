# Password Cracking

## Note:
We heard your feedback and we wont feature as many password cracking challenges next year. The ones we do feature we will ensure can be run on a CPU in < 15 minutes to ensure no one without a GPU is at a competitive disadvantage.

## Easy

Title: Getting Started  
Description: This is an MD5 hash of a easy password to crack.  
Hint: This hash is probably listed online and can be viewed from a Google search.  
Hash: 5f4dcc3b5aa765d61d8327deb882cf99  
Password: password  
Solution: Just google searching the hash will yield the result.

Title: rockyou  
Description: This is a SHA256 of a password found in a common dictionary attack. If new to password cracking, lookup how to use the tool 'hashcat'.  
Hint: This hash can probably be cracked from using a very common dictionary attack that will rockyou  
Hash: f70a65c5a5329bbac866029a2a500a70c4e5d538b638aaa463fd8b90028547fd  
Password: jacksbbq1  
Solution: Run the famous wordlist `rockyou.txt`. `hashcat -a 0 -m 1400 "f70a65c5a5329bbac866029a2a500a70c4e5d538b638aaa463fd8b90028547fd" rockyou.txt`

Title: Missing Digits  
Description: This is a SHA256 of a password that we know starts with `thepigbbq` and then has 5 digits after, but we don't know what the digits are.  
Hint: This hash can be cracked by using a mask attack to brute force all combinations of 5 digits.  
Hash: 0ccddfc1db55e9b637a89ff6328c234e6de98add7ace9acd64ecd32ba9883927  
Password: thepigbbq29174  
Solution: `hashcat -a 3 -m 1400 "0ccddfc1db55e9b637a89ff6328c234e6de98add7ace9acd64ecd32ba9883927" thepigbbq?d?d?d?d?d`


## Medium

Title: Qwerty  
Description: This is a SHA256 of a password that we know is 10 - 16 characters and only uses the lowercase letters q, w, e, r, t, & y.  
Hint: This hash can be cracked from using the -i --increment-min=10 flags for hashcat.  
Hash: b97bdf8221d5c3913f95a8624e679d09be49bed43a2fbb87145cc1211cf67c17  
Password: twyqwrewyqwtewt  
Solution: `hashcat -a 3 -m 1400 "b97bdf8221d5c3913f95a8624e679d09be49bed43a2fbb87145cc1211cf67c17" -i --increment-min=10 -1 qwerty ?1?1?1?1?1?1?1?1?1?1?1?1?1?1?1?1`
Admittedly, the keyspace for this challenge was too large and was out of reach for those without GPUs. With a GPU it can be solved in under 5 minutes, but it can easily take 4 hours on a CPU. We are sorry for that and will fix the problem next time around.

Title: Names  
Description: This is a SHA256 of a password that we know is a person's name followed by 4 digits.  
Hint: This hash can be cracked using a wordlist of lowercase names and a mask attack adding 4 digits to the end.  
Hash: e80742e2a469e8c5a7108149c6c48f5b2d1c65d0181c57b70649a6c7a97042ab  
Password: shayla1986  
Solution: `hashcat -a 6 -m 1400 "e80742e2a469e8c5a7108149c6c48f5b2d1c65d0181c57b70649a6c7a97042ab" NAMES.DIC ?d?d?d?d`
Any wordlist of common names can be used here. I used [NAMES.DIC](https://www.outpost9.com/files/wordlists/names.zip).

Title: Known Pattern  
Description: This is a SHA256 of a password that we know follows the pattern lowercase adjective + lowercase noun + 3 digits  
Hint: This hash can be cracked from using a wordlist of adjective + noun that phrases that follow a similar pattern to netgear/spectrum default wifi router passwords.  
Hash: a6684c79789af49701747d9c20a95bdf460fdf9de58e89203ffd72d307083b45  
Password: redlightning284  
Solution: `hashcat -a 6 -m 1400 "a6684c79789af49701747d9c20a95bdf460fdf9de58e89203ffd72d307083b45" netgear-spectrum.txt ?d?d?d`  
While other methods can be used to pair adjectives with nouns, I picked `redlighting` from the [netgear-spectrum](https://raw.githubusercontent.com/soxrok2212/PSKracker/master/dicts/netgear-spectrum/netgear-spectrum.txt) wordlist.


## Hard

Title: 1337  
Description: This is a SHA256 of a password that is an all lowercase word with the o's swapped with 0s and the e's replaced with 3's. Ex: wonder --> w0nd3r  
Hint: Because the word is unknown a large wordlist should be used with a rule that changes o's to 0's and e's to 3's.  
Hash: 12c160ace0afb52e9597bbeb98126706227bebf2002d7b460c2f3d69b0c94086  
Password: st0n3h3ng3  
Solution: `hashcat -a 0 -m 1400 "12c160ace0afb52e9597bbeb98126706227bebf2002d7b460c2f3d69b0c94086" -r 0-3.rule words_alpha.txt`  
`0-3.rule` only has one rule in it and that is `so0 se3`, which replaces o's with 0 and e's with 3's.
Any wordlist of common words can be used here, but I used [words_alpha.txt](https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt).

Title: Passphrases  
Description: This is a SHA256 of a password that was made following the guidelines of using a passphrase. In this case, they used a series of lowercase words separated by spaces. The words are not random.
Hint: Perhaps a passphrase dictionary attack could be useful.  
Hash: 5500914759d0c942feb88ee404d1e377e1ce65bcd2dd2c6d3b0e6d7cfe159995  
Password: rapid eye movement during sleep  
Solution: `hashcat -a 0 -m 1400 "5500914759d0c942feb88ee404d1e377e1ce65bcd2dd2c6d3b0e6d7cfe159995" passphrases.txt`  
This challenge was supposed to be solved using the [passphrase wordlist](https://github.com/initstring/passphrase-wordlist).

Title: Guidelines
Description: This is a SHA256 of a password that was made following the guidelines to be at least 10 characters with at least 1 uppercase, lowercase, symbol, & number.  
Hint: It can be assumed that this password follows the guidelines in ways that people tend follow them. Uppercase letters are usually first, the symbol usually comes after the word followed by some digits.  
Hash: f5f0f5a663df92d0518db4ee573675ecd4b22812285c11672d50a641bd93de80  
Password: Pigh3ad!839  
Solution: `hashcat -a 3 -m 1400 "f5f0f5a663df92d0518db4ee573675ecd4b22812285c11672d50a641bd93de80"  ?u?l?l?l?d?l?l!?d?d?d`
This challenge was designed to be very difficult, but in hindsight that puts a huge disadvantage on those without GPUs and that is a bad idea for a challenge. So, during the CTF very generous hints were given to help come up with a mask attack similar to the one above that will produce the result.