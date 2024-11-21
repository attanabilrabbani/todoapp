# Todo-App
ToDo web app

![image](https://github.com/user-attachments/assets/38cb9bef-815a-4acc-bf99-020977c99f22)


go version = 1.22.2
db = postgres

Before Running:
- Create a database in posrgresql
- Navigate to pkg/config/db.go

addr := "user={username} password={password} dbname={dbname} sslmode=disable"

- replace parameters within curly brackets:
- {username} = your db username (usually postgres)
- {password} = your db password
- {dbname} = name of the database you created previously in postgresql

- go to terminal, type cd cmd
- type go run main.go

- run html files.


