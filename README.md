## Go Microservice to fetch and serve for Frontend Go Projects..

```bash
[musti@musti-deputyp25 real-currency]$ go run main.go 
2026/10/27 14:49:18 Server started on port :8080

```

```bash
[musti@musti-deputyp25 real-currency]$ curl -s localhost:8080|jq
{
  "Time": "2026-10-27T14:49:42.756507873+01:00",
  "Status": "200 OK",
  "PrimeAssets": {
    "Altın Gümüş": {
      "name": "Altın Gümüş",
      "price": "74,7500",
      "change": "%0,43"
    },
    "Dolar": {
      "name": "Dolar",
      "price": "49,7930",
      "change": "%-0,01"
    },
    "Euro": {
      "name": "Euro",
      "price": "55,5774",
      "change": "%-0,24"
    },
    "Euro Dolar": {
      "name": "Euro Dolar",
      "price": "1,1145",
      "change": "%-0,27"
    },
    "Gram Altın": {
      "name": "Gram Altın",
      "price": "6.442,44",
      "change": "%0,26"
    },
    "Gümüş Gram": {
      "name": "Gümüş Gram",
      "price": "86,2917",
      "change": "%-0,09"
    },
    "ONS Altın": {
      "name": "ONS Altın",
      "price": "4.025,41",
      "change": "%0,20"
    },
    "İngiliz Sterlini": {
      "name": "İngiliz Sterlini",
      "price": "66,4044",
      "change": "%-0,30"
    }
  },
  "Others": {
    "Altın Gümüş": {
      "name": "Altın Gümüş",
      "price": "74,7500",
      "change": "%0,43"
    },
    "BHD": {
      "name": "BHD",
      "price": "134,45",
      "change": "%1,59"
    },
    "BNB": {
      "name": "BNB",
      "price": "742,02",
      "change": "%0,25"
    },
    "BRL": {
      "name": "BRL",
      "price": "9,3585",
      "change": "%1,34"
    },
    "BTC": {
      "name": "BTC",
      "price": "80.179,66",
      "change": "%0,42"
    },
    "Bahreyn Dinarı": {
      "name": "Bahreyn Dinarı",
      "price": "134,45",
      "change": "%1,59"
    },
    "Bosna-Hersek Markı": {
      "name": "Bosna-Hersek Markı",
      "price": "28,3876",
      "change": "%-0,56"
    },
    "Brezilya Reali": {
      "name": "Brezilya Reali",
      "price": "9,3585",
      "change": "%1,34"
    },
    "CAD": {
      "name": "CAD",
      "price": "35,5006",
      "change": "%-0,34"
    },
    "CHF": {
      "name": "CHF",
      "price": "60,0858",
      "change": "%-0,36"
    },
    "CORN": {
      "name": "CORN",
      "price": "451,25",
      "change": "%-0,82"
    },
    "COTTON": {
      "name": "COTTON",
      "price": "78,6900",
      "change": "%-0,29"
    },
    "DXY": {
      "name": "DXY",
      "price": "102,23",
      "change": "%0,21"
    },
    "Danimarka Kronu": {
      "name": "Danimarka Kronu",
      "price": "7,4263",
      "change": "%-0,49"
    },
    "Dolar": {
      "name": "Dolar",
      "price": "49,7930",
      "change": "%-0,01"
    },
    "ETH": {
      "name": "ETH",
      "price": "2.324,61",
      "change": "%0,95"
    },
    "EUR": {
      "name": "EUR",
      "price": "55,5774",
      "change": "%-0,24"
    },
    "EUR/USD": {
      "name": "EUR/USD",
      "price": "1,1145",
      "change": "%-0,27"
    },
    "Euro": {
      "name": "Euro",
      "price": "55,5774",
      "change": "%-0,24"
    },
    "Euro Dolar": {
      "name": "Euro Dolar",
      "price": "1,1145",
      "change": "%-0,27"
    },
    "GAU": {
      "name": "GAU",
      "price": "6.442,44",
      "change": "%0,26"
    },
    "GBP": {
      "name": "GBP",
      "price": "66,4044",
      "change": "%-0,30"
    },
    "GBP/EUR": {
      "name": "GBP/EUR",
      "price": "1,1951",
      "change": "%0,07"
    },
    "GBP/USD": {
      "name": "GBP/USD",
      "price": "1,3318",
      "change": "%-0,18"
    },
    "GRPLTN": {
      "name": "GRPLTN",
      "price": "2.440,71",
      "change": "%-0,55"
    },
    "Gram Altın": {
      "name": "Gram Altın",
      "price": "6.442,44",
      "change": "%0,26"
    },
    "Gümüş Gram": {
      "name": "Gümüş Gram",
      "price": "86,2917",
      "change": "%-0,09"
    },
    "Has Altın": {
      "name": "Has Altın",
      "price": "6.410,23",
      "change": "%0,31"
    },
    "ILS": {
      "name": "ILS",
      "price": "16,3631",
      "change": "%1,25"
    },
    "ISCTR": {
      "name": "ISCTR",
      "price": "12,2400",
      "change": "%-0,24"
    },
    "KCHOL": {
      "name": "KCHOL",
      "price": "200,60",
      "change": "%-0,05"
    },
    "Kanada Doları": {
      "name": "Kanada Doları",
      "price": "35,5006",
      "change": "%-0,34"
    },
    "Katar Riyali": {
      "name": "Katar Riyali",
      "price": "14,1893",
      "change": "%3,77"
    },
    "Kazak Tengesi": {
      "name": "Kazak Tengesi",
      "price": "0,1093",
      "change": "%-1,23"
    },
    "Meksika Pesosu": {
      "name": "Meksika Pesosu",
      "price": "2,6699",
      "change": "%-0,67"
    },
    "ONS Altın": {
      "name": "ONS Altın",
      "price": "4.025,41",
      "change": "%0,20"
    },
    "PKR": {
      "name": "PKR",
      "price": "0,1859",
      "change": "%3,22"
    },
    "Pakistan Rupisi": {
      "name": "Pakistan Rupisi",
      "price": "0,1859",
      "change": "%3,22"
    },
    "Polonya Zlotisi": {
      "name": "Polonya Zlotisi",
      "price": "12,6103",
      "change": "%-0,44"
    },
    "QAR": {
      "name": "QAR",
      "price": "14,1893",
      "change": "%3,77"
    },
    "RUB": {
      "name": "RUB",
      "price": "0,6198",
      "change": "%0,93"
    },
    "Rus Rublesi": {
      "name": "Rus Rublesi",
      "price": "0,6198",
      "change": "%0,93"
    },
    "SAHOL": {
      "name": "SAHOL",
      "price": "86,4000",
      "change": "%-1,26"
    },
    "SOL": {
      "name": "SOL",
      "price": "100,38",
      "change": "%0,17"
    },
    "SPX": {
      "name": "SPX",
      "price": "7.694,36",
      "change": "%0,23"
    },
    "SUGAR": {
      "name": "SUGAR",
      "price": "19,1200",
      "change": "%0,10"
    },
    "THYAO": {
      "name": "THYAO",
      "price": "255,25",
      "change": "%-1,16"
    },
    "TUPRS": {
      "name": "TUPRS",
      "price": "360,50",
      "change": "%-1,10"
    },
    "TWD": {
      "name": "TWD",
      "price": "1,5353",
      "change": "%0,45"
    },
    "Tam Altın": {
      "name": "Tam Altın",
      "price": "42.281,66",
      "change": "%0,21"
    },
    "USD": {
      "name": "USD",
      "price": "49,7930",
      "change": "%-0,01"
    },
    "USD/JPY": {
      "name": "USD/JPY",
      "price": "157,86",
      "change": "%0,23"
    },
    "USD/RUB": {
      "name": "USD/RUB",
      "price": "80,7839",
      "change": "%-0,46"
    },
    "USDT": {
      "name": "USDT",
      "price": "0,9984",
      "change": "%0,01"
    },
    "XBR/USD": {
      "name": "XBR/USD",
      "price": "110,37",
      "change": "%0,13"
    },
    "XPT/USD": {
      "name": "XPT/USD",
      "price": "1.523,91",
      "change": "%-0,58"
    },
    "XRP": {
      "name": "XRP",
      "price": "1,3144",
      "change": "%2,03"
    },
    "XU030": {
      "name": "XU030",
      "price": "13.994,62",
      "change": "%0,33"
    },
    "XU050": {
      "name": "XU050",
      "price": "10.498,62",
      "change": "%0,32"
    },
    "XU100": {
      "name": "XU100",
      "price": "11.191,82",
      "change": "%0,36"
    },
    "YKBNK": {
      "name": "YKBNK",
      "price": "35,4800",
      "change": "%-1,00"
    },
    "Yarım Altın": {
      "name": "Yarım Altın",
      "price": "21.237,51",
      "change": "%0,21"
    },
    "Yeni Tayvan Doları": {
      "name": "Yeni Tayvan Doları",
      "price": "1,5353",
      "change": "%0,45"
    },
    "Çeyrek Altın": {
      "name": "Çeyrek Altın",
      "price": "10.602,64",
      "change": "%0,21"
    },
    "İngiliz Sterlini": {
      "name": "İngiliz Sterlini",
      "price": "66,4044",
      "change": "%-0,30"
    },
    "İsrail Şekeli": {
      "name": "İsrail Şekeli",
      "price": "16,3631",
      "change": "%1,25"
    },
    "İsviçre Frangı": {
      "name": "İsviçre Frangı",
      "price": "60,0858",
      "change": "%-0,36"
    }
  },
  "Main": {
    "Altın Gümüş": {
      "name": "Altın Gümüş",
      "price": "74,7500",
      "change": "%0,43"
    },
    "BIST 100": {
      "name": "BIST 100",
      "price": "11.191,82",
      "change": "%0,36"
    },
    "Bitcoin": {
      "name": "Bitcoin",
      "price": "80.179,66",
      "change": "%0,42"
    },
    "Brent Petrol": {
      "name": "Brent Petrol",
      "price": "110,37",
      "change": "%0,13"
    },
    "Gümüş Gram": {
      "name": "Gümüş Gram",
      "price": "86,2917",
      "change": "%-0,09"
    },
    "ONS Altın": {
      "name": "ONS Altın",
      "price": "4.025,41",
      "change": "%0,20"
    },
    "Sepet Kur": {
      "name": "Sepet Kur",
      "price": "52,6835",
      "change": "%-0,13"
    }
  }
}
```
