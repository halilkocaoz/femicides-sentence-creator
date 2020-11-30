# femicides-sentence-creator
It uses https://femicidesinturkey.com API system, to create sentence about femicide victims.

It takes response like following from the API,

```json
{
    "fullName": "Nuray Altınay",
    "city": "Mersin",
    "killer": {
        "definition": "Groom",
        "status": "Suicide"
    },
    "methods": [
        {
            "method": "Firearm"
        }
    ],
    "causes": [
        {
            "cause": "Protecting her daughter"
        }
    ],
    "year": "2020",
}
```

and It creates sentence like that is 'Nuray Altınay was murdered by her groom while protecting her daughter, in 2020, Mersin.'

Look at twitter: [Femicides In Turkey Twitter](https://twitter.com/FemicidesTurkey)
