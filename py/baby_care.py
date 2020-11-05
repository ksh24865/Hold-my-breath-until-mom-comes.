"""
REQUIRMENTS
Adafruit_DHT, MYSQL, PYMYSQL, savetoDB.py


GPIO SETTINGS
02(DC 5V)   (LOADCELL DC, Adafruit_DHT.DHT11 DC, RELAY DC) -> NEED GPIO DISTRIBUTER!!
04(DC 5V)   RSP FAN DC
06(GND)     RSP FAN GND
30(GND)     RELAY GND
29(GPIO05)  RELAY INPUT1 (LAMP)
31(GPIO06)  RELAY INPUT4 (BLANKET)
33(GPIO13)  RELAY INPUT3 (HUMIDITY)
34(GND)     LOADCELL GND
36(GPIO16)  LOADCELL SCK
38(GPIO20)  LOADCELL DT
39(GND)     Adafruit_DHT.DHT11 GND
40(GPIO21)  Adafruit_DHT.DHT11 DT
"""

import Adafruit_DHT
import RPi.GPIO as GPIO
import os
import time
import sys
from sql_savetoDB import savetoDB

#BLANKET SETUP
ONOFF_BLANKET = False
GPIO.setmode(GPIO.BCM)
GPIO.setup(6,GPIO.OUT)
GPIO.setup(13,GPIO.OUT)
#HUMID, TEMP SETUP

Reference_humidity = 50
sensor = Adafruit_DHT.DHT11
pin = 21
ONOFF_MOIST = False
t = 10

dummydata = 0 #used for dummy value
#WEIGHT SETUP
EMULATE_HX711=False
referenceUnit = -470
if not EMULATE_HX711:
    import RPi.GPIO as GPIO
    from hx711 import HX711
    from gpiozero import LED, Button
    from signal import pause
else:
    from emulated_hx711 import HX711

hx = HX711(20, 16) #HX711 PIN NUM DT 20, SCK 16
hx.set_reading_format("MSB", "MSB") 
hx.set_reference_unit(referenceUnit)
hx.reset()
hx.tare()
print("Tare done! Add weight now...")
#setup end

#MAIN
print("---START---")

while 1 :
    #HX711 GET WEIGHT PART
    weight = int(hx.get_weight(5))
    print("Weight : {0}".format(weight))
    
    #TEMP, HUMID PART
    
    humid, temper = Adafruit_DHT.read_retry(sensor, pin)
    
    if humid is not None and temper is not None :
        print("Temperature = {0:0.1f}*C Humidity = {1:0.1f}%".format(temper, humid))
        if ONOFF_MOIST == False and humid < Reference_humidity:
            print("MOIST---ON---")
            GPIO.output(13,False)
            ONOFF_MOIST = True
        if ONOFF_MOIST == True and humid >= Reference_humidity:
            print("MOIST---OFF---")
            GPIO.output(13,True)
            ONOFF_MOIST = False
        #Blanket
        if ONOFF_BLANKET == False and temper < 35:
            GPIO.output(6,False)
            print("BLANKET---ON---")
            ONOFF_BLANKET = True
        if ONOFF_BLANKET == True and temper > 35:
            GPIO.output(6,True)
            print("BLANKET---OFF---")
            
            ONOFF_BLANKET = False
    
    #DUMMYDATA
    #humid = dummydata
    #temper = dummydata
    #print("humid {0}, temper {1}".format(humid, temper))
    
    #Save to SQL
    #savetoDB(DBname, Tablename, temperature, humidity, weight)
    savetoDB('BABY_VIBE', 'SIGNALS', temper, humid, weight)
    
    
    #HX711 RESET 
    hx.power_down()
    hx.power_up()
    
    #wait 3 sec
    dummydata = dummydata + 1
    time.sleep(3)
        
        
