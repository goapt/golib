```
	logger.Debug("cccc")
	or
	log := logger.NewLogger(func(c *logger.Config){
	    c.LogName= "log"
	})
	log.Debug("ddd")
```