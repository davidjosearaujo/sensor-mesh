# Messages

## Whisper

Online messages are **periodic and immutable by the user**, and used as a means of making sure that the node has not **died or disconnected**. If it has been lost connection to the node, it serves as way o knowing it's **last location and if it was moving**.

``` json
{
    "type": "whisper",
    "name": "0101_VFCI01",
    "time": 1562212768,
    "cam": {
        "accEngaged": true,
        "acceleration": 0,
        "altitude": 800001,
        "altitudeConf": 15,
        "brakePedal": true,
        "collisionWarning": true,
        "cruiseControl": true,
        "curvature": 1023,
        "driveDirection": "FORWARD",
        "emergencyBrake": true,
        "gasPedal": false,
        "heading": 3601,
        "headingConf": 127,
        "latitude": 40.0000000,
        "length": 10.0,
        "longitude": -8.0000000,
        "semiMajorConf": 4095,
        "semiMajorOrient": 3601,
        "semiMinorConf": 4095,
        "specialVehicle": {
            "publicTransportContainer": {
                "embarkationStatus": false
            }
        },
        "speed": 16383,
        "speedConf": 127,
        "speedLimiter": true,
        "stationID": 1,
        "stationType": 15,
        "width": 3.0,
        "yawRate": 0
    }
}
```

## Reading

Reading messages log the **values read from all the sensors** of the vehicle, and the time of that read. Plus, a binary flag is included to inform of the status of the sensor.

``` json
{
    "type": "online",
    "name": "0101_VFCI01",
    "time": 1562212768,
    "sensors":[
        {
            "name": "Sensor 1",
            "type": "temperature",
            "value": 30,
            "unit": "celsius",
            "read": 1562212758,
            "status": 1
        },
        {
            "name": "Sensor 2",
            "type": "accelerometer",
            "value": 10,
            "unit": "m/ss",
            "read": 1562212560,
            "status": 0
        }
    ]
}
```