box.cfg{
    listen = 3301,
}

local s = box.schema.space.create('test')
s:create_index('primary', {type = 'HASH'})