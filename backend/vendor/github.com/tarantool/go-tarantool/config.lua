-- Do not set listen for now so connector won't be
-- able to send requests until everything is configured.
box.cfg{
    work_dir = os.getenv("TEST_TNT_WORK_DIR"),
}

box.once("init", function()
    local s = box.schema.space.create('test', {
        id = 517,
        if_not_exists = true,
    })
    s:create_index('primary', {type = 'tree', parts = {1, 'uint'}, if_not_exists = true})

    local sp = box.schema.space.create('SQL_TEST', {
        id = 519,
        if_not_exists = true,
        format = {
            {name = "NAME0", type = "unsigned"},
            {name = "NAME1", type = "string"},
            {name = "NAME2", type = "string"},
        }
    })
    sp:create_index('primary', {type = 'tree', parts = {1, 'uint'}, if_not_exists = true})
    sp:insert{1, "test", "test"}

    local st = box.schema.space.create('schematest', {
        id = 516,
        temporary = true,
        if_not_exists = true,
        field_count = 7,
        format = {
            {name = "name0", type = "unsigned"},
            {name = "name1", type = "unsigned"},
            {name = "name2", type = "string"},
            {name = "name3", type = "unsigned"},
            {name = "name4", type = "unsigned"},
            {name = "name5", type = "string"},
        },
    })
    st:create_index('primary', {
        type = 'hash',
        parts = {1, 'uint'},
        unique = true,
        if_not_exists = true,
    })
    st:create_index('secondary', {
        id = 3,
        type = 'tree',
        unique = false,
        parts = { 2, 'uint', 3, 'string' },
        if_not_exists = true,
    })
    st:truncate()

    local s2 = box.schema.space.create('test_perf', {
        id = 520,
        temporary = true,
        if_not_exists = true,
        field_count = 3,
        format = {
            {name = "id", type = "unsigned"},
            {name = "name", type = "string"},
            {name = "arr1", type = "array"},
        },
    })
    s2:create_index('primary', {type = 'tree', unique = true, parts = {1, 'unsigned'}, if_not_exists = true})
    s2:create_index('secondary', {id = 5, type = 'tree', unique = false, parts = {2, 'string'}, if_not_exists = true})
    local arr_data = {}
    for i = 1,100 do
        arr_data[i] = i
    end
    for i = 1,1000 do
        s2:insert{
            i,
            'test_name',
            arr_data,
        }
    end

    --box.schema.user.grant('guest', 'read,write,execute', 'universe')
    box.schema.func.create('box.info')
    box.schema.func.create('simple_incr')

    -- auth testing: access control
    box.schema.user.create('test', {password = 'test'})
    box.schema.user.grant('test', 'execute', 'universe')
    box.schema.user.grant('test', 'read,write', 'space', 'test')
    box.schema.user.grant('test', 'read,write', 'space', 'schematest')
    box.schema.user.grant('test', 'read,write', 'space', 'test_perf')

    -- grants for sql tests
    box.schema.user.grant('test', 'create,read,write,drop,alter', 'space')
    box.schema.user.grant('test', 'create', 'sequence')
end)

local function func_name()
    return {
        {221, "", {
                {"Moscow", 34},
                {"Minsk", 23},
                {"Kiev", 31},
            }
        }
    }
end
rawset(_G, 'func_name', func_name)

local function simple_incr(a)
    return a + 1
end
rawset(_G, 'simple_incr', simple_incr)

box.space.test:truncate()

--box.schema.user.revoke('guest', 'read,write,execute', 'universe')

-- Set listen only when every other thing is configured.
box.cfg{
    listen = os.getenv("TEST_TNT_LISTEN"),
}
