# safer

## Restful API for an open-source "safe-space" platform where users can get heard/helped with their mental/quotidian problems by volunteers

## To use this app you will have to configure a MySql database using the following configs

```
CREATE TABLE users(
    Role ENUM( 'volunteer', 'client','admin' ) Default 'client',
    Id varchar(255), 
	Email varchar(255),
	Password varchar(255),
	FirstName varchar(255),
	LastName varchar(255),
	Phone varchar(255),
	PRIMARY KEY (Id)
);



CREATE TABLE cases(
    Status ENUM( 'open', 'closed','in-progress' ) Default 'open',
    Id varchar(255), 
	AssigneeId varchar(255),
	ReporterId varchar(255),
	PRIMARY KEY (Id),
	FOREIGN KEY (AssigneeId) REFERENCES  users(Id),
	FOREIGN KEY (ReporterId) REFERENCES  users(Id)
);
	
CREATE TABLE messages(
    Id varchar(255), 
	SenderId varchar(255),
	CaseId varchar(255),
	Message varchar(1024),
	Data varchar(1024),
	PRIMARY KEY (Id),
	FOREIGN KEY (SenderId) REFERENCES  users(Id),
	FOREIGN KEY (CaseId) REFERENCES  cases(Id)
);
```
## Env file: create a .env file in the root folder of the project with the following properties

```
PORT=:8080
SECRET_KEY=Your secret key
```

## API usage: sql.NullString should be used like this

```
{
    "assigneeId":{
        "String":"0eb9055b-ff08-49b7-8a29-c7259fbcc767",
        "Valid":true
        }
}
```

## Install all packages and then to start the app run
```
go run cmd/main.go
``` 
in the root folder
