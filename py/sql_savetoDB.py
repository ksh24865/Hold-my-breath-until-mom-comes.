import pymysql
import datetime
"""
columns in table must be same!!!!!!
(temp, humd, weight)
"""
def savetoDB(DBname, Tablename, temperature, humidity, weight):
    now = datetime.datetime.now()
    nowdate = now.strftime('%Y%m%d')
    if(now.day == '1'):
        db = pymysql.connect(host='localhost', user = 'root', password='0000', db = DBname ,charset='utf8')
        cursor = db.cursor()
        sql = "INSERT INTO WEIGHT (weight, date) VALUES (%s, "+now.strftime('%m')+")"
        cursor.execute(sql, (weight))
        db.commit()
        sql = "UPDATE WEIGHT SET weight = " +str(weight)+ ", date = " +now.strftime('%m')+ " where id = 1"
        cursor.execute(sql)
        db.commit()
        db.close()
    db = pymysql.connect(host='localhost', user = 'root', password='0000', db = DBname ,charset='utf8')
    cursor = db.cursor()
    sql = "INSERT INTO " + Tablename + " (temperature, humidity, weight, dates) VALUES (%s, %s, %s, "+nowdate+")"
    cursor.execute(sql, (temperature, humidity, weight))
    db.commit()
    sql = "UPDATE SIGNALS SET temperature = " + str(temperature) + ", humidity = " +str(humidity)+ \
    ", weight = " +str(weight)+ ", dates = " + nowdate + " where id = 1"
    cursor.execute(sql)
    db.commit()
    db.close()
