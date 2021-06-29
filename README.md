# hkimmd-scrapper
scrapping airport passenger traffic statistics on https://www.immd.gov.hk/

## How to build
```bash
go build .
```

## How to run
```bash
./hkimmd-scrapper -b "2021-05-01" -e "2021-05-07" -f "test.csv"
```

## Flags

|name|description|example|
|---|---|---|
|b|begin date|`2021-05-01`|
|e|end date|`2021-05-07`|
|f|(optional) csv file name, default is `data.csv`|`data.csv`|