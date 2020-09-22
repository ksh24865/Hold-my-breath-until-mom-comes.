# Smart Baby Care System
Check and manage your baby's condition in real time.

Table of contents
=================
<!--ts-->
   * [Concept drawing](#concept-drawing)
   * [Role](#role)
   * [Result](#result)
   * [Requirement](#requirement)
   * [Installation](#installation)
   * [Run](#run)
   
<!--te-->
Concept drawing
===============
  * Concept drawing
  
![Concept_drawing](https://user-images.githubusercontent.com/55729930/93132617-f29af380-f710-11ea-8908-0b2027bf75bd.png)

  * To do
  
    * 웹 서버 구축 (데이터 가시화, 캠 조작, 젖병데이터 저장)
    * 라즈베리파이 + esp32 캠 -> 캠 조작 서버  
    * 라즈베리파이 온도, 습도 센서연결 -> 서버에 온습도 정보 전송  
    * 라즈베리파이 가습기 모듈 연결 -> 라즈베리파이 자체적으로 습도판단해서 작동
    * 라즈베리파이 중량 모듈 연결 -> push 중량체크 push ->서버에 보냄 
    * 라즈베리파이, esp 쿨링 
    * 라즈베리파이 전기장판(보류)

Role
=======

|Name|Role| 
|:----:|:----:|
|민성|중량모듈, 하드웨어 쿨링|
|성호|온도&습도 센서, 가습기 모듈|
|은석|라즈베리파이 + esp32 캠|

Result
=======



Requirement
=======

```
Python >= 3.0
```

Installation
=======
* clone
```sh
$ git clone https://github.com/ksh24865/I_will_hold_my_breath_until_my_mom_comes.git
```

* set
  * usb-control (usb포트 제어)
  ```sh
  $ apt-get install libusb-dev
  $ sudo apt-get update
  $ sudo apt-get install libusb-dev
  $ git clone https://github.com/codazoda/hub-ctrl.c
  $ cd hub-ctrl.c
  $ gcc -o hub-ctrl hub-ctrl.c -lusb
  $ cp hub-ctrl ~
  ```
  * adafruit_dht (온습도 측정 제어)
  ```sh
  $ sudo pip3 install adafruit_dht
  ```
  

Run
=======
