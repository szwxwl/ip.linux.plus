FROM openresty/openresty:latest
COPY ./nginx/html /usr/local/openresty/html
COPY ./nginx/lua/src /usr/local/openresty/nginx/lua_src
COPY ./nginx/lua/lib/ip2location.lua /usr/local/openresty/lualib/ip2location.lua
COPY ./nginx/lua/lib/inspect.lua /usr/local/openresty/lualib/inspect.lua
COPY ./nginx/lua/lib/nums/bn.lua /usr/local/openresty/lualib/nums/bn.lua
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY ./nginx/ip_database /usr/local/openresty/nginx/ip_database
CMD ["/usr/bin/openresty", "-g", "daemon off;"]
