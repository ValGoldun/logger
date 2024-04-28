```go
logger, err := logger.New(Info)
if err != nil {
	return err
}

logger.Info("info log", 
	Field{
	    Key: "some field", 
	    Value: "some field value",
	},
)
```