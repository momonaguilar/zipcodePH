# :philippines:🚩 zipcodePH

Returns an array of Philippine zipcode details given a search key.

## Use cases:
### 1. Retrieves details for specific zipcode.

Sample request:
```
curl -X GET http://localhost:8085/zipcode?key=6023
```

Result:
```
[{"zipcode":"6023","area":"Alcoy","provinceCity":"Cebu"}]
```

### 2. Retrieves all possible areas given a search key that matches several 'area' results
Search key is case insensitive.

Sample request:
```
curl -X GET http://localhost:8085/zipcode?key=pilar
```

Result:
```
[
 {"zipcode":"2101","area":"Pilar","provinceCity":"Bataan"},
 {"zipcode":"2812","area":"Pilar","provinceCity":"Abra"},
 {"zipcode":"4714","area":"Pilar","provinceCity":"Sorsogon"},
 {"zipcode":"5804","area":"Pilar","provinceCity":"Capiz"},
 {"zipcode":"6048","area":"Pilar","provinceCity":"Cebu"},
 {"zipcode":"6321","area":"Pilar","provinceCity":"Bohol"},
 {"zipcode":"8420","area":"Pilar","provinceCity":"Surigao del Norte"}
]
```

### 3. No matching zipcode or area keys will return empty array

Sample request:
```
curl -X GET http://localhost:8085/zipcode?key=momon
```

Result:
```
[]
```
### 4. Default endpoint

Sample request:
```
curl -X GET http://localhost:8085/zipcode?key=momon
```

Result:
```

```