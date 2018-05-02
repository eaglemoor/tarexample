# tarantool v 1.9
FROM tarantool/tarantool:1

# tarantool v 1.7
#FROM tarantool/tarantool:1.7

# tarantool v 2
#FROM tarantool/tarantool:2.0

COPY app.lua /opt/tarantool
CMD ["tarantool", "/opt/tarantool/app.lua"]