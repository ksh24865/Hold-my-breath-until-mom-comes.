#usb_control.py
import time
import os

os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 0")
time.sleep(5)
os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 1")