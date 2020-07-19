Block Scout
======

Block Scout is an Ethereum block explorer providing realtime block and transaction information.

Live site: https://guarded-plains-39628.herokuapp.com


### Technologies


- Frontend: React
- Backend: Golang, Gorilla WebSockets, Geth
- Database and Storage: MongoDB
- Hosting: Heroku


### Latest Blocks and Transactions

The home page shows the latest 4 blocks and some recently confirmed transactions. To get the latest block information, Geth is used to setup a websocket subscription to an infura node which broadcasts the latest blocks to the connection. These blocks along with their transactions are saved to the DB, and then broadcasted to all the clients via a seperate websocket connection.  

