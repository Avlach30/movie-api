# Movie API     
Simple repository for Movie API with Go      
## Tech stack   
 1. Gin-Gonic
 2. GORM
 3. MySql
 3. Bcrypt
 4. JWT

## How to Migrate Database
 1. Type command in terminal  
 migrate create -ext sql -dir db/migration -seq nameMigrationProcess  
you can change 'nameMigrationProcess' as your desired 
 2. Insert All SQL query for updating database as your desired in file nameMigrationProcess.up.sql
 3. You can run up migration for update database structure with command in terminal  
migrate -path db/migration -database "mysql://mysql-user:mysql-password@tcp(localhost:3306)/database-name" -verbose up  
 mysql-user, mysql-password, and database-name is customized from your configuration