#!/usr/bin/python3
import sys


class Rotator:
    commands = {
    'Q': 'quit',
    'q': 'quit',
    '?': 'help',
    'help': 'help',
    'P': 'set_pos',
    'set_pos': 'set_pos',
    'p': 'get_pos',
    'get_pos': 'get_pos',
    'S': 'stop',
    'stop': 'stop',
    }
    def __getitem__(self, key):
        func_name = self.commands.get(key, '')
        if func_name: return getattr(self, func_name)

    def help(self):
        print("""
Commands (some may not be available for this rotator):
P: set_pos     (Azimuth, Elevation)
p: get_pos     ()
K: park        ()
S: stop        ()
R: reset       (Reset)
""")


    def __init__(self):
        self.azimuth = 0.0
        self.elevation = 0.0
        self.running = True

    def quit(self):
        self.running = False

    def stop(self):
        pass

    def get_pos(self):
        output = 'Azimuth: %.6f\nElevation: %.6f' % (self.azimuth, self.elevation)    
        print(output)

    def set_pos(self, azimuth=None, elevation=None):
        if azimuth is None:
            azimuth = input("Azimuth: ").strip()
        if elevation is None:
            elevation = input("Elevation: ").strip()
            
        azi, ele = float(azimuth), float(elevation)
        self.azimuth = azi
        self.elevation = ele

        if (21 <= ele <= 23):
            if (274 <= azi <= 276) or (-85 <= azi <= -83):
                print("flag{tIrtwxaAK_pkNZ1jRvLITIOZKL6U2PFq}")
            


if __name__ == "__main__":
    rotator = Rotator()
    while rotator.running:
        line = input('Rotator command: ')
        #print(line, file=sys.stderr)
        line = line.split('#')[0]
        cmd, *args = line.strip().split(' ')
        func = rotator[cmd]

        if func:
            func(*args)
            print('')

        
        
