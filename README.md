## Colors Api
The Api for colors 

## Summary
This is the Api for colors! 🎨 <br/>
Written in Go ..

## Platforms
| Maintainer | Email | Platform | Register | Generator |
| ---------- | ----- | -------- | -------- | --------- |
| [@aikon001](https://github.com/aikon001 "@aikon001") | | Docker | aikon001/colorApi |

## Build
### Through docker
Run in local
```bash
docker run aikon001/colorApi
```
Or with docker compose 
```bash
docker-compose up --build
```
### From source
Run with Go
```bash
git clone github.com/aikon001/colorApi
cd colorApi
go run main.go
```

## Usage
### Examples
```bash
curl https://tylegroup.com/colors                           # get all colors
curl https://tylegroup.com/colors/5                         # get color with id 5
curl -X POST https://tylegroup.com/colors       \
     -H "Content-type: application/json" -d     \
     '{ "name": "cyan","hexadecimal":"00FFFF"}'             # add new color with hexadecimal
curl -X POST https://tylegroup.com/colors       \
     -H "Content-type: application/json" -d     \
     '{ "name": "cyan","r":0,"g":255,"b":255}'              # add new color with rgb
curl -X DELETE https://tylegroup.com/colors/5               # delete color with id 5

```

## License
This package contains code licensed under the Apache License, Version 2.0 (the "License"). You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0 and may also view the License in the LICENSE file within this package.
