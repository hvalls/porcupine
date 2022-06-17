# PORCUPINE

Event-sourcing database

## HTTP API

### **GET /streams/{streamId}/events**

Get stream events

Response:

|Property|Type|
|---|---|
|success|boolean|  
|events|Event[]|  
|   |   |  

### **POST /streams/{streamId}/events**

Append stream event

Request:
|Property|Type|
|---|---|
|type|String|  
|data|string|  
|   |   |  

Response:

|Property|Type|
|---|---|
|success|boolean|  
|   |   |  

### **Type: Event**
|Property|Type|
|---|---|
|id|String|  
|number|int|  
|streamId|string|  
|type|string|  
|data|string|  
|   |   |  