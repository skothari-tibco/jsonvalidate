---
title: JSON-Validator

---

# JSON-Validator
This activity allows you to validate JSON




## Schema
Inputs and Outputs:

```json
"input":[
      {
        "name": "text",
        "type": "string",
        "value": ""
      },
      {
        "name": "path",
        "type":"string",
        "value":""
      }
    ],
    "output": [
      {
        "name": "isValid",
        "type": "bool"
      }
    ]
```

## Settings
### Validator:
| Input     | Description    |
|:------------|:---------------|
| test | The JSON to validate | 
| path | The path of JSON schema. Can be a URL or a file Path. Can also be null  |            
