# periodic-table-api
API to fetch elements of the periodic table in JSON format. Developed using Golang for faster execution 

## Overview
The following document will specify how to use the API to fetch the periodic elements in JSON. Also it will state different methods throught which the elements can be fetch.

|     **Data Object**    |
|------------------------|
| symbol                 |
| name                   |
| atomicMass             |
| atomicNumber           |
| atomicRadius           |
| boilingPoint           |
| bondingType            |
| cpkHexColor            |
| density                |
| electronAffinity       |
| electronegativity      |
| electronicConfiguration|
| groupBlock             |
| ionRadius              |
| ionizationEnergy       |
| meltingPoint           |
| oxidationStates        |
| standardState          |
| vanDerWaalsRadius      |
| yearDiscovered         |


## Methods
There are total of 6 methods by which you can fetch the data :

### All

- [https://periodic-table-api.herokuapp.com](https://periodic-table-api.herokuapp.com)

This will fetch all the 118 elements from periodic table.

### Atomic Number


- [https://periodic-table-api.herokuapp.com/atomicNumber/20](https://periodic-table-api.herokuapp.com/atomicNumber/20)

This will fetch element from periodic table having atomic number 20. Replace 20 with any other atomic number to fetch that element from 118.

### Atomic Name


- [https://periodic-table-api.herokuapp.com/atomicName/Mercury](https://periodic-table-api.herokuapp.com/atomicName/Mercury)

This will fetch element from periodic table having atomic name "Mercury". Replace "Mercury" with any other atomic name to fetch that element.

### Atomic Symbol


- [https://periodic-table-api.herokuapp.com/atomicSymbol/H](https://periodic-table-api.herokuapp.com/atomicSymbol/H)

This will fetch element from periodic table having atomic symbol "H" i.e. Hydrogen. Replace "H" with any other atomic symbol to fetch that element.

### Bonding Type


- [https://periodic-table-api.herokuapp.com/atomicBonding/Metallic](https://periodic-table-api.herokuapp.com/atomicBonding/Metallic)

This will fetch all elements from periodic table having  Metallic bonding. Replace metallic with any other bonding type to fetch elements.

### Group Block


- [https://periodic-table-api.herokuapp.com/atomicGroup/Metal](https://periodic-table-api.herokuapp.com/atomicGroup/Metal)

This will fetch all elements from periodic table belongs to metal group. Replace metal with any other bonding type to fetch elements.

### State


- [https://periodic-table-api.herokuapp.com/atomicState/gas](https://periodic-table-api.herokuapp.com/atomicState/gas)

This will fetch all elements from periodic table belongs to gas state. Replace gas with any other state to fetch elements.


