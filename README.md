## Simple Numeral System Translator API

This API just translate one integer number from one numeral system to another

### Available endpoints

- /api/translate

### Request format

```json
{
    "number": {
        "value":"35",
        "base":10
    },
    "toBase":3
}
```

### Response format

Response example for endpoint /api/translate:
```json
{
    "result": "1022"
}
```
