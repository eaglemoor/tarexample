box.cfg{
    listen = 3301,
}

local s = box.schema.space.create('test', {
        id = 512,
        if_not_exists = true,
})
s:create_index('primary', {type = 'HASH', parts = {1, 'uint'}, if_not_exists = true})