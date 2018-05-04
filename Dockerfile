FROM tarantool:1.10.1

COPY app.lua /opt/tarantool
CMD ["tarantool", "/opt/tarantool/app.lua"]