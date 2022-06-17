# PORCUPINE

Event-sourcing database

## HTTP API

### **GET /streams/{streamId}/events**

Get stream events.

Request params: 
|Property|Type|
|---|---|
|streamId|String|  


Response body:

|Property|Type|
|---|---|
|success|boolean|  
|events|Event[]|  

### **POST /streams/{streamId}/events**

Append stream event.

Request params: 
|Property|Type|
|---|---|
|streamId|String|  

Request body:
|Property|Type|
|---|---|
|type|String|  
|data|string|  

Response body:

|Property|Type|
|---|---|
|success|boolean|  

### **Type: Event**
|Property|Type|
|---|---|
|id|String|  
|number|int|  
|streamId|string|  
|type|string|  
|data|string|  