package main

type Plugin interface {
	getName() string
}

// Get

type beforeGet interface {
	beforeGet(key string)
}

type afterGet interface {
	afterGet(key string, value string)
}

// Set

type beforePost interface {
	beforePost(key, value string)
}

type guardPost interface {
	guardPost(key string, value string) (bool, string, int)
}

type afterPost interface {
	afterPost(key, value string)
}

// Put

type beforePut interface {
	beforePut(key, value string)
}

type guardPut interface {
	guardPut(key string, oldValue string, newValue string) (bool, string, int)
}

type afterPut interface {
	afterPut(key, oldValue string, newValue string)
}

// Delete

type beforeDelete interface {
	beforeDelete(key string, value string)
}

type afterDelete interface {
	afterDelete(key string, value string)
}
