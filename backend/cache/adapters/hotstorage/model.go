package hotstorage

func (c *Cache) Update(store string, id string, info any) error {
	return c.i.Upsert(store, []any{id, info}, []any{[]any{"=", 1, info}})
}

func (c *Cache) Get(store string, id string) ([]any, error) {
	return c.i.Get(store, []any{id})
}

func (c *Cache) Delete(store string, id string) error {
	return c.i.Delete(store, []any{id})
}
