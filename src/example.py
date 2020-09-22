import time
import Adafruit_DHT
import os

#setup
Reference_humidity = 50
sensor = Adafruit_DHT.DHT11
pin = 5
ONOFF = False
t = 10



os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 0")
print("---START---")

while t>0 :
        humid, temper = Adafruit_DHT.read_retry(sensor, pin)
        
        if humid is not None and temper is not None :
            print("Temperature = {0:0.1f}*C Humidity = {1:0.1f}%".format(temper, humid))
            if ONOFF == False and humid < Reference_humidity:
                print("---ON---")
                ONOFF = True
                os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 1")
                
            if ONOFF == True and humid >= Reference_humidity:
                print("---OFF---")
                ONOFF = False
                os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 0")
        else :
            print('Read error')
        time.sleep(3)
        t=t-1
        
print("---END---")
os.system("sudo /home/pi/hub-ctrl -h 0 -P 2 -p 1")