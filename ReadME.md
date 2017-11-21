With this API you can consult all transactions of one SKU, and define what currency you want for his all Tip calculation.
This Code generate a backup to consult if you can't connect or find the SKU, in this case, i use a file backup becouse is a simple example.
This API is based on a generic selection process task.


- This API is used to define a Tip Calculator.
- You have 3 EndPoint:
	- /api/rates to consume the existing Endpoint in appsettings.json.
	- /api/Transactions to consume the existing Endpoint in appsettings.json.
	- /api/{sku}/{currency} to consume the calculator Tip.
- The code is based on SOLID Go Structured.
- In this case i don't develop Unit testing, but i usually do that.


*"myCalculator" run in Linux