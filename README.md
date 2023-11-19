# Rulers Compendium
[![Build, Tests & Docs](https://github.com/grqphical07/rulers-compendium/actions/workflows/go.yml/badge.svg)](https://github.com/grqphical07/rulers-compendium/actions/workflows/go.yml)

A free to use database API for information about Sid Meier's civilization VI. I gathered data by scraping the civ6 wiki so some of it may be inaccurate. If it is feel free to submit an issue detailing the error.

## Contents

- [Self-Hosting Usage](#self-hosting-usage)
- [Docs](#docs)
- [Basic API Usage](#basic-api-usage)
    - [Example](#example)
- [Roadmap](#roadmap)
- [License](#license)

## Self-Hosting Usage

1. Install go

2. Run ```$ go run .``` to start the API

## Docs

API Documentation is available at ```/api/v1/docs/index.html```

## Basic API Usage

The API is split into different categories each pertaining to a subject area in Civilization VI. Currently there are three categories:

- Districts
- Leaders
- Civilizations

Each one has two endpoints: one for getting all available items and one for finding specific items

### Example:

```bash
$ curl -X GET "http://127.0.0.1:8000/api/v1/leaders?limit=2"
```

Returns:
```json
[
  {
    "name":"Abraham Lincoln",
    "civ":"America",
    "ability":{
      "name":"Emancipation Proclamation",
      "text":"Industrial Zones grant +2  Amenities and +3 Loyalty per turn but Plantations give -2 Loyalty. Receives a free Melee unit after constructing Industrial Zones and their buildings. The free unit does not require resources when created or to maintain and receives +5  Combat Strength."
    },
    "agenda":{
      "name":"Preserver of the Union",
      "text":"Likes civilizations with the same form of government, dislikes those with a different form of government and really dislikes ones with a different government of the same era as his own."
    }
  },
  {
    "name":"Alexander",
    "civ":"Macedonia",
    "ability":{
      "name":"To the World's End",
      "text":"Macedonian cities never incur war weariness. All military units heal completely when a city with a Wonder is captured. Gains the Hetairoi unique unit with Horseback Riding."
    },
    "agenda":{
      "name":"Short Life of Glory",
      "text":"Likes civilizations at war with powers other than Macedon. Dislikes civilizations at peace.  Grievances against this leader decay at twice the usual rate."
    }
  }
]
```

## Roadmap

- [x] Districts
- [x] Civilizations
- [x] Leaders
- [ ] Units
- [ ] Buildings
- [ ] Wonders
- [ ] Improvements
- [ ] Great People
- [ ] Governors
- [ ] DLC Filters
- [ ] General Database Search

## License

Ruler's Compendium is Licensed under the [MIT License](https://github.com/grqphical07/rulers-compendium/blob/main/LICENSE)