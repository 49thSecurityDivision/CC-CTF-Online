# noREST
I made this challenge with the goal of the user learning how to use ZAPs (or Burps) swagger import tool. I specifically made it so it would alert the sqli. I haven't really seen swaggers in CTFs, but they are pretty prevalent in the real world imo.

## Description
The goal of this challenge is to do a SQL injection to get a password from the `users` table. The table is sqlite and can be reset by restarting the application. 
The entire challenge is run from the one webserver.py file

## Prompt
Clarence can't remember his password to login to the cash register. I found some documentation on the backend api. Can you help him out?
+ Attach swagger.json
