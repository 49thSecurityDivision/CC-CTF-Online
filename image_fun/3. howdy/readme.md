
## solution
They need to remove the magicnum\n from the start of the png
it can be done by running `tail -c +10 howdy.png > how.png` skipping the bad bytes

## hint1
That magic number looks weird

## hint2
Did you know you can skip bytes with tail -C +<num>

## answer
flag{howdythere,ya'll_didagoodjob}
